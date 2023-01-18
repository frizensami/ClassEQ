# ClassEQ: Genuinely real-time classroom mood / pacing / content understanding feedback system

## Motivation 

Polling software such as PollEverywhere, Kahoot, and Quizlet are increasingly common ways to get targeted feedback during classes for specific questions.

However, the *frequency* of this feedback may be too low. For instance:
- Did students understand the slides you just went through? (you can't continually have polls to test every single concept!)
- Do they think you're suddenly speaking too fast? (as opposed to polling for this once during a class)
- Do they have a question? (but don't want to interrupt your flow?)

**Core idea**: Instructors and students may need a more real-time way of receiving and giving feedback on the current class.

## Goal

**ClassEQ** aims to be a system for genuine **real-time** classroom feedback that is **not structured around polls / quizzes / etc**. 

It should be used to *augment* the existing targeted feedback tools mentioned above.

## User Experience Overview

### Students / Talk Attendees / etc

- Students should be able to join a "session" based on an ID or link name 
- Once joined, they should be presented with an interface with a minimal number of buttons that give live feedback to the presenter
- Such buttons could be grouped as:
	- Speed: OK (default) / Too fast / Too slow
	- Understanding: OK (default) / Do not understand recent concept
	- Question: No question (default) / Have a question

### Teachers

- Should be able to set up a session
- Should have an interface during the talk (a window) to view the current live status of the students


## Implementation Ideas

- React + Websockets frontend for students
- React + Websockets frontend for instructors
- Websockets backend to link the two together via unique link codes / names.

