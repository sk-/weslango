package cld3

import (
	"fmt"
	"testing"
)

func TestOkay(t *testing.T) {
	det, err := New(0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	defer Free(det)
	res := det.FindLanguage("Hey, this is an english sentence")
	if res.Language != "eng" {
		t.Errorf("Unexpected language: wanted \"eng\", got %#v", res.Language)
	}
	if !res.IsReliable {
		t.Errorf("Answer is not reliable")
	}
}

func TestUnknown(t *testing.T) {
	det, err := New(10, 1000)
	if err != nil {
		t.Fatal(err)
	}
	defer Free(det)
	res := det.FindLanguage("Hey")
	if res.Language != "" {
		t.Errorf("Language was supposed to be unknown, got %#v", res.Language)
	}
}

func TestErrors(t *testing.T) {
	cases := []struct {
		min int
		max int
		err error
	}{
		{0, 0, ErrMaxLessThanOrEqToZero},
		{-1, 2, ErrMinLessThanZero},
		{1, -1, ErrMaxLessThanOrEqToZero},
		{2, 1, ErrMaxSmallerOrEqualToMin},
	}
	for _, c := range cases {
		_, err := New(c.min, c.max)
		if err != c.err {
			t.Errorf("Unexpected error: wanted %s, got %s", c.err, err)
		}
	}
}

func ExampleBasic() {
	det, err := New(0, 512)
	if err != nil {
		fmt.Println("whoops, couldn't create a new LanguageIdentifier:", err)
	}
	defer Free(det)
	res := det.FindLanguage("Hey, this is an english sentence")
	if res.IsReliable {
		fmt.Println("pretty sure we've got text written in", res.Language)
	}
	res = det.FindLanguage("Muy bien, gracias.")
	if res.IsReliable {
		fmt.Println("ah, and this one is", res.Language)
	}
	// Output:
	// pretty sure we've got text written in eng
	// ah, and this one is spa
}

const (
	afr    = "Dit is 'n kort stukkie van die teks wat gebruik sal word vir die toets van die akkuraatheid van die nuwe benadering."
	amh    = "እኔ ግን ሩቅ ለሆኑ ነገሮች ዘላለማዊ ማሳከክ ተሰቃያለሁ ፡፡ የተከለከሉ ባሕሮችን መጓዝ እና በባህር ዳርቻዎች ላይ መጓዝ እወዳለሁ ፡፡"
	ara    = "احتيالية بيع أي حساب"
	aze    = " a az qalıb breyn rinq intellektual oyunu üzrə yarışın zona mərhələləri  keçirilib miq un qalıqlarının dənizdən çıxarılması davam edir məhəmməd  peyğəmbərin karikaturalarını çap edən qəzetin baş redaktoru iş otağında  ölüb"
	bel    = " а друкаваць іх не было тэхнічна магчыма бліжэй за вільню тым самым часам  нямецкае кіраўніцтва прапаноўвала апроч ўвядзення лацінкі яе"
	bul    = " а дума попада в състояние на изпитание ключовите думи с предсказана  малко под то изискване на страниците за търсене в"
	bulLat = "a duma popada v sŭstoyanie na izpitanie klyuchovite dumi s predskazana  malko pod to iziskvane na stranitsite za tŭrsene v"
	ben    = "গ্যালারির ৩৮ বছর পূর্তিতে মূল্যছাড় অর্থনীতি বিএনপির ওয়াক আউট তপন  চৌধুরী হারবাল অ্যাসোসিয়েশনের সভাপতি আন্তর্জাতিক পরামর্শক  বোর্ড দিয়ে শরিয়াহ্ ইনন্ডেক্স করবে  সিএসই মালিকপক্ষের কান্না, শ্রমিকের অনিশ্চয়তা মতিঝিলে সমাবেশ নিষিদ্ধ:  এফবিসিসিআইয়ের ধন্যবাদ বিনোদন বিশেষ প্রতিবেদন বাংলালিংকের গ্র্যান্ডমাস্টার  সিজন-৩ ব্রাজিলে বিশ্বকাপ ফুটবল আয়োজনবিরোধী বিক্ষোভ দেশের নিরাপত্তার   চেয়ে অনেক বেশি সচেতন । প্রার্থীদের দক্ষতা  ও যোগ্যতার  পাশাপাশি তারা জাতীয় ইস্যুগুলোতে প্রাধান্য দিয়েছেন । ” পাঁচটি সিটিতে ২০  লাখ ভোটারদের দিয়ে জাতীয় নির্বাচনে ৮ কোটি ভোটারদের  সঙ্গে তুলনা করা যাবে কি একজন দর্শকের এমন প্রশ্নে জবাবে আব্দুল্লাহ  আল নোমান বলেন , “ এই পাঁচটি সিটি কর্পোরেশন নির্বাচন দেশের পাঁচটি বড়  বিভাগের প্রতিনিধিত্ব করছে । এছাড়া এখানকার ভোটার রা সবাই সচেতন । তারা"
	bos    = "Za današnje područje Visočkog polja se pretpostavlja da je u 10. vijeku predstavljalo zametak u razvoju srednjovjekovne bosanske države koju spominje Bizantski car Konstantin Porfirogenit. Proširena dolina rijeke Bosne oko današnjeg Visokog je bila središte najvećeg poljoprivrednog regiona u srednjoj Bosni, stoga je prostrano i plodno visočko polje bilo idealno za razvoj političkog centra. Naselje smješteno u Visočkom polju i njegovoj okolini, dugo je imalo naziv Bosna što je predstavljao najstariji i najuži sadržaj pojma Bosne. Visočka dolina je imala više kraljevskih dvorova, i bila je jedno od najvažnijih političkih centara srednjovjekovne države Bosne. U Milima se krunisao prvi bosanski kralj Tvrtko Kotromanić. Stari grad Visoki koji se nalazi na brdu Visočica je bio politički važna tvrđava, a njegovo podgrađe Podvisoki je bio jedan od najranijih primjera srednjovjekovne urbane sredine na užem području Bosne. "
	cat    = "al final en un únic lloc nhorabona l correu electrònic està concebut com  a eina de productivitat aleshores per què perdre el temps arxivant  missatges per després intentar recordar on els veu desar i per què heu d  eliminar missatges importants per l"
	ceb    = "Ang Sugbo usa sa mga labing ugmad nga lalawigan sa nasod. Kini ang sentro  sa komersyo, edukasyon ug industriya sa sentral ug habagatang dapit sa  kapupod-an. Ang mipadayag sa Sugbo isip ikapito nga labing nindot nga  pulo sa , ang nag-inusarang pulo sa Pilipinas nga napasidunggan sa maong  magasin sukad pa sa tuig"
	cos    = "In quantu à mè, mi sentu turmentatu cù una piccula eterna per e cose remoti. Mi piace à navigà mari proibiti, è a terra nant'à i costi barbari."
	ces    = " a akci opakujte film uložen vykreslit gmail tokio smazat obsah adresáře  nelze načíst systémový profil jednotky smoot okud používáte pro určení  polokoule značky z západ nebo v východ používejte nezáporné hodnoty  zeměpisné délky nelze"
	cym    = " a chofrestru eich cyfrif ymwelwch a unwaith i chi greu eich cyfrif mi  fydd yn cael ei hysbysu o ch cyfeiriad ebost newydd fel eich bod yn gallu  cadw mewn cysylltiad drwy gmail os nad ydych chi wedi clywed yn barod am  gmail mae n gwasanaeth gwebost"
	dan    = " a z tallene og punktummer der er tilladte log ud angiv den ønskede  adgangskode igen november gem personlige oplysninger kontrolspørgsmål det  sidste tegn i dit brugernavn skal være et bogstav a z eller tal skriv de  tegn du kan se i billedet nedenfor"
	deu    = " abschnitt ordner aktivieren werden die ordnereinstellungen im  farbabschnitt deaktiviert öchten sie wirklich fortfahren eldtypen angeben  optional n diesem schritt geben sie für jedesfeld aus dem datenset den  typ an ieser schritt ist optional eldtypen"
	ell    = " ή αρνητική αναζήτηση λέξης κλειδιού καταστήστε τις μεμονωμένες λέξεις  κλειδιά περισσότερο στοχοθετημένες με τη μετατροπή τους σε"
	ellLat = " O Balzák égrapse to Loí Lampér to kalokaíri tou 1832,  tin período pou émine me tous phílous tou ston Pírgo tou Sassé.  Dimosiéftike se tris sinoliká ekdósis me tris diaphoretikoús títlous.  To mithistórima periékhi mia elákhisti mónon plokí kai sto megalítero méros tou díni émphasi stis metaphisikés idées tou idiophioús tou protagonistí kai tou monadikoú tou phílou (o opíos,  teliká,  phaínetai na ínai o ídios o Balzák). "
	eng    = " a backup credit card by visiting your billing preferences page or visit  the adwords help centre for more details https adwords google com support  bin answer py answer hl en we were unable to process the payment of for  your outstanding google adwords"
	epo    = " a jarcento refoje per enmetado de koncerna pastro tiam de reformita  konfesio ekde refoje ekzistis luteranaj komunumanoj tamen tiuj fondis  propran komunumon nur en ambaŭ apartenis ekde al la evangela eklezio en  prusio resp ties rejnlanda provinceklezio en"
	spa    = " a continuación haz clic en el botón obtener ruta también puedes  desplazarte hasta el final de la página para cambiar tus opciones de  búsqueda gráfico y detalles ésta es una lista de los vídeos que te  recomendamos nuestras recomendaciones se basan"
	est    = " a niipea kui sinu maksimaalne igakuine krediidi limiit on meie poolt  heaks kiidetud on sinu kohustuseks see krediidilimiit"
	eus    = " a den eraso bat honen kontra hortaz eragiketa bakarrik behar dituen  eraso batek aes apurtuko luke nahiz eta oraingoz eraso bideraezina izan  gaur egungo teknologiaren mugak direla eta oraingoz kezka hauek alde  batera utzi daitezke orain arteko indar"
	fas    = " آب خوردن عجله می کردند به جای باز ی کتک کاری می کردند و همه چيز مثل قبل  بود فقط من ماندم و يک دنيا حرف و انتظار تا عاقبت رسيد احضاريه ی ای با"
	fin    = " a joilla olet käynyt tämä kerro meille kuka ä olet ei tunnistettavia  käyttötietoja kuten virheraportteja käytetään google desktopin  parantamiseen etsi näyttää mukautettuja uutisia google desktop  keskivaihto leikkaa voit kaksoisnapsauttaa"
	fil    = "Ito ay isang maikling piraso ng teksto na ito ay gagamitin para sa  pagsubok ang kawastuhan ng mga bagong diskarte."
	fra    = " a accès aux collections et aux frontaux qui lui ont été attribués il  peut consulter et modifier ses collections et exporter des configurations  de collection toutefois il ne peut pas créer ni supprimer des collections  enfin il a accès aux fonctions"
	fry    = "Wat my oanbelanget, bin ik pynige mei in ivige jeuk foar dingen op ôfstân. Ik hâld fan farre ferbeane seeën, en lân op barbaarske kusten."
	gle    = " a bhfuil na focail go léir i do cheist le fáil orthu ní gá ach focail  breise a chur leis na cinn a cuardaíodh cheana chun an cuardach a  bheachtú nó a chúngú má chuirtear focal breise isteach aimseofar fo aicme  ar leith de na torthaí a fuarthas"
	gla    = "Dhòmhsa, tha mi air mo shàrachadh le itch shìorraidh airson rudan iomallach. Is toigh leam a bhith a ’seòladh cuantan toirmisgte, agus a’ tighinn air tìr air oirthirean borb."
	glg    = "  debe ser como mínimo taranto tendas de venda polo miúdo cociñas  servizos bordado canadá viaxes parques de vehículos de recreo hotel  oriental habitación recibir unha postal no enderezo indicado  anteriormente"
	guj    = " આના પરિણામ પ્રમાણસર ફોન્ટ અવતરણ ચિન્હવાળા પાઠને છુપાવો બધા સમૂહો શોધાયા  હાલનો જ સંદેશ વિષયની"
	hau    = " a cikin a kan sakamako daga sakwannin a kan sakamako daga sakwannin daga  ranar zuwa a kan sakamako daga guda daga ranar zuwa a kan sakamako daga  shafukan daga ranar zuwa a kan sakamako daga guda a cikin last hour a kan  sakamako daga guda daga kafar"
	haw    = "ʻO wau nei, ʻeha wau me ka ʻeha mau loa no nā mea mamao. Makemake wau e holo i nā kai i pāpā ʻia, a i ka ʻāina ma nā kahakai."
	hin    = " ं ऐडवर्ड्स विज्ञापनों के अनुभव पर आधारित हैं और इनकी मदद से आपको अपने  विज्ञापनों का अधिकतम लाभ"
	hinLat = "n aidavards vigyaapanon ke anubhav par aadhaarit hain aur inakee madad se aapako apane  vigyaapanon ka adhikatam laabh"
	hmn    = "Qhov no yog ib tug luv luv daim ntawv nyeem uas yuav siv tau rau kev soj  ntsuam qhov tseeb ntawm tus tshiab mus kom ze."
	hrv    = "Posljednja dva vladara su Kijaksar (Κυαξαρης 625-585 prije Krista),  fraortov sin koji će proširiti teritorij Medije i Astijag. Kijaksar je  imao kćer ili unuku koja se zvala Amitis a postala je ženom  Nabukodonosora II. kojoj je ovaj izgradio Viseće vrtove Babilona.  Kijaksar je modernizirao svoju vojsku i uništio Ninivu 612. prije Krista.  Naslijedio ga je njegov sin, posljednji medijski kralj, Astijag, kojega  je detronizirao (srušio sa vlasti) njegov unuk Kir Veliki. Zemljom su  zavladali Perzijanci. Hrvatska je zemlja situacija u Europi. Ona ima  bogatu kulturu i ukusna jela."
	hat    = " ak pitit tout sosyete a chita se pou sa leta dwe pwoteje yo nimewo leta  fèt pou li pwoteje tout paran ak pitit nan peyi a menm jan kit paran yo  marye kit yo pa marye tout manman ki fè pitit leta fèt pou ba yo konkoul  menm jan tou pou timoun piti ak pou"
	hun    = " a felhasználóim a google azonosító szöveget ikor látják a felhasználóim  a google azonosító szöveget felhasználók a google azonosító szöveget  fogják látni minden tranzakció után ha a vásárlását regisztrációját  oldalunk"
	hye    = " ա յ եվ նա հիացած աչքերով նայում է հինգհարկանի շենքի տարօրինակ փոքրիկ  քառակուսի պատուհաններին դեռ մենք շատ ենք հետամնաց ասում է նա այսպես է"
	ind    = "Sistem imun adalah sebuah sistem yang terdiri dari sel dan banyak struktur biologis lainnya yang bertanggung jawab atas pertahanan suatu organisme dari penyakit. Untuk berfungsi dengan baik, sistem imun mengenali dan membunuh berbagai macam pengaruh biologis luar seperti dari infeksi, bakteri, virus sampai parasit, agar sel dan jaringan organisme yang sehat tetap dapat berfungsi dengan normal. Patogen dapat berevolusi dan beradaptasi agar terhindar dari penghancuran oleh sistem imun, tetapi mekanisme pertahanan tubuh juga berevolusi untuk menetralkannya. "
	ibo    = "Chineke bụ aha ọzọ ndï omenala Igbo kpọro Chukwu. Mgbe ndị bekee bịara,  ha mee ya nke ndi Christian. N'echiche ndi ekpere chi Omenala Ndi Igbo,  Christianity, Judaism, ma Islam, Chineke nwere ọtụtụ utu aha, ma nwee  nanị otu aha. Ụzọ abụọ e si akpọ aha ahụ bụ Jehovah ma Ọ bụ Yahweh. Na  ọtụtụ Akwụkwọ Nsọ, e wepụla aha Chineke ma jiri utu aha bụ Onyenwe Anyị  ma ọ bụ Chineke dochie ya. Ma mgbe e dere akwụkwọ nsọ, aha ahụ bụ Jehova  pụtara n’ime ya, ihe dị ka ugboro pụkụ asaa(7,000)."
	isl    = " a afköst leitarorða þinna leitarorð neikvæð leitarorð auglýsingahópa  byggja upp aðallista yfir ný leitarorð fyrir auglýsingahópana og skoða  ítarleg gögn um árangur leitarorða eins og samkeppni auglýsenda og  leitarmagn er krafist notkun"
	ita    = " a causa di un intervento di manutenzione del sistema fino alle ore circa  ora legale costa del pacifico del novembre le campagne esistenti  continueranno a essere pubblicate come di consueto anche durante questo  breve periodo di inattività ci scusiamo per"
	heb    = " או לערוך את העדפות ההפצה אנא עקוב אחרי השלבים הבאים כנס לחשבון האישי שלך  ב"
	jpn    = " このペ ジでは アカウントに指定された予算の履歴を一覧にしています  それぞれの項目には 予算額と特定期間のステ タスが表示されます  現在または今後の予算を設定するには"
	jpnLat = "Kono pe jide wa akaunto ni shitei sa reta yosan no rireki o ichiran ni shite imasu sorezore no kōmoku ni wa yosan-gaku to tokutei kikan no Sute tasu ga hyōji sa remasu genzai matawa kongo no yosan o settei suru ni wa"
	jav    = "Iki Piece cendhak teks sing bakal digunakake kanggo Testing akurasi  pendekatan anyar."
	kat    = " ა ბირთვიდან მიღებული ელემენტი მენდელეევის პერიოდულ სიტემაში  გადაინაცვლებს ორი უჯრით"
	kaz    = " а билердің өзіне рұқсат берілмеген егер халық талап етсе ғана хан  келісім берген өздеріңіз білесіздер қр қыл мыс тық кодексінде жазаның"
	khm    = "នេះគឺជាបំណែកខ្លីនៃអត្ថបទដែលនឹងត្រូវបានប្រើសម្រាប់ការធ្វើតេស្តភាពត្រឹមត្រូវ នៃវិធីសាស្រ្តថ្មីនេះ។"
	kan    = " ಂಠಯ್ಯನವರು ತುಮಕೂರು ಜಿಲ್ಲೆಯ ಚಿಕ್ಕನಾಯಕನಹಳ್ಳಿ ತಾಲ್ಲೂಕಿನ ತೀರ್ಥಪುರ ವೆಂಬ ಸಾಧಾರಣ  ಹಳ್ಳಿಯ ಶ್ಯಾನುಭೋಗರ"
	kor    = " 개별적으로 리포트 액세스 권한을 부여할 수 있습니다 액세스 권한  부여사용자에게 프로필 리포트에 액세스할 수 있는 권한을 부여하시려면 가용  프로필 상자에서 프로필 이름을 선택한 다음"
	kur    = "Wekî ku ji min re ye, ez ji xezîneyên herheyî ji bo tiştên dûr ve tengahiyê dikim. Ez hez dikim ku li peravên qedexekirî bar bikim, û li peravên barbar bibim."
	kir    = "Өзүм жөнүндө айтсам, мен алыстан нерселер үчүн түбөлүк, таз менен кыйналып жатам. Мен тыюу дарыя менен сүзүп, жана бөлөк элдер жер жакшы."
	lat    = " a deo qui enim nocendi causa mentiri solet si iam consulendi causa  mentiatur multum profecit sed aliud est quod per se ipsum laudabile  proponitur aliud quod in deterioris comparatione praeponitur aliter enim  gratulamur cum sanus est homo aliter cum melius"
	ltz    = "Wat mech ugeet, sinn ech gepflanzt mat engem éiwege Jucken fir Saachen op Fern. Ech hu gär verbueden Mier ze segelen, a landen op barbaresche Küsten."
	lao    = " ກຫາທົ່ວທັງເວັບ ແລະໃນເວັບໄຮ້ສາຍ ທຳອິດໃຫ້ທຳການຊອກຫາກ່ອນ ຈາກນັ້ນ  ໃຫ້ກົດປຸ່ມເມນູ ໃນໜ້າຜົນໄດ້"
	lit    = " a išsijungia mano idėja dėl geriausio laiko po pastarųjų savo santykių  pasimokiau penki dalykai be kurių negaliu gyventi mano miegamajame tu  surasi ideali pora išsilavinimas aukštoji mokykla koledžas universitetas  pagrindinis laipsnis metai"
	lav    = " a gadskārtējā izpārdošana slēpošana jāņi atlaide izmaiņas trafikā kas  saistītas ar sezonas izpārdošanu speciālajām atlaidēm u c ir parastas un  atslēgvārdi kas ir populāri noteiktos laika posmos šajā laikā saņems  lielāku klikšķu"
	mlg    = " amporisihin i ianao mba hijery ny dika teksta ranofotsiny an ity  lahatsoratra ity tsy ilaina ny opérateur efa karohina daholo ny teny  rehetra nosoratanao ampiasao anaovana dokambarotra i google telugu datin  ny takelaka fikarohana sary renitakelak i"
	mri    = " haere ki te kainga o o haere ki te kainga o o haere ki te kainga o te  rapunga ahua o haere ki te kainga o ka tangohia he ki to rapunga kaore au  mohio te tikanga whakatiki o te ra he whakaharuru te pai rapunga a te  rapunga ahua a e kainga o nga awhina o te"
	mkd    = " гласовите коалицијата на вмро дпмне како партија со најмногу освоени  гласови ќе добие евра а на сметката на коализијата за македонија"
	mal    = " ം അങ്ങനെ ഞങ്ങള് അവരുടെ മുമ്പില് നിന്നു ഔടും ഉടനെ നിങ്ങള് പതിയിരിപ്പില്  നിന്നു എഴുന്നേറ്റു"
	mon    = " а боловсронгуй болгох орон нутгийн ажил үйлсийг уялдуулж зохицуулах  дүрэм журам боловсруулах орон нутгийн өмч хөрөнгө санхүүгийн"
	mar    = "हैदराबाद  उच्चार ऐका (सहाय्य·माहिती)तेलुगू: హైదరాబాదు , उर्दू:  حیدر آباد हे भारतातील आंध्र प्रदेश राज्याच्या राजधानीचे शहर  आहे. हैदराबादची लोकसंख्या ७७ लाख ४० हजार ३३४ आहे. मोत्यांचे शहर  अशी एकेकाळी ओळख असलेल्या या शहराला ऐतिहासिक, सांस्कृतिक आणि  स्थापत्यशास्त्रीय वारसा लाभला आहे. १९९० नंतर शिक्षण आणि माहिती तंत्रज्ञान  त्याचप्रमाणे औषधनिर्मिती आणि जैवतंत्रज्ञान क्षेत्रातील उद्योगधंद्यांची  वाढ शहरात झाली. दक्षिण मध्य भारतातील पर्यटन आणि तेलुगू चित्रपटनिर्मितीचे  हैदराबाद हे केंद्र आहे"
	msa    = "pengampunan beramai-ramai supaya mereka pulang ke rumah masing-masing.  Orang-orang besarnya enggan mengiktiraf sultan yang dilantik oleh Belanda  sebagai Yang DiPertuan Selangor. Orang ramai pula tidak mahu menjalankan  perniagaan bijih timah dengan Belanda, selagi raja yang berhak tidak  ditabalkan. Perdagang yang lain dibekukan terus kerana untuk membalas  jasa beliau yang membantu Belanda menentang Riau, Johor dan Selangor. Di  antara tiga orang Sultan juga dipandang oleh rakyat sebagai seorang  sultan yang paling gigih. 1 | 2 SULTAN Sebagai ganti Sultan Ibrahim  ditabalkan Raja Muhammad iaitu Raja Muda. Walaupun baginda bukan anak  isteri pertama bergelar Sultan Muhammad bersemayam di Kuala Selangor  juga. Pentadbiran baginda yang lemah itu menyebabkan Kuala Selangor  menjadi sarang ioleh Cina di Lukut tidak diambil tindakan, sedangkan  baginda sendiri banyak berhutang kepada 1"
	mlt    = " ata ikteb messaġġ lil indirizzi differenti billi tagħżilhom u tagħfas il  buttuna ikteb żid numri tfittxijja tal kotba mur print home kotba minn  pagni ghal pagna minn ghall ktieb ta aċċessa stieden habib iehor grazzi  it tim tal gruppi google"
	mya    = " တက္ကသုိလ္ မ္ဟ ပ္ရန္ လာ္ရပီးေနာက္ န္ဟစ္ အရ္ဝယ္ ဦးသန္ ့သည္ ပန္  းတနော္ အမ္ယုိးသား ေက္ယာင္ း"
	nep    = "अरू ठाऊँबाटपनि खुलेको छ यो खाता अर अरू ठाऊँबाटपनि खुलेको छ यो खाता अर ू"
	nld    = " a als volgt te werk om een configuratiebestand te maken sitemap gen py  ebruik filters om de s op te geven die moeten worden toegevoegd of  uitgesloten op basis van de opmaaktaal elke sitemap mag alleen de s  bevatten voor een bepaalde opmaaktaal dit"
	nor    = " a er obligatorisk tidsforskyvning plassering av katalogsøk  planinformasjon loggfilbane gruppenavn kontoinformasjon passord domene  gruppeinformasjon alle kampanjesporing alternativ bruker grupper  oppgaveplanlegger oppgavehistorikk kontosammendrag antall"
	nya    = "Boma ndi gawo la dziko lomwe linapangidwa ndi cholinga chothandiza  ntchito yolamulira. Kuŵalako kulikuunikabe mandita, Edipo nyima  unalephera kugonjetsa kuŵalako."
	pan    = " ਂ ਦਿਨਾਂ ਵਿਚ ਭਾਈ ਸਾਹਿਬ ਦੀ ਬੁੱਚੜ ਗੋਬਿੰਦ ਰਾਮ ਨਾਲ ਅੜਫਸ ਚੱਲ ਰਹੀ ਸੀ ਗੋਬਿੰਦ  ਰਾਮ ਨੇ ਭਾਈ ਸਾਹਿਬ ਦੀਆਂ ਭੈਣਾ"
	pol    = " a australii będzie widział inne reklamy niż użytkownik z kanady  kierowanie geograficzne sprawia że reklamy są lepiej dopasowane do  użytkownika twojej strony oznacza to także że możesz nie zobaczyć  wszystkich reklam które są wyświetlane na"
	pus    = "لکه څنګه چې زما لپاره ، زه د لرې پرتو شیانو لپاره د تل پاتې خارش سره شکنجه شوی یم. زه د منع شوي بحریو سیل کولو سره مینه لرم ، او وحشي ساحلونو ته ځم."
	por    = " a abit prevê que a entrada desses produtos estrangeiros no mercado  têxtil e vestuário do brasil possa reduzir os preços em cerca de a partir  de má notícia para os empresários que terão que lutar para garantir suas  margens de lucro mas boa notícia"
	ron    = " a anunţurilor reţineţi nu plătiţi pentru clicuri sau impresii ci numai  atunci când pe site ul dvs survine o acţiune dorită site urile negative  nu pot avea uri de destinaţie daţi instrucţiuni societăţii dvs bancare  sau constructoare să"
	rus    = " а неправильный формат идентификатора дн назад"
	rusLat = "a nepravil'nyy format identifikatora dn nazad"
	snd    = "مون لاء، مون کي ٻاهرين شين لاء دائمي چرچ سان عذاب ڪيو ويو آهي. مون کي حرام درياء حرام ڪرڻ، ۽ جڙڙيل رستن تي زمين ڏيڻ."
	sin    = " අනුරාධ මිහිඳුකුල නමින් සකුරා ට ලිපියක් තැපෑලෙන් එවා තිබුණා කි  ් රස්ටි ෂෙල්ටන් ප ් රනාන්දු ද"
	slk    = " a aktivovať reklamnú kampaň ak chcete kampaň pred spustením ešte  prispôsobiť uložte ju ako šablónu a pokračujte v úprave vyberte si jednu  z možností nižšie a kliknite na tlačidlo uložiť kampaň nastavenia kampane  môžete ľubovoľne"
	slv    = " adsense stanje prijave za google adsense google adsense račun je bil  začasno zamrznjen pozdravljeni hvala za vaše zanimanje v google adsense  po pregledu vaše prijavnice so naši strokovnjaki ugotovili da spletna  stran ki je trenutno povezana z vašim"
	smo    = "A o aʻu, o loʻo faʻasauā aʻu i se mea e faʻavavau mo mea mamao. Ou te fiafia e folau vaʻa faasaina, ma lauʻeleʻele i luga o le sami."
	sna    = "Kana ndirini, ndinotambudzwa neicho chisingaperi nekuda kwezvinhu zviri kure. Ini ndinofarira kufamba muchikepe chakarambidzwa, uye nyika pamhenderekedzo isina mhaka."
	som    = " a oo maanta bogga koobaad ugu qoran yahey beesha caalamka laakiin si  kata oo beesha caalamku ula guntato soomaaliya waxa aan shaki ku jirin in  aakhirataanka dadka soomaalida oo kaliya ay yihiin ku soomaaliya ka saari  kara dhibka ay ku jirto"
	sqi    = " a do të kërkoni nga beogradi që të njohë pavarësinë e kosovës zoti thaçi  prishtina është gati ta njoh pavarësinë e serbisë ndërsa natyrisht se do  të kërkohet një gjë e tillë që edhe beogradi ta njoh shtetin e pavarur  dhe sovran të"
	srp    = "балчак балчак на мапи србије уреди демографија у насељу балчак живи  пунолетна становника а просечна старост становништва износи година"
	sot    = " bang ba nang le thahasello matshwao a sehlooho thuto e thehilweng hodima  diphetho ke tsela ya ho ruta le ho ithuta e totobatsang hantle seo  baithuti ba lokelang ho se fihlella ntlhatheo eo e sebetsang ka yona ke  ya hore titjhere o hlakisa pele seo"
	sun    = "Nu ngatur kahirupan warga, keur kapentingan pamarentahan diatur ku RT, RW  jeung Kepala Dusun, sedengkeun urusan adat dipupuhuan ku Kuncen jeung  kepala adat. Sanajan Kampung Kuta teu pati anggang jeung lembur sejenna  nu aya di wewengkon Desa Pasir Angin, tapi boh wangunan imah atawa  tradisi kahirupan masarakatna nenggang ti nu lian."
	swe    = " a bort objekt från google desktop post äldst meny öretag dress etaljer  alternativ för vad är inne yaste google skrivbord plugin program för  nyheter google visa nyheter som är anpassade efter de artiklar som du  läser om du till exempel läser"
	swa    = " a ujumbe mpya jumla unda tafuta na angalia vikundi vya kujadiliana na  kushiriki mawazo iliyopangwa kwa tarehe watumiaji wapya futa orodha hizi  lugha hoja vishikanisho vilivyo dhaminiwa ujumbe sanaa na tamasha toka  udhibitisho wa neno kwa haraka fikia"
	tam    = " அங்கு ராஜேந்திர சோழனால் கட்டப்பட்ட பிரம்மாண்டமான சிவன் கோவில் ஒன்றும்  உள்ளது தொகு"
	tel    = " ఁ దనర జయించిన తత్వ మరసి చూడఁ దాన యగును రాజయోగి యిట్లు తేజరిల్లుచు నుండు  విశ్వదాభిరామ వినర వేమ"
	tgk    = " адолат ва инсондӯстиро бар фашизм нажодпарастӣ ва адоват тарҷеҳ додааст  чоп кунед ба дигарон фиристед чоп кунед ба дигарон фиристед"
	tha    = " กฏในการค้นหา หรือหน้าเนื้อหา หากท่านเลือกลงโฆษณา  ท่านอาจจะปรับต้องเพิ่มงบประมาณรายวันตา"
	tur    = " a ayarlarınızı görmeniz ve yönetmeniz içindir eğer kampanyanız için  günlük bütçenizi gözden geçirebileceğiniz yeri arıyorsanız kampanya  yönetimi ne gidin kampanyanızı seçin ve kampanya ayarlarını düzenle yi  tıklayın sunumu"
	ukr    = " а більший бюджет щоб забезпечити собі максимум прибутків від переходів  відстежуйте свої об яви за датою географічним розташуванням"
	urd    = " آپ کو کم سے کم ممکنہ رقم چارج کرتا ہے اس کی مثال کے طور پر فرض کریں اگر  آپ کی زیادہ سے زیادہ قیمت فی کلِک امریکی ڈالر اور کلِک کرنے کی شرح ہو تو"
	uzb    = " abadiylashtirildi aqsh ayol prezidentga tayyormi markaziy osiyo afg  onistonga qanday yordam berishi mumkin ukrainada o zbekistonlik  muhojirlar tazyiqdan shikoyat qilmoqda gruziya va ukraina hozircha natoga  qabul qilinmaydi afg oniston o zbekistonni g"
	vie    = " adsense cho nội dung nhà cung cấp dịch vụ di động xác minh tín  dụng thay đổi nhãn kg các ô xem chi phí cho từ chối các đơn đặt  hàng dạng cấp dữ liệu ác minh trang web của bạn để xem"
	xho    = "Mna ke, ndihlushwe yitrone yaphakade ngenxa yezinto ezikude. Ndiyakuthanda ukuhamba ngolwandle ngokungavumelekanga, kunye nomhlaba kwiindlela ezixineneyo."
	yid    = "אן פאנטאזיע ער איז באקאנט צים מערסטן פאר זיינע באַלאַדעס ער האָט געוווינט  אין ווארשע יעס פאריס ליווערפול און לאנדאן סוף כל סוף איז ער"
	yor    = " abinibi han ikawe alantakun le ni opolopo ede abinibi ti a to lesese bi  eniyan to fe lo se fe lati se atunse jowo mo pe awon oju iwe itakunagbaye  miran ti ako ni oniruru ede abinibi le faragba nipa atunse ninu se iwadi  blogs ni ori itakun agbaye ti e ba"
	zho    = "产品的简报和公告 提交该申请后无法进行更改 请确认您的选择是正确的  对于要提交的图书 我确认 我是版权所有者或已得到版权所有者的授权  要更改您的国家 地区 请在此表的最上端更改您的"
	zhoLat = "chanpin de jianbao he gonggao tijiao gai shenqing hou wufa jinxing genggai qing queren nin de xuanze shi zhengque de duiyu yao tijiao de tushu wo queren woshi banquansuoyou zhe huo yi dedao banquansuoyou zhe de shouquan yao genggai nin de guojia diqu qing zaici biao de zuishang duan genggai nin de"
	zul    = " ana engu uma inkinga iqhubeka siza ubike kwi isexwayiso ngenxa yephutha  lomlekeleli sikwazi ukubuyisela emuva kuphela imiphumela engaqediwe  ukuthola imiphumela eqediwe zama ukulayisha kabusha leli khasi emizuzwini  engu uma inkinga iqhubeka siza uthumele"
)

func TestAllLanguages(t *testing.T) {
	det, err := New(0, 1024)
	if err != nil {
		fmt.Println("whoops, couldn't create a new LanguageIdentifier:", err)
	}
	defer Free(det)
	cases := []struct {
		lang string
		text string
	}{
		{"afr", afr},
		{"amh", amh},
		{"ara", ara},
		{"aze", aze},
		{"bel", bel},
		{"bul", bul},
		{"ben", ben},
		// {"bos", bos},
		{"cat", cat},
		{"ceb", ceb},
		{"cos", cos},
		{"ces", ces},
		{"cym", cym},
		{"dan", dan},
		{"deu", deu},
		{"ell", ell},
		{"eng", eng},
		{"epo", epo},
		{"spa", spa},
		{"est", est},
		{"eus", eus},
		{"fas", fas},
		{"fin", fin},
		{"fil", fil},
		{"fra", fra},
		{"fry", fry},
		{"gle", gle},
		{"gla", gla},
		{"glg", glg},
		{"guj", guj},
		{"hau", hau},
		{"haw", haw},
		{"hin", hin},
		{"hmn", hmn},
		{"hrv", hrv},
		{"hat", hat},
		{"hun", hun},
		{"hye", hye},
		// {"ind", ind},
		{"ibo", ibo},
		{"isl", isl},
		{"ita", ita},
		{"heb", heb},
		{"jpn", jpn},
		{"jav", jav},
		{"kat", kat},
		{"kaz", kaz},
		{"khm", khm},
		{"kan", kan},
		{"kor", kor},
		{"kur", kur},
		{"kir", kir},
		{"lat", lat},
		{"ltz", ltz},
		{"lao", lao},
		{"lit", lit},
		{"lav", lav},
		{"mlg", mlg},
		{"mri", mri},
		{"mkd", mkd},
		{"mal", mal},
		{"mon", mon},
		{"mar", mar},
		{"msa", msa},
		{"mlt", mlt},
		{"mya", mya},
		{"nep", nep},
		{"nld", nld},
		{"nor", nor},
		{"nya", nya},
		{"pan", pan},
		{"pol", pol},
		{"pus", pus},
		{"por", por},
		{"ron", ron},
		{"rus", rus},
		{"snd", snd},
		{"sin", sin},
		{"slk", slk},
		{"slv", slv},
		{"smo", smo},
		{"sna", sna},
		{"som", som},
		{"sqi", sqi},
		{"srp", srp},
		{"sot", sot},
		{"sun", sun},
		{"swe", swe},
		{"swa", swa},
		{"tam", tam},
		{"tel", tel},
		{"tgk", tgk},
		{"tha", tha},
		{"tur", tur},
		{"ukr", ukr},
		{"urd", urd},
		{"uzb", uzb},
		{"vie", vie},
		{"xho", xho},
		{"yid", yid},
		{"yor", yor},
		{"zho", zho},
		{"zul", zul},
	}
	for _, c := range cases {
		res := det.FindLanguage(c.text)
		if res.Language != c.lang {
			t.Errorf("Unexpected language: wanted \"%s\", got %#v", c.lang, res.Language)
		}
		if !res.IsReliable {
			t.Errorf("Result for language \"%s\" is not reliable", c.lang)
		}
		if res.Latin {
			t.Errorf("Result for language \"%s\" should not have been detected in latin script", c.lang)
		}
	}
}

func TestAllLanguagesLatin(t *testing.T) {
	det, err := New(0, 1024)
	if err != nil {
		fmt.Println("whoops, couldn't create a new LanguageIdentifier:", err)
	}
	defer Free(det)
	cases := []struct {
		lang string
		text string
	}{
		{"bul", bulLat},
		// {"ell", ellLat},
		// {"hin", hinLat},
		{"jpn", jpnLat},
		{"rus", rusLat},
		{"zho", zhoLat},
	}
	for _, c := range cases {
		res := det.FindLanguage(c.text)
		if res.Language != c.lang {
			t.Errorf("Unexpected language: wanted \"%s\", got %#v", c.lang, res.Language)
		}
		if !res.IsReliable {
			t.Errorf("Result for language \"%s\" is not reliable", c.lang)
		}
		if !res.Latin {
			t.Errorf("Result for language \"%s\" should have been detected in latin script", c.lang)
		}
	}
}
