package main
import (
    "redovalnica/redovanje"
    "github.com/urfave/cli/v3"
)
    
func main(){
    cmd := &cli.Command{
		Name:  "Redovalnica",
		Usage: "Manažiranje redovalnice ocen študentov glede na vpisne številke",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Najmanjše število ocen potrebnih za pozitivno oceno",
				Value: 5,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Najmanjša možna ocena",
				Value: 1,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Največja možna ocena",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")
            redovanje.SetSt(stOcen)
            redovanje.SetMin(minOcena)
            redovanje.SetMax(maxOcena)
            var studenti map[string]Student
            studenti=make(map[string]Student)
            studenti["63230000"]=Student{"Tone","Kovač",[]int{7}}
            studenti["63230002"]=Student{"Eni","Drugi",[]int{10}}
            studenti["63230004"]=Student{"Ime","Priimek",[]int{4,6}}
            fmt.Println(studenti)
            redovanje.dodajOceno(studenti,"42",3)
            redovanje.dodajOceno(studenti,"63230000",8)
            redovanje.dodajOceno(studenti,"63230000",42)
            fmt.Println(studenti)
            redovanje.izpisRedovalnice(studenti)
            redovanje.izpisiKoncniUspeh(studenti)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

    