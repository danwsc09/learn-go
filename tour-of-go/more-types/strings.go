package main

import (
	"fmt"
	"os"
	"strings"
)

/*
var lyrics string = `Yeah!
And if your heart stops beating
I'll be here wondering
Did you get what you deserve?
The ending of your life
And if you get to heaven
I'll be here waiting, babe
Did you get what you deserve?
The end, and if your life won't wait
Then your heart can't take this
Have you heard the news that you're dead?
No one ever had much nice to say, I
Think they never liked you anyway
Oh, take me from the hospital bed
Wouldn't it be grand? It ain't exactly what you planned
And wouldn't it be great if we were dead?
Oh, dead
Tongue-tied and oh so squeamish
You never fell in love
Did you get what you deserve?
The ending of your life
And if you get to heaven
I'll be here waiting, babe
Did you get what you deserve?
The end, and if your life won't wait
Then your heart can't take this
Have you heard the news that you're dead?
No one ever had much nice to say, I
Think they never liked you anyway
Oh, take me from the hospital bed
Wouldn't it be grand to take a pistol by the hand?
And wouldn't it be great if we were dead?
And in my honest observation
During this operation
Found a complication in your heart, so long
'Cause now you've got (now you've got)
Maybe just two weeks to live
Is that the most the both of you can give?
One, two, one, two, three, four
La-la-la la-la
La-la-la-la la-la
La-la-la-la la-la-la
Well, come on
La-la-la la-la
La-la-la-la la-la
La-la-la-la la-la-la
Oh, motherfucker
If life ain't just a joke
Then why are we laughing?
If life ain't just a joke
Then why are we laughing?
If life ain't just a joke
Then why are we laughing?
If life ain't just a joke
Then why am I dead?
Dead!`
*/

func main() {
	fileName := "dead.lyrics"
	file, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error while opening file.")
		return
	}

	contents := string(file)

	res := strings.Fields(contents)
	counts := make(map[string]int)
	for _, str := range res {
		counts[str]++
	}
	fmt.Println(counts)
}
