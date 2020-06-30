# Small Box

[![codecov](https://codecov.io/gh/roger-russel/smallbox/branch/master/graph/badge.svg)](https://codecov.io/gh/roger-russel/smallbox) [![Maintainability](https://api.codeclimate.com/v1/badges/9738317709cabed40cc5/maintainability)](https://codeclimate.com/github/roger-russel/smallbox/maintainability) ![download](https://img.shields.io/github/downloads/roger-russel/smallbox/total.svg)

Small Box is a small library to embed static content into Golang Binary.

By generating go files with it content.

## When use it

The small box is not suitable to high performance and it will increase the size of binary compiled.

But there is situations when it is suitable to use like when:

* Have a static content and as planning to add it into a variable with a huge content.

* Prefer to have everything into the binary instead of download external content, or have install a lib folder somewhere.

## How use

smallbox -f ./foo/boo.tpl

### Limitations

Some limitations found into this project.

#### Folder name must be "box"

Becouse it will try to create a box folder because of how package works in Golang, if a folder with this name already exists then try to create a parent folder like "./autogenerate" to use like this: "./autogenerate/box".

## Images Credits

There is some images that I took to test to add images into static files.

<span>Photo by <a href="https://unsplash.com/@filipz?utm_source=unsplash&amp;utm_medium=referral&amp;utm_content=creditCopyText">Filip Zrnzević</a> on <a href="/?utm_source=unsplash&amp;utm_medium=referral&amp;utm_content=creditCopyText">Unsplash</a></span>

<span>Photo by <a href="https://unsplash.com/@itfeelslikefilm?utm_source=unsplash&amp;utm_medium=referral&amp;utm_content=creditCopyText">🇸🇮 Janko Ferlič</a> on <a href="/?utm_source=unsplash&amp;utm_medium=referral&amp;utm_content=creditCopyText">Unsplash</a></span>

<span>Photo by <a href="https://unsplash.com/@iankeefe?utm_source=unsplash&amp;utm_medium=referral&amp;utm_content=creditCopyText">Ian Keefe</a> on <a href="/?utm_source=unsplash&amp;utm_medium=referral&amp;utm_content=creditCopyText">Unsplash</a></span>

<span>Photo by <a href="https://unsplash.com/@meric?utm_source=unsplash&amp;utm_medium=referral&amp;utm_content=creditCopyText">Meriç Dağlı</a> on <a href="/?utm_source=unsplash&amp;utm_medium=referral&amp;utm_content=creditCopyText">Unsplash</a></span>

