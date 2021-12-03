---
title: Format One point one
date: 2021-06-16
---

In considering the format, I think a small adjustment should be made to improve extensibility. Of course, like any extensibility improvement there comes the cost of potentially adding unnecessary complexity. For now, this change feels worth it.

Another small change is a default celltype. If the cell sigil is present, _alone_ on it's line, then this line can be considered to be a plaintext cell header.

Like many other formats, Birch Version One will support a header key/value store.

The syntax parsing should work as below.

1. Find the first whitespace in the document. All characters before this are considered the Cell Sigil for the document.
2. Parse the document line-by-line. When finding a line that begins with the Cell Sigil, consider this the Header Line of a new cell.
3. When parsing a Header Line, expect a Cell Sigil. If there is nothing else on the line, consider the cell a plaintext type. If more is on the line, expect whitespace, then a Cell Type. 
4. After the Cell Type is optional header key/value pairs. Separating the Cell Type from these is whitespace. Both the keys and values can be quoteless words or quoted strings. Expect a Header Key, optional whitespace, an Equals sign, optional whitespace, then finally a Header Value. Additional whitespace separates this Header Value from the next Header Key if one exists.
