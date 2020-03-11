# Multi-Armed Bandit

Solutions to the [multi-armed bandit problem](https://en.wikipedia.org/wiki/Multi-armed_bandit).

## Introduction

The multi-armed bandit problem is well-studied problem in the area of reinforcement learning. The name of the problem is comes from the name "one-armed bandit" for slot machines. The problem involves an agent and a set of ![equation](http://latex.codecogs.com/png.latex?n) levers. On each round, the agent must choose one lever to pull. Each lever may have different payout amounts, and distributions and the goal for the agent is to maximize total reward over time.

The key to the problem is finding a balance between exploration and exploitation. An agent needs to pull the best lever available to it as many times as possible in order to maximize reward, but the agent must explore the available levers in order to determine which is best.

### Notation and Definitions

An agent is an entity that exists in an environment and can make decisions and take actions within that environment. In these examples time progresses discretely in the environment. For each time ![equation](http://latex.codecogs.com/png.latex?t%20%5Cin%20%5C%7B%200%2C%201%2C%202%2C%20%5Cdots,%20T%20%5C%7D) the state of the environment is ![equation](http://latex.codecogs.com/png.latex?s_t%20%5Cin%20S), where ![equation](http://latex.codecogs.com/png.latex?S) is the *state space*. In each round (each timestep), the agent can take an action ![equation](http://latex.codecogs.com/png.latex?a%20%5Cin%20A), where ![equation](http://latex.codecogs.com/png.latex?A) is the *action space*. Taking an action results in a reward ![equation](http://latex.codecogs.com/png.latex?r%20%5Cin%20R), where ![equation](http://latex.codecogs.com/png.latex?R) is the *reward space*.

In this project, an action is the choice of the lever to pull. The variable ![equation](http://latex.codecogs.com/png.latex?n) typically denotes the number of levers. Taking an action results in a reward.

The true value of an action ![equation](http://latex.codecogs.com/png.latex?a) is denoted by ![equation](http://latex.codecogs.com/png.latex?Q%5E%7B%2A%7D%28a%29) and the value of the action, as estimated by the agent, at time ![equation](http://latex.codecogs.com/png.latex?t) is ![equation](http://latex.codecogs.com/png.latex?Q_%7Bt%7D%28a%29).

One simple formula for estimating the value of an action, is to calculate the sample average for the action. If an action ![equation](http://latex.codecogs.com/png.latex?a) has been taken ![equation](http://latex.codecogs.com/png.latex?k_a) times, resulting in the rewards ![equation](http://latex.codecogs.com/png.latex?r_1,%20r_2,%20%5Cdots,%20r_{k_a}), the agent can use the formula:

![equation](http://latex.codecogs.com/png.latex?Q_%7Bt%7D%28a%29%20=%20%5Cfrac%7B%20r_1%20%2B%20r_2%20%2B%20%5Cdots%20%2B%20r_%7Bk_a%7D%20%7D%7B%20k_a%20%7D)

By estimating the value of each lever at a given time, the agent can identify the lever with the highest estimated value. The action of pulling this lever is called the *greedy* action. The action with the greatest true value is denoted by ![equation](http://latex.codecogs.com/png.latex?a%5E%2A).

A *policy* defines an agent's behavior. On each timestep ![equation](http://latex.codecogs.com/png.latex?t) an agent uses its policy to examine the state of the environment ![equation](http://latex.codecogs.com/png.latex?s_t) and take an action ![equation](http://latex.codecogs.com/png.latex?a_t), resulting in a reward ![equation](http://latex.codecogs.com/png.latex?r_%7B%20t%20%2B%201%20%7D).

Formally, the goal of the problem is to maximize the cumulative reward function:

![equation](http://latex.codecogs.com/png.latex?%5Csum_%7B%20t%3D0%20%7D%5E%7B%20T%20%7D%20r_t)

## Solutions

### Epsilon-Greedy Strategy

> The best lever is selected for a proportion `1 - ε`  of the trials, and a lever is selected at random (with uniform probability) for a proportion `ε` . A typical parameter value might be `ε = 0.1`, but this can vary widely depending on circumstances and predilections.

### Epsilon-First Strategy

> A pure exploration phase is followed by a pure exploitation phase. For `N` trials in total, the exploration phase occupies `εN` trials and the exploitation phase `(1-ε)N` trials. During the exploration phase, a lever is randomly selected (with uniform probability); during the exploitation phase, the best lever is always selected.

### Epsilon-Decreasing Strategy

> Similar to the epsilon-greedy strategy, except that the value of ![equation](http://latex.codecogs.com/png.latex?%5Cvarepsilon) decreases as the experiment progresses, resulting in highly explorative behaviour at the start and highly exploitative behaviour at the finish.

### Softmax Boltzmann Strategy

Select actions according to probabilities, rather than selecting the best action or a random action (as in the epsilon-_ strategies). On each round, calculate a probability of selection for each possible action (using the action's value estimate), and then select an action according to those probabilities. The preceding description applies for all softmax strategies, and the difference between difference softmax strategies is the calculation of the probabilities. In the Boltzmann case, for each action ![equation](http://latex.codecogs.com/png.latex?a), the probability of selection is:



![equation](http://latex.codecogs.com/png.latex?%5Cfrac%7B%20e%5E%7BQ_%7Bt%7D%28a%29%2F%5Ctau%7D%20%7D%7B%20%5Csum_%7Bb%3D1%7D%5E%7Bn%7D%20e%5E%7BQ_%7Bt%7D%28b%29%2F%5Ctau%7D%20%7D)



### Value-Difference Based Exploration (VDBE) Strategy

The Value-Difference Based Exploration strategy is built upon the epsilon-greedy strategy, and, similar to the epsilon-decreasing strategy, the value of epsilon is updated over time. In the VDBE strategy, the value of epsilon is decreased as the agent becomes more sure of its environment and increases if the agent finds that its understanding of the environment is wrong. This is calculated by taking the difference between the estimated value of an action and the reward received from taking that action (this term is called the *temporal-difference error*). For large temporal difference errors, the value of epsilon is increased, and for small errors, epsilon is decreased. See [2] for details.



## Sources

1. Sutton, Richard S. & Barto, Andrew G. (1998). *Reinforcement learning: an introduction*. Cambridge, MA: The MIT Press.
2. Tokic M. (2010) Adaptive *ε*-Greedy Exploration in Reinforcement Learning Based on Value Differences. In: Dillmann R., Beyerer J., Hanebeck U.D., Schultz T. (eds) KI 2010: Advances in Artificial Intelligence. KI 2010. Lecture Notes in Computer Science, vol 6359. Springer, Berlin, Heidelberg