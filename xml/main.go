package main

import (
	"encoding/xml"
	"fmt"
)

type Imposto struct {
	Icms []string `xml:"imposto>ICMS>ICMS00>vICMS"`
}

func main(){

	loadXML(Imposto{Icms: []string{}})

}

func loadXML(refImposto Imposto){

	data := `
        <det>
                <prod>
                    <cProd>GL-R96HG</cProd>
                    <cEAN></cEAN>
                    <xProd>CORTINA DE SEGURANCA ( TRANSMI</xProd>
                    <NCM>85365090</NCM>
                    <CFOP>5106</CFOP>
                    <uCom>UN</uCom>
                    <qCom>1.0000</qCom>
                    <vUnCom>8483.73000000</vUnCom>
                    <vProd>8483.73</vProd>
                    <cEANTrib></cEANTrib>
                    <uTrib>UN</uTrib>
                    <qTrib>1.0000</qTrib>
                    <vUnTrib>8483.73000000</vUnTrib>
                    <indTot>1</indTot>
                </prod>
                <imposto>
                    <vTotTrib>4074.23</vTotTrib>
                    <ICMS>
                        <ICMS00>
                            <orig>1</orig>
                            <CST>00</CST>
                            <modBC>3</modBC>
                            <vBC>9756.29</vBC>
                            <pICMS>18.0000</pICMS>
                            <vICMS>1756.13</vICMS>
                        </ICMS00>
                    </ICMS>
                    <IPI>
                        <cEnq>999</cEnq>
                        <IPITrib>
                            <CST>50</CST>
                            <vBC>8483.73</vBC>
                            <pIPI>15.0000</pIPI>
                            <vIPI>1272.56</vIPI>
                        </IPITrib>
                    </IPI>
                    <PIS>
                        <PISAliq>
                            <CST>01</CST>
                            <vBC>8483.73</vBC>
                            <pPIS>1.6500</pPIS>
                            <vPIS>139.98</vPIS>
                        </PISAliq>
                    </PIS>
                    <COFINS>
                        <COFINSAliq>
                            <CST>01</CST>
                            <vBC>8483.73</vBC>
                            <pCOFINS>7.6000</pCOFINS>
                            <vCOFINS>644.76</vCOFINS>
                        </COFINSAliq>
                    </COFINS>
                </imposto>
            </det>`

	err := xml.Unmarshal([]byte(data), &refImposto)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("Names of people: %q", refImposto)
}