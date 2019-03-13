/*
	Copyright © 2019 The Goca.io team

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	    http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package ods

import (
	"os"
	"strings"
	"testing"

	"github.com/gocaio/goca"
	"github.com/gocaio/goca/gocaTesting"
)

// Test server URL.
var testserver = os.Getenv("GOCA_TEST_SERVER")

// T is a global reference for the test. This allows us to use *testing.T
// methods anywhere
var T *testing.T

// TestReadODS tests the reading over ODS files
func TestReadODS(t *testing.T) {
	T = t // Assignment t (*testing.T to a global T variable)
	// Get a controller
	ctrl := goca.NewControllerTest()
	// Subscribe a processOutput. The propper test will be placed in proccessOutput
	ctrl.Subscribe(goca.Topics["NewOutput"], processOutput)

	// Call the plugin entrypoint
	setup(ctrl)

	gocatesting.GetAssets(t, ctrl, testserver, plugName)
}

func processOutput(module, url string, out *goca.Output) {
	// We have to validate goca.Output according to the resource
	parts := strings.Split(out.Target, "/")

	switch parts[len(parts)-1] {
	case "Ods1.ods":
		validateCaseA(out)
	case "Ods2.ods":
		validateCaseB(out)
	case "Ods3.ods":
		validateCaseC(out)
	}
}

func validateCaseA(out *goca.Output) {
	if out.MainType != "ODS" {
		T.Errorf("expected ODS but found %s", out.MainType)
	}
	if out.Title != "WEIDMULLEREXT141010" {
		T.Errorf("expected \"WEIDMULLEREXT141010\" but found %s", out.Title)
	}
	if out.Description != "para promoción y venta de productos, aceptamos solicitudes en:  ventas@controlfr.com  y en   mario_paredes@detodoelectrico.com  no deje de visitarnos en:   www.controlfr.com       www.detodoelectrico.com      www.freda.com.mx           teléfonos:  5440 6090, 5440 6120, 5440 6144 y 5440 6145" {
		T.Errorf("expected long text but found %s", out.Description)
	}
	if out.Comment != "Inventario" {
		T.Errorf("expected \"Inventario\" but found %s", out.Comment)
	}
	if out.Lang != "es-ES" {
		T.Errorf("expected \"es-ES\" but found %s", out.Lang)
	}
	if out.CreatorTool != "OpenOffice.org/3.2$Win32 OpenOffice.org_project/320m18$Build-9502" {
		T.Errorf("expected \"OpenOffice.org/3.2$Win32 OpenOffice.org_project/320m18$Build-9502\" but found %s", out.CreatorTool)
	}
	if out.Producer != "" {
		T.Errorf("expected \"\" but found %s", out.Producer)
	}
	if out.Keywords != "WEIDMÜLLER" {
		T.Errorf("expected \"WEIDMÜLLER\" but found %s", out.Keywords)
	}
	if out.ModifiedBy != "" {
		T.Errorf("expected \"\" but found %s", out.ModifiedBy)
	}
	if out.CreateDate != "2009-05-06T15:55:29" {
		T.Errorf("expected \"2009-05-06T15:55:29\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2015-03-12T13:35:45.05" {
		T.Errorf("expected \"2015-03-12T13:35:45.05\" but found %s", out.ModifyDate)
	}
}

func validateCaseB(out *goca.Output) {
	if out.MainType != "ODS" {
		T.Errorf("expected ODS but found %s", out.MainType)
	}
	if out.Title != "Textzeugen" {
		T.Errorf("expected \"Textzeugen\" but found %s", out.Title)
	}
	if out.Description != "Diese Tabelle ist Bestandteil des Projekts \"Meister Eckhart und seine Zeit\", http://www.eckhart.deStand: 11. November 2017© Eckhart Triebel, 2008-2017" {
		T.Errorf("expected long text but found %s", out.Description)
	}
	if out.Comment != "Meister Eckhart" {
		T.Errorf("expected \"Meister Eckhart\" but found %s", out.Comment)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.CreatorTool != "LibreOffice/5.1.6.2$Linux_X86_64 LibreOffice_project/10m0$Build-2" {
		T.Errorf("expected \"LibreOffice/5.1.6.2$Linux_X86_64 LibreOffice_project/10m0$Build-2\" but found %s", out.CreatorTool)
	}
	if out.Producer != "Eckhart Triebel" {
		T.Errorf("expected \"Eckhart Triebel\" but found %s", out.Producer)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.ModifiedBy != "Eckhart Triebel" {
		T.Errorf("expected \"Eckhart Triebel\" but found %s", out.ModifiedBy)
	}
	if out.CreateDate != "2015-04-22T22:25:25" {
		T.Errorf("expected \"2015-04-22T22:25:25\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2017-11-12T17:33:52.578515368" {
		T.Errorf("expected \"2017-11-12T17:33:52.578515368\" but found %s", out.ModifyDate)
	}
}

func validateCaseC(out *goca.Output) {
	if out.MainType != "ODS" {
		T.Errorf("expected ODS but found %s", out.MainType)
	}
	if out.Title != "Überweisungsformular mit Scalc" {
		T.Errorf("expected \"Überweisungsformular mit Scalc\" but found %s", out.Title)
	}
	if out.Description != "Vorliegendes Formular habe ich auf der Basis von Scalc entwickelt, weil das von mir bisher benutzte Formular mit der neuen Version von OpenOffice 3.3 nicht mehr funktionierte und ich nicht auf Fehlersuche gehen wollte..Im oberen Bereich ist ein normales Überweisungsformular abgebildet. Die einzelnen Felder können mit Daten gefüllt werden. Im unteren Bereich werden im Doppel automatisch die gleichen Daten eingetragen und kann somit als Kopie genutzt werden. Die Größe reicht auch aus um den Beleg gelocht abzulegen.Das Formular kann frei benutzt werden, bitte aber darum, die Erläuterungenunverändert zu lassen. Sollten weitere Erläuterungen notwendig sein, kann ja einzusätzliches Tabellenblatt eingefügt werden.Bei Bank wird per SVERWEIS automatisch der Bankname eingetragen sobald das Feld Bankleitzahl mit einer gültigen BLZ gefüllt ist. SVERWEIS benutzt dabei die Daten aus demTabellenblatt BLZ.Falls die Formel in Tabelle1.F25 versehentlich gelöscht werden sollte. Hier ist sie:=WENN(BN20<>\"\";SVERWEIS(BN20;BLZ.B5:BLZ.E19865;2;1);\"\")Viel Spaß mit der Datei.Günter Rohler, Januar 2011" {
		T.Errorf("expected long text but found %s", out.Description)
	}
	if out.Comment != "Überweisungsvordruck zum Ausdruck und Versand" {
		T.Errorf("expected \"Überweisungsvordruck zum Ausdruck und Versand\" but found %s", out.Comment)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.CreatorTool != "OpenOffice.org/3.2$Win32 OpenOffice.org_project/320m18$Build-9502" {
		T.Errorf("expected \"OpenOffice.org/3.2$Win32 OpenOffice.org_project/320m18$Build-9502\" but found %s", out.CreatorTool)
	}
	if out.Producer != "" {
		T.Errorf("expected \"\" but found %s", out.Producer)
	}
	if out.Keywords != "Überweisungsvordruck mit Kopie" {
		T.Errorf("expected \"Überweisungsvordruck mit Kopie\" but found %s", out.Keywords)
	}
	if out.ModifiedBy != "" {
		T.Errorf("expected \"\" but found %s", out.ModifiedBy)
	}
	if out.CreateDate != "2011-02-12T11:59:46" {
		T.Errorf("expected \"2011-02-12T11:59:46\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2011-02-12T12:02:34" {
		T.Errorf("expected \"2011-02-12T12:02:34\" but found %s", out.ModifyDate)
	}
}
