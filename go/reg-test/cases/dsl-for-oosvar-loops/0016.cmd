mlr --from reg-test/input/xyz2 put -q 'func f() { return {"a"."b":"c"."d",3:4}}; for (k,v in f()){print "k=".k.",v=".v}'
