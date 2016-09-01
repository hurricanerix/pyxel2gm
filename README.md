pyxel2gm
========

Utility for importing Pyxel Edit tiles into Game Maker.

Tiles 0-N in foo.pyxel will be copied to the project-dir/sprites/images/ directory as "spr_foo_0.png", "spr_foo_1.png", ... "spr_foo_N.png".

Usage
-----

```
C:\>pyxel2gm -h
Usage of pyxel2gm:
  -prefix string
        prefix to add to filenames (default "spr_")
  -assets-dir string
        directory to scan for .pyxel files (default "assets")
  -project-dir string
        Game Maker project directory to export into (default ".")
  -version
        display the version of this app
exit status 2
```
