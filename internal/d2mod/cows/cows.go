package cows

import (
	"github.com/tlentz/d2modmaker/internal/d2fs"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/cubeMain"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/superUniques"
)

func AddTpRecipe(d2files d2fs.Files) {
	// Add New Recipe for Cow Poral (tp scroll -> cow portal)
	cubeF := d2files.Get(cubeMain.FileName)
	// newCubeRows := make([][]string, 0)
	for _, row := range cubeF.Rows {
		description := row[cubeMain.Description]

		if description == cubeMain.CowPortalWirt {
			tmp := make([]string, len(row))
			// // copy cow row to tmp
			copy(tmp, row)

			// change tmp to remove wirts leg
			tmp[cubeMain.Description] = cubeMain.CowPortalNoWirt
			tmp[cubeMain.NumInputs] = "1"
			tmp[cubeMain.Input1] = "tsc"
			tmp[cubeMain.Input2] = ""

			cubeF.Rows = append(cubeF.Rows, tmp)
		}
	}
}

func AllowKingKill(d2files d2fs.Files) {
	// Enable ability to kill cow king and still create portal
	suF := d2files.Get(superUniques.FileName)
	for idx, row := range suF.Rows {
		name := row[superUniques.Name]

		if name == superUniques.CowKing {
			suF.Rows[idx][superUniques.HcIdx] = "1"
		}
	}
}
