# Zombie Attack

## Description
The point of this exercise is to model zombie attack using golang. The objective is that a user can spawn a zombie and then those zombies remain active until destroyed by a zombie hunter.  Each zombie has a chance to create additional zombies to be destroyed by a zombie hunter every set period of time.

## Back-End
- When the app starts, there are no zombies
- There is a command that spawns a zombie when called
    - While the zombie is active, it has chance to create another zombie or be destroyed by the zombie hunter
    - Every 5 seconds, each zombie has 20% chance for creating a new zombie
    - Every 10 seconds, each zombie has a 25% change of being destroyed by the zombie hunter
    - Ther is also a commnd for manually destroying a zombie
    
