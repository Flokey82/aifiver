# aifiver: Simplified five factor model

This package provides a simplified [five factor model](https://en.wikipedia.org/wiki/Big_Five_personality_traits) (also known as the big five personality traits) for behavioral modeling and whatnot.

I developed this package to somewhat mimic a more "scientific" AI behavior ruleset somewhat akin to Crusader Kings' AI with its very impressive [traits and skills system](https://ck3.paradoxwikis.com/Traits) in order to give some realism to the decisions of Kings, Queens, and Tyrants.

It is still very much bare bones right now, but it lays the ground work for my overall vision.

## Model variations

There are two variants of the model:

* A detailed one with all facets and individual weights
* A simplified one with only the five factors and overall weights

Both provide the same interface, so use whatever floats your boat.

## Traits

One can freely define traits that are only expressed by someone who fulfills specified criteria (like minimum or maximum trait scores required).

Examples for traits are:
* Paranoid
* Gullable
* Optimistic
* Emotional
* Aggressive
* etc.

## Skills

Skills are more or less how skillful a certain personality is, given their preferences, strengths, and weaknesses.

Skills include:
* Diplomacy
* Martial
* Stewardship
* Intrigue
* etc.
