# Complex Vortices

Simulation of the 2D nonlinear Schrödinger equation using the [split operator method](https://www.algorithm-archive.org/contents/split-operator_method/split-operator_method.html). Run this program with `go run .`. The output will be saved as a series of TGA files, which you can then merge to a video using an external program.

This project is done similar to a previous [Odin](https://github.com/marl0ny/Odin-Projects/tree/main/qm2d) and [Rust](https://github.com/marl0ny/Rust-Sims/tree/main/qm2d_split_op) implementation.  

## References:

### Split-Operator Method:

- James Schloss. [The Split Operator Method - Arcane Algorithm Archive.](https://www.algorithm-archive.org/contents/split-operator_method/split-operator_method.html)

- Xavier Antoine, Weizhu Bao, Christophe Besse.
  [Computational methods for the dynamics of the nonlinear Schrodinger/Gross-Pitaevskii equations.](https://arxiv.org/abs/1305.1093)

### Fast Fourier Transform (Used in the Split-Operator method):

- [Wikipedia - Cooley–Tukey FFT algorithm](https://en.wikipedia.org/wiki/Cooley%E2%80%93Tukey_FFT_algorithm)

- [MathWorld Wolfram - Fast Fourier Transform](http://mathworld.wolfram.com/FastFourierTransform.html)

- William Press et al. [12.2 Fast Fourier Transform (FFT) - in Numerical Recipes](https://websites.pmc.ucsc.edu/~fnimmo/eart290c_17/NumericalRecipesinF77.pdf)

### Domain coloring method for visualizing complex-valued functions:

- [Wikipedia - Domain coloring](https://en.wikipedia.org/wiki/Domain_coloring)

- [Wikipedia - Hue](https://en.wikipedia.org/wiki/Hue)

- [https://en.wikipedia.org/wiki/Hue#/media/File:HSV-RGB-comparison.svg](https://en.wikipedia.org/wiki/Hue#/media/File:HSV-RGB-comparison.svg)

### TGA file format:

- [Wikipedia - Truevision TGA](https://en.wikipedia.org/wiki/Truevision_TGA)
