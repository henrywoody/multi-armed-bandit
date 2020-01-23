# Multi-Armed Bandit

Solutions to the [multi-armed bandit problem](https://en.wikipedia.org/wiki/Multi-armed_bandit).

## Solutions

### Epsilon-greedy strategy

> The best lever is selected for a proportion `1 - ε`  of the trials, and a lever is selected at random (with uniform probability) for a proportion `ε` . A typical parameter value might be `ε = 0.1`, but this can vary widely depending on circumstances and predilections.

### Epsilon-first strategy

> A pure exploration phase is followed by a pure exploitation phase. For `N` trials in total, the exploration phase occupies `εN` trials and the exploitation phase `(1-ε)N` trials. During the exploration phase, a lever is randomly selected (with uniform probability); during the exploitation phase, the best lever is always selected.

### Epsilon-decreasing strategy

> Similar to the epsilon-greedy strategy, except that the value of `ε`  decreases as the experiment progresses, resulting in highly explorative behaviour at the start and highly exploitative behaviour at the finish.