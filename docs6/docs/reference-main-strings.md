<!---  PLEASE DO NOT EDIT DIRECTLY. EDIT THE .md.in FILE PLEASE. --->
<div>
<span class="quicklinks">
Quick links:
&nbsp;
<a class="quicklink" href="../reference-verbs/index.html">Verb list</a>
&nbsp;
<a class="quicklink" href="../reference-dsl-builtin-functions/index.html">Function list</a>
&nbsp;
<a class="quicklink" href="../glossary/index.html">Glossary</a>
&nbsp;
<a class="quicklink" href="https://github.com/johnkerl/miller" target="_blank">Repository ↗</a>
</span>
</div>
# Strings

TODO

xxx concat

xxx index and slice 1-up

xxx lib functions

always double-quote

single-quote for shell; see windows page

dot operator ...

## Escape sequences for string literals

You can use the following backslash escapes for strings such as between the double quotes in contexts such as `mlr filter '$name =~ "..."'`, `mlr put '$name = $othername . "..."'`, `mlr put '$name = sub($name, "...", "...")`, etc.:

* `\a`: ASCII code 0x07 (alarm/bell)
* `\b`: ASCII code 0x08 (backspace)
* `\f`: ASCII code 0x0c (formfeed)
* `\n`: ASCII code 0x0a (LF/linefeed/newline)
* `\r`: ASCII code 0x0d (CR/carriage return)
* `\t`: ASCII code 0x09 (tab)
* `\v`: ASCII code 0x0b (vertical tab)
* `\\`: backslash
* `\"`: double quote
* `\123`: Octal 123, etc. for `\000` up to `\377`
* `\x7f`: Hexadecimal 7f, etc. for `\x00` up to `\xff`

See also [https://en.wikipedia.org/wiki/Escape_sequences_in_C](https://en.wikipedia.org/wiki/Escape_sequences_in_C).

These replacements apply only to strings you key in for the DSL expressions for `filter` and `put`: that is, if you type `\t` in a string literal for a `filter`/`put` expression, it will be turned into a tab character. If you want a backslash followed by a `t`, then please type `\\t`.

However, these replacements are done automatically only for string literals within DSL expressions -- they are not done automatically to fields within your data stream.  If you wish to make these replacements, you can do (for example) `mlr put '$field = gsub($field, "\\t", "\t")'`. If you need to make such a replacement for all fields in your data, you should probably use the system `sed` command instead. 
