run_mlr -n --ojson put '
  end {
    a = [1,2,3,4,5,6,7];
    m = {"a": 1, "b": 2};
    s = "abcdefg";

    emit {
      "a35": a[3:5],
      "a37": a[3:7],
      "a17": a[1:7],
      "a07": a[0:7],
      "a39": a[3:9],
      "a53": a[5:3],
      "a93": a[9:3],
    };

    emit {
      "m11": m[1:1],
    };

    emit {
      "s35": s[3:5],
      "s37": s[3:7],
      "s17": s[1:7],
      "s07": s[0:7],
      "s39": s[3:9],
      "s53": s[5:3],
      "s93": s[9:3],
    };

    emit {
      "u35": substr(s,3,5),
      "u37": substr(s,3,7),
      "u17": substr(s,1,7),
      "u07": substr(s,0,7),
      "u39": substr(s,3,9),
      "u53": substr(s,5,3),
      "u93": substr(s,9,3),
    };
  }
'
