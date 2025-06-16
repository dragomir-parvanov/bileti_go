package permit_table_parser

import (
	"bileti_go/utils"
	utils_test "bileti_go/utils/test"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestShouldReturnUrls(t *testing.T) {
	html := `
<div style="padding-top:22px;">
			<span class="font11">Открити издадени позволителни за сеч: <b>3</b>
	</div>
				
		<div style="padding-top:10px;">
			<div class="resultGrid">
			    <table width="929" cellpadding="0" cellspacing="0" border="1">
			      <tr valign="middle" style="background-color:#cbd1b2;">
				<td rowspan="2" align="center">№ / дата:</td>
					<td rowspan="2" align="center">Подотдел</td>
					<td rowspan="2" align="center">Кад. №</td>
					<td rowspan="2" align="center">Собственост</td>
					<td rowspan="2" align="center">куб. м.</td>
					<td rowspan="2" align="center">Площ <br>ха</td>
					<td rowspan="2" align="center">Срок сеч</td>
					<td rowspan="2" align="center">Срок извоз</td>
					<td rowspan="2" align="center">ДГС / ДЛС</td>
					<td colspan="4" align="center"><b>Протокол за освидетелстване</b></td>
			      </tr>
	<tr style="background-color:#cbd1b2;">
 	<td align="center" >№</td>
     <td align="center">Дата</td>
     <td align="center">Добито</td>
     <td align="center">Площ</td>
	</tr>
	<tr valign="middle" style="">
			<td align="center"><a href="Print_PozvSech.cgi?asdfuyoiyoisgdhuioy=798787&jklfhlkasjfhduwioefuoived=3ea9e30072b506964eeb5913b82165b72308ff00" target="new">
			<img src="/images/print_on.gif" height="15" border="0" alt="Принтиране"></a>&nbsp;798787 /
			05.01.2025 г.</td>
      <td align="center">182 л</td>
      <td align="left">02676.193.23</td>
      <td align="left">ОГТ</td>
      <td align="right">53.0</td>
      <td align="right">2.100</td>
      <td align="center">31.12.2025</td>
      <td align="center">31.12.2025</td>
			<td align="center">ДГС Добринище</td>
      <td align="right"></td>
      <td align="right"></td>
      <td align="right"></td>
      <td align="right"></td>
 </tr>
	<tr valign="middle" style="">
			<td align="center"><a href="Print_PozvSech.cgi?asdfuyoiyoisgdhuioy=798789&jklfhlkasjfhduwioefuoived=cdb39e0303dc8140e38137cf91833585d0052cc7" target="new">
			<img src="/images/print_on.gif" height="15" border="0" alt="Принтиране"></a>&nbsp;798789 /
			05.01.2025 г.</td>
      <td align="center">182 и</td>
      <td align="left">02676.185.7</td>
      <td align="left">ОГТ</td>
      <td align="right">95.2</td>
      <td align="right">4.000</td>
      <td align="center">31.12.2025</td>
      <td align="center">31.12.2025</td>
			<td align="center">ДГС Добринище</td>
      <td align="right"></td>
      <td align="right"></td>
      <td align="right"></td>
      <td align="right"></td>
 </tr>
	<tr valign="middle" style="">
			<td align="center"><a href="Print_PozvSech.cgi?asdfuyoiyoisgdhuioy=816701&jklfhlkasjfhduwioefuoived=9cd11b506d3e0ed583524d2a932fb77fdecfee73" target="new">
			<img src="/images/print_on.gif" height="15" border="0" alt="Принтиране"></a>&nbsp;816701 /
			25.04.2025 г.</td>
      <td align="center">182 и</td>
      <td align="left"></td>
      <td align="left">ДГТ</td>
      <td align="right">446.0</td>
      <td align="right">21.500</td>
      <td align="center">31.12.2025</td>
      <td align="center">31.12.2025</td>
			<td align="center">ДГС Добринище</td>
      <td align="right"></td>
      <td align="right"></td>
      <td align="right"></td>
      <td align="right"></td>
 </tr>
			    </table>
			  </div>
			</div>

`
	actual := utils.Must(ParsePermitTable(strings.NewReader(html)))

	expected := []string{
		`Print_PozvSech.cgi?asdfuyoiyoisgdhuioy=798787&jklfhlkasjfhduwioefuoived=3ea9e30072b506964eeb5913b82165b72308ff00`,
		`Print_PozvSech.cgi?asdfuyoiyoisgdhuioy=798789&jklfhlkasjfhduwioefuoived=cdb39e0303dc8140e38137cf91833585d0052cc7`,
		`Print_PozvSech.cgi?asdfuyoiyoisgdhuioy=816701&jklfhlkasjfhduwioefuoived=9cd11b506d3e0ed583524d2a932fb77fdecfee73`,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(`Expected  %v but got %v`, expected, actual)
	}
}

func TestShouldReturnEmptyArrayOnNoUrls(t *testing.T) {
	html := `
		<div style="padding-top:22px;">
			<span class="font11">Открити издадени позволителни за сеч: <b>0</b>
 <br> Няма издадено позволително за сеч по така зададените критерии.</span>
	</div>
				
		<div style="padding-top:10px;">
			<div class="resultGrid">
			    <table width="929" cellpadding="0" cellspacing="0" border="1">
			      <tr valign="middle" style="background-color:#cbd1b2;">
				<td rowspan="2" align="center">№ / дата:</td>
					<td rowspan="2" align="center">Подотдел</td>
					<td rowspan="2" align="center">Кад. №</td>
					<td rowspan="2" align="center">Собственост</td>
					<td rowspan="2" align="center">куб. м.</td>
					<td rowspan="2" align="center">Площ <br>ха</td>
					<td rowspan="2" align="center">Срок сеч</td>
					<td rowspan="2" align="center">Срок извоз</td>
					<td rowspan="2" align="center">ДГС / ДЛС</td>
					<td colspan="4" align="center"><b>Протокол за освидетелстване</b></td>
			      </tr>
	<tr style="background-color:#cbd1b2;">
 	<td align="center" >№</td>
     <td align="center">Дата</td>
     <td align="center">Добито</td>
     <td align="center">Площ</td>
	</tr>
			    </table>
			  </div>
			</div>


`

	actual := utils.Must(ParsePermitTable(strings.NewReader(html)))

	expected := make([]string, 0)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf(`Expected  %v but got %v`, expected, actual)
	}
}

func TestShouldReturnErrorOnReaderError(t *testing.T) {
	_, err := ParsePermitTable(&utils_test.ErrReader{Error: errors.New("reader error")})

	if err == nil {
		t.Fatalf(`Expected error, but got nill. Error %v`, err)
	}
}
