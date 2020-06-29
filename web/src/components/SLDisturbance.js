//https://api.sl.se/api2/deviations.JSON?key=750131196e424fadb988372cb83df24c&transportMode=metro&lineNumber=17,18,19

import React from 'react';
import axios from 'axios';

class SLDisturbance extends React.Component {
  constructor(props) {
    super(props);
    this.state = { disturbanceData : [] };
  }

  componentDidMount() {
    if(this.props.refreshRate) {
      this.refreshInterval = setInterval(this.fetchSLDisturbances, this.props.refreshRate);
    }

    this.fetchSLDisturbances();
  }

  componentWillUnmount() {
    clearInterval(this.refreshInterval);
  }

  fetchSLDisturbances = () => {
    const urlParameters = new URLSearchParams({
      'key': this.props.apiKey,
      'transportMode': this.props.transportMode,
      'lineNumber': this.props.lineNumber.join(','),
    });

    const url = `http://api.sl.se/api2/deviations.JSON?${urlParameters}`;
    console.log(url, 1);
    
//    axios.get('http://192.168.0.88:80/tun', null, {headers: {Tunnel: 'dof'}})

    this.setState({'disturbanceData': []});

    let fakeData = {"StatusCode":0,"Message":null,"ExecutionTime":121,"ResponseData":[{"Created":"2020-06-22T11:06:30.05+02:00","MainNews":false,"SortOrder":1,"Header":"Inställd trafik på linje 19 mellan Globen och Gullmarsplan 22 juni-16 augusti","Details":"Inställd trafik på tunnelbanans gröna linje 19 mellan Globen  från och med måndag 22 juni till och med söndag 16 augusti, på grund av utbyggnad av nya tunnelbanan.\n\nSå här påverkas trafiken:\n\nLinje 19 går endast mellan Hagsätra och Globen i båda riktningarna, som tätast var 10:e minut. Linjerna 17 och 18 påverkas inte av arbetet utan stannar som vanligt vid Gullmarsplan, det går däremot inte att byta till linje 19 vid Gullmarsplan.\n\nBussar ersätter tunnelbanan under helgnätter\nTunnelbanetrafiken mellan Hagsätra och Globen ersätts av nattbuss 195 mellan Hagsätra och Gullmarsplan, nätter mot lördagar och söndagar ca kl. 01-06.\nBussarna avgår ca var 15:e minut fram till kl. 03:45, därefter var 30:e minut. Restiden mellan Gullmarsplan och Hagsätra är ungefär 25 minuter.\n\nI båda riktningarna trafikerar buss 195 dessa hållplatser: \n\nGullmarsplan - Enskede gård - Sockenplan - Svedmyra - Postiljonsvägen - Söråkersvägen (Stureby) - Örbyleden - Bandhagen - Bäckahagens skola - Sjösavägen - Skebokvarnsvägen - Högdalen - Rågsved - Lerbäcksgränd - Hagsätra.\n\nI riktning mot Hagsätra trafikeras även Globen (hållplats Slakthuset), Porlabacken och Frövigatan.\nI riktning mot Gullmarsplan trafikeras även Bolidenvägen.\n\nAlternativa resvägar\n\nVi hänvisar till promenad mellan Globen och Gullmarsplan. Flera lokala busslinjer går mellan stationerna längs Hagsätragrenen och andra tunnelbanestationer eller till pendeltågsstationer, använd reseplaneraren på sl.se eller i valfri app.\nMellan stationerna Högdalen, Bandhagen, Stureby, Svedmyra, Sockenplan och innerstaden (Cityterminalen) kommer busslinje 169C att gå under avstängningen, på vardagar ca kl. 06:30-21:00. Observera att bussen inte kommer att trafikera Hagsätra, Rågsved, Enskede gård, Globen eller Gullmarsplan. \n\nI Cityterminalen avgår buss 169C från hållplats C (utgång genom gate B).\n\nÖvriga lokala bussar\n-Buss 143 går mellan Hökarängen och Älvsjö station via Gubbängen och Hagsätra, för byte till pendeltåg.\n-Blåbuss 172 går mellan Rågsved/Högdalen och Hökarängen, för byte till tunnelbanans gröna linje 18.\n-Blåbuss 173 går mellan Bandhagen och Hökarängen för byte till tunnelbanans gröna linje 18, alternativt mellan Bandhagen och Älvsjö station för byte till pendeltåg.\n\nGångavstånd:\nAtt promenera ca 700 meter mellan Globen och Gullmarsplan tar ungefär 8-10 minuter.\n\nPlanera din resa i god tid. Använd sökverktyget på sl.se eller appar för att hitta din bästa resväg.\n\n\n\n\n","Scope":"Gröna linjen","DevCaseGid":9076001025802433,"DevMessageVersionNumber":6,"ScopeElements":"Tunnelbanans gröna linje 19","FromDateTime":"2020-06-22T11:06:30.037","UpToDateTime":"2020-08-17T02:00:00","Updated":"2020-06-22T11:06:30.05+02:00"},{"Created":"2020-06-08T11:28:47.72+02:00","MainNews":false,"SortOrder":1,"Header":"Reducerad trafik mellan Alvik och Hässelby strand sena kvällar i juni och juli","Details":"Reducerad trafik mellan Alvik och Hässelby strand sena kvällar i juni och juli, på grund av spårarbete vid Åkeshov. Vid varje tillfälle är trafiken glesare från kl. 22:10 och fram till trafikens slut. Varje morgon går tågen som vanligt och arbetet har ingen påverkan under dagtid.\nArbetet pågår under dessa dagar och tider:\n- Söndag den 21 juni, från kl. 22:10 fram till trafikens slut.\n\n-Söndag till och med torsdag 19-23 juli, varje kväll från kl. 22:10 fram till trafikens slut.\n\nArbetet medför att tågen i båda riktningarna samsas på endast ett av normalt två spår mellan Abrahamsberg och Hässelby strand. Spåret som normalt används för tågen i riktning mot T-Centralen är avstängt.\nSå här påverkas trafiken berörda kvällar:\n\nReducerad trafik och förlängda restider\nMellan Alvik och Hässelby strand går tågen var 30:e minut. Din resa kan ta längre tid än normalt då tågen kör med sänkt hastighet mellan Abrahamsberg och Hässelby strand. Vid Stora mossen och Vällingby kan tågen bli stående längre tid än vanligt, för att vänta in mötande tåg. Färre avgångar än vanligt förekommer också mellan Hötorget och Alvik, enstaka tåg som normalt  skulle ha gått till Hässelby strand vänder i stället vid Hötorget.\n\nFörändrade avgångstider\nMellan Skarpnäck/Farsta strand och Hässelby strand ändras avgångstiderna för samtliga resor efter kl. 22:10. Sök din resa i förväg för att se gällande avgångstider.\n\nSpårändringar\nFrån Abrahamsberg till och med Hässelby strand avgår tågen i båda riktningarna från spår 1 (normalt spår för tågen i riktning mot Hässelby strand).\nVällingby: Samtliga tåg avgår från plattformen där tågen mot Hässelby strand normalt går. Tågen mot Hässelby strand från spår 1 (ordinarie spår) och tågen mot T-Centralen från spår 2 (mellanspåret).\nJohannelund: Samtliga tåg avgår från plattformen där tågen mot Hässelby strand normalt går.\n\nAlternativa resvägar\nFlera lokala busslinjer går från stationerna längs Hässelbygrenen till andra tunnelbanestationer eller till pendeltågsstationer:\n-Buss 113 mellan Islandstorget och Sundbyberg, Huvudsta, Västra skogen och Solna centrum.\n-Buss 115 mellan Vällingby och Råcksta.\n-Buss 116 mellan Vällingby och Spånga station.\n-Buss 118 mellan Vällingby och Spånga station, Rissne och Hallonbergen.\n-Buss 124 mellan Abrahamsberg och Alvik.\n-Buss 179 mellan Vällingby och Spånga station, Tensta/Rinkeby, Kista och Sollentuna.\n-Buss 541 mellan Vällingby och Barkarby/Jakobsberg station.\n\n\nRäkna med förlängda restider och att du behöver vara uppmärksam på spårändringarna och tågens destinationer. Använd sökverktyget på sl.se eller i valfri app för att hitta din bästa resväg.","Scope":"Gröna linjen","DevCaseGid":9076001025799902,"DevMessageVersionNumber":2,"ScopeElements":"Tunnelbanans gröna linje 17, 18, 19","FromDateTime":"2020-06-21T22:10:00","UpToDateTime":"2020-07-24T02:00:00","Updated":"2020-06-08T11:28:47.72+02:00"},{"Created":"2020-06-29T11:16:02.517+02:00","MainNews":false,"SortOrder":1,"Header":"Avstängd hiss ","Details":"Hissarna vid Vällingby till och från plattformen mot Hässelby strand är avstängda på grund av tekniskt fel. Resenärerna i behov av hiss hänvisas till plattformen mot Skarpnäck och Farsta strand eller till  angränsande stationer samt till tillgängligshetsgarantin. \n\nVi saknar prognos på när hissen åter kan tas i bruk.\n\n\n","Scope":"Vällingby","DevCaseGid":9076001027036883,"DevMessageVersionNumber":3,"ScopeElements":"Tunnelbanans gröna linje 17, 18, 19","FromDateTime":"2020-06-29T11:16:02.5","UpToDateTime":"2020-07-06T23:30:00","Updated":"2020-06-29T11:16:02.517+02:00"}]};
//    this.setState({'disturbanceData': fakeData['ResponseData']});
    axios({
      method: 'GET',
      url: 'http://127.0.0.1:80/tun',
      headers: {
        'Tunnel': url
      }
    })
    .then(res => {
      this.setState({'disturbanceData': res.data['ResponseData']});
    });
  }

  render() {
    return (
      <div className='sl-disturbance'>
        <button onClick={this.fetchSLDisturbances}>SL Disturbances</button>
        {this.state.disturbanceData.map((item, iterator) => 
          <div key={iterator}>
            <div className='header'>
              <h4>{item['Scope']} - {item['Header']}</h4>
            </div>
            <div className='details'>
              <p>{item['Details']}</p>
            </div>
          </div>
        )}
      </div>
    );
  }
}

export default SLDisturbance;