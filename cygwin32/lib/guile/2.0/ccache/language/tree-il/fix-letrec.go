GOOF----LE-4-2.0;T      ] ° 4      hß      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  language¤	g  tree-il¤	g  
fix-letrec¤		 ¤	
g  filenameS¤	f  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm¤	g  importsS¤	g  system¤	g  base¤	g  syntax¤	 ¤	 ¤	g  srfi¤	g  srfi-1¤	 ¤	 ¤	g  srfi-11¤	 ¤	 ¤	 ¤	 ¤	g  effects¤	 ¤	 ¤	 ¤	g  exportsS¤	 g  fix-letrec!¤	!  ¤	"g  set-current-module¤	#" ¤	$" ¤	%g  <lexical-set>¤	&% ¤	'% ¤	(g  lexical-set-exp¤	)g  <module-set>¤	*) ¤	+) ¤	,g  module-set-exp¤	-g  <toplevel-set>¤	.- ¤	/- ¤	0g  toplevel-set-exp¤	1g  <toplevel-define>¤	21 ¤	31 ¤	4g  toplevel-define-exp¤	5g  <conditional>¤	65 ¤	75 ¤	8g  conditional-test¤	9g  conditional-consequent¤	:g  conditional-alternate¤	;g  <application>¤	<; ¤	=; ¤	>g  application-proc¤	?g  application-args¤	@g  
<sequence>¤	A@ ¤	B@ ¤	Cg  sequence-exps¤	Dg  <lambda>¤	ED ¤	FD ¤	Gg  lambda-body¤	Hg  <lambda-case>¤	IH ¤	JH ¤	Kg  lambda-case-inits¤	Lg  lambda-case-body¤	Mg  lambda-case-alternate¤	Ng  <let>¤	ON ¤	PN ¤	Qg  let-vals¤	Rg  let-body¤	Sg  <letrec>¤	TS ¤	US ¤	Vg  letrec-vals¤	Wg  letrec-body¤	Xg  <fix>¤	YX ¤	ZX ¤	[g  fix-vals¤	\g  fix-body¤	]g  <let-values>¤	^] ¤	_] ¤	`g  let-values-exp¤	ag  let-values-body¤	bg  	<dynwind>¤	cb ¤	db ¤	eg  dynwind-body¤	fg  dynwind-winder¤	gg  dynwind-unwinder¤	hg  <dynlet>¤	ih ¤	jh ¤	kg  dynlet-fluids¤	lg  dynlet-vals¤	mg  dynlet-body¤	ng  <dynref>¤	on ¤	pn ¤	qg  dynref-fluid¤	rg  <dynset>¤	sr ¤	tr ¤	ug  dynset-fluid¤	vg  
dynset-exp¤	wg  <prompt>¤	xw ¤	yw ¤	zg  
prompt-tag¤	{g  prompt-body¤	|g  prompt-handler¤	}g  <abort>¤	~} ¤	} ¤ g  	abort-tag¤ g  
abort-args¤ g  
abort-tail¤ g  fix-fold¤ g  <void>¤ g  <const>¤ g  <lexical-ref>¤ g  lexical-ref-gensym¤ g  memq¤ g  simple-expression?¤ g  and-map¤ g  primitive-ref?¤ g  delq¤ g  lset-adjoin¤ g  eq?¤ g  lexical-set-gensym¤ g  letrec-gensyms¤ g  append¤ g  let-gensyms¤ g  letrec-in-order?¤ g  effect-free?¤ g  exclude-effects¤ g  make-effects-analyzer¤ g  lset-difference¤ g  lambda?¤ g  partition-vars¤ g  make-sequence¤ g  reverse¤ g  make-sequence*¤ g  post-order!¤ g  	make-void¤ g  
letrec-src¤  g  letrec-names¤ ¡g  map¤ ¢g  list¤ £g  make-let¤ ¤g  cadr¤ ¥g  car¤ ¦g  caddr¤ §g  make-fix¤ ¨g  make-lexical-set¤ ©g  module-gensym¤ ªf  fixlr¤ «g  make-lexical-ref¤ ¬g  let-src¤ ­g  	let-names¤ ®g  assq¤ ¯g  lset-intersection¤C 5    h°F  Ñ   ]4	
!5 4$ >  "  G    hh   ·  , 3 (  D4 >  G 
	 	
 "ÿÿ   ¯      g  proc
		f g  exps		f g  unref			f g  ref			f g  set			f g  simple			f g  lambda			f g  complex			f g  unref		2	f g  ref			2	f g  set	
	2	f g  simple		2	f g  lambda		2	f g  complex		2	f  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
		 		f	  g  nameg  fold-values C'(+,/034789:=>?BCFGJKLMPQRUVWZ[\_`adefgjklmpqtuvyz{| 9    h0    ]·4L >  G 
	 §&  #4 54L 	
> ¹" ´&  #4 54L 	
> " &  #4 54L 	
> c" ^&  #4 54L 	
> 8" 3	&  4
 54 54 54L 	
>  G 4L >  G 4L > ¯" ª&  T4 54 54L 	
>  G 4LL > S" N&  %4 54LL 	
> &" !&  ?4 5$  4L 	
> õ"  	
	" ß" Ú&  £4 54 54 54LL 	
>  G $  @4L >  G 4L > N"  4L > 4" /&  T4 54 54LL 	
>  G 4L > Ø" Ó&  T4 54 54LL 	
>  G 4L > |" w&  T4 54  54LL 	
>  G 4L >  " !&  R4" 54# 54L 	
>  G 4L > Æ" Á$&  4% 54& 54' 54L 	
>  G 4L >  G 4L > =" 8(&  4) 54* 54+ 54LL 	
>  G 4LL >  G 4L > °" «,&  #4- 54L 	
> " .&  R4/ 540 54L 	
>  G 4L > +" &1&  42 543 544 54L 	
>  G 4L >  G 4L >  ¢"  5&  46 547 548 54L 	
>  G 4LL >  G 4L >  "  	
	"  G L 6        g  tree
	. g  unref	. g  ref		. g  set		. g  simple		. g  lambda		. g  complex		. g  unref		. g  ref		. g  set			. g  simple	
	. g  lambda		. g  complex		. g  rtd		+ g  exp		<	T g  exp		g	 g  exp	  ª g  exp	 ½ Õ g  test	 ö^ g  
consequent	 ö^ g  	alternate	 ö^ g  unref	^ g  ref	^ g  set	^ g  simple	^ g  lambda	^ g  complex	^ g  unref	9^ g  ref	9^ g  set	9^ g  simple	9^ g  lambda	9^ g  complex	9^ g  proc	xº g  args	xº g  unref	º g  ref	º g  set	º g  simple	º g  lambda	º g  complex	º g  exps	Íç g  body	ú. g  inits	OÙ g  body	OÙ g  	alternate	OÙ g  unref	nÙ g  ref	nÙ g  set	nÙ g  simple	nÙ g  lambda	nÙ g  complex	nÙ g  unref	¿ g  ref	¿ g  set	¿ g  simple	¿ g  lambda	¿ g  complex	¿ g  vals	ó5 g  body	ó5 g  unref	5 g  ref	5 g  set	5 g  simple	5 g  lambda	5 g  complex	5 g  vals	O g  body	O g  unref	l g  ref	l g  set	l g  simple	l g  lambda	l g  complex	l g  vals	«í g  body	«í g  unref	Èí g  ref	Èí g  set	Èí g  simple	Èí g  lambda	Èí g  complex	Èí g  exp	G g  body	G g  unref	"G g  ref	"G g  set	"G g  simple	"G g  lambda	"G g  complex	"G g  body	hÐ g  winder	hÐ g  unwinder	hÐ g  unref	Ð g  ref	Ð g  set	Ð g  simple	Ð g  lambda	Ð g  complex	Ð g  unref	«Ð g  ref	«Ð g  set	«Ð g  simple	«Ð g  lambda	«Ð g  complex	«Ð g  fluids	ñ] g  vals	ñ] g  body	ñ] g  unref	] g  ref	] g  set	] g  simple	] g  lambda	] g  complex	] g  unref	8] g  ref	8] g  set	8] g  simple	8] g  lambda	8] g  complex	8] g  fluid	p g  fluid	¢â g  exp	¢â g  unref	½â g  ref	½â g  set	½â g  simple	½â g  lambda	½â g  complex	½â g  tag	k g  body	k g  handler	k g  unref	 k g  ref	 k g  set	 k g  simple	 k g  lambda	 k g  complex	 k g  unref	Fk g  ref	Fk g  set	Fk g  simple	Fk g  lambda	Fk g  complex	Fk g  tag	ö g  args	ö g  tail	ö g  unref	©ö g  ref	©ö g  set	©ö g  simple	©ö g  lambda	©ö g  complex	©ö g  unref	Ñö g  ref	Ñö g  set	Ñö g  simple	Ñö g  lambda	Ñö g  complex	Ñö g  unref	. g  ref	. g  set	. g  simple	. g  lambda	. g  complex	.  ¥g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
		 	.	  g  nameg  foldts C     h@     , 	3 O 		Q 	O 

	Q 

 6     w      g  tree
		; g  down		; g  up			; g  unref			; g  ref			; g  set			; g  simple			; g  lambda			; g  complex			; g  fold-values				; g  foldts	
		;  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
		 		;		  g  nameg  fix-fold CR589:@C     h   ®   ] LL 6    ¦       g  x
		  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
	,			,	 		   C;>?    h   ®   ] LL 6    ¦       g  x
		  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
	1			2	 		   C       hð   5  ]# §&  C&  C&  4 545C&  G4 54 54	 54
5$  4
5$  

6CC&  4 5O 6&  ;4 54 545$  4 5$  O 6CCC     -      g  x
	 ë g  
bound-vars	 ë g  simple-primcall?		 ë g  rtd		 ë g  gensym		*	7 g  test		T  g  
consequent		T  g  	alternate		T  g  exps	  ¦ g  proc	 ¼ é g  args	 ¼ é  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
	!
		"		-	&	
	6	&		?	"		[	(	
	i	(		j	)	
	x	(	 	*	
 	"	 ¦	,	 ®	"	 Á	/	
 Ë	/	 Ì	0	
 Ö	/	 å	1	
 	 ë	  g  nameg  simple-expression? CR%SN h¸   ý  ] §&  '4 54545D&   4 545D&  4	 54
5D&  4 54
5DDõ      g  x
	 ¸ g  unref	 ¸ g  ref		 ¸ g  set		 ¸ g  simple		 ¸ g  lambda*		 ¸ g  complex		 ¸ g  rtd		 ¸ g  gensym			6 g  gensym		E	^ g  gensyms		m  g  gensyms	  ª  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
	:			;			=		"	>		6	=		>	;		L	F		^	D		f	;		p	K	 	K	 	;	 	R	 ª	R	 ¸	Y	 	 ¸	   CSV  h      ]	4M  54 56       þ       g  x
		 g  effects			  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
	d			e	(			e			g			f	 		  g  nameg  effect+exception-free-primcall? C  h   ®   ] L 6      ¦       g  x
		
  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
	^	0	
	^	< 		
   C      h      ]	4M  54      56       g  x
		 g  effects			  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
	_			`	(			`			b			b	5		b			a	 		  g  nameg  effect-free-primcall? CNQ   hð    ]_ §& ½4 54 54 5
	HO Q 4O 5K" k(  8$  4	5"  4
54
54
5D45$  l"  "ÿÿ$  D45$  "  4	5$  "ÿÿP"ÿÿ"ÿÿ45$  "ÿÿ45$  "ÿþò4	$  "  O 5$  "ÿþ³"ÿþ	
"ÿþ& 
4 54 5	"  ß
(  #4
54
54
5D4
5$  

"ÿÿ¯4
5$  


"ÿÿ"  


"ÿÿa45$  04
5$  "ÿÿÄ


"ÿÿ%"ÿÿ¢	
"ÿÿD          g  x
	ì g  unref	ì g  ref		ì g  set		ì g  simple		ì g  lambda*		ì g  complex		ì g  rtd		ì g  	in-order?		$Ì g  orig-gensyms			$Ì g  vals	
	$Ì g  compute-effects		+Ì g  effect+exception-free-primcall?		5Ì g  gensyms		L· g  vals		L· g  s		L· g  l		L· g  c		L· g  orig-gensyms	âÞ g  vals		âÞ g  gensyms	
êÉ g  vals	êÉ g  s	êÉ g  l	êÉ g  c	êÉ  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
	Z			[		;	^		L	h		R	j		X	p	#	Y	q	'	n	u	#	w	v	# 	w	# 	p	 	x	 	x	! 	x	 	j	  	# ¢ 	1 ¶ 	 ¶	z	 ½	{	) Â	{	2 Ä	{	) È	z	 Î	|	) Ó	}	* Ù	|	) Ú	|	$ Þ	z	 á		# ä		1 ë 	- î 	' ü		 	
 	! 		j	 	 	- 	)" 	#0 	1 	6 	$8 	<	j	? 	B 	-G 	'J 	!Z 	[ 	` 	h 	w 	{	j	~ 	 	- 	% 	 	 	 	-¦ 	)© 	#· 	·	h	¼	i	#½	i	+¾	i	3Ì	h	Ô	[	ê 	ð 	÷ ¤	#  ¥	#	 ¦	# ¡	 §	 §	! §	! 	$ ¨	' ¨	-; ¨	< ª	A ª	!E ª	I 	L «	O «	-V ¬	)Y ¬	#g «	n µ	q µ	-x ¶	){ ¶	# µ	 	 ­	  ­	) ­	  	 ®	% ®	+ ®	%£ ­	ª ¯	­ ¯	-² °	'µ °	!Å ¯	É 	Î 	#Ï 	+Ð 	3Þ 	ì ¸	 x	ì	   C h0     ]14 >	  G D      g  x
		0 g  unref		0 g  ref			0 g  set			0 g  simple			0 g  lambda*			0 g  complex			0  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
	6
		9		 ¹		 º		 »		 ¼		 ½		 ¾			9			7		0 ¿	 		0  g  nameg  partition-vars CRD    h¨   5  ]""  (  (  C 456§&  "ÿÿ¿&  "ÿÿª&  "ÿÿ&  "ÿÿ"ÿÿp"ÿÿe       -      g  src
	 ¡ g  exps	 ¡ g  in		  g  out		  g  head		+  g  rtd		0   g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
 Á
	 Â			 Ã		 Ã		 Ä		 Å		 Æ		! Æ	.	$ Æ	(	& Æ		( Æ		+ Ç		+ Ç		0 È	
	= É		G É		O È	
	R Ê		\ Ê		d È	
	g Ë	!	q Ë		y È	
	| Ì	  Ì	  Í	  Í	  Í	  Â	  Â	 ¡ Â	  	 ¡	  g  nameg  make-sequence* CR%(S VW¡¢     h0   /  ] (  C4 L 5$   4L 5C  "ÿÿÓ '      g  binds
		/  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
 á		 â			 ã	 	 ä		 ä		 ä		 â		 å		 å	$	" å	(	$ å	$	% å		) æ		/ æ	 		/  g  nameg  lp C      h   Ô   ]	O  Q L 6Ì       g  set
		 g  lp		  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
 à		 á	 		  g  nameg  lookup C£¤¥¦    h   °   ]6 ¨       g  x
		  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
 ñ	*	 ñ	6 		   C§¨        h   Ë   ]   6     Ã       g  c
		  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
			1		:		.		 		   C©ª     h   ¹   ]6±       g  x
		  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
	&		7		( 		   C¨«     h    î   ]  4 56      æ       g  x
		 g  tmp		  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
			$		-		!		6		!		  			   CN¬­QR®        h   °   ] L 6      ¨       g  v
		
  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
		
	 		
   C¯®      h   °   ] L 6      ¨       g  v
		
  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
		
	 		
   C®  h   °   ] L 6      ¨       g  v
		
  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
		
	 		
   C '    h   ¤  ]q §&  24 54 54L5$  45 6 C& b4 54	 54
 54 54 54 545O 	Q 	4	L54	L54	L54	L 5
4454554454554454554454545444
5(  	 "  [$  445 5"  >454454544555 5556&  ä4 54 54 54  54! 5454"O 4#$L554%O 4#$L554&O 4#$L 55
	44544	54	54	544
54
54
555 56 C            g  x
	 g  rtd	 g  gensym			A g  exp			A g  src		s« g  	in-order?		s« g  names		s« g  gensyms		s« g  vals		s« g  body		s« g  binds	 « g  lookup		 « g  u	
 ·« g  s	 ·« g  l	 ·« g  c	 ·« g  tmps	n g  src	Ö g  names	Ö g  gensyms	Ö g  vals	Ö g  body	Ö g  binds	í g  u	4 g  l		4 g  c	
4  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
 Ò		 Ó		" Ø		. Ø	
	4 Ù	+	< Ù	!	> Ù		I Ó	  Ý	  Ý	
  ç	 £ è	 ª é	 ± ê	 · ç	 Ä ï	 Ç ï	 Ð ï	$ Ú ï	 Û ð	 Þ ð	 ç ð	# ñ ð	 ò ñ	 õ ñ	 þ ñ	% ñ		 ó	 ô	 ô	!  ô	-) õ	. ÷	1 ú	? û	D þ	N û	O	R	_	a	f	!n	q	u		-				£	¥ ÷	§ õ	© ó	« í	³ Ó	á	í	
ð	ú				(	4	4 	?%	B'	K*	P*	Y*	,b*	8k,	p,	)y,	6,	B,	*	(	%	#	 R	   C h0   )  ]!4 >  G O  6     !      g  x
		+ g  unref		+ g  simple			+ g  lambda*			+ g  complex			+  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm
 Ï
	 Ð	/	 Ð		+ Ñ	 		+  g  nameg  fix-letrec! C RC  É       g  m
		,  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/language/tree-il/fix-letrec.scm		
ê	
è	!
1	6
4 Á
F¬ Ï
 	F®
   C6 