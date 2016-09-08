pyxel2gm
========

Tools to streamline use of [Pyxel Edit](http://pyxeledit.com/) with [Game Maker Studio](http://www.yoyogames.com/gamemaker).

Installation
------------

For now, you are required to have the [Go programming language](https://golang.org/) installed.
(Eventually I will probably provide pre-compiled binaries somewhere.)

```
C:\> go get github.com/hurricanerix/pyxel2gm/...
```

pyxel2gm-export
---------------

Exports tilesets located in .pyxel files as sprites in a Game Maker Studio
project directory.

Setup:
```
Example.gmx/assets/sprites/foo/spr_bar.pyxel
Example.gmx/assets/sprites/spr_baz.pyxel
Example.gmx/sprites/images/spr_bar_0.png
Example.gmx/sprites/images/spr_bar_1.png
...
Example.gmx/sprites/images/spr_bar_N.png
Example.gmx/sprites/images/spr_baz_0.png
Example.gmx/sprites/images/spr_baz_1.png
...
Example.gmx/sprites/images/spr_baz_N.png
```
When run, tiles 0-N in spr_bar.pyxel and spr_baz.pyxel will be copied to Example.gmx/sprites/images/ directory as "spr_bar_0.png", "spr_bar_1.png", ...


pyxel2gm-editor
----------------

This utility maps sprites back to their corresponding .pyxel files and opens
them with Pyxel Edit.  This allows you to have your external editor in Game
Maker Studio open Pyxel Edit with the .pyxel file instead of the exported png.
Once the instance of Pyxel Edit launched by the utility closes, it will export
the tiles in the opened .pyxel file into the projects "sprites/images" directory.
