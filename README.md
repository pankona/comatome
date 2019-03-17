# comatome

Shows your GitHub contributions summary in various text format.  
Sorry, under construction...

## How this looks like

```bash
# e.g) Show contributions on 2019/03
$ comatome -from 2019-03 -to 2019-03 -all

Created 80 commits in 13 repositories
1	CPSPlatform/CloudConnectAgent
28	CPSPlatform/CoreService
8	CPSPlatform/PT_Gas_US
1	CPSPlatform/PostProcessService
1	google/go-github
1	pankona/TIL
7	pankona/comatome
4	pankona/dotfiles
1	pankona/godev
3	pankona/gomo-simra
12	pankona/hashira
1	pankona/orderedmap
12	pankona/slides

Created 2 repositories
pankona/go-github
pankona/TIL

Opened 12 pull requests in 7 repositories
1	CPSPlatform/CloudConnectAgent
4	CPSPlatform/CoreService
1	CPSPlatform/PT_Gas_US
1	google/go-github
1	pankona/comatome
1	pankona/gomo-simra
3	pankona/hashira

Reviewed 26 pull requests in 10 repositories
1	CPSPlatform/AppTemplate
1	CPSPlatform/CPSPAgent
4	CPSPlatform/CloudConnectAgent
4	CPSPlatform/CoreService
8	CPSPlatform/PT_DigitalManufacture
1	CPSPlatform/PT_Gas_US
2	CPSPlatform/PostProcessService
2	CPSPlatform/godev
2	access-company/one_hccex
1	pankona/godev

Opened 6 issues in 4 repositories
1	CPSPlatform/PT_Gas_US
1	google/go-github
1	pankona/gomo-simra
3	pankona/hashira
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
