# text-game

Simple text game made using [Go](https://golang.org).

## How to play

```bash
make run
```

## Make your own!

This engine is created for everyone, you only need to define your own map in `src/map/{your_map}.json` to create your own game.

## Contribute

### Development

Checkout from latest `main` branch
```bash
git checkout main 
git pull origin main 
git checkout -b <your_branch>
```
Hint: please take a look at [Branch Convention](#branch-convention)

If you add other dependencies, run:
```bash
make update-dep 
```

Before raise a Pull Request, please make sure you already suffice the tests of your code.

```bash
make test
```

### Branch Convention

Format:
> [prefix]_[feature_name]

Prefix:
- `f_` for new feature implementation
- `i_` for adding code improvement
- `b_` for fixing bug

Examples:
- f_add_scoring_feature
- i_enable_custom_effect
- b_fix_error_stuck_in_map

### Issue / Feature Request

Please raise an issue and explains the issue / feature that you want to be supported.
Give detail explanation about the issue / feature.

## Contact

If you have anything to ask / discuss, please contact me below, thanks!   
Aji Imawan Omi  
GitHub: cymon1997

## License

GNU GENERAL PUBLIC LICENSE - Aji Imawan Omi