# Multi-Armed Bandit

Solutions to the [multi-armed bandit problem](https://en.wikipedia.org/wiki/Multi-armed_bandit).

## Solutions

### Epsilon-Greedy Strategy

> The best lever is selected for a proportion `1 - ε`  of the trials, and a lever is selected at random (with uniform probability) for a proportion `ε` . A typical parameter value might be `ε = 0.1`, but this can vary widely depending on circumstances and predilections.

### Epsilon-First Strategy

> A pure exploration phase is followed by a pure exploitation phase. For `N` trials in total, the exploration phase occupies `εN` trials and the exploitation phase `(1-ε)N` trials. During the exploration phase, a lever is randomly selected (with uniform probability); during the exploitation phase, the best lever is always selected.

### Epsilon-Decreasing Strategy

> Similar to the epsilon-greedy strategy, except that the value of `ε`  decreases as the experiment progresses, resulting in highly explorative behaviour at the start and highly exploitative behaviour at the finish.

### Softmax Boltzmann Strategy

Select actions according to probabilities, rather than selecting the best action or a random action (as in the epsilon-_ strategies). On each round, calculate a probability of selection for each possible action (using the action's value estimate), and then select an action according to those probabilities. The preceding description applies for all softmax strategies, and the difference between difference softmax strategies is the calculation of the probabilities. In the Boltzmann case, for each action `a`, the probability of selection is:



![equation](http://latex.codecogs.com/gif.latex?%5Cfrac%7B%20e%5E%7BQ_%7Bt%7D%28a%29%2F%5Ctau%7D%20%7D%7B%20%5Csum_%7Bb%3D1%7D%5E%7Bn%7D%20e%5E%7BQ_%7Bt%7D%28b%29%2F%5Ctau%7D%20%7D)