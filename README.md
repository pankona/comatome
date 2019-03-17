# comatome

Shows your GitHub contributions summary in various text format.  
Sorry, under construction...

## How this looks like

```bash
# Shows commits per repository committed on specified term
$ comatome -from 2019-01 -to 2019-02 -all
12	pankona/dotfiles
25	pankona/comatome
3	pankona/hashira
30	pankona/ki
22	pankona/kodama
92 commits

pankona/ki
pankona/comatome
pankona/dotfiles
3 repositories created

1	pankona/comatome
2	pankona/gocui
3 pull requests opened
```

## Usage

```bash
Usage of comatome:
  -all
    	show all
  -co
    	show commits
  -from string
    	specify time contributed from (yyyy-mm)
  -oi
    	show opened issues
  -op
    	show opened pull requests
  -re
    	show created repositories
  -rp
    	show reviewed pull requests
  -to string
    	specify time contributed to (yyyy-mm)
```

## License

MIT

## Author

Yosuke Akatsuka (a.k.a [@pankona](https://github.com/pankona))
