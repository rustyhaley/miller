package stream

import (
	"errors"
	"fmt"
	"os"

	"miller/cliutil"
	"miller/input"
	"miller/lib"
	"miller/output"
	"miller/transforming"
	"miller/types"
)

// Since Go is concurrent, the context struct (AWK-like variables such as
// FILENAME, NF, NF, FNR, etc.) needs to be duplicated and passed through the
// channels along with each record.
//
// * Record-readers update FILENAME, FILENUM, NF, NR, FNR within context structs.
//
// * Record-transformers can read these from the context structs.
//
// * Record-writers don't need them (OPS et al. are already in the
//   writer-options struct). However, we have chained transformers using the
//   'then' command-line syntax. This means a given transformer might be piping
//   its output to a record-writer, or another transformer. So, the
//   record-and-context pair goes to the record-writers even though they don't
//   need the contexts.

func Stream(
	options cliutil.TOptions,
	recordTransformers []transforming.IRecordTransformer,
) error {

	// Since Go is concurrent, the context struct needs to be duplicated and
	// passed through the channels along with each record.
	initialContext := types.NewContext(&options)

	// Instantiate the record-reader
	recordReader := input.Create(&options.ReaderOptions)
	if recordReader == nil {
		return errors.New("Input format not found: " + options.ReaderOptions.InputFileFormat)
	}

	// Instantiate the record-writer
	recordWriter := output.Create(&options.WriterOptions)
	if recordWriter == nil {
		return errors.New("Output format not found: " + options.WriterOptions.OutputFileFormat)
	}

	// Set up the reader-to-transformer and transformer-to-writer channels.
	inputChannel := make(chan *types.RecordAndContext, 10)
	outputChannel := make(chan *types.RecordAndContext, 1)

	// We're done when a fatal error is registered on input (file not found,
	// etc) or when the record-writer has written all its output. We use
	// channels to communicate both of these conditions.
	errorChannel := make(chan error, 1)
	doneChannel := make(chan bool, 1)

	// Start the reader, transformer, and writer. Let them run until fatal input
	// error or end-of-processing happens.

	go recordReader.Read(options.FileNames, *initialContext, inputChannel, errorChannel)
	go transforming.ChainTransformer(inputChannel, recordTransformers, outputChannel)
	go output.ChannelWriter(outputChannel, recordWriter, doneChannel, os.Stdout)

	done := false
	for !done {
		select {
		case err := <-errorChannel:
			fmt.Fprintln(os.Stderr, lib.MlrExeName(), ": ", err)
			os.Exit(1)
		case _ = <-doneChannel:
			done = true
			break
		}
	}

	return nil
}
