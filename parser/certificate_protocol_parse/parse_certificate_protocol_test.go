package certificate_protocol_parse

import (
	"bileti_go/parser"
	"reflect"
	"strings"
	"testing"
	"time"
)

const html = `
<html><head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title></title>
<link rel="StyleSheet" type="text/css" href="/style.css">
</head>
<body data-new-gr-c-s-check-loaded="14.1239.0" data-gr-ext-installed="">
									<table width="650" align="center" border="0" summary="" class="sech">
										<tbody><tr>
											<td rowspan="2" align="center" width="80" height="53"><img src="/images/logo.png"></td>
											<td align="center"><font size="+1">ИЗПЪЛНИТЕЛНА АГЕНЦИЯ ПО ГОРИТЕ</font></td>
										</tr>
										<tr>
											<td colspan="2" align="center"><font size="2px">РЕГИОНАЛНА ДИРЕНЦИЯ ПО ГОРИТЕ <b>Варна</b></font></td>
										</tr>
										<tr>
											<td colspan="2" align="center"><hr></td>
										</tr>
									</tbody></table>
<table border="0" align="center" width="650" cellpadding="0" cellspacing="8" class="sech">
	<tbody><tr>
		<td colspan="2" align="center"><br><font size="2"><i><b>ПРОТОКОЛ за освидетелстване на сечище</b></i> № <b>0800519</b></font></td>
	</tr>
    <tr>
      <td colspan="2" align="center">Към позволително за сеч № 0800703/08.01.2025 год.</td>
    </tr>
    <tr>
      <td colspan="2">&nbsp;</td>
    </tr>
<tr>
  <td colspan="2" align="left">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Днес: <b>08.04.2025</b> год., на основание чл. 109 от Закона за горите се състави настоящият</td>
</tr>
<tr>
  <td colspan="2" align="left">констативен протокол от: <b>...........</b></td>
</tr>
<tr>
  <td colspan="2" align="left">в присъствието на: <b>...........</b></td>
</tr>
<tr>
  <td colspan="2" align="left">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>1.</b> След съвместно обхождане и проверка на сечището в Община: <b>Балчик</b>;</td>
</tr>
<tr>
  <td colspan="2" align="left">землище: <b>Балчик</b>; Отдел: <b>53</b>; Подотдел: <b>10</b>; Кад. № <b>02508.62.38</b>;</td>
</tr>
<tr>
  <td colspan="2" align="left">собственост на: <b>ДГТ</b>, и на всички документи</td>
</tr>
<tr>
  <td colspan="2" align="left">за сечта и добива на материалите констатира:</td>
</tr>
<tr>
  <td align="left" colspan="2">Площ по позволително за сеч (ха): <b>0.090</b>&nbsp;&nbsp;|&nbsp;&nbsp;
Сечта е проведена на площ (ха): <b>0.080</b></td>
</tr>
<tr><td colspan="2">
	<table align="center" width="640" border="0" cellpadding="0" cellspacing="0" class="ex">
		<tbody><tr>
			<td rowspan="2" align="center" width="20">№</td>
			<td rowspan="2" align="center">Категория дървесина</td>
			<td rowspan="2" align="center">Дървесен вид</td>
			<td colspan="4" align="center">Лежаща маса в куб. м.</td>
		</tr>
		<tr>
			<td align="center">По позволително <br>за сеч</td>
			<td align="center">Действително <br>отсечено</td>
			<td align="center">Наличност в <br>сечището</td>
			<td align="center">Налично на <br>временен склад</td>
		</tr>
		<tr>
			<td align="center">1.</td>
			<td align="left">&nbsp;Едра</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
		</tr>
		<tr>
			<td align="center">2.</td>
			<td align="left">&nbsp;Средна</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
		</tr>
		<tr>
			<td align="center">3.</td>
			<td align="left">&nbsp;Дребна</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
		</tr>
		<tr>
			<td rowspan="4" align="center">4.</td>
			<td rowspan="4">&nbsp;Дърва</td>
			<td align="right"><b>Ясен</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
		</tr>
		<tr>
			<td align="right"><b>Други широколистни</b></td>
			<td align="right"><b>0.1</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
		</tr>
		<tr>
			<td align="right"><b>Айлант</b></td>
			<td align="right"><b>0.1</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
		</tr>
		<tr>
			<td align="right"><b>Джанка</b></td>
			<td align="right"><b>0.3</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
		</tr>
		<tr>
			<td align="center">5.</td>
			<td align="left">&nbsp;Вършина</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
		</tr>
		<tr>
			<td colspan="3" align="right"><b>Общо количество:</b></td>
			<td align="right"><b>0.5</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
		</tr>
	</tbody></table>
</td></tr>
<tr>
  <td align="left" colspan="2">Налични неотсечени маркирани стъбла : 
<b>Да</b></td>
</tr>
<tr>
  <td colspan="2" align="left">Разликите между количествата и категориите дървесина, по позволително за сеч и фактически добити в повече или в по-малко се дължът на следното:&nbsp;
<b>Сечта е изведена по КП на РДГ Варна съставен за установяване на опасни дървета в имот 02508.62.38.Дървесината е негодна за употреба</b></td>
</tr>
<tr>
  <td colspan="2" align="left">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>2.</b> При сечта и извоза на дървесина са допуснати следните нарушения на Закона за горите:&nbsp;
  <b>Дарамдудай</b></td>
</tr>
<tr>
  <td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>3.</b> Провеждането на сечта и почистването на сечището е извършено според определения начин в технологичния план и позволителното за сеч:
<b>Задоволително</b>
</td></tr>
<tr>
  <td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>4.</b> За допуснатите нарушения е съставен акт за установяване на административно<br> 
 нарушение № <b>1337</b> / <b>&nbsp;</b>
 на: <b>Иван Петков</b>
</td></tr>
<tr>
  <td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>5.</b> За подобряване състоянието на насаждението в срок до: 
  <b>&nbsp;</b>
 следва да се проведът следните мероприятия:&nbsp;
 <b></b></td>
</tr>
<tr>
		<td colspan="2">&nbsp;</td>
</tr>
<tr>
  <td align="left">Съставил: ..........................................</td>
  <td align="left">Присъствал: ..........................................</td>
</tr>
<tr>
  <td align="left"><b>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;...........</b></td>
  <td align="left"><b>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;...........</b></td>
</tr>
	</tbody></table>


</body><grammarly-desktop-integration data-grammarly-shadow-root="true"></grammarly-desktop-integration></html>
`

func TestShouldMatchObject(t *testing.T) {
	actual := ParseCertificateProtocol(strings.NewReader(html))
	expected := CertificateProtocolParseResult{
		id:                          "0800519",
		regionalForestryDirectorate: "Варна",
		permittedForFellingId:       "0800703",
		permittedForFellingDate:     time.Date(2025, 1, 8, 0, 0, 0, 0, parser.GetLocation()),
		when:                        time.Date(2025, 4, 8, 0, 0, 0, 0, parser.GetLocation()),
		findingsRecordWitnessedBy:   "",
		findingsRecordIssuedBy:      "",
		municipality:                "Балчик",
		land:                        "Балчик",
		section:                     "53",
		subsection:                  "10",
		cadastreId:                  "02508.62.38",
		ownership:                   "ДГТ",
		areaOfPermit:                0.09,
		areaDone:                    0.08,
		notMarkedLoggedTreesExist:   true,
		volumeDifferenceDescription: `Сечта е изведена по КП на РДГ Варна съставен за установяване на опасни дървета в имот 02508.62.38.Дървесината е негодна за употреба`,
		violationDescription:        "Дарамдудай",
		loggingFinishDescription:    "Задоволително",
		violationId:                 "1337",
		violationRecipient:          "Иван Петков",
	}

	if reflect.DeepEqual(actual, expected) == false {
		t.Error("Objects are not equal")
	}
}
