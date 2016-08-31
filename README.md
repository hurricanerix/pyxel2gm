pyxel2gm
========

Utility for importing Pyxel Edit files into Game Maker.

Usage
-----

```
C:\>pyxel2gm -h
Usage of pyxel2gm:
  -assets-dir string
        directory to scan for .pyxel files (default "assets")
  -dry
        display report of files to be created
  -ignore-gmx
        don't create or modify .gmx files
  -project-dir string
        Game Maker project directory to export into (default ".")
  -version
        display the version of this app
exit status 2
```

Tech Details
------------

.pyxel files are really just zip files with a specific format:

```
file.pyxel
|- docData.json
|- layer0.png
|- layer1.png
|- ...
|- layerN.png
|- tile0.png
|- tile1.png
|- ...
|- tileN.png
```

docData.json contains the preferred exportName for tiles.

Assuming the Pyxel files are under a directory "assets" and the Game Maker project is in a directory "project"

This app will treat each Pyxel file as a sprite in the context of game maker.  So for every Pyxel file, there should be a corresponding GMX file, with a number of corresponding PNGs.

This application does two things.

1. It creates/modifies the .gmx file for each sprite.
"assets/*.pyxel" -> "project/sprites/{exportName}.gmx"

2. It creates/modifies the image data for the GMX file.
"assets/*.pyxel/tile{N}.png" -> "project/sprites/images/{exportName}_{N}.png"
