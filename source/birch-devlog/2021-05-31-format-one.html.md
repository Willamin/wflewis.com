---
title: Format One
date: 2021-05-31
---

I’ve been considering how I want to prioritize the goals of the project. In my imagination, I can see so many aspects that I would love to see realized. I have to ground myself though. I’m only one person and I’ll only be working on this in part of my spare time.

To start with, I want to leverage existing tools where possible, forming a sort of glued-together proof of concept. From there I’ll begin to slowly replace tools with improved self-built versions. I don’t mean that I’ll use Notion and AirTable – rather, I mean that I’ll start with a file format that’s comfortable to edit in a text editor, etc.

## Version One organization
More specifically, for Version One, I’ll organize the project as:
- A Notebook represents an idea, document, page, etc. This is the level that is represented by a single file.
- A Notebook consists of Cells.
- A Cell has a type and a content body.
- The Cell’s type determines how it’s content body is perceived, handled, or displayed.

With the goal of comfortable editing in mind, the format for Version One will be as follows:
- The first line of the document is always the first cell header.
- A cell header begins with a sigil, then 1 or more whitespace characters, then the cell type.
- Beginning on the next line is the cell body.
- The cell body continues until the next cell header, denoted by the sigil.
- The sigil is determined by the first cell header line in the document.

An example:

<pre><code>§ markdown
# A Quick Sample File

§ plaintext
For this file, I’ve chosen the § symbol as the cell sigil. This is an arbitrary choice.

§ crystal
# I can include code snippets

puts "hello world"

§ plaintext
As there is no program to process this file or syntax highlighting, there currently appears no difference in the cell types. But imagine, a specific editor for the file could render cell types appropriately: html viewing for markdown, monospace for plaintext, maybe even a source/output interface for code in supported languages.

In fact, it might make sense to cache the output of code somehow. This has proven difficult to get right as I’ve played around with this format, so it may be something to deal with on the next file format version.
</code></pre>
