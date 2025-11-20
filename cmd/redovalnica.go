package main
import (
	"context"
	"fmt"
    "github.com/DRGBarrel/vpsa-nal5/redovalnica"
    "github.com/urfave/cli/v3"
	"log"
	"os"
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
            var studenti map[string]redovanje.Student
            studenti=make(map[string]redovanje.Student)
            redovanje.AddStudent(studenti,"63230000","Tone","Kovač",[]int{7})
            redovanje.AddStudent(studenti,"63230002","Eni","Drugi",[]int{10})
            redovanje.AddStudent(studenti,"63230004","Ime","Priimek",[]int{4,6})
            fmt.Println(studenti)
            redovanje.DodajOceno(studenti,"42",3)
            redovanje.DodajOceno(studenti,"63230000",8)
            redovanje.DodajOceno(studenti,"63230000",42)
            fmt.Println(studenti)
            redovanje.IzpisVsehOcen(studenti)
            redovanje.IzpisiKoncniUspeh(studenti)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

    