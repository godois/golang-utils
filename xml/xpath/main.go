package main

import (
	"log"
	"fmt"
	"gopkg.in/xmlpath.v1"
	"strings"
)

func main(){

	path := xmlpath.MustCompile("/det/imposto/ICMS/ICMS00/vICMS")

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

	root, err := xmlpath.Parse(strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	if value, ok := path.String(root); ok {
		fmt.Println("Found:", value)
	}

}
