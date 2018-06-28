# Video Link: https://www.udemy.com/learn-how-to-code/

Why choose Go?

Performant - Right up there with C (Go is way easier to code in)
Multiple Core Support
Made to run well in Network environments

## Learning Go

Currently I am learning more about Go, and these are the notes that I am taking as i move along

Thought I would put them on github so I can use them while i work

Important things about Go
Everything in Go is passed "by value" (passed by balue of memory)

The "comma ok" idiom
if seconds, ok := timeZone[tz]; ok {
    return seconds
}

Concurrency vs Parallelism
Concurrency: Doing multiple things at different times (drinking coffee & coding)
Parallelism: Doing many things at exactly the same time

Race condition, where two functions add a number or something to the same variable and they both go at the same time, causing only one of them to register