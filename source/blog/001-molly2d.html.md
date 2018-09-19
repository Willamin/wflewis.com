---
title: Molly2d
date: 2018-09-19
---

Ever since I can remember, I've enjoyed making games. When I was younger this mostly meant changing the rules of basketball to something else. As I got older, I started to learn how to make computer games, starting with GameMaker. GameMaker provided a drag and drop way to describe game logic along with a simple programming language that could be used instead. While I focused on other aspects of programming, I still kept my love for making games.

Later, I came across the game framework Löve2d. It's a small application that wraps Lua scripts, while providing a simple API for drawing things, playing sounds, etc. Not to make a pun out of the name, but I loved it. I was able to write games using only code and without a lot of boilerplate! I soon found it's weakness - objects. I'll leave arguments about the virtues and detriments of Object Oriented Programming in general programming out of it, but for video games, I hold a strong opinion that OOP is one of the best ways to structure code. Entities that interact with one another forms the core of video games. These entities map exceedingly well to objects.

One of the core aspects of Lua is that "everything is a table". Tables in Lua are a list of mappings from one thing to another. To group data together in Lua, tables are used. To group functions together in Lua, tables are used. To make lists of things, tables are used. I found that using tables like this to build objects was troublesome. I could make tables that held all the appropriate data and functions that make up an object, but it felt cumbersome. If you know me, you know that I am not a fan of cumbersome programming. (That should be an entire article itself)

I did my best to ease the process of defining new objects and constructing inheritance trees using Lua's prototypical inheritance, but it always felt like I was fighting with Lua itself, when it was supposed to be my tool. Coincidentally, I was beginning to explore the new Crystal language. Crystal claimed to have beautiful syntax like Ruby while keeping the speed and efficiency of languages like C. While Crystal still lacks some important features, it has definitely met and exceeded my expectations.

So I set about building a lightweight game engine in Crystal, my goals simple:  stay idiomatically Crystal, at least from a user's perspective be uncomplicated, and get out of the way when possible. If my API ever began to felt complicated, I've looked to Löve2d as a baseline. To make sure it gets out of the way, I've been working on a simple game in conjunction with my engine.

While it's not quite finished, fully featured, or meticulously documented, I present [Molly2d](https://github.com/willamin/molly2d).

![](/images/blog/001/molly2d.png)
