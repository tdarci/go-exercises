package data

var Booklist = `
author, filename, title
"William Shakespeare", "as-you-like-it_TXT_FolgerShakespeare.txt", "As You Like It"
"William Shakespeare", "hamlet_TXT_FolgerShakespeare.txt", "Hamlet"
"William Shakespeare", "loves-labors-lost_TXT_FolgerShakespeare.txt", "Love's Labours Lost"
"William Shakespeare", "richard-iii_TXT_FolgerShakespeare.txt", "Richard III"
"William Shakespeare", "the-tempest_TXT_FolgerShakespeare.txt", "The Tempest"
"Zora Neale Hurston", "hurston-forty-yards.txt", "Forty Yards"
"Zora Neale Hurston", "hurston-lawing-and-jawing.txt", "Lawing and Jawing"
"Zora Neale Hurston", "hurston-woofing.txt", "Woofing"
"Zora Neale Hurston", "hurston-de-turkey-and-de-law.txt", "De Turkey and De Law"
"Zora Neale Hurston", "hurston-poker.txt", "Poker!"
`

var Books = initBookMap()

func initBookMap() map[string]string {
	b := make(map[string]string)

	b["as-you-like-it_TXT_FolgerShakespeare.txt"] = `
As You Like It
by William Shakespeare
Edited by Barbara A. Mowat and Paul Werstine
  with Michael Poston and Rebecca Niles
Folger Shakespeare Library
https://shakespeare.folger.edu/shakespeares-works/as-you-like-it/
Created on Jul 31, 2015, from FDT version 0.9.2

Characters in the Play
======================
ORLANDO, youngest son of Sir Rowland de Boys
OLIVER, his elder brother
SECOND BROTHER, brother to Orlando and Oliver, named Jaques
ADAM, servant to Oliver and friend to Orlando
DENNIS, servant to Oliver
ROSALIND, daughter to Duke Senior
CELIA, Rosalind's cousin, daughter to Duke Frederick
TOUCHSTONE, a court Fool
DUKE FREDERICK, the usurping duke
CHARLES, wrestler at Duke Frederick's court
LE BEAU, a courtier at Duke Frederick's court
Attending Duke Frederick:
  FIRST LORD
  SECOND LORD
DUKE SENIOR, the exiled duke, brother to Duke Frederick
Lords attending Duke Senior in exile:
  JAQUES
  AMIENS
  FIRST LORD
  SECOND LORD
Attending Duke Senior in exile:
  FIRST PAGE
  SECOND PAGE
CORIN, a shepherd
SILVIUS, a young shepherd in love
PHOEBE, a disdainful shepherdess
AUDREY, a goat-keeper
WILLIAM, a country youth in love with Audrey
SIR OLIVER MARTEXT, a parish priest
HYMEN, god of marriage
Lords, Attendants, Musicians


ACT 1
=====

Scene 1
=======
[Enter Orlando and Adam.]


ORLANDO  As I remember, Adam, it was upon this
fashion bequeathed me by will but poor a thousand
crowns, and, as thou sayst, charged my brother on
his blessing to breed me well. And there begins my
sadness. My brother Jaques he keeps at school, and
report speaks goldenly of his profit. For my part, he
keeps me rustically at home, or, to speak more
properly, stays me here at home unkept; for call you
that "keeping," for a gentleman of my birth, that
differs not from the stalling of an ox? His horses are
bred better, for, besides that they are fair with their
feeding, they are taught their manage and, to that
end, riders dearly hired. But I, his brother, gain
nothing under him but growth, for the which his
animals on his dunghills are as much bound to him
as I. Besides this nothing that he so plentifully gives
me, the something that nature gave me his countenance
seems to take from me. He lets me feed with
his hinds, bars me the place of a brother, and, as
much as in him lies, mines my gentility with my
education. This is it, Adam, that grieves me, and the
spirit of my father, which I think is within me,
begins to mutiny against this servitude. I will no
longer endure it, though yet I know no wise remedy
how to avoid it.

[Enter Oliver.]


ADAM  Yonder comes my master, your brother.

ORLANDO  Go apart, Adam, and thou shalt hear how he
will shake me up.	[Adam steps aside.]

OLIVER  Now, sir, what make you here?

ORLANDO  Nothing. I am not taught to make anything.

OLIVER  What mar you then, sir?

ORLANDO  Marry, sir, I am helping you to mar that
which God made, a poor unworthy brother of
yours, with idleness.

OLIVER  Marry, sir, be better employed, and be naught
awhile.

ORLANDO  Shall I keep your hogs and eat husks with
them? What prodigal portion have I spent that I
should come to such penury?

OLIVER  Know you where you are, sir?

ORLANDO  O, sir, very well: here in your orchard.

OLIVER  Know you before whom, sir?

ORLANDO  Ay, better than him I am before knows me. I
know you are my eldest brother, and in the gentle
condition of blood you should so know me. The
courtesy of nations allows you my better in that you
are the first-born, but the same tradition takes not
away my blood, were there twenty brothers betwixt
us. I have as much of my father in me as you, albeit I
confess your coming before me is nearer to his
reverence.

OLIVER, [threatening Orlando]  What, boy!

ORLANDO, [holding off Oliver by the throat]  Come,
come, elder brother, you are too young in this.

OLIVER  Wilt thou lay hands on me, villain?

ORLANDO  I am no villain. I am the youngest son of Sir
Rowland de Boys. He was my father, and he is
thrice a villain that says such a father begot villains.
Wert thou not my brother, I would not take this
hand from thy throat till this other had pulled out
thy tongue for saying so. Thou hast railed on thyself.

ADAM, [coming forward]  Sweet masters, be patient. For
your father's remembrance, be at accord.

OLIVER, [to Orlando]  Let me go, I say.

ORLANDO  I will not till I please. You shall hear me. My
father charged you in his will to give me good
education. You have trained me like a peasant,
obscuring and hiding from me all gentlemanlike
qualities. The spirit of my father grows strong in
me, and I will no longer endure it. Therefore allow
me such exercises as may become a gentleman, or
give me the poor allottery my father left me by
testament. With that I will go buy my fortunes.
[Orlando releases Oliver.]

OLIVER  And what wilt thou do--beg when that is
spent? Well, sir, get you in. I will not long be
troubled with you. You shall have some part of your
will. I pray you leave me.

ORLANDO  I will no further offend you than becomes
me for my good.

OLIVER, [to Adam]  Get you with him, you old dog.

ADAM  Is "old dog" my reward? Most true, I have lost
my teeth in your service. God be with my old
master. He would not have spoke such a word.
[Orlando and Adam exit.]

OLIVER  Is it even so? Begin you to grow upon me? I
will physic your rankness, and yet give no thousand
crowns neither.--Holla, Dennis!

[Enter Dennis.]


DENNIS  Calls your Worship?

OLIVER  Was not Charles, the Duke's wrestler, here to
speak with me?

DENNIS  So please you, he is here at the door and
importunes access to you.

OLIVER  Call him in. [Dennis exits.] 'Twill be a good
way, and tomorrow the wrestling is.

[Enter Charles.]


CHARLES  Good morrow to your Worship.

OLIVER  Good Monsieur Charles, what's the new news
at the new court?

CHARLES  There's no news at the court, sir, but the old
news. That is, the old duke is banished by his
younger brother the new duke, and three or four
loving lords have put themselves into voluntary
exile with him, whose lands and revenues enrich
the new duke. Therefore he gives them good leave
to wander.

OLIVER  Can you tell if Rosalind, the Duke's daughter,
be banished with her father?

CHARLES  O, no, for the Duke's daughter her cousin so
loves her, being ever from their cradles bred together,
that she would have followed her exile or have
died to stay behind her. She is at the court and no
less beloved of her uncle than his own daughter,
and never two ladies loved as they do.

OLIVER  Where will the old duke live?

CHARLES  They say he is already in the Forest of Arden,
and a many merry men with him; and there they
live like the old Robin Hood of England. They say
many young gentlemen flock to him every day and
fleet the time carelessly, as they did in the golden
world.

OLIVER  What, you wrestle tomorrow before the new
duke?

CHARLES  Marry, do I, sir, and I came to acquaint you
with a matter. I am given, sir, secretly to understand
that your younger brother Orlando hath a
disposition to come in disguised against me to try a
fall. Tomorrow, sir, I wrestle for my credit, and he
that escapes me without some broken limb shall
acquit him well. Your brother is but young and
tender, and for your love I would be loath to foil
him, as I must for my own honor if he come in.
Therefore, out of my love to you, I came hither to
acquaint you withal, that either you might stay him
from his intendment, or brook such disgrace well
as he shall run into, in that it is a thing of his own
search and altogether against my will.

OLIVER  Charles, I thank thee for thy love to me, which
thou shalt find I will most kindly requite. I had
myself notice of my brother's purpose herein, and
have by underhand means labored to dissuade him
from it; but he is resolute. I'll tell thee, Charles, it is
the stubbornest young fellow of France, full of
ambition, an envious emulator of every man's good
parts, a secret and villainous contriver against me
his natural brother. Therefore use thy discretion. I
had as lief thou didst break his neck as his finger.
And thou wert best look to 't, for if thou dost him
any slight disgrace, or if he do not mightily grace
himself on thee, he will practice against thee by
poison, entrap thee by some treacherous device,
and never leave thee till he hath ta'en thy life by
some indirect means or other. For I assure thee--
and almost with tears I speak it--there is not one so
young and so villainous this day living. I speak but
brotherly of him, but should I anatomize him to
thee as he is, I must blush and weep, and thou must
look pale and wonder.

CHARLES  I am heartily glad I came hither to you. If he
come tomorrow, I'll give him his payment. If ever
he go alone again, I'll never wrestle for prize more.
And so God keep your Worship.

OLIVER  Farewell, good Charles.	[Charles exits.]
Now will I stir this gamester. I hope I shall see an
end of him, for my soul--yet I know not why--
hates nothing more than he. Yet he's gentle, never
schooled and yet learned, full of noble device, of all
sorts enchantingly beloved, and indeed so much in
the heart of the world, and especially of my own
people, who best know him, that I am altogether
misprized. But it shall not be so long; this wrestler
shall clear all. Nothing remains but that I kindle the
boy thither, which now I'll go about.
[He exits.]

Scene 2
=======
[Enter Rosalind and Celia.]


CELIA  I pray thee, Rosalind, sweet my coz, be merry.

ROSALIND  Dear Celia, I show more mirth than I am
mistress of, and would you yet I were merrier?
Unless you could teach me to forget a banished
father, you must not learn me how to remember
any extraordinary pleasure.

CELIA  Herein I see thou lov'st me not with the full
weight that I love thee. If my uncle, thy banished
father, had banished thy uncle, the Duke my father,
so thou hadst been still with me, I could have taught
my love to take thy father for mine. So wouldst thou,
if the truth of thy love to me were so righteously
tempered as mine is to thee.

ROSALIND  Well, I will forget the condition of my estate
to rejoice in yours.

CELIA  You know my father hath no child but I, nor
none is like to have; and truly, when he dies, thou
shalt be his heir, for what he hath taken away from
thy father perforce, I will render thee again in
affection. By mine honor I will, and when I break
that oath, let me turn monster. Therefore, my sweet
Rose, my dear Rose, be merry.

ROSALIND  From henceforth I will, coz, and devise
sports. Let me see--what think you of falling in
love?

CELIA  Marry, I prithee do, to make sport withal; but
love no man in good earnest, nor no further in
sport neither than with safety of a pure blush thou
mayst in honor come off again.

ROSALIND  What shall be our sport, then?

CELIA  Let us sit and mock the good housewife Fortune
from her wheel, that her gifts may henceforth be
bestowed equally.

ROSALIND  I would we could do so, for her benefits are
mightily misplaced, and the bountiful blind woman
doth most mistake in her gifts to women.

CELIA  'Tis true, for those that she makes fair she scarce
makes honest, and those that she makes honest she
makes very ill-favoredly.

ROSALIND  Nay, now thou goest from Fortune's office to
Nature's. Fortune reigns in gifts of the world, not in
the lineaments of nature.

CELIA  No? When Nature hath made a fair creature,
may she not by fortune fall into the fire?

[Enter Touchstone.]

Though Nature hath given us wit to flout at Fortune,
hath not Fortune sent in this fool to cut off the
argument?

ROSALIND  Indeed, there is Fortune too hard for Nature,
when Fortune makes Nature's natural the
cutter-off of Nature's wit.

CELIA  Peradventure this is not Fortune's work neither,
but Nature's, who perceiveth our natural wits too
dull to reason of such goddesses, and hath sent
this natural for our whetstone, for always the dullness
of the fool is the whetstone of the wits. [To
Touchstone.] How now, wit, whither wander you?

TOUCHSTONE  Mistress, you must come away to your
father.

CELIA  Were you made the messenger?

TOUCHSTONE  No, by mine honor, but I was bid to come
for you.

ROSALIND  Where learned you that oath, fool?

TOUCHSTONE  Of a certain knight that swore by his
honor they were good pancakes, and swore by his
honor the mustard was naught. Now, I'll stand to it,
the pancakes were naught and the mustard was
good, and yet was not the knight forsworn.

CELIA  How prove you that in the great heap of your
knowledge?

ROSALIND  Ay, marry, now unmuzzle your wisdom.

TOUCHSTONE  Stand you both forth now: stroke your
chins, and swear by your beards that I am a knave.

CELIA  By our beards (if we had them), thou art.

TOUCHSTONE  By my knavery (if I had it), then I were.
But if you swear by that that is not, you are not
forsworn. No more was this knight swearing by his
honor, for he never had any, or if he had, he had
sworn it away before ever he saw those pancakes or
that mustard.

CELIA  Prithee, who is 't that thou mean'st?

TOUCHSTONE  One that old Frederick, your father, loves.

CELIA  My father's love is enough to honor him.
Enough. Speak no more of him; you'll be whipped
for taxation one of these days.

TOUCHSTONE  The more pity that fools may not speak
wisely what wise men do foolishly.

CELIA  By my troth, thou sayest true. For, since the little
wit that fools have was silenced, the little foolery
that wise men have makes a great show. Here
comes Monsieur Le Beau.

[Enter Le Beau.]


ROSALIND  With his mouth full of news.

CELIA  Which he will put on us as pigeons feed their
young.

ROSALIND  Then shall we be news-crammed.

CELIA  All the better. We shall be the more
marketable.--Bonjour, Monsieur Le Beau. What's
the news?

LE BEAU  Fair princess, you have lost much good sport.

CELIA  Sport? Of what color?

LE BEAU  What color, madam? How shall I answer you?

ROSALIND  As wit and fortune will.

TOUCHSTONE  Or as the destinies decrees.

CELIA  Well said. That was laid on with a trowel.

TOUCHSTONE  Nay, if I keep not my rank--

ROSALIND  Thou losest thy old smell.

LE BEAU  You amaze me, ladies. I would have told you of
good wrestling, which you have lost the sight of.

ROSALIND  Yet tell us the manner of the wrestling.

LE BEAU  I will tell you the beginning, and if it please
your Ladyships, you may see the end, for the best is
yet to do, and here, where you are, they are coming
to perform it.

CELIA  Well, the beginning that is dead and buried.

LE BEAU  There comes an old man and his three sons--

CELIA  I could match this beginning with an old tale.

LE BEAU  Three proper young men of excellent growth
and presence.

ROSALIND  With bills on their necks: "Be it known unto
all men by these presents."

LE BEAU  The eldest of the three wrestled with Charles,
the Duke's wrestler, which Charles in a moment
threw him and broke three of his ribs, that there is
little hope of life in him. So he served the second,
and so the third. Yonder they lie, the poor old man
their father making such pitiful dole over them that
all the beholders take his part with weeping.

ROSALIND  Alas!

TOUCHSTONE  But what is the sport, monsieur, that the
ladies have lost?

LE BEAU  Why, this that I speak of.

TOUCHSTONE  Thus men may grow wiser every day. It is
the first time that ever I heard breaking of ribs was
sport for ladies.

CELIA  Or I, I promise thee.

ROSALIND  But is there any else longs to see this broken
music in his sides? Is there yet another dotes upon
rib-breaking? Shall we see this wrestling, cousin?

LE BEAU  You must if you stay here, for here is the place
appointed for the wrestling, and they are ready to
perform it.

CELIA  Yonder sure they are coming. Let us now stay
and see it.

[Flourish. Enter Duke Frederick, Lords, Orlando,
Charles, and Attendants.]


DUKE FREDERICK  Come on. Since the youth will not be
entreated, his own peril on his forwardness.

ROSALIND, [to Le Beau]  Is yonder the man?

LE BEAU  Even he, madam.

CELIA  Alas, he is too young. Yet he looks successfully.

DUKE FREDERICK  How now, daughter and cousin? Are
you crept hither to see the wrestling?

ROSALIND  Ay, my liege, so please you give us leave.

DUKE FREDERICK  You will take little delight in it, I can
tell you, there is such odds in the man. In pity of the
challenger's youth, I would fain dissuade him, but
he will not be entreated. Speak to him, ladies; see if
you can move him.

CELIA  Call him hither, good Monsieur Le Beau.

DUKE FREDERICK  Do so. I'll not be by.
[He steps aside.]

LE BEAU, [to Orlando]  Monsieur the challenger, the
Princess calls for you.

ORLANDO  I attend them with all respect and duty.

ROSALIND  Young man, have you challenged Charles the
wrestler?

ORLANDO  No, fair princess. He is the general challenger.
I come but in as others do, to try with him the
strength of my youth.

CELIA  Young gentleman, your spirits are too bold for
your years. You have seen cruel proof of this man's
strength. If you saw yourself with your eyes or knew
yourself with your judgment, the fear of your adventure
would counsel you to a more equal enterprise.
We pray you for your own sake to embrace your
own safety and give over this attempt.

ROSALIND  Do, young sir. Your reputation shall not
therefore be misprized. We will make it our suit to
the Duke that the wrestling might not go forward.

ORLANDO  I beseech you, punish me not with your hard
thoughts, wherein I confess me much guilty to deny
so fair and excellent ladies anything. But let your
fair eyes and gentle wishes go with me to my trial,
wherein, if I be foiled, there is but one shamed that
was never gracious; if killed, but one dead that is
willing to be so. I shall do my friends no wrong, for
I have none to lament me; the world no injury, for
in it I have nothing. Only in the world I fill up a
place which may be better supplied when I have
made it empty.

ROSALIND  The little strength that I have, I would it
were with you.

CELIA  And mine, to eke out hers.

ROSALIND  Fare you well. Pray heaven I be deceived in
you.

CELIA  Your heart's desires be with you.

CHARLES  Come, where is this young gallant that is so
desirous to lie with his mother Earth?

ORLANDO  Ready, sir; but his will hath in it a more
modest working.

DUKE FREDERICK, [coming forward]  You shall try but
one fall.

CHARLES  No, I warrant your Grace you shall not entreat
him to a second, that have so mightily persuaded
him from a first.

ORLANDO  You mean to mock me after, you should not
have mocked me before. But come your ways.

ROSALIND  Now Hercules be thy speed, young man!

CELIA  I would I were invisible, to catch the strong
fellow by the leg.
[Orlando and Charles wrestle.]

ROSALIND  O excellent young man!

CELIA  If I had a thunderbolt in mine eye, I can tell who
should down.
[Orlando throws Charles. Shout.]

DUKE FREDERICK  No more, no more.

ORLANDO  Yes, I beseech your Grace. I am not yet well
breathed.

DUKE FREDERICK  How dost thou, Charles?

LE BEAU  He cannot speak, my lord.

DUKE FREDERICK  Bear him away.
[Charles is carried off by Attendants.]
What is thy name, young man?

ORLANDO  Orlando, my liege, the youngest son of Sir
Rowland de Boys.

DUKE FREDERICK
I would thou hadst been son to some man else.
The world esteemed thy father honorable,
But I did find him still mine enemy.
Thou shouldst have better pleased me with this
deed
Hadst thou descended from another house.
But fare thee well. Thou art a gallant youth.
I would thou hadst told me of another father.
[Duke exits with Touchstone, Le Beau,
Lords, and Attendants.]

CELIA, [to Rosalind]
Were I my father, coz, would I do this?

ORLANDO
I am more proud to be Sir Rowland's son,
His youngest son, and would not change that calling
To be adopted heir to Frederick.

ROSALIND, [to Celia]
My father loved Sir Rowland as his soul,
And all the world was of my father's mind.
Had I before known this young man his son,
I should have given him tears unto entreaties
Ere he should thus have ventured.

CELIA  Gentle cousin,
Let us go thank him and encourage him.
My father's rough and envious disposition
Sticks me at heart.--Sir, you have well deserved.
If you do keep your promises in love
But justly, as you have exceeded all promise,
Your mistress shall be happy.

ROSALIND, [giving Orlando a chain from her neck]
Gentleman,
Wear this for me--one out of suits with Fortune,
That could give more but that her hand lacks
means.--
Shall we go, coz?

CELIA  Ay.--Fare you well, fair gentleman.

ORLANDO, [aside]
Can I not say "I thank you"? My better parts
Are all thrown down, and that which here stands up
Is but a quintain, a mere lifeless block.

ROSALIND, [to Celia]
He calls us back. My pride fell with my fortunes.
I'll ask him what he would.--Did you call, sir?
Sir, you have wrestled well and overthrown
More than your enemies.

CELIA  Will you go, coz?

ROSALIND  Have with you. [To Orlando.] Fare you well.
[Rosalind and Celia exit.]

ORLANDO
What passion hangs these weights upon my tongue?
I cannot speak to her, yet she urged conference.
O poor Orlando! Thou art overthrown.
Or Charles or something weaker masters thee.

[Enter Le Beau.]


LE BEAU
Good sir, I do in friendship counsel you
To leave this place. Albeit you have deserved
High commendation, true applause, and love,
Yet such is now the Duke's condition
That he misconsters all that you have done.
The Duke is humorous. What he is indeed
More suits you to conceive than I to speak of.

ORLANDO
I thank you, sir, and pray you tell me this:
Which of the two was daughter of the duke
That here was at the wrestling?

LE BEAU
Neither his daughter, if we judge by manners,
But yet indeed the smaller is his daughter.
The other is daughter to the banished duke,
And here detained by her usurping uncle
To keep his daughter company, whose loves
Are dearer than the natural bond of sisters.
But I can tell you that of late this duke
Hath ta'en displeasure 'gainst his gentle niece,
Grounded upon no other argument
But that the people praise her for her virtues
And pity her for her good father's sake;
And, on my life, his malice 'gainst the lady
Will suddenly break forth. Sir, fare you well.
Hereafter, in a better world than this,
I shall desire more love and knowledge of you.

ORLANDO
I rest much bounden to you. Fare you well.
[Le Beau exits.]
Thus must I from the smoke into the smother,
From tyrant duke unto a tyrant brother.
But heavenly Rosalind!
[He exits.]

Scene 3
=======
[Enter Celia and Rosalind.]


CELIA  Why, cousin! Why, Rosalind! Cupid have mercy,
not a word?

ROSALIND  Not one to throw at a dog.

CELIA  No, thy words are too precious to be cast away
upon curs. Throw some of them at me. Come, lame
me with reasons.

ROSALIND  Then there were two cousins laid up, when
the one should be lamed with reasons, and the
other mad without any.

CELIA  But is all this for your father?

ROSALIND  No, some of it is for my child's father. O,
how full of briers is this working-day world!

CELIA  They are but burs, cousin, thrown upon thee in
holiday foolery. If we walk not in the trodden paths,
our very petticoats will catch them.

ROSALIND  I could shake them off my coat. These burs
are in my heart.

CELIA  Hem them away.

ROSALIND  I would try, if I could cry "hem" and have
him.

CELIA  Come, come, wrestle with thy affections.

ROSALIND  O, they take the part of a better wrestler
than myself.

CELIA  O, a good wish upon you. You will try in time, in
despite of a fall. But turning these jests out of
service, let us talk in good earnest. Is it possible on
such a sudden you should fall into so strong a liking
with old Sir Rowland's youngest son?

ROSALIND  The Duke my father loved his father dearly.

CELIA  Doth it therefore ensue that you should love his
son dearly? By this kind of chase I should hate him,
for my father hated his father dearly. Yet I hate not
Orlando.

ROSALIND  No, faith, hate him not, for my sake.

CELIA  Why should I not? Doth he not deserve well?

ROSALIND  Let me love him for that, and do you love
him because I do.

[Enter Duke Frederick with Lords.]

Look, here comes the Duke.

CELIA  With his eyes full of anger.

DUKE FREDERICK, [to Rosalind]
Mistress, dispatch you with your safest haste,
And get you from our court.

ROSALIND  Me, uncle?

DUKE FREDERICK  You, cousin.
Within these ten days if that thou beest found
So near our public court as twenty miles,
Thou diest for it.

ROSALIND  I do beseech your Grace,
Let me the knowledge of my fault bear with me.
If with myself I hold intelligence
Or have acquaintance with mine own desires,
If that I do not dream or be not frantic--
As I do trust I am not--then, dear uncle,
Never so much as in a thought unborn
Did I offend your Highness.

DUKE FREDERICK  Thus do all traitors.
If their purgation did consist in words,
They are as innocent as grace itself.
Let it suffice thee that I trust thee not.

ROSALIND
Yet your mistrust cannot make me a traitor.
Tell me whereon the likelihood depends.

DUKE FREDERICK
Thou art thy father's daughter. There's enough.

ROSALIND
So was I when your Highness took his dukedom.
So was I when your Highness banished him.
Treason is not inherited, my lord,
Or if we did derive it from our friends,
What's that to me? My father was no traitor.
Then, good my liege, mistake me not so much
To think my poverty is treacherous.

CELIA  Dear sovereign, hear me speak.

DUKE FREDERICK
Ay, Celia, we stayed her for your sake;
Else had she with her father ranged along.

CELIA
I did not then entreat to have her stay.
It was your pleasure and your own remorse.
I was too young that time to value her,
But now I know her. If she be a traitor,
Why, so am I. We still have slept together,
Rose at an instant, learned, played, eat together,
And, wheresoe'er we went, like Juno's swans
Still we went coupled and inseparable.

DUKE FREDERICK
She is too subtle for thee, and her smoothness,
Her very silence, and her patience
Speak to the people, and they pity her.
Thou art a fool. She robs thee of thy name,
And thou wilt show more bright and seem more
virtuous
When she is gone. Then open not thy lips.
Firm and irrevocable is my doom
Which I have passed upon her. She is banished.

CELIA
Pronounce that sentence then on me, my liege.
I cannot live out of her company.

DUKE FREDERICK
You are a fool.--You, niece, provide yourself.
If you outstay the time, upon mine honor
And in the greatness of my word, you die.
[Duke and Lords exit.]

CELIA
O my poor Rosalind, whither wilt thou go?
Wilt thou change fathers? I will give thee mine.
I charge thee, be not thou more grieved than I am.

ROSALIND  I have more cause.

CELIA  Thou hast not, cousin.
Prithee, be cheerful. Know'st thou not the Duke
Hath banished me, his daughter?

ROSALIND  That he hath not.

CELIA
No, hath not? Rosalind lacks then the love
Which teacheth thee that thou and I am one.
Shall we be sundered? Shall we part, sweet girl?
No, let my father seek another heir.
Therefore devise with me how we may fly,
Whither to go, and what to bear with us,
And do not seek to take your change upon you,
To bear your griefs yourself and leave me out.
For, by this heaven, now at our sorrows pale,
Say what thou canst, I'll go along with thee.

ROSALIND  Why, whither shall we go?

CELIA
To seek my uncle in the Forest of Arden.

ROSALIND
Alas, what danger will it be to us,
Maids as we are, to travel forth so far?
Beauty provoketh thieves sooner than gold.

CELIA
I'll put myself in poor and mean attire,
And with a kind of umber smirch my face.
The like do you. So shall we pass along
And never stir assailants.

ROSALIND  Were it not better,
Because that I am more than common tall,
That I did suit me all points like a man?
A gallant curtal-ax upon my thigh,
A boar-spear in my hand, and in my heart
Lie there what hidden woman's fear there will,
We'll have a swashing and a martial outside--
As many other mannish cowards have
That do outface it with their semblances.

CELIA
What shall I call thee when thou art a man?

ROSALIND
I'll have no worse a name than Jove's own page,
And therefore look you call me Ganymede.
But what will you be called?

CELIA
Something that hath a reference to my state:
No longer Celia, but Aliena.

ROSALIND
But, cousin, what if we assayed to steal
The clownish fool out of your father's court?
Would he not be a comfort to our travel?

CELIA
He'll go along o'er the wide world with me.
Leave me alone to woo him. Let's away
And get our jewels and our wealth together,
Devise the fittest time and safest way
To hide us from pursuit that will be made
After my flight. Now go we in content
To liberty, and not to banishment.
[They exit.]


ACT 2
=====

Scene 1
=======
[Enter Duke Senior, Amiens, and two or three Lords, like
foresters.]


DUKE SENIOR
Now, my co-mates and brothers in exile,
Hath not old custom made this life more sweet
Than that of painted pomp? Are not these woods
More free from peril than the envious court?
Here feel we not the penalty of Adam,
The seasons' difference, as the icy fang
And churlish chiding of the winter's wind,
Which when it bites and blows upon my body
Even till I shrink with cold, I smile and say
"This is no flattery. These are counselors
That feelingly persuade me what I am."
Sweet are the uses of adversity,
Which, like the toad, ugly and venomous,
Wears yet a precious jewel in his head.
And this our life, exempt from public haunt,
Finds tongues in trees, books in the running brooks,
Sermons in stones, and good in everything.

AMIENS
I would not change it. Happy is your Grace,
That can translate the stubbornness of fortune
Into so quiet and so sweet a style.

DUKE SENIOR
Come, shall we go and kill us venison?
And yet it irks me the poor dappled fools,
Being native burghers of this desert city,
Should in their own confines with forked heads
Have their round haunches gored.

FIRST LORD  Indeed, my lord,
The melancholy Jaques grieves at that,
And in that kind swears you do more usurp
Than doth your brother that hath banished you.
Today my Lord of Amiens and myself
Did steal behind him as he lay along
Under an oak, whose antique root peeps out
Upon the brook that brawls along this wood;
To the which place a poor sequestered stag
That from the hunter's aim had ta'en a hurt
Did come to languish. And indeed, my lord,
The wretched animal heaved forth such groans
That their discharge did stretch his leathern coat
Almost to bursting, and the big round tears
Coursed one another down his innocent nose
In piteous chase. And thus the hairy fool,
Much marked of the melancholy Jaques,
Stood on th' extremest verge of the swift brook,
Augmenting it with tears.

DUKE SENIOR  But what said Jaques?
Did he not moralize this spectacle?

FIRST LORD
O yes, into a thousand similes.
First, for his weeping into the needless stream:
"Poor deer," quoth he, "thou mak'st a testament
As worldlings do, giving thy sum of more
To that which had too much." Then, being there
alone,
Left and abandoned of his velvet friends:
"'Tis right," quoth he. "Thus misery doth part
The flux of company." Anon a careless herd,
Full of the pasture, jumps along by him
And never stays to greet him. "Ay," quoth Jaques,
"Sweep on, you fat and greasy citizens.
'Tis just the fashion. Wherefore do you look
Upon that poor and broken bankrupt there?"
Thus most invectively he pierceth through
The body of country, city, court,
Yea, and of this our life, swearing that we
Are mere usurpers, tyrants, and what's worse,
To fright the animals and to kill them up
In their assigned and native dwelling place.

DUKE SENIOR
And did you leave him in this contemplation?

SECOND LORD
We did, my lord, weeping and commenting
Upon the sobbing deer.

DUKE SENIOR  Show me the place.
I love to cope him in these sullen fits,
For then he's full of matter.

FIRST LORD  I'll bring you to him straight.
[They exit.]

Scene 2
=======
[Enter Duke Frederick with Lords.]


DUKE FREDERICK
Can it be possible that no man saw them?
It cannot be. Some villains of my court
Are of consent and sufferance in this.

FIRST LORD
I cannot hear of any that did see her.
The ladies her attendants of her chamber
Saw her abed, and in the morning early
They found the bed untreasured of their mistress.

SECOND LORD
My lord, the roinish clown at whom so oft
Your Grace was wont to laugh is also missing.
Hisperia, the Princess' gentlewoman,
Confesses that she secretly o'erheard
Your daughter and her cousin much commend
The parts and graces of the wrestler
That did but lately foil the sinewy Charles,
And she believes wherever they are gone
That youth is surely in their company.

DUKE FREDERICK
Send to his brother. Fetch that gallant hither.
If he be absent, bring his brother to me.
I'll make him find him. Do this suddenly,
And let not search and inquisition quail
To bring again these foolish runaways.
[They exit.]

Scene 3
=======
[Enter Orlando and Adam, meeting.]


ORLANDO  Who's there?

ADAM
What, my young master, O my gentle master,
O my sweet master, O you memory
Of old Sir Rowland! Why, what make you here?
Why are you virtuous? Why do people love you?
And wherefore are you gentle, strong, and valiant?
Why would you be so fond to overcome
The bonny prizer of the humorous duke?
Your praise is come too swiftly home before you.
Know you not, master, to some kind of men
Their graces serve them but as enemies?
No more do yours. Your virtues, gentle master,
Are sanctified and holy traitors to you.
O, what a world is this when what is comely
Envenoms him that bears it!

ORLANDO  Why, what's the matter?

ADAM  O unhappy youth,
Come not within these doors. Within this roof
The enemy of all your graces lives.
Your brother--no, no brother--yet the son--
Yet not the son, I will not call him son--
Of him I was about to call his father,
Hath heard your praises, and this night he means
To burn the lodging where you use to lie,
And you within it. If he fail of that,
He will have other means to cut you off.
I overheard him and his practices.
This is no place, this house is but a butchery.
Abhor it, fear it, do not enter it.

ORLANDO
Why, whither, Adam, wouldst thou have me go?

ADAM
No matter whither, so you come not here.

ORLANDO
What, wouldst thou have me go and beg my food,
Or with a base and boist'rous sword enforce
A thievish living on the common road?
This I must do, or know not what to do;
Yet this I will not do, do how I can.
I rather will subject me to the malice
Of a diverted blood and bloody brother.

ADAM
But do not so. I have five hundred crowns,
The thrifty hire I saved under your father,
Which I did store to be my foster nurse
When service should in my old limbs lie lame,
And unregarded age in corners thrown.
Take that, and He that doth the ravens feed,
Yea, providently caters for the sparrow,
Be comfort to my age. Here is the gold.
All this I give you. Let me be your servant.
Though I look old, yet I am strong and lusty,
For in my youth I never did apply
Hot and rebellious liquors in my blood,
Nor did not with unbashful forehead woo
The means of weakness and debility.
Therefore my age is as a lusty winter,
Frosty but kindly. Let me go with you.
I'll do the service of a younger man
In all your business and necessities.

ORLANDO
O good old man, how well in thee appears
The constant service of the antique world,
When service sweat for duty, not for meed.
Thou art not for the fashion of these times,
Where none will sweat but for promotion,
And having that do choke their service up
Even with the having. It is not so with thee.
But, poor old man, thou prun'st a rotten tree
That cannot so much as a blossom yield
In lieu of all thy pains and husbandry.
But come thy ways. We'll go along together,
And ere we have thy youthful wages spent,
We'll light upon some settled low content.

ADAM
Master, go on, and I will follow thee
To the last gasp with truth and loyalty.
From seventeen years till now almost fourscore
Here lived I, but now live here no more.
At seventeen years, many their fortunes seek,
But at fourscore, it is too late a week.
Yet fortune cannot recompense me better
Than to die well, and not my master's debtor.
[They exit.]

Scene 4
=======
[Enter Rosalind for Ganymede, Celia for Aliena, and
Clown, alias Touchstone.]


ROSALIND
O Jupiter, how weary are my spirits!

TOUCHSTONE  I care not for my spirits, if my legs were
not weary.

ROSALIND  I could find in my heart to disgrace my
man's apparel and to cry like a woman, but I must
comfort the weaker vessel, as doublet and hose
ought to show itself courageous to petticoat. Therefore
courage, good Aliena.

CELIA  I pray you bear with me. I cannot go no further.

TOUCHSTONE  For my part, I had rather bear with you
than bear you. Yet I should bear no cross if I did
bear you, for I think you have no money in your
purse.

ROSALIND  Well, this is the Forest of Arden.

TOUCHSTONE  Ay, now am I in Arden, the more fool I.
When I was at home I was in a better place, but
travelers must be content.

ROSALIND  Ay, be so, good Touchstone.

[Enter Corin and Silvius.]

Look you who comes here, a young man and an old
in solemn talk.

[Rosalind, Celia, and Touchstone step aside and
eavesdrop.]



CORIN, [to Silvius]
That is the way to make her scorn you still.

SILVIUS
O Corin, that thou knew'st how I do love her!

CORIN
I partly guess, for I have loved ere now.

SILVIUS
No, Corin, being old, thou canst not guess,
Though in thy youth thou wast as true a lover
As ever sighed upon a midnight pillow.
But if thy love were ever like to mine--
As sure I think did never man love so--
How many actions most ridiculous
Hast thou been drawn to by thy fantasy?

CORIN
Into a thousand that I have forgotten.

SILVIUS
O, thou didst then never love so heartily.
If thou rememb'rest not the slightest folly
That ever love did make thee run into,
Thou hast not loved.
Or if thou hast not sat as I do now,
Wearing thy hearer in thy mistress' praise,
Thou hast not loved.
Or if thou hast not broke from company
Abruptly, as my passion now makes me,
Thou hast not loved.
O Phoebe, Phoebe, Phoebe!	[He exits.]

ROSALIND
Alas, poor shepherd, searching of thy wound,
I have by hard adventure found mine own.

TOUCHSTONE  And I mine. I remember when I was in
love I broke my sword upon a stone and bid him
take that for coming a-night to Jane Smile; and I
remember the kissing of her batler, and the cow's
dugs that her pretty chopped hands had milked;
and I remember the wooing of a peascod instead of
her, from whom I took two cods and, giving her
them again, said with weeping tears "Wear these for
my sake." We that are true lovers run into strange
capers. But as all is mortal in nature, so is all nature
in love mortal in folly.

ROSALIND  Thou speak'st wiser than thou art ware of.

TOUCHSTONE  Nay, I shall ne'er be ware of mine own
wit till I break my shins against it.

ROSALIND
Jove, Jove, this shepherd's passion
Is much upon my fashion.

TOUCHSTONE  And mine, but it grows something stale
with me.

CELIA  I pray you, one of you question yond man, if he
for gold will give us any food. I faint almost to death.

TOUCHSTONE, [to Corin]  Holla, you clown!

ROSALIND  Peace, fool. He's not thy kinsman.

CORIN  Who calls?

TOUCHSTONE  Your betters, sir.

CORIN  Else are they very wretched.

ROSALIND, [to Touchstone]
Peace, I say. [As Ganymede, to Corin.]
 Good even toyou, friend.

CORIN
And to you, gentle sir, and to you all.

ROSALIND, [as Ganymede]
I prithee, shepherd, if that love or gold
Can in this desert place buy entertainment,
Bring us where we may rest ourselves and feed.
Here's a young maid with travel much oppressed,
And faints for succor.

CORIN  Fair sir, I pity her
And wish for her sake more than for mine own
My fortunes were more able to relieve her.
But I am shepherd to another man
And do not shear the fleeces that I graze.
My master is of churlish disposition
And little recks to find the way to heaven
By doing deeds of hospitality.
Besides, his cote, his flocks, and bounds of feed
Are now on sale, and at our sheepcote now,
By reason of his absence, there is nothing
That you will feed on. But what is, come see,
And in my voice most welcome shall you be.

ROSALIND, [as Ganymede]
What is he that shall buy his flock and pasture?

CORIN
That young swain that you saw here but erewhile,
That little cares for buying anything.

ROSALIND, [as Ganymede]
I pray thee, if it stand with honesty,
Buy thou the cottage, pasture, and the flock,
And thou shalt have to pay for it of us.

CELIA, [as Aliena]
And we will mend thy wages. I like this place,
And willingly could waste my time in it.

CORIN
Assuredly the thing is to be sold.
Go with me. If you like upon report
The soil, the profit, and this kind of life,
I will your very faithful feeder be
And buy it with your gold right suddenly.
[They exit.]

Scene 5
=======
[Enter Amiens, Jaques, and others.]

Song.

AMIENS [sings]
	   Under the greenwood tree
	   Who loves to lie with me
	   And turn his merry note
	   Unto the sweet bird's throat,
	Come hither, come hither, come hither.
	      Here shall he see
	      No enemy
	But winter and rough weather.

JAQUES  More, more, I prithee, more.

AMIENS  It will make you melancholy, Monsieur
Jaques.

JAQUES  I thank it. More, I prithee, more. I can suck
melancholy out of a song as a weasel sucks eggs.
More, I prithee, more.

AMIENS  My voice is ragged. I know I cannot please you.

JAQUES  I do not desire you to please me. I do desire
you to sing. Come, more, another stanzo. Call you
'em "stanzos"?

AMIENS  What you will, Monsieur Jaques.

JAQUES  Nay, I care not for their names. They owe me
nothing. Will you sing?

AMIENS  More at your request than to please myself.

JAQUES  Well then, if ever I thank any man, I'll thank
you. But that they call "compliment" is like th'
encounter of two dog-apes. And when a man thanks
me heartily, methinks I have given him a penny and
he renders me the beggarly thanks. Come, sing. And
you that will not, hold your tongues.

AMIENS  Well, I'll end the song.--Sirs, cover the while;
the Duke will drink under this tree.--He hath been
all this day to look you.

JAQUES  And I have been all this day to avoid him. He is
too disputable for my company. I think of as many
matters as he, but I give heaven thanks and make no
boast of them. Come, warble, come.

Song.


ALL [together here.]
	   Who doth ambition shun
	   And loves to live i' th' sun,
	   Seeking the food he eats
	   And pleased with what he gets,
	Come hither, come hither, come hither.
	      Here shall he see
	      No enemy
	But winter and rough weather.

JAQUES  I'll give you a verse to this note that I made
yesterday in despite of my invention.

AMIENS  And I'll sing it.

JAQUES  Thus it goes:
	   If it do come to pass
	   That any man turn ass,
	   Leaving his wealth and ease
	   A stubborn will to please,
	Ducdame, ducdame, ducdame.
	      Here shall he see
	      Gross fools as he,
	An if he will come to me.

AMIENS  What's that "ducdame"?

JAQUES  'Tis a Greek invocation to call fools into a
circle. I'll go sleep if I can. If I cannot, I'll rail
against all the first-born of Egypt.

AMIENS  And I'll go seek the Duke. His banquet is
prepared.
[They exit.]

Scene 6
=======
[Enter Orlando and Adam.]


ADAM  Dear master, I can go no further. O, I die for
food. Here lie I down and measure out my grave.
Farewell, kind master.	[He lies down.]

ORLANDO  Why, how now, Adam? No greater heart in
thee? Live a little, comfort a little, cheer thyself a
little. If this uncouth forest yield anything savage, I
will either be food for it or bring it for food to thee.
Thy conceit is nearer death than thy powers. For my
sake, be comfortable. Hold death awhile at the
arm's end. I will here be with thee presently, and if
I bring thee not something to eat, I will give thee
leave to die. But if thou diest before I come, thou art
a mocker of my labor. Well said. Thou look'st
cheerly, and I'll be with thee quickly. Yet thou liest
in the bleak air. Come, I will bear thee to some
shelter, and thou shalt not die for lack of a dinner if
there live anything in this desert. Cheerly, good
Adam.
[They exit.]

Scene 7
=======
[Enter Duke Senior and Lords, like outlaws.]


DUKE SENIOR
I think he be transformed into a beast,
For I can nowhere find him like a man.

FIRST LORD
My lord, he is but even now gone hence.
Here was he merry, hearing of a song.

DUKE SENIOR
If he, compact of jars, grow musical,
We shall have shortly discord in the spheres.
Go seek him. Tell him I would speak with him.

[Enter Jaques.]


FIRST LORD
He saves my labor by his own approach.

DUKE SENIOR, [to Jaques]
Why, how now, monsieur? What a life is this
That your poor friends must woo your company?
What, you look merrily.

JAQUES
A fool, a fool, I met a fool i' th' forest,
A motley fool. A miserable world!
As I do live by food, I met a fool,
Who laid him down and basked him in the sun
And railed on Lady Fortune in good terms,
In good set terms, and yet a motley fool.
"Good morrow, fool," quoth I. "No, sir," quoth he,
"Call me not 'fool' till heaven hath sent me
fortune."
And then he drew a dial from his poke
And, looking on it with lack-luster eye,
Says very wisely "It is ten o'clock.
Thus we may see," quoth he, "how the world wags.
'Tis but an hour ago since it was nine,
And after one hour more 'twill be eleven.
And so from hour to hour we ripe and ripe,
And then from hour to hour we rot and rot,
And thereby hangs a tale." When I did hear
The motley fool thus moral on the time,
My lungs began to crow like chanticleer
That fools should be so deep-contemplative,
And I did laugh sans intermission
An hour by his dial. O noble fool!
A worthy fool! Motley's the only wear.

DUKE SENIOR  What fool is this?

JAQUES
O worthy fool!--One that hath been a courtier,
And says "If ladies be but young and fair,
They have the gift to know it." And in his brain,
Which is as dry as the remainder biscuit
After a voyage, he hath strange places crammed
With observation, the which he vents
In mangled forms. O, that I were a fool!
I am ambitious for a motley coat.

DUKE SENIOR
Thou shalt have one.

JAQUES  It is my only suit,
Provided that you weed your better judgments
Of all opinion that grows rank in them
That I am wise. I must have liberty
Withal, as large a charter as the wind,
To blow on whom I please, for so fools have.
And they that are most galled with my folly,
They most must laugh. And why, sir, must they so?
The "why" is plain as way to parish church:
He that a fool doth very wisely hit
Doth very foolishly, although he smart,
Not to seem senseless of the bob. If not,
The wise man's folly is anatomized
Even by the squand'ring glances of the fool.
Invest me in my motley. Give me leave
To speak my mind, and I will through and through
Cleanse the foul body of th' infected world,
If they will patiently receive my medicine.

DUKE SENIOR
Fie on thee! I can tell what thou wouldst do.

JAQUES
What, for a counter, would I do but good?

DUKE SENIOR
Most mischievous foul sin in chiding sin;
For thou thyself hast been a libertine,
As sensual as the brutish sting itself,
And all th' embossed sores and headed evils
That thou with license of free foot hast caught
Wouldst thou disgorge into the general world.

JAQUES  Why, who cries out on pride
That can therein tax any private party?
Doth it not flow as hugely as the sea
Till that the weary very means do ebb?
What woman in the city do I name
When that I say the city-woman bears
The cost of princes on unworthy shoulders?
Who can come in and say that I mean her,
When such a one as she such is her neighbor?
Or what is he of basest function
That says his bravery is not on my cost,
Thinking that I mean him, but therein suits
His folly to the mettle of my speech?
There then. How then, what then? Let me see
wherein
My tongue hath wronged him. If it do him right,
Then he hath wronged himself. If he be free,
Why then my taxing like a wild goose flies
Unclaimed of any man.

[Enter Orlando, brandishing a sword.]

But who comes here?

ORLANDO  Forbear, and eat no more.

JAQUES  Why, I have eat none yet.

ORLANDO
Nor shalt not till necessity be served.

JAQUES  Of what kind should this cock come of?

DUKE SENIOR, [to Orlando]
Art thou thus boldened, man, by thy distress,
Or else a rude despiser of good manners,
That in civility thou seem'st so empty?

ORLANDO
You touched my vein at first. The thorny point
Of bare distress hath ta'en from me the show
Of smooth civility, yet am I inland bred
And know some nurture. But forbear, I say.
He dies that touches any of this fruit
Till I and my affairs are answered.

JAQUES  An you will not be answered with reason, I
must die.

DUKE SENIOR, [to Orlando]
What would you have? Your gentleness shall force
More than your force move us to gentleness.

ORLANDO
I almost die for food, and let me have it.

DUKE SENIOR
Sit down and feed, and welcome to our table.

ORLANDO
Speak you so gently? Pardon me, I pray you.
I thought that all things had been savage here,
And therefore put I on the countenance
Of stern commandment. But whate'er you are
That in this desert inaccessible,
Under the shade of melancholy boughs,
Lose and neglect the creeping hours of time,
If ever you have looked on better days,
If ever been where bells have knolled to church,
If ever sat at any good man's feast,
If ever from your eyelids wiped a tear
And know what 'tis to pity and be pitied,
Let gentleness my strong enforcement be,
In the which hope I blush and hide my sword.
[He sheathes his sword.]

DUKE SENIOR
True is it that we have seen better days,
And have with holy bell been knolled to church,
And sat at good men's feasts and wiped our eyes
Of drops that sacred pity hath engendered.
And therefore sit you down in gentleness,
And take upon command what help we have
That to your wanting may be ministered.

ORLANDO
Then but forbear your food a little while
Whiles, like a doe, I go to find my fawn
And give it food. There is an old poor man
Who after me hath many a weary step
Limped in pure love. Till he be first sufficed,
Oppressed with two weak evils, age and hunger,
I will not touch a bit.

DUKE SENIOR  Go find him out,
And we will nothing waste till you return.

ORLANDO
I thank you; and be blessed for your good comfort.
[He exits.]

DUKE SENIOR
Thou seest we are not all alone unhappy.
This wide and universal theater
Presents more woeful pageants than the scene
Wherein we play in.

JAQUES  All the world's a stage,
And all the men and women merely players.
They have their exits and their entrances,
And one man in his time plays many parts,
His acts being seven ages. At first the infant,
Mewling and puking in the nurse's arms.
Then the whining schoolboy with his satchel
And shining morning face, creeping like snail
Unwillingly to school. And then the lover,
Sighing like furnace, with a woeful ballad
Made to his mistress' eyebrow. Then a soldier,
Full of strange oaths and bearded like the pard,
Jealous in honor, sudden and quick in quarrel,
Seeking the bubble reputation
Even in the cannon's mouth. And then the justice,
In fair round belly with good capon lined,
With eyes severe and beard of formal cut,
Full of wise saws and modern instances;
And so he plays his part. The sixth age shifts
Into the lean and slippered pantaloon
With spectacles on nose and pouch on side,
His youthful hose, well saved, a world too wide
For his shrunk shank, and his big manly voice,
Turning again toward childish treble, pipes
And whistles in his sound. Last scene of all,
That ends this strange eventful history,
Is second childishness and mere oblivion,
Sans teeth, sans eyes, sans taste, sans everything.

[Enter Orlando, carrying Adam.]


DUKE SENIOR
Welcome. Set down your venerable burden,
And let him feed.

ORLANDO  I thank you most for him.

ADAM  So had you need.--
I scarce can speak to thank you for myself.

DUKE SENIOR
Welcome. Fall to. I will not trouble you
As yet to question you about your fortunes.--
Give us some music, and, good cousin, sing.

[The Duke and Orlando continue their conversation,
apart.]


Song.


AMIENS [sings]
	   Blow, blow, thou winter wind.
	   Thou art not so unkind
	      As man's ingratitude.
	   Thy tooth is not so keen,
	   Because thou art not seen,
	      Although thy breath be rude.
	Heigh-ho, sing heigh-ho, unto the green holly.
	Most friendship is feigning, most loving mere folly.
	      Then heigh-ho, the holly.
	      This life is most jolly.

	   Freeze, freeze, thou bitter sky,
	   That dost not bite so nigh
	      As benefits forgot.
	   Though thou the waters warp,
	   Thy sting is not so sharp
	      As friend remembered not.
	Heigh-ho, sing heigh-ho, unto the green holly.
	Most friendship is feigning, most loving mere folly.
	      Then heigh-ho, the holly.
	      This life is most jolly.

DUKE SENIOR, [to Orlando]
If that you were the good Sir Rowland's son,
As you have whispered faithfully you were,
And as mine eye doth his effigies witness
Most truly limned and living in your face,
Be truly welcome hither. I am the duke
That loved your father. The residue of your fortune
Go to my cave and tell me.--Good old man,
Thou art right welcome as thy master is.
[To Lords.] Support him by the arm. [To Orlando.]
Give me your hand,
And let me all your fortunes understand.
[They exit.]


ACT 3
=====

Scene 1
=======
[Enter Duke Frederick, Lords, and Oliver.]


DUKE FREDERICK, [to Oliver]
Not see him since? Sir, sir, that cannot be.
But were I not the better part made mercy,
I should not seek an absent argument
Of my revenge, thou present. But look to it:
Find out thy brother wheresoe'er he is.
Seek him with candle. Bring him, dead or living,
Within this twelvemonth, or turn thou no more
To seek a living in our territory.
Thy lands and all things that thou dost call thine,
Worth seizure, do we seize into our hands
Till thou canst quit thee by thy brother's mouth
Of what we think against thee.

OLIVER
O, that your Highness knew my heart in this:
I never loved my brother in my life.

DUKE FREDERICK
More villain thou.--Well, push him out of doors,
And let my officers of such a nature
Make an extent upon his house and lands.
Do this expediently, and turn him going.
[They exit.]

Scene 2
=======
[Enter Orlando, with a paper.]


ORLANDO
Hang there, my verse, in witness of my love.
   And thou, thrice-crowned queen of night, survey
With thy chaste eye, from thy pale sphere above,
   Thy huntress' name that my full life doth sway.
O Rosalind, these trees shall be my books,
   And in their barks my thoughts I'll character,
That every eye which in this forest looks
   Shall see thy virtue witnessed everywhere.
Run, run, Orlando, carve on every tree
The fair, the chaste, and unexpressive she.
[He exits.]

[Enter Corin and Touchstone.]


CORIN  And how like you this shepherd's life, Master
Touchstone?

TOUCHSTONE  Truly, shepherd, in respect of itself, it is a
good life; but in respect that it is a shepherd's life, it
is naught. In respect that it is solitary, I like it very
well; but in respect that it is private, it is a very vile
life. Now in respect it is in the fields, it pleaseth me
well; but in respect it is not in the court, it is
tedious. As it is a spare life, look you, it fits my
humor well; but as there is no more plenty in it, it
goes much against my stomach. Hast any philosophy
in thee, shepherd?

CORIN  No more but that I know the more one sickens,
the worse at ease he is, and that he that wants
money, means, and content is without three good
friends; that the property of rain is to wet, and fire
to burn; that good pasture makes fat sheep; and that
a great cause of the night is lack of the sun; that he
that hath learned no wit by nature nor art may
complain of good breeding or comes of a very dull
kindred.

TOUCHSTONE  Such a one is a natural philosopher. Wast
ever in court, shepherd?

CORIN  No, truly.

TOUCHSTONE  Then thou art damned.

CORIN  Nay, I hope.

TOUCHSTONE  Truly, thou art damned, like an ill-roasted
egg, all on one side.

CORIN  For not being at court? Your reason.

TOUCHSTONE  Why, if thou never wast at court, thou
never saw'st good manners; if thou never saw'st
good manners, then thy manners must be wicked,
and wickedness is sin, and sin is damnation. Thou
art in a parlous state, shepherd.

CORIN  Not a whit, Touchstone. Those that are good
manners at the court are as ridiculous in the
country as the behavior of the country is most
mockable at the court. You told me you salute not at
the court but you kiss your hands. That courtesy
would be uncleanly if courtiers were shepherds.

TOUCHSTONE  Instance, briefly. Come, instance.

CORIN  Why, we are still handling our ewes, and their
fells, you know, are greasy.

TOUCHSTONE  Why, do not your courtier's hands sweat?
And is not the grease of a mutton as wholesome as
the sweat of a man? Shallow, shallow. A better
instance, I say. Come.

CORIN  Besides, our hands are hard.

TOUCHSTONE  Your lips will feel them the sooner. Shallow
again. A more sounder instance. Come.

CORIN  And they are often tarred over with the surgery
of our sheep; and would you have us kiss tar? The
courtier's hands are perfumed with civet.

TOUCHSTONE  Most shallow man. Thou worms' meat in
respect of a good piece of flesh, indeed. Learn of the
wise and perpend: civet is of a baser birth than tar,
the very uncleanly flux of a cat. Mend the instance,
shepherd.

CORIN  You have too courtly a wit for me. I'll rest.

TOUCHSTONE  Wilt thou rest damned? God help thee,
shallow man. God make incision in thee; thou art
raw.

CORIN  Sir, I am a true laborer. I earn that I eat, get that
I wear, owe no man hate, envy no man's happiness,
glad of other men's good, content with my harm,
and the greatest of my pride is to see my ewes graze
and my lambs suck.

TOUCHSTONE  That is another simple sin in you, to bring
the ewes and the rams together and to offer to get
your living by the copulation of cattle; to be bawd to
a bell-wether and to betray a she-lamb of a twelvemonth
to a crooked-pated old cuckoldly ram, out of
all reasonable match. If thou be'st not damned for
this, the devil himself will have no shepherds. I
cannot see else how thou shouldst 'scape.

[Enter Rosalind, as Ganymede.]


CORIN  Here comes young Master Ganymede, my new
mistress's brother.

ROSALIND, [as Ganymede, reading a paper]
	From the east to western Ind
	No jewel is like Rosalind.
	Her worth being mounted on the wind,
	Through all the world bears Rosalind.
	All the pictures fairest lined
	Are but black to Rosalind.
	Let no face be kept in mind
	But the fair of Rosalind.

TOUCHSTONE  I'll rhyme you so eight years together,
dinners and suppers and sleeping hours excepted.
It is the right butter-women's rank to market.

ROSALIND, [as Ganymede]  Out, fool.

TOUCHSTONE  For a taste:
	If a hart do lack a hind,
	Let him seek out Rosalind.
	If the cat will after kind,
	So be sure will Rosalind.
	Wintered garments must be lined;
	So must slender Rosalind.
	They that reap must sheaf and bind;
	Then to cart with Rosalind.
	Sweetest nut hath sourest rind;
	Such a nut is Rosalind.
	He that sweetest rose will find
	Must find love's prick, and Rosalind.
This is the very false gallop of verses. Why do you
infect yourself with them?

ROSALIND, [as Ganymede]  Peace, you dull fool. I found
them on a tree.

TOUCHSTONE  Truly, the tree yields bad fruit.

ROSALIND, [as Ganymede]  I'll graft it with you, and
then I shall graft it with a medlar. Then it will be
the earliest fruit i' th' country, for you'll be rotten
ere you be half ripe, and that's the right virtue of
the medlar.

TOUCHSTONE  You have said, but whether wisely or no,
let the forest judge.

[Enter Celia, as Aliena, with a writing.]


ROSALIND, [as Ganymede]  Peace. Here comes my sister
reading. Stand aside.

CELIA, [as Aliena, reads]
	Why should this a desert be?
	   For it is unpeopled? No.
	Tongues I'll hang on every tree
	   That shall civil sayings show.
	Some how brief the life of man
	   Runs his erring pilgrimage,
	That the stretching of a span
	   Buckles in his sum of age;
	Some of violated vows
	   'Twixt the souls of friend and friend.
	But upon the fairest boughs,
	   Or at every sentence' end,
	Will I "Rosalinda" write,
	   Teaching all that read to know
	The quintessence of every sprite
	   Heaven would in little show.
	Therefore heaven nature charged
	   That one body should be filled
	With all graces wide-enlarged.
	   Nature presently distilled
	Helen's cheek, but not her heart,
	   Cleopatra's majesty,
	Atalanta's better part,
	   Sad Lucretia's modesty.
	Thus Rosalind of many parts
	   By heavenly synod was devised
	Of many faces, eyes, and hearts
	   To have the touches dearest prized.
	Heaven would that she these gifts should have
	And I to live and die her slave.

ROSALIND, [as Ganymede]  O most gentle Jupiter, what
tedious homily of love have you wearied your parishioners
withal, and never cried "Have patience,
good people!"

CELIA, [as Aliena]  How now?--Back, friends. Shepherd,
go off a little.--Go with him, sirrah.

TOUCHSTONE  Come, shepherd, let us make an honorable
retreat, though not with bag and baggage, yet
with scrip and scrippage.
[Touchstone and Corin exit.]

CELIA  Didst thou hear these verses?

ROSALIND  O yes, I heard them all, and more too, for
some of them had in them more feet than the verses
would bear.

CELIA  That's no matter. The feet might bear the verses.

ROSALIND  Ay, but the feet were lame and could not
bear themselves without the verse, and therefore
stood lamely in the verse.

CELIA  But didst thou hear without wondering how thy
name should be hanged and carved upon these
trees?

ROSALIND  I was seven of the nine days out of the
wonder before you came, for look here what I
found on a palm tree. [She shows the paper she
read.] I was never so berhymed since Pythagoras'
time that I was an Irish rat, which I can hardly
remember.

CELIA  Trow you who hath done this?

ROSALIND  Is it a man?

CELIA  And a chain, that you once wore, about his neck.
Change you color?

ROSALIND  I prithee, who?

CELIA  O Lord, Lord, it is a hard matter for friends to
meet, but mountains may be removed with earthquakes
and so encounter.

ROSALIND  Nay, but who is it?

CELIA  Is it possible?

ROSALIND  Nay, I prithee now, with most petitionary
vehemence, tell me who it is.

CELIA  O wonderful, wonderful, and most wonderful
wonderful, and yet again wonderful, and after that
out of all whooping!

ROSALIND  Good my complexion, dost thou think
though I am caparisoned like a man, I have a
doublet and hose in my disposition? One inch of
delay more is a South Sea of discovery. I prithee,
tell me who is it quickly, and speak apace. I would
thou couldst stammer, that thou might'st pour this
concealed man out of thy mouth as wine comes out
of a narrow-mouthed bottle--either too much at
once, or none at all. I prithee take the cork out of
thy mouth, that I may drink thy tidings.

CELIA  So you may put a man in your belly.

ROSALIND  Is he of God's making? What manner of
man? Is his head worth a hat, or his chin worth a
beard?

CELIA  Nay, he hath but a little beard.

ROSALIND  Why, God will send more, if the man will be
thankful. Let me stay the growth of his beard, if
thou delay me not the knowledge of his chin.

CELIA  It is young Orlando, that tripped up the wrestler's
heels and your heart both in an instant.

ROSALIND  Nay, but the devil take mocking. Speak sad
brow and true maid.

CELIA  I' faith, coz, 'tis he.

ROSALIND  Orlando?

CELIA  Orlando.

ROSALIND  Alas the day, what shall I do with my doublet
and hose? What did he when thou saw'st him? What
said he? How looked he? Wherein went he? What
makes he here? Did he ask for me? Where remains
he? How parted he with thee? And when shalt thou
see him again? Answer me in one word.

CELIA  You must borrow me Gargantua's mouth first.
'Tis a word too great for any mouth of this age's size.
To say ay and no to these particulars is more than to
answer in a catechism.

ROSALIND  But doth he know that I am in this forest and
in man's apparel? Looks he as freshly as he did the
day he wrestled?

CELIA  It is as easy to count atomies as to resolve the
propositions of a lover. But take a taste of my
finding him, and relish it with good observance. I
found him under a tree like a dropped acorn.

ROSALIND  It may well be called Jove's tree when it
drops forth such fruit.

CELIA  Give me audience, good madam.

ROSALIND  Proceed.

CELIA  There lay he, stretched along like a wounded
knight.

ROSALIND  Though it be pity to see such a sight, it well
becomes the ground.

CELIA  Cry "holla" to thy tongue, I prithee. It curvets
unseasonably. He was furnished like a hunter.

ROSALIND  O, ominous! He comes to kill my heart.

CELIA  I would sing my song without a burden. Thou
bring'st me out of tune.

ROSALIND  Do you not know I am a woman? When I
think, I must speak. Sweet, say on.

CELIA  You bring me out.

[Enter Orlando and Jaques.]

Soft, comes he not here?

ROSALIND  'Tis he. Slink by, and note him.
[Rosalind and Celia step aside.]

JAQUES, [to Orlando]  I thank you for your company,
but, good faith, I had as lief have been myself alone.

ORLANDO  And so had I, but yet, for fashion sake, I
thank you too for your society.

JAQUES  God be wi' you. Let's meet as little as we can.

ORLANDO  I do desire we may be better strangers.

JAQUES  I pray you mar no more trees with writing love
songs in their barks.

ORLANDO  I pray you mar no more of my verses with
reading them ill-favoredly.

JAQUES  Rosalind is your love's name?

ORLANDO  Yes, just.

JAQUES  I do not like her name.

ORLANDO  There was no thought of pleasing you when
she was christened.

JAQUES  What stature is she of?

ORLANDO  Just as high as my heart.

JAQUES  You are full of pretty answers. Have you not
been acquainted with goldsmiths' wives and
conned them out of rings?

ORLANDO  Not so. But I answer you right painted cloth,
from whence you have studied your questions.

JAQUES  You have a nimble wit. I think 'twas made of
Atalanta's heels. Will you sit down with me? And we
two will rail against our mistress the world and all
our misery.

ORLANDO  I will chide no breather in the world but
myself, against whom I know most faults.

JAQUES  The worst fault you have is to be in love.

ORLANDO  'Tis a fault I will not change for your best
virtue. I am weary of you.

JAQUES  By my troth, I was seeking for a fool when I
found you.

ORLANDO  He is drowned in the brook. Look but in, and
you shall see him.

JAQUES  There I shall see mine own figure.

ORLANDO  Which I take to be either a fool or a cipher.

JAQUES  I'll tarry no longer with you. Farewell, good
Signior Love.

ORLANDO  I am glad of your departure. Adieu, good
Monsieur Melancholy.	[Jaques exits.]

ROSALIND, [aside to Celia]  I will speak to him like a
saucy lackey, and under that habit play the knave
with him. [As Ganymede.] Do you hear, forester?

ORLANDO  Very well. What would you?

ROSALIND, [as Ganymede]  I pray you, what is 't
o'clock?

ORLANDO  You should ask me what time o' day. There's
no clock in the forest.

ROSALIND, [as Ganymede]  Then there is no true lover
in the forest; else sighing every minute and
groaning every hour would detect the lazy foot of
time as well as a clock.

ORLANDO  And why not the swift foot of time? Had not
that been as proper?

ROSALIND, [as Ganymede]  By no means, sir. Time
travels in divers paces with divers persons. I'll tell
you who time ambles withal, who time trots withal,
who time gallops withal, and who he stands still
withal.

ORLANDO  I prithee, who doth he trot withal?

ROSALIND, [as Ganymede]  Marry, he trots hard with a
young maid between the contract of her marriage
and the day it is solemnized. If the interim be but a
se'nnight, time's pace is so hard that it seems the
length of seven year.

ORLANDO  Who ambles time withal?

ROSALIND, [as Ganymede]  With a priest that lacks Latin
and a rich man that hath not the gout, for the one
sleeps easily because he cannot study, and the other
lives merrily because he feels no pain--the one
lacking the burden of lean and wasteful learning,
the other knowing no burden of heavy tedious
penury. These time ambles withal.

ORLANDO  Who doth he gallop withal?

ROSALIND, [as Ganymede]  With a thief to the gallows,
for though he go as softly as foot can fall, he thinks
himself too soon there.

ORLANDO  Who stays it still withal?

ROSALIND, [as Ganymede]  With lawyers in the vacation,
for they sleep between term and term, and
then they perceive not how time moves.

ORLANDO  Where dwell you, pretty youth?

ROSALIND, [as Ganymede]  With this shepherdess, my
sister, here in the skirts of the forest, like fringe
upon a petticoat.

ORLANDO  Are you native of this place?

ROSALIND, [as Ganymede]  As the cony that you see
dwell where she is kindled.

ORLANDO  Your accent is something finer than you
could purchase in so removed a dwelling.

ROSALIND, [as Ganymede]  I have been told so of many.
But indeed an old religious uncle of mine taught
me to speak, who was in his youth an inland man,
one that knew courtship too well, for there he fell in
love. I have heard him read many lectures against it,
and I thank God I am not a woman, to be touched
with so many giddy offenses as he hath generally
taxed their whole sex withal.

ORLANDO  Can you remember any of the principal evils
that he laid to the charge of women?

ROSALIND, [as Ganymede]  There were none principal.
They were all like one another as halfpence are,
every one fault seeming monstrous till his fellow
fault came to match it.

ORLANDO  I prithee recount some of them.

ROSALIND, [as Ganymede]  No, I will not cast away my
physic but on those that are sick. There is a man
haunts the forest that abuses our young plants with
carving "Rosalind" on their barks, hangs odes upon
hawthorns and elegies on brambles, all, forsooth,
deifying the name of Rosalind. If I could meet
that fancy-monger, I would give him some good
counsel, for he seems to have the quotidian of love
upon him.

ORLANDO  I am he that is so love-shaked. I pray you tell
me your remedy.

ROSALIND, [as Ganymede]  There is none of my uncle's
marks upon you. He taught me how to know a man
in love, in which cage of rushes I am sure you are
not prisoner.

ORLANDO  What were his marks?

ROSALIND, [as Ganymede]  A lean cheek, which you
have not; a blue eye and sunken, which you have
not; an unquestionable spirit, which you have not; a
beard neglected, which you have not--but I pardon
you for that, for simply your having in beard is a
younger brother's revenue. Then your hose should
be ungartered, your bonnet unbanded, your sleeve
unbuttoned, your shoe untied, and everything
about you demonstrating a careless desolation. But
you are no such man. You are rather point-device in
your accouterments, as loving yourself than seeming
the lover of any other.

ORLANDO  Fair youth, I would I could make thee believe
I love.

ROSALIND, [as Ganymede]  Me believe it? You may as
soon make her that you love believe it, which I
warrant she is apter to do than to confess she does.
That is one of the points in the which women still
give the lie to their consciences. But, in good sooth,
are you he that hangs the verses on the trees
wherein Rosalind is so admired?

ORLANDO  I swear to thee, youth, by the white hand of
Rosalind, I am that he, that unfortunate he.

ROSALIND, [as Ganymede]  But are you so much in love
as your rhymes speak?

ORLANDO  Neither rhyme nor reason can express how
much.

ROSALIND, [as Ganymede]  Love is merely a madness,
and, I tell you, deserves as well a dark house and a
whip as madmen do; and the reason why they are
not so punished and cured is that the lunacy is so
ordinary that the whippers are in love too. Yet I
profess curing it by counsel.

ORLANDO  Did you ever cure any so?

ROSALIND, [as Ganymede]  Yes, one, and in this manner.
He was to imagine me his love, his mistress,
and I set him every day to woo me; at which time
would I, being but a moonish youth, grieve, be
effeminate, changeable, longing and liking, proud,
fantastical, apish, shallow, inconstant, full of tears,
full of smiles; for every passion something, and for
no passion truly anything, as boys and women are,
for the most part, cattle of this color; would now
like him, now loathe him; then entertain him, then
forswear him; now weep for him, then spit at him,
that I drave my suitor from his mad humor of love
to a living humor of madness, which was to forswear
the full stream of the world and to live in a
nook merely monastic. And thus I cured him, and
this way will I take upon me to wash your liver as
clean as a sound sheep's heart, that there shall not
be one spot of love in 't.

ORLANDO  I would not be cured, youth.

ROSALIND, [as Ganymede]  I would cure you if you
would but call me Rosalind and come every day to
my cote and woo me.

ORLANDO  Now, by the faith of my love, I will. Tell me
where it is.

ROSALIND, [as Ganymede]  Go with me to it, and I'll
show it you; and by the way you shall tell me where
in the forest you live. Will you go?

ORLANDO  With all my heart, good youth.

ROSALIND, [as Ganymede]  Nay, you must call me
Rosalind.--Come, sister, will you go?
[They exit.]

Scene 3
=======
[Enter Touchstone and Audrey, followed by Jaques.]


TOUCHSTONE  Come apace, good Audrey. I will fetch up
your goats, Audrey. And how, Audrey? Am I the
man yet? Doth my simple feature content you?

AUDREY  Your features, Lord warrant us! What
features?

TOUCHSTONE  I am here with thee and thy goats, as the
most capricious poet, honest Ovid, was among the
Goths.

JAQUES, [aside]  O knowledge ill-inhabited, worse than
Jove in a thatched house.

TOUCHSTONE  When a man's verses cannot be understood,
nor a man's good wit seconded with the
forward child, understanding, it strikes a man more
dead than a great reckoning in a little room. Truly, I
would the gods had made thee poetical.

AUDREY  I do not know what "poetical" is. Is it honest
in deed and word? Is it a true thing?

TOUCHSTONE  No, truly, for the truest poetry is the most
feigning, and lovers are given to poetry, and what
they swear in poetry may be said as lovers they do
feign.

AUDREY  Do you wish, then, that the gods had made me
poetical?

TOUCHSTONE  I do, truly, for thou swear'st to me thou
art honest. Now if thou wert a poet, I might have
some hope thou didst feign.

AUDREY  Would you not have me honest?

TOUCHSTONE  No, truly, unless thou wert hard-favored;
for honesty coupled to beauty is to have honey a
sauce to sugar.

JAQUES, [aside]  A material fool.

AUDREY  Well, I am not fair, and therefore I pray the
gods make me honest.

TOUCHSTONE  Truly, and to cast away honesty upon a
foul slut were to put good meat into an unclean
dish.

AUDREY  I am not a slut, though I thank the gods I am
foul.

TOUCHSTONE  Well, praised be the gods for thy foulness;
sluttishness may come hereafter. But be it as it may
be, I will marry thee; and to that end I have been
with Sir Oliver Martext, the vicar of the next village,
who hath promised to meet me in this place of the
forest and to couple us.

JAQUES, [aside]  I would fain see this meeting.

AUDREY  Well, the gods give us joy.

TOUCHSTONE  Amen. A man may, if he were of a fearful
heart, stagger in this attempt, for here we have no
temple but the wood, no assembly but horn-beasts.
But what though? Courage. As horns are odious,
they are necessary. It is said "Many a man knows no
end of his goods." Right: many a man has good
horns and knows no end of them. Well, that is the
dowry of his wife; 'tis none of his own getting.
Horns? Even so. Poor men alone? No, no. The
noblest deer hath them as huge as the rascal. Is the
single man therefore blessed? No. As a walled town
is more worthier than a village, so is the forehead of
a married man more honorable than the bare brow
of a bachelor. And by how much defense is better
than no skill, by so much is a horn more precious
than to want.

[Enter Sir Oliver Martext.]

Here comes Sir Oliver.--Sir Oliver Martext, you are
well met. Will you dispatch us here under this tree,
or shall we go with you to your chapel?

OLIVER MARTEXT  Is there none here to give the
woman?

TOUCHSTONE  I will not take her on gift of any man.

OLIVER MARTEXT  Truly, she must be given, or the
marriage is not lawful.

JAQUES, [coming forward]  Proceed, proceed. I'll give
her.

TOUCHSTONE  Good even, good Monsieur What-you-call-'t.
How do you, sir? You are very well met. God
'ild you for your last company. I am very glad to see
you. Even a toy in hand here, sir. Nay, pray be
covered.

JAQUES  Will you be married, motley?

TOUCHSTONE  As the ox hath his bow, sir, the horse his
curb, and the falcon her bells, so man hath his
desires; and as pigeons bill, so wedlock would be
nibbling.

JAQUES  And will you, being a man of your breeding, be
married under a bush like a beggar? Get you to
church, and have a good priest that can tell you
what marriage is. This fellow will but join you
together as they join wainscot. Then one of you will
prove a shrunk panel and, like green timber, warp,
warp.

TOUCHSTONE  I am not in the mind but I were better to
be married of him than of another, for he is not like
to marry me well, and not being well married, it
will be a good excuse for me hereafter to leave my
wife.

JAQUES  Go thou with me, and let me counsel thee.

TOUCHSTONE  Come, sweet Audrey. We must be married,
or we must live in bawdry.--Farewell, good
Master Oliver, not
	      O sweet Oliver,
	      O brave Oliver,
	   Leave me not behind thee,
But
	      Wind away,
	      Begone, I say,
	   I will not to wedding with thee.
[Audrey, Touchstone, and Jaques exit.]

OLIVER MARTEXT  'Tis no matter. Ne'er a fantastical
knave of them all shall flout me out of my calling.
[He exits.]

Scene 4
=======
[Enter Rosalind, dressed as Ganymede, and Celia,
dressed as Aliena.]


ROSALIND  Never talk to me. I will weep.

CELIA  Do, I prithee, but yet have the grace to consider
that tears do not become a man.

ROSALIND  But have I not cause to weep?

CELIA  As good cause as one would desire. Therefore
weep.

ROSALIND  His very hair is of the dissembling color.

CELIA  Something browner than Judas's. Marry, his
kisses are Judas's own children.

ROSALIND  I' faith, his hair is of a good color.

CELIA  An excellent color. Your chestnut was ever the
only color.

ROSALIND  And his kissing is as full of sanctity as the
touch of holy bread.

CELIA  He hath bought a pair of cast lips of Diana. A
nun of winter's sisterhood kisses not more religiously.
The very ice of chastity is in them.

ROSALIND  But why did he swear he would come this
morning, and comes not?

CELIA  Nay, certainly, there is no truth in him.

ROSALIND  Do you think so?

CELIA  Yes, I think he is not a pickpurse nor a horse-stealer,
but for his verity in love, I do think him as
concave as a covered goblet or a worm-eaten nut.

ROSALIND  Not true in love?

CELIA  Yes, when he is in, but I think he is not in.

ROSALIND  You have heard him swear downright he
was.

CELIA  "Was" is not "is." Besides, the oath of a lover is
no stronger than the word of a tapster. They are
both the confirmer of false reckonings. He attends
here in the forest on the Duke your father.

ROSALIND  I met the Duke yesterday and had much
question with him. He asked me of what parentage
I was. I told him, of as good as he. So he laughed
and let me go. But what talk we of fathers when
there is such a man as Orlando?

CELIA  O, that's a brave man. He writes brave verses,
speaks brave words, swears brave oaths, and breaks
them bravely, quite traverse, athwart the heart of
his lover, as a puny tilter that spurs his horse but on
one side breaks his staff like a noble goose; but all's
brave that youth mounts and folly guides.

[Enter Corin.]

Who comes here?

CORIN
Mistress and master, you have oft inquired
After the shepherd that complained of love,
Who you saw sitting by me on the turf,
Praising the proud disdainful shepherdess
That was his mistress.

CELIA, [as Aliena]  Well, and what of him?

CORIN
If you will see a pageant truly played
Between the pale complexion of true love
And the red glow of scorn and proud disdain,
Go hence a little, and I shall conduct you
If you will mark it.

ROSALIND, [aside to Celia]  O come, let us remove.
The sight of lovers feedeth those in love.
[As Ganymede, to Corin.] 
Bring us to this sight, andyou shall say
I'll prove a busy actor in their play.
[They exit.]

Scene 5
=======
[Enter Silvius and Phoebe.]


SILVIUS
Sweet Phoebe, do not scorn me. Do not, Phoebe.
Say that you love me not, but say not so
In bitterness. The common executioner,
Whose heart th' accustomed sight of death makes
hard,
Falls not the axe upon the humbled neck
But first begs pardon. Will you sterner be
Than he that dies and lives by bloody drops?

[Enter, unobserved, Rosalind as Ganymede, Celia as
Aliena, and Corin.]


PHOEBE
I would not be thy executioner.
I fly thee, for I would not injure thee.
Thou tell'st me there is murder in mine eye.
'Tis pretty, sure, and very probable
That eyes, that are the frail'st and softest things,
Who shut their coward gates on atomies,
Should be called tyrants, butchers, murderers.
Now I do frown on thee with all my heart,
And if mine eyes can wound, now let them kill thee.
Now counterfeit to swoon; why, now fall down;
Or if thou canst not, O, for shame, for shame,
Lie not, to say mine eyes are murderers.
Now show the wound mine eye hath made in thee.
Scratch thee but with a pin, and there remains
Some scar of it. Lean upon a rush,
The cicatrice and capable impressure
Thy palm some moment keeps. But now mine eyes,
Which I have darted at thee, hurt thee not;
Nor I am sure there is no force in eyes
That can do hurt.

SILVIUS  O dear Phoebe,
If ever--as that ever may be near--
You meet in some fresh cheek the power of fancy,
Then shall you know the wounds invisible
That love's keen arrows make.

PHOEBE  But till that time
Come not thou near me. And when that time
comes,
Afflict me with thy mocks, pity me not,
As till that time I shall not pity thee.

ROSALIND, [as Ganymede, coming forward]
And why, I pray you? Who might be your mother,
That you insult, exult, and all at once,
Over the wretched? What though you have no
beauty--
As, by my faith, I see no more in you
Than without candle may go dark to bed--
Must you be therefore proud and pitiless?
Why, what means this? Why do you look on me?
I see no more in you than in the ordinary
Of nature's sale-work.--'Od's my little life,
I think she means to tangle my eyes, too.--
No, faith, proud mistress, hope not after it.
'Tis not your inky brows, your black silk hair,
Your bugle eyeballs, nor your cheek of cream
That can entame my spirits to your worship.--
You foolish shepherd, wherefore do you follow her,
Like foggy south puffing with wind and rain?
You are a thousand times a properer man
Than she a woman. 'Tis such fools as you
That makes the world full of ill-favored children.
'Tis not her glass but you that flatters her,
And out of you she sees herself more proper
Than any of her lineaments can show her.--
But, mistress, know yourself. Down on your knees
And thank heaven, fasting, for a good man's love,
For I must tell you friendly in your ear,
Sell when you can; you are not for all markets.
Cry the man mercy, love him, take his offer.
Foul is most foul, being foul to be a scoffer.--
So take her to thee, shepherd. Fare you well.

PHOEBE
Sweet youth, I pray you chide a year together.
I had rather hear you chide than this man woo.

ROSALIND[,as Ganymede]  He's fall'n in love with your
foulness. [(To Silvius.)] And she'll fall in love with
my anger. If it be so, as fast as she answers thee with
frowning looks, I'll sauce her with bitter words. [(To
Phoebe.)] Why look you so upon me?

PHOEBE  For no ill will I bear you.

ROSALIND, [as Ganymede]
I pray you, do not fall in love with me,
For I am falser than vows made in wine.
Besides, I like you not. If you will know my house,
'Tis at the tuft of olives, here hard by.--
Will you go, sister?--Shepherd, ply her hard.--
Come, sister.--Shepherdess, look on him better,
And be not proud. Though all the world could see,
None could be so abused in sight as he.--
Come, to our flock.
[She exits, with Celia and Corin.]

PHOEBE, [aside]
Dead shepherd, now I find thy saw of might:
"Who ever loved that loved not at first sight?"

SILVIUS
Sweet Phoebe--

PHOEBE  Ha, what sayst thou, Silvius?

SILVIUS  Sweet Phoebe, pity me.

PHOEBE
Why, I am sorry for thee, gentle Silvius.

SILVIUS
Wherever sorrow is, relief would be.
If you do sorrow at my grief in love,
By giving love your sorrow and my grief
Were both extermined.

PHOEBE
Thou hast my love. Is not that neighborly?

SILVIUS
I would have you.

PHOEBE  Why, that were covetousness.
Silvius, the time was that I hated thee;
And yet it is not that I bear thee love;
But since that thou canst talk of love so well,
Thy company, which erst was irksome to me,
I will endure, and I'll employ thee too.
But do not look for further recompense
Than thine own gladness that thou art employed.

SILVIUS
So holy and so perfect is my love,
And I in such a poverty of grace,
That I shall think it a most plenteous crop
To glean the broken ears after the man
That the main harvest reaps. Loose now and then
A scattered smile, and that I'll live upon.

PHOEBE
Know'st thou the youth that spoke to me erewhile?

SILVIUS
Not very well, but I have met him oft,
And he hath bought the cottage and the bounds
That the old carlot once was master of.

PHOEBE
Think not I love him, though I ask for him.
'Tis but a peevish boy--yet he talks well--
But what care I for words? Yet words do well
When he that speaks them pleases those that hear.
It is a pretty youth--not very pretty--
But sure he's proud--and yet his pride becomes
him.
He'll make a proper man. The best thing in him
Is his complexion; and faster than his tongue
Did make offense, his eye did heal it up.
He is not very tall--yet for his years he's tall.
His leg is but so-so--and yet 'tis well.
There was a pretty redness in his lip,
A little riper and more lusty red
Than that mixed in his cheek: 'twas just the
difference
Betwixt the constant red and mingled damask.
There be some women, Silvius, had they marked
him
In parcels as I did, would have gone near
To fall in love with him; but for my part
I love him not nor hate him not; and yet
I have more cause to hate him than to love him.
For what had he to do to chide at me?
He said mine eyes were black and my hair black,
And now I am remembered, scorned at me.
I marvel why I answered not again.
But that's all one: omittance is no quittance.
I'll write to him a very taunting letter,
And thou shalt bear it. Wilt thou, Silvius?

SILVIUS
Phoebe, with all my heart.

PHOEBE  I'll write it straight.
The matter's in my head and in my heart.
I will be bitter with him and passing short.
Go with me, Silvius.
[They exit.]


ACT 4
=====

Scene 1
=======
[Enter Rosalind as Ganymede, and Celia as Aliena,
and Jaques.]


JAQUES  I prithee, pretty youth, let me be better
acquainted with thee.

ROSALIND, [as Ganymede]  They say you are a melancholy
fellow.

JAQUES  I am so. I do love it better than laughing.

ROSALIND, [as Ganymede]  Those that are in extremity
of either are abominable fellows and betray
themselves to every modern censure worse than
drunkards.

JAQUES  Why, 'tis good to be sad and say nothing.

ROSALIND, [as Ganymede]  Why then, 'tis good to be a
post.

JAQUES  I have neither the scholar's melancholy, which
is emulation; nor the musician's, which is fantastical;
nor the courtier's, which is proud; nor the
soldier's, which is ambitious; nor the lawyer's,
which is politic; nor the lady's, which is nice; nor
the lover's, which is all these; but it is a melancholy
of mine own, compounded of many simples, extracted
from many objects, and indeed the sundry
contemplation of my travels, in which my often
rumination wraps me in a most humorous sadness.

ROSALIND, [as Ganymede]  A traveller. By my faith, you
have great reason to be sad. I fear you have sold
your own lands to see other men's. Then to have
seen much and to have nothing is to have rich eyes
and poor hands.

JAQUES  Yes, I have gained my experience.

ROSALIND, [as Ganymede]  And your experience makes
you sad. I had rather have a fool to make me merry
than experience to make me sad--and to travel for
it too.

[Enter Orlando.]


ORLANDO
Good day and happiness, dear Rosalind.

JAQUES  Nay then, God be wi' you, an you talk in blank
verse.

ROSALIND, [as Ganymede]  Farewell, Monsieur Traveller.
Look you lisp and wear strange suits, disable all
the benefits of your own country, be out of love with
your nativity, and almost chide God for making you
that countenance you are, or I will scarce think you
have swam in a gondola.
[Jaques exits.]
Why, how now, Orlando, where have you been all
this while? You a lover? An you serve me such
another trick, never come in my sight more.

ORLANDO  My fair Rosalind, I come within an hour of
my promise.

ROSALIND, [as Ganymede]  Break an hour's promise in
love? He that will divide a minute into a thousand
parts and break but a part of the thousand part of a
minute in the affairs of love, it may be said of him
that Cupid hath clapped him o' th' shoulder, but I'll
warrant him heart-whole.

ORLANDO  Pardon me, dear Rosalind.

ROSALIND, [as Ganymede]  Nay, an you be so tardy,
come no more in my sight. I had as lief be wooed of
a snail.

ORLANDO  Of a snail?

ROSALIND, [as Ganymede]  Ay, of a snail, for though he
comes slowly, he carries his house on his head--a
better jointure, I think, than you make a woman.
Besides, he brings his destiny with him.

ORLANDO  What's that?

ROSALIND, [as Ganymede]  Why, horns, which such as
you are fain to be beholding to your wives for. But
he comes armed in his fortune and prevents the
slander of his wife.

ORLANDO  Virtue is no hornmaker, and my Rosalind is
virtuous.

ROSALIND, [as Ganymede]  And I am your Rosalind.

CELIA, [as Aliena]  It pleases him to call you so, but he
hath a Rosalind of a better leer than you.

ROSALIND, [as Ganymede, to Orlando]  Come, woo me,
woo me, for now I am in a holiday humor, and like
enough to consent. What would you say to me now
an I were your very, very Rosalind?

ORLANDO  I would kiss before I spoke.

ROSALIND, [as Ganymede]  Nay, you were better speak
first, and when you were gravelled for lack of
matter, you might take occasion to kiss. Very good
orators, when they are out, they will spit; and for
lovers lacking--God warn us--matter, the cleanliest
shift is to kiss.

ORLANDO  How if the kiss be denied?

ROSALIND, [as Ganymede]  Then she puts you to entreaty,
and there begins new matter.

ORLANDO  Who could be out, being before his beloved
mistress?

ROSALIND, [as Ganymede]  Marry, that should you if I
were your mistress, or I should think my honesty
ranker than my wit.

ORLANDO  What, of my suit?

ROSALIND, [as Ganymede]  Not out of your apparel, and
yet out of your suit. Am not I your Rosalind?

ORLANDO  I take some joy to say you are because I
would be talking of her.

ROSALIND, [as Ganymede]  Well, in her person I say I
will not have you.

ORLANDO  Then, in mine own person I die.

ROSALIND, [as Ganymede]  No, faith, die by attorney.
The poor world is almost six thousand years old,
and in all this time there was not any man died in
his own person, videlicet, in a love cause. Troilus
had his brains dashed out with a Grecian club, yet
he did what he could to die before, and he is one of
the patterns of love. Leander, he would have lived
many a fair year though Hero had turned nun, if it
had not been for a hot midsummer night, for, good
youth, he went but forth to wash him in the Hellespont
and, being taken with the cramp, was
drowned; and the foolish chroniclers of that age
found it was Hero of Sestos. But these are all lies.
Men have died from time to time and worms have
eaten them, but not for love.

ORLANDO  I would not have my right Rosalind of this
mind, for I protest her frown might kill me.

ROSALIND, [as Ganymede]  By this hand, it will not kill a
fly. But come; now I will be your Rosalind in a more
coming-on disposition, and ask me what you will, I
will grant it.

ORLANDO  Then love me, Rosalind.

ROSALIND, [as Ganymede]  Yes, faith, will I, Fridays and
Saturdays and all.

ORLANDO  And wilt thou have me?

ROSALIND, [as Ganymede]  Ay, and twenty such.

ORLANDO  What sayest thou?

ROSALIND, [as Ganymede]  Are you not good?

ORLANDO  I hope so.

ROSALIND, [as Ganymede]  Why then, can one desire
too much of a good thing?--Come, sister, you shall
be the priest and marry us.--Give me your hand,
Orlando.--What do you say, sister?

ORLANDO, [to Celia]  Pray thee marry us.

CELIA, [as Aliena]  I cannot say the words.

ROSALIND, [as Ganymede]  You must begin "Will you,
Orlando--"

CELIA, [as Aliena]  Go to.--Will you, Orlando, have to
wife this Rosalind?

ORLANDO  I will.

ROSALIND, [as Ganymede]  Ay, but when?

ORLANDO  Why now, as fast as she can marry us.

ROSALIND, [as Ganymede]  Then you must say "I take
thee, Rosalind, for wife."

ORLANDO  I take thee, Rosalind, for wife.

ROSALIND, [as Ganymede]  I might ask you for your
commission, but I do take thee, Orlando, for my
husband. There's a girl goes before the priest, and
certainly a woman's thought runs before her
actions.

ORLANDO  So do all thoughts. They are winged.

ROSALIND, [as Ganymede]  Now tell me how long you
would have her after you have possessed her?

ORLANDO  Forever and a day.

ROSALIND, [as Ganymede]  Say "a day" without the
"ever." No, no, Orlando, men are April when they
woo, December when they wed. Maids are May
when they are maids, but the sky changes when
they are wives. I will be more jealous of thee than a
Barbary cock-pigeon over his hen, more clamorous
than a parrot against rain, more newfangled than
an ape, more giddy in my desires than a monkey. I
will weep for nothing, like Diana in the fountain,
and I will do that when you are disposed to be
merry. I will laugh like a hyena, and that when thou
art inclined to sleep.

ORLANDO  But will my Rosalind do so?

ROSALIND, [as Ganymede]  By my life, she will do as I
do.

ORLANDO  O, but she is wise.

ROSALIND, [as Ganymede]  Or else she could not have
the wit to do this. The wiser, the waywarder. Make
the doors upon a woman's wit, and it will out at the
casement. Shut that, and 'twill out at the keyhole.
Stop that, 'twill fly with the smoke out at the
chimney.

ORLANDO  A man that had a wife with such a wit, he
might say "Wit, whither wilt?"

ROSALIND, [as Ganymede]  Nay, you might keep that
check for it till you met your wife's wit going to
your neighbor's bed.

ORLANDO  And what wit could wit have to excuse that?

ROSALIND, [as Ganymede]  Marry, to say she came to
seek you there. You shall never take her without her
answer unless you take her without her tongue. O,
that woman that cannot make her fault her husband's
occasion, let her never nurse her child
herself, for she will breed it like a fool.

ORLANDO  For these two hours, Rosalind, I will leave
thee.

ROSALIND, [as Ganymede]  Alas, dear love, I cannot lack
thee two hours.

ORLANDO  I must attend the Duke at dinner. By two
o'clock I will be with thee again.

ROSALIND, [as Ganymede]  Ay, go your ways, go your
ways. I knew what you would prove. My friends told
me as much, and I thought no less. That flattering
tongue of yours won me. 'Tis but one cast away, and
so, come, death. Two o'clock is your hour?

ORLANDO  Ay, sweet Rosalind.

ROSALIND, [as Ganymede]  By my troth, and in good
earnest, and so God mend me, and by all pretty
oaths that are not dangerous, if you break one jot of
your promise or come one minute behind your
hour, I will think you the most pathetical break-promise,
and the most hollow lover, and the most
unworthy of her you call Rosalind that may be
chosen out of the gross band of the unfaithful.
Therefore beware my censure, and keep your
promise.

ORLANDO  With no less religion than if thou wert indeed
my Rosalind. So, adieu.

ROSALIND, [as Ganymede]  Well, time is the old justice
that examines all such offenders, and let time try.
Adieu.
[Orlando exits.]

CELIA  You have simply misused our sex in your love-prate.
We must have your doublet and hose plucked
over your head and show the world what the bird
hath done to her own nest.

ROSALIND  O coz, coz, coz, my pretty little coz, that thou
didst know how many fathom deep I am in love. But
it cannot be sounded; my affection hath an
unknown bottom, like the Bay of Portugal.

CELIA  Or rather bottomless, that as fast as you pour
affection in, it runs out.

ROSALIND  No, that same wicked bastard of Venus, that
was begot of thought, conceived of spleen, and born
of madness, that blind rascally boy that abuses
everyone's eyes because his own are out, let him be
judge how deep I am in love. I'll tell thee, Aliena, I
cannot be out of the sight of Orlando. I'll go find a
shadow and sigh till he come.

CELIA  And I'll sleep.
[They exit.]

Scene 2
=======
[Enter Jaques and Lords, like foresters.]


JAQUES  Which is he that killed the deer?

FIRST LORD  Sir, it was I.

JAQUES, [to the other Lords]  Let's present him to the
Duke like a Roman conqueror. And it would do well
to set the deer's horns upon his head for a branch of
victory.--Have you no song, forester, for this
purpose?

SECOND LORD  Yes, sir.

JAQUES  Sing it. 'Tis no matter how it be in tune, so it
make noise enough.

Music. Song.


SECOND LORD [sings]
	What shall he have that killed the deer?
	His leather skin and horns to wear.
	   Then sing him home.

[The rest shall bear this burden:]


	Take thou no scorn to wear the horn.
	It was a crest ere thou wast born.
	   Thy father's father wore it,
	   And thy father bore it.
	The horn, the horn, the lusty horn
	Is not a thing to laugh to scorn.
[They exit.]

Scene 3
=======
[Enter Rosalind dressed as Ganymede and Celia
dressed as Aliena.]


ROSALIND  How say you now? Is it not past two o'clock?
And here much Orlando.

CELIA  I warrant you, with pure love and troubled brain
he hath ta'en his bow and arrows and is gone forth
to sleep.

[Enter Silvius.]

Look who comes here.

SILVIUS, [to Rosalind]
My errand is to you, fair youth.
My gentle Phoebe did bid me give you this.
[He gives Rosalind a paper.]
I know not the contents, but as I guess
By the stern brow and waspish action
Which she did use as she was writing of it,
It bears an angry tenor. Pardon me.
I am but as a guiltless messenger.
[Rosalind reads the letter.]

ROSALIND, [as Ganymede]
Patience herself would startle at this letter
And play the swaggerer. Bear this, bear all.
She says I am not fair, that I lack manners.
She calls me proud, and that she could not love me
Were man as rare as phoenix. 'Od's my will,
Her love is not the hare that I do hunt.
Why writes she so to me? Well, shepherd, well,
This is a letter of your own device.

SILVIUS
No, I protest. I know not the contents.
Phoebe did write it.

ROSALIND, [as Ganymede]  Come, come, you are a
fool,
And turned into the extremity of love.
I saw her hand. She has a leathern hand,
A freestone-colored hand. I verily did think
That her old gloves were on, but 'twas her hands.
She has a huswife's hand--but that's no matter.
I say she never did invent this letter.
This is a man's invention, and his hand.

SILVIUS  Sure it is hers.

ROSALIND, [as Ganymede]
Why, 'tis a boisterous and a cruel style,
A style for challengers. Why, she defies me
Like Turk to Christian. Women's gentle brain
Could not drop forth such giant-rude invention,
Such Ethiop words, blacker in their effect
Than in their countenance. Will you hear the letter?

SILVIUS
So please you, for I never heard it yet,
Yet heard too much of Phoebe's cruelty.

ROSALIND, [as Ganymede]
She Phoebes me. Mark how the tyrant writes.
[Read.]
	Art thou god to shepherd turned,
	That a maiden's heart hath burned?
Can a woman rail thus?

SILVIUS  Call you this railing?

ROSALIND, [as Ganymede]
[Read.]
	Why, thy godhead laid apart,
	Warr'st thou with a woman's heart?
Did you ever hear such railing?
	Whiles the eye of man did woo me,
	That could do no vengeance to me.
Meaning me a beast.
	If the scorn of your bright eyne
	Have power to raise such love in mine,
	Alack, in me what strange effect
	Would they work in mild aspect?
	Whiles you chid me, I did love.
	How then might your prayers move?
	He that brings this love to thee
	Little knows this love in me,
	And by him seal up thy mind
	Whether that thy youth and kind
	Will the faithful offer take
	Of me, and all that I can make,
	Or else by him my love deny,
	And then I'll study how to die.

SILVIUS  Call you this chiding?

CELIA, [as Aliena]  Alas, poor shepherd.

ROSALIND, [as Ganymede]  Do you pity him? No, he
deserves no pity.--Wilt thou love such a woman?
What, to make thee an instrument and play false
strains upon thee? Not to be endured. Well, go your
way to her, for I see love hath made thee a tame
snake, and say this to her: that if she love me, I
charge her to love thee; if she will not, I will never
have her unless thou entreat for her. If you be a true
lover, hence, and not a word, for here comes more
company.	[Silvius exits.]

[Enter Oliver.]


OLIVER
Good morrow, fair ones. Pray you, if you know,
Where in the purlieus of this forest stands
A sheepcote fenced about with olive trees?

CELIA, [as Aliena]
West of this place, down in the neighbor bottom;
The rank of osiers by the murmuring stream
Left on your right hand brings you to the place.
But at this hour the house doth keep itself.
There's none within.

OLIVER
If that an eye may profit by a tongue,
Then should I know you by description--
Such garments, and such years. "The boy is fair,
Of female favor, and bestows himself
Like a ripe sister; the woman low
And browner than her brother." Are not you
The owner of the house I did inquire for?

CELIA, [as Aliena]
It is no boast, being asked, to say we are.

OLIVER
Orlando doth commend him to you both,
And to that youth he calls his Rosalind
He sends this bloody napkin. Are you he?
[He shows a stained handkerchief.]

ROSALIND, [as Ganymede]
I am. What must we understand by this?

OLIVER
Some of my shame, if you will know of me
What man I am, and how, and why, and where
This handkercher was stained.

CELIA, [as Aliena]  I pray you tell it.

OLIVER
When last the young Orlando parted from you,
He left a promise to return again
Within an hour, and pacing through the forest,
Chewing the food of sweet and bitter fancy,
Lo, what befell. He threw his eye aside--
And mark what object did present itself:
Under an old oak, whose boughs were mossed with
age
And high top bald with dry antiquity,
A wretched, ragged man, o'ergrown with hair,
Lay sleeping on his back. About his neck
A green and gilded snake had wreathed itself,
Who with her head, nimble in threats, approached
The opening of his mouth. But suddenly,
Seeing Orlando, it unlinked itself
And, with indented glides, did slip away
Into a bush, under which bush's shade
A lioness, with udders all drawn dry,
Lay couching, head on ground, with catlike watch
When that the sleeping man should stir--for 'tis
The royal disposition of that beast
To prey on nothing that doth seem as dead.
This seen, Orlando did approach the man
And found it was his brother, his elder brother.

CELIA, [as Aliena]
O, I have heard him speak of that same brother,
And he did render him the most unnatural
That lived amongst men.

OLIVER  And well he might so do,
For well I know he was unnatural.

ROSALIND, [as Ganymede]
But to Orlando: did he leave him there,
Food to the sucked and hungry lioness?

OLIVER
Twice did he turn his back and purposed so,
But kindness, nobler ever than revenge,
And nature, stronger than his just occasion,
Made him give battle to the lioness,
Who quickly fell before him; in which hurtling,
From miserable slumber I awaked.

CELIA, [as Aliena]  Are you his brother?

ROSALIND, [as Ganymede]  Was 't you he rescued?

CELIA, [as Aliena]
Was 't you that did so oft contrive to kill him?

OLIVER
'Twas I, but 'tis not I. I do not shame
To tell you what I was, since my conversion
So sweetly tastes, being the thing I am.

ROSALIND, [as Ganymede]
But for the bloody napkin?

OLIVER  By and by.
When from the first to last betwixt us two
Tears our recountments had most kindly bathed--
As how I came into that desert place--
In brief, he led me to the gentle duke,
Who gave me fresh array and entertainment,
Committing me unto my brother's love;
Who led me instantly unto his cave,
There stripped himself, and here upon his arm
The lioness had torn some flesh away,
Which all this while had bled; and now he fainted,
And cried in fainting upon Rosalind.
Brief, I recovered him, bound up his wound,
And after some small space, being strong at heart,
He sent me hither, stranger as I am,
To tell this story, that you might excuse
His broken promise, and to give this napkin
Dyed in his blood unto the shepherd youth
That he in sport doth call his Rosalind.
[Rosalind faints.]

CELIA, [as Aliena]
Why, how now, Ganymede, sweet Ganymede?

OLIVER
Many will swoon when they do look on blood.

CELIA, [as Aliena]
There is more in it.--Cousin Ganymede.

OLIVER  Look, he recovers.

ROSALIND  I would I were at home.

CELIA, [as Aliena]  We'll lead you thither.--I pray you,
will you take him by the arm?

OLIVER, [helping Rosalind to rise]  Be of good cheer,
youth. You a man? You lack a man's heart.

ROSALIND, [as Ganymede]  I do so, I confess it. Ah,
sirrah, a body would think this was well-counterfeited.
I pray you tell your brother how well I
counterfeited. Heigh-ho.

OLIVER  This was not counterfeit. There is too great
testimony in your complexion that it was a passion
of earnest.

ROSALIND, [as Ganymede]  Counterfeit, I assure you.

OLIVER  Well then, take a good heart, and counterfeit to
be a man.

ROSALIND, [as Ganymede]  So I do; but, i' faith, I should
have been a woman by right.

CELIA, [as Aliena]  Come, you look paler and paler. Pray
you draw homewards.--Good sir, go with us.

OLIVER
That will I, for I must bear answer back
How you excuse my brother, Rosalind.

ROSALIND, [as Ganymede]  I shall devise something.
But I pray you commend my counterfeiting to him.
Will you go?
[They exit.]


ACT 5
=====

Scene 1
=======
[Enter Touchstone and Audrey.]


TOUCHSTONE  We shall find a time, Audrey. Patience,
gentle Audrey.

AUDREY  Faith, the priest was good enough, for all the
old gentleman's saying.

TOUCHSTONE  A most wicked Sir Oliver, Audrey, a most
vile Martext. But Audrey, there is a youth here in
the forest lays claim to you.

AUDREY  Ay, I know who 'tis. He hath no interest in me
in the world.

[Enter William.]

Here comes the man you mean.

TOUCHSTONE  It is meat and drink to me to see a clown.
By my troth, we that have good wits have much to
answer for. We shall be flouting. We cannot hold.

WILLIAM  Good ev'n, Audrey.

AUDREY  God gi' good ev'n, William.

WILLIAM, [to Touchstone]  And good ev'n to you, sir.

TOUCHSTONE  Good ev'n, gentle friend. Cover thy head,
cover thy head. Nay, prithee, be covered. How old
are you, friend?

WILLIAM  Five-and-twenty, sir.

TOUCHSTONE  A ripe age. Is thy name William?

WILLIAM  William, sir.

TOUCHSTONE  A fair name. Wast born i' th' forest here?

WILLIAM  Ay, sir, I thank God.

TOUCHSTONE  "Thank God." A good answer. Art rich?

WILLIAM  'Faith sir, so-so.

TOUCHSTONE  "So-so" is good, very good, very excellent
good. And yet it is not: it is but so-so. Art thou wise?

WILLIAM  Ay, sir, I have a pretty wit.

TOUCHSTONE  Why, thou sayst well. I do now remember
a saying: "The fool doth think he is wise, but the
wise man knows himself to be a fool." The heathen
philosopher, when he had a desire to eat a grape,
would open his lips when he put it into his mouth,
meaning thereby that grapes were made to eat and
lips to open. You do love this maid?

WILLIAM  I do, sir.

TOUCHSTONE  Give me your hand. Art thou learned?

WILLIAM  No, sir.

TOUCHSTONE  Then learn this of me: to have is to have.
For it is a figure in rhetoric that drink, being poured
out of a cup into a glass, by filling the one doth
empty the other. For all your writers do consent
that ipse is "he." Now, you are not ipse, for I am he.

WILLIAM  Which he, sir?

TOUCHSTONE  He, sir, that must marry this woman.
Therefore, you clown, abandon--which is in the
vulgar "leave"--the society--which in the boorish
is "company"--of this female--which in the common
is "woman"; which together is, abandon the
society of this female, or, clown, thou perishest; or,
to thy better understanding, diest; or, to wit, I kill
thee, make thee away, translate thy life into death,
thy liberty into bondage. I will deal in poison with
thee, or in bastinado, or in steel. I will bandy with
thee in faction. I will o'errun thee with policy. I
will kill thee a hundred and fifty ways. Therefore
tremble and depart.

AUDREY  Do, good William.

WILLIAM, [to Touchstone]  God rest you merry, sir.
[He exits.]

[Enter Corin.]


CORIN  Our master and mistress seeks you. Come away,
away.

TOUCHSTONE  Trip, Audrey, trip, Audrey.--I attend, I
attend.
[They exit.]

Scene 2
=======
[Enter Orlando, with his arm in a sling, and Oliver.]


ORLANDO  Is 't possible that on so little acquaintance
you should like her? That, but seeing, you should
love her? And loving, woo? And wooing, she should
grant? And will you persever to enjoy her?

OLIVER  Neither call the giddiness of it in question, the
poverty of her, the small acquaintance, my sudden
wooing, nor her sudden consenting, but say with
me "I love Aliena"; say with her that she loves me;
consent with both that we may enjoy each other. It
shall be to your good, for my father's house and all
the revenue that was old Sir Rowland's will I estate
upon you, and here live and die a shepherd.

[Enter Rosalind, as Ganymede.]


ORLANDO  You have my consent. Let your wedding be
tomorrow. Thither will I invite the Duke and all 's
contented followers. Go you and prepare Aliena,
for, look you, here comes my Rosalind.

ROSALIND, [as Ganymede, to Oliver]  God save you,
brother.

OLIVER  And you, fair sister.	[He exits.]

ROSALIND, [as Ganymede]  O my dear Orlando, how it
grieves me to see thee wear thy heart in a scarf.

ORLANDO  It is my arm.

ROSALIND, [as Ganymede]  I thought thy heart had been
wounded with the claws of a lion.

ORLANDO  Wounded it is, but with the eyes of a lady.

ROSALIND, [as Ganymede]  Did your brother tell you
how I counterfeited to swoon when he showed me
your handkercher?

ORLANDO  Ay, and greater wonders than that.

ROSALIND, [as Ganymede]  O, I know where you are.
Nay, 'tis true. There was never anything so sudden
but the fight of two rams, and Caesar's thrasonical
brag of "I came, saw, and overcame." For your
brother and my sister no sooner met but they
looked, no sooner looked but they loved, no sooner
loved but they sighed, no sooner sighed but they
asked one another the reason, no sooner knew the
reason but they sought the remedy; and in these
degrees have they made a pair of stairs to marriage,
which they will climb incontinent, or else be incontinent
before marriage. They are in the very wrath
of love, and they will together. Clubs cannot part
them.

ORLANDO  They shall be married tomorrow, and I will
bid the Duke to the nuptial. But O, how bitter a
thing it is to look into happiness through another
man's eyes. By so much the more shall I tomorrow
be at the height of heart-heaviness by how much I
shall think my brother happy in having what he
wishes for.

ROSALIND, [as Ganymede]  Why, then, tomorrow I cannot
serve your turn for Rosalind?

ORLANDO  I can live no longer by thinking.

ROSALIND, [as Ganymede]  I will weary you then no
longer with idle talking. Know of me then--for
now I speak to some purpose--that I know you are
a gentleman of good conceit. I speak not this that
you should bear a good opinion of my knowledge,
insomuch I say I know you are. Neither do I labor
for a greater esteem than may in some little measure
draw a belief from you to do yourself good, and
not to grace me. Believe then, if you please, that I
can do strange things. I have, since I was three year
old, conversed with a magician, most profound in
his art and yet not damnable. If you do love Rosalind
so near the heart as your gesture cries it out,
when your brother marries Aliena shall you marry
her. I know into what straits of fortune she is
driven, and it is not impossible to me, if it appear
not inconvenient to you, to set her before your eyes
tomorrow, human as she is, and without any
danger.

ORLANDO  Speak'st thou in sober meanings?

ROSALIND, [as Ganymede]  By my life I do, which I
tender dearly, though I say I am a magician. Therefore
put you in your best array, bid your friends; for
if you will be married tomorrow, you shall, and to
Rosalind, if you will.

[Enter Silvius and Phoebe.]

Look, here comes a lover of mine and a lover of
hers.

PHOEBE, [to Rosalind]
Youth, you have done me much ungentleness
To show the letter that I writ to you.

ROSALIND, [as Ganymede]
I care not if I have. It is my study
To seem despiteful and ungentle to you.
You are there followed by a faithful shepherd.
Look upon him, love him; he worships you.

PHOEBE, [to Silvius]
Good shepherd, tell this youth what 'tis to love.

SILVIUS
It is to be all made of sighs and tears,
And so am I for Phoebe.

PHOEBE  And I for Ganymede.

ORLANDO  And I for Rosalind.

ROSALIND, [as Ganymede]  And I for no woman.

SILVIUS
It is to be all made of faith and service,
And so am I for Phoebe.

PHOEBE  And I for Ganymede.

ORLANDO  And I for Rosalind.

ROSALIND, [as Ganymede]  And I for no woman.

SILVIUS
It is to be all made of fantasy,
All made of passion and all made of wishes,
All adoration, duty, and observance,
All humbleness, all patience and impatience,
All purity, all trial, all observance,
And so am I for Phoebe.

PHOEBE  And so am I for Ganymede.

ORLANDO  And so am I for Rosalind.

ROSALIND, [as Ganymede]  And so am I for no
woman.

PHOEBE
If this be so, why blame you me to love you?

SILVIUS
If this be so, why blame you me to love you?

ORLANDO
If this be so, why blame you me to love you?

ROSALIND, [as Ganymede]  Why do you speak too,
"Why blame you me to love you?"

ORLANDO  To her that is not here, nor doth not hear.

ROSALIND, [as Ganymede]  Pray you, no more of this.
'Tis like the howling of Irish wolves against the
moon. [(To Silvius.)] I will help you if I can. [(To
Phoebe.)] I would love you if I could.--Tomorrow
meet me all together. [(To Phoebe.)] I will marry
you if ever I marry woman, and I'll be married
tomorrow. [(To Orlando.)] I will satisfy you if ever I
satisfy man, and you shall be married tomorrow.
[(To Silvius.)] I will content you, if what pleases you
contents you, and you shall be married tomorrow.
[(To Orlando.)] As you love Rosalind, meet. [(To
Silvius.)] As you love Phoebe, meet.--And as I love
no woman, I'll meet. So fare you well. I have left
you commands.

SILVIUS  I'll not fail, if I live.

PHOEBE  Nor I.

ORLANDO  Nor I.
[They exit.]

Scene 3
=======
[Enter Touchstone and Audrey.]


TOUCHSTONE  Tomorrow is the joyful day, Audrey. Tomorrow
will we be married.

AUDREY  I do desire it with all my heart, and I hope it is
no dishonest desire to desire to be a woman of the
world.

[Enter two Pages.]

Here come two of the banished duke's pages.

FIRST PAGE  Well met, honest gentleman.

TOUCHSTONE  By my troth, well met. Come, sit, sit, and
a song.

SECOND PAGE  We are for you. Sit i' th' middle.
[They sit.]

FIRST PAGE  Shall we clap into 't roundly, without
hawking or spitting or saying we are hoarse, which
are the only prologues to a bad voice?

SECOND PAGE  I' faith, i' faith, and both in a tune like
two gypsies on a horse.

Song.


PAGES [sing]
	It was a lover and his lass,
	   With a hey, and a ho, and a hey-nonny-no,
	That o'er the green cornfield did pass
	   In springtime, the only pretty ring time,
	When birds do sing, hey ding a ding, ding.
	Sweet lovers love the spring.

	Between the acres of the rye,
	   With a hey, and a ho, and a hey-nonny-no,
	These pretty country folks would lie
	   In springtime, the only pretty ring time,
	When birds do sing, hey ding a ding, ding.
	Sweet lovers love the spring.

	This carol they began that hour,
	   With a hey, and a ho, and a hey-nonny-no,
	How that a life was but a flower
	   In springtime, the only pretty ring time,
	When birds do sing, hey ding a ding, ding.
	Sweet lovers love the spring.

	And therefore take the present time,
	   With a hey, and a ho, and a hey-nonny-no,
	For love is crowned with the prime,
	   In springtime, the only pretty ring time,
	When birds do sing, hey ding a ding, ding.
	Sweet lovers love the spring.

TOUCHSTONE  Truly, young gentlemen, though there
was no great matter in the ditty, yet the note was
very untunable.

FIRST PAGE  You are deceived, sir. We kept time. We lost
not our time.

TOUCHSTONE  By my troth, yes. I count it but time lost
to hear such a foolish song. God be wi' you, and
God mend your voices.--Come, Audrey.
[They rise and exit.]

Scene 4
=======
[Enter Duke Senior, Amiens, Jaques, Orlando, Oliver,
and Celia as Aliena.]


DUKE SENIOR
Dost thou believe, Orlando, that the boy
Can do all this that he hath promised?

ORLANDO
I sometimes do believe and sometimes do not,
As those that fear they hope, and know they fear.

[Enter Rosalind as Ganymede, Silvius, and Phoebe.]


ROSALIND, [as Ganymede]
Patience once more whiles our compact is urged.
[To Duke.] You say, if I bring in your Rosalind,
You will bestow her on Orlando here?

DUKE SENIOR
That would I, had I kingdoms to give with her.

ROSALIND, [as Ganymede, to Orlando]
And you say you will have her when I bring her?

ORLANDO
That would I, were I of all kingdoms king.

ROSALIND, [as Ganymede, to Phoebe]
You say you'll marry me if I be willing?

PHOEBE
That will I, should I die the hour after.

ROSALIND, [as Ganymede]
But if you do refuse to marry me,
You'll give yourself to this most faithful shepherd?

PHOEBE  So is the bargain.

ROSALIND, [as Ganymede, to Silvius]
You say that you'll have Phoebe if she will?

SILVIUS
Though to have her and death were both one thing.

ROSALIND, [as Ganymede]
I have promised to make all this matter even.
Keep you your word, O duke, to give your
daughter,--
You yours, Orlando, to receive his daughter.--
Keep you your word, Phoebe, that you'll marry me,
Or else, refusing me, to wed this shepherd.--
Keep your word, Silvius, that you'll marry her
If she refuse me. And from hence I go
To make these doubts all even.
[Rosalind and Celia exit.]

DUKE SENIOR
I do remember in this shepherd boy
Some lively touches of my daughter's favor.

ORLANDO
My lord, the first time that I ever saw him
Methought he was a brother to your daughter.
But, my good lord, this boy is forest-born
And hath been tutored in the rudiments
Of many desperate studies by his uncle,
Whom he reports to be a great magician
Obscured in the circle of this forest.

[Enter Touchstone and Audrey.]


JAQUES  There is sure another flood toward, and these
couples are coming to the ark. Here comes a pair of
very strange beasts, which in all tongues are called
fools.

TOUCHSTONE  Salutation and greeting to you all.

JAQUES, [to Duke]  Good my lord, bid him welcome.
This is the motley-minded gentleman that I have so
often met in the forest. He hath been a courtier, he
swears.

TOUCHSTONE  If any man doubt that, let him put me to
my purgation. I have trod a measure. I have flattered
a lady. I have been politic with my friend,
smooth with mine enemy. I have undone three
tailors. I have had four quarrels, and like to have
fought one.

JAQUES  And how was that ta'en up?

TOUCHSTONE  Faith, we met and found the quarrel was
upon the seventh cause.

JAQUES  How "seventh cause"?--Good my lord, like
this fellow.

DUKE SENIOR  I like him very well.

TOUCHSTONE  God 'ild you, sir. I desire you of the like. I
press in here, sir, amongst the rest of the country
copulatives, to swear and to forswear, according as
marriage binds and blood breaks. A poor virgin, sir,
an ill-favored thing, sir, but mine own. A poor
humor of mine, sir, to take that that no man else
will. Rich honesty dwells like a miser, sir, in a poor
house, as your pearl in your foul oyster.

DUKE SENIOR  By my faith, he is very swift and
sententious.

TOUCHSTONE  According to the fool's bolt, sir, and such
dulcet diseases.

JAQUES  But for the seventh cause. How did you find the
quarrel on the seventh cause?

TOUCHSTONE  Upon a lie seven times removed.--Bear
your body more seeming, Audrey.--As thus, sir: I
did dislike the cut of a certain courtier's beard. He
sent me word if I said his beard was not cut well, he
was in the mind it was. This is called "the retort
courteous." If I sent him word again it was not well
cut, he would send me word he cut it to please
himself. This is called "the quip modest." If again it
was not well cut, he disabled my judgment. This is
called "the reply churlish." If again it was not well
cut, he would answer I spake not true. This is called
"the reproof valiant." If again it was not well cut, he
would say I lie. This is called "the countercheck
quarrelsome," and so to "the lie circumstantial,"
and "the lie direct."

JAQUES  And how oft did you say his beard was not well
cut?

TOUCHSTONE  I durst go no further than the lie circumstantial,
nor he durst not give me the lie direct, and
so we measured swords and parted.

JAQUES  Can you nominate in order now the degrees of
the lie?

TOUCHSTONE  O sir, we quarrel in print, by the book, as
you have books for good manners. I will name you
the degrees: the first, "the retort courteous"; the
second, "the quip modest"; the third, "the reply
churlish"; the fourth, "the reproof valiant"; the
fifth, "the countercheck quarrelsome"; the sixth,
"the lie with circumstance"; the seventh, "the lie
direct." All these you may avoid but the lie direct,
and you may avoid that too with an "if." I knew
when seven justices could not take up a quarrel, but
when the parties were met themselves, one of them
thought but of an "if," as: "If you said so, then I said
so." And they shook hands and swore brothers.
Your "if" is the only peacemaker: much virtue in
"if."

JAQUES, [to Duke]  Is not this a rare fellow, my lord?
He's as good at anything and yet a fool.

DUKE SENIOR  He uses his folly like a stalking-horse,
and under the presentation of that he shoots his wit.

[Enter Hymen, Rosalind, and Celia. Still music.]


HYMEN
	Then is there mirth in heaven
	When earthly things made even
	   Atone together.
	Good duke, receive thy daughter.
	Hymen from heaven brought her,
	   Yea, brought her hither,
	That thou mightst join her hand with his,
	Whose heart within his bosom is.

ROSALIND, [to Duke]
To you I give myself, for I am yours.
[To Orlando.] To you I give myself, for I am yours.

DUKE SENIOR
If there be truth in sight, you are my daughter.

ORLANDO
If there be truth in sight, you are my Rosalind.

PHOEBE
If sight and shape be true,
Why then, my love adieu.

ROSALIND, [to Duke]
I'll have no father, if you be not he.
[To Orlando.] I'll have no husband, if you be not he,
[To Phoebe.] Nor ne'er wed woman, if you be not
she.

HYMEN
	Peace, ho! I bar confusion.
	'Tis I must make conclusion
	   Of these most strange events.
	Here's eight that must take hands
	To join in Hymen's bands,
	   If truth holds true contents.
[To Rosalind and Orlando.]
	You and you no cross shall part.
[To Celia and Oliver.]
	You and you are heart in heart.
[To Phoebe.]
	You to his love must accord
	Or have a woman to your lord.
[To Audrey and Touchstone.]
	You and you are sure together
	As the winter to foul weather.
[To All.]
	Whiles a wedlock hymn we sing,
	Feed yourselves with questioning,
	That reason wonder may diminish
	How thus we met, and these things finish.

Song.


	Wedding is great Juno's crown,
	   O blessed bond of board and bed.
	'Tis Hymen peoples every town.
	   High wedlock then be honored.
	Honor, high honor, and renown
	To Hymen, god of every town.


DUKE SENIOR, [to Celia]
O my dear niece, welcome thou art to me,
Even daughter, welcome in no less degree.

PHOEBE, [to Silvius]
I will not eat my word. Now thou art mine,
Thy faith my fancy to thee doth combine.

[Enter Second Brother, Jaques de Boys.]


SECOND BROTHER
Let me have audience for a word or two.
I am the second son of old Sir Rowland,
That bring these tidings to this fair assembly.
Duke Frederick, hearing how that every day
Men of great worth resorted to this forest,
Addressed a mighty power, which were on foot
In his own conduct, purposely to take
His brother here and put him to the sword;
And to the skirts of this wild wood he came,
Where, meeting with an old religious man,
After some question with him, was converted
Both from his enterprise and from the world,
His crown bequeathing to his banished brother,
And all their lands restored to them again
That were with him exiled. This to be true
I do engage my life.

DUKE SENIOR  Welcome, young man.
Thou offer'st fairly to thy brothers' wedding:
To one his lands withheld, and to the other
A land itself at large, a potent dukedom.--
First, in this forest let us do those ends
That here were well begun and well begot,
And, after, every of this happy number
That have endured shrewd days and nights with us
Shall share the good of our returned fortune
According to the measure of their states.
Meantime, forget this new-fall'n dignity,
And fall into our rustic revelry.--
Play, music.--And you brides and bridegrooms all,
With measure heaped in joy to th' measures fall.

JAQUES, [to Second Brother]
Sir, by your patience: if I heard you rightly,
The Duke hath put on a religious life
And thrown into neglect the pompous court.

SECOND BROTHER  He hath.

JAQUES
To him will I. Out of these convertites
There is much matter to be heard and learned.
[To Duke.] You to your former honor I bequeath;
Your patience and your virtue well deserves it.
[To Orlando.] You to a love that your true faith doth
merit.
[To Oliver.] You to your land, and love, and great
allies.
[To Silvius.] You to a long and well-deserved bed.
[To Touchstone.] And you to wrangling, for thy
loving voyage
Is but for two months victualled.--So to your
pleasures.
I am for other than for dancing measures.

DUKE SENIOR  Stay, Jaques, stay.

JAQUES
To see no pastime, I. What you would have
I'll stay to know at your abandoned cave.	[He exits.]

DUKE SENIOR
Proceed, proceed. We'll begin these rites,
As we do trust they'll end, in true delights.
[Dance. All but Rosalind exit.]

EPILOGUE.
=========

ROSALIND  It is not the fashion to see the lady the
epilogue, but it is no more unhandsome than to see
the lord the prologue. If it be true that good wine
needs no bush, 'tis true that a good play needs no
epilogue. Yet to good wine they do use good bushes,
and good plays prove the better by the help of good
epilogues. What a case am I in then that am neither
a good epilogue nor cannot insinuate with you in
the behalf of a good play! I am not furnished like a
beggar; therefore to beg will not become me. My
way is to conjure you, and I'll begin with the
women. I charge you, O women, for the love you
bear to men, to like as much of this play as please
you. And I charge you, O men, for the love you bear
to women--as I perceive by your simpering, none
of you hates them--that between you and the
women the play may please. If I were a woman, I
would kiss as many of you as had beards that
pleased me, complexions that liked me, and breaths
that I defied not. And I am sure as many as have
good beards, or good faces, or sweet breaths will for
my kind offer, when I make curtsy, bid me farewell.
[She exits.]
`

	b["hamlet_TXT_FolgerShakespeare.txt"] = `
Hamlet
by William Shakespeare
Edited by Barbara A. Mowat and Paul Werstine
  with Michael Poston and Rebecca Niles
Folger Shakespeare Library
https://shakespeare.folger.edu/shakespeares-works/hamlet/
Created on Jul 31, 2015, from FDT version 0.9.2

Characters in the Play
======================
THE GHOST
HAMLET, Prince of Denmark, son of the late King Hamlet and Queen Gertrude
QUEEN GERTRUDE, widow of King Hamlet, now married to Claudius
KING CLAUDIUS, brother to the late King Hamlet
OPHELIA
LAERTES, her brother
POLONIUS, father of Ophelia and Laertes, councillor to King Claudius
REYNALDO, servant to Polonius
HORATIO, Hamlet's friend and confidant
Courtiers at the Danish court:
  VOLTEMAND
  CORNELIUS
  ROSENCRANTZ
  GUILDENSTERN
  OSRIC
  Gentlemen
  A Lord
Danish soldiers:
  FRANCISCO
  BARNARDO
  MARCELLUS
FORTINBRAS, Prince of Norway
A Captain in Fortinbras's army
Ambassadors to Denmark from England
Players who take the roles of Prologue, Player King, Player Queen, and Lucianus in <title>The Murder of Gonzago</title>
Two Messengers
Sailors
Gravedigger
Gravedigger's companion
Doctor of Divinity
Attendants, Lords, Guards, Musicians, Laertes's Followers, Soldiers, Officers


ACT 1
=====

Scene 1
=======
[Enter Barnardo and Francisco, two sentinels.]


BARNARDO  Who's there?

FRANCISCO
Nay, answer me. Stand and unfold yourself.

BARNARDO  Long live the King!

FRANCISCO  Barnardo?

BARNARDO  He.

FRANCISCO
You come most carefully upon your hour.

BARNARDO
'Tis now struck twelve. Get thee to bed, Francisco.

FRANCISCO
For this relief much thanks. 'Tis bitter cold,
And I am sick at heart.

BARNARDO  Have you had quiet guard?

FRANCISCO  Not a mouse stirring.

BARNARDO  Well, good night.
If you do meet Horatio and Marcellus,
The rivals of my watch, bid them make haste.

[Enter Horatio and Marcellus.]


FRANCISCO
I think I hear them.--Stand ho! Who is there?

HORATIO  Friends to this ground.

MARCELLUS  And liegemen to the Dane.

FRANCISCO  Give you good night.

MARCELLUS
O farewell, honest soldier. Who hath relieved
you?

FRANCISCO
Barnardo hath my place. Give you good night.
[Francisco exits.]

MARCELLUS  Holla, Barnardo.

BARNARDO  Say, what, is Horatio there?

HORATIO  A piece of him.

BARNARDO
Welcome, Horatio.--Welcome, good Marcellus.

HORATIO
What, has this thing appeared again tonight?

BARNARDO  I have seen nothing.

MARCELLUS
Horatio says 'tis but our fantasy
And will not let belief take hold of him
Touching this dreaded sight twice seen of us.
Therefore I have entreated him along
With us to watch the minutes of this night,
That, if again this apparition come,
He may approve our eyes and speak to it.

HORATIO
Tush, tush, 'twill not appear.

BARNARDO  Sit down awhile,
And let us once again assail your ears,
That are so fortified against our story,
What we have two nights seen.

HORATIO  Well, sit we down,
And let us hear Barnardo speak of this.

BARNARDO  Last night of all,
When yond same star that's westward from the pole
Had made his course t' illume that part of heaven
Where now it burns, Marcellus and myself,
The bell then beating one--

[Enter Ghost.]


MARCELLUS
Peace, break thee off! Look where it comes again.

BARNARDO
In the same figure like the King that's dead.

MARCELLUS, [to Horatio]
Thou art a scholar. Speak to it, Horatio.

BARNARDO
Looks he not like the King? Mark it, Horatio.

HORATIO
Most like. It harrows me with fear and wonder.

BARNARDO
It would be spoke to.

MARCELLUS  Speak to it, Horatio.

HORATIO
What art thou that usurp'st this time of night,
Together with that fair and warlike form
In which the majesty of buried Denmark
Did sometimes march? By heaven, I charge thee,
speak.

MARCELLUS
It is offended.

BARNARDO  See, it stalks away.

HORATIO
Stay! speak! speak! I charge thee, speak!
[Ghost exits.]

MARCELLUS  'Tis gone and will not answer.

BARNARDO
How now, Horatio, you tremble and look pale.
Is not this something more than fantasy?
What think you on 't?

HORATIO
Before my God, I might not this believe
Without the sensible and true avouch
Of mine own eyes.

MARCELLUS  Is it not like the King?

HORATIO  As thou art to thyself.
Such was the very armor he had on
When he the ambitious Norway combated.
So frowned he once when, in an angry parle,
He smote the sledded Polacks on the ice.
'Tis strange.

MARCELLUS
Thus twice before, and jump at this dead hour,
With martial stalk hath he gone by our watch.

HORATIO
In what particular thought to work I know not,
But in the gross and scope of mine opinion
This bodes some strange eruption to our state.

MARCELLUS
Good now, sit down, and tell me, he that knows,
Why this same strict and most observant watch
So nightly toils the subject of the land,
And why such daily cast of brazen cannon
And foreign mart for implements of war,
Why such impress of shipwrights, whose sore task
Does not divide the Sunday from the week.
What might be toward that this sweaty haste
Doth make the night joint laborer with the day?
Who is 't that can inform me?

HORATIO  That can I.
At least the whisper goes so: our last king,
Whose image even but now appeared to us,
Was, as you know, by Fortinbras of Norway,
Thereto pricked on by a most emulate pride,
Dared to the combat; in which our valiant Hamlet
(For so this side of our known world esteemed him)
Did slay this Fortinbras, who by a sealed compact,
Well ratified by law and heraldry,
Did forfeit, with his life, all those his lands
Which he stood seized of, to the conqueror.
Against the which a moiety competent
Was gaged by our king, which had returned
To the inheritance of Fortinbras
Had he been vanquisher, as, by the same comart
And carriage of the article designed,
His fell to Hamlet. Now, sir, young Fortinbras,
Of unimproved mettle hot and full,
Hath in the skirts of Norway here and there
Sharked up a list of lawless resolutes
For food and diet to some enterprise
That hath a stomach in 't; which is no other
(As it doth well appear unto our state)
But to recover of us, by strong hand
And terms compulsatory, those foresaid lands
So by his father lost. And this, I take it,
Is the main motive of our preparations,
The source of this our watch, and the chief head
Of this posthaste and rummage in the land.

BARNARDO
I think it be no other but e'en so.
Well may it sort that this portentous figure
Comes armed through our watch so like the king
That was and is the question of these wars.

HORATIO
A mote it is to trouble the mind's eye.
In the most high and palmy state of Rome,
A little ere the mightiest Julius fell,
The graves stood tenantless, and the sheeted dead
Did squeak and gibber in the Roman streets;
As stars with trains of fire and dews of blood,
Disasters in the sun; and the moist star,
Upon whose influence Neptune's empire stands,
Was sick almost to doomsday with eclipse.
And even the like precurse of feared events,
As harbingers preceding still the fates
And prologue to the omen coming on,
Have heaven and Earth together demonstrated
Unto our climatures and countrymen.

[Enter Ghost.]

But soft, behold! Lo, where it comes again!
I'll cross it though it blast me.--Stay, illusion!
[It spreads his arms.]
If thou hast any sound or use of voice,
Speak to me.
If there be any good thing to be done
That may to thee do ease and grace to me,
Speak to me.
If thou art privy to thy country's fate,
Which happily foreknowing may avoid,
O, speak!
Or if thou hast uphoarded in thy life
Extorted treasure in the womb of earth,
For which, they say, you spirits oft walk in death,
Speak of it.	[The cock crows.]
Stay and speak!--Stop it, Marcellus.

MARCELLUS
Shall I strike it with my partisan?

HORATIO  Do, if it will not stand.

BARNARDO  'Tis here.

HORATIO  'Tis here.
[Ghost exits.]

MARCELLUS  'Tis gone.
We do it wrong, being so majestical,
To offer it the show of violence,
For it is as the air, invulnerable,
And our vain blows malicious mockery.

BARNARDO
It was about to speak when the cock crew.

HORATIO
And then it started like a guilty thing
Upon a fearful summons. I have heard
The cock, that is the trumpet to the morn,
Doth with his lofty and shrill-sounding throat
Awake the god of day, and at his warning,
Whether in sea or fire, in earth or air,
Th' extravagant and erring spirit hies
To his confine, and of the truth herein
This present object made probation.

MARCELLUS
It faded on the crowing of the cock.
Some say that ever 'gainst that season comes
Wherein our Savior's birth is celebrated,
This bird of dawning singeth all night long;
And then, they say, no spirit dare stir abroad,
The nights are wholesome; then no planets strike,
No fairy takes, nor witch hath power to charm,
So hallowed and so gracious is that time.

HORATIO
So have I heard and do in part believe it.
But look, the morn in russet mantle clad
Walks o'er the dew of yon high eastward hill.
Break we our watch up, and by my advice
Let us impart what we have seen tonight
Unto young Hamlet; for, upon my life,
This spirit, dumb to us, will speak to him.
Do you consent we shall acquaint him with it
As needful in our loves, fitting our duty?

MARCELLUS
Let's do 't, I pray, and I this morning know
Where we shall find him most convenient.
[They exit.]

Scene 2
=======
[Flourish. Enter Claudius, King of Denmark, Gertrude the
Queen, the Council, as Polonius, and his son Laertes,
Hamlet, with others, among them Voltemand and
Cornelius.]


KING
Though yet of Hamlet our dear brother's death
The memory be green, and that it us befitted
To bear our hearts in grief, and our whole kingdom
To be contracted in one brow of woe,
Yet so far hath discretion fought with nature
That we with wisest sorrow think on him
Together with remembrance of ourselves.
Therefore our sometime sister, now our queen,
Th' imperial jointress to this warlike state,
Have we (as 'twere with a defeated joy,
With an auspicious and a dropping eye,
With mirth in funeral and with dirge in marriage,
In equal scale weighing delight and dole)
Taken to wife. Nor have we herein barred
Your better wisdoms, which have freely gone
With this affair along. For all, our thanks.
Now follows that you know. Young Fortinbras,
Holding a weak supposal of our worth
Or thinking by our late dear brother's death
Our state to be disjoint and out of frame,
Colleagued with this dream of his advantage,
He hath not failed to pester us with message
Importing the surrender of those lands
Lost by his father, with all bonds of law,
To our most valiant brother--so much for him.
Now for ourself and for this time of meeting.
Thus much the business is: we have here writ
To Norway, uncle of young Fortinbras,
Who, impotent and bedrid, scarcely hears
Of this his nephew's purpose, to suppress
His further gait herein, in that the levies,
The lists, and full proportions are all made
Out of his subject; and we here dispatch
You, good Cornelius, and you, Voltemand,
For bearers of this greeting to old Norway,
Giving to you no further personal power
To business with the King more than the scope
Of these dilated articles allow.
[Giving them a paper.]
Farewell, and let your haste commend your duty.

CORNELIUS/VOLTEMAND
In that and all things will we show our duty.

KING
We doubt it nothing. Heartily farewell.
[Voltemand and Cornelius exit.]
And now, Laertes, what's the news with you?
You told us of some suit. What is 't, Laertes?
You cannot speak of reason to the Dane
And lose your voice. What wouldst thou beg,
Laertes,
That shall not be my offer, not thy asking?
The head is not more native to the heart,
The hand more instrumental to the mouth,
Than is the throne of Denmark to thy father.
What wouldst thou have, Laertes?

LAERTES  My dread lord,
Your leave and favor to return to France,
From whence though willingly I came to Denmark
To show my duty in your coronation,
Yet now I must confess, that duty done,
My thoughts and wishes bend again toward France
And bow them to your gracious leave and pardon.

KING
Have you your father's leave? What says Polonius?

POLONIUS
Hath, my lord, wrung from me my slow leave
By laborsome petition, and at last
Upon his will I sealed my hard consent.
I do beseech you give him leave to go.

KING
Take thy fair hour, Laertes. Time be thine,
And thy best graces spend it at thy will.--
But now, my cousin Hamlet and my son--

HAMLET, [aside]
A little more than kin and less than kind.

KING
How is it that the clouds still hang on you?

HAMLET
Not so, my lord; I am too much in the sun.

QUEEN
Good Hamlet, cast thy nighted color off,
And let thine eye look like a friend on Denmark.
Do not forever with thy vailed lids
Seek for thy noble father in the dust.
Thou know'st 'tis common; all that lives must die,
Passing through nature to eternity.

HAMLET
Ay, madam, it is common.

QUEEN  If it be,
Why seems it so particular with thee?

HAMLET
"Seems," madam? Nay, it is. I know not "seems."
'Tis not alone my inky cloak, good mother,
Nor customary suits of solemn black,
Nor windy suspiration of forced breath,
No, nor the fruitful river in the eye,
Nor the dejected havior of the visage,
Together with all forms, moods, shapes of grief,
That can denote me truly. These indeed "seem,"
For they are actions that a man might play;
But I have that within which passes show,
These but the trappings and the suits of woe.

KING
'Tis sweet and commendable in your nature,
Hamlet,
To give these mourning duties to your father.
But you must know your father lost a father,
That father lost, lost his, and the survivor bound
In filial obligation for some term
To do obsequious sorrow. But to persever
In obstinate condolement is a course
Of impious stubbornness. 'Tis unmanly grief.
It shows a will most incorrect to heaven,
A heart unfortified, a mind impatient,
An understanding simple and unschooled.
For what we know must be and is as common
As any the most vulgar thing to sense,
Why should we in our peevish opposition
Take it to heart? Fie, 'tis a fault to heaven,
A fault against the dead, a fault to nature,
To reason most absurd, whose common theme
Is death of fathers, and who still hath cried,
From the first corse till he that died today,
"This must be so." We pray you, throw to earth
This unprevailing woe and think of us
As of a father; for let the world take note,
You are the most immediate to our throne,
And with no less nobility of love
Than that which dearest father bears his son
Do I impart toward you. For your intent
In going back to school in Wittenberg,
It is most retrograde to our desire,
And we beseech you, bend you to remain
Here in the cheer and comfort of our eye,
Our chiefest courtier, cousin, and our son.

QUEEN
Let not thy mother lose her prayers, Hamlet.
I pray thee, stay with us. Go not to Wittenberg.

HAMLET
I shall in all my best obey you, madam.

KING
Why, 'tis a loving and a fair reply.
Be as ourself in Denmark.--Madam, come.
This gentle and unforced accord of Hamlet
Sits smiling to my heart, in grace whereof
No jocund health that Denmark drinks today
But the great cannon to the clouds shall tell,
And the King's rouse the heaven shall bruit again,
Respeaking earthly thunder. Come away.
[Flourish. All but Hamlet exit.]

HAMLET
O, that this too, too sullied flesh would melt,
Thaw, and resolve itself into a dew,
Or that the Everlasting had not fixed
His canon 'gainst self-slaughter! O God, God,
How weary, stale, flat, and unprofitable
Seem to me all the uses of this world!
Fie on 't, ah fie! 'Tis an unweeded garden
That grows to seed. Things rank and gross in nature
Possess it merely. That it should come to this:
But two months dead--nay, not so much, not two.
So excellent a king, that was to this
Hyperion to a satyr; so loving to my mother
That he might not beteem the winds of heaven
Visit her face too roughly. Heaven and Earth,
Must I remember? Why, she would hang on him
As if increase of appetite had grown
By what it fed on. And yet, within a month
(Let me not think on 't; frailty, thy name is woman!),
A little month, or ere those shoes were old
With which she followed my poor father's body,
Like Niobe, all tears--why she, even she
(O God, a beast that wants discourse of reason
Would have mourned longer!), married with my
uncle,
My father's brother, but no more like my father
Than I to Hercules. Within a month,
Ere yet the salt of most unrighteous tears
Had left the flushing in her galled eyes,
She married. O, most wicked speed, to post
With such dexterity to incestuous sheets!
It is not, nor it cannot come to good.
But break, my heart, for I must hold my tongue.

[Enter Horatio, Marcellus, and Barnardo.]


HORATIO  Hail to your Lordship.

HAMLET  I am glad to see you well.
Horatio--or I do forget myself!

HORATIO
The same, my lord, and your poor servant ever.

HAMLET
Sir, my good friend. I'll change that name with you.
And what make you from Wittenberg, Horatio?--
Marcellus?

MARCELLUS  My good lord.

HAMLET
I am very glad to see you. [To Barnardo.] Good
even, sir.--
But what, in faith, make you from Wittenberg?

HORATIO
A truant disposition, good my lord.

HAMLET
I would not hear your enemy say so,
Nor shall you do my ear that violence
To make it truster of your own report
Against yourself. I know you are no truant.
But what is your affair in Elsinore?
We'll teach you to drink deep ere you depart.

HORATIO
My lord, I came to see your father's funeral.

HAMLET
I prithee, do not mock me, fellow student.
I think it was to see my mother's wedding.

HORATIO
Indeed, my lord, it followed hard upon.

HAMLET
Thrift, thrift, Horatio. The funeral baked meats
Did coldly furnish forth the marriage tables.
Would I had met my dearest foe in heaven
Or ever I had seen that day, Horatio!
My father--methinks I see my father.

HORATIO
Where, my lord?

HAMLET  In my mind's eye, Horatio.

HORATIO
I saw him once. He was a goodly king.

HAMLET
He was a man. Take him for all in all,
I shall not look upon his like again.

HORATIO
My lord, I think I saw him yesternight.

HAMLET  Saw who?

HORATIO
My lord, the King your father.

HAMLET  The King my father?

HORATIO
Season your admiration for a while
With an attent ear, till I may deliver
Upon the witness of these gentlemen
This marvel to you.

HAMLET  For God's love, let me hear!

HORATIO
Two nights together had these gentlemen,
Marcellus and Barnardo, on their watch,
In the dead waste and middle of the night,
Been thus encountered: a figure like your father,
Armed at point exactly, cap-a-pie,
Appears before them and with solemn march
Goes slow and stately by them. Thrice he walked
By their oppressed and fear-surprised eyes
Within his truncheon's length, whilst they, distilled
Almost to jelly with the act of fear,
Stand dumb and speak not to him. This to me
In dreadful secrecy impart they did,
And I with them the third night kept the watch,
Where, as they had delivered, both in time,
Form of the thing (each word made true and good),
The apparition comes. I knew your father;
These hands are not more like.

HAMLET  But where was this?

MARCELLUS
My lord, upon the platform where we watch.

HAMLET
Did you not speak to it?

HORATIO  My lord, I did,
But answer made it none. Yet once methought
It lifted up its head and did address
Itself to motion, like as it would speak;
But even then the morning cock crew loud,
And at the sound it shrunk in haste away
And vanished from our sight.

HAMLET  'Tis very strange.

HORATIO
As I do live, my honored lord, 'tis true.
And we did think it writ down in our duty
To let you know of it.

HAMLET  Indeed, sirs, but this troubles me.
Hold you the watch tonight?

ALL  We do, my lord.

HAMLET
Armed, say you?

ALL  Armed, my lord.

HAMLET  From top to toe?

ALL  My lord, from head to foot.

HAMLET  Then saw you not his face?

HORATIO
O, yes, my lord, he wore his beaver up.

HAMLET  What, looked he frowningly?

HORATIO
A countenance more in sorrow than in anger.

HAMLET  Pale or red?

HORATIO
Nay, very pale.

HAMLET  And fixed his eyes upon you?

HORATIO
Most constantly.

HAMLET  I would I had been there.

HORATIO  It would have much amazed you.

HAMLET  Very like. Stayed it long?

HORATIO
While one with moderate haste might tell a
hundred.

BARNARDO/MARCELLUS  Longer, longer.

HORATIO
Not when I saw 't.

HAMLET  His beard was grizzled, no?

HORATIO
It was as I have seen it in his life,
A sable silvered.

HAMLET  I will watch tonight.
Perchance 'twill walk again.

HORATIO  I warrant it will.

HAMLET
If it assume my noble father's person,
I'll speak to it, though hell itself should gape
And bid me hold my peace. I pray you all,
If you have hitherto concealed this sight,
Let it be tenable in your silence still;
And whatsomever else shall hap tonight,
Give it an understanding but no tongue.
I will requite your loves. So fare you well.
Upon the platform, 'twixt eleven and twelve,
I'll visit you.

ALL  Our duty to your Honor.

HAMLET
Your loves, as mine to you. Farewell.
[All but Hamlet exit.]
My father's spirit--in arms! All is not well.
I doubt some foul play. Would the night were come!
Till then, sit still, my soul. Foul deeds will rise,
Though all the earth o'erwhelm them, to men's
eyes.
[He exits.]

Scene 3
=======
[Enter Laertes and Ophelia, his sister.]


LAERTES
My necessaries are embarked. Farewell.
And, sister, as the winds give benefit
And convey is assistant, do not sleep,
But let me hear from you.

OPHELIA  Do you doubt that?

LAERTES
For Hamlet, and the trifling of his favor,
Hold it a fashion and a toy in blood,
A violet in the youth of primy nature,
Forward, not permanent, sweet, not lasting,
The perfume and suppliance of a minute,
No more.

OPHELIA  No more but so?

LAERTES  Think it no more.
For nature, crescent, does not grow alone
In thews and bulk, but, as this temple waxes,
The inward service of the mind and soul
Grows wide withal. Perhaps he loves you now,
And now no soil nor cautel doth besmirch
The virtue of his will; but you must fear,
His greatness weighed, his will is not his own,
For he himself is subject to his birth.
He may not, as unvalued persons do,
Carve for himself, for on his choice depends
The safety and the health of this whole state.
And therefore must his choice be circumscribed
Unto the voice and yielding of that body
Whereof he is the head. Then, if he says he loves
you,
It fits your wisdom so far to believe it
As he in his particular act and place
May give his saying deed, which is no further
Than the main voice of Denmark goes withal.
Then weigh what loss your honor may sustain
If with too credent ear you list his songs
Or lose your heart or your chaste treasure open
To his unmastered importunity.
Fear it, Ophelia; fear it, my dear sister,
And keep you in the rear of your affection,
Out of the shot and danger of desire.
The chariest maid is prodigal enough
If she unmask her beauty to the moon.
Virtue itself 'scapes not calumnious strokes.
The canker galls the infants of the spring
Too oft before their buttons be disclosed,
And, in the morn and liquid dew of youth,
Contagious blastments are most imminent.
Be wary, then; best safety lies in fear.
Youth to itself rebels, though none else near.

OPHELIA
I shall the effect of this good lesson keep
As watchman to my heart. But, good my brother,
Do not, as some ungracious pastors do,
Show me the steep and thorny way to heaven,
Whiles, like a puffed and reckless libertine,
Himself the primrose path of dalliance treads
And recks not his own rede.

LAERTES  O, fear me not.

[Enter Polonius.]

I stay too long. But here my father comes.
A double blessing is a double grace.
Occasion smiles upon a second leave.

POLONIUS
Yet here, Laertes? Aboard, aboard, for shame!
The wind sits in the shoulder of your sail,
And you are stayed for. There, my blessing with
thee.
And these few precepts in thy memory
Look thou character. Give thy thoughts no tongue,
Nor any unproportioned thought his act.
Be thou familiar, but by no means vulgar.
Those friends thou hast, and their adoption tried,
Grapple them unto thy soul with hoops of steel,
But do not dull thy palm with entertainment
Of each new-hatched, unfledged courage. Beware
Of entrance to a quarrel, but, being in,
Bear 't that th' opposed may beware of thee.
Give every man thy ear, but few thy voice.
Take each man's censure, but reserve thy judgment.
Costly thy habit as thy purse can buy,
But not expressed in fancy (rich, not gaudy),
For the apparel oft proclaims the man,
And they in France of the best rank and station
Are of a most select and generous chief in that.
Neither a borrower nor a lender be,
For loan oft loses both itself and friend,
And borrowing dulls the edge of husbandry.
This above all: to thine own self be true,
And it must follow, as the night the day,
Thou canst not then be false to any man.
Farewell. My blessing season this in thee.

LAERTES
Most humbly do I take my leave, my lord.

POLONIUS
The time invests you. Go, your servants tend.

LAERTES
Farewell, Ophelia, and remember well
What I have said to you.

OPHELIA  'Tis in my memory locked,
And you yourself shall keep the key of it.

LAERTES  Farewell.	[Laertes exits.]

POLONIUS
What is 't, Ophelia, he hath said to you?

OPHELIA
So please you, something touching the Lord
Hamlet.

POLONIUS  Marry, well bethought.
'Tis told me he hath very oft of late
Given private time to you, and you yourself
Have of your audience been most free and
bounteous.
If it be so (as so 'tis put on me,
And that in way of caution), I must tell you
You do not understand yourself so clearly
As it behooves my daughter and your honor.
What is between you? Give me up the truth.

OPHELIA
He hath, my lord, of late made many tenders
Of his affection to me.

POLONIUS
Affection, puh! You speak like a green girl
Unsifted in such perilous circumstance.
Do you believe his "tenders," as you call them?

OPHELIA
I do not know, my lord, what I should think.

POLONIUS
Marry, I will teach you. Think yourself a baby
That you have ta'en these tenders for true pay,
Which are not sterling. Tender yourself more dearly,
Or (not to crack the wind of the poor phrase,
Running it thus) you'll tender me a fool.

OPHELIA
My lord, he hath importuned me with love
In honorable fashion--

POLONIUS
Ay, "fashion" you may call it. Go to, go to!

OPHELIA
And hath given countenance to his speech, my lord,
With almost all the holy vows of heaven.

POLONIUS
Ay, springes to catch woodcocks. I do know,
When the blood burns, how prodigal the soul
Lends the tongue vows. These blazes, daughter,
Giving more light than heat, extinct in both
Even in their promise as it is a-making,
You must not take for fire. From this time
Be something scanter of your maiden presence.
Set your entreatments at a higher rate
Than a command to parle. For Lord Hamlet,
Believe so much in him that he is young,
And with a larger tether may he walk
Than may be given you. In few, Ophelia,
Do not believe his vows, for they are brokers,
Not of that dye which their investments show,
But mere implorators of unholy suits,
Breathing like sanctified and pious bawds
The better to beguile. This is for all:
I would not, in plain terms, from this time forth
Have you so slander any moment leisure
As to give words or talk with the Lord Hamlet.
Look to 't, I charge you. Come your ways.

OPHELIA  I shall obey, my lord.
[They exit.]

Scene 4
=======
[Enter Hamlet, Horatio, and Marcellus.]


HAMLET
The air bites shrewdly; it is very cold.

HORATIO
It is a nipping and an eager air.

HAMLET  What hour now?

HORATIO  I think it lacks of twelve.

MARCELLUS  No, it is struck.

HORATIO
Indeed, I heard it not. It then draws near the season
Wherein the spirit held his wont to walk.
[A flourish of trumpets and two pieces goes off.]
What does this mean, my lord?

HAMLET
The King doth wake tonight and takes his rouse,
Keeps wassail, and the swagg'ring upspring reels;
And, as he drains his draughts of Rhenish down,
The kettledrum and trumpet thus bray out
The triumph of his pledge.

HORATIO  Is it a custom?

HAMLET  Ay, marry, is 't,
But, to my mind, though I am native here
And to the manner born, it is a custom
More honored in the breach than the observance.
This heavy-headed revel east and west
Makes us traduced and taxed of other nations.
They clepe us drunkards and with swinish phrase
Soil our addition. And, indeed, it takes
From our achievements, though performed at
height,
The pith and marrow of our attribute.
So oft it chances in particular men
That for some vicious mole of nature in them,
As in their birth (wherein they are not guilty,
Since nature cannot choose his origin),
By the o'ergrowth of some complexion
(Oft breaking down the pales and forts of reason),
Or by some habit that too much o'erleavens
The form of plausive manners--that these men,
Carrying, I say, the stamp of one defect,
Being nature's livery or fortune's star,
His virtues else, be they as pure as grace,
As infinite as man may undergo,
Shall in the general censure take corruption
From that particular fault. The dram of evil
Doth all the noble substance of a doubt
To his own scandal.

[Enter Ghost.]


HORATIO  Look, my lord, it comes.

HAMLET
Angels and ministers of grace, defend us!
Be thou a spirit of health or goblin damned,
Bring with thee airs from heaven or blasts from
hell,
Be thy intents wicked or charitable,
Thou com'st in such a questionable shape
That I will speak to thee. I'll call thee "Hamlet,"
"King," "Father," "Royal Dane." O, answer me!
Let me not burst in ignorance, but tell
Why thy canonized bones, hearsed in death,
Have burst their cerements; why the sepulcher,
Wherein we saw thee quietly interred,
Hath oped his ponderous and marble jaws
To cast thee up again. What may this mean
That thou, dead corse, again in complete steel,
Revisits thus the glimpses of the moon,
Making night hideous, and we fools of nature
So horridly to shake our disposition
With thoughts beyond the reaches of our souls?
Say, why is this? Wherefore? What should we do?
[Ghost beckons.]

HORATIO
It beckons you to go away with it
As if it some impartment did desire
To you alone.

MARCELLUS  Look with what courteous action
It waves you to a more removed ground.
But do not go with it.

HORATIO  No, by no means.

HAMLET
It will not speak. Then I will follow it.

HORATIO
Do not, my lord.

HAMLET  Why, what should be the fear?
I do not set my life at a pin's fee.
And for my soul, what can it do to that,
Being a thing immortal as itself?
It waves me forth again. I'll follow it.

HORATIO
What if it tempt you toward the flood, my lord?
Or to the dreadful summit of the cliff
That beetles o'er his base into the sea,
And there assume some other horrible form
Which might deprive your sovereignty of reason
And draw you into madness? Think of it.
The very place puts toys of desperation,
Without more motive, into every brain
That looks so many fathoms to the sea
And hears it roar beneath.

HAMLET
It waves me still.--Go on, I'll follow thee.

MARCELLUS
You shall not go, my lord.	[They hold back Hamlet.]

HAMLET  Hold off your hands.

HORATIO
Be ruled. You shall not go.

HAMLET  My fate cries out
And makes each petty arture in this body
As hardy as the Nemean lion's nerve.
Still am I called. Unhand me, gentlemen.
By heaven, I'll make a ghost of him that lets me!
I say, away!--Go on. I'll follow thee.
[Ghost and Hamlet exit.]

HORATIO
He waxes desperate with imagination.

MARCELLUS
Let's follow. 'Tis not fit thus to obey him.

HORATIO
Have after. To what issue will this come?

MARCELLUS
Something is rotten in the state of Denmark.

HORATIO
Heaven will direct it.

MARCELLUS  Nay, let's follow him.
[They exit.]

Scene 5
=======
[Enter Ghost and Hamlet.]


HAMLET
Whither wilt thou lead me? Speak. I'll go no
further.

GHOST
Mark me.

HAMLET  I will.

GHOST  My hour is almost come
When I to sulf'rous and tormenting flames
Must render up myself.

HAMLET  Alas, poor ghost!

GHOST
Pity me not, but lend thy serious hearing
To what I shall unfold.

HAMLET  Speak. I am bound to hear.

GHOST
So art thou to revenge, when thou shalt hear.

HAMLET  What?

GHOST  I am thy father's spirit,
Doomed for a certain term to walk the night
And for the day confined to fast in fires
Till the foul crimes done in my days of nature
Are burnt and purged away. But that I am forbid
To tell the secrets of my prison house,
I could a tale unfold whose lightest word
Would harrow up thy soul, freeze thy young blood,
Make thy two eyes, like stars, start from their
spheres,
Thy knotted and combined locks to part,
And each particular hair to stand an end,
Like quills upon the fearful porpentine.
But this eternal blazon must not be
To ears of flesh and blood. List, list, O list!
If thou didst ever thy dear father love--

HAMLET  O God!

GHOST
Revenge his foul and most unnatural murder.

HAMLET  Murder?

GHOST
Murder most foul, as in the best it is,
But this most foul, strange, and unnatural.

HAMLET
Haste me to know 't, that I, with wings as swift
As meditation or the thoughts of love,
May sweep to my revenge.

GHOST  I find thee apt;
And duller shouldst thou be than the fat weed
That roots itself in ease on Lethe wharf,
Wouldst thou not stir in this. Now, Hamlet, hear.
'Tis given out that, sleeping in my orchard,
A serpent stung me. So the whole ear of Denmark
Is by a forged process of my death
Rankly abused. But know, thou noble youth,
The serpent that did sting thy father's life
Now wears his crown.

HAMLET  O, my prophetic soul! My uncle!

GHOST
Ay, that incestuous, that adulterate beast,
With witchcraft of his wits, with traitorous gifts--
O wicked wit and gifts, that have the power
So to seduce!--won to his shameful lust
The will of my most seeming-virtuous queen.
O Hamlet, what a falling off was there!
From me, whose love was of that dignity
That it went hand in hand even with the vow
I made to her in marriage, and to decline
Upon a wretch whose natural gifts were poor
To those of mine.
But virtue, as it never will be moved,
Though lewdness court it in a shape of heaven,
So, lust, though to a radiant angel linked,
Will sate itself in a celestial bed
And prey on garbage.
But soft, methinks I scent the morning air.
Brief let me be. Sleeping within my orchard,
My custom always of the afternoon,
Upon my secure hour thy uncle stole
With juice of cursed hebona in a vial
And in the porches of my ears did pour
The leprous distilment, whose effect
Holds such an enmity with blood of man
That swift as quicksilver it courses through
The natural gates and alleys of the body,
And with a sudden vigor it doth posset
And curd, like eager droppings into milk,
The thin and wholesome blood. So did it mine,
And a most instant tetter barked about,
Most lazar-like, with vile and loathsome crust
All my smooth body.
Thus was I, sleeping, by a brother's hand
Of life, of crown, of queen at once dispatched,
Cut off, even in the blossoms of my sin,
Unhouseled, disappointed, unaneled,
No reck'ning made, but sent to my account
With all my imperfections on my head.
O horrible, O horrible, most horrible!
If thou hast nature in thee, bear it not.
Let not the royal bed of Denmark be
A couch for luxury and damned incest.
But, howsomever thou pursues this act,
Taint not thy mind, nor let thy soul contrive
Against thy mother aught. Leave her to heaven
And to those thorns that in her bosom lodge
To prick and sting her. Fare thee well at once.
The glowworm shows the matin to be near
And 'gins to pale his uneffectual fire.
Adieu, adieu, adieu. Remember me.	[He exits.]

HAMLET
O all you host of heaven! O Earth! What else?
And shall I couple hell? O fie! Hold, hold, my heart,
And you, my sinews, grow not instant old,
But bear me stiffly up. Remember thee?
Ay, thou poor ghost, whiles memory holds a seat
In this distracted globe. Remember thee?
Yea, from the table of my memory
I'll wipe away all trivial, fond records,
All saws of books, all forms, all pressures past,
That youth and observation copied there,
And thy commandment all alone shall live
Within the book and volume of my brain,
Unmixed with baser matter. Yes, by heaven!
O most pernicious woman!
O villain, villain, smiling, damned villain!
My tables--meet it is I set it down
That one may smile and smile and be a villain.
At least I am sure it may be so in Denmark.
[He writes.]
So, uncle, there you are. Now to my word.
It is "adieu, adieu, remember me."
I have sworn 't.

[Enter Horatio and Marcellus.]


HORATIO  My lord, my lord!

MARCELLUS  Lord Hamlet.

HORATIO  Heavens secure him!

HAMLET  So be it.

MARCELLUS  Illo, ho, ho, my lord!

HAMLET  Hillo, ho, ho, boy! Come, bird, come!

MARCELLUS
How is 't, my noble lord?

HORATIO  What news, my lord?

HAMLET  O, wonderful!

HORATIO
Good my lord, tell it.

HAMLET  No, you will reveal it.

HORATIO
Not I, my lord, by heaven.

MARCELLUS  Nor I, my lord.

HAMLET
How say you, then? Would heart of man once think
it?
But you'll be secret?

HORATIO/MARCELLUS   Ay, by heaven, my lord.

HAMLET
There's never a villain dwelling in all Denmark
But he's an arrant knave.

HORATIO
There needs no ghost, my lord, come from the grave
To tell us this.

HAMLET  Why, right, you are in the right.
And so, without more circumstance at all,
I hold it fit that we shake hands and part,
You, as your business and desire shall point you
(For every man hath business and desire,
Such as it is), and for my own poor part,
I will go pray.

HORATIO
These are but wild and whirling words, my lord.

HAMLET
I am sorry they offend you, heartily;
Yes, faith, heartily.

HORATIO  There's no offense, my lord.

HAMLET
Yes, by Saint Patrick, but there is, Horatio,
And much offense, too. Touching this vision here,
It is an honest ghost--that let me tell you.
For your desire to know what is between us,
O'ermaster 't as you may. And now, good friends,
As you are friends, scholars, and soldiers,
Give me one poor request.

HORATIO  What is 't, my lord? We will.

HAMLET
Never make known what you have seen tonight.

HORATIO/MARCELLUS   My lord, we will not.

HAMLET  Nay, but swear 't.

HORATIO  In faith, my lord, not I.

MARCELLUS  Nor I, my lord, in faith.

HAMLET
Upon my sword.

MARCELLUS  We have sworn, my lord, already.

HAMLET  Indeed, upon my sword, indeed.

GHOST [cries under the stage]  Swear.

HAMLET
Ha, ha, boy, sayst thou so? Art thou there,
truepenny?
Come on, you hear this fellow in the cellarage.
Consent to swear.

HORATIO  Propose the oath, my lord.

HAMLET
Never to speak of this that you have seen,
Swear by my sword.

GHOST, [beneath]  Swear.

HAMLET
Hic et ubique? Then we'll shift our ground.
Come hither, gentlemen,
And lay your hands again upon my sword.
Swear by my sword
Never to speak of this that you have heard.

GHOST, [beneath]  Swear by his sword.

HAMLET
Well said, old mole. Canst work i' th' earth so fast?--
A worthy pioner! Once more remove, good friends.

HORATIO
O day and night, but this is wondrous strange.

HAMLET
And therefore as a stranger give it welcome.
There are more things in heaven and earth, Horatio,
Than are dreamt of in your philosophy. But come.
Here, as before, never, so help you mercy,
How strange or odd some'er I bear myself
(As I perchance hereafter shall think meet
To put an antic disposition on)
That you, at such times seeing me, never shall,
With arms encumbered thus, or this headshake,
Or by pronouncing of some doubtful phrase,
As "Well, well, we know," or "We could an if we
would,"
Or "If we list to speak," or "There be an if they
might,"
Or such ambiguous giving-out, to note
That you know aught of me--this do swear,
So grace and mercy at your most need help you.

GHOST, [beneath]  Swear.

HAMLET
Rest, rest, perturbed spirit.--So, gentlemen,
With all my love I do commend me to you,
And what so poor a man as Hamlet is
May do t' express his love and friending to you,
God willing, shall not lack. Let us go in together,
And still your fingers on your lips, I pray.
The time is out of joint. O cursed spite
That ever I was born to set it right!
Nay, come, let's go together.
[They exit.]


ACT 2
=====

Scene 1
=======
[Enter old Polonius with his man Reynaldo.]


POLONIUS
Give him this money and these notes, Reynaldo.

REYNALDO  I will, my lord.

POLONIUS
You shall do marvelous wisely, good Reynaldo,
Before you visit him, to make inquire
Of his behavior.

REYNALDO  My lord, I did intend it.

POLONIUS
Marry, well said, very well said. Look you, sir,
Inquire me first what Danskers are in Paris;
And how, and who, what means, and where they
keep,
What company, at what expense; and finding
By this encompassment and drift of question
That they do know my son, come you more nearer
Than your particular demands will touch it.
Take you, as 'twere, some distant knowledge of him,
As thus: "I know his father and his friends
And, in part, him." Do you mark this, Reynaldo?

REYNALDO  Ay, very well, my lord.

POLONIUS
"And, in part, him, but," you may say, "not well.
But if 't be he I mean, he's very wild,
Addicted so and so." And there put on him
What forgeries you please--marry, none so rank
As may dishonor him, take heed of that,
But, sir, such wanton, wild, and usual slips
As are companions noted and most known
To youth and liberty.

REYNALDO  As gaming, my lord.

POLONIUS  Ay, or drinking, fencing, swearing,
Quarreling, drabbing--you may go so far.

REYNALDO  My lord, that would dishonor him.

POLONIUS
Faith, no, as you may season it in the charge.
You must not put another scandal on him
That he is open to incontinency;
That's not my meaning. But breathe his faults so
quaintly
That they may seem the taints of liberty,
The flash and outbreak of a fiery mind,
A savageness in unreclaimed blood,
Of general assault.

REYNALDO  But, my good lord--

POLONIUS  Wherefore should you do this?

REYNALDO  Ay, my lord, I would know that.

POLONIUS  Marry, sir, here's my drift,
And I believe it is a fetch of wit.
You, laying these slight sullies on my son,
As 'twere a thing a little soiled i' th' working,
Mark you, your party in converse, him you would
sound,
Having ever seen in the prenominate crimes
The youth you breathe of guilty, be assured
He closes with you in this consequence:
"Good sir," or so, or "friend," or "gentleman,"
According to the phrase or the addition
Of man and country--

REYNALDO  Very good, my lord.

POLONIUS  And then, sir, does he this, he does--what
was I about to say? By the Mass, I was about to say
something. Where did I leave?

REYNALDO  At "closes in the consequence," at "friend,
or so," and "gentleman."

POLONIUS
At "closes in the consequence"--ay, marry--
He closes thus: "I know the gentleman.
I saw him yesterday," or "th' other day"
(Or then, or then, with such or such), "and as you
say,
There was he gaming, there o'ertook in 's rouse,
There falling out at tennis"; or perchance
"I saw him enter such a house of sale"--
Videlicet, a brothel--or so forth. See you now
Your bait of falsehood take this carp of truth;
And thus do we of wisdom and of reach,
With windlasses and with assays of bias,
By indirections find directions out.
So by my former lecture and advice
Shall you my son. You have me, have you not?

REYNALDO
My lord, I have.

POLONIUS  God be wi' you. Fare you well.

REYNALDO  Good my lord.

POLONIUS
Observe his inclination in yourself.

REYNALDO  I shall, my lord.

POLONIUS  And let him ply his music.

REYNALDO  Well, my lord.

POLONIUS
Farewell.	[Reynaldo exits.]

[Enter Ophelia.]

How now, Ophelia, what's the matter?

OPHELIA
O, my lord, my lord, I have been so affrighted!

POLONIUS  With what, i' th' name of God?

OPHELIA
My lord, as I was sewing in my closet,
Lord Hamlet, with his doublet all unbraced,
No hat upon his head, his stockings fouled,
Ungartered, and down-gyved to his ankle,
Pale as his shirt, his knees knocking each other,
And with a look so piteous in purport
As if he had been loosed out of hell
To speak of horrors--he comes before me.

POLONIUS
Mad for thy love?

OPHELIA  My lord, I do not know,
But truly I do fear it.

POLONIUS  What said he?

OPHELIA
He took me by the wrist and held me hard.
Then goes he to the length of all his arm,
And, with his other hand thus o'er his brow,
He falls to such perusal of my face
As he would draw it. Long stayed he so.
At last, a little shaking of mine arm,
And thrice his head thus waving up and down,
He raised a sigh so piteous and profound
As it did seem to shatter all his bulk
And end his being. That done, he lets me go,
And, with his head over his shoulder turned,
He seemed to find his way without his eyes,
For out o' doors he went without their helps
And to the last bended their light on me.

POLONIUS
Come, go with me. I will go seek the King.
This is the very ecstasy of love,
Whose violent property fordoes itself
And leads the will to desperate undertakings
As oft as any passions under heaven
That does afflict our natures. I am sorry.
What, have you given him any hard words of late?

OPHELIA
No, my good lord, but as you did command
I did repel his letters and denied
His access to me.

POLONIUS  That hath made him mad.
I am sorry that with better heed and judgment
I had not coted him. I feared he did but trifle
And meant to wrack thee. But beshrew my jealousy!
By heaven, it is as proper to our age
To cast beyond ourselves in our opinions
As it is common for the younger sort
To lack discretion. Come, go we to the King.
This must be known, which, being kept close, might
move
More grief to hide than hate to utter love.
Come.
[They exit.]

Scene 2
=======
[Flourish. Enter King and Queen, Rosencrantz and
Guildenstern and Attendants.]


KING
Welcome, dear Rosencrantz and Guildenstern.
Moreover that we much did long to see you,
The need we have to use you did provoke
Our hasty sending. Something have you heard
Of Hamlet's transformation, so call it,
Sith nor th' exterior nor the inward man
Resembles that it was. What it should be,
More than his father's death, that thus hath put him
So much from th' understanding of himself
I cannot dream of. I entreat you both
That, being of so young days brought up with him
And sith so neighbored to his youth and havior,
That you vouchsafe your rest here in our court
Some little time, so by your companies
To draw him on to pleasures, and to gather
So much as from occasion you may glean,
Whether aught to us unknown afflicts him thus
That, opened, lies within our remedy.

QUEEN
Good gentlemen, he hath much talked of you,
And sure I am two men there is not living
To whom he more adheres. If it will please you
To show us so much gentry and goodwill
As to expend your time with us awhile
For the supply and profit of our hope,
Your visitation shall receive such thanks
As fits a king's remembrance.

ROSENCRANTZ  Both your Majesties
Might, by the sovereign power you have of us,
Put your dread pleasures more into command
Than to entreaty.

GUILDENSTERN  But we both obey,
And here give up ourselves in the full bent
To lay our service freely at your feet,
To be commanded.

KING
Thanks, Rosencrantz and gentle Guildenstern.

QUEEN
Thanks, Guildenstern and gentle Rosencrantz.
And I beseech you instantly to visit
My too much changed son.--Go, some of you,
And bring these gentlemen where Hamlet is.

GUILDENSTERN
Heavens make our presence and our practices
Pleasant and helpful to him!

QUEEN  Ay, amen!
[Rosencrantz and Guildenstern exit
with some Attendants.]

[Enter Polonius.]


POLONIUS
Th' ambassadors from Norway, my good lord,
Are joyfully returned.

KING
Thou still hast been the father of good news.

POLONIUS
Have I, my lord? I assure my good liege
I hold my duty as I hold my soul,
Both to my God and to my gracious king,
And I do think, or else this brain of mine
Hunts not the trail of policy so sure
As it hath used to do, that I have found
The very cause of Hamlet's lunacy.

KING
O, speak of that! That do I long to hear.

POLONIUS
Give first admittance to th' ambassadors.
My news shall be the fruit to that great feast.

KING
Thyself do grace to them and bring them in.
[Polonius exits.]
He tells me, my dear Gertrude, he hath found
The head and source of all your son's distemper.

QUEEN
I doubt it is no other but the main--
His father's death and our o'erhasty marriage.

KING
Well, we shall sift him.

[Enter Ambassadors Voltemand and Cornelius with
Polonius.]

Welcome, my good friends.
Say, Voltemand, what from our brother Norway?

VOLTEMAND
Most fair return of greetings and desires.
Upon our first, he sent out to suppress
His nephew's levies, which to him appeared
To be a preparation 'gainst the Polack,
But, better looked into, he truly found
It was against your Highness. Whereat, grieved
That so his sickness, age, and impotence
Was falsely borne in hand, sends out arrests
On Fortinbras, which he, in brief, obeys,
Receives rebuke from Norway, and, in fine,
Makes vow before his uncle never more
To give th' assay of arms against your Majesty.
Whereon old Norway, overcome with joy,
Gives him three-score thousand crowns in annual
fee
And his commission to employ those soldiers,
So levied as before, against the Polack,
With an entreaty, herein further shown,
[He gives a paper.]
That it might please you to give quiet pass
Through your dominions for this enterprise,
On such regards of safety and allowance
As therein are set down.

KING  It likes us well,
And, at our more considered time, we'll read,
Answer, and think upon this business.
Meantime, we thank you for your well-took labor.
Go to your rest. At night we'll feast together.
Most welcome home!
[Voltemand and Cornelius exit.]

POLONIUS  This business is well ended.
My liege, and madam, to expostulate
What majesty should be, what duty is,
Why day is day, night night, and time is time
Were nothing but to waste night, day, and time.
Therefore, since brevity is the soul of wit,
And tediousness the limbs and outward flourishes,
I will be brief. Your noble son is mad.
"Mad" call I it, for, to define true madness,
What is 't but to be nothing else but mad?
But let that go.

QUEEN  More matter with less art.

POLONIUS
Madam, I swear I use no art at all.
That he's mad, 'tis true; 'tis true 'tis pity,
And pity 'tis 'tis true--a foolish figure,
But farewell it, for I will use no art.
Mad let us grant him then, and now remains
That we find out the cause of this effect,
Or, rather say, the cause of this defect,
For this effect defective comes by cause.
Thus it remains, and the remainder thus.
Perpend.
I have a daughter (have while she is mine)
Who, in her duty and obedience, mark,
Hath given me this. Now gather and surmise.
[He reads.] To the celestial, and my soul's idol, the
most beautified Ophelia--
That's an ill phrase, a vile phrase; "beautified" is a
vile phrase. But you shall hear. Thus: [He reads.]
In her excellent white bosom, these, etc.--

QUEEN  Came this from Hamlet to her?

POLONIUS
Good madam, stay awhile. I will be faithful.
[He reads the letter.]
	Doubt thou the stars are fire,
	   Doubt that the sun doth move,
	Doubt truth to be a liar,
	   But never doubt I love.
O dear Ophelia, I am ill at these numbers. I have not
art to reckon my groans, but that I love thee best, O
most best, believe it. Adieu.
Thine evermore, most dear lady, whilst
this machine is to him, Hamlet.
This, in obedience, hath my daughter shown me,
And more above, hath his solicitings,
As they fell out by time, by means, and place,
All given to mine ear.

KING  But how hath she received his love?

POLONIUS  What do you think of me?

KING
As of a man faithful and honorable.

POLONIUS
I would fain prove so. But what might you think,
When I had seen this hot love on the wing
(As I perceived it, I must tell you that,
Before my daughter told me), what might you,
Or my dear Majesty your queen here, think,
If I had played the desk or table-book
Or given my heart a winking, mute and dumb,
Or looked upon this love with idle sight?
What might you think? No, I went round to work,
And my young mistress thus I did bespeak:
"Lord Hamlet is a prince, out of thy star.
This must not be." And then I prescripts gave her,
That she should lock herself from his resort,
Admit no messengers, receive no tokens;
Which done, she took the fruits of my advice,
And he, repelled (a short tale to make),
Fell into a sadness, then into a fast,
Thence to a watch, thence into a weakness,
Thence to a lightness, and, by this declension,
Into the madness wherein now he raves
And all we mourn for.

KING, [to Queen]  Do you think 'tis this?

QUEEN  It may be, very like.

POLONIUS
Hath there been such a time (I would fain know
that)
That I have positively said "'Tis so,"
When it proved otherwise?

KING  Not that I know.

POLONIUS
Take this from this, if this be otherwise.
If circumstances lead me, I will find
Where truth is hid, though it were hid, indeed,
Within the center.

KING  How may we try it further?

POLONIUS
You know sometimes he walks four hours together
Here in the lobby.

QUEEN  So he does indeed.

POLONIUS
At such a time I'll loose my daughter to him.
[To the King.] Be you and I behind an arras then.
Mark the encounter. If he love her not,
And be not from his reason fall'n thereon,
Let me be no assistant for a state,
But keep a farm and carters.

KING  We will try it.

[Enter Hamlet reading on a book.]


QUEEN
But look where sadly the poor wretch comes
reading.

POLONIUS
Away, I do beseech you both, away.
I'll board him presently. O, give me leave.
[King and Queen exit with Attendants.]
How does my good Lord Hamlet?

HAMLET  Well, God-a-mercy.

POLONIUS  Do you know me, my lord?

HAMLET  Excellent well. You are a fishmonger.

POLONIUS  Not I, my lord.

HAMLET  Then I would you were so honest a man.

POLONIUS  Honest, my lord?

HAMLET  Ay, sir. To be honest, as this world goes, is to
be one man picked out of ten thousand.

POLONIUS  That's very true, my lord.

HAMLET  For if the sun breed maggots in a dead
dog, being a good kissing carrion--Have you a
daughter?

POLONIUS  I have, my lord.

HAMLET  Let her not walk i' th' sun. Conception is a
blessing, but, as your daughter may conceive,
friend, look to 't.

POLONIUS, [aside]  How say you by that? Still harping on
my daughter. Yet he knew me not at first; he said I
was a fishmonger. He is far gone. And truly, in my
youth, I suffered much extremity for love, very near
this. I'll speak to him again.--What do you read, my
lord?

HAMLET  Words, words, words.

POLONIUS  What is the matter, my lord?

HAMLET  Between who?

POLONIUS  I mean the matter that you read, my lord.

HAMLET  Slanders, sir; for the satirical rogue says here
that old men have gray beards, that their faces are
wrinkled, their eyes purging thick amber and
plum-tree gum, and that they have a plentiful lack of
wit, together with most weak hams; all which, sir,
though I most powerfully and potently believe, yet I
hold it not honesty to have it thus set down; for
yourself, sir, shall grow old as I am, if, like a crab,
you could go backward.

POLONIUS, [aside]  Though this be madness, yet there is
method in 't.--Will you walk out of the air, my lord?

HAMLET  Into my grave?

POLONIUS  Indeed, that's out of the air. [Aside.] How
pregnant sometimes his replies are! A happiness
that often madness hits on, which reason and
sanity could not so prosperously be delivered of. I
will leave him and suddenly contrive the means of
meeting between him and my daughter.--My lord,
I will take my leave of you.

HAMLET  You cannot, sir, take from me anything that I
will more willingly part withal--except my life,
except my life, except my life.

POLONIUS  Fare you well, my lord.

HAMLET, [aside]  These tedious old fools.

[Enter Guildenstern and Rosencrantz.]


POLONIUS  You go to seek the Lord Hamlet. There he is.

ROSENCRANTZ, [to Polonius]  God save you, sir.
[Polonius exits.]

GUILDENSTERN  My honored lord.

ROSENCRANTZ  My most dear lord.

HAMLET  My excellent good friends! How dost thou,
Guildenstern? Ah, Rosencrantz! Good lads, how do
you both?

ROSENCRANTZ
As the indifferent children of the earth.

GUILDENSTERN
Happy in that we are not overhappy.
On Fortune's cap, we are not the very button.

HAMLET  Nor the soles of her shoe?

ROSENCRANTZ  Neither, my lord.

HAMLET  Then you live about her waist, or in the
middle of her favors?

GUILDENSTERN  Faith, her privates we.

HAMLET  In the secret parts of Fortune? O, most true!
She is a strumpet. What news?

ROSENCRANTZ  None, my lord, but that the world's
grown honest.

HAMLET  Then is doomsday near. But your news is not
true. Let me question more in particular. What
have you, my good friends, deserved at the hands of
Fortune that she sends you to prison hither?

GUILDENSTERN  Prison, my lord?

HAMLET  Denmark's a prison.

ROSENCRANTZ  Then is the world one.

HAMLET  A goodly one, in which there are many confines,
wards, and dungeons, Denmark being one o'
th' worst.

ROSENCRANTZ  We think not so, my lord.

HAMLET  Why, then, 'tis none to you, for there is
nothing either good or bad but thinking makes it
so. To me, it is a prison.

ROSENCRANTZ  Why, then, your ambition makes it one.
'Tis too narrow for your mind.

HAMLET  O God, I could be bounded in a nutshell and
count myself a king of infinite space, were it not
that I have bad dreams.

GUILDENSTERN  Which dreams, indeed, are ambition,
for the very substance of the ambitious is merely
the shadow of a dream.

HAMLET  A dream itself is but a shadow.

ROSENCRANTZ  Truly, and I hold ambition of so airy
and light a quality that it is but a shadow's shadow.

HAMLET  Then are our beggars bodies, and our monarchs
and outstretched heroes the beggars' shadows.
Shall we to th' court? For, by my fay, I cannot
reason.

ROSENCRANTZ/GUILDENSTERN  We'll wait upon you.

HAMLET  No such matter. I will not sort you with the
rest of my servants, for, to speak to you like an
honest man, I am most dreadfully attended. But,
in the beaten way of friendship, what make you at
Elsinore?

ROSENCRANTZ  To visit you, my lord, no other occasion.

HAMLET  Beggar that I am, I am even poor in thanks;
but I thank you, and sure, dear friends, my thanks
are too dear a halfpenny. Were you not sent for?
Is it your own inclining? Is it a free visitation?
Come, come, deal justly with me. Come, come; nay,
speak.

GUILDENSTERN  What should we say, my lord?

HAMLET  Anything but to th' purpose. You were sent
for, and there is a kind of confession in your looks
which your modesties have not craft enough to
color. I know the good king and queen have sent for
you.

ROSENCRANTZ  To what end, my lord?

HAMLET  That you must teach me. But let me conjure
you by the rights of our fellowship, by the consonancy
of our youth, by the obligation of our ever-preserved
love, and by what more dear a better
proposer can charge you withal: be even and direct
with me whether you were sent for or no.

ROSENCRANTZ, [to Guildenstern]  What say you?

HAMLET, [aside]  Nay, then, I have an eye of you.--If
you love me, hold not off.

GUILDENSTERN  My lord, we were sent for.

HAMLET  I will tell you why; so shall my anticipation
prevent your discovery, and your secrecy to the
King and Queen molt no feather. I have of late, but
wherefore I know not, lost all my mirth, forgone all
custom of exercises, and, indeed, it goes so heavily
with my disposition that this goodly frame, the
Earth, seems to me a sterile promontory; this most
excellent canopy, the air, look you, this brave o'erhanging
firmament, this majestical roof, fretted
with golden fire--why, it appeareth nothing to me
but a foul and pestilent congregation of vapors.
What a piece of work is a man, how noble in
reason, how infinite in faculties, in form and moving
how express and admirable; in action how like
an angel, in apprehension how like a god: the
beauty of the world, the paragon of animals--and
yet, to me, what is this quintessence of dust? Man
delights not me, no, nor women neither, though by
your smiling you seem to say so.

ROSENCRANTZ  My lord, there was no such stuff in my
thoughts.

HAMLET  Why did you laugh, then, when I said "man
delights not me"?

ROSENCRANTZ  To think, my lord, if you delight not in
man, what Lenten entertainment the players shall
receive from you. We coted them on the way, and
hither are they coming to offer you service.

HAMLET  He that plays the king shall be welcome--his
Majesty shall have tribute on me. The adventurous
knight shall use his foil and target, the lover shall
not sigh gratis, the humorous man shall end his
part in peace, the clown shall make those laugh
whose lungs are tickle o' th' sear, and the lady
shall say her mind freely, or the blank verse shall
halt for 't. What players are they?

ROSENCRANTZ  Even those you were wont to take such
delight in, the tragedians of the city.

HAMLET  How chances it they travel? Their residence,
both in reputation and profit, was better both ways.

ROSENCRANTZ  I think their inhibition comes by the
means of the late innovation.

HAMLET  Do they hold the same estimation they did
when I was in the city? Are they so followed?

ROSENCRANTZ  No, indeed are they not.

HAMLET  How comes it? Do they grow rusty?

ROSENCRANTZ  Nay, their endeavor keeps in the wonted
pace. But there is, sir, an aerie of children, little
eyases, that cry out on the top of question and are
most tyrannically clapped for 't. These are now the
fashion and so berattle the common stages (so
they call them) that many wearing rapiers are afraid
of goose quills and dare scarce come thither.

HAMLET  What, are they children? Who maintains 'em?
How are they escoted? Will they pursue the quality
no longer than they can sing? Will they not say
afterwards, if they should grow themselves to common
players (as it is most like, if their means are
no better), their writers do them wrong to make
them exclaim against their own succession?

ROSENCRANTZ  Faith, there has been much to-do on
both sides, and the nation holds it no sin to tar
them to controversy. There was for a while no
money bid for argument unless the poet and the
player went to cuffs in the question.

HAMLET  Is 't possible?

GUILDENSTERN  O, there has been much throwing
about of brains.

HAMLET  Do the boys carry it away?

ROSENCRANTZ  Ay, that they do, my lord--Hercules
and his load too.

HAMLET  It is not very strange; for my uncle is King of
Denmark, and those that would make mouths at
him while my father lived give twenty, forty, fifty,
a hundred ducats apiece for his picture in little.
'Sblood, there is something in this more than natural,
if philosophy could find it out.
[A flourish for the Players.]

GUILDENSTERN  There are the players.

HAMLET  Gentlemen, you are welcome to Elsinore.
Your hands, come then. Th' appurtenance of welcome
is fashion and ceremony. Let me comply
with you in this garb, lest my extent to the players,
which, I tell you, must show fairly outwards, should
more appear like entertainment than yours. You are
welcome. But my uncle-father and aunt-mother are
deceived.

GUILDENSTERN  In what, my dear lord?

HAMLET  I am but mad north-north-west. When the
wind is southerly, I know a hawk from a handsaw.

[Enter Polonius.]


POLONIUS  Well be with you, gentlemen.

HAMLET  Hark you, Guildenstern, and you too--at
each ear a hearer! That great baby you see there is
not yet out of his swaddling clouts.

ROSENCRANTZ  Haply he is the second time come to
them, for they say an old man is twice a child.

HAMLET  I will prophesy he comes to tell me of the
players; mark it.--You say right, sir, a Monday
morning, 'twas then indeed.

POLONIUS  My lord, I have news to tell you.

HAMLET  My lord, I have news to tell you: when Roscius
was an actor in Rome--

POLONIUS  The actors are come hither, my lord.

HAMLET  Buzz, buzz.

POLONIUS  Upon my honor--

HAMLET  Then came each actor on his ass.

POLONIUS  The best actors in the world, either for
tragedy, comedy, history, pastoral, pastoral-comical,
historical-pastoral, tragical-historical,
tragical-comical-historical-pastoral, scene individable, or
poem unlimited. Seneca cannot be too heavy, nor
Plautus too light. For the law of writ and the liberty,
these are the only men.

HAMLET  O Jephthah, judge of Israel, what a treasure
hadst thou!

POLONIUS  What a treasure had he, my lord?

HAMLET  Why,
	One fair daughter, and no more,
	The which he loved passing well.

POLONIUS, [aside]  Still on my daughter.

HAMLET  Am I not i' th' right, old Jephthah?

POLONIUS  If you call me "Jephthah," my lord: I have a
daughter that I love passing well.

HAMLET  Nay, that follows not.

POLONIUS  What follows then, my lord?

HAMLET  Why,
	As by lot, God wot
and then, you know,
	It came to pass, as most like it was--
the first row of the pious chanson will show you
more, for look where my abridgment comes.

[Enter the Players.]

You are welcome, masters; welcome all.--I am glad
to see thee well.--Welcome, good friends.--O my
old friend! Why, thy face is valanced since I saw thee
last. Com'st thou to beard me in Denmark?--What,
my young lady and mistress! By 'r Lady, your Ladyship
is nearer to heaven than when I saw you last, by
the altitude of a chopine. Pray God your voice, like a
piece of uncurrent gold, be not cracked within the
ring. Masters, you are all welcome. We'll e'en to 't
like French falconers, fly at anything we see. We'll
have a speech straight. Come, give us a taste of your
quality. Come, a passionate speech.

FIRST PLAYER  What speech, my good lord?

HAMLET  I heard thee speak me a speech once, but it
was never acted, or, if it was, not above once; for
the play, I remember, pleased not the million:
'twas caviary to the general. But it was (as I
received it, and others whose judgments in such
matters cried in the top of mine) an excellent play,
well digested in the scenes, set down with as much
modesty as cunning. I remember one said there
were no sallets in the lines to make the matter
savory, nor no matter in the phrase that might indict
the author of affection, but called it an honest
method, as wholesome as sweet and, by very much,
more handsome than fine. One speech in 't I
chiefly loved. 'Twas Aeneas' tale to Dido, and
thereabout of it especially when he speaks of
Priam's slaughter. If it live in your memory, begin at
this line--let me see, let me see:
The rugged Pyrrhus, like th' Hyrcanian beast--
'tis not so; it begins with Pyrrhus:
The rugged Pyrrhus, he whose sable arms,
Black as his purpose, did the night resemble
When he lay couched in th' ominous horse,
Hath now this dread and black complexion smeared
With heraldry more dismal. Head to foot,
Now is he total gules, horridly tricked
With blood of fathers, mothers, daughters, sons,
Baked and impasted with the parching streets,
That lend a tyrannous and a damned light
To their lord's murder. Roasted in wrath and fire,
And thus o'ersized with coagulate gore,
With eyes like carbuncles, the hellish Pyrrhus
Old grandsire Priam seeks.
So, proceed you.

POLONIUS  'Fore God, my lord, well spoken, with good
accent and good discretion.

FIRST PLAYER  Anon he finds him
Striking too short at Greeks. His antique sword,
Rebellious to his arm, lies where it falls,
Repugnant to command. Unequal matched,
Pyrrhus at Priam drives, in rage strikes wide;
But with the whiff and wind of his fell sword
Th' unnerved father falls. Then senseless Ilium,
Seeming to feel this blow, with flaming top
Stoops to his base, and with a hideous crash
Takes prisoner Pyrrhus' ear. For lo, his sword,
Which was declining on the milky head
Of reverend Priam, seemed i' th' air to stick.
So as a painted tyrant Pyrrhus stood
And, like a neutral to his will and matter,
Did nothing.
But as we often see against some storm
A silence in the heavens, the rack stand still,
The bold winds speechless, and the orb below
As hush as death, anon the dreadful thunder
Doth rend the region; so, after Pyrrhus' pause,
Aroused vengeance sets him new a-work,
And never did the Cyclops' hammers fall
On Mars's armor, forged for proof eterne,
With less remorse than Pyrrhus' bleeding sword
Now falls on Priam.
Out, out, thou strumpet Fortune! All you gods
In general synod take away her power,
Break all the spokes and fellies from her wheel,
And bowl the round nave down the hill of heaven
As low as to the fiends!

POLONIUS  This is too long.

HAMLET  It shall to the barber's with your beard.--
Prithee say on. He's for a jig or a tale of bawdry, or
he sleeps. Say on; come to Hecuba.

FIRST PLAYER
But who, ah woe, had seen the mobled queen--

HAMLET  "The mobled queen"?

POLONIUS  That's good. "Mobled queen" is good.

FIRST PLAYER
Run barefoot up and down, threat'ning the flames
With bisson rheum, a clout upon that head
Where late the diadem stood, and for a robe,
About her lank and all o'erteemed loins
A blanket, in the alarm of fear caught up--
Who this had seen, with tongue in venom steeped,
'Gainst Fortune's state would treason have
pronounced.
But if the gods themselves did see her then
When she saw Pyrrhus make malicious sport
In mincing with his sword her husband's limbs,
The instant burst of clamor that she made
(Unless things mortal move them not at all)
Would have made milch the burning eyes of heaven
And passion in the gods.

POLONIUS  Look whe'er he has not turned his color and
has tears in 's eyes. Prithee, no more.

HAMLET  'Tis well. I'll have thee speak out the rest of
this soon.--Good my lord, will you see the players
well bestowed? Do you hear, let them be well used,
for they are the abstract and brief chronicles of the
time. After your death you were better have a bad
epitaph than their ill report while you live.

POLONIUS  My lord, I will use them according to their
desert.

HAMLET  God's bodykins, man, much better! Use every
man after his desert and who shall 'scape
whipping? Use them after your own honor and
dignity. The less they deserve, the more merit is in
your bounty. Take them in.

POLONIUS  Come, sirs.

HAMLET  Follow him, friends. We'll hear a play
tomorrow. [As Polonius and Players exit, Hamlet speaks to
the First Player.] Dost thou hear me, old friend? Can
you play "The Murder of Gonzago"?

FIRST PLAYER  Ay, my lord.

HAMLET  We'll ha 't tomorrow night. You could, for a
need, study a speech of some dozen or sixteen
lines, which I would set down and insert in 't,
could you not?

FIRST PLAYER  Ay, my lord.

HAMLET  Very well. Follow that lord--and look you
mock him not. [First Player exits.] My good friends,
I'll leave you till night. You are welcome to Elsinore.

ROSENCRANTZ  Good my lord.

HAMLET
Ay, so, good-bye to you.
[Rosencrantz and Guildenstern exit.]
Now I am alone.
O, what a rogue and peasant slave am I!
Is it not monstrous that this player here,
But in a fiction, in a dream of passion,
Could force his soul so to his own conceit
That from her working all his visage wanned,
Tears in his eyes, distraction in his aspect,
A broken voice, and his whole function suiting
With forms to his conceit--and all for nothing!
For Hecuba!
What's Hecuba to him, or he to Hecuba,
That he should weep for her? What would he do
Had he the motive and the cue for passion
That I have? He would drown the stage with tears
And cleave the general ear with horrid speech,
Make mad the guilty and appall the free,
Confound the ignorant and amaze indeed
The very faculties of eyes and ears. Yet I,
A dull and muddy-mettled rascal, peak
Like John-a-dreams, unpregnant of my cause,
And can say nothing--no, not for a king
Upon whose property and most dear life
A damned defeat was made. Am I a coward?
Who calls me "villain"? breaks my pate across?
Plucks off my beard and blows it in my face?
Tweaks me by the nose? gives me the lie i' th' throat
As deep as to the lungs? Who does me this?
Ha! 'Swounds, I should take it! For it cannot be
But I am pigeon-livered and lack gall
To make oppression bitter, or ere this
I should have fatted all the region kites
With this slave's offal. Bloody, bawdy villain!
Remorseless, treacherous, lecherous, kindless
villain!
O vengeance!
Why, what an ass am I! This is most brave,
That I, the son of a dear father murdered,
Prompted to my revenge by heaven and hell,
Must, like a whore, unpack my heart with words
And fall a-cursing like a very drab,
A stallion! Fie upon 't! Foh!
About, my brains!--Hum, I have heard
That guilty creatures sitting at a play
Have, by the very cunning of the scene,
Been struck so to the soul that presently
They have proclaimed their malefactions;
For murder, though it have no tongue, will speak
With most miraculous organ. I'll have these players
Play something like the murder of my father
Before mine uncle. I'll observe his looks;
I'll tent him to the quick. If he do blench,
I know my course. The spirit that I have seen
May be a devil, and the devil hath power
T' assume a pleasing shape; yea, and perhaps,
Out of my weakness and my melancholy,
As he is very potent with such spirits,
Abuses me to damn me. I'll have grounds
More relative than this. The play's the thing
Wherein I'll catch the conscience of the King.
[He exits.]


ACT 3
=====

Scene 1
=======
[Enter King, Queen, Polonius, Ophelia, Rosencrantz,
Guildenstern, and Lords.]


KING
And can you by no drift of conference
Get from him why he puts on this confusion,
Grating so harshly all his days of quiet
With turbulent and dangerous lunacy?

ROSENCRANTZ
He does confess he feels himself distracted,
But from what cause he will by no means speak.

GUILDENSTERN
Nor do we find him forward to be sounded,
But with a crafty madness keeps aloof
When we would bring him on to some confession
Of his true state.

QUEEN  Did he receive you well?

ROSENCRANTZ  Most like a gentleman.

GUILDENSTERN
But with much forcing of his disposition.

ROSENCRANTZ
Niggard of question, but of our demands
Most free in his reply.

QUEEN  Did you assay him to any pastime?

ROSENCRANTZ
Madam, it so fell out that certain players
We o'erraught on the way. Of these we told him,
And there did seem in him a kind of joy
To hear of it. They are here about the court,
And, as I think, they have already order
This night to play before him.

POLONIUS  'Tis most true,
And he beseeched me to entreat your Majesties
To hear and see the matter.

KING
With all my heart, and it doth much content me
To hear him so inclined.
Good gentlemen, give him a further edge
And drive his purpose into these delights.

ROSENCRANTZ
We shall, my lord.	[Rosencrantz and Guildenstern
and Lords exit.]

KING  Sweet Gertrude, leave us too,
For we have closely sent for Hamlet hither,
That he, as 'twere by accident, may here
Affront Ophelia.
Her father and myself, lawful espials,
Will so bestow ourselves that, seeing unseen,
We may of their encounter frankly judge
And gather by him, as he is behaved,
If 't be th' affliction of his love or no
That thus he suffers for.

QUEEN  I shall obey you.
And for your part, Ophelia, I do wish
That your good beauties be the happy cause
Of Hamlet's wildness. So shall I hope your virtues
Will bring him to his wonted way again,
To both your honors.

OPHELIA  Madam, I wish it may.
[Queen exits.]

POLONIUS
Ophelia, walk you here.--Gracious, so please you,
We will bestow ourselves. [To Ophelia.] Read on this
book,
That show of such an exercise may color
Your loneliness.--We are oft to blame in this
('Tis too much proved), that with devotion's visage
And pious action we do sugar o'er
The devil himself.

KING, [aside]  O, 'tis too true!
How smart a lash that speech doth give my
conscience.
The harlot's cheek beautied with plast'ring art
Is not more ugly to the thing that helps it
Than is my deed to my most painted word.
O heavy burden!

POLONIUS
I hear him coming. Let's withdraw, my lord.
[They withdraw.]

[Enter Hamlet.]


HAMLET
To be or not to be--that is the question:
Whether 'tis nobler in the mind to suffer
The slings and arrows of outrageous fortune,
Or to take arms against a sea of troubles
And, by opposing, end them. To die, to sleep--
No more--and by a sleep to say we end
The heartache and the thousand natural shocks
That flesh is heir to--'tis a consummation
Devoutly to be wished. To die, to sleep--
To sleep, perchance to dream. Ay, there's the rub,
For in that sleep of death what dreams may come,
When we have shuffled off this mortal coil,
Must give us pause. There's the respect
That makes calamity of so long life.
For who would bear the whips and scorns of time,
Th' oppressor's wrong, the proud man's contumely,
The pangs of despised love, the law's delay,
The insolence of office, and the spurns
That patient merit of th' unworthy takes,
When he himself might his quietus make
With a bare bodkin? Who would fardels bear,
To grunt and sweat under a weary life,
But that the dread of something after death,
The undiscovered country from whose bourn
No traveler returns, puzzles the will
And makes us rather bear those ills we have
Than fly to others that we know not of?
Thus conscience does make cowards of us all,
And thus the native hue of resolution
Is sicklied o'er with the pale cast of thought,
And enterprises of great pitch and moment
With this regard their currents turn awry
And lose the name of action.--Soft you now,
The fair Ophelia.--Nymph, in thy orisons
Be all my sins remembered.

OPHELIA  Good my lord,
How does your Honor for this many a day?

HAMLET  I humbly thank you, well.

OPHELIA
My lord, I have remembrances of yours
That I have longed long to redeliver.
I pray you now receive them.

HAMLET
No, not I. I never gave you aught.

OPHELIA
My honored lord, you know right well you did,
And with them words of so sweet breath composed
As made the things more rich. Their perfume
lost,
Take these again, for to the noble mind
Rich gifts wax poor when givers prove unkind.
There, my lord.

HAMLET  Ha, ha, are you honest?

OPHELIA  My lord?

HAMLET  Are you fair?

OPHELIA  What means your Lordship?

HAMLET  That if you be honest and fair, your honesty
should admit no discourse to your beauty.

OPHELIA  Could beauty, my lord, have better commerce
than with honesty?

HAMLET  Ay, truly, for the power of beauty will sooner
transform honesty from what it is to a bawd than
the force of honesty can translate beauty into his
likeness. This was sometime a paradox, but now
the time gives it proof. I did love you once.

OPHELIA  Indeed, my lord, you made me believe so.

HAMLET  You should not have believed me, for virtue
cannot so inoculate our old stock but we shall
relish of it. I loved you not.

OPHELIA  I was the more deceived.

HAMLET  Get thee to a nunnery. Why wouldst thou be
a breeder of sinners? I am myself indifferent honest,
but yet I could accuse me of such things that it
were better my mother had not borne me: I am
very proud, revengeful, ambitious, with more offenses
at my beck than I have thoughts to put them
in, imagination to give them shape, or time to act
them in. What should such fellows as I do crawling
between earth and heaven? We are arrant knaves
all; believe none of us. Go thy ways to a nunnery.
Where's your father?

OPHELIA  At home, my lord.

HAMLET  Let the doors be shut upon him that he may
play the fool nowhere but in 's own house. Farewell.

OPHELIA  O, help him, you sweet heavens!

HAMLET  If thou dost marry, I'll give thee this plague
for thy dowry: be thou as chaste as ice, as pure as
snow, thou shalt not escape calumny. Get thee to a
nunnery, farewell. Or if thou wilt needs marry,
marry a fool, for wise men know well enough what
monsters you make of them. To a nunnery, go, and
quickly too. Farewell.

OPHELIA  Heavenly powers, restore him!

HAMLET  I have heard of your paintings too, well
enough. God hath given you one face, and you
make yourselves another. You jig and amble, and
you lisp; you nickname God's creatures and make
your wantonness your ignorance. Go to, I'll no
more on 't. It hath made me mad. I say we will have
no more marriage. Those that are married already,
all but one, shall live. The rest shall keep as they are.
To a nunnery, go.	[He exits.]

OPHELIA
O, what a noble mind is here o'erthrown!
The courtier's, soldier's, scholar's, eye, tongue,
sword,
Th' expectancy and rose of the fair state,
The glass of fashion and the mold of form,
Th' observed of all observers, quite, quite down!
And I, of ladies most deject and wretched,
That sucked the honey of his musicked vows,
Now see that noble and most sovereign reason,
Like sweet bells jangled, out of time and harsh;
That unmatched form and stature of blown youth
Blasted with ecstasy. O, woe is me
T' have seen what I have seen, see what I see!

KING, [advancing with Polonius]
Love? His affections do not that way tend;
Nor what he spake, though it lacked form a little,
Was not like madness. There's something in his soul
O'er which his melancholy sits on brood,
And I do doubt the hatch and the disclose
Will be some danger; which for to prevent,
I have in quick determination
Thus set it down: he shall with speed to England
For the demand of our neglected tribute.
Haply the seas, and countries different,
With variable objects, shall expel
This something-settled matter in his heart,
Whereon his brains still beating puts him thus
From fashion of himself. What think you on 't?

POLONIUS
It shall do well. But yet do I believe
The origin and commencement of his grief
Sprung from neglected love.--How now, Ophelia?
You need not tell us what Lord Hamlet said;
We heard it all.--My lord, do as you please,
But, if you hold it fit, after the play
Let his queen-mother all alone entreat him
To show his grief. Let her be round with him;
And I'll be placed, so please you, in the ear
Of all their conference. If she find him not,
To England send him, or confine him where
Your wisdom best shall think.

KING  It shall be so.
Madness in great ones must not unwatched go.
[They exit.]

Scene 2
=======
[Enter Hamlet and three of the Players.]


HAMLET  Speak the speech, I pray you, as I pronounced
it to you, trippingly on the tongue; but if you mouth
it, as many of our players do, I had as lief the
town-crier spoke my lines. Nor do not saw the air
too much with your hand, thus, but use all gently;
for in the very torrent, tempest, and, as I may say,
whirlwind of your passion, you must acquire and
beget a temperance that may give it smoothness. O,
it offends me to the soul to hear a robustious,
periwig-pated fellow tear a passion to tatters, to very
rags, to split the ears of the groundlings, who for the
most part are capable of nothing but inexplicable
dumb shows and noise. I would have such a fellow
whipped for o'erdoing Termagant. It out-Herods
Herod. Pray you, avoid it.

PLAYER  I warrant your Honor.

HAMLET  Be not too tame neither, but let your own
discretion be your tutor. Suit the action to the
word, the word to the action, with this special
observance, that you o'erstep not the modesty of
nature. For anything so o'erdone is from the purpose
of playing, whose end, both at the first and
now, was and is to hold, as 'twere, the mirror up to
nature, to show virtue her own feature, scorn her
own image, and the very age and body of the time
his form and pressure. Now this overdone or come
tardy off, though it makes the unskillful laugh,
cannot but make the judicious grieve, the censure
of the which one must in your allowance o'erweigh
a whole theater of others. O, there be players that I
have seen play and heard others praise (and that
highly), not to speak it profanely, that, neither
having th' accent of Christians nor the gait of
Christian, pagan, nor man, have so strutted and
bellowed that I have thought some of nature's
journeymen had made men, and not made them
well, they imitated humanity so abominably.

PLAYER  I hope we have reformed that indifferently
with us, sir.

HAMLET  O, reform it altogether. And let those that play
your clowns speak no more than is set down for
them, for there be of them that will themselves
laugh, to set on some quantity of barren spectators
to laugh too, though in the meantime some necessary
question of the play be then to be considered.
That's villainous and shows a most pitiful ambition
in the fool that uses it. Go make you ready.
[Players exit.]

[Enter Polonius, Guildenstern, and Rosencrantz.]

How now, my lord, will the King hear this piece of
work?

POLONIUS  And the Queen too, and that presently.

HAMLET  Bid the players make haste.	[Polonius exits.]
Will you two help to hasten them?

ROSENCRANTZ  Ay, my lord.	[They exit.]

HAMLET  What ho, Horatio!

[Enter Horatio.]


HORATIO  Here, sweet lord, at your service.

HAMLET
Horatio, thou art e'en as just a man
As e'er my conversation coped withal.

HORATIO
O, my dear lord--

HAMLET  Nay, do not think I flatter,
For what advancement may I hope from thee
That no revenue hast but thy good spirits
To feed and clothe thee? Why should the poor be
flattered?
No, let the candied tongue lick absurd pomp
And crook the pregnant hinges of the knee
Where thrift may follow fawning. Dost thou hear?
Since my dear soul was mistress of her choice
And could of men distinguish, her election
Hath sealed thee for herself. For thou hast been
As one in suffering all that suffers nothing,
A man that Fortune's buffets and rewards
Hast ta'en with equal thanks; and blessed are those
Whose blood and judgment are so well
commeddled
That they are not a pipe for Fortune's finger
To sound what stop she please. Give me that man
That is not passion's slave, and I will wear him
In my heart's core, ay, in my heart of heart,
As I do thee.--Something too much of this.--
There is a play tonight before the King.
One scene of it comes near the circumstance
Which I have told thee of my father's death.
I prithee, when thou seest that act afoot,
Even with the very comment of thy soul
Observe my uncle. If his occulted guilt
Do not itself unkennel in one speech,
It is a damned ghost that we have seen,
And my imaginations are as foul
As Vulcan's stithy. Give him heedful note,
For I mine eyes will rivet to his face,
And, after, we will both our judgments join
In censure of his seeming.

HORATIO  Well, my lord.
If he steal aught the whilst this play is playing
And 'scape detecting, I will pay the theft.
[Sound a flourish.]

HAMLET  They are coming to the play. I must be idle.
Get you a place.

[Enter Trumpets and Kettle Drums. Enter King, Queen,
Polonius, Ophelia, Rosencrantz, Guildenstern, and other
Lords attendant with the King's guard carrying
torches.]


KING  How fares our cousin Hamlet?

HAMLET  Excellent, i' faith, of the chameleon's dish. I
eat the air, promise-crammed. You cannot feed
capons so.

KING  I have nothing with this answer, Hamlet. These
words are not mine.

HAMLET  No, nor mine now. [To Polonius.] My lord, you
played once i' th' university, you say?

POLONIUS  That did I, my lord, and was accounted a
good actor.

HAMLET  What did you enact?

POLONIUS  I did enact Julius Caesar. I was killed i' th'
Capitol. Brutus killed me.

HAMLET  It was a brute part of him to kill so capital a
calf there.--Be the players ready?

ROSENCRANTZ  Ay, my lord. They stay upon your
patience.

QUEEN  Come hither, my dear Hamlet, sit by me.

HAMLET  No, good mother. Here's metal more
attractive.	[Hamlet takes a place near Ophelia.]

POLONIUS, [to the King]  Oh, ho! Do you mark that?

HAMLET  Lady, shall I lie in your lap?

OPHELIA  No, my lord.

HAMLET  I mean, my head upon your lap?

OPHELIA  Ay, my lord.

HAMLET  Do you think I meant country matters?

OPHELIA  I think nothing, my lord.

HAMLET  That's a fair thought to lie between maids'
legs.

OPHELIA  What is, my lord?

HAMLET  Nothing.

OPHELIA  You are merry, my lord.

HAMLET  Who, I?

OPHELIA  Ay, my lord.

HAMLET  O God, your only jig-maker. What should a
man do but be merry? For look you how cheerfully
my mother looks, and my father died within 's two
hours.

OPHELIA  Nay, 'tis twice two months, my lord.

HAMLET  So long? Nay, then, let the devil wear black,
for I'll have a suit of sables. O heavens, die two
months ago, and not forgotten yet? Then there's
hope a great man's memory may outlive his life half
a year. But, by 'r Lady, he must build churches, then,
or else shall he suffer not thinking on, with the
hobby-horse, whose epitaph is "For oh, for oh, the
hobby-horse is forgot."
[The trumpets sounds. Dumb show follows.]

[Enter a King and a Queen, very lovingly, the Queen
embracing him and he her. She kneels and makes show of
protestation unto him. He takes her up and declines his
head upon her neck. He lies him down upon a bank of
flowers. She, seeing him asleep, leaves him. Anon
comes in another man, takes off his crown, kisses it, pours
poison in the sleeper's ears, and leaves him. The Queen
returns, finds the King dead, makes passionate action. The
poisoner with some three or four come in again, seem to
condole with her. The dead body is carried away. The
poisoner woos the Queen with gifts. She seems harsh
awhile but in the end accepts his love.]
[Players exit.]

OPHELIA  What means this, my lord?

HAMLET  Marry, this is miching mallecho. It means
mischief.

OPHELIA  Belike this show imports the argument of the
play.

[Enter Prologue.]


HAMLET  We shall know by this fellow. The players
cannot keep counsel; they'll tell all.

OPHELIA  Will he tell us what this show meant?

HAMLET  Ay, or any show that you will show him. Be
not you ashamed to show, he'll not shame to tell you
what it means.

OPHELIA  You are naught, you are naught. I'll mark the
play.

PROLOGUE
	For us and for our tragedy,
	Here stooping to your clemency,
	We beg your hearing patiently.	[He exits.]

HAMLET  Is this a prologue or the posy of a ring?

OPHELIA  'Tis brief, my lord.

HAMLET  As woman's love.

[Enter the Player King and Queen.]


PLAYER KING
Full thirty times hath Phoebus' cart gone round
Neptune's salt wash and Tellus' orbed ground,
And thirty dozen moons with borrowed sheen
About the world have times twelve thirties been
Since love our hearts and Hymen did our hands
Unite commutual in most sacred bands.

PLAYER QUEEN
So many journeys may the sun and moon
Make us again count o'er ere love be done!
But woe is me! You are so sick of late,
So far from cheer and from your former state,
That I distrust you. Yet, though I distrust,
Discomfort you, my lord, it nothing must.
For women fear too much, even as they love,
And women's fear and love hold quantity,
In neither aught, or in extremity.
Now what my love is, proof hath made you know,
And, as my love is sized, my fear is so:
Where love is great, the littlest doubts are fear;
Where little fears grow great, great love grows there.

PLAYER KING
Faith, I must leave thee, love, and shortly too.
My operant powers their functions leave to do.
And thou shalt live in this fair world behind,
Honored, beloved; and haply one as kind
For husband shalt thou--

PLAYER QUEEN  O, confound the rest!
Such love must needs be treason in my breast.
In second husband let me be accurst.
None wed the second but who killed the first.

HAMLET  That's wormwood!

PLAYER QUEEN
The instances that second marriage move
Are base respects of thrift, but none of love.
A second time I kill my husband dead
When second husband kisses me in bed.

PLAYER KING
I do believe you think what now you speak,
But what we do determine oft we break.
Purpose is but the slave to memory,
Of violent birth, but poor validity,
Which now, the fruit unripe, sticks on the tree
But fall unshaken when they mellow be.
Most necessary 'tis that we forget
To pay ourselves what to ourselves is debt.
What to ourselves in passion we propose,
The passion ending, doth the purpose lose.
The violence of either grief or joy
Their own enactures with themselves destroy.
Where joy most revels, grief doth most lament;
Grief joys, joy grieves, on slender accident.
This world is not for aye, nor 'tis not strange
That even our loves should with our fortunes change;
For 'tis a question left us yet to prove
Whether love lead fortune or else fortune love.
The great man down, you mark his favorite flies;
The poor, advanced, makes friends of enemies.
And hitherto doth love on fortune tend,
For who not needs shall never lack a friend,
And who in want a hollow friend doth try
Directly seasons him his enemy.
But, orderly to end where I begun:
Our wills and fates do so contrary run
That our devices still are overthrown;
Our thoughts are ours, their ends none of our own.
So think thou wilt no second husband wed,
But die thy thoughts when thy first lord is dead.

PLAYER QUEEN
Nor Earth to me give food, nor heaven light,
Sport and repose lock from me day and night,
To desperation turn my trust and hope,
An anchor's cheer in prison be my scope.
Each opposite that blanks the face of joy
Meet what I would have well and it destroy.
Both here and hence pursue me lasting strife,
If, once a widow, ever I be wife.

HAMLET  If she should break it now!

PLAYER KING
'Tis deeply sworn. Sweet, leave me here awhile.
My spirits grow dull, and fain I would beguile
The tedious day with sleep.	[Sleeps.]

PLAYER QUEEN  Sleep rock thy brain,
And never come mischance between us twain.
[Player Queen exits.]

HAMLET  Madam, how like you this play?

QUEEN  The lady doth protest too much, methinks.

HAMLET  O, but she'll keep her word.

KING  Have you heard the argument? Is there no
offense in 't?

HAMLET  No, no, they do but jest, poison in jest. No
offense i' th' world.

KING  What do you call the play?

HAMLET  "The Mousetrap." Marry, how? Tropically.
This play is the image of a murder done in Vienna.
Gonzago is the duke's name, his wife Baptista. You
shall see anon. 'Tis a knavish piece of work, but
what of that? Your Majesty and we that have free
souls, it touches us not. Let the galled jade wince;
our withers are unwrung.

[Enter Lucianus.]

This is one Lucianus, nephew to the king.

OPHELIA  You are as good as a chorus, my lord.

HAMLET  I could interpret between you and your love,
if I could see the puppets dallying.

OPHELIA  You are keen, my lord, you are keen.

HAMLET  It would cost you a groaning to take off mine
edge.

OPHELIA  Still better and worse.

HAMLET  So you mis-take your husbands.--Begin,
murderer. Pox, leave thy damnable faces and
begin. Come, the croaking raven doth bellow for
revenge.

LUCIANUS
Thoughts black, hands apt, drugs fit, and time
agreeing,
Confederate season, else no creature seeing,
Thou mixture rank, of midnight weeds collected,
With Hecate's ban thrice blasted, thrice infected,
Thy natural magic and dire property
On wholesome life usurp immediately.
[Pours the poison in his ears.]

HAMLET  He poisons him i' th' garden for his estate. His
name's Gonzago. The story is extant and written in
very choice Italian. You shall see anon how the
murderer gets the love of Gonzago's wife.
[Claudius rises.]

OPHELIA  The King rises.

HAMLET  What, frighted with false fire?

QUEEN  How fares my lord?

POLONIUS  Give o'er the play.

KING  Give me some light. Away!

POLONIUS  Lights, lights, lights!
[All but Hamlet and Horatio exit.]

HAMLET
	Why, let the strucken deer go weep,
	   The hart ungalled play.
	For some must watch, while some must sleep:
	   Thus runs the world away.
Would not this, sir, and a forest of feathers (if the
rest of my fortunes turn Turk with me) with two
Provincial roses on my razed shoes, get me a
fellowship in a cry of players?

HORATIO  Half a share.

HAMLET  A whole one, I.
	For thou dost know, O Damon dear,
	   This realm dismantled was
	Of Jove himself, and now reigns here
	   A very very--pajock.

HORATIO  You might have rhymed.

HAMLET  O good Horatio, I'll take the ghost's word for
a thousand pound. Didst perceive?

HORATIO  Very well, my lord.

HAMLET  Upon the talk of the poisoning?

HORATIO  I did very well note him.

HAMLET  Ah ha! Come, some music! Come, the
recorders!
	For if the King like not the comedy,
	Why, then, belike he likes it not, perdy.
Come, some music!

[Enter Rosencrantz and Guildenstern.]


GUILDENSTERN  Good my lord, vouchsafe me a word
with you.

HAMLET  Sir, a whole history.

GUILDENSTERN  The King, sir--

HAMLET  Ay, sir, what of him?

GUILDENSTERN  Is in his retirement marvelous
distempered.

HAMLET  With drink, sir?

GUILDENSTERN  No, my lord, with choler.

HAMLET  Your wisdom should show itself more richer
to signify this to the doctor, for for me to put him to
his purgation would perhaps plunge him into more
choler.

GUILDENSTERN  Good my lord, put your discourse into
some frame and start not so wildly from my
affair.

HAMLET  I am tame, sir. Pronounce.

GUILDENSTERN  The Queen your mother, in most great
affliction of spirit, hath sent me to you.

HAMLET  You are welcome.

GUILDENSTERN  Nay, good my lord, this courtesy is not
of the right breed. If it shall please you to make me
a wholesome answer, I will do your mother's
commandment. If not, your pardon and my return
shall be the end of my business.

HAMLET  Sir, I cannot.

ROSENCRANTZ  What, my lord?

HAMLET  Make you a wholesome answer. My wit's
diseased. But, sir, such answer as I can make, you
shall command--or, rather, as you say, my mother.
Therefore no more but to the matter. My mother,
you say--

ROSENCRANTZ  Then thus she says: your behavior hath
struck her into amazement and admiration.

HAMLET  O wonderful son that can so 'stonish a mother!
But is there no sequel at the heels of this
mother's admiration? Impart.

ROSENCRANTZ  She desires to speak with you in her
closet ere you go to bed.

HAMLET  We shall obey, were she ten times our mother.
Have you any further trade with us?

ROSENCRANTZ  My lord, you once did love me.

HAMLET  And do still, by these pickers and stealers.

ROSENCRANTZ  Good my lord, what is your cause of
distemper? You do surely bar the door upon your
own liberty if you deny your griefs to your friend.

HAMLET  Sir, I lack advancement.

ROSENCRANTZ  How can that be, when you have the
voice of the King himself for your succession in
Denmark?

HAMLET  Ay, sir, but "While the grass grows"--the
proverb is something musty.

[Enter the Players with recorders.]

O, the recorders! Let me see one. [He takes a
recorder and turns to Guildenstern.] To withdraw
with you: why do you go about to recover the wind
of me, as if you would drive me into a toil?

GUILDENSTERN  O, my lord, if my duty be too bold, my
love is too unmannerly.

HAMLET  I do not well understand that. Will you play
upon this pipe?

GUILDENSTERN  My lord, I cannot.

HAMLET  I pray you.

GUILDENSTERN  Believe me, I cannot.

HAMLET  I do beseech you.

GUILDENSTERN  I know no touch of it, my lord.

HAMLET  It is as easy as lying. Govern these ventages
with your fingers and thumb, give it breath with
your mouth, and it will discourse most eloquent
music. Look you, these are the stops.

GUILDENSTERN  But these cannot I command to any
utt'rance of harmony. I have not the skill.

HAMLET  Why, look you now, how unworthy a thing
you make of me! You would play upon me, you
would seem to know my stops, you would pluck
out the heart of my mystery, you would sound me
from my lowest note to the top of my compass;
and there is much music, excellent voice, in this
little organ, yet cannot you make it speak. 'Sblood,
do you think I am easier to be played on than a pipe?
Call me what instrument you will, though you can
fret me, you cannot play upon me.

[Enter Polonius.]

God bless you, sir.

POLONIUS  My lord, the Queen would speak with you,
and presently.

HAMLET  Do you see yonder cloud that's almost in
shape of a camel?

POLONIUS  By th' Mass, and 'tis like a camel indeed.

HAMLET  Methinks it is like a weasel.

POLONIUS  It is backed like a weasel.

HAMLET  Or like a whale.

POLONIUS  Very like a whale.

HAMLET  Then I will come to my mother by and by.
[Aside.] They fool me to the top of my bent.--I will
come by and by.

POLONIUS  I will say so.

HAMLET  "By and by" is easily said. Leave me,
friends.
[All but Hamlet exit.]
'Tis now the very witching time of night,
When churchyards yawn and hell itself breathes
out
Contagion to this world. Now could I drink hot
blood
And do such bitter business as the day
Would quake to look on. Soft, now to my mother.
O heart, lose not thy nature; let not ever
The soul of Nero enter this firm bosom.
Let me be cruel, not unnatural.
I will speak daggers to her, but use none.
My tongue and soul in this be hypocrites:
How in my words somever she be shent,
To give them seals never, my soul, consent.
[He exits.]

Scene 3
=======
[Enter King, Rosencrantz, and Guildenstern.]


KING
I like him not, nor stands it safe with us
To let his madness range. Therefore prepare you.
I your commission will forthwith dispatch,
And he to England shall along with you.
The terms of our estate may not endure
Hazard so near 's as doth hourly grow
Out of his brows.

GUILDENSTERN  We will ourselves provide.
Most holy and religious fear it is
To keep those many many bodies safe
That live and feed upon your Majesty.

ROSENCRANTZ
The single and peculiar life is bound
With all the strength and armor of the mind
To keep itself from noyance, but much more
That spirit upon whose weal depends and rests
The lives of many. The cess of majesty
Dies not alone, but like a gulf doth draw
What's near it with it; or it is a massy wheel
Fixed on the summit of the highest mount,
To whose huge spokes ten thousand lesser things
Are mortised and adjoined, which, when it falls,
Each small annexment, petty consequence,
Attends the boist'rous ruin. Never alone
Did the king sigh, but with a general groan.

KING
Arm you, I pray you, to this speedy voyage,
For we will fetters put about this fear,
Which now goes too free-footed.

ROSENCRANTZ  We will haste us.
[Rosencrantz and Guildenstern exit.]

[Enter Polonius.]


POLONIUS
My lord, he's going to his mother's closet.
Behind the arras I'll convey myself
To hear the process. I'll warrant she'll tax him
home;
And, as you said (and wisely was it said),
'Tis meet that some more audience than a mother,
Since nature makes them partial, should o'erhear
The speech of vantage. Fare you well, my liege.
I'll call upon you ere you go to bed
And tell you what I know.

KING  Thanks, dear my lord.
[Polonius exits.]
O, my offense is rank, it smells to heaven;
It hath the primal eldest curse upon 't,
A brother's murder. Pray can I not,
Though inclination be as sharp as will.
My stronger guilt defeats my strong intent,
And, like a man to double business bound,
I stand in pause where I shall first begin
And both neglect. What if this cursed hand
Were thicker than itself with brother's blood?
Is there not rain enough in the sweet heavens
To wash it white as snow? Whereto serves mercy
But to confront the visage of offense?
And what's in prayer but this twofold force,
To be forestalled ere we come to fall,
Or pardoned being down? Then I'll look up.
My fault is past. But, O, what form of prayer
Can serve my turn? "Forgive me my foul murder"?
That cannot be, since I am still possessed
Of those effects for which I did the murder:
My crown, mine own ambition, and my queen.
May one be pardoned and retain th' offense?
In the corrupted currents of this world,
Offense's gilded hand may shove by justice,
And oft 'tis seen the wicked prize itself
Buys out the law. But 'tis not so above:
There is no shuffling; there the action lies
In his true nature, and we ourselves compelled,
Even to the teeth and forehead of our faults,
To give in evidence. What then? What rests?
Try what repentance can. What can it not?
Yet what can it, when one cannot repent?
O wretched state! O bosom black as death!
O limed soul, that, struggling to be free,
Art more engaged! Help, angels! Make assay.
Bow, stubborn knees, and heart with strings of steel
Be soft as sinews of the newborn babe.
All may be well.	[He kneels.]

[Enter Hamlet.]


HAMLET
Now might I do it pat, now he is a-praying,
And now I'll do 't.	[He draws his sword.]
And so he goes to heaven,
And so am I revenged. That would be scanned:
A villain kills my father, and for that,
I, his sole son, do this same villain send
To heaven.
Why, this is hire and salary, not revenge.
He took my father grossly, full of bread,
With all his crimes broad blown, as flush as May;
And how his audit stands who knows save heaven.
But in our circumstance and course of thought
'Tis heavy with him. And am I then revenged
To take him in the purging of his soul,
When he is fit and seasoned for his passage?
No.
Up sword, and know thou a more horrid hent.
[He sheathes his sword.]
When he is drunk asleep, or in his rage,
Or in th' incestuous pleasure of his bed,
At game, a-swearing, or about some act
That has no relish of salvation in 't--
Then trip him, that his heels may kick at heaven,
And that his soul may be as damned and black
As hell, whereto it goes. My mother stays.
This physic but prolongs thy sickly days.
[Hamlet exits.]

KING, [rising]
My words fly up, my thoughts remain below;
Words without thoughts never to heaven go.
[He exits.]

Scene 4
=======
[Enter Queen and Polonius.]


POLONIUS
He will come straight. Look you lay home to him.
Tell him his pranks have been too broad to bear
with
And that your Grace hath screened and stood
between
Much heat and him. I'll silence me even here.
Pray you, be round with him.

HAMLET, [within]  Mother, mother, mother!

QUEEN  I'll warrant you. Fear me not. Withdraw,
I hear him coming.
[Polonius hides behind the arras.]

[Enter Hamlet.]


HAMLET  Now, mother, what's the matter?

QUEEN
Hamlet, thou hast thy father much offended.

HAMLET
Mother, you have my father much offended.

QUEEN
Come, come, you answer with an idle tongue.

HAMLET
Go, go, you question with a wicked tongue.

QUEEN
Why, how now, Hamlet?

HAMLET  What's the matter now?

QUEEN
Have you forgot me?

HAMLET  No, by the rood, not so.
You are the Queen, your husband's brother's wife,
And (would it were not so) you are my mother.

QUEEN
Nay, then I'll set those to you that can speak.

HAMLET
Come, come, and sit you down; you shall not budge.
You go not till I set you up a glass
Where you may see the inmost part of you.

QUEEN
What wilt thou do? Thou wilt not murder me?
Help, ho!

POLONIUS, [behind the arras]  What ho! Help!

HAMLET
How now, a rat? Dead for a ducat, dead.
[He kills Polonius by thrusting a rapier
through the arras.]

POLONIUS, [behind the arras]
O, I am slain!

QUEEN  O me, what hast thou done?

HAMLET  Nay, I know not. Is it the King?

QUEEN
O, what a rash and bloody deed is this!

HAMLET
A bloody deed--almost as bad, good mother,
As kill a king and marry with his brother.

QUEEN
As kill a king?

HAMLET  Ay, lady, it was my word.
[He pulls Polonius' body from behind the arras.]
Thou wretched, rash, intruding fool, farewell.
I took thee for thy better. Take thy fortune.
Thou find'st to be too busy is some danger.
[To Queen.] Leave wringing of your hands. Peace, sit
you down,
And let me wring your heart; for so I shall
If it be made of penetrable stuff,
If damned custom have not brazed it so
That it be proof and bulwark against sense.

QUEEN
What have I done, that thou dar'st wag thy tongue
In noise so rude against me?

HAMLET  Such an act
That blurs the grace and blush of modesty,
Calls virtue hypocrite, takes off the rose
From the fair forehead of an innocent love
And sets a blister there, makes marriage vows
As false as dicers' oaths--O, such a deed
As from the body of contraction plucks
The very soul, and sweet religion makes
A rhapsody of words! Heaven's face does glow
O'er this solidity and compound mass
With heated visage, as against the doom,
Is thought-sick at the act.

QUEEN  Ay me, what act
That roars so loud and thunders in the index?

HAMLET
Look here upon this picture and on this,
The counterfeit presentment of two brothers.
See what a grace was seated on this brow,
Hyperion's curls, the front of Jove himself,
An eye like Mars' to threaten and command,
A station like the herald Mercury
New-lighted on a heaven-kissing hill,
A combination and a form indeed
Where every god did seem to set his seal
To give the world assurance of a man.
This was your husband. Look you now what follows.
Here is your husband, like a mildewed ear
Blasting his wholesome brother. Have you eyes?
Could you on this fair mountain leave to feed
And batten on this moor? Ha! Have you eyes?
You cannot call it love, for at your age
The heyday in the blood is tame, it's humble
And waits upon the judgment; and what judgment
Would step from this to this? Sense sure you have,
Else could you not have motion; but sure that sense
Is apoplexed; for madness would not err,
Nor sense to ecstasy was ne'er so thralled,
But it reserved some quantity of choice
To serve in such a difference. What devil was 't
That thus hath cozened you at hoodman-blind?
Eyes without feeling, feeling without sight,
Ears without hands or eyes, smelling sans all,
Or but a sickly part of one true sense
Could not so mope. O shame, where is thy blush?
Rebellious hell,
If thou canst mutine in a matron's bones,
To flaming youth let virtue be as wax
And melt in her own fire. Proclaim no shame
When the compulsive ardor gives the charge,
Since frost itself as actively doth burn,
And reason panders will.

QUEEN  O Hamlet, speak no more!
Thou turn'st my eyes into my very soul,
And there I see such black and grained spots
As will not leave their tinct.

HAMLET  Nay, but to live
In the rank sweat of an enseamed bed,
Stewed in corruption, honeying and making love
Over the nasty sty!

QUEEN  O, speak to me no more!
These words like daggers enter in my ears.
No more, sweet Hamlet!

HAMLET  A murderer and a villain,
A slave that is not twentieth part the tithe
Of your precedent lord; a vice of kings,
A cutpurse of the empire and the rule,
That from a shelf the precious diadem stole
And put it in his pocket--

QUEEN  No more!

HAMLET  A king of shreds and patches--

[Enter Ghost.]

Save me and hover o'er me with your wings,
You heavenly guards!--What would your gracious
figure?

QUEEN  Alas, he's mad.

HAMLET
Do you not come your tardy son to chide,
That, lapsed in time and passion, lets go by
Th' important acting of your dread command?
O, say!

GHOST  Do not forget. This visitation
Is but to whet thy almost blunted purpose.
But look, amazement on thy mother sits.
O, step between her and her fighting soul.
Conceit in weakest bodies strongest works.
Speak to her, Hamlet.

HAMLET  How is it with you, lady?

QUEEN  Alas, how is 't with you,
That you do bend your eye on vacancy
And with th' incorporal air do hold discourse?
Forth at your eyes your spirits wildly peep,
And, as the sleeping soldiers in th' alarm,
Your bedded hair, like life in excrements,
Start up and stand an end. O gentle son,
Upon the heat and flame of thy distemper
Sprinkle cool patience! Whereon do you look?

HAMLET
On him, on him! Look you how pale he glares.
His form and cause conjoined, preaching to stones,
Would make them capable. [To the Ghost.] Do not
look upon me,
Lest with this piteous action you convert
My stern effects. Then what I have to do
Will want true color--tears perchance for blood.

QUEEN  To whom do you speak this?

HAMLET  Do you see nothing there?

QUEEN
Nothing at all; yet all that is I see.

HAMLET  Nor did you nothing hear?

QUEEN  No, nothing but ourselves.

HAMLET
Why, look you there, look how it steals away!
My father, in his habit as he lived!
Look where he goes even now out at the portal!
[Ghost exits.]

QUEEN
This is the very coinage of your brain.
This bodiless creation ecstasy
Is very cunning in.

HAMLET  Ecstasy?
My pulse as yours doth temperately keep time
And makes as healthful music. It is not madness
That I have uttered. Bring me to the test,
And I the matter will reword, which madness
Would gambol from. Mother, for love of grace,
Lay not that flattering unction to your soul
That not your trespass but my madness speaks.
It will but skin and film the ulcerous place,
Whiles rank corruption, mining all within,
Infects unseen. Confess yourself to heaven,
Repent what's past, avoid what is to come,
And do not spread the compost on the weeds
To make them ranker. Forgive me this my virtue,
For, in the fatness of these pursy times,
Virtue itself of vice must pardon beg,
Yea, curb and woo for leave to do him good.

QUEEN
O Hamlet, thou hast cleft my heart in twain!

HAMLET
O, throw away the worser part of it,
And live the purer with the other half!
Good night. But go not to my uncle's bed.
Assume a virtue if you have it not.
That monster, custom, who all sense doth eat,
Of habits devil, is angel yet in this,
That to the use of actions fair and good
He likewise gives a frock or livery
That aptly is put on. Refrain tonight,
And that shall lend a kind of easiness
To the next abstinence, the next more easy;
For use almost can change the stamp of nature
And either ... the devil or throw him out
With wondrous potency. Once more, good night,
And, when you are desirous to be blest,
I'll blessing beg of you. For this same lord
[Pointing to Polonius.]
I do repent; but heaven hath pleased it so
To punish me with this and this with me,
That I must be their scourge and minister.
I will bestow him and will answer well
The death I gave him. So, again, good night.
I must be cruel only to be kind.
This bad begins, and worse remains behind.
One word more, good lady.

QUEEN  What shall I do?

HAMLET
Not this by no means that I bid you do:
Let the bloat king tempt you again to bed,
Pinch wanton on your cheek, call you his mouse,
And let him, for a pair of reechy kisses
Or paddling in your neck with his damned fingers,
Make you to ravel all this matter out
That I essentially am not in madness,
But mad in craft. 'Twere good you let him know,
For who that's but a queen, fair, sober, wise,
Would from a paddock, from a bat, a gib,
Such dear concernings hide? Who would do so?
No, in despite of sense and secrecy,
Unpeg the basket on the house's top,
Let the birds fly, and like the famous ape,
To try conclusions, in the basket creep
And break your own neck down.

QUEEN
Be thou assured, if words be made of breath
And breath of life, I have no life to breathe
What thou hast said to me.

HAMLET
I must to England, you know that.

QUEEN  Alack,
I had forgot! 'Tis so concluded on.

HAMLET
There's letters sealed; and my two schoolfellows,
Whom I will trust as I will adders fanged,
They bear the mandate; they must sweep my way
And marshal me to knavery. Let it work,
For 'tis the sport to have the enginer
Hoist with his own petard; and 't shall go hard
But I will delve one yard below their mines
And blow them at the moon. O, 'tis most sweet
When in one line two crafts directly meet.
This man shall set me packing.
I'll lug the guts into the neighbor room.
Mother, good night indeed. This counselor
Is now most still, most secret, and most grave,
Who was in life a foolish prating knave.--
Come, sir, to draw toward an end with you.--
Good night, mother.
[They exit, Hamlet tugging in Polonius.]


ACT 4
=====

Scene 1
=======
[Enter King and Queen, with Rosencrantz and
Guildenstern.]


KING
There's matter in these sighs; these profound heaves
You must translate; 'tis fit we understand them.
Where is your son?

QUEEN
Bestow this place on us a little while.
[Rosencrantz and Guildenstern exit.]
Ah, mine own lord, what have I seen tonight!

KING  What, Gertrude? How does Hamlet?

QUEEN
Mad as the sea and wind when both contend
Which is the mightier. In his lawless fit,
Behind the arras hearing something stir,
Whips out his rapier, cries "A rat, a rat,"
And in this brainish apprehension kills
The unseen good old man.

KING  O heavy deed!
It had been so with us, had we been there.
His liberty is full of threats to all--
To you yourself, to us, to everyone.
Alas, how shall this bloody deed be answered?
It will be laid to us, whose providence
Should have kept short, restrained, and out of haunt
This mad young man. But so much was our love,
We would not understand what was most fit,
But, like the owner of a foul disease,
To keep it from divulging, let it feed
Even on the pith of life. Where is he gone?

QUEEN
To draw apart the body he hath killed,
O'er whom his very madness, like some ore
Among a mineral of metals base,
Shows itself pure: he weeps for what is done.

KING  O Gertrude, come away!
The sun no sooner shall the mountains touch
But we will ship him hence; and this vile deed
We must with all our majesty and skill
Both countenance and excuse.--Ho, Guildenstern!

[Enter Rosencrantz and Guildenstern.]

Friends both, go join you with some further aid.
Hamlet in madness hath Polonius slain,
And from his mother's closet hath he dragged him.
Go seek him out, speak fair, and bring the body
Into the chapel. I pray you, haste in this.
[Rosencrantz and Guildenstern exit.]
Come, Gertrude, we'll call up our wisest friends
And let them know both what we mean to do
And what's untimely done. ...
Whose whisper o'er the world's diameter,
As level as the cannon to his blank
Transports his poisoned shot, may miss our name
And hit the woundless air. O, come away!
My soul is full of discord and dismay.
[They exit.]

Scene 2
=======
[Enter Hamlet.]


HAMLET  Safely stowed.

GENTLEMEN, [within]  Hamlet! Lord Hamlet!

HAMLET  But soft, what noise? Who calls on Hamlet?
O, here they come.

[Enter Rosencrantz, Guildenstern, and others.]


ROSENCRANTZ
What have you done, my lord, with the dead body?

HAMLET
Compounded it with dust, whereto 'tis kin.

ROSENCRANTZ
Tell us where 'tis, that we may take it thence
And bear it to the chapel.

HAMLET  Do not believe it.

ROSENCRANTZ  Believe what?

HAMLET  That I can keep your counsel and not mine
own. Besides, to be demanded of a sponge, what
replication should be made by the son of a king?

ROSENCRANTZ  Take you me for a sponge, my lord?

HAMLET  Ay, sir, that soaks up the King's countenance,
his rewards, his authorities. But such officers do the
King best service in the end. He keeps them like an
ape an apple in the corner of his jaw, first mouthed,
to be last swallowed. When he needs what you have
gleaned, it is but squeezing you, and, sponge, you
shall be dry again.

ROSENCRANTZ  I understand you not, my lord.

HAMLET  I am glad of it. A knavish speech sleeps in a
foolish ear.

ROSENCRANTZ  My lord, you must tell us where the
body is and go with us to the King.

HAMLET  The body is with the King, but the King is not
with the body. The King is a thing--

GUILDENSTERN  A "thing," my lord?

HAMLET  Of nothing. Bring me to him. Hide fox, and
all after!
[They exit.]

Scene 3
=======
[Enter King and two or three.]


KING
I have sent to seek him and to find the body.
How dangerous is it that this man goes loose!
Yet must not we put the strong law on him.
He's loved of the distracted multitude,
Who like not in their judgment, but their eyes;
And, where 'tis so, th' offender's scourge is weighed,
But never the offense. To bear all smooth and even,
This sudden sending him away must seem
Deliberate pause. Diseases desperate grown
By desperate appliance are relieved
Or not at all.

[Enter Rosencrantz.]

How now, what hath befallen?

ROSENCRANTZ
Where the dead body is bestowed, my lord,
We cannot get from him.

KING  But where is he?

ROSENCRANTZ
Without, my lord; guarded, to know your pleasure.

KING
Bring him before us.

ROSENCRANTZ  Ho! Bring in the lord.

[They enter with Hamlet.]


KING  Now, Hamlet, where's Polonius?

HAMLET  At supper.

KING  At supper where?

HAMLET  Not where he eats, but where he is eaten. A
certain convocation of politic worms are e'en at
him. Your worm is your only emperor for diet. We
fat all creatures else to fat us, and we fat ourselves
for maggots. Your fat king and your lean beggar is
but variable service--two dishes but to one table.
That's the end.

KING  Alas, alas!

HAMLET  A man may fish with the worm that hath eat
of a king and eat of the fish that hath fed of that
worm.

KING  What dost thou mean by this?

HAMLET  Nothing but to show you how a king may go a
progress through the guts of a beggar.

KING  Where is Polonius?

HAMLET  In heaven. Send thither to see. If your messenger
find him not there, seek him i' th' other
place yourself. But if, indeed, you find him not
within this month, you shall nose him as you go up
the stairs into the lobby.

KING, [to Attendants.]  Go, seek him there.

HAMLET  He will stay till you come.	[Attendants exit.]

KING
Hamlet, this deed, for thine especial safety
(Which we do tender, as we dearly grieve
For that which thou hast done) must send thee
hence
With fiery quickness. Therefore prepare thyself.
The bark is ready, and the wind at help,
Th' associates tend, and everything is bent
For England.

HAMLET  For England?

KING  Ay, Hamlet.

HAMLET  Good.

KING
So is it, if thou knew'st our purposes.

HAMLET
I see a cherub that sees them. But come, for
England.
Farewell, dear mother.

KING  Thy loving father, Hamlet.

HAMLET
My mother. Father and mother is man and wife,
Man and wife is one flesh, and so, my mother.--
Come, for England.	[He exits.]

KING
Follow him at foot; tempt him with speed aboard.
Delay it not. I'll have him hence tonight.
Away, for everything is sealed and done
That else leans on th' affair. Pray you, make haste.
[All but the King exit.]
And England, if my love thou hold'st at aught
(As my great power thereof may give thee sense,
Since yet thy cicatrice looks raw and red
After the Danish sword, and thy free awe
Pays homage to us), thou mayst not coldly set
Our sovereign process, which imports at full,
By letters congruing to that effect,
The present death of Hamlet. Do it, England,
For like the hectic in my blood he rages,
And thou must cure me. Till I know 'tis done,
Howe'er my haps, my joys will ne'er begin.
[He exits.]

Scene 4
=======
[Enter Fortinbras with his army over the stage.]


FORTINBRAS
Go, Captain, from me greet the Danish king.
Tell him that by his license Fortinbras
Craves the conveyance of a promised march
Over his kingdom. You know the rendezvous.
If that his Majesty would aught with us,
We shall express our duty in his eye;
And let him know so.

CAPTAIN  I will do 't, my lord.

FORTINBRAS  Go softly on.	[All but the Captain exit.]

[Enter Hamlet, Rosencrantz, Guildenstern, and others.]


HAMLET  Good sir, whose powers are these?

CAPTAIN  They are of Norway, sir.

HAMLET  How purposed, sir, I pray you?

CAPTAIN  Against some part of Poland.

HAMLET  Who commands them, sir?

CAPTAIN
The nephew to old Norway, Fortinbras.

HAMLET
Goes it against the main of Poland, sir,
Or for some frontier?

CAPTAIN
Truly to speak, and with no addition,
We go to gain a little patch of ground
That hath in it no profit but the name.
To pay five ducats, five, I would not farm it;
Nor will it yield to Norway or the Pole
A ranker rate, should it be sold in fee.

HAMLET
Why, then, the Polack never will defend it.

CAPTAIN
Yes, it is already garrisoned.

HAMLET
Two thousand souls and twenty thousand ducats
Will not debate the question of this straw.
This is th' impostume of much wealth and peace,
That inward breaks and shows no cause without
Why the man dies.--I humbly thank you, sir.

CAPTAIN  God be wi' you, sir.	[He exits.]

ROSENCRANTZ  Will 't please you go, my lord?

HAMLET
I'll be with you straight. Go a little before.
[All but Hamlet exit.]
How all occasions do inform against me
And spur my dull revenge. What is a man
If his chief good and market of his time
Be but to sleep and feed? A beast, no more.
Sure He that made us with such large discourse,
Looking before and after, gave us not
That capability and godlike reason
To fust in us unused. Now whether it be
Bestial oblivion or some craven scruple
Of thinking too precisely on th' event
(A thought which, quartered, hath but one part
wisdom
And ever three parts coward), I do not know
Why yet I live to say "This thing's to do,"
Sith I have cause, and will, and strength, and means
To do 't. Examples gross as Earth exhort me:
Witness this army of such mass and charge,
Led by a delicate and tender prince,
Whose spirit with divine ambition puffed
Makes mouths at the invisible event,
Exposing what is mortal and unsure
To all that fortune, death, and danger dare,
Even for an eggshell. Rightly to be great
Is not to stir without great argument,
But greatly to find quarrel in a straw
When honor's at the stake. How stand I, then,
That have a father killed, a mother stained,
Excitements of my reason and my blood,
And let all sleep, while to my shame I see
The imminent death of twenty thousand men
That for a fantasy and trick of fame
Go to their graves like beds, fight for a plot
Whereon the numbers cannot try the cause,
Which is not tomb enough and continent
To hide the slain? O, from this time forth
My thoughts be bloody or be nothing worth!
[He exits.]

Scene 5
=======
[Enter Horatio, Queen, and a Gentleman.]


QUEEN  I will not speak with her.

GENTLEMAN  She is importunate,
Indeed distract; her mood will needs be pitied.

QUEEN  What would she have?

GENTLEMAN
She speaks much of her father, says she hears
There's tricks i' th' world, and hems, and beats her
heart,
Spurns enviously at straws, speaks things in doubt
That carry but half sense. Her speech is nothing,
Yet the unshaped use of it doth move
The hearers to collection. They aim at it
And botch the words up fit to their own thoughts;
Which, as her winks and nods and gestures yield
them,
Indeed would make one think there might be
thought,
Though nothing sure, yet much unhappily.

HORATIO
'Twere good she were spoken with, for she may
strew
Dangerous conjectures in ill-breeding minds.

QUEEN  Let her come in.	[Gentleman exits.]
[Aside.] To my sick soul (as sin's true nature is),
Each toy seems prologue to some great amiss.
So full of artless jealousy is guilt,
It spills itself in fearing to be spilt.

[Enter Ophelia distracted.]


OPHELIA
Where is the beauteous Majesty of Denmark?

QUEEN  How now, Ophelia?

OPHELIA [sings]
	How should I your true love know
	   From another one?
	By his cockle hat and staff
	   And his sandal shoon.

QUEEN
Alas, sweet lady, what imports this song?

OPHELIA  Say you? Nay, pray you, mark.
[Sings.]	He is dead and gone, lady,
	   He is dead and gone;
	At his head a grass-green turf,
	   At his heels a stone.
Oh, ho!

QUEEN  Nay, but Ophelia--

OPHELIA  Pray you, mark.
[Sings.]	White his shroud as the mountain snow--

[Enter King.]


QUEEN  Alas, look here, my lord.

OPHELIA [sings]
	   Larded all with sweet flowers;
	Which bewept to the ground did not go
	   With true-love showers.

KING  How do you, pretty lady?

OPHELIA  Well, God dild you. They say the owl was a
baker's daughter. Lord, we know what we are but
know not what we may be. God be at your table.

KING  Conceit upon her father.

OPHELIA  Pray let's have no words of this, but when
they ask you what it means, say you this:
[Sings.]	Tomorrow is Saint Valentine's day,
	   All in the morning betime,
	And I a maid at your window,
	   To be your Valentine.
	Then up he rose and donned his clothes
	   And dupped the chamber door,
	Let in the maid, that out a maid
	   Never departed more.

KING  Pretty Ophelia--

OPHELIA
Indeed, without an oath, I'll make an end on 't:
[Sings.]	By Gis and by Saint Charity,
	   Alack and fie for shame,
	Young men will do 't, if they come to 't;
	   By Cock, they are to blame.
	Quoth she "Before you tumbled me,
	   You promised me to wed."
He answers:
	"So would I 'a done, by yonder sun,
	   An thou hadst not come to my bed."

KING  How long hath she been thus?

OPHELIA  I hope all will be well. We must be patient,
but I cannot choose but weep to think they would
lay him i' th' cold ground. My brother shall know of
it. And so I thank you for your good counsel. Come,
my coach! Good night, ladies, good night, sweet
ladies, good night, good night.	[She exits.]

KING
Follow her close; give her good watch, I pray you.
[Horatio exits.]
O, this is the poison of deep grief. It springs
All from her father's death, and now behold!
O Gertrude, Gertrude,
When sorrows come, they come not single spies,
But in battalions: first, her father slain;
Next, your son gone, and he most violent author
Of his own just remove; the people muddied,
Thick, and unwholesome in their thoughts and
whispers
For good Polonius' death, and we have done but
greenly
In hugger-mugger to inter him; poor Ophelia
Divided from herself and her fair judgment,
Without the which we are pictures or mere beasts;
Last, and as much containing as all these,
Her brother is in secret come from France,
Feeds on his wonder, keeps himself in clouds,
And wants not buzzers to infect his ear
With pestilent speeches of his father's death,
Wherein necessity, of matter beggared,
Will nothing stick our person to arraign
In ear and ear. O, my dear Gertrude, this,
Like to a murd'ring piece, in many places
Gives me superfluous death.
[A noise within.]

QUEEN  Alack, what noise is this?

KING  Attend!
Where is my Switzers? Let them guard the door.

[Enter a Messenger.]

What is the matter?

MESSENGER  Save yourself, my lord.
The ocean, overpeering of his list,
Eats not the flats with more impiteous haste
Than young Laertes, in a riotous head,
O'erbears your officers. The rabble call him "lord,"
And, as the world were now but to begin,
Antiquity forgot, custom not known,
The ratifiers and props of every word,
They cry "Choose we, Laertes shall be king!"
Caps, hands, and tongues applaud it to the clouds,
"Laertes shall be king! Laertes king!"
[A noise within.]

QUEEN
How cheerfully on the false trail they cry.
O, this is counter, you false Danish dogs!

KING  The doors are broke.

[Enter Laertes with others.
]

LAERTES
Where is this king?--Sirs, stand you all without.

ALL  No, let's come in!

LAERTES  I pray you, give me leave.

ALL  We will, we will.

LAERTES
I thank you. Keep the door. [Followers exit.] O, thou
vile king,
Give me my father!

QUEEN  Calmly, good Laertes.

LAERTES
That drop of blood that's calm proclaims me
bastard,
Cries "cuckold" to my father, brands the harlot
Even here between the chaste unsmirched brow
Of my true mother.

KING  What is the cause, Laertes,
That thy rebellion looks so giant-like?--
Let him go, Gertrude. Do not fear our person.
There's such divinity doth hedge a king
That treason can but peep to what it would,
Acts little of his will.--Tell me, Laertes,
Why thou art thus incensed.--Let him go,
Gertrude.--
Speak, man.

LAERTES  Where is my father?

KING  Dead.

QUEEN
But not by him.

KING  Let him demand his fill.

LAERTES
How came he dead? I'll not be juggled with.
To hell, allegiance! Vows, to the blackest devil!
Conscience and grace, to the profoundest pit!
I dare damnation. To this point I stand,
That both the worlds I give to negligence,
Let come what comes, only I'll be revenged
Most throughly for my father.

KING  Who shall stay you?

LAERTES  My will, not all the world.
And for my means, I'll husband them so well
They shall go far with little.

KING  Good Laertes,
If you desire to know the certainty
Of your dear father, is 't writ in your revenge
That, swoopstake, you will draw both friend and
foe,
Winner and loser?

LAERTES  None but his enemies.

KING  Will you know them, then?

LAERTES
To his good friends thus wide I'll ope my arms
And, like the kind life-rend'ring pelican,
Repast them with my blood.

KING  Why, now you speak
Like a good child and a true gentleman.
That I am guiltless of your father's death
And am most sensibly in grief for it,
It shall as level to your judgment 'pear
As day does to your eye.

  [A noise within:] "Let her come in!"

LAERTES  How now, what noise is that?

[Enter Ophelia.]

O heat, dry up my brains! Tears seven times salt
Burn out the sense and virtue of mine eye!
By heaven, thy madness shall be paid with weight
Till our scale turn the beam! O rose of May,
Dear maid, kind sister, sweet Ophelia!
O heavens, is 't possible a young maid's wits
Should be as mortal as an old man's life?
Nature is fine in love, and, where 'tis fine,
It sends some precious instance of itself
After the thing it loves.

OPHELIA [sings]
	They bore him barefaced on the bier,
	Hey non nonny, nonny, hey nonny,
	And in his grave rained many a tear.
Fare you well, my dove.

LAERTES
Hadst thou thy wits and didst persuade revenge,
It could not move thus.

OPHELIA  You must sing "A-down a-down"--and you
"Call him a-down-a."--O, how the wheel becomes
it! It is the false steward that stole his master's
daughter.

LAERTES  This nothing's more than matter.

OPHELIA  There's rosemary, that's for remembrance.
Pray you, love, remember. And there is pansies,
that's for thoughts.

LAERTES  A document in madness: thoughts and remembrance
fitted.

OPHELIA  There's fennel for you, and columbines.
There's rue for you, and here's some for me; we
may call it herb of grace o' Sundays. You must wear
your rue with a difference. There's a daisy. I would
give you some violets, but they withered all when
my father died. They say he made a good end.
[Sings.] For bonny sweet Robin is all my joy.

LAERTES
Thought and afflictions, passion, hell itself
She turns to favor and to prettiness.

OPHELIA [sings]
	And will he not come again?
	And will he not come again?
	   No, no, he is dead.
	   Go to thy deathbed.
	He never will come again.

	His beard was as white as snow,
	All flaxen was his poll.
	   He is gone, he is gone,
	   And we cast away moan.
	God 'a mercy on his soul.
And of all Christians' souls, I pray God. God be wi'
you.	[She exits.]

LAERTES  Do you see this, O God?

KING
Laertes, I must commune with your grief,
Or you deny me right. Go but apart,
Make choice of whom your wisest friends you will,
And they shall hear and judge 'twixt you and me.
If by direct or by collateral hand
They find us touched, we will our kingdom give,
Our crown, our life, and all that we call ours,
To you in satisfaction; but if not,
Be you content to lend your patience to us,
And we shall jointly labor with your soul
To give it due content.

LAERTES  Let this be so.
His means of death, his obscure funeral
(No trophy, sword, nor hatchment o'er his bones,
No noble rite nor formal ostentation)
Cry to be heard, as 'twere from heaven to earth,
That I must call 't in question.

KING  So you shall,
And where th' offense is, let the great ax fall.
I pray you, go with me.
[They exit.]

Scene 6
=======
[Enter Horatio and others.]


HORATIO  What are they that would speak with me?

GENTLEMAN  Seafaring men, sir. They say they have
letters for you.

HORATIO  Let them come in. [Gentleman exits.] I do not
know from what part of the world I should be
greeted, if not from Lord Hamlet.

[Enter Sailors.]


SAILOR  God bless you, sir.

HORATIO  Let Him bless thee too.

SAILOR  He shall, sir, an 't please Him. There's a letter
for you, sir. It came from th' ambassador that was
bound for England--if your name be Horatio, as I
am let to know it is.	[He hands Horatio a letter.]

HORATIO [reads the letter]  Horatio, when thou shalt have
overlooked this, give these fellows some means to the
King. They have letters for him. Ere we were two days
old at sea, a pirate of very warlike appointment gave
us chase. Finding ourselves too slow of sail, we put on
a compelled valor, and in the grapple I boarded them.
On the instant, they got clear of our ship; so I alone
became their prisoner. They have dealt with me like
thieves of mercy, but they knew what they did: I am to
do a good turn for them. Let the King have the letters
I have sent, and repair thou to me with as much speed
as thou wouldst fly death. I have words to speak in
thine ear will make thee dumb; yet are they much too
light for the bore of the matter. These good fellows
will bring thee where I am. Rosencrantz and Guildenstern
hold their course for England; of them I have
much to tell thee. Farewell.
He that thou knowest thine,
Hamlet.
Come, I will give you way for these your letters
And do 't the speedier that you may direct me
To him from whom you brought them.
[They exit.]

Scene 7
=======
[Enter King and Laertes.]


KING
Now must your conscience my acquittance seal,
And you must put me in your heart for friend,
Sith you have heard, and with a knowing ear,
That he which hath your noble father slain
Pursued my life.

LAERTES  It well appears. But tell me
Why you proceeded not against these feats,
So criminal and so capital in nature,
As by your safety, greatness, wisdom, all things else,
You mainly were stirred up.

KING  O, for two special reasons,
Which may to you perhaps seem much unsinewed,
But yet to me they're strong. The Queen his mother
Lives almost by his looks, and for myself
(My virtue or my plague, be it either which),
She is so conjunctive to my life and soul
That, as the star moves not but in his sphere,
I could not but by her. The other motive
Why to a public count I might not go
Is the great love the general gender bear him,
Who, dipping all his faults in their affection,
Work like the spring that turneth wood to stone,
Convert his gyves to graces, so that my arrows,
Too slightly timbered for so loud a wind,
Would have reverted to my bow again,
But not where I have aimed them.

LAERTES
And so have I a noble father lost,
A sister driven into desp'rate terms,
Whose worth, if praises may go back again,
Stood challenger on mount of all the age
For her perfections. But my revenge will come.

KING
Break not your sleeps for that. You must not think
That we are made of stuff so flat and dull
That we can let our beard be shook with danger
And think it pastime. You shortly shall hear more.
I loved your father, and we love ourself,
And that, I hope, will teach you to imagine--

[Enter a Messenger with letters.]

How now? What news?

MESSENGER  Letters, my lord, from
Hamlet.
These to your Majesty, this to the Queen.

KING  From Hamlet? Who brought them?

MESSENGER
Sailors, my lord, they say. I saw them not.
They were given me by Claudio. He received them
Of him that brought them.

KING  Laertes, you shall hear
them.--
Leave us.	[Messenger exits.]
[Reads.] High and mighty, you shall know I am set
naked on your kingdom. Tomorrow shall I beg leave to
see your kingly eyes, when I shall (first asking your
pardon) thereunto recount the occasion of my sudden
and more strange return. Hamlet.
What should this mean? Are all the rest come back?
Or is it some abuse and no such thing?

LAERTES  Know you the hand?

KING  'Tis Hamlet's character. "Naked"--
And in a postscript here, he says "alone."
Can you advise me?

LAERTES
I am lost in it, my lord. But let him come.
It warms the very sickness in my heart
That I shall live and tell him to his teeth
"Thus didst thou."

KING  If it be so, Laertes
(As how should it be so? how otherwise?),
Will you be ruled by me?

LAERTES  Ay, my lord,
So you will not o'errule me to a peace.

KING
To thine own peace. If he be now returned,
As checking at his voyage, and that he means
No more to undertake it, I will work him
To an exploit, now ripe in my device,
Under the which he shall not choose but fall;
And for his death no wind of blame shall breathe,
But even his mother shall uncharge the practice
And call it accident.

LAERTES  My lord, I will be ruled,
The rather if you could devise it so
That I might be the organ.

KING  It falls right.
You have been talked of since your travel much,
And that in Hamlet's hearing, for a quality
Wherein they say you shine. Your sum of parts
Did not together pluck such envy from him
As did that one, and that, in my regard,
Of the unworthiest siege.

LAERTES  What part is that, my lord?

KING
A very ribbon in the cap of youth--
Yet needful too, for youth no less becomes
The light and careless livery that it wears
Than settled age his sables and his weeds,
Importing health and graveness. Two months since
Here was a gentleman of Normandy.
I have seen myself, and served against, the French,
And they can well on horseback, but this gallant
Had witchcraft in 't. He grew unto his seat,
And to such wondrous doing brought his horse
As had he been encorpsed and demi-natured
With the brave beast. So far he topped my thought
That I in forgery of shapes and tricks
Come short of what he did.

LAERTES  A Norman was 't?

KING  A Norman.

LAERTES
Upon my life, Lamord.

KING  The very same.

LAERTES
I know him well. He is the brooch indeed
And gem of all the nation.

KING  He made confession of you
And gave you such a masterly report
For art and exercise in your defense,
And for your rapier most especial,
That he cried out 'twould be a sight indeed
If one could match you. The 'scrimers of their
nation
He swore had neither motion, guard, nor eye,
If you opposed them. Sir, this report of his
Did Hamlet so envenom with his envy
That he could nothing do but wish and beg
Your sudden coming-o'er, to play with you.
Now out of this--

LAERTES  What out of this, my lord?

KING
Laertes, was your father dear to you?
Or are you like the painting of a sorrow,
A face without a heart?

LAERTES  Why ask you this?

KING
Not that I think you did not love your father,
But that I know love is begun by time
And that I see, in passages of proof,
Time qualifies the spark and fire of it.
There lives within the very flame of love
A kind of wick or snuff that will abate it,
And nothing is at a like goodness still;
For goodness, growing to a pleurisy,
Dies in his own too-much. That we would do
We should do when we would; for this "would"
changes
And hath abatements and delays as many
As there are tongues, are hands, are accidents;
And then this "should" is like a spendthrift sigh,
That hurts by easing. But to the quick of th' ulcer:
Hamlet comes back; what would you undertake
To show yourself indeed your father's son
More than in words?

LAERTES  To cut his throat i' th' church.

KING
No place indeed should murder sanctuarize;
Revenge should have no bounds. But, good Laertes,
Will you do this? Keep close within your chamber.
Hamlet, returned, shall know you are come home.
We'll put on those shall praise your excellence
And set a double varnish on the fame
The Frenchman gave you; bring you, in fine,
together
And wager on your heads. He, being remiss,
Most generous, and free from all contriving,
Will not peruse the foils, so that with ease,
Or with a little shuffling, you may choose
A sword unbated, and in a pass of practice
Requite him for your father.

LAERTES  I will do 't,
And for that purpose I'll anoint my sword.
I bought an unction of a mountebank
So mortal that, but dip a knife in it,
Where it draws blood no cataplasm so rare,
Collected from all simples that have virtue
Under the moon, can save the thing from death
That is but scratched withal. I'll touch my point
With this contagion, that, if I gall him slightly,
It may be death.

KING  Let's further think of this,
Weigh what convenience both of time and means
May fit us to our shape. If this should fail,
And that our drift look through our bad
performance,
'Twere better not assayed. Therefore this project
Should have a back or second that might hold
If this did blast in proof. Soft, let me see.
We'll make a solemn wager on your cunnings--
I ha 't!
When in your motion you are hot and dry
(As make your bouts more violent to that end)
And that he calls for drink, I'll have prepared
him
A chalice for the nonce, whereon but sipping,
If he by chance escape your venomed stuck,
Our purpose may hold there.--But stay, what
noise?

[Enter Queen.]


QUEEN
One woe doth tread upon another's heel,
So fast they follow. Your sister's drowned, Laertes.

LAERTES  Drowned? O, where?

QUEEN
There is a willow grows askant the brook
That shows his hoar leaves in the glassy stream.
Therewith fantastic garlands did she make
Of crowflowers, nettles, daisies, and long purples,
That liberal shepherds give a grosser name,
But our cold maids do "dead men's fingers" call
them.
There on the pendant boughs her coronet weeds
Clamb'ring to hang, an envious sliver broke,
When down her weedy trophies and herself
Fell in the weeping brook. Her clothes spread wide,
And mermaid-like awhile they bore her up,
Which time she chanted snatches of old lauds,
As one incapable of her own distress
Or like a creature native and endued
Unto that element. But long it could not be
Till that her garments, heavy with their drink,
Pulled the poor wretch from her melodious lay
To muddy death.

LAERTES  Alas, then she is drowned.

QUEEN  Drowned, drowned.

LAERTES
Too much of water hast thou, poor Ophelia,
And therefore I forbid my tears. But yet
It is our trick; nature her custom holds,
Let shame say what it will. When these are gone,
The woman will be out.--Adieu, my lord.
I have a speech o' fire that fain would blaze,
But that this folly drowns it.	[He exits.]

KING  Let's follow, Gertrude.
How much I had to do to calm his rage!
Now fear I this will give it start again.
Therefore, let's follow.
[They exit.]


ACT 5
=====

Scene 1
=======
[Enter Gravedigger and Another.]


GRAVEDIGGER  Is she to be buried in Christian burial,
when she willfully seeks her own salvation?

OTHER  I tell thee she is. Therefore make her grave
straight. The crowner hath sat on her and finds it
Christian burial.

GRAVEDIGGER  How can that be, unless she drowned
herself in her own defense?

OTHER  Why, 'tis found so.

GRAVEDIGGER  It must be se offendendo; it cannot be
else. For here lies the point: if I drown myself
wittingly, it argues an act, and an act hath three
branches--it is to act, to do, to perform. Argal, she
drowned herself wittingly.

OTHER  Nay, but hear you, goodman delver--

GRAVEDIGGER  Give me leave. Here lies the water;
good. Here stands the man; good. If the man go to
this water and drown himself, it is (will he, nill he)
he goes; mark you that. But if the water come to him
and drown him, he drowns not himself. Argal, he
that is not guilty of his own death shortens not his
own life.

OTHER  But is this law?

GRAVEDIGGER  Ay, marry, is 't--crowner's 'quest law.

OTHER  Will you ha' the truth on 't? If this had not been
a gentlewoman, she should have been buried out o'
Christian burial.

GRAVEDIGGER  Why, there thou sayst. And the more
pity that great folk should have count'nance in this
world to drown or hang themselves more than
their even-Christian. Come, my spade. There is no
ancient gentlemen but gard'ners, ditchers, and
grave-makers. They hold up Adam's profession.

OTHER  Was he a gentleman?

GRAVEDIGGER  He was the first that ever bore arms.

OTHER  Why, he had none.

GRAVEDIGGER  What, art a heathen? How dost thou
understand the scripture? The scripture says Adam
digged. Could he dig without arms? I'll put another
question to thee. If thou answerest me not to the
purpose, confess thyself--

OTHER  Go to!

GRAVEDIGGER  What is he that builds stronger than
either the mason, the shipwright, or the carpenter?

OTHER  The gallows-maker; for that frame outlives a
thousand tenants.

GRAVEDIGGER  I like thy wit well, in good faith. The
gallows does well. But how does it well? It does
well to those that do ill. Now, thou dost ill to say the
gallows is built stronger than the church. Argal, the
gallows may do well to thee. To 't again, come.

OTHER  "Who builds stronger than a mason, a shipwright,
or a carpenter?"

GRAVEDIGGER  Ay, tell me that, and unyoke.

OTHER  Marry, now I can tell.

GRAVEDIGGER  To 't.

OTHER  Mass, I cannot tell.

[Enter Hamlet and Horatio afar off.]


GRAVEDIGGER  Cudgel thy brains no more about it,
for your dull ass will not mend his pace with
beating. And, when you are asked this question
next, say "a grave-maker." The houses he makes
lasts till doomsday. Go, get thee in, and fetch me a
stoup of liquor.
[The Other Man exits
and the Gravedigger digs and sings.]
	In youth when I did love, did love,
	   Methought it was very sweet
	To contract--O--the time for--a--my behove,
	   O, methought there--a--was nothing--a--meet.

HAMLET  Has this fellow no feeling of his business? He
sings in grave-making.

HORATIO  Custom hath made it in him a property of
easiness.

HAMLET  'Tis e'en so. The hand of little employment
hath the daintier sense.

GRAVEDIGGER [sings]
	But age with his stealing steps
	Hath clawed me in his clutch,
	And hath shipped me into the land,
	As if I had never been such.
[He digs up a skull.]

HAMLET  That skull had a tongue in it and could sing
once. How the knave jowls it to the ground as if
'twere Cain's jawbone, that did the first murder!
This might be the pate of a politician which this ass
now o'erreaches, one that would circumvent God,
might it not?

HORATIO  It might, my lord.

HAMLET  Or of a courtier, which could say "Good
morrow, sweet lord! How dost thou, sweet lord?"
This might be my Lord Such-a-one that praised my
Lord Such-a-one's horse when he went to beg it,
might it not?

HORATIO  Ay, my lord.

HAMLET  Why, e'en so. And now my Lady Worm's,
chapless and knocked about the mazard with a
sexton's spade. Here's fine revolution, an we had
the trick to see 't. Did these bones cost no more the
breeding but to play at loggets with them? Mine
ache to think on 't.

GRAVEDIGGER [sings]
	A pickax and a spade, a spade,
	For and a shrouding sheet,
	O, a pit of clay for to be made
	For such a guest is meet.
[He digs up more skulls.]

HAMLET  There's another. Why may not that be the
skull of a lawyer? Where be his quiddities now, his
quillities, his cases, his tenures, and his tricks? Why
does he suffer this mad knave now to knock him
about the sconce with a dirty shovel and will not tell
him of his action of battery? Hum, this fellow might
be in 's time a great buyer of land, with his statutes,
his recognizances, his fines, his double vouchers,
his recoveries. Is this the fine of his fines and the
recovery of his recoveries, to have his fine pate full
of fine dirt? Will his vouchers vouch him no more
of his purchases, and double ones too, than the
length and breadth of a pair of indentures? The very
conveyances of his lands will scarcely lie in this box,
and must th' inheritor himself have no more, ha?

HORATIO  Not a jot more, my lord.

HAMLET  Is not parchment made of sheepskins?

HORATIO  Ay, my lord, and of calves' skins too.

HAMLET  They are sheep and calves which seek out
assurance in that. I will speak to this fellow.--
Whose grave's this, sirrah?

GRAVEDIGGER  Mine, sir.
[Sings.]	O, a pit of clay for to be made
	For such a guest is meet.

HAMLET  I think it be thine indeed, for thou liest in 't.

GRAVEDIGGER  You lie out on 't, sir, and therefore 'tis
not yours. For my part, I do not lie in 't, yet it is
mine.

HAMLET  Thou dost lie in 't, to be in 't and say it is thine.
'Tis for the dead, not for the quick; therefore thou
liest.

GRAVEDIGGER  'Tis a quick lie, sir; 'twill away again
from me to you.

HAMLET  What man dost thou dig it for?

GRAVEDIGGER  For no man, sir.

HAMLET  What woman then?

GRAVEDIGGER  For none, neither.

HAMLET  Who is to be buried in 't?

GRAVEDIGGER  One that was a woman, sir, but, rest
her soul, she's dead.

HAMLET  How absolute the knave is! We must speak by
the card, or equivocation will undo us. By the
Lord, Horatio, this three years I have took note of
it: the age is grown so picked that the toe of the
peasant comes so near the heel of the courtier, he
galls his kibe.--How long hast thou been
grave-maker?

GRAVEDIGGER  Of all the days i' th' year, I came to 't
that day that our last King Hamlet overcame
Fortinbras.

HAMLET  How long is that since?

GRAVEDIGGER  Cannot you tell that? Every fool can
tell that. It was that very day that young Hamlet
was born--he that is mad, and sent into England.

HAMLET  Ay, marry, why was he sent into England?

GRAVEDIGGER  Why, because he was mad. He shall
recover his wits there. Or if he do not, 'tis no great
matter there.

HAMLET  Why?

GRAVEDIGGER  'Twill not be seen in him there. There
the men are as mad as he.

HAMLET  How came he mad?

GRAVEDIGGER  Very strangely, they say.

HAMLET  How "strangely"?

GRAVEDIGGER  Faith, e'en with losing his wits.

HAMLET  Upon what ground?

GRAVEDIGGER  Why, here in Denmark. I have been
sexton here, man and boy, thirty years.

HAMLET  How long will a man lie i' th' earth ere he rot?

GRAVEDIGGER  Faith, if he be not rotten before he die
(as we have many pocky corses nowadays that will
scarce hold the laying in), he will last you some
eight year or nine year. A tanner will last you nine
year.

HAMLET  Why he more than another?

GRAVEDIGGER  Why, sir, his hide is so tanned with his
trade that he will keep out water a great while; and
your water is a sore decayer of your whoreson dead
body. Here's a skull now hath lien you i' th' earth
three-and-twenty years.

HAMLET  Whose was it?

GRAVEDIGGER  A whoreson mad fellow's it was.
Whose do you think it was?

HAMLET  Nay, I know not.

GRAVEDIGGER  A pestilence on him for a mad rogue!
He poured a flagon of Rhenish on my head once.
This same skull, sir, was, sir, Yorick's skull, the
King's jester.

HAMLET  This?

GRAVEDIGGER  E'en that.

HAMLET, [taking the skull]  Let me see. Alas, poor
Yorick! I knew him, Horatio--a fellow of infinite
jest, of most excellent fancy. He hath bore me on his
back a thousand times, and now how abhorred in
my imagination it is! My gorge rises at it. Here hung
those lips that I have kissed I know not how oft.
Where be your gibes now? your gambols? your
songs? your flashes of merriment that were wont to
set the table on a roar? Not one now to mock your
own grinning? Quite chapfallen? Now get you to my
lady's chamber, and tell her, let her paint an inch
thick, to this favor she must come. Make her laugh
at that.--Prithee, Horatio, tell me one thing.

HORATIO  What's that, my lord?

HAMLET  Dost thou think Alexander looked o' this
fashion i' th' earth?

HORATIO  E'en so.

HAMLET  And smelt so? Pah!	[He puts the skull down.]

HORATIO  E'en so, my lord.

HAMLET  To what base uses we may return, Horatio!
Why may not imagination trace the noble dust of
Alexander till he find it stopping a bunghole?

HORATIO  'Twere to consider too curiously to consider
so.

HAMLET  No, faith, not a jot; but to follow him thither,
with modesty enough and likelihood to lead it, as
thus: Alexander died, Alexander was buried, Alexander
returneth to dust; the dust is earth; of earth
we make loam; and why of that loam whereto he
was converted might they not stop a beer barrel?
Imperious Caesar, dead and turned to clay,
Might stop a hole to keep the wind away.
O, that that earth which kept the world in awe
Should patch a wall t' expel the winter's flaw!

[Enter King, Queen, Laertes, Lords attendant, and the
corpse of Ophelia, with a Doctor of Divinity.]

But soft, but soft awhile! Here comes the King,
The Queen, the courtiers. Who is this they follow?
And with such maimed rites? This doth betoken
The corse they follow did with desp'rate hand
Fordo its own life. 'Twas of some estate.
Couch we awhile and mark.	[They step aside.]

LAERTES  What ceremony else?

HAMLET  That is Laertes, a very noble youth. Mark.

LAERTES  What ceremony else?

DOCTOR
Her obsequies have been as far enlarged
As we have warranty. Her death was doubtful,
And, but that great command o'ersways the order,
She should in ground unsanctified been lodged
Till the last trumpet. For charitable prayers
Shards, flints, and pebbles should be thrown on
her.
Yet here she is allowed her virgin crants,
Her maiden strewments, and the bringing home
Of bell and burial.

LAERTES
Must there no more be done?

DOCTOR  No more be done.
We should profane the service of the dead
To sing a requiem and such rest to her
As to peace-parted souls.

LAERTES  Lay her i' th' earth,
And from her fair and unpolluted flesh
May violets spring! I tell thee, churlish priest,
A minist'ring angel shall my sister be
When thou liest howling.

HAMLET, [to Horatio]  What, the fair Ophelia?

QUEEN  Sweets to the sweet, farewell!
[She scatters flowers.]
I hoped thou shouldst have been my Hamlet's wife;
I thought thy bride-bed to have decked, sweet maid,
And not have strewed thy grave.

LAERTES  O, treble woe
Fall ten times treble on that cursed head
Whose wicked deed thy most ingenious sense
Deprived thee of!--Hold off the earth awhile,
Till I have caught her once more in mine arms.
[Leaps in the grave.]
Now pile your dust upon the quick and dead,
Till of this flat a mountain you have made
T' o'ertop old Pelion or the skyish head
Of blue Olympus.

HAMLET, [advancing]
What is he whose grief
Bears such an emphasis, whose phrase of sorrow
Conjures the wand'ring stars and makes them stand
Like wonder-wounded hearers? This is I,
Hamlet the Dane.

LAERTES, [coming out of the grave]
The devil take thy soul!

HAMLET  Thou pray'st not well.	[They grapple.]
I prithee take thy fingers from my throat,
For though I am not splenitive and rash,
Yet have I in me something dangerous,
Which let thy wisdom fear. Hold off thy hand.

KING  Pluck them asunder.

QUEEN  Hamlet! Hamlet!

ALL  Gentlemen!

HORATIO  Good my lord, be quiet.
[Hamlet and Laertes are separated.]

HAMLET
Why, I will fight with him upon this theme
Until my eyelids will no longer wag!

QUEEN  O my son, what theme?

HAMLET
I loved Ophelia. Forty thousand brothers
Could not with all their quantity of love
Make up my sum. What wilt thou do for her?

KING  O, he is mad, Laertes!

QUEEN  For love of God, forbear him.

HAMLET  'Swounds, show me what thou 't do.
Woo't weep, woo't fight, woo't fast, woo't tear
thyself,
Woo't drink up eisel, eat a crocodile?
I'll do 't. Dost thou come here to whine?
To outface me with leaping in her grave?
Be buried quick with her, and so will I.
And if thou prate of mountains, let them throw
Millions of acres on us, till our ground,
Singeing his pate against the burning zone,
Make Ossa like a wart. Nay, an thou 'lt mouth,
I'll rant as well as thou.

QUEEN  This is mere madness;
And thus awhile the fit will work on him.
Anon, as patient as the female dove
When that her golden couplets are disclosed,
His silence will sit drooping.

HAMLET  Hear you, sir,
What is the reason that you use me thus?
I loved you ever. But it is no matter.
Let Hercules himself do what he may,
The cat will mew, and dog will have his day.
[Hamlet exits.]

KING
I pray thee, good Horatio, wait upon him.
[Horatio exits.]
[To Laertes.] Strengthen your patience in our last
night's speech.
We'll put the matter to the present push.--
Good Gertrude, set some watch over your son.--
This grave shall have a living monument.
An hour of quiet thereby shall we see.
Till then in patience our proceeding be.
[They exit.]

Scene 2
=======
[Enter Hamlet and Horatio.]


HAMLET
So much for this, sir. Now shall you see the other.
You do remember all the circumstance?

HORATIO  Remember it, my lord!

HAMLET
Sir, in my heart there was a kind of fighting
That would not let me sleep. Methought I lay
Worse than the mutines in the bilboes. Rashly--
And praised be rashness for it; let us know,
Our indiscretion sometime serves us well
When our deep plots do pall; and that should learn
us
There's a divinity that shapes our ends,
Rough-hew them how we will--

HORATIO  That is most
certain.

HAMLET  Up from my cabin,
My sea-gown scarfed about me, in the dark
Groped I to find out them; had my desire,
Fingered their packet, and in fine withdrew
To mine own room again, making so bold
(My fears forgetting manners) to unfold
Their grand commission; where I found, Horatio,
A royal knavery--an exact command,
Larded with many several sorts of reasons
Importing Denmark's health and England's too,
With--ho!--such bugs and goblins in my life,
That on the supervise, no leisure bated,
No, not to stay the grinding of the ax,
My head should be struck off.

HORATIO  Is 't possible?

HAMLET
Here's the commission. Read it at more leisure.
[Handing him a paper.]
But wilt thou hear now how I did proceed?

HORATIO  I beseech you.

HAMLET
Being thus benetted round with villainies,
Or I could make a prologue to my brains,
They had begun the play. I sat me down,
Devised a new commission, wrote it fair--
I once did hold it, as our statists do,
A baseness to write fair, and labored much
How to forget that learning; but, sir, now
It did me yeoman's service. Wilt thou know
Th' effect of what I wrote?

HORATIO  Ay, good my lord.

HAMLET
An earnest conjuration from the King,
As England was his faithful tributary,
As love between them like the palm might flourish,
As peace should still her wheaten garland wear
And stand a comma 'tween their amities,
And many suchlike ases of great charge,
That, on the view and knowing of these contents,
Without debatement further, more or less,
He should those bearers put to sudden death,
Not shriving time allowed.

HORATIO  How was this sealed?

HAMLET
Why, even in that was heaven ordinant.
I had my father's signet in my purse,
Which was the model of that Danish seal;
Folded the writ up in the form of th' other,
Subscribed it, gave 't th' impression, placed it
safely,
The changeling never known. Now, the next day
Was our sea-fight; and what to this was sequent
Thou knowest already.

HORATIO
So Guildenstern and Rosencrantz go to 't.

HAMLET
Why, man, they did make love to this employment.
They are not near my conscience. Their defeat
Does by their own insinuation grow.
'Tis dangerous when the baser nature comes
Between the pass and fell incensed points
Of mighty opposites.

HORATIO  Why, what a king is this!

HAMLET
Does it not, think thee, stand me now upon--
He that hath killed my king and whored my mother,
Popped in between th' election and my hopes,
Thrown out his angle for my proper life,
And with such cozenage--is 't not perfect
conscience
To quit him with this arm? And is 't not to be
damned
To let this canker of our nature come
In further evil?

HORATIO
It must be shortly known to him from England
What is the issue of the business there.

HAMLET
It will be short. The interim's mine,
And a man's life's no more than to say "one."
But I am very sorry, good Horatio,
That to Laertes I forgot myself,
For by the image of my cause I see
The portraiture of his. I'll court his favors.
But, sure, the bravery of his grief did put me
Into a tow'ring passion.

HORATIO  Peace, who comes here?

[Enter Osric, a courtier.]


OSRIC  Your Lordship is right welcome back to
Denmark.

HAMLET  I humbly thank you, sir. [Aside to Horatio.]
Dost know this waterfly?

HORATIO, [aside to Hamlet]  No, my good lord.

HAMLET, [aside to Horatio]  Thy state is the more gracious,
for 'tis a vice to know him. He hath much
land, and fertile. Let a beast be lord of beasts and his
crib shall stand at the king's mess. 'Tis a chough,
but, as I say, spacious in the possession of dirt.

OSRIC  Sweet lord, if your Lordship were at leisure, I
should impart a thing to you from his Majesty.

HAMLET  I will receive it, sir, with all diligence of
spirit. Put your bonnet to his right use: 'tis for the
head.

OSRIC  I thank your Lordship; it is very hot.

HAMLET  No, believe me, 'tis very cold; the wind is
northerly.

OSRIC  It is indifferent cold, my lord, indeed.

HAMLET  But yet methinks it is very sultry and hot for
my complexion.

OSRIC  Exceedingly, my lord; it is very sultry, as
'twere--I cannot tell how. My lord, his Majesty
bade me signify to you that he has laid a great wager
on your head. Sir, this is the matter--

HAMLET  I beseech you, remember. [He motions to
Osric to put on his hat.]

OSRIC  Nay, good my lord, for my ease, in good faith.
Sir, here is newly come to court Laertes--believe
me, an absolute gentleman, full of most excellent
differences, of very soft society and great showing.
Indeed, to speak feelingly of him, he is the card or
calendar of gentry, for you shall find in him the
continent of what part a gentleman would see.

HAMLET  Sir, his definement suffers no perdition in
you, though I know to divide him inventorially
would dozy th' arithmetic of memory, and yet but
yaw neither, in respect of his quick sail. But, in the
verity of extolment, I take him to be a soul of great
article, and his infusion of such dearth and rareness
as, to make true diction of him, his semblable is his
mirror, and who else would trace him, his umbrage,
nothing more.

OSRIC  Your Lordship speaks most infallibly of him.

HAMLET  The concernancy, sir? Why do we wrap the
gentleman in our more rawer breath?

OSRIC  Sir?

HORATIO  Is 't not possible to understand in another
tongue? You will to 't, sir, really.

HAMLET, [to Osric]  What imports the nomination of
this gentleman?

OSRIC  Of Laertes?

HORATIO  His purse is empty already; all 's golden words
are spent.

HAMLET  Of him, sir.

OSRIC  I know you are not ignorant--

HAMLET  I would you did, sir. Yet, in faith, if you did, it
would not much approve me. Well, sir?

OSRIC  You are not ignorant of what excellence Laertes
is--

HAMLET  I dare not confess that, lest I should compare
with him in excellence. But to know a man well
were to know himself.

OSRIC  I mean, sir, for his weapon. But in the imputation
laid on him by them, in his meed he's
unfellowed.

HAMLET  What's his weapon?

OSRIC  Rapier and dagger.

HAMLET  That's two of his weapons. But, well--

OSRIC  The King, sir, hath wagered with him six Barbary
horses, against the which he has impawned, as I
take it, six French rapiers and poniards, with their
assigns, as girdle, hangers, and so. Three of the
carriages, in faith, are very dear to fancy, very
responsive to the hilts, most delicate carriages, and
of very liberal conceit.

HAMLET  What call you the "carriages"?

HORATIO  I knew you must be edified by the margent
ere you had done.

OSRIC  The carriages, sir, are the hangers.

HAMLET  The phrase would be more germane to the
matter if we could carry a cannon by our sides. I
would it might be "hangers" till then. But on. Six
Barbary horses against six French swords, their
assigns, and three liberal-conceited carriages--
that's the French bet against the Danish. Why is this
all "impawned," as you call it?

OSRIC  The King, sir, hath laid, sir, that in a dozen
passes between yourself and him, he shall not
exceed you three hits. He hath laid on twelve for
nine, and it would come to immediate trial if your
Lordship would vouchsafe the answer.

HAMLET  How if I answer no?

OSRIC  I mean, my lord, the opposition of your person
in trial.

HAMLET  Sir, I will walk here in the hall. If it please his
Majesty, it is the breathing time of day with me. Let
the foils be brought, the gentleman willing, and the
King hold his purpose, I will win for him, an I can.
If not, I will gain nothing but my shame and the odd
hits.

OSRIC  Shall I deliver you e'en so?

HAMLET  To this effect, sir, after what flourish your
nature will.

OSRIC  I commend my duty to your Lordship.

HAMLET  Yours. [Osric exits.] He does well to commend
it himself. There are no tongues else for 's
turn.

HORATIO  This lapwing runs away with the shell on his
head.

HAMLET  He did comply, sir, with his dug before he
sucked it. Thus has he (and many more of the same
breed that I know the drossy age dotes on) only got
the tune of the time, and, out of an habit of
encounter, a kind of yeasty collection, which carries
them through and through the most fanned
and winnowed opinions; and do but blow them to
their trial, the bubbles are out.

[Enter a Lord.]


LORD  My lord, his Majesty commended him to you by
young Osric, who brings back to him that you
attend him in the hall. He sends to know if your
pleasure hold to play with Laertes, or that you will
take longer time.

HAMLET  I am constant to my purposes. They follow
the King's pleasure. If his fitness speaks, mine is
ready now or whensoever, provided I be so able as
now.

LORD  The King and Queen and all are coming down.

HAMLET  In happy time.

LORD  The Queen desires you to use some gentle
entertainment to Laertes before you fall to play.

HAMLET  She well instructs me.	[Lord exits.]

HORATIO  You will lose, my lord.

HAMLET  I do not think so. Since he went into France, I
have been in continual practice. I shall win at the
odds; but thou wouldst not think how ill all's here
about my heart. But it is no matter.

HORATIO  Nay, good my lord--

HAMLET  It is but foolery, but it is such a kind of
gaingiving as would perhaps trouble a woman.

HORATIO  If your mind dislike anything, obey it. I will
forestall their repair hither and say you are not fit.

HAMLET  Not a whit. We defy augury. There is a
special providence in the fall of a sparrow. If it be
now, 'tis not to come; if it be not to come, it will be
now; if it be not now, yet it will come. The
readiness is all. Since no man of aught he leaves
knows, what is 't to leave betimes? Let be.

[A table prepared. Enter Trumpets, Drums, and Officers
with cushions, King, Queen, Osric, and all the state,
foils, daggers, flagons of wine, and Laertes.]


KING
Come, Hamlet, come and take this hand from me.
[He puts Laertes' hand into Hamlet's.]

HAMLET, [to Laertes]
Give me your pardon, sir. I have done you wrong;
But pardon 't as you are a gentleman. This presence
knows,
And you must needs have heard, how I am punished
With a sore distraction. What I have done
That might your nature, honor, and exception
Roughly awake, I here proclaim was madness.
Was 't Hamlet wronged Laertes? Never Hamlet.
If Hamlet from himself be ta'en away,
And when he's not himself does wrong Laertes,
Then Hamlet does it not; Hamlet denies it.
Who does it, then? His madness. If 't be so,
Hamlet is of the faction that is wronged;
His madness is poor Hamlet's enemy.
Sir, in this audience
Let my disclaiming from a purposed evil
Free me so far in your most generous thoughts
That I have shot my arrow o'er the house
And hurt my brother.

LAERTES  I am satisfied in nature,
Whose motive in this case should stir me most
To my revenge; but in my terms of honor
I stand aloof and will no reconcilement
Till by some elder masters of known honor
I have a voice and precedent of peace
To keep my name ungored. But till that time
I do receive your offered love like love
And will not wrong it.

HAMLET  I embrace it freely
And will this brothers' wager frankly play.--
Give us the foils. Come on.

LAERTES  Come, one for me.

HAMLET
I'll be your foil, Laertes; in mine ignorance
Your skill shall, like a star i' th' darkest night,
Stick fiery off indeed.

LAERTES  You mock me, sir.

HAMLET  No, by this hand.

KING
Give them the foils, young Osric. Cousin Hamlet,
You know the wager?

HAMLET  Very well, my lord.
Your Grace has laid the odds o' th' weaker side.

KING
I do not fear it; I have seen you both.
But, since he is better, we have therefore odds.

LAERTES
This is too heavy. Let me see another.

HAMLET
This likes me well. These foils have all a length?

OSRIC  Ay, my good lord.
[Prepare to play.]

KING
Set me the stoups of wine upon that table.--
If Hamlet give the first or second hit
Or quit in answer of the third exchange,
Let all the battlements their ordnance fire.
The King shall drink to Hamlet's better breath,
And in the cup an union shall he throw,
Richer than that which four successive kings
In Denmark's crown have worn. Give me the cups,
And let the kettle to the trumpet speak,
The trumpet to the cannoneer without,
The cannons to the heavens, the heaven to earth,
"Now the King drinks to Hamlet." Come, begin.
And you, the judges, bear a wary eye.
[Trumpets the while.]

HAMLET  Come on, sir.

LAERTES  Come, my lord.	[They play.]

HAMLET  One.

LAERTES  No.

HAMLET  Judgment!

OSRIC  A hit, a very palpable hit.

LAERTES  Well, again.

KING
Stay, give me drink.--Hamlet, this pearl is thine.
Here's to thy health.
[He drinks and then drops the pearl in the cup.]
[Drum, trumpets, and shot.]
Give him the cup.

HAMLET
I'll play this bout first. Set it by awhile.
Come. [They play.] Another hit. What say you?

LAERTES
A touch, a touch. I do confess 't.

KING
Our son shall win.

QUEEN  He's fat and scant of breath.--
Here, Hamlet, take my napkin; rub thy brows.
The Queen carouses to thy fortune, Hamlet.
[She lifts the cup.]

HAMLET  Good madam.

KING  Gertrude, do not drink.

QUEEN
I will, my lord; I pray you pardon me.	[She drinks.]

KING, [aside]
It is the poisoned cup. It is too late.

HAMLET
I dare not drink yet, madam--by and by.

QUEEN  Come, let me wipe thy face.

LAERTES, [to Claudius]
My lord, I'll hit him now.

KING  I do not think 't.

LAERTES, [aside]
And yet it is almost against my conscience.

HAMLET
Come, for the third, Laertes. You do but dally.
I pray you pass with your best violence.
I am afeard you make a wanton of me.

LAERTES  Say you so? Come on.	[Play.]

OSRIC  Nothing neither way.

LAERTES  Have at you now!
[Laertes wounds Hamlet. Then in scuffling they change
rapiers, and Hamlet wounds Laertes.]

KING  Part them. They are incensed.

HAMLET  Nay, come again.
[The Queen falls.]

OSRIC  Look to the Queen there, ho!

HORATIO
They bleed on both sides.--How is it, my lord?

OSRIC  How is 't, Laertes?

LAERTES
Why as a woodcock to mine own springe, Osric.
[He falls.]
I am justly killed with mine own treachery.

HAMLET
How does the Queen?

KING  She swoons to see them bleed.

QUEEN
No, no, the drink, the drink! O, my dear Hamlet!
The drink, the drink! I am poisoned.	[She dies.]

HAMLET
O villainy! Ho! Let the door be locked.	[Osric exits.]
Treachery! Seek it out.

LAERTES
It is here, Hamlet. Hamlet, thou art slain.
No med'cine in the world can do thee good.
In thee there is not half an hour's life.
The treacherous instrument is in thy hand,
Unbated and envenomed. The foul practice
Hath turned itself on me. Lo, here I lie,
Never to rise again. Thy mother's poisoned.
I can no more. The King, the King's to blame.

HAMLET
The point envenomed too! Then, venom, to thy
work.	[Hurts the King.]

ALL  Treason, treason!

KING
O, yet defend me, friends! I am but hurt.

HAMLET
Here, thou incestuous, murd'rous, damned Dane,
Drink off this potion. Is thy union here?
[Forcing him to drink the poison.]
Follow my mother.	[King dies.]

LAERTES  He is justly served.
It is a poison tempered by himself.
Exchange forgiveness with me, noble Hamlet.
Mine and my father's death come not upon thee,
Nor thine on me.	[Dies.]

HAMLET
Heaven make thee free of it. I follow thee.--
I am dead, Horatio.--Wretched queen, adieu.--
You that look pale and tremble at this chance,
That are but mutes or audience to this act,
Had I but time (as this fell sergeant, Death,
Is strict in his arrest), O, I could tell you--
But let it be.--Horatio, I am dead.
Thou livest; report me and my cause aright
To the unsatisfied.

HORATIO  Never believe it.
I am more an antique Roman than a Dane.
Here's yet some liquor left.	[He picks up the cup.]

HAMLET  As thou 'rt a man,
Give me the cup. Let go! By heaven, I'll ha 't.
O God, Horatio, what a wounded name,
Things standing thus unknown, shall I leave behind
me!
If thou didst ever hold me in thy heart,
Absent thee from felicity awhile
And in this harsh world draw thy breath in pain
To tell my story.
[A march afar off and shot within.]
What warlike noise is this?

[Enter Osric.]


OSRIC
Young Fortinbras, with conquest come from Poland,
To th' ambassadors of England gives
This warlike volley.

HAMLET  O, I die, Horatio!
The potent poison quite o'ercrows my spirit.
I cannot live to hear the news from England.
But I do prophesy th' election lights
On Fortinbras; he has my dying voice.
So tell him, with th' occurrents, more and less,
Which have solicited--the rest is silence.
O, O, O, O!	[Dies.]

HORATIO
Now cracks a noble heart. Good night, sweet prince,
And flights of angels sing thee to thy rest.
[March within.]
Why does the drum come hither?

[Enter Fortinbras with the English Ambassadors with
Drum, Colors, and Attendants.]


FORTINBRAS  Where is this sight?

HORATIO  What is it you would see?
If aught of woe or wonder, cease your search.

FORTINBRAS
This quarry cries on havoc. O proud Death,
What feast is toward in thine eternal cell
That thou so many princes at a shot
So bloodily hast struck?

AMBASSADOR  The sight is dismal,
And our affairs from England come too late.
The ears are senseless that should give us hearing
To tell him his commandment is fulfilled,
That Rosencrantz and Guildenstern are dead.
Where should we have our thanks?

HORATIO  Not from his
mouth,
Had it th' ability of life to thank you.
He never gave commandment for their death.
But since, so jump upon this bloody question,
You from the Polack wars, and you from England,
Are here arrived, give order that these bodies
High on a stage be placed to the view,
And let me speak to th' yet unknowing world
How these things came about. So shall you hear
Of carnal, bloody, and unnatural acts,
Of accidental judgments, casual slaughters,
Of deaths put on by cunning and forced cause,
And, in this upshot, purposes mistook
Fall'n on th' inventors' heads. All this can I
Truly deliver.

FORTINBRAS  Let us haste to hear it
And call the noblest to the audience.
For me, with sorrow I embrace my fortune.
I have some rights of memory in this kingdom,
Which now to claim my vantage doth invite me.

HORATIO
Of that I shall have also cause to speak,
And from his mouth whose voice will draw on
more.
But let this same be presently performed
Even while men's minds are wild, lest more
mischance
On plots and errors happen.

FORTINBRAS  Let four captains
Bear Hamlet like a soldier to the stage,
For he was likely, had he been put on,
To have proved most royal; and for his passage,
The soldier's music and the rite of war
Speak loudly for him.
Take up the bodies. Such a sight as this
Becomes the field but here shows much amiss.
Go, bid the soldiers shoot.
[They exit, marching, after the which, a peal of
ordnance are shot off.]
`

	b["hurston-de-turkey-and-de-law.txt"] = `
The Project Gutenberg EBook of De Turkey and De Law, by Zora Neale Hurston

This eBook is for the use of anyone anywhere at no cost and with
almost no restrictions whatsoever.  You may copy it, give it away or
re-use it under the terms of the Project Gutenberg License included
with this eBook or online at www.gutenberg.org

Title: De Turkey and De Law
       A Comedy in Three Acts

Author: Zora Neale Hurston

Release Date: July 25, 2007 [EBook #22146]

Language: English


*** START OF THIS PROJECT GUTENBERG EBOOK DE TURKEY AND DE LAW ***




Produced by Charlene Taylor and the Online Distributed
Proofreading Team at http://www.pgdp.net (This file was
produced from images generously made available by the
Library of Congress)





[Transcriber's Notes: This play transcribed from an original
typewritten manuscript at the Library of Congress in the Zora Neale
Hurston collection. There are pencilled notations probably by Ms.
Hurston herself. These pencilled edits have been transcribed as [Note:
(text)] Any other questionable transcription is similarly noted.
Nothing in the dialect has been changed. Occasional obvious typos in
the stage directions have been corrected. There are inconsistencies in
both bracketing and punctuation, which have been left as in the
original manuscript. There seems to be an irregularity in the spelling
of "Simms"; "Sims" seems randomly substituted.]




DE TURKEY AND DE LAW

A COMEDY IN THREE ACTS

by

ZORA HURSTON




CAST

Jim Weston             A young man and the town bully (A Methodist)

Dave Carter            The town's best hunter and fisherman (Baptist)

Joe Clarke             The Mayor, Postmaster, storekeeper

Daisy Blunt            The town vamp

Lum Boger              The Marshall

Walter Thomas          A villager (Methodist)

Lige Moseley           A villager (Methodist)

Joe Lindsay            A villager (Baptist)

Della Lewis            A villager (Baptist)

Tod Hambo              A villager (Baptist)

Lucy Taylor            A villager (Methodist)

Rev. Singletary                   (Baptist)

Rev. Simms                        (Methodist)

Villagers, children, dogs.




ACT I

SETTING: A Negro village in Florida in our own time. All action from
viewpoint of an actor facing audience.

PLACE: Joe Clarke's store porch in the village. A frame building with
a false front. A low porch with two steps up. Door in center of porch.
A window on each side of the door. A bench on each side of the porch.
Axhandles, hoes and shovels, etc. are displayed leaning against the
wall. Exits right and left. Street is unpaved. Grass and weeds growing
all over.

TIME: It is late afternoon on a Saturday in summer.

Before the curtain rises the voices of children are heard, boisterous
at play. Shouts and laughter.


VOICE OF ONE BOY
Naw, I don't want to play wringing no dish rag! We gointer play chick
mah chick mah craney crow.

GIRL'S VOICE
Yeah, less play dat, and I'm gointer to be de hen.

BOY'S VOICE
And I'm gointer be de hawk. Lemme git myself a stick to mark wid. (The
curtain rises slowly. As it goes up the game is being organized. The
boy who is the hawk is squatting center stage in the street before the
store with a short twig in his hand. The largest girl is lining up the
other children behind her.)

THE MOTHER HEN
(looking back over her flock) Y'all ketch holt of one 'nother's
clothes so de hauk can't git yuh. (They do.) Y'all straight now?

CHORUS
Yeah. (The march around the hawk commences.)

HEN AND CHICKS
     Chick mah chick mah craney crow
     Went to de well to wash my toe
     When I come back my chick was gone.
     What time ole witch?

HAWK
(making a tally on the ground) One!

HEN AND CHICKS
Chick mah chick etc.--(While this is going on Walter Thomas from the
store door eating peanuts from a bag appears and seats himself on the
porch beside the steps.)

HAWK
(Scoring again) Two!
                       (Enter a little girl right. She trots up to
the big girl.)

LITTLE GIRL
(officiously) Titter, mama say if you don't come on wid dat soap she
gointer wear you out.

HEN AND CHICKS
Chick mah chick etc. (While this is being sung, enter Joe Lindsay and
seats himself on right bench. He lights his pipe. The little girl
stands b by the fence rubbing her leg with her foot.

HAWK
(scoring) Three!

LITTLE GIRL
(insistent) Titter, titter! Mama say to tell you to come on home wid
dat soap and rake up dat yard. I bet she gointer beat you good.

BIG GIRL
(angrily) Aw naw, mama ain't sent you after me, nothin' of de kind!
Gwan home and leave me alone.

LITTLE GIRL
You better come on! I'm gointer tell mama how 'omanish you actin
cause you in front of dese boys.

BIG GIRL
(makes a threatenin' gesture) Aw don't be so fast and showin' off in
company. Ack lak you ain't got no sense!

LITTLE GIRL
(starts to cry) Dat's all right. I'm going home and tell mama you down
here playing wid boys and she sho gointer whup you good, too. I'm
gointer tell her you called me a fool too, now. (She walks off, wiping
her eyes and nose with the back of her hand) Yeah, I'm goin' tell her!
Jus' showin' off in front of ole John Wesley Taylor. I'm going to tell
her too, now.

BIG GIRL
(flounces her skirt) Tell her! Tell her! Turn her up and smell her!
(Game resumed) Chick mah chick etc.

HAWK
Four! (He arises and imitates a hawk flying and trying to catch a
chicken. Calling in a high voice.) Chickie!!

HEN
(Flapping her wings to protect her young) My chickens' sleep.

HAWK
Chickie!!

HEN
My chickens' sleep.

HAWK
I shall have a chick.

HEN
You shan't have a chick.

HAWK
I'm going home. (flies off)

HEN
There's de road.

HAWK
I'm comin' back.

(During this dialog the hawk is feinting and darting in his efforts to
catch a chicken and the chickens are dancing defensively.)

HEN
Don't keer if you do.

HAWK
My pot's a boiling.

HEN
Let it boil.

HAWK
My guts a growling

HEN
Let 'em growl.

HAWK
I must have a chick.

HEN
You shan't have nairn.

HAWK
My mama's sick.

HEN
Let her die.

HAWK
Chickie!!

HEN
My chicken's sleep.

(Hawk darts quickly around the hen and grabs a chicken and leads him
off and places the captive on his knees at the store porch. After a
brief bit of dancing he catches another, then a third who is a chubby
little boy. The little boy begins to cry.)

LITTLE BOY
I ain't gointer play cause you hurt me.

HAWK
Aw, naw, I din't hurt you.

LITTLE BOY
Yeah you did too. You pecked me right here. (points to top of his
head)

HAWK
Well if you so touchous you got to cry every time anybody look at you,
you can't play wid us.

LITTLE BOY
(smothering sobs) I ain't cryin'. (He is placed with the other
captives. Hawk returns to game.)

HAWK
Chickie.

HEN
My chickens sleep!

VOICE FROM A DISTANCE
Titter! You Titter!!!

BIG GIRL
Yessum

VOICE
If you don't come here wid dat soap you better!

BIG GIRL
(shakes herself poutingly, half sobs) Soon's I git grown I'm gointer
run away. Everytime a person gits to havin' fun, it's "come here,
Titter and rake de yard." She don't never make Bubber do nothin. (She
exits into the store.)

HAWK
Now we ain't got no hen.

ALL THE GIRLS
(in a clamor) I'll be de mama hen! Lemme be it! (Enter Hambo left and
stands looking at the children.)

HAMBO
Can't dese young uns keep up a powerful racket, Joe?

LINDSAY
They sho kin. They kin git round so vi'grous when they whoopin and
hollerin and rompin and racin, but just put 'em to work now and you
kin count dead lice fallin' off of 'em.

(Enter Tillie from the store with the soap. Hambo pulls out a plug of
tobacco from his hip pocket and bites a chunk from it.)

HAMBO
De way dese chillun is dese days is,--eat? Yes! Squall and holler?
Yes! Kick out shoes? Yes! Work? No!!

LINDSAY
You sho is tellin' de truth. Now look at dese! I'll bet everyone of
'em's mammies sent 'em to de store an' they out here frollickin'. If
one of 'em was mine, I'd whup 'em till they couldn't set down. (to the
children) Shet up dat racket and gwan home! (The children pay no
attention and the game gets hotter.)

DISTANT VOICE
(off stage) You Tit-ter!! You Tit-Ter!!

WALTER
Titter, don't you hear yo' ma callin' you?

ESSIE
Yessuh, I mean naw suh.

LINDSAY
How come you can't answer, then? Lawd knows de folks just ruins
chilluns dese days. Deys skeered tuh whup 'em right. Den before they
gits twenty de gals done come up wid somethin' in dey arms an' de boys
on de chain gang. If you don't whup 'em, they'll whip you.

HAMBO
Dat sho is whut de Lawd loves. When I wuz a boy they _raised_ chillen
then. Now they lets 'em do as they please. There ain't no real
chestizing no more. They takes a lil tee-ninchy switch and tickles em.
No wonder de world is in sich uh mess.

VOICE OFF STAGE
You Tit-ter!! Aw Titter!!

ESSIE
(stops to listen) Yessum!!

VOICE OFF STAGE
If you don't come here, you better!

ESSIE
Yessum! (to her playmates) Aw shucks! I got to go home. (She exits
right, walking sullenly. The game has stopped.)

LINDSAY
(pointing at Essie) You see dat gal shakin' herself at her mammy? De
sassy lil binch needs her guts stomped out. (to Essie) Run! I'm comin'
on down there an' tell yo' ma how 'omanish you is, shakin' yo'self at
grown folks. (Essie walks slower and shakes her skirt contemptously.
Lindsay jumps to his feet as if to pursue her.) You must smell
yo'self! (Essie exits.) Now de rest of you haitians scatter way from
in front dis store. Dis ain't no place for chillen, nohow. (gesture of
shooing) Gwan! Thin out! Every time a grownperson open they mouf y'all
right dere to gaze down they throat. Git! (The children exit sullenly
right. In the silence that follows the cracking of Walter's peanut
shells can be heard very plainly.)

HAMBO
Walter, God a' mighty! You better quit eatin' em ground peas de way
you do. You gointer die wid de colic.

LINDSAY
Aw, taint gointer hurt him. I don't b'lieve uh cord uh wood would lay
heavy on Walter's belly. He kin eat mo' penders than Brazzle's mule.

WALTER
(laughing) Aw naw, don't throw me in wid dat mule. He could eat up
camp-meetin, back off scociation and drink Jurdan dry.

LINDSAY
And still stay so po' till he wuzn't nothin atall but a mule frame.
(Enter Lige Moseley right) Taint never been no mule in de world lak
dat ole yaller mule since Jonah went to joppy.

(Lige seats himself on the floor on the other side of the steps. Pulls
out a bone toothpick and begins to pick his teeth)

LIGE
Y'all still talkin bout Brazzle's ole useter-be mule?

HAMBO
Yeah. Memeber dat time Brazzle hitched him to de plow and took him to
Eshleman's new ground?

LIGE
And he laid down before he'd plow a lick. Sho I do! But who ever seen
him work? All you ever did see was him and Brazzle fightin up and down
de furrows. (all laugh) He was so mean he would even try to kick you
if you went in his stall to carry him some corn.

WALTER
Nothin but pure concentrated meanness stuffed into uh mule hide. Thass
de reason he wouldn't git fat--just too mean.

LIGE
Sho was skinny now. You could use his ribs for a washboard and hang de
clothes up on his hips to dry. (all laugh)

HAMBO
Lige, you kin lie [Note: "like" crossed out] lak cross ties [Note
inserted text: from Jacksonville to Key West.]. But layin all sides to
jokes, when they told me dat mule was dead, uh just took and knocked
off from work to see him drug out lak all de rest of de folks, and
folkses dat mule wuz too contrary to lay down on his side and die. He
laid on his raw-boney back wid his foots stickin straight up in de air
lak he wuz fightin something.

LINDSAY
He wuz--bet he fought ole death lak a natural man. Ah seen his bones
yistiddy, out dere on de edge of de cypress swamp. De buzzards done
picked em clean and de elements done bleached em.

LIGE
Everybody went to dat draggin out. Even Joe Clarke shet up his store
dat mornin and went (turns his head and calls into the store) didn't
you, Mr. Clarke?

CLARKE'S VOICE
Didn't I whut? (enters and stands in door)

LIGE
Shet up yo' store and go to de draggin out of Brazzle's ole mule.

CLARKE
I, God, Yeah. It was worth it. (sees Hambo) I didn't know you was out
here. Lemme beat you uh game of checkers.

HAMBO
Lissen at de ole tush hawg! Well, go git de board, and lemme beat you
a pair of games befo' de mail gits in.

CLARKE
(to the others) Beat old me! (to Hambo) Come on here, youse my fish.
(calls into store) Mattie bring me dat checker-board and de checkers!
(to men on porch) You got to talk to wimmen-folks lak dat--tell 'em
every lil' thing-do she'd come rackin out here wid de board by itself.

(Enter Mrs. Clarke with homemade checker-board and coffee can
containing the much-used checkers. Clarke sits on a keg and faces
Hambo. They put the board on their knees and pour out the checkers)

HAMBO
You want black or red?

CLARKE
Oh, I don't keer which--I'm gointer beat you anyhow. You take de
black. (they arrange them. The others get near to look on. Hambo sits
looking at the board without moving.)

HAMBO
Who's first move?

CLARKE
Black folks always go to work first. Move! (Hambo moves and the same
proceeds with the spectators very interested. Enter Lum Boger [Note:
Handwritten correction: Bailey] right and joins the spectators. A
woman enters left with a market basket and goes on in the store. The
checkers click on the board. A girl about twelve enters right and goes
into the store and comes out with a stick of peppermint candy.

WALTER
Naw you don't Hambo!--Don't you go in dere! Dats a trap--(pointing)
come right here and you got him.

LIGE
Back dat man up (pointing) Hambo do he'll git et up.

(there is the noise of the checkers for a half minute then a general
shout of triumph)

SPECTATORS
You got him now, Hambo! Clarke, he's sho got you.

CLARKE
(Chagrined) Aw, he aint done nothin! Jes' watch ME.

HAMBO
(Jeering) Yeah, gwan move! Ha! Ha! go head and move.

SPECTATORS
Aw, he got you, Bro. Mayor--might as well give up. He got you in de
Louisville loop.

CLARKE
Give up what? He can't beat me? (peeved) de rest of y'all git from
over me, whoopin and hollerin! I God, a man can't hear his ears.

(The men fall back revealing the players clearly)

HAMBO
Aw, neb mind bout them, Joe, go head and move. You aint got but one
move to make nohow--go head on and take it.

CLARKE
(moving a checker) Aw, here.

HAMBO
(triumphant) Now! watch me boys whut Ahm gonna do to him. Ahm gonna
laff in notes, while Ah work on him. (he lifts a checker high in the
air preparatory to the jump, laughing to the scale and counting each
checker he jumps out loud) Do, sol, fa, me, la! One! (jumps a checker)
la, sol, fa, me, do! Two! (jumps another) Do, re, fa, me, do, Three!
Me, re, la, so, fa! Four! (the crowd is roaring with laughter) Sol,
fa, me, la, sol, do! Five! Ha! Ha! boys I got [Note: "the" x-ed out]
de ole tush hawg! I got him in de go-long. (He slaps his leg and
accidently knocks the board off his knee and spills the checkers.)

CLARKE
Too bad you done dat, Hambo, cause Ah was gointer beat you at dat (he
rises and starts towards the door of the store as the crowd roars in
laughter)

HAMBO
You mean you was gointer beat me to de door, not a game Of checkers.
Ah done run de ole coon in his hole.

LIGE
Well, Hambo, you done got to be so hard at checkers, come on less see
whut you can do wid de cards. (He pulls out a soiled deck from his
coat pocket and moves toward the bench at the left of the porch) You
take Lum and me and Walter will wear you out.

HAMBO
You know I don't play no cards.

LUM
We aint playin for no money, just a lil Florida flip.

HAMBO
Y'all can't play no Florida flip. 'Fore Ah joined de church there
wasn't a man in de state could beat me wid de cards. But Ahm a deacon
now, in Macedonia Baptist--Ah don't bother wid de cards no mo". (He
and Joe Lindsay go inside store)

LIGE
Well, come on Lum. Walter, git yo'self a partner.

WALTER
(Looking about) Taint nobody to git (looks off right) Here come Dave
Carter.

LIGE
You can't do nothin wid him dese days. He useter choose a game of
cards when he wasn't out huntin, but now when he ain't out huntin
varmints he's huntin' Daisy Blunt. (Enter Dave right with a shot-gun
slung over his shoulder.)

WALTER
Come on, fish, lemme bend a five-up over yo' head. You looks just like
my meat.

DAVE
Ahm on mah way to kill me a turkey gobbler, but if you and Lum thinks
y'all's tush hawgs Ah'll stop long enough to take you down a
button-hole lower. (He sets his gun down and finds a seat and draws it
up to the card table)

WALTER
Naw, Dave, we aint going to fool wid no button-holes we gointer tear
off de whole piece dat de button-holes is in. (They all get set) All
right boys, turn it on and let de bad luck happen.

LIGE
(Probbing the deck) My deal.

WALTER
Watch yo'self Dave, don't get to worryin bout Daisy and let 'em ketch
yo' jack.

LUM
(Winking) What you reckon he gointer be worryin' bout Daisy for? Dot's
Jim's gal.

DAVE
Air Lawd, a heap sees but a few knows. Deal de cards man--you
shufflin' a mighty lot.

WALTER
Sho is--must be tryin' to carry de cut to us.

LIGE
Aw, we ain't gonna cheat you, we gonna beat you. (He slams down the
cards for Dave to cut) Wanna cut 'em?

DAVE
Nope. Taint no use cuttin' a rabbit out when you kin twist him out.
Deal 'em! (Lige deals and turns up Jack of spades.)

WALTER
Yee-ee! Did you snatch dat Jack?

LIGE
Man, you know I ain't snatched no Jack. Whut you doin'?

WALTER
I'm beggin!

LIGE
Go ahead and tell 'em I sent you.

WALTER
Play just like ahm in New York, partner. (scratches his head) We
oughter try to ketch dat Jack.

LIGE
Stick out yo' hand an' you'll draw back a nub.

WALTER
Whut you want me to play for you, partner?

DAVE
Play me a baby diamond.
                   (Walter plays, then Lum, then Dave)

LUM
(Triumphant) Looka pardner, they doin all dat woofin on uh
queen--sendin' women to do uh man's work. Watch me stomp her wid mah
king (He slams his card down and collects the trick.) Now come un
under dis ace! (They all play and he collects the trick.) Now whut you
want me to play for you, pardner?

LIGE
How many times you seen de deck.

LUM
Twice

LIGE
Pull off wid yo' king.

(Lum plays the king of spades. All the others play.)
Look at ole low pardner. Ah knowed ah wuz gointer ketch him! Come
right back at 'em.

LUM
(stands up and slams down the ace) Pack up, pardner. Ahm playin' mah
knots, now all play now. Ho! Ho! Dere goes de queen'. De Jack's a
gentleman! (Lige takes the Jack and sticks it up on his forehead in
braggadocia.) Here comes de ten spot, pardner, ahm dumpin to yuh!

LIGE
(as he plays the Jack) Everybody git up off it and dump. High, low,
Jack, game and gone from de first four.

WALTER
Gimme dem cards! Y'all carried de cub to us dat time. (riffles the
cards elaborately) but de deal is in de high, tall house now. Dis is
Booker T Washington spreadin his mess. (offers cards to Lige) Cut?

LIGE
Yeah, cut 'em and shoot 'em. I'd cut behind mah ma. (He cuts and
Walter deals.)

WALTER
Well, whut sayin'?

LUM
I'm beggin.

WALTER
Get up off yo' knees. Youse dat one.

LIGE
Walter, you sho stacked dese cards.

WALTER
Aw, stop cryin' and play, man. Youse too old to be hollerin'
titty-mama.

LUM
Dis ain't no hand, dis is a foot. What you want me to play for you
partner?

LIGE
Play yo' own hand partner--I ain't nobody. Lead yo' bosses. (He leads
the ace of clubs. Play goes round to dealer and Walter takes the card
off the deck and slams it down.)

WALTER
Get up ol' deuce of diamonds and gallop off wid yo' load. Pardner, how
many times you seen de deck?

DAVE
(Two times--(they make signals.)

WALTER
Watch dis ol' queen. Less go! (He begins to sing--Dave joins in.) When
yo' card gits lucky, oh pardner, you oughter be in a rollin' game. (He
speaks.) Ha! Ha! Wash day and no soap! (He sticks the Jack upon his
forehead. He stands up and sings again.) Ahm goin' to de 'Bama Lawd.
Pardner don't want no change. (He collects that trick and plays again.
Dave also stands.)

DAVE
Here come de man from de White House--ol' king of diamonds. (Sings,
all join.) Ahm goin' back to de Bama, Lawd. Pardner won't be worried
wid you. (He collects the trick.) Never had no trouble, Lour pardner,
till I stopped by here.

(They all stand hilariously slam down their cards.

WALTER
Aw, wese just too hard for you boys--we eats our dinner out de
blacksmith shop. Y'all can't bully dis game. (He solemnly reaches over
and takes Dave's hand.)

DAVE
(to Walter) Mr. Hoover, you sho is a noble president. We done stuck
dese shad-moufs full of cobs. They skeered to play us any mo'.

LIGE
Who skeered? Y'all jus' playin ketch up nohow. Git back down and lemme
wrap uh five-up round yo' neck.

DAVE
(looking off right) Squat dat rabbit an' less jump another one. Here
come Daisy.

WALTER
Aw Lord, you ain't no mo' good now. But Ah don't blame you, Dave, she
looks warm.

(Enter Daisy right with a scarlet hibiscus over each ear and smiling
broadly.)

LIGE
(jumps down and takes Daisy by the arm) Come on up here, Daisy and
ease Dave's pain. He's so crazy 'bout you his heart 'bout to burn a
hole in his shirt. (She steps up on the porch)

DAVE
(Bashfully) Aw, y'all gwan. Ah kin talk.

DAISY
(Arms akimbo, impudently) Oh kin you? (She gets up close to Dave)

DAVE
(Pleased) You better git way from me fore Jim come long.

DAISY
(Coquettishly) Ain't you man enough to cover de ground you stand
on?

DAVE
Oh, Ah can back my crap! Don't worry 'bout me. Where you headed for?

DAISY
Where _you_ goin? (Audaciously)

DAVE
Out by de cypress swamp to kill us uh turkey. Its uh great big ole
gobbler--been slurring me fer six months. Ahm gointer git him today
for you, and yo' mama gointer cook him.

DAISY
Ah sho would love the ham of turkey.

DAVE
(Patting his gun barrel) Well me an' ole Hannah sho gointer git you
one. Look here, Daisy, will you choose uh bag of ground peas?

DAISY
I jus love goobers

DAVE
(Sticking out his right elbow) You lak chicken?

DAISY
Yeah

DAVE
Take uh wing. (She locks arms with him and they strut inside the
store.)

LIGE
Ah blieve dat fool is got some gumption. Jim Weston better watch out.

WALTER
Oh I ain't never figgered Dave was no fool. He's uh bottom fish. Jim
talks all de time but Dave will run him uh hot--here he come now.
(Looks off left. All look the same way.)

LUM
Lawd, don't he look mean? (She chuckles) Ah bet he know Daisy's here
wid Dave. Ah wouldn't take nothin' for dis.

(Enter Jim Weston left with a guitar looking very glum. He stops
beside the step for a moment. Takes off his hat and fans with it.)

JIM
Howdy do, folks.

ALL
Howdy do, Jim.

JIM
Don't do all they say. (He sees the gun leaning against the rail) Who
gun dat? (Points at the gun)

LIGE
You know so well whose gun dat is. Ah jus' heard him say he's goin out
to git his gal uh ham of a turkey gobbler out round de cypress swamp.
He's inside now treatin her to penders and candy. (He winks at the
others and they wink back)

WALTER
(Turns and calls into the store) Say, Dave! Don't try to keep Daisy in
dere all day. Her feller out here waitin to scorch her home.

DAVE (from inside store)
Let him come git her if she want him.

LIGE
Umph! dere now, de mule done kicked Rucker! (Calls inside to Dave) I
hear you crowin, rooster. I know yo' nest aint far.

HAMBO
(From inside store) Yeah, dis rooster must know something--he's gittin
plenty grit in his craw.

(General laughter)

(There is a gay burst of laughter from inside the store. In a moment
Dave enters from the store with Daisy on his left arm. With his right
he is stuffing shells into his pocket. The air is tense. Lindsay,
Hambo and Joe Clarke all enter behind the couple)

DAVE
(Releases Daisy and steps to the edge of the porch right in front of
Jim and looks up at the sky) Well, sun's gettin low--better git on out
to de swamp and git dat gobbler. (He turns and picks up de gun and
breaks it)

JIM
Lo Daisy. (Sullenly)

DAISY
(Brightly) Hello Jimmy (She is eating peanuts) Ain't Dave smart? He's
gonna kill me uh turkey an' ah kin eat all ah wants.

JIM
He aint de onliest person kin shoot round here.

LIGE
Yeah, but he's best marksman just de same. Taint no use talkin, Jim.
You can't buck Dave in de woods. But you got de world beat wid uh
git-fiddle. Yessuh, Dave is uh sworn marksman but you kin really beat
de box. Less have uh tune.

JIM
Oh I ain't for pickin no box. I come to git some shells for my rifle.
Sorta figgered on uh wild turkey or two. (He comes up on the porch and
starts in the store)

DAISY
If Dave go git me dat big ole turkey an' you go git me one too--gee!
Wont I have uh turkey fit?

LINDSAY
Lord, Daisy, you gointer have dese boys killin up every turkey in
Orange County.

WALTER
You mean _Dave_. Jim couldn't hit de side of uh barn wid uh brass
fiddle.

JIM
(Hitching up his trousers) Who can't shoot? (to Clarke) Come on an'
gimme un box uh shells. I'll show yuh who kin shoot! (He exits into
store with Clarke behind him)

DAVE
(To Daisy) You wait here till ah git back wid yo' turkey.

DAISY
Ahm skeered.

DAVE
Whut you skeered of? Jim? He aint no booger boo, if his ears do flop
lak uh mule.

DAISY
Naw. Ah aint skeered uh no Jim. Ah got tuh git back tuh de white folks
an Ahm skeered tuh go round dat lake at night by myself. (Enter Jim
from store and stands in door with box of shells in his hand)

JIM
No girl look like you don't have to go home by yo' self, if it was
midnight.

DAVE
(Gun in hand and ready to exit) Naw, cause Ahm right here--

JIM
Daisy don't you trust yo'self round dat lake after dark, wid dat
(points at Dave) breath and--britches. You needs uh real man to
perteck you from dem 'gators and moccasins.

DAVE
Let somethin happen and she'll find out who got rabbit blood and who
aint. Well, Ahm gone. (He steps down off the steps but looks back at
Daisy).

JIM
Ahm goin too--git you uh great big ole turkey-rooster. (Dave takes a
step or two towards left exit).

DAISY
Jim, aint you gointer knock off a li'l tune fo' you go? Ahm lonesome
for some music.

(Dave stops in his tracks and looks wistful. Jim sets down the shells
on the bench and picks up his box with a swagger and tunes a bit.)

WALTER
Georgy Buck!

JIM
(Plays the air thru once then starts to sing. Dave leans his gun
against the fence and stands there.)

  1.
  Georgy Buck is dead, last word he said
  I don't want no shortenin in my bread.

  2.
  Rabbit on de log--Aint got no dog
  How am I goin git him, God knows.

(Dave walks on back near the step, and begins to buck a wing. Daisy
comes down the step admiring both the playing and the dancing. All the
men goin in singing and clapping)

  3.
  Rabbit on de log--aint got no dog
  Shoot him wid my rifle, bam! bam!

  4.
  Oh Georgie Buck is dead, last word he said
  Never let a woman have her way

(The tempo rises. As Dave does a good break he brings up directly in
front of Daisy. He grabs her and swings her into a slow drag. The
porch cheers. Jim stops abruptly. (Enter two women, right and hurry up
to the porch)

1st WOMAN (LULU)
Don't stop, Jim! Hit dat box a couple mo' licks so some of dese men
kin scorch us in de store and treat us.

JIM
Aw, I dont feel lak no playin.

DAVE
(Grinning triumphantly) Ahm gone dis time to git dat turkey. Daisy run
tell yo' ma to put on de hot water kittle (He exits left with gun on
shoulder)

DAISY
Oh lemme see if I got a letter in de postoffice (She exits into store)

JIM
He better git for home fore ah bust dis box over his head.

2nd WOMAN (Jenny)
(Grabbing Lige) Aw, don't worry bout Dave Carter. Play us some music
so I kin make Lige buy me some soda water. (She is playfully dragging
Lige towards the door). Jenny you grab Walter.

(Walter makes a break to jump off the porch and run. The woman catches
him and there is a very gay bit of tussling as the men are dragged
towards the door)

1st WOMAN (Miss Lulu)
I bet if this was Daisy, they'd uh done halted inside and toted out
half de store.

JENNY
Yeah. (gets Walter to the door) Everything you hear is Daisy, Daisy,
Daisy! Just cause she got a walk on her like she done gone crazy thru
de hips! (Yanks Walter into the door) Yeah, y'all goin treat us. Come
on!

WALTER
Yeah, but Daisy's uh young pullet and you gittin gray headed.

JENNY
Thank God I aint gray elsewhere! Come right on. You gointer buy me
some soda water nigger. (to Jim) Play us some music, Jim, so we kin
grand march up to de counter.

JIM
I can't play nothin--mad as I is. I'm one minute to boilin and two
minutes to steam. I smell blood!

MISS LULU
You don't want to fight, do you?

JIM
Sho do. You aint never seen a Weston yet dat wouldn't fight, have you?

LIGE
Thats whut they all got run outa town for--fightin. (Calls into store)
Hey, Joe, give Jenny and Lulu some soda water and ground peas on me so
they'll turn us loose. (to Jim) Yeah, y'all Westons blieves in
fightin.

JIM
Ahd ruther get run out for fightin than to be uh coward. (He slings
the guitar round his neck an' picks up his box of shells.) Well, Ah
reckon Ah'll go git Daisy her turkey cause she sho wont git none less
Ah go git it. Here come Elder Simms anyhow now taint no mo' pickin de
box. (to Daisy) Don't git lonesome whilst Ahm gone.

(Enter Daisy from the store smiling, and walks down to where Jim is
standing)

DAISY
Whuts all dis talk about fightin?

JIM
Lige throwin it up to me bout all my folks been run outa town for
fightin. But I don't keer!

DAISY
Mah mouf done got lonesome already. Buy me some chewing gum to keep
mah mouf comp'ny till y'all gits back wid dat turkey.

JIM
Don't hafta buy none. (reaches in his pocket and pulls out a stick)
What it takes tuh satisfy de ladies, Ah totes it. (He hands her the
gum tenderly) 'By, Daisy. (He walks to left exit)

DAISY
(Coyly) Bye, till you come back.

(Enter Elder Simms right)

Good evenin' everybody.

ALL
Good evenin', Elder Sims.

LUM
(Getting up from his seat on the porch) Have mah seat, Elder. Sims
takes it with a sigh of pleasure. Lum steps off the porch and sets his
hat over one eye) Say, Daisy, you aint goin to sprain yo' lil mouf on
dat tough chewin gum, is yuh? Not wid de help _you_ got. Better lemme
kinda tender dat gum up for yuh so yo' lil mouf won't hafta strain wid
it. (He places himself exactly in front of her. She glances up coyly
at him)

DAISY
Ain't you crazy, now? (Lum tries to snatch the gum but she pops it
into her mouth and laughs as he seizes her hands.

LUM
You don't need no gum to keep yo' mouf company wid me around. Ahm all
de compny yo' mouf need. Ahm sweet papa chewin and sweetness change.

DAISY
Tell dat to Bootsie Pitts, you cant fool me. (turns right) Guess Ah
better go home and see mama. Ah ain't been round since Ah come from de
white folk. You goin walk round there wid me?

LUM
Naw, Ah aint gointer _walk_. When Ahm wid de angels ah puts on mah
hosanna wings and flies round heben lak de rest. (He falls in beside
her and catches her elbow) Less go! (to the porch) See you later and
tell you straighter.

LINDSAY
Don't stay round to Daisy's too long, Lum, and get run out from under
yo' hat!

LUM
Who run?

HAMBO
Taint no use in you hollerin "who". Yo' feet don't fit no limb.
(General laughter) (Exit Lum and Daisy right)

WALTER
Lawd! Daisy sho is propaganda. She really handles a lot of traffic. Ah
don't blame de boys. If Ah was uh single man Ah'd be round there
myself.

LIGE
Ahm willin tuh serve some time on her gang as it is, but mah wife
won't lissen to reason. (Laughter) Ah tries to show her dis deep point
where taint right for one woman to be harboring uh whole man all to
herself when theres heaps uh po' young girls aint got no husband
atall. But Ah just can't sense her into it.

(Laughter)

HAMBO
Now take Jim and Dave for instant. Here they is, old friends, done
fell out and ready to fight--all over Daisy.

WALTER
Thass me all over. I don't want no partnership when it comes to my
women. Its whole hawg uh none. Lawd, what wimmen makes us do!

LINDSAY
What is it dey don't make us do. Now take for instant Jim Weston. He
know he can't hunt wid Dave--Dave is uh sworn marksman, but jes' so as
not to be outdone here he go trying to shoot turkeys--wild turkeys
mind you, 'ginst Dave.

JOE CLARKE
I God, I hope he finds 'em too. If he get to killin turkeys maybe
he'll stay way from my hen house. I God, I done lost nine uh my best
layin' hens in three weeks.

(General Laughter)

WALTER
Did Jim git em?

CLARKE
I ain't personatin' nobody but I been told dat Jim's got uh powerful
lot uh chicken feathers buried in his back yard. I know one thing if
I ever ketch his toe-nails in my chicken yard, I God, he's gointer
follow his pappy and his four brothers. He's got to git from dis town
of mine.

(Enter a little girl right, very neat and starchy. She runs up to Rev.
Sims.)

GIRL
Papa, mama say send her dat witch hazely oil she sent you after right
quick.

LINDSAY
Whuss matter wid Sister Sims--poly today?

SIMS
She don't keep so well since we been here, but I reckon she's on
de mend.

HAMBO
Don't look like she never would be sick. She look so big and portly.

CLARKE
Size don't mean nothin'. My wife is portly and she be's on de sick
list all de time. It's "Jody, pain in de belly all day. Jody, pain in
de back all night.

LIGE
Besides, Mrs. Simms ain't very large. She wouldn't weigh more'n two
hundred. You ain't seen no big woman. I seen one so big she went to
whip her lil boy an' he run up under her belly and stayed up under
dere for six months.

(General laughter)

WALTER
You seen de biggest ones. But I seen uh woman so little till she could
go out in uh shower uh rain and run between de drops. She had tuh git
up on uh box tuh look over uh grain uh sand.

SIMMS
Y'all boys better read yo' Bibles 'stead of studyin foolishness. (He
gets up and starts into the store. Clarke and the little girl follow
him.) Reckon Ah better git dat medicine. (The three exit into store)

HAMBO
Well, y'all done seen so much--be y'all ain't never seen uh snake big
as de one Ah seen down round Kissimnee. He was so big he couldn't
hardly move his self. He laid in one spot so long he growed moss on
him and everybody thought he was uh log layin' there; till one day Ah
set down on him and went to sleep. When Ah woke up ah wuz in Middle
Georgy.

(General laughter. Two women enter left and go in store after
everybody has spoken to them)

LINDSAY
Layin' all sides to jokes now, y'all remember dat rattlesnake Ah kilt
on Lake Hope was 'most big as dat one.

WALTER
(Nudgin' Lige and winking at the crowd) How big did you say it was, Joe?

LINDSAY
He mought not uh been quite as big as dat one--but jes' bout fourteen
feet.

HAMBO
Gimme dat lyin' snake! He wasn't but fo' foot long when you kilt him
and here you done growed him ten feet after he's dead.

(Enter Simms followed by the girl with an all day sucker. Simms has a
small package in his hand.

SIMMS

(Gives the package to the child and resumes his seat.)

Run 'long home now. Tell yo' ma to put on uh pot uh peas.

(Child exits right trotting and sucking her candy.)

WALTER
They's some powerful big snakes round here. We was choppin' down de
weeds in front of our parsonage yistiddy and kilt uh great big ol'
cotton mouf moccasin.

SIMMS
Yeah, look like me or some of my fambly 'bout to git snake-bit right
at our own front do'.

LIGE
An' bit by uh Baptist snake at dat.

LINDSAY
How you make him out uh Baptist snake?

LIGE
Nobody don't love water lak uh Baptist an' uh Moccasin.

(General laughter)

HAMBO
An' nobody don't hate it lak de devil, uh rattlesnake an uh Meth'dis.

(General laughter. Enter Joe Clark from store. Stands in door)

SIMMS
Dis town needs uh cleanin' in more ways than one. Now if this town was
run right, when folks misbehaves, they oughter be locked up in jail
and if they can't pay no fine, they oughter be made to work it out on
de streets--chopping weeds.

LINDSAY
How we gointer do all dat when we ain't got no jail?

SIMMS
Well, you orta _have_ uh jail. Y'all needs uh whole heap of
improvements in dis town. Ah ain't never pastored no town so way back
as this one here.

CLARKE
(Stepping out before Simms) What improvements you figgers we needs?

SIMMS
A whole heap. Now for one thing, we really does need uh jail, Brother
Mayor. Taint no sense in runnin' people out of town that cuts up. We
oughter have jails like other towns. Every town I ever pastored had
uh jail.

CLARKE
(Angrily) Now hold on uh minute, Simms! Don't you reckon uh man dat
knows how to start uh town knows how to run it? You ain't been here
long enough to find out who started dis town yet. (Very emphatic,
beating of his palm with other fist) Do you know who started dis town?
(Does not pause for an answer) Me! I started _dis_ town. I went to de
white folks and wid _dis_ right hand I laid down two hundred dollars
for de land and walked out and started dis town. I ain't like some
folks--come here when grapes was ripe. I was here to cut new ground.

SIMMS
Well, tain't no sense in one man stayin' Mayor all de time, nohow.

CLARKE
(Triumphantly) So dat de tree you barkin' up? Why, you ain't nothin'
but uh trunk man. You can't be no mayor. I got roots here.

SIMMS
You ain't all de voters, tho, Brother Mayor.

CLARKE
(Arrogantly) I don't hafta be. I God, it's my town and I kin be Mayor
jes' as long as I want to. (Slaps his chest) I God, it was _me_ dat
put dis town on de map.

SIMMS
What map you put it on, Brother Clarke? You musta misplaced it. I
ain't seen it on no map.

CLARKE
Tain't on no map, hunh? I God, everytime I go to Maitland de white
folks calls me Mayor. Otherwise, Simms, I God, if you so dissatisfied
wid de way I run dis town, just take yo' Bible and flat foots and git
younder cross de woods.

SIMMS
(Aggressively) Naw, Ah don't like it. You ack lack tain't nobody in de
corporation but you? Now look. (Points at the street lamp) Tain't but
one street light in town an' you got it in front of yo' place. We pays
de taxes an' you got de lamp.

CLARKE
I God, nobody can't tell me how to run dis town. I 'lected myself and
I'm gonna run it to suit myself. (Looks all about) Where is dat
Marshall? He ain't lit de lamp?

WALTER
Scorched Daisy Blunt home and ain' got back.

CLARKE
I God, call him there, some of you boys.

(Lige steps to edge of porch left and calls "Lum! Lum!" Lum's voice at
a distance: "What!" Lige: "Come on and light de lamp it gittin dark.")

SIMMS
Now, when I pastored in Ocala you oughter seen de lovely jail dey had.

HAMBO
Thass all right for white folks. We colored folks don't need no jail.

WALTER
Aw, yes we do too. Elder Simms is right. We ain't a bit bottern white
folks. (Enter the two women from the store.) You wimmen folks been in
dat store uh mighty long time.

MRS. LULU
We been makin' our market.

HAMBO
Looks mighty bad for some man's pocket. But y'all ain't had no treat
on me. Go back and tell Mrs. Clark tuh give you some candy.

LINDSAY
Have somethin' on me too. Money ain't no good lessen de women kin help
you use it. (Hollers inside) Every lady in there take a treat on me.

MRS. JENNY
Ain't y'all comin' in tuh help us eat de treat. Come on, Elder Simms!

HAMBO
(Getting up quickly. Lindsay and Joe Clarke also get up. They go
inside laughing.) Here, lemme git hold of somebody. (Grabs one of the
women by the arm as they exit into the store.)

LIGE
(Pointing his thumb after the women) Ah wouldn't way lay nothin' lak
dat. Too old even tuh chew peanuts if Ah was tuh buy it.

WALTER
Preach it, Brother. But they's all right for mullet heads like Lindsay
and Hambo. (Sings)

  When they git old, when they [Note: corrected missing space.] git old
  Old folks turns tuh monkeys
      When they git old.

(Looks off right) Lawd! They must be havin' recess in heben! Look at
dese lil ground angels! (Yells off right) Hello Big 'Oman, an' Teets
and Bootsie! Hurry up! My money jumpin' up and down in my pocket lak
uh mule in uh tin stable. (Enter three girls right, dressed in cool
cotton dresses. They are all locked armed and giggling)

LIGE
Hello, folkses.

BOOTSIE
(Coquettishly) Hello yo'self--Want uh piece uh corn bread look on de
shelf. (Great burst of laughter from inside the store)

LIGE
(Catching Bootsie's arm) Lemme scorch y'all inside en' treat yuh.

BOOTSIE
(Looks at the other girls for confirmation) Not yet, after while.

WALTER
Well, come set on de piazza an' les' have some chat.

TEETS
We ain't got time. We come tuh git our mail out de postoffice.

LIGE
Youse uh Got-dat-wrong! You come after Dave an' Jim an' Lum. But Daisy
done treed de las' one of 'em. She got Jim and Dave out in de swamp
where de mule was drugged out huntin' her uh turkey. An' she got Lum
at her house. Thass how come de light ain't lit.

BIG 'OMAN
Oh, Ah ain't worried 'bout Lum. Ah b'lieve Ah kin straighten him out.

WALTER
Some wimmen kin git yo' man so he won't stand uh straightenin'.

LIGE
Don't come rollin' yo' eyes at me an' gittin' all mad cause y'all
stuck on de boys and de boys is stuck on Daisy. (makes a sly face
at Walter)

TEETS
Who? Me? Nobody ain't studyin' 'bout ole Daisy. She come before me
like a gnat in a whirlwind.

WALTER
(in mock seriousness) Better stop dat talkin' 'bout Daisy, do I'll
tell her whut you say. I think I better call her anyhow and see
whether you gointer talk dat big talk to her face. (Makes a move as if
to call Daisy)

LIGE
(keeping up the raillery, grabs Walter) Don't do dat, Walter. We don't
want no trouble round here. But sho nuff, [Note: corrected missing
space.] girls, y'all ain't got no time wid Daisy. Know what Lum say?
Says Daisy is a bucket flower--jes' _made_ him to set up on de
porch an' look pritty. I ast him how 'bout de rest an' he says "Oh de
rest is yard flowers jes' plant them any which a way.

BOOTSIE
I don't b'lieve Lum said no sich uh thing.

LIGE
You tellin' dat flat--Ah knows. (Looks off left) Here come Lum, now,
in uh big hurry jus' lak he ain't been gone two hours.

BIG 'OMAN
Less we all go git our treat! (They start up on the porch. At that
moment Hambo, Lindsay, Clarke, Simms, and the two women enter from
the store.)

CLARKE
(to Lige) Looks here, I God! Ain't Lum lit dat lamp yet? (Enter Lum
left hurriedly. Clarke stands akimbo glaring at him. Lum fumbles for a
match, strikes it and drops it. Gets another from his pocket and goes
to the lamp and strikes it.) Somebody reach de numbskull uh box.
(Walter hands Lum a box of the porch and he gets up on it and opens
the lamp to light it.)

LUM
(to Clarke) Reckon Ah better put some oil in de lamp. Tain't much in
it.

CLARKE
(Impatiently) Oh, that'll do! That'll do. It'll be time tuh put it out
befo' you git it lit, I God.

(Lum lights the lamp. The men have resumed their seats and the women
are on the ground and near right exit. Walter and Lige and the three
girls are at the door about to enter the store. Lum has the box in his
hand and is still under the lamp. He walks slowly towards the step,
box in hand. At the step he looks off left.)

LUM
Here come Dave. (All look left. Walter and Lige and the girls abandon
the idea of the treat and wait for Dave)

HAMBO
But ah ain't seen no turkey yet. Dat ole gobbler's too smart for Dave.

(Enter Dave with gun over his shoulder and holding his head. A little
blood is on his shoulder. He pauses under the lamp a moment then comes
to the step)

HAMBO
Whuss de matter, Dave? Dat ole turkey gobbler done pecked you in de
head? Whut kind of a huntsman is you?

(General laughter)

DAVE
Naw, ain't no turkey pecked me. It's Jim. Ah wuz out in de woods and
hand don squatted down before he got dere. Ah know jus' where dat ole
gobbler roost at. Soon's he hit de limb an' squatted hisself, Ah let
'im have it. He flopped his wings an' tried to fly off but here he
come tumblin' down right by dem ole mule bones. Jim, he was jus'
comin' up when Ah fired. So when he seen dat turkey fallin', whut do
he do? He fires off his gun an' make out he kilt dat turkey. Ah beat
him tuh de bird and we got tuh tusslin'. He tries tuh make _me_ give
him _mah_ turkey so's he kin run tuh Daisy an' make out he done kilt
it. So we got tuh fightin' an' Ah wuz beatin' him too till he retched
down an' got de hock bone uh dat mule an' lammed me over de head an'
fore Ah could git up, he done took mah turkey an' went wid it. (to
Clarke) Mist Clarke Ah wants tuh swear out uh warrant ginst Jim
Weston. Ahm gointer law him outa dis town, too.

SIMMS
Dat wuz uh low-down caper, Jim, cut sho nuff.

CLARKE
Sho its uh ugly caper tuh cut. Come on inside, Dave, an Ah'll make out
de papers. He ain't goin' to carry on lak dat in _my_ town.

(Exit Dave and Clarke into de store)

LINDSAY
(Jokingly to Sims) See whut capers you Meth'dis niggers'll
cut--lammin' folks over de head wid mule bones an' stealin' they
turkeys.

SIMMS
Oh you Baptist ain't uh lot better'n nobody else. You steals an'
fights too.

LINDSAY
(still bantering) Yeah, but we done kotched dis Meth'dis nigger an' we
gointer run him right on outa town too. Jus' wait an' see. Yeah, boy.
Dat Jim'll be uh gone gator 'fore tomorrow night.

WALTER
Oh, I don't know whether he's gointer be gone or not. We Meth'dis got
jus' as much say-so in dis town as anybody else.

LIGE
Yeah. You Baptis run yo' mouf but you don't run de town. Furthermo' we
ain't heard nothin' but Dave's lie. Better wait till we see Jim an'
git de straight of dis thing.

HAMBO
Will you lissen at dat? Dese half-washed Christians hates de truth lak
uh bed-bug hates de light. God a' mighty! (rising) Ahm goin' in an'
see to it dat de Mayor makes dem papers out right. (He exits angrily
into the store. Simms and all the men rise too)

SIMMS
Come on Walter, you an Lige. Less we go inside too. Dat po' boy got
tuh git jestice. An' 'tween de Mayor an' dese Baptists he ain't got
much chance. (They exit into the store)

1st WOMAN
Come on you young gals, whut y'all wanta be hangin' round de store an'
its way after black dark. Yo' mammies oughter take an frail de las'
one of yuh! Come along! (The girls come downoff the porch and join the
women. Loud angry voices inside the store)

2nd WOMAN
Lawd, lemme git home an' tell my husban' bout all dis. Umph! Umph!
(The women and girls exit as the men all emerge from the store. Lum
comes first with the warrant in his hand. Clarke emerges last.)

CLARKE
Can't have all dat fuss an' racket in my store. All of you git outside
dat wants tuh fight? (He begins to close up)

SIMMS
But Brother Mayor, I said it, an' I'll say it agin, tain't right--

CLARKE
(turns angrily) I God, Clarke [Hand written correction: Simms], Ah
don't keer whut you say. 'Taint worth uh hill uh beans nohow. Jim is
gointer be 'rested for hittin' Dave an' takin' his turkey, an' if he's
found guilty he's goin' way from here. Tain't no use uh you swellin'
up neither. (to Lum) Go get him, Lum, an' lock 'im in my barn an' put
dat turkey under arrest too. I God, de law is gointer be law in my
town. (Exit Lum with an important air.)

WALTER
Where de trial gointer be, Brother Clarke, in de hall?

CLARKE
Nope, it's too little. It'll hafta be in de Baptist Church. Ah reckon
dat's de bigges' place in town. Three o'clock Monday evening. Now,
y'all git on off my porch tuh fuss. Lige, outen dat lamp for Lum.

(The stage goes black. The crowd is dispersing slowly. Angry voices
are heard. The curtain is descending slowly. Off-stage right the voice
of Lum is heard calling Daisy.)

LUM
Oh Daisy! Daisy!

DAISY
(at a distance) What you want, Lum?

LUM
Tell yo' mama to put on de hot water kittle. I'll be round there
before long.


_CURTAIN_




ACT II


Scene I

SETTING: Village street scene. Huge oak tree upstage center. A house
or two on backdrop. When curtain goes up Sister Lucy Taylor is seen
standing under the tree trying to read a notice posted on the tree.
She is painfully spelling it out. Enter Sister Thomas--a younger woman
(in her thirties) at left.

SISTER THOMAS
Evenin', Sis Taylor.

SISTER TAYLOR
Evenin'. (returns to the notice)

SISTER THOMAS
Whut you doin'? Readin' dat notice Joe Clarke put up 'bout de meetin'?
(approaches tree)

SISTER TAYLOR
Is dat whut it says? I ain't much on readin' since I had my teeth
pulled out. You know if you pull out dem eye teeth you ruins yo' eye
sight. (turns back to notice) Whut it say?

SISTER THOMAS
(Reading notice) The trial of Jim Weston for assault and battery on
Dave Carter wid a dangerous weapon will be held at Macedonia Baptist
Church on Monday November 10, at three o'clock. All are welcome--by
order of J. Clarke, Mayor of Eatonville, Fla. (turning to Sister
Taylor) Hit's makin' on to three now.

SISTER TAYLOR
You mean its right _now_. (looks up at sun to tell time) Lemme go git
ready to be at de trial--cause I'm sho going to be there and I ain't
goin' to bite my tongue neither.

SISTER THOMAS
I done went and crapped a mess of collard greens for supper--I better
go put em on--cause Lawd knows when we goin' to git outa there--and my
husband is one of them dats gointer eat don't keer whut happen. I bet
if Judgment day was to happen tomorrow, he'd speck I orter fix him a
bucket to carry long.

(She moves to exit right)

SISTER TAYLOR
All men favors they guts, chile. But whut you think of all dis mess
they got going on round here?

SISTER THOMAS
I just think its a sin and a shame before de livin justice de way dese
Baptis' niggers is runnin' round here carryin' on.

SISTER TAYLOR
Oh they been puttin out they brags ever since Sat'day night bout whut
they gointer do to Jim. They thinks they runs this town. They tell me
Rev. Singleton preached a sermon on it yesterday.

SISTER THOMAS
Lawd help us! He can't preach and he look like 10¢ worth of have-mercy,
let lone gittin' up dare tryin' to throw slams at us. Now all Elder
Sims done was to explain to us our rights--Whut you think bout Joe
Clarke running round here takin' up for those ole Baptist niggers?

SISTER TAYLOR
De puzzle-gut rascal--we oughter have him up in conference and put him
out de Meth'dis' faith. He don't blong in there--Wanta run dat boy
outa town for nothin'.

SISTER THOMAS
But we all know how come he so hot to law Jim outa town--hits to dig
de foundation out from under Elder Sims--

SISTER TAYLOR
What he wanta do dat for?

SISTER THOMAS
Cause he wants to be a God-knows-it-all an' a God-do-it-all and Simms
is de onliest one in this town whut will buck up to him.

(Enter Sister Jones, walking leisurely)

SISTER JONES
Hello Hoyt, Hello Lucy.

SISTER TAYLOR
Goin' to de meetin'?

SISTER JONES
Done got my clothes on de line and I'm bound to be dere--

SISTER THOMAS
Gointer testify for Jim?

SISTER JONES
Naw. I reckon--Don't make much difference to me which way de drop
fall--Taint neither one of 'em much good.

SISTER TAYLOR
I know it. I know it, Ida. But dat ain't de point. De crow we wants to
pick is, is we gointer set still and let dese Baptist tell us when to
plant and when to pluck up?

SISTER JONES
Dat _is_ something to think about when you come to think about it.
(starts to move on) Guess I better go ahead--See y'all later and tell
you straighter. (Enter Elder Simms right, walking fast, Bible under
his arm, almost collides with Mrs. Jones. She nods and smiles and
exits.)

ELDER SIMMS
How you do, Sister Taylor, Sister Thomas.

BOTH
Good evenin', Elder

SIMMS
Sho is a hot day

SISTER TAYLOR
Yeah, de bear is walkin' de earth lak a natural man.

SISTER THOMAS
Reverend, look like you headed de wrong way. It's almost time for de
trial and youse all de dependence we got.

ELDER SIMMS
I know it. I'm trying to find de Marshall so we kin go after Jim. I
wants a chance to talk wid him a minute before court sits.

SISTER TAYLOR
Y'think he'll come clear?

ELDER SIMMS
(proudly) I _know_ it! (shakes the Bible) I'm going to law 'em from
Genesis to Revelation.

SISTER THOMAS
Give it to 'em, Elder. Wear 'em out!

ELDER SIMMS
We'se liable to have a new Mayor when all dis dust settle. Well, I
better scuffle on down de road.

(Exit Sims left)

SISTER THOMAS
Lord, lemme gwan home and put dese greens on. (looks off stage left)
Here come Mayor Clark now, wid his belly settin' out in front of him
like a cow-catcher. His name oughter be Mayor Belly.

SISTER TAYLOR
(akimbo) Jus' look at him! Trying to look like a jigadier Breneral.

(Enter Clarke hot and perspiring. They look at him coldly.)

CLARKE
I God, de bear got me! (silence for a moment) How y'all feelin'
ladies?

SISTER TAYLOR
Brother Mayor, I ain't one of these folks dat bite my tongue and bust
my gall--Whuts inside got to come out! I can't see to my rest why you
cloakin' in wid dese Baptist buzzards ginst yo' own Church.

MAYOR CLARKE
I ain't cloakin' in wid _none_. I'm de Mayor of dis whole town.
I stands for de right and against de wrong. I don't keer who it kill
or cure.

SISTER THOMAS
You think it's right to be runnin' dat boy off for nothin?

MAYOR CLARKE
I God! You call knockin' a man in de head wid a mule bone nothin'?
'Nother thing--I done missed nine of my best-layin' hens. I ain't
sayin' Jim got 'em--but different people has told me he buries a
powerful lot of feathers in his back yard. I God, I'm a ruint man! (He
starts towards the right exit, but Lum Rogers enters right.) I God,
Lum, I been lookin' for you all day. It's almost three o'clock. (hands
him a key from his ring) Take dis key and go fetch Jim Weston on to de
church.

LUM
Have you got yo' gavel from de lodge-room?

CLARKE
I God, that's right, Lum. I'll go get it from de lodge room whilst you
go git de bone an' de prisoner. Hurry up! You walk like dead lice
droppin' off you! (He exits right while Lum crosses stage towards
left)

SISTER TAYLOR
Lum, Elder Simms been huntin' you--he's gone on down bout de barn.
(She gestures.)

LUM
I reckon I'll overtake him. (Exit left)

SISTER THOMAS
I better go put dese greens on--my husband will kill me if he don't
find no supper ready. Here come Mrs. Blunt. She oughter feel like a
penny's worth of have-mercy wid all dis stink behind her daughter.

SISTER TAYLOR
Chile, some folks don't keer. They don't raise they chillen, they
drags 'em up. God knows if dat Daisy was mine, I'd throw her down and
put a hundred lashes on her back wid a plow-line. Here she come in de
store Sat'day night (acts coy and coquettish, burlesques Daisy's walk)
a wringing and a twisting!

(Enter Mrs. Blunt left.)

MRS. BLUNT
How y'all sisters?

SISTER THOMAS
Very well, Miz Blunt, how you?

MRS. BLUNT
Oh so-so.

SISTER TAYLOR
I'm kickin' but not high.

MRS. BLUNT
Well, thank God you still on prayin' ground and in a Bible
Country--Me, I ain't many today. De niggers got my Daisy's name all
mixed up in diss mess.

SISTER TAYLOR
You musn't mind dat, Sister Blunt. People just _will_ talk. They's
talkin' in New York and they's talkin' in Georgy and they's talkin' in
Italy.

SISTER THOMAS
Chile, if you talk after niggers they'll have you in de graveyard or
in Chattahoochee one. You can't pay no tention to talk.

MRS. BLUNT
Well, I know one thing--de man or woman, chick or child, grizzly or
gray that tells me to my face anything wrong bout _my_ chile--I'm
going to take _my_ fist (rolls up right sleeve and gestures with right
fist) and knock they teeth down they throat. (She looks ferocious.)
Cause y'll know I raised my Daisy right round my feet till I let her
go up north last year wid them white folks. I'd ruther her to be in de
white folks kitchen than walkin' de streets like some of dese girls
round here. If I do say so, I done raised a lady. She can't help it if
all dese men get stuck on her.

SISTER TAYLOR
You'se telling de truth, Sister Blunt--that's what I always say--Don't
confidence dese niggers, do they'll sho put you in de street.

SISTER THOMAS
Naw indeed. Never syndicate wid niggers--do--they will distriminate
you. They'll be an _anybody_. You goin to de trial, ain't you?

MRS. BLUNT
Just as sho as you snore, and they better leave Daisy's name outer dis
too. I done told her and told her to come straight home from her work.
Naw, she had to stop by dat store and skin her gums back wid dem
trashy niggers. She better not leave them white [Corrected missing
space.] folks today to come praipsin over here scornin her name all up
wid dis nigger mess--do, I'll kill her. No daughter of mine ain't
going to do as she please long as she live under de sound of my voice.
(She crosses to right.)

SISTER THOMAS
That's right, Sister Blunt--I glory in yo' spunk. Lord, I better go
put on my supper. (As Mrs. Blunt exits right, Rev. Singletary enters
left with Dave and Deacon Lindsay and Sister Lewis. Very hostile
glances from Sisters Thomas and Taylor towards the others.

ELDER SINGLETARY
Good evening, folks.

(Sister Thomas and Sister Taylor just grunt. Sister Thomas moves a
step or two towards exit. Flirts her skirts and exits.)

LINDSAY
(Angrily) Whuts de matter, y'all? Cat got yo' tongue?

SISTER TAYLOR
More matter than you kin scatter all over Cincinnatti.

LINDSAY
Go head on, Lucy Taylor, go head on. You know a very little of yo'
sugar sweetens my coffee. Go head on. Everytime you lift yo' arm you
smell like a nest of yellow hammers.

SISTER TAYLOR
Go head on yo'self. Yo' head look like it done wore out three
bodies--talking bout _me_ smelling--you smell lak a nest of grand
daddies yo'self.

LINDSAY
Aw, rack on down de road, 'oman. Ah don't wantuh change words wid yuh.
You'se too ugly.

MRS. TAYLOR
You ain't nobody's pretty baby yo'self. You so ugly I betcha yo' wife
have to spread uh sheet over yo' head tuh let sleep slip up on yuh.

LINDSAY
(Threatening) You better git 'way from me while you able. I done tole
you I don't wants break a mouth wid you. It's a whole heap better tuh
walk off on yo own legs than it is to be toted off. I'm tired of yo'
achin round here. You fool wid me now an' I'll knock you into doll
rags, Tony or no Tony.

SISTER TAYLOR
(jumping up in his face) Hit me! Hit me! I dare you tuh hit me. If you
take dat dare you'll steal a hawg an' eat his hair.

LINDSAY
Lemme gwan down to dat church befo' you make me stomp you.
(He exits right.)

SISTER TAYLOR
You mean you'll _git_ stomped. Ahm going to de trial too. De nex trial
gointer be _me_ for kickin some uh you Baptis niggers around.

(A great noise is heard off stage left. The angry and jeering voices
of children. Mrs. Taylor looks off left and takes a step or two
towards left exit as the noise comes nearer.)

VOICE OF ONE CHILD
Tell her! Tell her! Turn her up and smell her. Yo' mama ain't got
nothin to do wid me.

SISTER TAYLOR
(Hollering off left) You lil Baptis haitians, leave them chillun
alone. If you don't, you better!

(Enter about 10 chidren struggling and wrestling in a bunch. Mrs.
Taylor looks about on the ground for a stick to strike the children
with.)

VOICE OF CHILD IN CROWD
Hey! Hey! He's skeered tuh knock it off. Coward!

SISTER TAYLOR
If y'all don't git on home!

SASSY LITTLE GIRL
(Standing akimbo) I know you better not touch me, do my mama will tend
to you.

SISTER TAYLOR
(Making as if to strike her) Shet up, you nasty lil heifer, sassing
me! You ain't half raised.

(The little girl shakes herself at Mrs. Taylor and is joined by two or
three others.)

SISTER TAYLOR
(Walking towards right exit) I'm going on down to de church an' tell
yo' mammy. But she ain't been half raised herself. (She exits right
with several children making faces behind her.)

A BOY
(to sassy girl) Aw haw! Y'all ol' Baptis ain't got no book case in yo'
church. We went there one day an' I saw uh soda cracker box settin' up
in de corner so I set down on it. (pointing at sassy girl) Know whut
ole Mary Ella say? (jeering laughter) Willie, you git up off our
library! Haw! Haw!

MARY ELLA
Y'all ole Meth'dis' ain't got no window panes in yo' ole church.

A GIRL
(Takes center of stage and hands akimbo shakes her hips.) I don't keer
whut y'allsay. I'm a Methdis' bred an' uh Methdis' born an' when I'm
dead there'll be uh Methdis' gone.

MARY ELLA
(snaps fingers under other girl's nose and starts singing. Several
join her.)

  Oh Baptis, Baptis is my name
  My name's written on high
  I got my lick in de Baptis church
  Gointer eat up de Methdis pie

(the Methodist children jeer and make faces. The Baptist camp make
faces back for a full minute there is silence while each camp tries to
outdo the other in face making. The Baptist makes the last face.

METHODIST BOY
Come on, less us don't notice em. Less gwan down to de church an' hear
de trial.

MARY ELLA
Y'all ain't the onliest ones kin go. We goin' too.

WILLIE
Aw Haw! Copy cats! (Makes face) Dat's right, follow on behind us lak
uh puppy dog tail. (They start walking toward right exit switching
their clothes behind.)

(Baptist children stage a rush and struggle to get in front of the
methodists. They finally succeed in flinging some of the Methodist
children to the ground and some behind them and walk towards right
exit haughtily switching their clothes.)

WILLIE
(whispers to his crowd) Less go round by Mosely's lot and beat 'em
there!

OTHERS
All right!

WILLIE
(Yelling to Baptists) We wouldn't walk behind no ole Baptists! (The
Methodists turn and walk off towards left exit switching their clothes
as the Baptists are doing.)


_SLOW CURTAIN_




ACT II


SCENE II

SETTING: Interior of Macedonia Baptist Church, a rectangular room,
windows on each side, two "Amen Corners", pulpit with a plush cover
with heavy fringe, practical door in pulpit, practical door in front
of church, two oil brackets with reflectors on each side wall with
lamps missing all but one, one big oil lamp in center.

ACTION: At the rise, church is about full. A buzz and hum fills the
church. Voices of children angry and jeering heard from the street.
The church bell begins to toll for death. Everybody looks shocked.

SISTER LEWIS
Lawd! Is Dave done died from dat lick?

SISTER THOMAS
(to her husband) Walter, go see. (He gets up and starts down the aisle
to front door. Enter Deacon Hambo by front door.)

WALTER
Who dead?[Note: correction to e]

HAMBO
(laughing) Nobody--jus' tollin' de bell for dat Meth'dis gopher dat's
gointer be long long gone after dis trial. (laughter from the Baptist
side)

WALTER
Y'all sho thinks you runs dis town, dontcher? But Elder Simms'll show
you somethin' t'day. If he don't, God's uh gopher.

HAMBO
He can't show us nothin' cause he don't know nothin' hisself.

WALTER
He got mo' book-learnin' than Rev. Singletary got.

HAMBO
He mought be unletter-learnt, but he kin drive over Sims like a
road plow.

METHODIST CHORUS
Aw, naw! Dat's a lie!

(Enter Rev. Simms by front door with open Bible in hand. A murmur of
applause arises on the Methodist side, grunts on the Baptist side.
Immediately behind him comes Lum Boger leading Jim Weston. They parade
up to the right Amen Corner and seat themselves on the same bench, Jim
between the Marshall and the preacher. A great rooster crowing and hen
cackling arises on the Baptist side. Jim Weston jumps angrily to his
feet.)

(Enter by front door Rev. Singletary and Dave. Dave's head is
bandaged, but he walks firmly and seems not ill at all. They sit in
the left Amen Corner. Jeering grunts from the Methodist side.)

SISTER THOMAS
Look at ol' Dave trying to make out he's hurt.

LIGE
Everybody know uh Baptis' head is hardern uh rock. Look like they'd be
skeered tuh go in swimmin', do they heads would drown 'em. (general
laughter on Methodist side)

(Enter Bro. Nixon with his jumper jacket on his arm and climbs over
the knees of a bench full of people and finds seat against the wall
directly beneath empty lamp bracket. He looks around for some place to
dispose of his coat. Sees the lamp-bracket and hangs up the coat,
hitches up his pants and sits down.)

SISTER LEWIS
(rising and glaring at Nixon) Shank Nixon, you take yo' lousy coat
down off these sacred walls. Ain't you Methdis' niggers got no
gumption in de house of Wash-up!

(Nixon mocks her by standing akimbo and shaking himself like a woman.
General laughter. He prepares to resume his seat but looks over and
sees Deacon Hambo on his feet, and glaring angrily at him. He quickly
reaches up and takes the coat down and folds it across his knees.)

(Sister Taylor looks very pointedly at Sister Lewis then takes a dip
of snuff and looks sneering at Lewis again.)

SISTER TAYLOR
Some folks is a whole lot more keerful bout a louse in de church than
[Note: corrected missing space] they is in they house. (Looks
pointedly at Sister Lewis.)

SISTER LEWIS
(bustling) Whut you gazin' at me for? Wid your pop-eyes looking like
skirt ginny-nuts.

SISTER TAYLOR
I hate to tell you whut yo' mouf looks like. I sho do you and soap and
soap and water musta had some words.

SISTER LEWIS
Talkin' bout other folks being dirty--yo' young 'uns must be sleep in
they draws cause you kin smell 'em a mile down de road.

SISTER TAYLOR
Taint no lice on 'em though.

SISTER LEWIS
You got just as many bed-bugs and chinches as anybody else, don't come
trying to hand me dat rough package bout yo' house so clean.

SISTER TAYLOR
Yeah, but I done seen de bed-bugs munchin' out yo' house in de
mornin', keepin' step just like soldiers drillin'. An' you got so many
lice I seen em on de dish-rag. One day you tried to pick up de
dish-rag and put it in de dish water and them lice pulled back and
tole you "Aw naw, damned if I'm going to let you drown me." (Loud
laughter from the Methodist side)

SISTER LEWIS
(furious--rises akimbo) Well, my house might not be exactly clean, but
there's no fly-specks on my character! They didn't have to sit de
sheriff to make Willie marry _me_ like they did to make Tony marry
_you_.

SISTER TAYLOR
(Jumping up and starts across the aisle. She is pulled back out of the
aisle by friends.) Yeah, they got de sheriff to make Tony marry me,
but he married me and made me a good husband, too. I sits in my
rocking cheer on my porch every Sat'day evening and say "here come
Tony and them--

SISTER LEWIS
Them what?

SISTER TAYLOR
Them dollars. Now you sho orter go git de sheriff and a shot-gun and
make some of dese men marry yo' daughter Ada.

SISTER LEWIS
(Jumping up and starting across the aisle. She is restrained, but
struggles hard.) Lemme go, Jim Merchant! Turn me go! I'm going to
stomp de black heifer till she can't sit down.

SISTER TAYLOR
(Also struggling) Let her come on! If I get my hands on her I'll turn
her every way but loose.

SISTER LEWIS
Just come on out dis church, Lucy Taylor. I'll beat you on everything
you got but yo' tongue and I'll bit dat a lick if you stick it out.
(to the men holding her) Turn me go! I'm going to fix her so her own
mammy won't know her. She ain't going to slip _me_ into de dozens and
laugh about it.

SISTER TAYLOR
(Trying to free herself) Why don't y'all turn dat ole twist mouth
'oman loose. All I wants to do is hit her one lick. I betcha I'll take
her 'way from here faster than de word of God.

SISTER LEWIS
(to men holding Mrs. Taylor) I don't see how come y'all want let ole
flat-behind Lucy Taylor aloose--make out she so bad, now. She may be
red hot but I kin cool her. I'll ride her just like Jesus rode a
jackass.

(They have subsided into their seats again, but are glaring at each
other. Enter Mayor Clarke thru the pulpit door and is annoyed at the
clamor going on. He tries to quell the noise with a frown.)

SISTER TAYLOR
Dat ain't nothin' but talk--You looks lak de Devil before day, but you
ain't so bad--not half as bad as you smell.

CLARKE
Order, please. Court is set.

SISTER LEWIS
You looks like all hell and de devil's doll baby, but all I want _you_
to do is to hit de ground and I'll crawl you. Put it where I kin git
it and I'll sho use it.

MAYOR CLARKE
(feeling everywhere for the gavel) Lum Boger! Where's dat gavel I told
you to put here?

LUM
(from beside prisoner) You said _you_ were going to git it yo'self.

CLARKE
I God, Lum, you gointer stand there like a bump on a log and see I
ain't got nothin' to open court wid? Go head--fetch me dat gavel. Make
haste quick before dese wimmen folks tote off dis church house. (Lum
exits by front door)

SISTER TAYLOR
(to Lewis) Aw, shut up, you big old he-looking rascal you! Nobody
don't know whether you'se a man or a woman.

CLARKE
You wimmen, shut up!

SISTER LEWIS
(to Taylor) Air Lawd! Dat ain't _yo_ trouble. They all _knows_ whut
_you_ is--eg-zackly!

LINDSAY
Aw, why don't you wimmen cut dat out in de church-house! Jus' jawin'
and chewin' de rag!

SISTER TAYLOR
Joe Lindsay, if you'd go home and feed dat raw-boned horse of yourn
you wouldn't have so much time to stick yo' bill in business that
ain't yourn.

LINDSAY
You ain't got nairn to feed--You better go hunt another dead dog and
git some mo' teeth. Great big ole empty mouf, and no cheers in de
parlar.

SISTER TAYLOR
I kin git all de teeth I wants--I'd ruther not have no cheers in my
parlor than to have them ole snags you got in yo' mouf. I'd ruther gum
it out.

LINDSAY
You don't _ruther_ gum it out, you _hafta_ gum it out. You ain't got
no teeth. Dey better send out to dat ole mule and git you some
teethes.

SISTER LEWIS
Joe Lindsay, don't you know no better than to strain wid folks ain't
got sense enough to tote guts to a bean? If they ain't born wid no
sense you cna't learn 'em none.

LINDSAY
You sho done tole whut God love now. (Glaring across the aisle) Ain't
got enough gumption to kill a buzzard.

(Enter Lum by front door with gavel in one hand and mule bone in the
other. He walks importantly up the aisles and hands Clarke the gavel
and lays the bone atop the pulpit.)

CLARKE
(rapping sharply with gavel) Here! You moufy wimmen shut up. (to Lum)
Lum, go on back there and shut dem wimmen up or put 'em outa here.

(Lum starts walking importantly down the aisle towards Sister Taylor.
she almost rises to meet him.)

SISTER TAYLOR
Lum Boger, you fresh little snot you! Don't you dast to come here
trying to put _me_ out--Many diapers as I done pinned on _you_! Git
way from me befo' I knock every nap off of yo' head, one by one.

(Lum hurries away from her apologetically. He turns towards Mrs.
Lewis.)

MRS. LEWIS
Deed Godknows you better not lay de weight of yo' hand on _me_, Lum.
Here you ain't dry behind de ears yert and come telling _me_ what to
do. Gwan way from here before I kick yo' clothes up round you' neck
like a horse collar.

(Lum goes on back and takes his seat beside the prisoner.)

CLARKE
(glaring ferociously) This court is set and I'm bound to have some
order or else. (The talking ceases. Absolute quiet)

CLARKE
Now less git down to business. We got folks in dis town dat's just
like a snake in de grass.

SISTER BOGER
Brother Mayor! We ain't got no business going into no trial nor
northin' else 'thout a word of prayer--to be sure de right spirit is
wid us.

VOICE ON METHODIST SIDE
Thass right,--Elder Simms, give us a word of prayer. (He rises
hurriedly.)

VOICE ON BAPTIST SIDE
This is a Baptist Church and de pastor is settin' right here--how come
he can't pray in his own church?

VOICE ON METHODIST SIDE
Y'all done started all dis mess--how you going to git de right spirit
here? Go head, Rev. Simms.

VOICE ON BAPTIST SIDE
He can't pray over me. Dis Church says one Lord, one faith, one
Baptism--and a man that ain't never been baptised atall ain't got no
business praying over nobody.

CLARKE
(rapping with gavel) Less sing! Somebody raise a tune.

(VOICE ON BAPTIST SIDE begins "Onward Christian Soldiers" and the
others join in.)

(VOICE ON METHODIST SIDE begins "All hail the power of Jesus name" and
the Methodists join in. Both shout as loud as they can to the end of
the verse.)

(Mayor Clarke raps loudly for order at the end of the verse and lifts
his hands as if to bless a table)

CLARKE
(praying) Lord be withus and bless these few remarks we are about to
receive, Amen. Now this court is open for business. All of us know we
came here on serious business. This town is bout to be tore up by
back-biting and malice. Now everybody that's a witness in this case
stand up. I wants the witness to take the front seat.

(Nearly everybody in the room rises. Brother Hambo frowns across the
aisle at Mrs. McDuffy, who is standing.)

BROTHER HAMBO
Whut _you_ doing standin' up for a witness? I know you wasn't there.
You don't know one thing about it.

SISTER McDUFFY
I got just as much right to testify as you is. I don't keer if I
wasn't there. Any man that treat they wife bad as _you_ can't tell
nobody else they eye is black. You clean round yo' _own_ door before
you go sweeping round other folks.

SISTER LINDSAY
(to Nixon) What you doin' up there testifying? When you done let yo'
hawg root up all my p'tater patch.

NIXON
Aw shut up woman--You ain't had no taters for no pit to root up.

SISTER LINDSAY
Who ain't had no taters? (To Lige) Look here, Lige, didn't I git a
whole crokus sack full of tater slips from yo' brother Sam?

LIGE
(reluctantly) Yeah.

SISTER LINDSAY
Course I had sweet p'taters! And if you stand up there and tell _me_ I
ain't had no p'taters I'll be all over you just like gravy over rice.

NIXON
Aw shut up--We ain't come here to talk about yo' tater vines, we come--

SISTER LINDSAY
(to her husband) Joe! What kind of a husband is you? Set here and let
Nixon 'buse me out lak dat!

WALTER
How is he going to give anybody a straightening when he needs
straightening hisself. I bought a load of compost from him and _paid
for it in advance_ and he come there when I wasn't home and dumped a
half-a-load in there and drove on off wid my money.

SISTER HAMBO
Aw, you ain't got no right to talk, Walter, not low down as you is--if
somebody stump their toe in dis town you won't let yo' shirt-tail
touch you till you bolt over to Maitland and puke yo' guts to de white
folks--and God knows I 'bominates a white folks nigger.

WALTER
Aw you just mad cause I wouldn't let your old starved-out cow eat up
my cow-peas.

SISTER HAMBO
(triumphantly) Unhumh! I knowed you was the one knocked my cow's horn
off! And you lied like a doodle-bug going backwards in his hole and
made out you didn't do it.

WALTER
I didn't do no such a thing.

SISTER HAMBO
I say you did and belong to Macedonia Baptist Church and I can't lie.

WALTER
Yo' mouf is cut cross ways, ain't it? Well then, yo' mouf ain't no
prayer-book even if yo' lips do flap like a Bible. You kin lie and
then re-lie.

DEACON HAMBO
Walter Thomas talk dat biggity talk to me, not to my wife. Maybe you
kin whip her, but if you can't whip me too, don't bring de mess up.

CLARKE
(rapping) Y'all men folks shut up before I put you both under arrest.
Come to order everybody.

LINDSAY
I just wants say this before we go any further. Nobody bet not slur my
wife in here--do I'll strow 'em all over de county.

MRS. NIXON
Aw, youse de nastiest threatener in three states but I ain't seen you
do nothin'. De seat of yo' pants is too close to de ground for you to
be crowin' so loud. You so short you smell right earthy.

MRS. LINDSAY
De seat of yo' husband's britches been draggin' de ground ever since I
knowed him. Don't like it dontcher take it, here's my collar come and
shake it. (She puts the palms of her hands together and holding the
heels together, flaps the fore part of her hands like a gator opening
and shutting its mouth. This infuriates Mrs. Nixon.

CLARKE
Shut up! We didn't come here to wash and iron niggers. We come here
for a trial. (raps)

MRS. NIXON
(to Clarke) I ain't going to shut up nothin' of de kind. Think I'm
going to let her low-rate me and I take it all? Naw indeed. I'm going
to sack dis female out before we any further go.

MRS. LINDSAY
Aw, I done dished you out too many times. Go head on and try to keep
yo' lil squatty husband away from down on de lake wid wimmens and
you'll have _all_ you can do. How does old heavy-hipted mama talk?

(snaps her fingers)

MRS. NIXON
Nobody wouldn't have you if he could get anybody else.
(She makes a circle with her thum and first finger and
holds it up for Mrs. Lindsay to see.) Come thru--don't
you feel cheap?

CLARKE
Sister Nixon, shut up!

SISTER NIXON
You can't shut me up, not the way you live. When you quit beatin Mrs.
Mattie and dominizing her all de time then you kin tell other folks
what to do. You ain't none of my boss. Don't let yo' wooden God and
corn-stalk Jesus fool you now. Now de way you sells rancid bacon for
fresh.

NIXON
Aw, honey, hush a while, please and less git started.

(A momentary quiet falls on the place. Mayor glowers all over the
place. Turns to Lum.)

CLARKE
Lum, git a piece of paper and a pencil and take de names of all de
witnesses _who was dere while de fight was going on_.

LUM
(Pulling a small tablet and pencil out of his coat pocket) I brought
it with me.

CLARKE
Now everybody who was at de fight hold up yo' hands so Lum can know
who you are.

(Several hands go up. Sister Anderson puts up her hand.)

CLARKE
You wasn't there, Sister Anderson, not at that time.

SISTER ANDERSON
I hadn't been gone more'n ten minutes 'fore Dave come in from de
woods.

CLARKE
But you didn't see it.

SISTER ANDERSON
It don't make no difference--my husband heered every word was spoke
and told me jes' lak it happen. Don't tell _me_ I can't testify.

DEACON HAMBO
Nobody can't testify but de two boys cause nobody wuz at de fight but
dem.

SISTER ANDERSON
Dat's all right too, Brother, but I know whut they wuz fightin' about
an' it wudn't no turkey neither. It wuz Daisy Blunt.

MRS. BLUNT
Just you take my chile's name right out yo' mouf, Becky Anderson. She
wuznt out in dat cypress swamp. Leave her out dis mess.

REV. SIMMS
You ain't got no call to be so touchous bout yo' girl, but you sho
said a mouthful, Sister Blunt. Dis sho is a mess. Can't help from
being uh mess. (glares at Mayor) Holdin' a trial in de Baptist Church!
Some folks ain't got sense enough todo 'em till four o'clock and its
way after half past tree right now.

MAYOR
Shet up, dere, Simms! Set down! Who ast yo' pot to boil, nohow! Court
is de best church they is, anyhow, cause you come in court. You better
have a good experience and a strong determination. (raps vigorously)
Now lemme tell _y'all_ something. When de Mayor sets Court--don't keer
when I sets it nor where I sets it, you got to git quiet and stay
quiet till I ast you tuh talk. I God, you sound lak a tree full uh
blackbirds! Dis ain't no barbecue, nor neither no camp meetin'. We
'sembled here tuh law uh boy on a serious charge. (A great buzz rises
from the congregation. Mayor raps hard for order and glares all about
him.) Hear! Hear! All of us kin sing at de same time, but can't but
one of us talk at a time. I'm doin' de talkin' now, so de rest of you
dry up till I git through. I God, you sound lak uh passle uh dog
fights! We ain't here for no form and no fashion and no outside show
to de world. Wese here to law. (to Lum) You done got all de witnesses
straight--Got they names down?

LUM
Yessuh, I got it all straightened out.

CLARKE
Well, read de names out and let de witnesses take de front seats.

LUM
Mr. Clarke, I done found out nobody wasn't at dat fight but Jim and
Dave and de mule bones. Dere's de bone Dave got hit wid up on de
rostrum and deres Jim and Dave in de Amen Corners.

DAVE
(rising excitedly) Mist' Clarke! Brother Mayor, I wants to ast uh
question right now to git some information.

MAYOR
All right, Dave, go head and ast it.

DAVE
Brother Mayor, I wanted to know whut become of my turkey gobbler?

MAYOR
I God, Dave, youse in order. Lum! I God, I been layin' off to ast you
whut you dont wid dat turkey. Where is it?

(A burst of knowing laughter from the house)

LUM
(very embarrassed) Well, when you tole me to go 'rrest Jim and de
turkey, I took and went on round to his ma's house and he wudnt dere
so I took and turnt round and made it t'wards Daisy's house an' I
caught up wid him under dat China-berry tree jest befo' you gits tuh
Daisy's house. He was makin' it on t'wards her house wid de turkey in
one hand--his gun crost his shoulder when I hailed 'im. I hollered
"Jim, hold on dere uh minute!" He dropped de turkey and wheeled and
throwed de gun on me.

MAYOR CLARKE
I God, he drawed uh gon on de City Marshall?

LUM
Yessir! He sho did. Thought I was Dave. Tole me: "Don't you come
another step unless you want to see yuh Jesus." I hollered back "It's
me, I ain't no Dave Carter." So he took de gun offa me and I went up
to him and put him under arrest, and locked him up in yo' barn and
brought _you_ de key, didn't I?

CLARKE
You sho did, but I God, I ast you whut become of de turkey?

LUM
De turkey wasn't picked or nothin', so I put him under 'rrest too,
jus' lak you tole me. (general laughter)

CLARKE
I God, Lum, whut did you _do_ wid de turkey after you put him under
'rrest?

LUM
Jim, he didn't want to come wid me till he could make it to Daisy's
house to give her det turkey but, bein so close up on him till he
couldn't draw his rifle, I throwed my 32:20 in his face an' tole him I
said "Don't you move! Don't you move uh pig do I'll burn you down! I
got my burner cocked dead in yo' face and I'll keer you down jus' lak
good gas went up. Come on wid me!" So I took his rifle and picked up
de turkey and marched him off to yo' cow-lot. Ast him didn't I do it.
I tole him, I said "I know you Westons goes for bad but I'm yo' match.
I said you may be slick but you kin stand another greasing. Now sir! I
ain't skeered uh nobody. I'll put de whole town under 'rrest.

MAYOR CLARKE
I God, Lum, if you don't tell me whut you done wid dat turkey, you
better! (draws back the gavel as if to hurl it at Lum) I'll lam you
over de head wid dis mallet! Whut did you do wid dat gobbler turkey?

LUM
Being as he wasn't picked or nothin', I know you didn't want to be
bothered wid it, so I took and carried it over to Mrs. Blunt's house
and she put on some hot water and we set up way Sat'day night pickin
de turkey and fixin him so nex' day she cooked him off--just sorta
baked him wid a lil stuffin an' such, so he'd keep.

MAYOR CLARKE
Didn't you know my wife knowed how to cook? Go fetch dat turkey here,
and don't let no dead lice fall off of you on de way.

LUM
(extremely embarrassed) I don't speck he's dere now, Mist' Clarke.

CLARKE
(ferociously) How come?

LUM
I passed by dere on Sunday and et a lil piece of shoulder offa him,
an' being everybody else was eatin' turkey too, I et some breast meat
an' uh mouf ful or two of stuffin' an' uh drum stick wid de ham part
of de leg hung on to it wid a lil gravy. (general laughter) I thought
I was doin' right cause [Note: corrected missing space] de turkey was
kilt for Daisy anyhow. So I jus' took it on to her. Dave was all hurt
up and Jim was locked up so--

CLARKE
Dat'll do! Dat'll do! Dry up, Suh! (turns to Dave) Stand up, Dave.
Since youse de one got hurted, you be de first witness and tell me
just whut went on out dere.

(Dave rises slowly.)

SISTER TAYLOR
Dat's right, Dave. Git up dere and lie lak de cross ties from New York
to Texas. You greasy rascal you! You better go wash yo'self before you
go testifying on people.

DAVE
I'm just as clean as you.

REV. SINGLETARY
(jumping to his feet) Wait a minute! Taint none of y'all got no call
to be throwin' off on dis boy. He come here to git justice, not to be
slurred and low-rated. He ain't 'ssaulted nobody. He ain't stole no
turkeys _nor_ chickens. He's a clean boy. He set at my feet in Sunday
school since he was so high, (measures knee height) and he come thru
religion under de sound of my voice an' I baptized him and I know he's
clean.

SISTER TAYLOR
It'll take more'n uh baptizin' to clean dat nigger.

DAVE
I goes in swimmin' nearly every day. I'm just as clean as anybody
else.

SISTER TAYLOR
(Mayor begins rapping for order. She shouts out) Swimmin! Dat ain't
gointer clean de crust offa _you_. You ain't had a good bath since de
devil was a hatchet. If you ain't been parboiled in de wash pot and
scoured wid Red Seal lye, don't bring de mess up.

CLARKE
I'm goin' to have order here or else! Gwan, Dave.

DAVE
It's just lak I tole you Sat'day night.

CLARKE
Yeah, but dat wuz at de store. Dis is in [Note: corrected missing
space] court and it's got to be tole agin.

ELDER SIMMS
Just uh minute, Brother Clarke, before we any further go I wants to
ast de witness uh question dat oughter be answered before he open his
mouf.

MAYOR CLARKE
Whut _kind_ of a question is dat?

SIMMS
Dave, tell de truth. Ain't yo' heart full of envy and malce 'gainst
dis chile? (Gestures towards Jim. Dave shakes his head and starts to
deny the charge but Simms hurries on.) Wait a minute now! Wait till I
git thru. Didn't y'all used to run around everywhere playin' and
singing andeverything till you got so full of envy and malce and
devilment till y'al broke up? Now, Brother Mayor, make him tell de
truth.

DAVE
Yeah, I useter be crazy bout Jim, and we was buddies till he tried to
back bite me wie, wid my girl.

JIM
Never _was_ yo girl. Nohow I ain't none of yo' buddy. I ain't got no
buddy. They kilt my buddy tryin' to raise me. But I did useter lak you
till you acted so low down tryin' to undermine me and root me out wid
my girl.

MAYOR
Aw, table dat business an' less open up new business. We ain't here to
find out whose girl it is. We wants to know 'bout dis fight and who
hit de first lick and how come. Go head on Dave and talk.

DAVE
Well, jus lak I tole yuh, Sat'day night, I been watchin' dat flock uh
wild turkeys ever since way last summer roostin' in de edge of dat
cypress swamp out by Howell Creek, where Brazzle's ole mule was
dragged out. It was a great, bit ole gobbler leadin' de flock. So last
time I seen him I said I was gointer git him for my girl if it taken
me uh year. So Sat'day, kinda late, I grabs ole Hannah, my gun, I
calls her Old Hannah, and come to de store to buy some shells. Y'all
know whut went on at de store. Well, it made me feel lak I wuz
gointergit dat ole gobbler if I had to follow him clean to Diddy war
Diddy or slap into Ginny-Gall. But I didn't have to do nothin'. When I
got out by de ole mule bones, I seen 'em flyin' round lak buzzards. So
I loaded both barrels, squatted down on uh log where I had dead aim on
dat big ole cypress pine where they roosts at. Sho nuff, soon's de sun
had done set, here dey come followin' de leader'. He lit way out on de
end of de limb kinda off from de rest and I eased ole Hannah up on
him. Man! I got so skeered I wuz gointer miss him till I got de all
overs. He gobbled two three times to see if all his fambly was safed
den he settled down and bam! I let him have it! He spread his wings
lak he wuz gointer fly on off an' I _cried_ lak a chile! But I got him
alright and down he come floppin, and me grabbin him before he quit
kickin. Gee, I was proud. He felt lak he weighed forty pounds. Whilst
I was kinda heftin him in my hands I heard uh rifle fire and I looked
and dere was Jim firin into de turkey flock dat was flyin round
skeered. He didn't hit a God's thing, but he seen me wid my gobbler
and come runnin up talking bout give him his turkey. I ast him "who
turkey you talkin bout?[Note: missing double quote?] He says dat one
of hisn I hed done grabbed. I tole him he must gone crazy in de head.
He says, I better give him his turkey before he beat my head off. I
tole him I wasn't gointer give nobody but Daisy Blunt dat turkey.
Otherwise, if he wanted to try my head, I wasn't runnin uh damn step.
Come on. So he jumped on me and tried to snatch de turkey. We fit all
over de place. First we was just tusslin for de bird, but when he
found out he couldn't take it he hit me wid his fist. Den I ups wid my
African soup bone and I bet I plowed up uh acre uh bushes wid his
head. He hit ker-bam! right in dat pack uh mule bones and I turnt and
started off, when lo and behold, he gits up wid dat hock bone and lams
me in de head and when I come to, him and my turkey was gone. So I
come swore out uh warrant aginst him cause didn't fight fair. I ain't
mad. I always lakted Jim, but he sho done dirty--lammin me wid uh mule
bone and takin' [Note: corrected missing space] my turkey.

(Dave resumes his seat and Jim drops his head for a moment, then
snatches it up arrogantly and glares at the Baptists. The whole place
is very silent for a moment. Then Mayor Clarke clears his throat, raps
with his gavel and looks sternly at Jim.)

CLARKE
Jim Weston, stand up suh! (Jim rises sullenly.) Youse charged wid
'saulting Dave Carter wid uh dangerous weapon and then stealin his
lawful turkey gobbler. You heard de charge--guilty or not guilty?

JIM
(arrogantly) Yeah, I hit him and I'll hit him agin if he crowd me. But
I ain't guilty uh no crime. (He hitches up his pants and sits down
arrogantly.)

CLARKE
(surprised) Whuts dat you say, Jim? (raps sharply) Git up from there
sir! Whuts dat you say?

JIM
(rising) I say, heah, I lammed ole Dave wid de mule bone, but I ain't
guilty uh nothin.

(There is a stark silence for a few seconds. Then Clarke raps
nervously.)

CLARKE
How come you ain't guilty?

(Jim sits down amid jubilant smiles of Methodists. Simms chuckles out
loud and wipes his face with his handkerchief. He gets to his feet
still gloating.)

SIMMS
(to Jim) Set down, Jim, and lemme show dese people dat walks in de
darkness wid sinners an' republicans de light.

SINGLETARY
You just as well tuh hush up befo' you start, then, Simms. You can't
show nobody uh light when you ain't got none tuh show.

HAMBO
Ain't dat de gospel?

NIXON
Aw, let de man talk. Y'all sound lak uh tree full uh blackbirds. Go
head on, Elder Simms.

WALTER
Yeah, you can't teach 'em nothin' but talk on. We know whut you
talkin' about.

CLARKE
(raps once or twice) I God, tell it. Whut ever tis you got tuh tell.

SISTER LEWIS
An yeah, hurry up and tell it. I know it ain't goin' tuh be nothin'
after you git it told but hurry up and say it so yo' egg-bag kin rest
easy.

WALTER
Aw shut up an' give de man uh chance.

SISTER LEWIS
My shetters ain't workin' good. Sposin' you come shet me up, Walter.
Den you'll know it's done right.

LIGE
Aw, whyn't y'all ack lak folks an' leave de man talk.

CLARKE
(rapping repeatedly) Order in dis court, I God, jus' like you was in
Orlando! (Silence falls.) Now, Simms, talk yo' chat.

SIMMS
(glances down into his open Bible then looks all around the room with
great deliberation. It is evident he enjoys being the center of
attraction. He smiles smugly as he turns his face towards the pulpit.
He speaks slowly and accents his words so that none will be lost on
his audience.) De Bible says, be sho' you're right, then go ahead. (He
looks all around to collect the admiration he feels he has earned.)
Now, we all done gethered and 'sembled here tuh law dis young lad of
uh boy on uh might serious charge. Uh whole passle of us is rarin tuh
drive him way from home lak you done done off his daddy an' his
brothers.

HAMBO
We never drove off his pappy. De white folks took an' hung him for
killin' dat man [Note: corrected missing space?] in Kissimmee for
nothin'.

SIMMS
Dat ain't de point, brother Hambo.

HAMBO
It's jes' as good uh point as any. If you gointer talk--tell de truth.
An if you can't tell de truth, set down an' leave Rev. Singletary
talk.

SIMMS
Brother Mayor, how come you let dese people run they mouf
lak uh passle uh cow-bells? Ain't I got de floor? I ain't no
breath-and-britches. I was _people_ in Middle Georgy befo' I ever come
to Floridy. Whut kind of Chairman is you, nohow?

CLARKE
(angrily) Heah! Heah! Don't you come tryin' show yo'self round me! I
God, I don't keer whut you wuz in Georgy. I God, I kin eat fried
chicken when you [Note: corrected missing spaces] caint git rain water
tuh drink. Hurry up an' say dat mess you got in yo' craw an' set down.
We needs yo' space more than we needs yo' comp'ny.

NIXON
Don't let him skeer you, Elder Sims. You got plenty shoulders tuh back
yo' fallin.

HAMBO
Well, each an' every shoulder kin hit de ground an' I'll git wid 'em.
Don't like it dontcher take, here my collar come an' shake it.

WALTER
Hambo, everybody in Orange County knows you love tuh fight. But dis is
uh law hearin'--not no wrassle.

HAMBO
Oh you Methdis' niggers wants tuh fight bad enough, but youse skeered.
Youse jus' as hot as Tucker when de mule kicked his mammy. But you
know you got plenty coolers.

SISTER TAYLOR
Aw, taint nobody skeered uh you half-pint Baptists. God knows Ahm
ready an' willin'. (She glares at Mrs. Lewis.)

(Sister Lewis jumps to her feet but is pulled back into her seat.
Mayor Clarke raps for order and the room gets quiet.)

CLARKE
Aw right now, Simms. I God, git through.

SIMMS
(pompously) Now, y'all done up an' took dis po' boy an' had him locked
up in uh barn ever since Sat'day night an' done got him 'coused uh
assault an' stealing uh turkey an' I don't know whut all an' you ain't
got no business wid yo' hands on him stell. He ain't done no crime,
an' if y'all knowed anything 'bout law, I wouldn't have tuh tell you so.

CLARKE
I God, he is done uh crime and he's gointer ketch it, too.

SIMMS
But not by law, Brother Mayor. You tryin' tuh lay uh hearin' on dis
boy an' you can't do it cause he ain't broke no law--I don't keer whut
he done so long as he don't break no law you can't tetch him.

SINGLETARY
He committed assault, didn't he? Dat sho is breakin' de law.

SIMMS
Naw, he ain't committed no 'sault. He jus' lammed Dave over de head
an' took his own turkey an' come on home, dat's all. (triumphantly)
Yuh see y'll don't knoww whut you talkin' 'bout. Now, I done set in de
court house an' heard de white folks law from mornin' till night. (He
flips his Bible shut.) I done read dis book from lid tuh lid an' I
knows de law. You got tuh have uh weepon tuh commit uh 'sault. An'
taint in no white folks law an taint in dis Bible dat no mule bone is
no weapon. I

CLARKE
(after a moment of dead silence) I God, whut's dat you say?

SIMMS
(sitting down and crossing his legs and folding his hands upon his
Bible) You heard me. I say you ain't got no case 'ginst dis boy an'
you got tuh turn him go.

SINGLETARY
(jumping up) Brother Chairman--

CLARKE
(raps once and nods recognition) You got de floor.

SINGLETARY
I ain't book-learnt an' I ain't rubbed de hair offen my head agin no
college wall, but I know when uh 'sault been committed. I says Jim
Weston did 'sault Davie. (He points at Dave's head.) An' steal his
turkey. Everybody knows Jim can't hunt wid Dave. An' he 'saulted Dave
too.

SIMMS
(arrogantly) Prove it!

(Singletary stands there silent and puzzled. The Methodist side breaks
into a triumphant shout of "Oh Mary, don't you weep, don't you moan,
Pharaoh's army got drownded." Singletary sinks into his seat. When
they have shouted out three choruses, Simms arises to speak.)

I move dat we sing doxology and bring dis meetin' to uh close. We'se
all workin' people, Brother Mayor. Dismiss us so we kin gwan back to
our work. De sun is two hours high yet. (looks towards the Methodist
side) I move dat we adjourn.

WALTER
I second de motion.

SINGLETARY
(arising slowly) Hold on there uh minute wid dat motion. Dis ain't no
lodge meetin'. Dis is uh court an' bofe sides got uh right tuh talk.
(motions towards Simms' Bible) Youse uh letter learnt man but I kin
read dat Bible some too. Lemme take it uh minute.

SIMMS
I ain't gointer do it. Any preacher dat amounts to uh hill uh beans
would have his own Bible.

CLARKE
I God, Singletary, you right here in yo' own church. Come on up here
an' read out yo' pulpit Bible. I God, don't mind me being up here.
Come on up.

(A great buzzing breaks out all over the church as Singletary mounts
the pulpit. Clarke raps for order. Simms begins to turn the leaves of
the Bible.)

SIMMS
Brother Mayor, you oughter let us outa here. You ain't got no case
'ginst dis boy. Don't waste our time for nothin'. Leave us go home.

CLARKE
Aw, dry up, Simms. You done talked yo' talk. I God, leave Singletary
talk his. (to Singletary) Step on out when you ready, Rev.

REV. SINGLETARY
(Reading) It says here in Judges 18:18 dat Samson slewed three
thousand [Note: corrected missing space] Philistines wid de jawbone of
an ass.

SIMMS
(on his feet) Yeah, but dis wasn't no ass. Dis was uh mule, Brother
Mayor. Dismiss dis meetin' and less all go home.

SINGLETARY
Yeah, but he was half-ass. A ass is uh mule's daddy and he's biggern
uh ass, too. (emphatic gestures) Everybody knows dat--even de lil
chillun.

SIMMS
(standing) Yeah, but we didn't come here to talk about no asses,
neither no half asses, nor no mule daddies. (laughter from de
Methodists) We come to law uh boy for 'sault an' larceny.

SINGLETARY
(very patiently) We'se comin' to dat pint now. Dat's de second claw uh
de sentence wese expoundin'. I say Jim Weston did have uh weepon in
his hand when he 'saulted Dave. Cause y'all knows if de daddy is
dangerous, den de son is dangerous too. An' y'all knows dat de further
back you gits on uh mule de more dangerous he gits an' if de jawbone
slewed three thousand people, by de time you gits back tuh his hocks,
its pizen enough tuh kill ten thousand. Taint no gun in de world ever
kilt dat many mens. Taint no knives nor no razors ever kilt no three
thousand people. Now, folkses, I ast y'all whut kin be mo' dangarous
dan uh mule bone? (to Clarke) Brother Mayor, Jim didn't jes' lam Dave
an walk off. (very emphatic) He 'saulted him wid de deadliest weepon
there is in de world an' while he was layin' unconscious, he stole his
turkey an' went. Brother Mayor, he's uh criminal an' oughter be run
outa dis peaceful town.

(Great chorus of approval from Baptist Clarke begins to rap for
order.)

SIMMS
(attending) Brother Mayor, I object. I have studied jury and I know
what I'm talkin' about.

CLARKE
Aw dry up, Simms. Youse entirely out of order. You may be slick, but
you kin stand another greasing. Rev. Singletary is right. I God, I
knows de law when I hear it. Stand up dere, Jim.

(Jim rises very slowly. Simms rises also.)

CLARKE
Set down, Simms. I God, I know where to find you when I want you.
(Simms sits.) Jim, I find you guilty as charged an' I wants you to git
outa my town and stay gone for two years. (to Lum) Brother Marshall,
you see dat he gits outa town befo' dark. An' you folks dats so
anxious to fight, git on off dis church grounds befo' you start. And
don't use no knives and no guns and no mule bones. Court's dismissed.


_CURTAIN_




ACT III


Scene I


SETTING: Curtain goes up on a stretch of railroad track with a
luxurious Florida forest on the backdrop. Entrances left and right. It
is near sundown.

ACTION: When the curtain goes up there is no one on the stage, but
there is a tremendous noise and hub-bub off stage right. There are
yells of derision, shouts of anger. Part of the mob is trying to keep
Jim in town and a part is driving him off. After a full minute of
this, Jim enters with his guitar hanging around his neck and his coat
over his shoulder. The sun is dropping low and red thru the forest. He
is looking back angrily and shouting back at the mob. A small missile
is thrown after him. Jim drops his coat and guitar and grabs up a
piece of brick and threatens to throw it.

JIM
(Running back the way he came and hurls the brick with all his might.)
I'll kill some of youole box-ankled niggers--(grabs up another piece
of brick) I'm out yo' ole town--now jus' some of you ole half-pint
Baptists let yo' wooden God and Cornstalk Jesus fool you to hit me!
(Threatens to throw. There are some frightened screams and the mob is
heard running back.) I'm glad I'm out yo' ole town, anyhow. I ain't
never comin' back no more, neither. You ole ugly-rump niggers done
ruint de town anyhow.

(There is complete silence off stage. Jim walks a few steps then sits
down on the railroad embankment facing the audience. Jim pulls off one
shoe and pours the sand out. He holds the shoe in his hand a moment
and looks wistfully back down the railroad track.)

JIM
Lawd, folks sho is deceitful. (He puts on the shoe and looks back down
the track again.) I never woulda thought people woulda acted lak dat.
(Laces up the shoe) Specially Dave Carter, much as me an' him done
proaged round together goin' in swimmin' and playin' ball an'
serenadin' de girls an' de white folks.

(He sits there gloomily silent for a while, then looks behind him and
picks up his guitar and begins to pick a tune. It is very sad. He
trails off into "You may leave an' go to Halimuhfack." When he
finishes he looks back at the sun and picks up his coat also.)

I never woulda thought people woulda acted lak dat. (laces up the
shoe) Specially Dave Carter, much as me an' him done proaged round
together, goin' in swimmin' and playin' ball an' serenadin' de girls
an' de white folks. (He sits there gloomily silent for a while then
looks behind him and picks up his guitar and beings to pick a tune. It
is very sad. He trails off into "You may leave and go to Halimuhfack."
When he finishes he looks back at the sun and picks up his coat also.)
(He looks back again towards the village.) Reckon I better git on down
de road an' git somewhere, Lawd knows where. (stops suddenly in his
tracks and turns back towards the village and takes a step or two.)
All dat mess and stink for nothin'. Dave knows good an' well I didn't
mean to hurt him much. (He takes off his cap and scratches his head
thoroughly, then turns again and starts on down the road towards left.
Enter Daisy left walking briskly.)

DAISY
Hello, Jim.

JIM
Hello, Daisy.

(Embarrassed silence)

DAISY
I was just coming over town to see how you come out.

JIM
You don't have to go way over there to find dat out--you and Dave done
got me run outa town for nothin'.

DAISY
(Putting her hand on his arm) Dey didn't run you outa town, did dey?

JIM
(Shaking her hand off) Whut you reckon I'm countin' Mr. Railroad's
ties for--just to find out how many ties between here and Orlando?

DAISY
(Hand on his arm again) Dey _cain't_ run you off like dat!

JIM
Take yo' hands off me, Daisy! How come they can't run me off wid you
and Dave an'--_everybody_ gainst me?

DAISY
I ain't opened my moff 'gainst you, Jim. I ain't said one word--I
wasn't even at de old trial. My madame wouldn't let me git off. I wuz
just comin' to see 'bout you now.

JIM
Aw, go 'head on. You figgered I was gone too long to talk about. You
was haulin' it over to town to see Dave--dat's whut was doin'--after
gittin _me_ all messed up.

DAISY
(Making as if to cry) I wasn't studying 'bout no Dave.

JIM
(Hopefully) Aw, don't tell me. (Sings) Ashes to ashes, dust to dust,
show me a woman that a man can trust.

(Daisy is crying now.)

JIM
Whut you crying for? You know you love Dave. I'm yo' monkey-man. He
always could do more wid you that I could.

DAISY
Naw, you ain't no monkey-man neither. I don't want you to leave town.
I didn't want y'all to be fightin' over me, nohow.

JIM
Aw, rock on down de road wid dat stuff. A two-timing cloaker like you
don't keer whut come off. Me and Dave been good friends ever since we
was born till you had to go flouncing yourself around.

DAISY
What did I do? All I did was to come over town to see you and git a
mouf-ful of gum. Next thing I now y'all is fighting and carrying on.

JIM
(stands silent for a while) Did you come over there Sat'day to see me
sho nuff, sugar babe?

DAISY
Everybody could see dat but you.

JIM
Just like I told you, Daisy. I'll say it before yo' face and behind
yo' back. I could kiss you every day--just as regular as pig-bracks.

DAISY
And I tole you I could stand it too--justa s regular as you could.

JIM
(Catching her by the arm and pulling her down with him onto the rail)
Set, down here, Daisy. Less talk some chat. You want me sho
nuff--honest to God?

DAISY
(coyly) 'Member whut I told you out on de lake last summer?

JIM
Sho nuff, Daisy?

(Daisy nods smilingly.)

JIM
(Sadly) But I got to go 'way. Whut we gointer to 'bout dat?

DAISY
Where you goin', Jim?

JIM
(Looking sadly down the track) God knows.

(Off stage from the same direction from which Jim entered comes the
sound of whistling and tramping of feet on the ties.)

JIM
(Brightening) Dat's Dave! (Frowning suspiciously) Wonder whut he doin'
walking dis track? (Looks accusingly at Daisy) I bet he's goin' to yo'
work-place.

DAISY
Whut for?

JIM
He ain't goin' to see de madame--must be goin' to see you. (He starts
to rise petulantly as Dave comes upon the scene. Daisy rises also.)

DAVE
(Looks accusingly from one to the other) Whut y'all jumpin' up
for? I....

JIM
Whut you got to do wid us business? Tain't none of yo' business if we
stand up, set down or fly like a skeeter hawk.

DAVE
Who said I keered? Dis railroad belongs to de _man_--I kin walk it
good as you, can't I?

JIM
(Laughing exultantly) Oh yeah, Mr. Do-Dirty! You figgered you had done
run me on off so you could git Daisy all by yo'self. You was headin'
right for her work-place.

DAVE
I wasn't no such a thing.

JIM
You was. Didn't I lear you coming down de track all whistling and
everything?

DAVE
Youse a big old Georgy-something-ain't-so! I done got my belly full of
Daisy Sat'day night. She can't snore in my ear no more.

DAISY
(Indignantly) Whut you come here low-rating me for, Dave Carter? I
ain't done nothin' to you but treat you white. Who come rubbed yo' ole
head for you yestiddy if it wasn't me?

DAVE
Yeah, you rubbed my head all right, and I lakted dat. But everybody
say you done toted a pan to Joe Clark's barn for Jim before I seen
you.

DAISY
Think I was going to let Jim there thout nothing fitten for a dog to
eat?

DAVE
That's all right, Daisy. If you want to pay Jim for r knockin' me in
de head, all right. But I'm a man in a class--in a class to myself and
nobody knows my name.

JIM
(Snatching Daisy around to face him) Was you over to Dave's house
yestiddy rubbing his ole head and cloaking wid him to run me outa
town--and me locked up in dat barn wid de cows and mules?

DAISY
(Sobbing) All both of y'all hollerin' at me an' fussin' me just
cause I tries to be nice--and neither one of y'all don't keer
nothin' bout me.

(Both boys glare at each other over Daisy's head and both try to hug
her at the same time. She violently wrenches herself away from both
and makes as if to move on.)

Leave me go! Take yo' rusty pams offen me. I'm going on back to my
work-place. I just got off to see bout y'all and look how y'all
treat me.

JIM
Wait a minute, Daisy. I love you like God loves Gabriel--and dat's His
best angel.

DAVE
Daisy, I love you harder than detthunder can bump a stump--if I
don't--God's a gopher.

DAISY
(Brightening) Dat's de first time you ever said so.

DAVE and JIM
Who?

JIM
Whut you hollering "who" for? Yo' foot don't fit no limb.

DAVE
Speak when you spoken to--come when you called, next fall you'll be my
coon houn' dog.

JIM
Table dat discussion. (Turning to Daisy) You ain't never give me no
chance to talk wid you right.

DAVE
You made _me_ feel like you was trying to put de Ned book on me all de
time. Do you love me sho nuff, Daisy?

DAISY
(Blooming again into coquetry) Aw, y'all better stop dat. You know you
don't mean it.

DAVE
Who don't mean it? Lemme tell you something, mama, if you was mine I
wouldn't have you counting no ties wid yo' pretty lil toes. Know whut
I'd do?

DAISY
(Coyly) Naw, whut would you do?

DAVE
I'd buy a whole passenger train and hire some mens to run it for you.

DAISY
(Happily) Oo-ooh, Dave.

JIM
(to Dave)

  De wind may blow, de door may slam
  Dat whut you shootin' ain't worth a dam.

(to Daisy) I'd buy you a great big ole ship--and then baby, I'd buy
you a ocean to[Note: corrected missing space] sail yo' ship on.

DAISY
(Happily) Oo-ooh, Jim.

DAVE
(to Jim)

  A long train, a short caboose
  Dat lie whut you shootin', ain't no use.

(to Daisy) Miss Daisy, know what I'd do for you?

DAISY
Naw, whut?

DAVE
I'd like uh job cleanin out de Atlantic Ocean jus for you.

DAISY
Don't fool me now, papa.

DAVE
I couldn't fool _you_, Daisy, cause anything I say bout lovin' you, I
don't keer how big it is, it wouldn't be half de truth. Y

DAVE
I'd come down de river riding a mud cat and leading a minnow.

DAISY
Lawd, Dave, you sho is propaganda.

JIM
(Peevishly) Naw he ain't--he's just lying--he's a noble liar. Know
whut I'd do if you was mine?

DAISY
Naw, Jim.

JIM
I'd make a panther wash yo' dishes and a 'gator chop yo' wood for you.

DAVE
Daisy, how come you [Note: corrected missing space] let Jim lie lak
dat? He's as big a liar as he is a [Note: corrected missing space]
man. But sho nuff now, laying all sides to jokes, Jim, there don't
even know how to answer you. If you don't b'lieve it, ast him
something.

DAISY
(to Jim) You like me much, Jim?

JIM
(Enthusiastically) Yeah, Daisy, I sho do.

DAVE
(Triumphant) See dat! I tole you he didn't know how to answer nobocy
like you. If he was talking to some of them ol' funny looking gals
over town he'd be answering 'em just right. But he got to learn how to
answer _you_. Now you ast _me_ something and see how I answer you.

DAISY
Do you like me, Dave?

DAVE
(Very properly in a falsetto voice) Yes ma'am! Dat's de way to answer
swell folks like you. Furthermore, less we prove which one [Note:
corrected missing space] of us love you de best right now. (To Jim)
Jim, how much time owuld you do on de chain-gang for dis 'oman?

JIM
Twenty years and like it.

DAVE
See dat, Daisy? Dat nigger ain't willing to do no time for you. I'd
_beg_ de judge to gimme life. (Both Jim and Dave laugh)

DAISY
Y'all doin' all dis bookooing out here on de railroad track but I bet
y'all crazy 'bout Bootsie and Teets and a whole heap of others.

JIM
Cross my feet and hope to die! I'd ruther see all de other wimmenfolks
in de world dead than for[Note: corrected missing space] you to have de
tooth-ache.

DAVE
If I was dead any any other woman come near my coffin de undertaker
would have to do his job all over--cause I'd git right up and walk
off. Furthermore, Miss Daisy, ma'am, also m'am, which would _you_
ruther be a lark a flying or a dove a settin'--ma'am also ma'am?

DAISY
'Course I'd ruther be a dove.

JIM
Miss Daisy, ma'am, also ma'am--if you marry dis nigger over my head,
I'm going to git me a green hickory club and season it over yo' head.

DAVE
Don't you be skeered, baby--papa kin take keer a _you_. (to Jim)
Counting from de finger (suiting the action to the word) back to the
thumb--start anything I got you some.

JIM
Aw, I don't want no more fight wid you, Dave.

DAVE
Who said anything about fighting? We just provin' who love Daisy de
best. (to Daisy) Now, which one of us you think love you de best?

DAISY
Deed I don't know, Dave.

DAVE
Baby, I'd walk de water for you--and tote a mountain on my head while
I'm walkin'.

JIM
Know whut I'd do, honey babe? If you was a thousand miles from home
and you didn't have no ready-made money and you had to walk all de
way, walkin' till yo' feet start to rolling, just like a wheel, and I
was riding way up in de sky, I'd step backwards offa dat airyplane
just to walk home wid you.

DAISY
(Falling on Jim's neck) Jim, when you talk to me like dat I just can't
stand it. Less us git married right now.

JIM
Now you talkin' like a blue-back speller. Less go!

DAVE
(Sadly) You gointer leave me lak dis, Daisy?

DAISY
(Sadly) I likes you, too, Dave, I sho do. But I can't marry both of
y'all at de same time.

JIM
Aw, come on, Daisy--sun's gettin' low. (He starts off pulling Daisy.)

DAVE
Whut's I'm gointer do? (Walking after them)

JIM
Gwan back and hunt turkeys--you make out you so touchous nobody can't
tell you yo' eye is black thout you got to run git de law.

DAVE
(Almost tearfully) Aw Jim, shucks! Where y'all going?

(Daisy comes to an abrupt halt and stops Jim)

DAISY
That's right, Honey. Where _is_ we goin' sho nuff?

JIM (Sadly)
Deed I don't know, baby. They just sentenced [Note: corrected missing
space] me to go--they didn't say where and I don't know.

DAISY
How we goin' know how to go when [Note: corrected missing space] we
don't know where we goin'?

(Jim looks at Dave as if he expects some help but Dave stands sadly
silent. Jim takes a few steps forward as if to go on. Daisy makes a
step or two, unwillingly, then looks behind her and stops. Dave looks
as if he will follow them.)

DAISY
Jim! (He stops and turns) Wait a minute! Whut we gointer do when we
git there?

JIM
Where?

DAISY
Where we goin'?

JIM
I done tole you I don't know where it is.

DAISY
But how we gointer git something to eat and a place to stay?

JIM
Play my box for de white folks and dance just like I been doing.

DAISY
You can't take keer of me on dat, not where we hafta pay rent.

JIM
(Looks appealingly at Dave, then away quickly) Well, I can't help
_dat_, can I?

DAISY
(Brightly) I tell you whut, Jim! Less us don't go nowhere. They
sentenced you to leave Eatonville and youse almost a mile from de city
limits already. Youse in Maitland now. Supposin' you come live on de
white folks' place wid me after we git married. Eatonville ain't got
nothin' to do wid you livin' in Maitland.

JIM
Dat's a good idea, Daisy.

DAISY
(Jumping into his arms) And lissen, honey, you don't have to be
beholden to nobody. You can throw dat ole box away if you want to. I
know where you can get a _swell_ job.

JIM
(Sheepishly) Doin' whut? (Looks lovingly at his guitar)

DAISY
(Almost dancing) Yard man. All you have to do is wash windows, and
sweep de sidewalk, and scrub off de steps and porch and hoe up de
weeds and rake up de leaves and dig a few holes now and then with a
spade--to plant some trees and things like that. It's a good steady
job.

JIM
(After a long deliberation) You see, Daisy, de mayor and corporation
told me to go on off and I oughter go.

DAISY
Well, I'm not going tippin' down no railroad track like a Maltese cat.
I wasn't brought up knockin' round from here to yonder.

JIM
Well, I wasn't brought up wid no spade in my hand--and ain't going to
start it now.

DAISY
But sweetheart, we got to live, ain't we? We got to git hold of money
before we kin do anything. I don't mean to stay in de white folks'
kitchen all my days.

JIM
Yeah, all dat's true, but you couldn't buy a flea a waltzing jacket
wid de money _I'm_ going to make wid a hoe and spade.

DAISY
(Getting tearful) You don't want me. You don't love me.

JIM
Yes, I do, darling, I love you. Youse de one letting a spade come
between us. (He caresses her.) I loves you and you only. You don't see
_me_ dragging a whole gang of farming tools into us business, do you?

DAISY
(stiffly) Well, I ain't going to marry no man that ain't going to work
and take care of me.

JIM
I don't mind working if de job ain't too heavy for me. I ain't going
to bother wid nothin' in my hands heavier than dis box--and I totes it
round my neck 'most of de time. I kin go out and hunt you some game
when times gits tight.

DAISY
Don't strain yo'self huntin' nothin' for me. I ain't goin' to eat
nobody's settin' hen. (She turns to DAVE finally.)

JIM
Whut ole sittin hen? Ain't you and Lum done et up de turkey
I--I--bought?

DAISY
You might of brought it, but Dave sho kilt it. You couldn't hit de
side of uh barn wid uh bass fiddle.

DAVE
Course I kilt it, and I kilt it for you, but I didn't kill none for
Lum Boger. De clean head hound!

(Daisy turns to Dave finally)

DAISY
Well, I reckon you loves me the best anyhow. You wouldn't talk to me
like Jim did, would you, Dave?

DAVE
Naw, I wouldn't say whut he said a-tall.

DAISY
(Cuddling up to him) Whut would _you_ say, honey?

DAVE
I'd say dat box was too heavy for me to fool wid. I wouldn't tote
nothing my gun and my hat and I feel like I'm 'busing myself sometie
totin' dat.

DAISY
(Outraged) Don't you mean to work none?

DAVE
Wouldn't hit a lick at a snake.

DAISY
I don't blame _you_, Dave (looks down at his feet) cause toting dem
feet of yourn is enough to break down your constitution.

DAVE
They carries me wherever I wants to go. Daisy, you marry Jim cause I
don't want to come between y'all. He's my buddy.

JIM
Come to think of it, Dave, she was yourn first. You take and handle
dat spade for her.

DAVE
You heard her say it is all I can do to lift up dese feets and put 'em
down. Where I'm going to git any time to wrassle wid any hoes and
shovels? You kin git round better'n me. You done won Daisy--I give in.
I ain't going to bite no friend of mine in de back.

DAISY
Both of you niggers can git yo' hat en' yo' heads an' git on down de
road. Neither one of y'all don't have to have me. I got a good job and
plenty men begging for yo' chance.

JIM
Dat's right, Daisy, you go git you one them mens whut don't mind
smelling mules--and beating de white folks to de barn every morning. I
don't wanta be bothered wid nothin' but dis box.

DAVE
And I can't strain wid nothin' but my feets and my gun. I kin git mo'
turkey gobblers, but never no job.

(Daisy walks slowly away in the direction from which she came. Both
watch her a little wistfully for a minute. The sun is setting.)

DAVE
Guess I better be gitin' on back--it's most dark. Where you goin, Jim?

JIM
I don't know, Dave. Down de road, I reckon.

DAVE
Whyncher come on back to town? Taint no use you proagin' up and down
[Note: corrected missing space] de railroad track when you got a home.

JIM
They done lawed me way from it for hittin' you wid dat bone.

DAVE
Dat ain't nothin'. It was my head you hit. An' if I don't keer whut
dem ole ugly-rump niggers got to do wid it?

JIM
They might not let me come in town.

DAVE
(Seizing Jim's arm and facing him back toward the town.) They better!
Look here, Jim, if they try to keep you out dat town we'll go out to
dat swamp and git us a mule bone a piece and come back and boil dat
stew down to a low gravy.

JIM
You mean dat Dave? (Dave nods his head eagerly.)

DAVE
Us wasn't mad wid one 'nother nohow. Come on less go back to town. Dem
mullet heads better leave me be, too. (Picks up a heavy stick) I wish
Lum would come tellin' me bout de law when I got all dis law in _my_
hands. An' de rest of dem 'gator-face jigs--if they ain't got a whole
set of mule bones and a good determination they better not bring de
mess up.


_CURTAIN_







End of Project Gutenberg's De Turkey and De Law, by Zora Neale Hurston

*** END OF THIS PROJECT GUTENBERG EBOOK DE TURKEY AND DE LAW ***

***** This file should be named 22146-8.txt or 22146-8.zip *****
This and all associated files of various formats will be found in:
        http://www.gutenberg.org/2/2/1/4/22146/

Produced by Charlene Taylor and the Online Distributed
Proofreading Team at http://www.pgdp.net (This file was
produced from images generously made available by the
Library of Congress)


Updated editions will replace the previous one--the old editions
will be renamed.

Creating the works from public domain print editions means that no
one owns a United States copyright in these works, so the Foundation
(and you!) can copy and distribute it in the United States without
permission and without paying copyright royalties.  Special rules,
set forth in the General Terms of Use part of this license, apply to
copying and distributing Project Gutenberg-tm electronic works to
protect the PROJECT GUTENBERG-tm concept and trademark.  Project
Gutenberg is a registered trademark, and may not be used if you
charge for the eBooks, unless you receive specific permission.  If you
do not charge anything for copies of this eBook, complying with the
rules is very easy.  You may use this eBook for nearly any purpose
such as creation of derivative works, reports, performances and
research.  They may be modified and printed and given away--you may do
practically ANYTHING with public domain eBooks.  Redistribution is
subject to the trademark license, especially commercial
redistribution.



*** START: FULL LICENSE ***

THE FULL PROJECT GUTENBERG LICENSE
PLEASE READ THIS BEFORE YOU DISTRIBUTE OR USE THIS WORK

To protect the Project Gutenberg-tm mission of promoting the free
distribution of electronic works, by using or distributing this work
(or any other work associated in any way with the phrase "Project
Gutenberg"), you agree to comply with all the terms of the Full Project
Gutenberg-tm License (available with this file or online at
http://gutenberg.org/license).


Section 1.  General Terms of Use and Redistributing Project Gutenberg-tm
electronic works

1.A.  By reading or using any part of this Project Gutenberg-tm
electronic work, you indicate that you have read, understand, agree to
and accept all the terms of this license and intellectual property
(trademark/copyright) agreement.  If you do not agree to abide by all
the terms of this agreement, you must cease using and return or destroy
all copies of Project Gutenberg-tm electronic works in your possession.
If you paid a fee for obtaining a copy of or access to a Project
Gutenberg-tm electronic work and you do not agree to be bound by the
terms of this agreement, you may obtain a refund from the person or
entity to whom you paid the fee as set forth in paragraph 1.E.8.

1.B.  "Project Gutenberg" is a registered trademark.  It may only be
used on or associated in any way with an electronic work by people who
agree to be bound by the terms of this agreement.  There are a few
things that you can do with most Project Gutenberg-tm electronic works
even without complying with the full terms of this agreement.  See
paragraph 1.C below.  There are a lot of things you can do with Project
Gutenberg-tm electronic works if you follow the terms of this agreement
and help preserve free future access to Project Gutenberg-tm electronic
works.  See paragraph 1.E below.

1.C.  The Project Gutenberg Literary Archive Foundation ("the Foundation"
or PGLAF), owns a compilation copyright in the collection of Project
Gutenberg-tm electronic works.  Nearly all the individual works in the
collection are in the public domain in the United States.  If an
individual work is in the public domain in the United States and you are
located in the United States, we do not claim a right to prevent you from
copying, distributing, performing, displaying or creating derivative
works based on the work as long as all references to Project Gutenberg
are removed.  Of course, we hope that you will support the Project
Gutenberg-tm mission of promoting free access to electronic works by
freely sharing Project Gutenberg-tm works in compliance with the terms of
this agreement for keeping the Project Gutenberg-tm name associated with
the work.  You can easily comply with the terms of this agreement by
keeping this work in the same format with its attached full Project
Gutenberg-tm License when you share it without charge with others.

1.D.  The copyright laws of the place where you are located also govern
what you can do with this work.  Copyright laws in most countries are in
a constant state of change.  If you are outside the United States, check
the laws of your country in addition to the terms of this agreement
before downloading, copying, displaying, performing, distributing or
creating derivative works based on this work or any other Project
Gutenberg-tm work.  The Foundation makes no representations concerning
the copyright status of any work in any country outside the United
States.

1.E.  Unless you have removed all references to Project Gutenberg:

1.E.1.  The following sentence, with active links to, or other immediate
access to, the full Project Gutenberg-tm License must appear prominently
whenever any copy of a Project Gutenberg-tm work (any work on which the
phrase "Project Gutenberg" appears, or with which the phrase "Project
Gutenberg" is associated) is accessed, displayed, performed, viewed,
copied or distributed:

This eBook is for the use of anyone anywhere at no cost and with
almost no restrictions whatsoever.  You may copy it, give it away or
re-use it under the terms of the Project Gutenberg License included
with this eBook or online at www.gutenberg.org

1.E.2.  If an individual Project Gutenberg-tm electronic work is derived
from the public domain (does not contain a notice indicating that it is
posted with permission of the copyright holder), the work can be copied
and distributed to anyone in the United States without paying any fees
or charges.  If you are redistributing or providing access to a work
with the phrase "Project Gutenberg" associated with or appearing on the
work, you must comply either with the requirements of paragraphs 1.E.1
through 1.E.7 or obtain permission for the use of the work and the
Project Gutenberg-tm trademark as set forth in paragraphs 1.E.8 or
1.E.9.

1.E.3.  If an individual Project Gutenberg-tm electronic work is posted
with the permission of the copyright holder, your use and distribution
must comply with both paragraphs 1.E.1 through 1.E.7 and any additional
terms imposed by the copyright holder.  Additional terms will be linked
to the Project Gutenberg-tm License for all works posted with the
permission of the copyright holder found at the beginning of this work.

1.E.4.  Do not unlink or detach or remove the full Project Gutenberg-tm
License terms from this work, or any files containing a part of this
work or any other work associated with Project Gutenberg-tm.

1.E.5.  Do not copy, display, perform, distribute or redistribute this
electronic work, or any part of this electronic work, without
prominently displaying the sentence set forth in paragraph 1.E.1 with
active links or immediate access to the full terms of the Project
Gutenberg-tm License.

1.E.6.  You may convert to and distribute this work in any binary,
compressed, marked up, nonproprietary or proprietary form, including any
word processing or hypertext form.  However, if you provide access to or
distribute copies of a Project Gutenberg-tm work in a format other than
"Plain Vanilla ASCII" or other format used in the official version
posted on the official Project Gutenberg-tm web site (www.gutenberg.org),
you must, at no additional cost, fee or expense to the user, provide a
copy, a means of exporting a copy, or a means of obtaining a copy upon
request, of the work in its original "Plain Vanilla ASCII" or other
form.  Any alternate format must include the full Project Gutenberg-tm
License as specified in paragraph 1.E.1.

1.E.7.  Do not charge a fee for access to, viewing, displaying,
performing, copying or distributing any Project Gutenberg-tm works
unless you comply with paragraph 1.E.8 or 1.E.9.

1.E.8.  You may charge a reasonable fee for copies of or providing
access to or distributing Project Gutenberg-tm electronic works provided
that

- You pay a royalty fee of 20% of the gross profits you derive from
     the use of Project Gutenberg-tm works calculated using the method
     you already use to calculate your applicable taxes.  The fee is
     owed to the owner of the Project Gutenberg-tm trademark, but he
     has agreed to donate royalties under this paragraph to the
     Project Gutenberg Literary Archive Foundation.  Royalty payments
     must be paid within 60 days following each date on which you
     prepare (or are legally required to prepare) your periodic tax
     returns.  Royalty payments should be clearly marked as such and
     sent to the Project Gutenberg Literary Archive Foundation at the
     address specified in Section 4, "Information about donations to
     the Project Gutenberg Literary Archive Foundation."

- You provide a full refund of any money paid by a user who notifies
     you in writing (or by e-mail) within 30 days of receipt that s/he
     does not agree to the terms of the full Project Gutenberg-tm
     License.  You must require such a user to return or
     destroy all copies of the works possessed in a physical medium
     and discontinue all use of and all access to other copies of
     Project Gutenberg-tm works.

- You provide, in accordance with paragraph 1.F.3, a full refund of any
     money paid for a work or a replacement copy, if a defect in the
     electronic work is discovered and reported to you within 90 days
     of receipt of the work.

- You comply with all other terms of this agreement for free
     distribution of Project Gutenberg-tm works.

1.E.9.  If you wish to charge a fee or distribute a Project Gutenberg-tm
electronic work or group of works on different terms than are set
forth in this agreement, you must obtain permission in writing from
both the Project Gutenberg Literary Archive Foundation and Michael
Hart, the owner of the Project Gutenberg-tm trademark.  Contact the
Foundation as set forth in Section 3 below.

1.F.

1.F.1.  Project Gutenberg volunteers and employees expend considerable
effort to identify, do copyright research on, transcribe and proofread
public domain works in creating the Project Gutenberg-tm
collection.  Despite these efforts, Project Gutenberg-tm electronic
works, and the medium on which they may be stored, may contain
"Defects," such as, but not limited to, incomplete, inaccurate or
corrupt data, transcription errors, a copyright or other intellectual
property infringement, a defective or damaged disk or other medium, a
computer virus, or computer codes that damage or cannot be read by
your equipment.

1.F.2.  LIMITED WARRANTY, DISCLAIMER OF DAMAGES - Except for the "Right
of Replacement or Refund" described in paragraph 1.F.3, the Project
Gutenberg Literary Archive Foundation, the owner of the Project
Gutenberg-tm trademark, and any other party distributing a Project
Gutenberg-tm electronic work under this agreement, disclaim all
liability to you for damages, costs and expenses, including legal
fees.  YOU AGREE THAT YOU HAVE NO REMEDIES FOR NEGLIGENCE, STRICT
LIABILITY, BREACH OF WARRANTY OR BREACH OF CONTRACT EXCEPT THOSE
PROVIDED IN PARAGRAPH F3.  YOU AGREE THAT THE FOUNDATION, THE
TRADEMARK OWNER, AND ANY DISTRIBUTOR UNDER THIS AGREEMENT WILL NOT BE
LIABLE TO YOU FOR ACTUAL, DIRECT, INDIRECT, CONSEQUENTIAL, PUNITIVE OR
INCIDENTAL DAMAGES EVEN IF YOU GIVE NOTICE OF THE POSSIBILITY OF SUCH
DAMAGE.

1.F.3.  LIMITED RIGHT OF REPLACEMENT OR REFUND - If you discover a
defect in this electronic work within 90 days of receiving it, you can
receive a refund of the money (if any) you paid for it by sending a
written explanation to the person you received the work from.  If you
received the work on a physical medium, you must return the medium with
your written explanation.  The person or entity that provided you with
the defective work may elect to provide a replacement copy in lieu of a
refund.  If you received the work electronically, the person or entity
providing it to you may choose to give you a second opportunity to
receive the work electronically in lieu of a refund.  If the second copy
is also defective, you may demand a refund in writing without further
opportunities to fix the problem.

1.F.4.  Except for the limited right of replacement or refund set forth
in paragraph 1.F.3, this work is provided to you 'AS-IS' WITH NO OTHER
WARRANTIES OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
WARRANTIES OF MERCHANTIBILITY OR FITNESS FOR ANY PURPOSE.

1.F.5.  Some states do not allow disclaimers of certain implied
warranties or the exclusion or limitation of certain types of damages.
If any disclaimer or limitation set forth in this agreement violates the
law of the state applicable to this agreement, the agreement shall be
interpreted to make the maximum disclaimer or limitation permitted by
the applicable state law.  The invalidity or unenforceability of any
provision of this agreement shall not void the remaining provisions.

1.F.6.  INDEMNITY - You agree to indemnify and hold the Foundation, the
trademark owner, any agent or employee of the Foundation, anyone
providing copies of Project Gutenberg-tm electronic works in accordance
with this agreement, and any volunteers associated with the production,
promotion and distribution of Project Gutenberg-tm electronic works,
harmless from all liability, costs and expenses, including legal fees,
that arise directly or indirectly from any of the following which you do
or cause to occur: (a) distribution of this or any Project Gutenberg-tm
work, (b) alteration, modification, or additions or deletions to any
Project Gutenberg-tm work, and (c) any Defect you cause.


Section  2.  Information about the Mission of Project Gutenberg-tm

Project Gutenberg-tm is synonymous with the free distribution of
electronic works in formats readable by the widest variety of computers
including obsolete, old, middle-aged and new computers.  It exists
because of the efforts of hundreds of volunteers and donations from
people in all walks of life.

Volunteers and financial support to provide volunteers with the
assistance they need, is critical to reaching Project Gutenberg-tm's
goals and ensuring that the Project Gutenberg-tm collection will
remain freely available for generations to come.  In 2001, the Project
Gutenberg Literary Archive Foundation was created to provide a secure
and permanent future for Project Gutenberg-tm and future generations.
To learn more about the Project Gutenberg Literary Archive Foundation
and how your efforts and donations can help, see Sections 3 and 4
and the Foundation web page at http://www.pglaf.org.


Section 3.  Information about the Project Gutenberg Literary Archive
Foundation

The Project Gutenberg Literary Archive Foundation is a non profit
501(c)(3) educational corporation organized under the laws of the
state of Mississippi and granted tax exempt status by the Internal
Revenue Service.  The Foundation's EIN or federal tax identification
number is 64-6221541.  Its 501(c)(3) letter is posted at
http://pglaf.org/fundraising.  Contributions to the Project Gutenberg
Literary Archive Foundation are tax deductible to the full extent
permitted by U.S. federal laws and your state's laws.

The Foundation's principal office is located at 4557 Melan Dr. S.
Fairbanks, AK, 99712., but its volunteers and employees are scattered
throughout numerous locations.  Its business office is located at
809 North 1500 West, Salt Lake City, UT 84116, (801) 596-1887, email
business@pglaf.org.  Email contact links and up to date contact
information can be found at the Foundation's web site and official
page at http://pglaf.org

For additional contact information:
     Dr. Gregory B. Newby
     Chief Executive and Director
     gbnewby@pglaf.org


Section 4.  Information about Donations to the Project Gutenberg
Literary Archive Foundation

Project Gutenberg-tm depends upon and cannot survive without wide
spread public support and donations to carry out its mission of
increasing the number of public domain and licensed works that can be
freely distributed in machine readable form accessible by the widest
array of equipment including outdated equipment.  Many small donations
($1 to $5,000) are particularly important to maintaining tax exempt
status with the IRS.

The Foundation is committed to complying with the laws regulating
charities and charitable donations in all 50 states of the United
States.  Compliance requirements are not uniform and it takes a
considerable effort, much paperwork and many fees to meet and keep up
with these requirements.  We do not solicit donations in locations
where we have not received written confirmation of compliance.  To
SEND DONATIONS or determine the status of compliance for any
particular state visit http://pglaf.org

While we cannot and do not solicit contributions from states where we
have not met the solicitation requirements, we know of no prohibition
against accepting unsolicited donations from donors in such states who
approach us with offers to donate.

International donations are gratefully accepted, but we cannot make
any statements concerning tax treatment of donations received from
outside the United States.  U.S. laws alone swamp our small staff.

Please check the Project Gutenberg Web pages for current donation
methods and addresses.  Donations are accepted in a number of other
ways including checks, online payments and credit card donations.
To donate, please visit: http://pglaf.org/donate


Section 5.  General Information About Project Gutenberg-tm electronic
works.

Professor Michael S. Hart is the originator of the Project Gutenberg-tm
concept of a library of electronic works that could be freely shared
with anyone.  For thirty years, he produced and distributed Project
Gutenberg-tm eBooks with only a loose network of volunteer support.


Project Gutenberg-tm eBooks are often created from several printed
editions, all of which are confirmed as Public Domain in the U.S.
unless a copyright notice is included.  Thus, we do not necessarily
keep eBooks in compliance with any particular paper edition.


Most people start at our Web site which has the main PG search facility:

     http://www.gutenberg.org

This Web site includes information about Project Gutenberg-tm,
including how to make donations to the Project Gutenberg Literary
Archive Foundation, how to help produce our new eBooks, and how to
subscribe to our email newsletter to hear about new eBooks.
`

	b["hurston-forty-yards.txt"] = `
"FORTY YARDS"

by

ZORA *[Handwritten: (Neale)] HURSTON



"FORTY YARDS"


                              (A Negro football game with the
                              popular concept of Negro life)



TIME:      Present

PLACE:     Washington, D.C.

SCENE:     The Ball Park

PERSONS:   The Howard and Lincoln teams, the Howard band, cheer
           leaders, spectators.

SETTING:   The park with grandstands on either sides and up-stage.

ACTION:    At rise, the grandstands are full, the cheer leaders
           are violently gyrating to whip up the mob. The
           Lincoln colors fly from the right. The Howard from the
           left. Both have cheer leaders. First is heard the
           Lincoln mob singing "DIDN'T HE RAMBLE, RAMBLE."


Lincoln Mob

    And didn't he ramble, ramble, ramble all around, in and out of town
    He rambled, he rambled, rambled till Ol' Lincoln cut him down


Howard Mob

    There'll be nothing but sweetmeats, for our football team
    There'll be nothing but sweetmeats for our football team
    Baked Hampton, boiled Shaw, fried Union, Lincoln Slaw,
    There'll be nothing but sweetmeats, for our football team.

                              (Enter the HOWARD BAND, led by a
                              hot-strutting drum major. They parade
                              the field and the men students pile
                              down and fall in behind the team.
                              They sing and shout to the TEAM
                              SONG:)

    This is the t-e-a-m team
    On which the hopes of Howard lean
    Beat Ol' Hampton, beat Ol' Union
    Sweep Ol' Lincoln clean

    We are the b-e-s-t best
    Of the r-e-s-t rest
    Come and watch us put Ol' Howard
    On top of Lincoln's chest.

    We'll hit the l-i-n-e line
    For a hundred ninety-nine
    For we love Ol' Howard, yes we love her
    All the t-i-m-e time.

                              (At the conclusion the teams takes
                              the field. The ball is put into play
                              and LINCOLN kicks off to Howard. As
                              the ball is caught and when the
                              player who is carrying the ball
                              plunges, followed by his team, the
                              Lincoln players fall on their knees
                              and begin to sing I COULDN'T HEAR
                              NOBODY PRAY. The HOWARD team charges
                              down shouting Joshua fit de battle of
                              Jericho. Whenever a player is tackled
                              there is a duet of dancing. Every
                              step is a dance. Finally the
                              grandstand catches fire and the
                              dancing and the shouting runs riot up
                              there. When the ball is on Lincoln's
                              ten-yard line, they hold Howard there
                              by rounding up both teams into a
                              huddle and the bunch-shout and sing
                              to a QUICK CURTAIN.)

CURTAIN

LINCOLN'S PRAYER:

    Ah, ah, they shall not ah pass us
    Lord, Lord, Lord, Lord
    They shall not pass us, Ah-h-h-h.



*** START: FULL LICENSE ***

THE FULL PROJECT GUTENBERG LICENSE
PLEASE READ THIS BEFORE YOU DISTRIBUTE OR USE THIS WORK

To protect the Project Gutenberg-tm mission of promoting the free
distribution of electronic works, by using or distributing this work
(or any other work associated in any way with the phrase "Project
Gutenberg"), you agree to comply with all the terms of the Full Project
Gutenberg-tm License (available with this file or online at
http://gutenberg.net/license).


Section 1.  General Terms of Use and Redistributing Project Gutenberg-tm
electronic works

1.A.  By reading or using any part of this Project Gutenberg-tm
electronic work, you indicate that you have read, understand, agree to
and accept all the terms of this license and intellectual property
(trademark/copyright) agreement.  If you do not agree to abide by all
the terms of this agreement, you must cease using and return or destroy
all copies of Project Gutenberg-tm electronic works in your possession.
If you paid a fee for obtaining a copy of or access to a Project
Gutenberg-tm electronic work and you do not agree to be bound by the
terms of this agreement, you may obtain a refund from the person or
entity to whom you paid the fee as set forth in paragraph 1.E.8.

1.B.  "Project Gutenberg" is a registered trademark.  It may only be
used on or associated in any way with an electronic work by people who
agree to be bound by the terms of this agreement.  There are a few
things that you can do with most Project Gutenberg-tm electronic works
even without complying with the full terms of this agreement.  See
paragraph 1.C below.  There are a lot of things you can do with Project
Gutenberg-tm electronic works if you follow the terms of this agreement
and help preserve free future access to Project Gutenberg-tm electronic
works.  See paragraph 1.E below.

1.C.  The Project Gutenberg Literary Archive Foundation ("the Foundation"
or PGLAF), owns a compilation copyright in the collection of Project
Gutenberg-tm electronic works.  Nearly all the individual works in the
collection are in the public domain in the United States.  If an
individual work is in the public domain in the United States and you are
located in the United States, we do not claim a right to prevent you from
copying, distributing, performing, displaying or creating derivative
works based on the work as long as all references to Project Gutenberg
are removed.  Of course, we hope that you will support the Project
Gutenberg-tm mission of promoting free access to electronic works by
freely sharing Project Gutenberg-tm works in compliance with the terms of
this agreement for keeping the Project Gutenberg-tm name associated with
the work.  You can easily comply with the terms of this agreement by
keeping this work in the same format with its attached full Project
Gutenberg-tm License when you share it without charge with others.

1.D.  The copyright laws of the place where you are located also govern
what you can do with this work.  Copyright laws in most countries are in
a constant state of change.  If you are outside the United States, check
the laws of your country in addition to the terms of this agreement
before downloading, copying, displaying, performing, distributing or
creating derivative works based on this work or any other Project
Gutenberg-tm work.  The Foundation makes no representations concerning
the copyright status of any work in any country outside the United
States.

1.E.  Unless you have removed all references to Project Gutenberg:

1.E.1.  The following sentence, with active links to, or other immediate
access to, the full Project Gutenberg-tm License must appear prominently
whenever any copy of a Project Gutenberg-tm work (any work on which the
phrase "Project Gutenberg" appears, or with which the phrase "Project
Gutenberg" is associated) is accessed, displayed, performed, viewed,
copied or distributed:

This eBook is for the use of anyone anywhere at no cost and with
almost no restrictions whatsoever.  You may copy it, give it away or
re-use it under the terms of the Project Gutenberg License included
with this eBook or online at www.gutenberg.net

1.E.2.  If an individual Project Gutenberg-tm electronic work is derived
from the public domain (does not contain a notice indicating that it is
posted with permission of the copyright holder), the work can be copied
and distributed to anyone in the United States without paying any fees
or charges.  If you are redistributing or providing access to a work
with the phrase "Project Gutenberg" associated with or appearing on the
work, you must comply either with the requirements of paragraphs 1.E.1
through 1.E.7 or obtain permission for the use of the work and the
Project Gutenberg-tm trademark as set forth in paragraphs 1.E.8 or
1.E.9.

1.E.3.  If an individual Project Gutenberg-tm electronic work is posted
with the permission of the copyright holder, your use and distribution
must comply with both paragraphs 1.E.1 through 1.E.7 and any additional
terms imposed by the copyright holder.  Additional terms will be linked
to the Project Gutenberg-tm License for all works posted with the
permission of the copyright holder found at the beginning of this work.

1.E.4.  Do not unlink or detach or remove the full Project Gutenberg-tm
License terms from this work, or any files containing a part of this
work or any other work associated with Project Gutenberg-tm.

1.E.5.  Do not copy, display, perform, distribute or redistribute this
electronic work, or any part of this electronic work, without
prominently displaying the sentence set forth in paragraph 1.E.1 with
active links or immediate access to the full terms of the Project
Gutenberg-tm License.

1.E.6.  You may convert to and distribute this work in any binary,
compressed, marked up, nonproprietary or proprietary form, including any
word processing or hypertext form.  However, if you provide access to or
distribute copies of a Project Gutenberg-tm work in a format other than
"Plain Vanilla ASCII" or other format used in the official version
posted on the official Project Gutenberg-tm web site (www.gutenberg.net),
you must, at no additional cost, fee or expense to the user, provide a
copy, a means of exporting a copy, or a means of obtaining a copy upon
request, of the work in its original "Plain Vanilla ASCII" or other
form.  Any alternate format must include the full Project Gutenberg-tm
License as specified in paragraph 1.E.1.

1.E.7.  Do not charge a fee for access to, viewing, displaying,
performing, copying or distributing any Project Gutenberg-tm works
unless you comply with paragraph 1.E.8 or 1.E.9.

1.E.8.  You may charge a reasonable fee for copies of or providing
access to or distributing Project Gutenberg-tm electronic works provided
that

- You pay a royalty fee of 20% of the gross profits you derive from
     the use of Project Gutenberg-tm works calculated using the method
     you already use to calculate your applicable taxes.  The fee is
     owed to the owner of the Project Gutenberg-tm trademark, but he
     has agreed to donate royalties under this paragraph to the
     Project Gutenberg Literary Archive Foundation.  Royalty payments
     must be paid within 60 days following each date on which you
     prepare (or are legally required to prepare) your periodic tax
     returns.  Royalty payments should be clearly marked as such and
     sent to the Project Gutenberg Literary Archive Foundation at the
     address specified in Section 4, "Information about donations to
     the Project Gutenberg Literary Archive Foundation."

- You provide a full refund of any money paid by a user who notifies
     you in writing (or by e-mail) within 30 days of receipt that s/he
     does not agree to the terms of the full Project Gutenberg-tm
     License.  You must require such a user to return or
     destroy all copies of the works possessed in a physical medium
     and discontinue all use of and all access to other copies of
     Project Gutenberg-tm works.

- You provide, in accordance with paragraph 1.F.3, a full refund of any
     money paid for a work or a replacement copy, if a defect in the
     electronic work is discovered and reported to you within 90 days
     of receipt of the work.

- You comply with all other terms of this agreement for free
     distribution of Project Gutenberg-tm works.

1.E.9.  If you wish to charge a fee or distribute a Project Gutenberg-tm
electronic work or group of works on different terms than are set
forth in this agreement, you must obtain permission in writing from
both the Project Gutenberg Literary Archive Foundation and Michael
Hart, the owner of the Project Gutenberg-tm trademark.  Contact the
Foundation as set forth in Section 3 below.

1.F.

1.F.1.  Project Gutenberg volunteers and employees expend considerable
effort to identify, do copyright research on, transcribe and proofread
public domain works in creating the Project Gutenberg-tm
collection.  Despite these efforts, Project Gutenberg-tm electronic
works, and the medium on which they may be stored, may contain
"Defects," such as, but not limited to, incomplete, inaccurate or
corrupt data, transcription errors, a copyright or other intellectual
property infringement, a defective or damaged disk or other medium, a
computer virus, or computer codes that damage or cannot be read by
your equipment.

1.F.2.  LIMITED WARRANTY, DISCLAIMER OF DAMAGES - Except for the "Right
of Replacement or Refund" described in paragraph 1.F.3, the Project
Gutenberg Literary Archive Foundation, the owner of the Project
Gutenberg-tm trademark, and any other party distributing a Project
Gutenberg-tm electronic work under this agreement, disclaim all
liability to you for damages, costs and expenses, including legal
fees.  YOU AGREE THAT YOU HAVE NO REMEDIES FOR NEGLIGENCE, STRICT
LIABILITY, BREACH OF WARRANTY OR BREACH OF CONTRACT EXCEPT THOSE
PROVIDED IN PARAGRAPH F3.  YOU AGREE THAT THE FOUNDATION, THE
TRADEMARK OWNER, AND ANY DISTRIBUTOR UNDER THIS AGREEMENT WILL NOT BE
LIABLE TO YOU FOR ACTUAL, DIRECT, INDIRECT, CONSEQUENTIAL, PUNITIVE OR
INCIDENTAL DAMAGES EVEN IF YOU GIVE NOTICE OF THE POSSIBILITY OF SUCH
DAMAGE.

1.F.3.  LIMITED RIGHT OF REPLACEMENT OR REFUND - If you discover a
defect in this electronic work within 90 days of receiving it, you can
receive a refund of the money (if any) you paid for it by sending a
written explanation to the person you received the work from.  If you
received the work on a physical medium, you must return the medium with
your written explanation.  The person or entity that provided you with
the defective work may elect to provide a replacement copy in lieu of a
refund.  If you received the work electronically, the person or entity
providing it to you may choose to give you a second opportunity to
receive the work electronically in lieu of a refund.  If the second copy
is also defective, you may demand a refund in writing without further
opportunities to fix the problem.

1.F.4.  Except for the limited right of replacement or refund set forth
in paragraph 1.F.3, this work is provided to you 'AS-IS', WITH NO OTHER
WARRANTIES OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
WARRANTIES OF MERCHANTIBILITY OR FITNESS FOR ANY PURPOSE.

1.F.5.  Some states do not allow disclaimers of certain implied
warranties or the exclusion or limitation of certain types of damages.
If any disclaimer or limitation set forth in this agreement violates the
law of the state applicable to this agreement, the agreement shall be
interpreted to make the maximum disclaimer or limitation permitted by
the applicable state law.  The invalidity or unenforceability of any
provision of this agreement shall not void the remaining provisions.

1.F.6.  INDEMNITY - You agree to indemnify and hold the Foundation, the
trademark owner, any agent or employee of the Foundation, anyone
providing copies of Project Gutenberg-tm electronic works in accordance
with this agreement, and any volunteers associated with the production,
promotion and distribution of Project Gutenberg-tm electronic works,
harmless from all liability, costs and expenses, including legal fees,
that arise directly or indirectly from any of the following which you do
or cause to occur: (a) distribution of this or any Project Gutenberg-tm
work, (b) alteration, modification, or additions or deletions to any
Project Gutenberg-tm work, and (c) any Defect you cause.


Section  2.  Information about the Mission of Project Gutenberg-tm

Project Gutenberg-tm is synonymous with the free distribution of
electronic works in formats readable by the widest variety of computers
including obsolete, old, middle-aged and new computers.  It exists
because of the efforts of hundreds of volunteers and donations from
people in all walks of life.

Volunteers and financial support to provide volunteers with the
assistance they need, is critical to reaching Project Gutenberg-tm's
goals and ensuring that the Project Gutenberg-tm collection will
remain freely available for generations to come.  In 2001, the Project
Gutenberg Literary Archive Foundation was created to provide a secure
and permanent future for Project Gutenberg-tm and future generations.
To learn more about the Project Gutenberg Literary Archive Foundation
and how your efforts and donations can help, see Sections 3 and 4
and the Foundation web page at http://www.gutenberg.net/fundraising/pglaf.


Section 3.  Information about the Project Gutenberg Literary Archive
Foundation

The Project Gutenberg Literary Archive Foundation is a non profit
501(c)(3) educational corporation organized under the laws of the
state of Mississippi and granted tax exempt status by the Internal
Revenue Service.  The Foundation's EIN or federal tax identification
number is 64-6221541.  Contributions to the Project Gutenberg
Literary Archive Foundation are tax deductible to the full extent
permitted by U.S. federal laws and your state's laws.

The Foundation's principal office is located at 4557 Melan Dr. S.
Fairbanks, AK, 99712., but its volunteers and employees are scattered
throughout numerous locations.  Its business office is located at
809 North 1500 West, Salt Lake City, UT 84116, (801) 596-1887, email
business@pglaf.org.  Email contact links and up to date contact
information can be found at the Foundation's web site and official
page at http://www.gutenberg.net/about/contact

For additional contact information:
     Dr. Gregory B. Newby
     Chief Executive and Director
     gbnewby@pglaf.org

Section 4.  Information about Donations to the Project Gutenberg
Literary Archive Foundation

Project Gutenberg-tm depends upon and cannot survive without wide
spread public support and donations to carry out its mission of
increasing the number of public domain and licensed works that can be
freely distributed in machine readable form accessible by the widest
array of equipment including outdated equipment.  Many small donations
($1 to $5,000) are particularly important to maintaining tax exempt
status with the IRS.

The Foundation is committed to complying with the laws regulating
charities and charitable donations in all 50 states of the United
States.  Compliance requirements are not uniform and it takes a
considerable effort, much paperwork and many fees to meet and keep up
with these requirements.  We do not solicit donations in locations
where we have not received written confirmation of compliance.  To
SEND DONATIONS or determine the status of compliance for any
particular state visit http://www.gutenberg.net/fundraising/donate

While we cannot and do not solicit contributions from states where we
have not met the solicitation requirements, we know of no prohibition
against accepting unsolicited donations from donors in such states who
approach us with offers to donate.

International donations are gratefully accepted, but we cannot make
any statements concerning tax treatment of donations received from
outside the United States.  U.S. laws alone swamp our small staff.

Please check the Project Gutenberg Web pages for current donation
methods and addresses.  Donations are accepted in a number of other
ways including including checks, online payments and credit card
donations.  To donate, please visit:
http://www.gutenberg.net/fundraising/donate


Section 5.  General Information About Project Gutenberg-tm electronic
works.

Professor Michael S. Hart is the originator of the Project Gutenberg-tm
concept of a library of electronic works that could be freely shared
with anyone.  For thirty years, he produced and distributed Project
Gutenberg-tm eBooks with only a loose network of volunteer support.

Project Gutenberg-tm eBooks are often created from several printed
editions, all of which are confirmed as Public Domain in the U.S.
unless a copyright notice is included.  Thus, we do not necessarily
keep eBooks in compliance with any particular paper edition.

Most people start at our Web site which has the main PG search facility:

     http://www.gutenberg.net

This Web site includes information about Project Gutenberg-tm,
including how to make donations to the Project Gutenberg Literary
Archive Foundation, how to help produce our new eBooks, and how to
subscribe to our email newsletter to hear about new eBooks.
`

	b["hurston-lawing-and-jawing.txt"] = `
LAWING AND JAWING


by Zora *[Handwritten: (Neale)] Hurston



TIME:          Present

PLACE:         Way cross Georgia

SCENE:         Judge Dunfumy's Court.

PERSONS:       Judge Dunfumy, Officer Simpson and another, Jemima
               Flapcakes, Cliff Mullins, John Barnes, two lawyers,
               a clerk, a pretty girl and her escort.

SETTING:       Usual court-room arrangement, except that there is a
               large red arrow pointing off-stage left, marked
               "To Jail."

ACTION:        At rise everybody is in place except the Judge.
               Suddenly the CLERK looks off-stage right and motions
               for everybody to rise. Enter the JUDGE. He wears a
               black cap and gown and has his gavel in his hand.
               The two POLICEMEN walk behind him holding up his gown.
               He mounts the bench and glares all about him before he
               seats himself. There is a PRETTY GIRL in the front row
               left, and he takes a good look at her, smiles, frowns
               at her escort. He motions the police to leave him and
               take their places with the spectators and he then raps
               vigorously with his gavel for order.



JUDGE
Hear! Hear! Court is set! My honor is on de bench. You moufy folks set
up!
                             (He glares at the boy with the pretty girl)
All right, Mr. Whistle-britches, just keep up dat jawing now and see
how much time I'll give you!


BOY
I wasn't talking, your honor.


JUDGE
Well, quit looking so moufy.
                             (to CLERK)
Call de first case. And I warn each and all dat my honor is in bad
humor dis mawnin'. I'd give a canary bird twenty years for peckin'
at a elephant.
                             (to CLERK)
Bring 'em on.


CLERK
                             (Reading)
Cliff Mullins, charged with assault upon his wife with a weapon and
disturbing the peace.
                             (As CLIFF is led to the bar by the
                             officer, the JUDGE glares ferociously
                             at the prisoner. His wife, all
                             bandages, limps up to the bar at the
                             same time.)


JUDGE
So youse one of dese hard-boiled wife-beaters, huh? Just a mean old
woman-Jessie! If I don't lay a hearing on you, God's a gopher! Now what
_made_ you cut such a caper?


CLIFF
Judge, I didn't go hunt her. Saturday night I was down on Dearborn
Street in a nasty ditch *[Handwritten: nasty ditch crossed out in
pencil, (buffet flat)]--


JUDGE
A nasty ditch? *[Handwritten: A nasty ditch crossed out in pencil,
(Buffet flat)]


CLIFF
Aw, at Emma Hayles' house.


JUDGE
Oh, yes. Go on.


CLIFF
Well,
                             (Points thumb at wife)
she come down dere and claim I took her money and she claimed I wuz
spending it on Emma.


CLIFF'S WIFE
And dat's just whut he was doing, too, Judge.


CLIFF
AW, she's tellin' a great big ole Georgia lie, Judge. I wasn't spendin'
no money of her'n.


WOMAN
Yes he was, Judge. There wasn't no money for him to git _but_ mine. He
ain't hit a lick of work since God been to Macon. Know whut he 'lowed
when I worry him 'bout workin'? Says he wouldn't take a job wid de
Careless Love Lumber Company, puttin' out whut make you do me lak you
do, do, do.


JUDGE
So, you goes for a sweet-back, do you?


CLIFF
Naw suh, Judge. I'd be glad to work if I could find a job.


JUDGE
How long you been outa work?


CLIFF
Seventeen years--


JUDGE
Seventeen years?
                             (to woman)
You been takin' keer of dis man for seventeen years?


WOMAN
Naw, but he been so mean to me, it seems lak seventeen
years.


JUDGE
Now you tell me just where he hurt you.


WOMAN
Judge, tell you de truth, I'm hurt all over.
                             (Rubs her buttocks)
Fact is I'm cut.


JUDGE
Did you git cut in de fracas?


WOMAN
                             (feeling the back of her left thigh
                             below her buttocks)
Not in de fracas, Judge--just below it.
                             (She starts to show the JUDGE where
                             she has been cut. He motions to stop
                             her.)


JUDGE
Stop!
                              (to Officer Simpson)
Grab him. Put him in de shade.


CLIFF
Judge, I'm unguilty! I ain't laid de weight of my hand on her in malice.
You got me 'cused of murder and I ain't harmed a child.


JUDGE
Lemme ast _you_ something. Didn't you know dat all de women in dis town
belongs to me? Beat my women and I'll stuff you in jail. 90 years. Take
'im away.
                             (CLIFF is led off to jail. JUDGE looks
                             angrily at the boy who is holding
                             hands with the pretty girl)
You runs me hot and I'm just dyin' to sit on _yo'_ case.
Whut you in here for?


BOY
Nothin'.


JUDGE
Well, whut you doin' in my court, you gater-faced rascal?


BOY
My girl wanted to see whut was goin' on, so I brought
her in.


JUDGE
Oh yeah!
                             (Smiles at GIRL)
She was usin' good sense to come see whut I'm doin', but how come _you_
come in here? You gointer have a hard time gittin' out.


BOY
I ain't done a thing. I ain't never done nothin'. I'm just as clean as a
fish, and he been bathin' all his life.


JUDGE
You ain't done nothin', hunh? Well den youse guilty of vacancy. Grab
'im, Simpson, and search 'im--and if he got any concealed weapons, I'm
gointer give 'im life-time and eight years mo'.
                             (The OFFICER seizes the boy and frisks
                             him. All he finds is a new deck of
                             cards. The JUDGE looks at them in
                             triumph.)
Unh hunh! I knowed it, one of dese skin game jelly-beans. Robbin' hard
workin' men out they money.


BOY
Judge, I ain't used 'em at all. See, dey's brand new.


JUDGE
Well, den youse charged wid totin' concealed cards and attempt to
gamble. Ten years at hard labor. Put him in de dark, Simpson, and throw
de key away.
                             (He looks at the girl and beams.)
Don't you worry bout how you gointer git home. You gointer be took home
right, 'cause I'm gointer take you myself. Bring on de next one, clerk.


CLERK
Jemima Flap-Cakes, charged with illegal possession and sale of alcoholic
liquors.


JUDGE
                             (She is a fat, black, belligerent
                             looking woman. JUDGE looks coldly at
                             her.)
Well, you heard whut he said. Is you guilty or unguilty? And I'm tellin'
you right now dat you come up befo' _me_ it's just like youse in church.
You better have a strong determination, and you better tell a good
experience.


JEMIMA
                             (Arms akimbo)
Yes, I sold it and I'll sell it again.
                             (snaps fingers and shakes hips)
How does ole booze-selling mama talk?


JUDGE
Yes, five thousand dollars and ten years in jail.
                             (Snaps fingers and shakes hips)
How does ole heavy fining papa talk?
                             (She is led away, shouting and
                             weeping)


CLERK
De Otis Blunt, charged wid stealin' a mule.
                             (LAWYER arises and comes forward with
                             the prisoner)


LAWYER
You can't convict this man. I'm here to represent him.


JUDGE
Yo' mouf might spout lak a coffee pot but I got a lawyer
                             (Looks at other lawyer)
dat kin beat your segastuatin'.
                             (Looks admiring at girl)
How am I chewin' my dictionary and minglin' my alphabets?


LAWYER
Well, I kin try, can't I?


JUDGE
Oh yeah, you kin try, but I kin see right now where he's gointer git all
de time dat God ever made dat ain't been used already. From now on.
                             (To LAWYER)
Go 'head, and spread yo' lungs all over Georgy, but he's goin' to jail!
Mules _must_ be respected.


LAWYER
                             (Striking a pose at the bar)
Your Honor,
                             (Looks at the pretty girl)
Ladies and Gentlemen--


JUDGE
Never mind 'bout dat lady. You talk yo' chat to _me_.


LAWYER
This is a clear case of syllogism! Again I say syllogism. My client is
innocent because it was a dark night when they say he stole the mule and
that's against all laws of syllogism.
                             (JUDGE looks impressed and laughs)


JUDGE
Dat ole fool do know somethin' 'bout law.


LAWYER
When George Washington was pleading de case of Marbury vs. Madison, what
did _he_ say? What _did_ he say? Scintillate, scintillate, Globule
orific. Fain would I fathom thy nature's specific. Loftily poised in
ether capacious, strongly resembling a gem carbonacious. What did
Abraham Lincoln say about mule-stealing? When torrid Phoebut refuses his
presence and ceases to lamp with fierce incandescence, then you illumine
the regions supernal, scintillate, scintillate, semper noctornal.
Syllogism, again I say syllogism.
                             (He takes his seat amid applause)


JUDGE
Man, youse a pleadin' fool. You knows yo' rules and by-laws.


OTHER LAWYER
Let me show my glory. Let me spread my habeas corpus.


JUDGE
'Tain't no use. Dis lawyer done convinced me.


OTHER LAWYER
But, lemme parade my material--


JUDGE
Parade yo' material anywhere you wants to exceptin' befo' me. Dis lil
girl wants to go home and I'm goin' with her and enjoy de consequences.
Court's adjourned.


_CURTAIN_







*** START: FULL LICENSE ***

THE FULL PROJECT GUTENBERG LICENSE
PLEASE READ THIS BEFORE YOU DISTRIBUTE OR USE THIS WORK

To protect the Project Gutenberg-tm mission of promoting the free
distribution of electronic works, by using or distributing this work
(or any other work associated in any way with the phrase "Project
Gutenberg"), you agree to comply with all the terms of the Full Project
Gutenberg-tm License (available with this file or online at
http://gutenberg.net/license).


Section 1.  General Terms of Use and Redistributing Project Gutenberg-tm
electronic works

1.A.  By reading or using any part of this Project Gutenberg-tm
electronic work, you indicate that you have read, understand, agree to
and accept all the terms of this license and intellectual property
(trademark/copyright) agreement.  If you do not agree to abide by all
the terms of this agreement, you must cease using and return or destroy
all copies of Project Gutenberg-tm electronic works in your possession.
If you paid a fee for obtaining a copy of or access to a Project
Gutenberg-tm electronic work and you do not agree to be bound by the
terms of this agreement, you may obtain a refund from the person or
entity to whom you paid the fee as set forth in paragraph 1.E.8.

1.B.  "Project Gutenberg" is a registered trademark.  It may only be
used on or associated in any way with an electronic work by people who
agree to be bound by the terms of this agreement.  There are a few
things that you can do with most Project Gutenberg-tm electronic works
even without complying with the full terms of this agreement.  See
paragraph 1.C below.  There are a lot of things you can do with Project
Gutenberg-tm electronic works if you follow the terms of this agreement
and help preserve free future access to Project Gutenberg-tm electronic
works.  See paragraph 1.E below.

1.C.  The Project Gutenberg Literary Archive Foundation ("the Foundation"
or PGLAF), owns a compilation copyright in the collection of Project
Gutenberg-tm electronic works.  Nearly all the individual works in the
collection are in the public domain in the United States.  If an
individual work is in the public domain in the United States and you are
located in the United States, we do not claim a right to prevent you from
copying, distributing, performing, displaying or creating derivative
works based on the work as long as all references to Project Gutenberg
are removed.  Of course, we hope that you will support the Project
Gutenberg-tm mission of promoting free access to electronic works by
freely sharing Project Gutenberg-tm works in compliance with the terms of
this agreement for keeping the Project Gutenberg-tm name associated with
the work.  You can easily comply with the terms of this agreement by
keeping this work in the same format with its attached full Project
Gutenberg-tm License when you share it without charge with others.

1.D.  The copyright laws of the place where you are located also govern
what you can do with this work.  Copyright laws in most countries are in
a constant state of change.  If you are outside the United States, check
the laws of your country in addition to the terms of this agreement
before downloading, copying, displaying, performing, distributing or
creating derivative works based on this work or any other Project
Gutenberg-tm work.  The Foundation makes no representations concerning
the copyright status of any work in any country outside the United
States.

1.E.  Unless you have removed all references to Project Gutenberg:

1.E.1.  The following sentence, with active links to, or other immediate
access to, the full Project Gutenberg-tm License must appear prominently
whenever any copy of a Project Gutenberg-tm work (any work on which the
phrase "Project Gutenberg" appears, or with which the phrase "Project
Gutenberg" is associated) is accessed, displayed, performed, viewed,
copied or distributed:

This eBook is for the use of anyone anywhere at no cost and with
almost no restrictions whatsoever.  You may copy it, give it away or
re-use it under the terms of the Project Gutenberg License included
with this eBook or online at www.gutenberg.net

1.E.2.  If an individual Project Gutenberg-tm electronic work is derived
from the public domain (does not contain a notice indicating that it is
posted with permission of the copyright holder), the work can be copied
and distributed to anyone in the United States without paying any fees
or charges.  If you are redistributing or providing access to a work
with the phrase "Project Gutenberg" associated with or appearing on the
work, you must comply either with the requirements of paragraphs 1.E.1
through 1.E.7 or obtain permission for the use of the work and the
Project Gutenberg-tm trademark as set forth in paragraphs 1.E.8 or
1.E.9.

1.E.3.  If an individual Project Gutenberg-tm electronic work is posted
with the permission of the copyright holder, your use and distribution
must comply with both paragraphs 1.E.1 through 1.E.7 and any additional
terms imposed by the copyright holder.  Additional terms will be linked
to the Project Gutenberg-tm License for all works posted with the
permission of the copyright holder found at the beginning of this work.

1.E.4.  Do not unlink or detach or remove the full Project Gutenberg-tm
License terms from this work, or any files containing a part of this
work or any other work associated with Project Gutenberg-tm.

1.E.5.  Do not copy, display, perform, distribute or redistribute this
electronic work, or any part of this electronic work, without
prominently displaying the sentence set forth in paragraph 1.E.1 with
active links or immediate access to the full terms of the Project
Gutenberg-tm License.

1.E.6.  You may convert to and distribute this work in any binary,
compressed, marked up, nonproprietary or proprietary form, including any
word processing or hypertext form.  However, if you provide access to or
distribute copies of a Project Gutenberg-tm work in a format other than
"Plain Vanilla ASCII" or other format used in the official version
posted on the official Project Gutenberg-tm web site (www.gutenberg.net),
you must, at no additional cost, fee or expense to the user, provide a
copy, a means of exporting a copy, or a means of obtaining a copy upon
request, of the work in its original "Plain Vanilla ASCII" or other
form.  Any alternate format must include the full Project Gutenberg-tm
License as specified in paragraph 1.E.1.

1.E.7.  Do not charge a fee for access to, viewing, displaying,
performing, copying or distributing any Project Gutenberg-tm works
unless you comply with paragraph 1.E.8 or 1.E.9.

1.E.8.  You may charge a reasonable fee for copies of or providing
access to or distributing Project Gutenberg-tm electronic works provided
that

- You pay a royalty fee of 20% of the gross profits you derive from
     the use of Project Gutenberg-tm works calculated using the method
     you already use to calculate your applicable taxes.  The fee is
     owed to the owner of the Project Gutenberg-tm trademark, but he
     has agreed to donate royalties under this paragraph to the
     Project Gutenberg Literary Archive Foundation.  Royalty payments
     must be paid within 60 days following each date on which you
     prepare (or are legally required to prepare) your periodic tax
     returns.  Royalty payments should be clearly marked as such and
     sent to the Project Gutenberg Literary Archive Foundation at the
     address specified in Section 4, "Information about donations to
     the Project Gutenberg Literary Archive Foundation."

- You provide a full refund of any money paid by a user who notifies
     you in writing (or by e-mail) within 30 days of receipt that s/he
     does not agree to the terms of the full Project Gutenberg-tm
     License.  You must require such a user to return or
     destroy all copies of the works possessed in a physical medium
     and discontinue all use of and all access to other copies of
     Project Gutenberg-tm works.

- You provide, in accordance with paragraph 1.F.3, a full refund of any
     money paid for a work or a replacement copy, if a defect in the
     electronic work is discovered and reported to you within 90 days
     of receipt of the work.

- You comply with all other terms of this agreement for free
     distribution of Project Gutenberg-tm works.

1.E.9.  If you wish to charge a fee or distribute a Project Gutenberg-tm
electronic work or group of works on different terms than are set
forth in this agreement, you must obtain permission in writing from
both the Project Gutenberg Literary Archive Foundation and Michael
Hart, the owner of the Project Gutenberg-tm trademark.  Contact the
Foundation as set forth in Section 3 below.

1.F.

1.F.1.  Project Gutenberg volunteers and employees expend considerable
effort to identify, do copyright research on, transcribe and proofread
public domain works in creating the Project Gutenberg-tm
collection.  Despite these efforts, Project Gutenberg-tm electronic
works, and the medium on which they may be stored, may contain
"Defects," such as, but not limited to, incomplete, inaccurate or
corrupt data, transcription errors, a copyright or other intellectual
property infringement, a defective or damaged disk or other medium, a
computer virus, or computer codes that damage or cannot be read by
your equipment.

1.F.2.  LIMITED WARRANTY, DISCLAIMER OF DAMAGES - Except for the "Right
of Replacement or Refund" described in paragraph 1.F.3, the Project
Gutenberg Literary Archive Foundation, the owner of the Project
Gutenberg-tm trademark, and any other party distributing a Project
Gutenberg-tm electronic work under this agreement, disclaim all
liability to you for damages, costs and expenses, including legal
fees.  YOU AGREE THAT YOU HAVE NO REMEDIES FOR NEGLIGENCE, STRICT
LIABILITY, BREACH OF WARRANTY OR BREACH OF CONTRACT EXCEPT THOSE
PROVIDED IN PARAGRAPH F3.  YOU AGREE THAT THE FOUNDATION, THE
TRADEMARK OWNER, AND ANY DISTRIBUTOR UNDER THIS AGREEMENT WILL NOT BE
LIABLE TO YOU FOR ACTUAL, DIRECT, INDIRECT, CONSEQUENTIAL, PUNITIVE OR
INCIDENTAL DAMAGES EVEN IF YOU GIVE NOTICE OF THE POSSIBILITY OF SUCH
DAMAGE.

1.F.3.  LIMITED RIGHT OF REPLACEMENT OR REFUND - If you discover a
defect in this electronic work within 90 days of receiving it, you can
receive a refund of the money (if any) you paid for it by sending a
written explanation to the person you received the work from.  If you
received the work on a physical medium, you must return the medium with
your written explanation.  The person or entity that provided you with
the defective work may elect to provide a replacement copy in lieu of a
refund.  If you received the work electronically, the person or entity
providing it to you may choose to give you a second opportunity to
receive the work electronically in lieu of a refund.  If the second copy
is also defective, you may demand a refund in writing without further
opportunities to fix the problem.

1.F.4.  Except for the limited right of replacement or refund set forth
in paragraph 1.F.3, this work is provided to you 'AS-IS', WITH NO OTHER
WARRANTIES OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
WARRANTIES OF MERCHANTIBILITY OR FITNESS FOR ANY PURPOSE.

1.F.5.  Some states do not allow disclaimers of certain implied
warranties or the exclusion or limitation of certain types of damages.
If any disclaimer or limitation set forth in this agreement violates the
law of the state applicable to this agreement, the agreement shall be
interpreted to make the maximum disclaimer or limitation permitted by
the applicable state law.  The invalidity or unenforceability of any
provision of this agreement shall not void the remaining provisions.

1.F.6.  INDEMNITY - You agree to indemnify and hold the Foundation, the
trademark owner, any agent or employee of the Foundation, anyone
providing copies of Project Gutenberg-tm electronic works in accordance
with this agreement, and any volunteers associated with the production,
promotion and distribution of Project Gutenberg-tm electronic works,
harmless from all liability, costs and expenses, including legal fees,
that arise directly or indirectly from any of the following which you do
or cause to occur: (a) distribution of this or any Project Gutenberg-tm
work, (b) alteration, modification, or additions or deletions to any
Project Gutenberg-tm work, and (c) any Defect you cause.


Section  2.  Information about the Mission of Project Gutenberg-tm

Project Gutenberg-tm is synonymous with the free distribution of
electronic works in formats readable by the widest variety of computers
including obsolete, old, middle-aged and new computers.  It exists
because of the efforts of hundreds of volunteers and donations from
people in all walks of life.

Volunteers and financial support to provide volunteers with the
assistance they need, is critical to reaching Project Gutenberg-tm's
goals and ensuring that the Project Gutenberg-tm collection will
remain freely available for generations to come.  In 2001, the Project
Gutenberg Literary Archive Foundation was created to provide a secure
and permanent future for Project Gutenberg-tm and future generations.
To learn more about the Project Gutenberg Literary Archive Foundation
and how your efforts and donations can help, see Sections 3 and 4
and the Foundation web page at http://www.gutenberg.net/fundraising/pglaf.


Section 3.  Information about the Project Gutenberg Literary Archive
Foundation

The Project Gutenberg Literary Archive Foundation is a non profit
501(c)(3) educational corporation organized under the laws of the
state of Mississippi and granted tax exempt status by the Internal
Revenue Service.  The Foundation's EIN or federal tax identification
number is 64-6221541.  Contributions to the Project Gutenberg
Literary Archive Foundation are tax deductible to the full extent
permitted by U.S. federal laws and your state's laws.

The Foundation's principal office is located at 4557 Melan Dr. S.
Fairbanks, AK, 99712., but its volunteers and employees are scattered
throughout numerous locations.  Its business office is located at
809 North 1500 West, Salt Lake City, UT 84116, (801) 596-1887, email
business@pglaf.org.  Email contact links and up to date contact
information can be found at the Foundation's web site and official
page at http://www.gutenberg.net/about/contact

For additional contact information:
     Dr. Gregory B. Newby
     Chief Executive and Director
     gbnewby@pglaf.org

Section 4.  Information about Donations to the Project Gutenberg
Literary Archive Foundation

Project Gutenberg-tm depends upon and cannot survive without wide
spread public support and donations to carry out its mission of
increasing the number of public domain and licensed works that can be
freely distributed in machine readable form accessible by the widest
array of equipment including outdated equipment.  Many small donations
($1 to $5,000) are particularly important to maintaining tax exempt
status with the IRS.

The Foundation is committed to complying with the laws regulating
charities and charitable donations in all 50 states of the United
States.  Compliance requirements are not uniform and it takes a
considerable effort, much paperwork and many fees to meet and keep up
with these requirements.  We do not solicit donations in locations
where we have not received written confirmation of compliance.  To
SEND DONATIONS or determine the status of compliance for any
particular state visit http://www.gutenberg.net/fundraising/donate

While we cannot and do not solicit contributions from states where we
have not met the solicitation requirements, we know of no prohibition
against accepting unsolicited donations from donors in such states who
approach us with offers to donate.

International donations are gratefully accepted, but we cannot make
any statements concerning tax treatment of donations received from
outside the United States.  U.S. laws alone swamp our small staff.

Please check the Project Gutenberg Web pages for current donation
methods and addresses.  Donations are accepted in a number of other
ways including including checks, online payments and credit card
donations.  To donate, please visit:
http://www.gutenberg.net/fundraising/donate


Section 5.  General Information About Project Gutenberg-tm electronic
works.

Professor Michael S. Hart is the originator of the Project Gutenberg-tm
concept of a library of electronic works that could be freely shared
with anyone.  For thirty years, he produced and distributed Project
Gutenberg-tm eBooks with only a loose network of volunteer support.

Project Gutenberg-tm eBooks are often created from several printed
editions, all of which are confirmed as Public Domain in the U.S.
unless a copyright notice is included.  Thus, we do not necessarily
keep eBooks in compliance with any particular paper edition.

Most people start at our Web site which has the main PG search facility:

     http://www.gutenberg.net

This Web site includes information about Project Gutenberg-tm,
including how to make donations to the Project Gutenberg Literary
Archive Foundation, how to help produce our new eBooks, and how to
subscribe to our email newsletter to hear about new eBooks.
`

	b["hurston-poker.txt"] = `
The Project Gutenberg EBook of Poker!, by Zora Hurston

This eBook is for the use of anyone anywhere at no cost and with
almost no restrictions whatsoever.  You may copy it, give it away or
re-use it under the terms of the Project Gutenberg License included
with this eBook or online at www.gutenberg.net


Title: Poker!

Author: Zora Hurston

Release Date: May 25, 2005 [EBook #15902]

Language: English


*** START OF THIS PROJECT GUTENBERG EBOOK POKER! ***




Produced by Library of Congress, American Memory Project,
Charlene Taylor and the Online Distributed Proofreading
Team.





[Transcriber's Note: This play transcribed from an original manuscript.
There are pencilled notations possibly by Ms. Hurston herself. These
pencilled edits have been transcribed as *[Handwritten: (text)]]




Copyright 1931 by Zora *[Handwritten: Neale] Hurston




POKER!


Time--Present

Place--New York

Cast of characters--
   Nunkie
   Too-Sweet
   Peckerwood
   Black Baby
   Sack Daddy
   Tush Hawg
   Aunt Dilsey

SCENE--

     A shabby front room in a shotgun house.

     A door covered by dingy portieres upstage C. Small panel
     window in side Wall L. Plain centre table with chairs drawn up
     about it. Gaudy calendars on wall. Battered piano against wall
     R. Kerosene lamp with reflector against wall on either side of
     room.

     At rise of curtain NUNKIE is at piano playing.... Others at
     table with small stacks of chips before each man. TUSH HAWG is
     seated at table so that he faces audience. He is expertly
     riffing the cards ... looks over his shoulder and speaks to
     NUNKIE.


TUSH HAWG
Come on here, Nunkie--and take a hand! You're holding up the game. You
been woofin' round here about the poker you can play--now do it!

NUNKIE
Yeah, I plays poker. I plays the piano and Gawd knows I plays the devil.
I'm Uncle Bob with a wooden leg!*[Handwritten: Last sentence crossed out
in pencil in manuscript.]

BLACK BABY
Aw, you can be had! Come on and get in the game! My
britches is cryin' for your money! Come on, don't give
the healer no trouble!*[Handwritten: last sentence crossed out in pencil]

NUNKIE
Soon as I play the deck I'm comin' and take you alls money! Don' rush
me.

 Ace means the first time that I met you
 Duece means there was nobody there but us two
 Trey means the third party--Charlie was his name
 Four spot means the fourth time you tried that same old game--
Five spot means five years you played me for a clown
Six spot means six feet of earth when the deal goes down
Now I'm holding the seven spot for each day of the week
Eight means eight hours that she Sheba-ed with your Sheik--
Nine spot means nine hours that I work hard every day--
Ten spot means tenth of every month I brought you home my pay--
The Jack is three-card Charlie who played me for a goat
The Queen, that's my pretty Mama, also trying to cut my throat--
The King stands for Sweet Papa Nunkie and he's goin' to wear the crown,
So be careful you all ain't broke when the deal goes down!
                              (He laughs--X'es to table, bringing
                              piano stool for seat)

TUSH HAWG
Aw now, brother, two dollars for your seat before you try to sit in this
game.

NUNKIE
                              (Laughs sheepishly--puts money
                              down--TUSH HAWG pushes stack of chips
                              toward him. Bus.)
I didn't put it down because I knew you all goin' to be puttin' it right
back in my pocket.

BECKERWOOD
Aw, Y'all go ahead and play.
                              (to TUSH HAWG)
Deal!
                              (TUSH HAWG begins to deal for draw
                              poker. The game gets tense. SACK
                              DADDY is first man at TUSH's left--he
                              throws back three cards and is dealt
                              three more)

SACK DADDY
My luck sure is rotten! My gal must be cheatin' on me. I ain't had a
pair since John Henry had a hammer!

BLACK BABY
                              (Drawing three new cards)
You might be fooling the rest with the cryin' you're doin' but I'm
squattin' for you! You're cryin' worse than cryin' Emma!

TOO-SWEET
                              (Studying his three new cards)
                              (Sings)
When yo' cards gets lucky, oh Partner, you oughter be in a rollin' game.
*[Handwritten: get you foot offa my chair etc]

AUNT DILSEY
                              (Enters through portieres--stands and
                              looks disapprovingly)
You all oughter be ashamed of yourself, gamblin' and carryin' on like
this!

BLACK BABY
Aw, this ain't no harm, Aunt Dilsey! You go on back to bed and git your
night's rest.

AUNT DILSEY
No harm! I know all about these no-harm sins! If you don't stop this
card playin', all of you all goin' to die and go to Hell.
                              (Shakes warning finger--exits through
                              portieres--while she is talking the
                              men have been hiding cards out of
                              their hands and pulling aces out of
                              sleeves and vest pockets and
                              shoes--it is done quickly, one does
                              not see the other do it)

NUNKIE
                              (Shoving a chip forward)
A dollar!

SACK DADDY
Raise you two!

BLACK BABY
I don't like to strain with nobody but it's goin' to cost you five. Come
on, you shag-nags! This hand I got is enough to pull a country man into
town. *[Handwritten: Last sentence crossed through in pencil.]

TOO-SWEET
You all act like you're spuddin'! Bet some money! Put your money where
your mouth is *[Handwritten: els my fist where yo mouf is.]

TUSH HAWG
Twenty-five dollars to keep my company! Dog-gone, I'm spreadin' my
knots!

SACK DADDY
And I bet you a fat man I'll take your money--I call you.
                              (Turns up his cards--he has four aces
                              and king)

TUSH HAWG
                              (showing his cards)
Youse a liar! I ain't dealt you no aces. Don't try to carry the Pam-Pam
to me 'cause I'll gently chain-gang for you!

SACK DADDY
Oh yeah! I ain't goin' to fit no jail for you and nobody else. I'm to
get me a green club and season it over your head. Then I'll give my case
to Miss Bush and let Mother Green stand my bond! I got deal them aces!

NUNKIE
That's a lie! Both of you is lyin'! Lyin' like the cross-ties from New
York to Key West! How can you all hold aces when I got four? Somebody is
goin' to West hell before midnight!

BECKERWOOD
Don't you woof at Tush Hawg. If you do I'm goin' to bust hell wide open
with a man!

BLACK BABY
                              (Pulls out razor--Bus.)
My chop-axe tells me I got the only clean aces they is on this table!
Before I'll leave you all rob me outa my money, I'm goin' to die it off!

TOO-SWEET
I promised the devil one man and I'm goin' to give him five!
                              (Draws gun)

TUSH HAWG
Don't draw your bosom on me! God sent me a pistol and I'm goin' to send
him a man!
                              (FIRES. Bus. for all)

AUNT DILSEY
                              (Enters after shooting bus. Stands.
                              Bus. drops to chair)
They wouldn't lissen--
                              (Looks men over--Bus.)
It sure is goin' to be a whole lot tougher in hell now!




CURTAIN





End of the Project Gutenberg EBook of Poker!, by Zora Hurston

*** END OF THIS PROJECT GUTENBERG EBOOK POKER! ***

***** This file should be named 15902.txt or 15902.zip *****
This and all associated files of various formats will be found in:
        http://www.gutenberg.org/1/5/9/0/15902/

Produced by Library of Congress, American Memory Project,
Charlene Taylor and the Online Distributed Proofreading
Team.


Updated editions will replace the previous one--the old editions
will be renamed.

Creating the works from public domain print editions means that no
one owns a United States copyright in these works, so the Foundation
(and you!) can copy and distribute it in the United States without
permission and without paying copyright royalties.  Special rules,
set forth in the General Terms of Use part of this license, apply to
copying and distributing Project Gutenberg-tm electronic works to
protect the PROJECT GUTENBERG-tm concept and trademark.  Project
Gutenberg is a registered trademark, and may not be used if you
charge for the eBooks, unless you receive specific permission.  If you
do not charge anything for copies of this eBook, complying with the
rules is very easy.  You may use this eBook for nearly any purpose
such as creation of derivative works, reports, performances and
research.  They may be modified and printed and given away--you may do
practically ANYTHING with public domain eBooks.  Redistribution is
subject to the trademark license, especially commercial
redistribution.



*** START: FULL LICENSE ***

THE FULL PROJECT GUTENBERG LICENSE
PLEASE READ THIS BEFORE YOU DISTRIBUTE OR USE THIS WORK

To protect the Project Gutenberg-tm mission of promoting the free
distribution of electronic works, by using or distributing this work
(or any other work associated in any way with the phrase "Project
Gutenberg"), you agree to comply with all the terms of the Full Project
Gutenberg-tm License (available with this file or online at
http://gutenberg.net/license).


Section 1.  General Terms of Use and Redistributing Project Gutenberg-tm
electronic works

1.A.  By reading or using any part of this Project Gutenberg-tm
electronic work, you indicate that you have read, understand, agree to
and accept all the terms of this license and intellectual property
(trademark/copyright) agreement.  If you do not agree to abide by all
the terms of this agreement, you must cease using and return or destroy
all copies of Project Gutenberg-tm electronic works in your possession.
If you paid a fee for obtaining a copy of or access to a Project
Gutenberg-tm electronic work and you do not agree to be bound by the
terms of this agreement, you may obtain a refund from the person or
entity to whom you paid the fee as set forth in paragraph 1.E.8.

1.B.  "Project Gutenberg" is a registered trademark.  It may only be
used on or associated in any way with an electronic work by people who
agree to be bound by the terms of this agreement.  There are a few
things that you can do with most Project Gutenberg-tm electronic works
even without complying with the full terms of this agreement.  See
paragraph 1.C below.  There are a lot of things you can do with Project
Gutenberg-tm electronic works if you follow the terms of this agreement
and help preserve free future access to Project Gutenberg-tm electronic
works.  See paragraph 1.E below.

1.C.  The Project Gutenberg Literary Archive Foundation ("the Foundation"
or PGLAF), owns a compilation copyright in the collection of Project
Gutenberg-tm electronic works.  Nearly all the individual works in the
collection are in the public domain in the United States.  If an
individual work is in the public domain in the United States and you are
located in the United States, we do not claim a right to prevent you from
copying, distributing, performing, displaying or creating derivative
works based on the work as long as all references to Project Gutenberg
are removed.  Of course, we hope that you will support the Project
Gutenberg-tm mission of promoting free access to electronic works by
freely sharing Project Gutenberg-tm works in compliance with the terms of
this agreement for keeping the Project Gutenberg-tm name associated with
the work.  You can easily comply with the terms of this agreement by
keeping this work in the same format with its attached full Project
Gutenberg-tm License when you share it without charge with others.

1.D.  The copyright laws of the place where you are located also govern
what you can do with this work.  Copyright laws in most countries are in
a constant state of change.  If you are outside the United States, check
the laws of your country in addition to the terms of this agreement
before downloading, copying, displaying, performing, distributing or
creating derivative works based on this work or any other Project
Gutenberg-tm work.  The Foundation makes no representations concerning
the copyright status of any work in any country outside the United
States.

1.E.  Unless you have removed all references to Project Gutenberg:

1.E.1.  The following sentence, with active links to, or other immediate
access to, the full Project Gutenberg-tm License must appear prominently
whenever any copy of a Project Gutenberg-tm work (any work on which the
phrase "Project Gutenberg" appears, or with which the phrase "Project
Gutenberg" is associated) is accessed, displayed, performed, viewed,
copied or distributed:

This eBook is for the use of anyone anywhere at no cost and with
almost no restrictions whatsoever.  You may copy it, give it away or
re-use it under the terms of the Project Gutenberg License included
with this eBook or online at www.gutenberg.net

1.E.2.  If an individual Project Gutenberg-tm electronic work is derived
from the public domain (does not contain a notice indicating that it is
posted with permission of the copyright holder), the work can be copied
and distributed to anyone in the United States without paying any fees
or charges.  If you are redistributing or providing access to a work
with the phrase "Project Gutenberg" associated with or appearing on the
work, you must comply either with the requirements of paragraphs 1.E.1
through 1.E.7 or obtain permission for the use of the work and the
Project Gutenberg-tm trademark as set forth in paragraphs 1.E.8 or
1.E.9.

1.E.3.  If an individual Project Gutenberg-tm electronic work is posted
with the permission of the copyright holder, your use and distribution
must comply with both paragraphs 1.E.1 through 1.E.7 and any additional
terms imposed by the copyright holder.  Additional terms will be linked
to the Project Gutenberg-tm License for all works posted with the
permission of the copyright holder found at the beginning of this work.

1.E.4.  Do not unlink or detach or remove the full Project Gutenberg-tm
License terms from this work, or any files containing a part of this
work or any other work associated with Project Gutenberg-tm.

1.E.5.  Do not copy, display, perform, distribute or redistribute this
electronic work, or any part of this electronic work, without
prominently displaying the sentence set forth in paragraph 1.E.1 with
active links or immediate access to the full terms of the Project
Gutenberg-tm License.

1.E.6.  You may convert to and distribute this work in any binary,
compressed, marked up, nonproprietary or proprietary form, including any
word processing or hypertext form.  However, if you provide access to or
distribute copies of a Project Gutenberg-tm work in a format other than
"Plain Vanilla ASCII" or other format used in the official version
posted on the official Project Gutenberg-tm web site (www.gutenberg.net),
you must, at no additional cost, fee or expense to the user, provide a
copy, a means of exporting a copy, or a means of obtaining a copy upon
request, of the work in its original "Plain Vanilla ASCII" or other
form.  Any alternate format must include the full Project Gutenberg-tm
License as specified in paragraph 1.E.1.

1.E.7.  Do not charge a fee for access to, viewing, displaying,
performing, copying or distributing any Project Gutenberg-tm works
unless you comply with paragraph 1.E.8 or 1.E.9.

1.E.8.  You may charge a reasonable fee for copies of or providing
access to or distributing Project Gutenberg-tm electronic works provided
that

- You pay a royalty fee of 20% of the gross profits you derive from
     the use of Project Gutenberg-tm works calculated using the method
     you already use to calculate your applicable taxes.  The fee is
     owed to the owner of the Project Gutenberg-tm trademark, but he
     has agreed to donate royalties under this paragraph to the
     Project Gutenberg Literary Archive Foundation.  Royalty payments
     must be paid within 60 days following each date on which you
     prepare (or are legally required to prepare) your periodic tax
     returns.  Royalty payments should be clearly marked as such and
     sent to the Project Gutenberg Literary Archive Foundation at the
     address specified in Section 4, "Information about donations to
     the Project Gutenberg Literary Archive Foundation."

- You provide a full refund of any money paid by a user who notifies
     you in writing (or by e-mail) within 30 days of receipt that s/he
     does not agree to the terms of the full Project Gutenberg-tm
     License.  You must require such a user to return or
     destroy all copies of the works possessed in a physical medium
     and discontinue all use of and all access to other copies of
     Project Gutenberg-tm works.

- You provide, in accordance with paragraph 1.F.3, a full refund of any
     money paid for a work or a replacement copy, if a defect in the
     electronic work is discovered and reported to you within 90 days
     of receipt of the work.

- You comply with all other terms of this agreement for free
     distribution of Project Gutenberg-tm works.

1.E.9.  If you wish to charge a fee or distribute a Project Gutenberg-tm
electronic work or group of works on different terms than are set
forth in this agreement, you must obtain permission in writing from
both the Project Gutenberg Literary Archive Foundation and Michael
Hart, the owner of the Project Gutenberg-tm trademark.  Contact the
Foundation as set forth in Section 3 below.

1.F.

1.F.1.  Project Gutenberg volunteers and employees expend considerable
effort to identify, do copyright research on, transcribe and proofread
public domain works in creating the Project Gutenberg-tm
collection.  Despite these efforts, Project Gutenberg-tm electronic
works, and the medium on which they may be stored, may contain
"Defects," such as, but not limited to, incomplete, inaccurate or
corrupt data, transcription errors, a copyright or other intellectual
property infringement, a defective or damaged disk or other medium, a
computer virus, or computer codes that damage or cannot be read by
your equipment.

1.F.2.  LIMITED WARRANTY, DISCLAIMER OF DAMAGES - Except for the "Right
of Replacement or Refund" described in paragraph 1.F.3, the Project
Gutenberg Literary Archive Foundation, the owner of the Project
Gutenberg-tm trademark, and any other party distributing a Project
Gutenberg-tm electronic work under this agreement, disclaim all
liability to you for damages, costs and expenses, including legal
fees.  YOU AGREE THAT YOU HAVE NO REMEDIES FOR NEGLIGENCE, STRICT
LIABILITY, BREACH OF WARRANTY OR BREACH OF CONTRACT EXCEPT THOSE
PROVIDED IN PARAGRAPH F3.  YOU AGREE THAT THE FOUNDATION, THE
TRADEMARK OWNER, AND ANY DISTRIBUTOR UNDER THIS AGREEMENT WILL NOT BE
LIABLE TO YOU FOR ACTUAL, DIRECT, INDIRECT, CONSEQUENTIAL, PUNITIVE OR
INCIDENTAL DAMAGES EVEN IF YOU GIVE NOTICE OF THE POSSIBILITY OF SUCH
DAMAGE.

1.F.3.  LIMITED RIGHT OF REPLACEMENT OR REFUND - If you discover a
defect in this electronic work within 90 days of receiving it, you can
receive a refund of the money (if any) you paid for it by sending a
written explanation to the person you received the work from.  If you
received the work on a physical medium, you must return the medium with
your written explanation.  The person or entity that provided you with
the defective work may elect to provide a replacement copy in lieu of a
refund.  If you received the work electronically, the person or entity
providing it to you may choose to give you a second opportunity to
receive the work electronically in lieu of a refund.  If the second copy
is also defective, you may demand a refund in writing without further
opportunities to fix the problem.

1.F.4.  Except for the limited right of replacement or refund set forth
in paragraph 1.F.3, this work is provided to you 'AS-IS', WITH NO OTHER
WARRANTIES OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
WARRANTIES OF MERCHANTIBILITY OR FITNESS FOR ANY PURPOSE.

1.F.5.  Some states do not allow disclaimers of certain implied
warranties or the exclusion or limitation of certain types of damages.
If any disclaimer or limitation set forth in this agreement violates the
law of the state applicable to this agreement, the agreement shall be
interpreted to make the maximum disclaimer or limitation permitted by
the applicable state law.  The invalidity or unenforceability of any
provision of this agreement shall not void the remaining provisions.

1.F.6.  INDEMNITY - You agree to indemnify and hold the Foundation, the
trademark owner, any agent or employee of the Foundation, anyone
providing copies of Project Gutenberg-tm electronic works in accordance
with this agreement, and any volunteers associated with the production,
promotion and distribution of Project Gutenberg-tm electronic works,
harmless from all liability, costs and expenses, including legal fees,
that arise directly or indirectly from any of the following which you do
or cause to occur: (a) distribution of this or any Project Gutenberg-tm
work, (b) alteration, modification, or additions or deletions to any
Project Gutenberg-tm work, and (c) any Defect you cause.


Section  2.  Information about the Mission of Project Gutenberg-tm

Project Gutenberg-tm is synonymous with the free distribution of
electronic works in formats readable by the widest variety of computers
including obsolete, old, middle-aged and new computers.  It exists
because of the efforts of hundreds of volunteers and donations from
people in all walks of life.

Volunteers and financial support to provide volunteers with the
assistance they need, is critical to reaching Project Gutenberg-tm's
goals and ensuring that the Project Gutenberg-tm collection will
remain freely available for generations to come.  In 2001, the Project
Gutenberg Literary Archive Foundation was created to provide a secure
and permanent future for Project Gutenberg-tm and future generations.
To learn more about the Project Gutenberg Literary Archive Foundation
and how your efforts and donations can help, see Sections 3 and 4
and the Foundation web page at http://www.pglaf.org.


Section 3.  Information about the Project Gutenberg Literary Archive
Foundation

The Project Gutenberg Literary Archive Foundation is a non profit
501(c)(3) educational corporation organized under the laws of the
state of Mississippi and granted tax exempt status by the Internal
Revenue Service.  The Foundation's EIN or federal tax identification
number is 64-6221541.  Its 501(c)(3) letter is posted at
http://pglaf.org/fundraising.  Contributions to the Project Gutenberg
Literary Archive Foundation are tax deductible to the full extent
permitted by U.S. federal laws and your state's laws.

The Foundation's principal office is located at 4557 Melan Dr. S.
Fairbanks, AK, 99712., but its volunteers and employees are scattered
throughout numerous locations.  Its business office is located at
809 North 1500 West, Salt Lake City, UT 84116, (801) 596-1887, email
business@pglaf.org.  Email contact links and up to date contact
information can be found at the Foundation's web site and official
page at http://pglaf.org

For additional contact information:
     Dr. Gregory B. Newby
     Chief Executive and Director
     gbnewby@pglaf.org

Section 4.  Information about Donations to the Project Gutenberg
Literary Archive Foundation

Project Gutenberg-tm depends upon and cannot survive without wide
spread public support and donations to carry out its mission of
increasing the number of public domain and licensed works that can be
freely distributed in machine readable form accessible by the widest
array of equipment including outdated equipment.  Many small donations
($1 to $5,000) are particularly important to maintaining tax exempt
status with the IRS.

The Foundation is committed to complying with the laws regulating
charities and charitable donations in all 50 states of the United
States.  Compliance requirements are not uniform and it takes a
considerable effort, much paperwork and many fees to meet and keep up
with these requirements.  We do not solicit donations in locations
where we have not received written confirmation of compliance.  To
SEND DONATIONS or determine the status of compliance for any
particular state visit http://pglaf.org

While we cannot and do not solicit contributions from states where we
have not met the solicitation requirements, we know of no prohibition
against accepting unsolicited donations from donors in such states who
approach us with offers to donate.

International donations are gratefully accepted, but we cannot make
any statements concerning tax treatment of donations received from
outside the United States.  U.S. laws alone swamp our small staff.

Please check the Project Gutenberg Web pages for current donation
methods and addresses.  Donations are accepted in a number of other
ways including including checks, online payments and credit card
donations.  To donate, please visit: http://pglaf.org/donate


Section 5.  General Information About Project Gutenberg-tm electronic
works.

Professor Michael S. Hart is the originator of the Project Gutenberg-tm
concept of a library of electronic works that could be freely shared
with anyone.  For thirty years, he produced and distributed Project
Gutenberg-tm eBooks with only a loose network of volunteer support.

Project Gutenberg-tm eBooks are often created from several printed
editions, all of which are confirmed as Public Domain in the U.S.
unless a copyright notice is included.  Thus, we do not necessarily
keep eBooks in compliance with any particular paper edition.

Most people start at our Web site which has the main PG search facility:

     http://www.gutenberg.net

This Web site includes information about Project Gutenberg-tm,
including how to make donations to the Project Gutenberg Literary
Archive Foundation, how to help produce our new eBooks, and how to
subscribe to our email newsletter to hear about new eBooks.

*** END: FULL LICENSE ***
`

	b["hurston-woofing.txt"] = `
_"WOOFING"_

By

_ZORA *[Handwritten (Neale)] HURSTON_



_"WOOFING"_


TIME:     Present.

PLACE:    Negro Street in Waycross, Ga.

PERSONS:  Loungers, two children, guitar
          players, women, band--

SETTING:  Porch and side walk, etc.

ACTION:   Thru the open window of 'one' of
          the shacks a WOMAN is discovered
          ironing. A MAN is sitting on the
          floor of the porch asleep.  She
          hums a bar or two, then comes to
          the window and calls to the man.



Woman
Good Black, why don't you git up from dere and carry dese white folks
clothes home? You always want money but you wouldn't hit a lick at a
snake!

Man
Aw, shut up woman. I'm tired of hearin' bout dem white folks clothes. I
don't keer if dey never git 'em.

Woman
You better keer! Dese very clothes took and brought _you_ out de crack.
'Cause de first time I saw you you was so hungry till you was walkin'
lap-legged. Man, you had de white-mouf, you was so hungry.

                             (Enter another MAN leisurely. Good
                             Black sees him and calls)

Good Black
Hey, Cliffert, where you headed for?

Cliffert
Oh, no where in particular.

Good Black
Come here then, fish, and lemme bend a checker game over yo' head. Come
on, youse my fish.

Cliffert
                             (Comes to the porch and sits)
Git de checkers and I'll have you any, some or none. I push a mean
chuck-a-luck myself.

Woman
                             (Voice inside quarreling)
Dress up and strut around! Yes! Play checkers? Yes! Eat? Yes! Work? No!!

                             (The game starts. A period of silence
                             in which they indicate their
                             concentration by frowns, cautious
                             moves, head scratching. GOOD BLACK is
                             pointing his index finger over the
                             board indicating moves. He wig-wags,
                             starts to move, scratches his head
                             thoroughly, changes his mind and fools
                             around without moving)

Cliffert
Police! Police! Come here and make dis man move!

Good Black
Aw, I got plenty moves.
                             (Scratches his head)
Jus' tryin' to see which one I want to make. But when I do move, it's
gointer be just too bad for you.

                             (A guitar is heard off stage and
                             Cliffert brightens. He cups his hand
                             and calls)

Cliffert
Hey Lonnie! Come here! Ha, ha, ha! I got me a fish.
                             (Enter LONNIE picking "East Coast" on
                             his box and stands watching the game.
                             He ceases to play as he stops walking)
Ha, ha! You see ol' Good Black goes for a hard guy. He tries to know
more than a mule and a mule's head longer'n his'n. Ha, ha! I set a trap
for him and he fell right in it. Trying to ride de britches! _Now_ look
at him.

Good Black
Aw, shut up! You tryin' to show yo' grandma how to milk ducks. You can't
beat me playin' no checkers.
                             (Scratches his head again)
Just watch me show my glory.

Woman
                             (Leans out of window)
Good Black! When you gointer come git dese clothes!
                             (He does not answer, he is trying to
                             concentrate)

Lonnie
You got him Cliffert. You got him in Louisville Loop. He's yo' fish all
right.

Cliffert
                             (Boastfully)
Man, didn't I push a mean chuck-a-luck dat time! I'm good, better, and
best. Move, Man!
                             (To Good Black)
I tole you not to do it.

Good Black
All dat noise ain't playin' checkers. You just wait till I make my move.

Woman
All right, now, Mr. Nappy-Chin! I don't want to have to call you no mo'
to come keer dese white folks clothes! I'm tired of takin' and takin'
affa you! My belly's full clear up to de neck. I don't need no lazy coon
lak you nohow. I'm a good woman, and I needs somebody dats gointer give
aid and assistance.

Good Black,
Aw, go head on', woman, and leave me be! Every Saturday it's de same
thing! Yo' mouth exhausting like a automobile. You worse than "cryin'
Emma". You kin whoop like de Seaboard and squall lak de Coast Line.
                             (Taps his head)
You ain't go all dat b'long to _you_, and nothin' dat b'long to nobody's
else. You better leave me 'lone before you make a bad man out of me.
Fool wid me and I'll go git me somebody else. I'm a much-right man.

Woman
Now you ain't no much right man neither. You didn't _git me_ wid no
saw-mill license--You went to de court house and paid a dollar and a
half for me. Tain't no other woman got as much right to you as I got. De
Man to tell you youse divorced befo' yo' kin play dat much-right on me!

Good Black
De Man don't have to tell me nothin'! I got divorce in my heels.

Woman
You ain't de only one dat knows where de railroad track is, I done made
up my mind, and I done promised Gabriel and a couple of other men dat
if yo' don't do no better than yo' been doin', I'm gointer pack me a
suit case and grab de first smoky thing I see. I'll be long gone.

Good Black
Aw, yo' ain't no trouble! Yo' can be had. Yo' ain't never gointer
leave me.

Woman
How come I won't? Just 'cause I been takin' keer of yo', don't make a
park ape out yo'self. I'll leave yo', just as sure as yo' snore!

Good Black
                             (Rises and hitches up his trousers)
Aw, yo' ain't gointer leave me, and if yo' go, yo' wouldn't stay, 'cause
I'm a damn sweet man, and yo' know it!

Lonnie
Hey, Hey!
                             (He begins to pick and Good Black
                             sings. Lonnie sings a line now and
                             then)

Good Black
  Yo' may leave and go to Hali-muh-fack
  But my slow drag will--uh bring yo' back
  Well yo' may go, but this will bring yo' back

  I been in de country but I moved to town
  I'm a tolo-shaker from my head on down
  Well, yo' may go, but this will bring yo' back

  Some folks call me a tolo-shaker
  It's a doggone lie I'm a back-bone breaker
  Well, yo' may go, but this will bring yo' back.

  Oh, ship on de sea, boat on de ocean
  I raise hell when I take a notion
  Well, yo' may go, but this will bring yo' back.

  Oh, who do, who do, who do wackin'
  Wid my hells a' poppin' and my toe-nails crackin'
  Well, yo' may go, but this will bring yo' back.

Woman
Dat's all right too, pap but if yo' can't make me tote dese clothes
home, don't bring de mess up. Yo'se abstifically a humbug.

Cliffert
Man, come on back here and move, or else own up to de folks yo' can't
push no checkers wid me.
                             (He sits and begins to lay out moves
                             with his fingers and scratch his head.
                             Enter another MAN and stands akimbo
                             looking over Cliff's shoulder)

Cliff
                             (Looking up)
Don't stand over me lak dat, ugly as yo' is.

Man (Skanko)
You ain't nobody's pretty baby yo'self!

Cliff
Dat's all right, I ain't as ugly as yo'--youse ugly enough to git behind
a Simpoon weed and hatch monkies.

Man (Skanko)
And youse ugly enough to git behind a tombstone and hatch hants.

Cliff
Youse so ugly dey have to cover yo' face up at night so sleep can slip
up on yo'.

Man (Skanko)
You look like ten cents worth of Have-Mercy. Yo' face look lak ole Uncle
Jump-off. Yo' mouth look lak a bunch of ruffles.

Cliff
Yeah, but yo' done passed me. Yo' so ugly till they could throw yo' in
de Mississippi River and skim ugly for six months.

Man (Skanko)
Look here, Cliff, don't yo' personate me! Counting from de little finger
back to de thumb--yo' start anythin', I go yo' some.

Cliff
Go head and grab me buddie, but if yo' don't know how to turn me loose
too, don't bring de mess up! If yo' hit me, I may not beat you, but
yo'll be so dirty when St. Peter git yo' dat he can't use yo'.

Man (Skanko)
Don't call _me_ buddy. Yo' buddy is huntin' coconuts. Don't yo' try to
throw me for a nap. Do. I'll kill yo' so stiff dead they'll have to push
yo' down. Yo' gointer to make me do some double cussin' on you.
                             (He picks up a heavy stick and walks
                             back towards Cliff)
Now I got dis farmer's choice in my hands, yo' better git outa my face.

Cliff
Yo' wanta fight?

Man
Yeah I wanta fight. Put it where I kin use it and I'll sho' use it. I'll
fight anybody. I get so hot sometimes I fights de corner of de house.
I'm so hot I totes a pistol to keep from gettin' in a fight wid myself.
I prints dangerous every time I sit down in, in a chair.

Cliff
Man, this ain't no fighting weather. Ha, ha, ha! Did yo' think I was mad
sho' nuff? Yo' can't fight me. They's got to be runnin' before fightin'
and they's got to be plenty _good_ runnin' before dis fight comes off.

Man
All right now. Yo' leave me alone and I'm a _good_ man. I'm just like an
old shoe. If yo' rain on me and cool me off I'm soft! If yo' shine on me
and git me hot, I'm hard.
                             (He drops the stick and exits)

                             (Cliff is shaking all over. He looks
                             after the Man to be sure he is gone)

Good Black
Kah, kah, kah. Whut yo' so scarred about? De way yo' was talkin' I
though yo' was mad enough to fight.

Cliff
I was. I gits hot real quick! But I'm very easy cooled when de man I'm
mad wid is bigger'n me.
                             (He drops into his seat, wiping his
                             face)
Man did yo' see how he grabbed up dat check? He done skeered me into a
three-week's spasm!

Good Black's Wife
Good Black, dese clothes is still waiting.

Good Black
Well, let 'em wait on, I done tole yo' once. Yo' kind run yo' mouf but
yo' can't run my business.

                             (Enter a PRETTY GIRL. She strolls
                             happily across without stopping. Good
                             Black pretends to cough)

Good Black
Who is dat?

Girl
                             (Turns and glares at him)
My old man got something for dat cough yo' got.

Cliff
Dat's right, tell dese old mullet hear married men to mind they own
business. Now, take _me_ for instance. I'm a much-right man.
                             (Gets up and approaches her
                             flirtatiously)
I didn't quite git yo' name straight. Yo' better tell it to me again.

Girl
My name is Bee Ethel, turned round to Jones.

Cliff
                             (Flirtatiously)
Yo' pretty lil ole ground angel yo'? Where did yo' come from?

Bee Ethel
Detroit. Yo' like me?

Cliff
Do I lak yo'? I love yo' just lak God loves Gabriel, and dat's his best
angel. Go 'head and say somethin'. I jus' love to hear yo' talk.

Bee Ethel
Gimme five dollars. I need some stockings.

Cliff
_Now_ Mama, dis ain't Gimme, Ga. Dis is Waycross. I'm just lak de
cemetery. I takes in but never no put out. I ain't puttin' out nothin'
but old folks eyes--and I don't do that till they's dead. Run
long, mama.
                             (The girl exits and he resumes his
                             seat)

Cliff
Come on, Good Black, lemme wrap dis checker roun yo' neck.

Good Black
Gimme time, gimme time! Don't try to rush me.
                             (He begins same business of figuring
                             out moves and scratching his head)

                             (Enter two or three girls and fellows.
                             The girls are dressed in cool summer
                             dresses, but nothing elaborate)

Lonnie
I know I'm gointer play something now.
                             (He tunes and plays "Cold Rainy Day".
                             He begins to sing and the others join
                             in. Not all. But all start to dancing.
                             They couple off as far as possible and
                             Lindy. The men unmated do hot solo
                             steps. The men cry out in ecstacy)

1. Shimmy! If you can't shimmy, shake your head.

2. Look, baby, look! Throw it in de alley

3. Look, if you can't look, stick out, and if you can't stick
out, git out.
                             (At the end of the son and dance, one
                             of the girls exclaim)

Girl
Aw, we got to go. Mama's looking for us.
                             (The three girls exit, walking
                             happily. The men watch them go)

Cliff
Oh boy, look at 'em! Switching it and looking back at it.
                             (He imitates the girl's walk)

Good Black
Yeah Lawd, ain't they specifyin'! They handles a lot of traffic.

Cliff
                             (Seating himself again)
Yeah, but dat don't play no checkers. Come on here, Good Black
and lemme finish wearing your ant.

Good Black's Wife
Good Black, yo' better come git dese clothes.

Lonnie
Good Black, yo' wife kin cold whoop for what she want.

Good Black
Yeah and if she don't git, she keep right on whoopin'. B'lieve
I wants a drink of water. Wisht I knowed where I could slip
up on me a drink.

Cliff
Aw man, come on back here and move. Yo' doin' everythin' but playin'
checkers. You'd ruther move a mountain wid a pry bar than to move
                             (Points)
dat man.

Good Black
                             (Seats himself)
Lemme hurry up and beat dis game befo' yo' bust yo' britches.
                             (He wags his finger to indicate moves,
                             scratches his head, but doesn't move.
                             Several men enter and group around the
                             players. All offer suggestions. One
                             says, "you got him Cliffert. He's
                             locked up just as tight as a keyhole".
                             Another: "Aw, man he kin break out!"
                             Another: "Yeah, but it'll cost him
                             plenty to git out of dat trap".)

Cliff
Police! Police! He won't move!

Another Voice
Aw, leave go de checkers and less shoot some crap.

                             (Enter a WOMAN in a house dress, head
                             rag on, run down house shoes. She goes
                             to the edge of the porch and calls
                             inside)

Woman
Him there Bertha, what yo' doin'.

Woman Inside
Still bumpin' de white folks clothes--hittin' for de sundown man. Come
on in and have some sit down.

Outside Woman
Ain't got time. Got a house full of company. I took a minute to see if
yo' could let me have a little skeeting garret.

Inside Woman
How come yo' didn't git yo'self some snuff whilst yo' was at de store?
De man ast yo' what else. I ain't no Piggly Wiggly. Reckon I kin spare
yo' a dip tho.
                             (She hands out the box and the outside
                             woman fills her lip and hands it back)

Outside Woman
Much obliged, I thank yo'. Reckon I better heel and toe it on back, to
see how de comp'ny is makin' out.

Inside Woman
Step inside a minute I want to put a bug in yo' ear.
                             (She makes an urgent gesture and the
                             other woman goes inside)

                             (Lonnie is sitting off to himself and
                             picking "Rabbit on de Log" softly. A
                             small BOY dashes on with a lolly pop
                             in his hand. He is licking it and
                             laughing. He is pursued by a little
                             GIRL yelling "you gimme my all day
                             sucker! Johnny! You gimme my candy,
                             now!" They run all over the stage. The
                             men take notice of them and one of
                             them seizes the boy and restores the
                             candy to the girl. She pokes out her
                             tongue at the boy and says "goody,
                             goody, goody, goody, goody!" She notes
                             the guitar playing and begins to
                             dance. The boy makes faces back at her
                             and dances back at her. The music gets
                             louder, dancing faster, check board
                             gets upset. General laughter at that.
                             When dance is over, boy snatches the
                             lolly pop again and races away and the
                             girl runs behind him yelling "Johnny!
                             You gimme my candy! Johnny!" The music
                             stops and the crap game gets under
                             way. Furious side bets for 5 and 10
                             cents each. Loud calls on Miss "Daisy
                             Dice", snake eyes, "Ada from Decatur".
                             Somebody suggests a soft roll, others
                             object on the ground that it's too
                             easy for the experts to cheat)

Good Black
Gimme de dice! I'm gointer play 'em like John Henry.

Lonnie
John Henry didn't bother wid de bones. He used to play Georgy Skin.

Good Black
He shot crap too. He played everythin' and everythin' he played, he
played it good. Just like he uster drive steel. If I could whip steel
like John Henry, I wouldn't stay here and nowhere else.

Cliff
Whut would yo' do?

Good Black
I'd go somewhere and keep books for somebody.

Lonnie
I know how to play John Henry.

Good Black
Well, turn it on and let de bad luck happen.
                             (As Lonnie plays thru a verse warming
                             up, all the men get interested and
                             start to hum. Cliffert shouts out)

Cliff
Lawd, Lawd, what evil have I done)
                             (They sing John Henry. At the close,
                             the woman who came to borrow snuff
                             emerges from the house still talking
                             back at the woman inside)

Woman
He ain't no trouble. I tole him, I says, "yo' must think youse de man
dat made side meat taste lak ham." See yo' later.
                             (She exits hurriedly. The crap game
                             goes on until a band is heard
                             approaching)

Lonnie
Who dead?

Cliff
Nobody. Don't you know de Imperial Elks is goin' to New York to de Elks
Grand Lodge? Yeah, bo, and they's takin' they band. Dat's supposed to be
de _finest_ band in de United States.
                             (The band approaches followed by a
                             great crow. The crap game is instantly
                             deserted and all follow the band)



***END OF THE PROJECT GUTENBERG EBOOK THREE PLAYS***


******* This file should be named 17187.txt or 17187.zip *******


This and all associated files of various formats will be found in:
http://www.gutenberg.org/dirs/1/7/1/8/17187



Updated editions will replace the previous one--the old editions
will be renamed.

Creating the works from public domain print editions means that no
one owns a United States copyright in these works, so the Foundation
(and you!) can copy and distribute it in the United States without
permission and without paying copyright royalties.  Special rules,
set forth in the General Terms of Use part of this license, apply to
copying and distributing Project Gutenberg-tm electronic works to
protect the PROJECT GUTENBERG-tm concept and trademark.  Project
Gutenberg is a registered trademark, and may not be used if you
charge for the eBooks, unless you receive specific permission.  If you
do not charge anything for copies of this eBook, complying with the
rules is very easy.  You may use this eBook for nearly any purpose
such as creation of derivative works, reports, performances and
research.  They may be modified and printed and given away--you may do
practically ANYTHING with public domain eBooks.  Redistribution is
subject to the trademark license, especially commercial
redistribution.



*** START: FULL LICENSE ***

THE FULL PROJECT GUTENBERG LICENSE
PLEASE READ THIS BEFORE YOU DISTRIBUTE OR USE THIS WORK

To protect the Project Gutenberg-tm mission of promoting the free
distribution of electronic works, by using or distributing this work
(or any other work associated in any way with the phrase "Project
Gutenberg"), you agree to comply with all the terms of the Full Project
Gutenberg-tm License (available with this file or online at
http://gutenberg.net/license).


Section 1.  General Terms of Use and Redistributing Project Gutenberg-tm
electronic works

1.A.  By reading or using any part of this Project Gutenberg-tm
electronic work, you indicate that you have read, understand, agree to
and accept all the terms of this license and intellectual property
(trademark/copyright) agreement.  If you do not agree to abide by all
the terms of this agreement, you must cease using and return or destroy
all copies of Project Gutenberg-tm electronic works in your possession.
If you paid a fee for obtaining a copy of or access to a Project
Gutenberg-tm electronic work and you do not agree to be bound by the
terms of this agreement, you may obtain a refund from the person or
entity to whom you paid the fee as set forth in paragraph 1.E.8.

1.B.  "Project Gutenberg" is a registered trademark.  It may only be
used on or associated in any way with an electronic work by people who
agree to be bound by the terms of this agreement.  There are a few
things that you can do with most Project Gutenberg-tm electronic works
even without complying with the full terms of this agreement.  See
paragraph 1.C below.  There are a lot of things you can do with Project
Gutenberg-tm electronic works if you follow the terms of this agreement
and help preserve free future access to Project Gutenberg-tm electronic
works.  See paragraph 1.E below.

1.C.  The Project Gutenberg Literary Archive Foundation ("the Foundation"
or PGLAF), owns a compilation copyright in the collection of Project
Gutenberg-tm electronic works.  Nearly all the individual works in the
collection are in the public domain in the United States.  If an
individual work is in the public domain in the United States and you are
located in the United States, we do not claim a right to prevent you from
copying, distributing, performing, displaying or creating derivative
works based on the work as long as all references to Project Gutenberg
are removed.  Of course, we hope that you will support the Project
Gutenberg-tm mission of promoting free access to electronic works by
freely sharing Project Gutenberg-tm works in compliance with the terms of
this agreement for keeping the Project Gutenberg-tm name associated with
the work.  You can easily comply with the terms of this agreement by
keeping this work in the same format with its attached full Project
Gutenberg-tm License when you share it without charge with others.

1.D.  The copyright laws of the place where you are located also govern
what you can do with this work.  Copyright laws in most countries are in
a constant state of change.  If you are outside the United States, check
the laws of your country in addition to the terms of this agreement
before downloading, copying, displaying, performing, distributing or
creating derivative works based on this work or any other Project
Gutenberg-tm work.  The Foundation makes no representations concerning
the copyright status of any work in any country outside the United
States.

1.E.  Unless you have removed all references to Project Gutenberg:

1.E.1.  The following sentence, with active links to, or other immediate
access to, the full Project Gutenberg-tm License must appear prominently
whenever any copy of a Project Gutenberg-tm work (any work on which the
phrase "Project Gutenberg" appears, or with which the phrase "Project
Gutenberg" is associated) is accessed, displayed, performed, viewed,
copied or distributed:

This eBook is for the use of anyone anywhere at no cost and with
almost no restrictions whatsoever.  You may copy it, give it away or
re-use it under the terms of the Project Gutenberg License included
with this eBook or online at www.gutenberg.net

1.E.2.  If an individual Project Gutenberg-tm electronic work is derived
from the public domain (does not contain a notice indicating that it is
posted with permission of the copyright holder), the work can be copied
and distributed to anyone in the United States without paying any fees
or charges.  If you are redistributing or providing access to a work
with the phrase "Project Gutenberg" associated with or appearing on the
work, you must comply either with the requirements of paragraphs 1.E.1
through 1.E.7 or obtain permission for the use of the work and the
Project Gutenberg-tm trademark as set forth in paragraphs 1.E.8 or
1.E.9.

1.E.3.  If an individual Project Gutenberg-tm electronic work is posted
with the permission of the copyright holder, your use and distribution
must comply with both paragraphs 1.E.1 through 1.E.7 and any additional
terms imposed by the copyright holder.  Additional terms will be linked
to the Project Gutenberg-tm License for all works posted with the
permission of the copyright holder found at the beginning of this work.

1.E.4.  Do not unlink or detach or remove the full Project Gutenberg-tm
License terms from this work, or any files containing a part of this
work or any other work associated with Project Gutenberg-tm.

1.E.5.  Do not copy, display, perform, distribute or redistribute this
electronic work, or any part of this electronic work, without
prominently displaying the sentence set forth in paragraph 1.E.1 with
active links or immediate access to the full terms of the Project
Gutenberg-tm License.

1.E.6.  You may convert to and distribute this work in any binary,
compressed, marked up, nonproprietary or proprietary form, including any
word processing or hypertext form.  However, if you provide access to or
distribute copies of a Project Gutenberg-tm work in a format other than
"Plain Vanilla ASCII" or other format used in the official version
posted on the official Project Gutenberg-tm web site (www.gutenberg.net),
you must, at no additional cost, fee or expense to the user, provide a
copy, a means of exporting a copy, or a means of obtaining a copy upon
request, of the work in its original "Plain Vanilla ASCII" or other
form.  Any alternate format must include the full Project Gutenberg-tm
License as specified in paragraph 1.E.1.

1.E.7.  Do not charge a fee for access to, viewing, displaying,
performing, copying or distributing any Project Gutenberg-tm works
unless you comply with paragraph 1.E.8 or 1.E.9.

1.E.8.  You may charge a reasonable fee for copies of or providing
access to or distributing Project Gutenberg-tm electronic works provided
that

- You pay a royalty fee of 20% of the gross profits you derive from
     the use of Project Gutenberg-tm works calculated using the method
     you already use to calculate your applicable taxes.  The fee is
     owed to the owner of the Project Gutenberg-tm trademark, but he
     has agreed to donate royalties under this paragraph to the
     Project Gutenberg Literary Archive Foundation.  Royalty payments
     must be paid within 60 days following each date on which you
     prepare (or are legally required to prepare) your periodic tax
     returns.  Royalty payments should be clearly marked as such and
     sent to the Project Gutenberg Literary Archive Foundation at the
     address specified in Section 4, "Information about donations to
     the Project Gutenberg Literary Archive Foundation."

- You provide a full refund of any money paid by a user who notifies
     you in writing (or by e-mail) within 30 days of receipt that s/he
     does not agree to the terms of the full Project Gutenberg-tm
     License.  You must require such a user to return or
     destroy all copies of the works possessed in a physical medium
     and discontinue all use of and all access to other copies of
     Project Gutenberg-tm works.

- You provide, in accordance with paragraph 1.F.3, a full refund of any
     money paid for a work or a replacement copy, if a defect in the
     electronic work is discovered and reported to you within 90 days
     of receipt of the work.

- You comply with all other terms of this agreement for free
     distribution of Project Gutenberg-tm works.

1.E.9.  If you wish to charge a fee or distribute a Project Gutenberg-tm
electronic work or group of works on different terms than are set
forth in this agreement, you must obtain permission in writing from
both the Project Gutenberg Literary Archive Foundation and Michael
Hart, the owner of the Project Gutenberg-tm trademark.  Contact the
Foundation as set forth in Section 3 below.

1.F.

1.F.1.  Project Gutenberg volunteers and employees expend considerable
effort to identify, do copyright research on, transcribe and proofread
public domain works in creating the Project Gutenberg-tm
collection.  Despite these efforts, Project Gutenberg-tm electronic
works, and the medium on which they may be stored, may contain
"Defects," such as, but not limited to, incomplete, inaccurate or
corrupt data, transcription errors, a copyright or other intellectual
property infringement, a defective or damaged disk or other medium, a
computer virus, or computer codes that damage or cannot be read by
your equipment.

1.F.2.  LIMITED WARRANTY, DISCLAIMER OF DAMAGES - Except for the "Right
of Replacement or Refund" described in paragraph 1.F.3, the Project
Gutenberg Literary Archive Foundation, the owner of the Project
Gutenberg-tm trademark, and any other party distributing a Project
Gutenberg-tm electronic work under this agreement, disclaim all
liability to you for damages, costs and expenses, including legal
fees.  YOU AGREE THAT YOU HAVE NO REMEDIES FOR NEGLIGENCE, STRICT
LIABILITY, BREACH OF WARRANTY OR BREACH OF CONTRACT EXCEPT THOSE
PROVIDED IN PARAGRAPH F3.  YOU AGREE THAT THE FOUNDATION, THE
TRADEMARK OWNER, AND ANY DISTRIBUTOR UNDER THIS AGREEMENT WILL NOT BE
LIABLE TO YOU FOR ACTUAL, DIRECT, INDIRECT, CONSEQUENTIAL, PUNITIVE OR
INCIDENTAL DAMAGES EVEN IF YOU GIVE NOTICE OF THE POSSIBILITY OF SUCH
DAMAGE.

1.F.3.  LIMITED RIGHT OF REPLACEMENT OR REFUND - If you discover a
defect in this electronic work within 90 days of receiving it, you can
receive a refund of the money (if any) you paid for it by sending a
written explanation to the person you received the work from.  If you
received the work on a physical medium, you must return the medium with
your written explanation.  The person or entity that provided you with
the defective work may elect to provide a replacement copy in lieu of a
refund.  If you received the work electronically, the person or entity
providing it to you may choose to give you a second opportunity to
receive the work electronically in lieu of a refund.  If the second copy
is also defective, you may demand a refund in writing without further
opportunities to fix the problem.

1.F.4.  Except for the limited right of replacement or refund set forth
in paragraph 1.F.3, this work is provided to you 'AS-IS', WITH NO OTHER
WARRANTIES OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
WARRANTIES OF MERCHANTIBILITY OR FITNESS FOR ANY PURPOSE.

1.F.5.  Some states do not allow disclaimers of certain implied
warranties or the exclusion or limitation of certain types of damages.
If any disclaimer or limitation set forth in this agreement violates the
law of the state applicable to this agreement, the agreement shall be
interpreted to make the maximum disclaimer or limitation permitted by
the applicable state law.  The invalidity or unenforceability of any
provision of this agreement shall not void the remaining provisions.

1.F.6.  INDEMNITY - You agree to indemnify and hold the Foundation, the
trademark owner, any agent or employee of the Foundation, anyone
providing copies of Project Gutenberg-tm electronic works in accordance
with this agreement, and any volunteers associated with the production,
promotion and distribution of Project Gutenberg-tm electronic works,
harmless from all liability, costs and expenses, including legal fees,
that arise directly or indirectly from any of the following which you do
or cause to occur: (a) distribution of this or any Project Gutenberg-tm
work, (b) alteration, modification, or additions or deletions to any
Project Gutenberg-tm work, and (c) any Defect you cause.


Section  2.  Information about the Mission of Project Gutenberg-tm

Project Gutenberg-tm is synonymous with the free distribution of
electronic works in formats readable by the widest variety of computers
including obsolete, old, middle-aged and new computers.  It exists
because of the efforts of hundreds of volunteers and donations from
people in all walks of life.

Volunteers and financial support to provide volunteers with the
assistance they need, is critical to reaching Project Gutenberg-tm's
goals and ensuring that the Project Gutenberg-tm collection will
remain freely available for generations to come.  In 2001, the Project
Gutenberg Literary Archive Foundation was created to provide a secure
and permanent future for Project Gutenberg-tm and future generations.
To learn more about the Project Gutenberg Literary Archive Foundation
and how your efforts and donations can help, see Sections 3 and 4
and the Foundation web page at http://www.gutenberg.net/fundraising/pglaf.


Section 3.  Information about the Project Gutenberg Literary Archive
Foundation

The Project Gutenberg Literary Archive Foundation is a non profit
501(c)(3) educational corporation organized under the laws of the
state of Mississippi and granted tax exempt status by the Internal
Revenue Service.  The Foundation's EIN or federal tax identification
number is 64-6221541.  Contributions to the Project Gutenberg
Literary Archive Foundation are tax deductible to the full extent
permitted by U.S. federal laws and your state's laws.

The Foundation's principal office is located at 4557 Melan Dr. S.
Fairbanks, AK, 99712., but its volunteers and employees are scattered
throughout numerous locations.  Its business office is located at
809 North 1500 West, Salt Lake City, UT 84116, (801) 596-1887, email
business@pglaf.org.  Email contact links and up to date contact
information can be found at the Foundation's web site and official
page at http://www.gutenberg.net/about/contact

For additional contact information:
     Dr. Gregory B. Newby
     Chief Executive and Director
     gbnewby@pglaf.org

Section 4.  Information about Donations to the Project Gutenberg
Literary Archive Foundation

Project Gutenberg-tm depends upon and cannot survive without wide
spread public support and donations to carry out its mission of
increasing the number of public domain and licensed works that can be
freely distributed in machine readable form accessible by the widest
array of equipment including outdated equipment.  Many small donations
($1 to $5,000) are particularly important to maintaining tax exempt
status with the IRS.

The Foundation is committed to complying with the laws regulating
charities and charitable donations in all 50 states of the United
States.  Compliance requirements are not uniform and it takes a
considerable effort, much paperwork and many fees to meet and keep up
with these requirements.  We do not solicit donations in locations
where we have not received written confirmation of compliance.  To
SEND DONATIONS or determine the status of compliance for any
particular state visit http://www.gutenberg.net/fundraising/donate

While we cannot and do not solicit contributions from states where we
have not met the solicitation requirements, we know of no prohibition
against accepting unsolicited donations from donors in such states who
approach us with offers to donate.

International donations are gratefully accepted, but we cannot make
any statements concerning tax treatment of donations received from
outside the United States.  U.S. laws alone swamp our small staff.

Please check the Project Gutenberg Web pages for current donation
methods and addresses.  Donations are accepted in a number of other
ways including including checks, online payments and credit card
donations.  To donate, please visit:
http://www.gutenberg.net/fundraising/donate


Section 5.  General Information About Project Gutenberg-tm electronic
works.

Professor Michael S. Hart is the originator of the Project Gutenberg-tm
concept of a library of electronic works that could be freely shared
with anyone.  For thirty years, he produced and distributed Project
Gutenberg-tm eBooks with only a loose network of volunteer support.

Project Gutenberg-tm eBooks are often created from several printed
editions, all of which are confirmed as Public Domain in the U.S.
unless a copyright notice is included.  Thus, we do not necessarily
keep eBooks in compliance with any particular paper edition.

Most people start at our Web site which has the main PG search facility:

     http://www.gutenberg.net

This Web site includes information about Project Gutenberg-tm,
including how to make donations to the Project Gutenberg Literary
Archive Foundation, how to help produce our new eBooks, and how to
subscribe to our email newsletter to hear about new eBooks.
`

	b["loves-labors-lost_TXT_FolgerShakespeare.txt"] = `
Love's Labor's Lost
by William Shakespeare
Edited by Barbara A. Mowat and Paul Werstine
  with Michael Poston and Rebecca Niles
Folger Shakespeare Library
https://shakespeare.folger.edu/shakespeares-works/loves-labors-lost/
Created on Jul 31, 2015, from FDT version 0.9.2

Characters in the Play
======================
KING of Navarre, also known as Ferdinand
Lords attending the King:
  BEROWNE
  LONGAVILLE
  DUMAINE
The PRINCESS of France
Ladies attending the Princess:
  ROSALINE
  MARIA
  KATHERINE
BOYET, a lord attending the Princess
ARMADO, the BRAGGART, also known as Don Adriano de Armado
BOY, Armado's PAGE, also known as MOTE
JAQUENETTA, the WENCH
COSTARD, the CLOWN or SWAIN
DULL, the CONSTABLE
HOLOFERNES, the PEDANT, or schoolmaster
NATHANIEL, the CURATE
FORESTER
MONSIEUR MARCADE, a messenger from France
Lords, Blackamoors, Musicians


ACT 1
=====

Scene 1
=======
[Enter Ferdinand, King of Navarre, Berowne,
Longaville, and Dumaine.]


KING
Let fame, that all hunt after in their lives,
Live registered upon our brazen tombs,
And then grace us in the disgrace of death,
When, spite of cormorant devouring time,
Th' endeavor of this present breath may buy
That honor which shall bate his scythe's keen edge
And make us heirs of all eternity.
Therefore, brave conquerors, for so you are
That war against your own affections
And the huge army of the world's desires,
Our late edict shall strongly stand in force.
Navarre shall be the wonder of the world;
Our court shall be a little academe,
Still and contemplative in living art.
You three, Berowne, Dumaine, and Longaville,
Have sworn for three years' term to live with me,
My fellow scholars, and to keep those statutes
That are recorded in this schedule here.
[He holds up a scroll.]
Your oaths are passed, and now subscribe your
names,
That his own hand may strike his honor down
That violates the smallest branch herein.
If you are armed to do as sworn to do,
Subscribe to your deep oaths, and keep it too.

LONGAVILLE
I am resolved. 'Tis but a three years' fast.
The mind shall banquet though the body pine.
Fat paunches have lean pates, and dainty bits
Make rich the ribs but bankrout quite the wits.
[He signs his name.]

DUMAINE
My loving lord, Dumaine is mortified.
The grosser manner of these world's delights
He throws upon the gross world's baser slaves.
To love, to wealth, to pomp I pine and die,
With all these living in philosophy.
[He signs his name.]

BEROWNE
I can but say their protestation over.
So much, dear liege, I have already sworn,
That is, to live and study here three years.
But there are other strict observances:
As not to see a woman in that term,
Which I hope well is not enrolled there;
And one day in a week to touch no food,
And but one meal on every day besides,
The which I hope is not enrolled there;
And then to sleep but three hours in the night,
And not be seen to wink of all the day--
When I was wont to think no harm all night,
And make a dark night too of half the day--
Which I hope well is not enrolled there.
O, these are barren tasks, too hard to keep,
Not to see ladies, study, fast, not sleep.

KING
Your oath is passed to pass away from these.

BEROWNE
Let me say no, my liege, an if you please.
I only swore to study with your Grace
And stay here in your court for three years' space.

LONGAVILLE
You swore to that, Berowne, and to the rest.

BEROWNE
By yea and nay, sir. Then I swore in jest.
What is the end of study, let me know?

KING
Why, that to know which else we should not know.

BEROWNE
Things hid and barred, you mean, from common
sense.

KING
Ay, that is study's godlike recompense.

BEROWNE
Come on, then, I will swear to study so,
To know the thing I am forbid to know:
As thus--to study where I well may dine,
   When I to feast expressly am forbid;
Or study where to meet some mistress fine
   When mistresses from common sense are hid;
Or having sworn too hard-a-keeping oath,
Study to break it, and not break my troth.
If study's gain be thus, and this be so,
Study knows that which yet it doth not know.
Swear me to this, and I will ne'er say no.

KING
These be the stops that hinder study quite,
And train our intellects to vain delight.

BEROWNE
Why, all delights are vain, and that most vain
Which with pain purchased doth inherit pain:
As painfully to pore upon a book
   To seek the light of truth, while truth the while
Doth falsely blind the eyesight of his look.
   Light seeking light doth light of light beguile.
So, ere you find where light in darkness lies,
Your light grows dark by losing of your eyes.
Study me how to please the eye indeed
   By fixing it upon a fairer eye,
Who dazzling so, that eye shall be his heed
   And give him light that it was blinded by.
Study is like the heaven's glorious sun,
   That will not be deep-searched with saucy looks.
Small have continual plodders ever won,
   Save base authority from others' books.
These earthly godfathers of heaven's lights,
   That give a name to every fixed star,
Have no more profit of their shining nights
   Than those that walk and wot not what they are.
Too much to know is to know naught but fame,
And every godfather can give a name.

KING
How well he's read to reason against reading.

DUMAINE
Proceeded well, to stop all good proceeding.

LONGAVILLE
He weeds the corn, and still lets grow the weeding.

BEROWNE
The spring is near when green geese are a-breeding.

DUMAINE
How follows that?

BEROWNE  Fit in his place and time.

DUMAINE
In reason nothing.

BEROWNE  Something then in rhyme.

KING
Berowne is like an envious sneaping frost
   That bites the firstborn infants of the spring.

BEROWNE
Well, say I am. Why should proud summer boast
   Before the birds have any cause to sing?
Why should I joy in any abortive birth?
At Christmas I no more desire a rose
Than wish a snow in May's new-fangled shows,
But like of each thing that in season grows.
So you, to study now it is too late,
Climb o'er the house to unlock the little gate.

KING
Well, sit you out. Go home, Berowne. Adieu.

BEROWNE
No, my good lord, I have sworn to stay with you.
And though I have for barbarism spoke more
   Than for that angel knowledge you can say,
Yet, confident, I'll keep what I have sworn
   And bide the penance of each three years' day.
Give me the paper. Let me read the same,
And to the strictest decrees I'll write my name.

KING
How well this yielding rescues thee from shame.

BEROWNE [reads]  Item, That no woman shall come within
a mile of my court. Hath this been proclaimed?

LONGAVILLE  Four days ago.

BEROWNE  Let's see the penalty. [Reads:] On pain of
losing her tongue. Who devised this penalty?

LONGAVILLE  Marry, that did I.

BEROWNE  Sweet lord, and why?

LONGAVILLE
To fright them hence with that dread penalty.

BEROWNE
A dangerous law against gentility.
[Reads:] Item, If any man be seen to talk with a
woman within the term of three years, he shall endure
such public shame as the rest of the court can possible
devise.
This article, my liege, yourself must break,
   For well you know here comes in embassy
The French king's daughter with yourself to speak--
   A maid of grace and complete majesty--
About surrender up of Aquitaine
   To her decrepit, sick, and bedrid father.
Therefore this article is made in vain,
   Or vainly comes th' admired princess hither.

KING
What say you, lords? Why, this was quite forgot.

BEROWNE
So study evermore is overshot.
While it doth study to have what it would,
It doth forget to do the thing it should.
And when it hath the thing it hunteth most,
'Tis won as towns with fire--so won, so lost.

KING
We must of force dispense with this decree.
She must lie here on mere necessity.

BEROWNE
Necessity will make us all forsworn
   Three thousand times within this three years'
   space;
For every man with his affects is born,
   Not by might mastered, but by special grace.
If I break faith, this word shall speak for me:
I am forsworn on mere necessity.
So to the laws at large I write my name,
   And he that breaks them in the least degree
Stands in attainder of eternal shame.
   Suggestions are to other as to me,
But I believe, although I seem so loath,
I am the last that will last keep his oath.
[He signs his name.]
But is there no quick recreation granted?

KING
Ay, that there is. Our court, you know, is haunted
   With a refined traveler of Spain,
A man in all the world's new fashion planted,
   That hath a mint of phrases in his brain;
One who the music of his own vain tongue
   Doth ravish like enchanting harmony,
A man of compliments, whom right and wrong
   Have chose as umpire of their mutiny.
This child of fancy, that Armado hight,
   For interim to our studies shall relate
In high-born words the worth of many a knight
   From tawny Spain lost in the world's debate.
How you delight, my lords, I know not, I,
But I protest I love to hear him lie,
And I will use him for my minstrelsy.

BEROWNE
Armado is a most illustrious wight,
A man of fire-new words, fashion's own knight.

LONGAVILLE
Costard the swain and he shall be our sport,
And so to study three years is but short.

[Enter Dull, a Constable, with a letter, and Costard.]


DULL  Which is the Duke's own person?

BEROWNE  This, fellow. What wouldst?

DULL  I myself reprehend his own person, for I am his
Grace's farborough. But I would see his own
person in flesh and blood.

BEROWNE  This is he.

DULL, [to King]  Signior Arm-, Arm-, commends you.
There's villainy abroad. This letter will tell you
more.	[He gives the letter to the King.]

COSTARD  Sir, the contempts thereof are as touching
me.

KING  A letter from the magnificent Armado.

BEROWNE  How low soever the matter, I hope in God
for high words.

LONGAVILLE  A high hope for a low heaven. God grant
us patience!

BEROWNE  To hear, or forbear hearing?

LONGAVILLE  To hear meekly, sir, and to laugh moderately,
or to forbear both.

BEROWNE  Well, sir, be it as the style shall give us cause
to climb in the merriness.

COSTARD  The matter is to me, sir, as concerning
Jaquenetta. The manner of it is, I was taken with
the manner.

BEROWNE  In what manner?

COSTARD  In manner and form following, sir, all those
three. I was seen with her in the manor house,
sitting with her upon the form, and taken following
her into the park, which, put together, is "in manner
and form following." Now, sir, for the manner.
It is the manner of a man to speak to a woman. For
the form--in some form.

BEROWNE  For the "following," sir?

COSTARD  As it shall follow in my correction, and God
defend the right.

KING  Will you hear this letter with attention?

BEROWNE  As we would hear an oracle.

COSTARD  Such is the sinplicity of man to hearken after
the flesh.

KING [reads]  Great deputy, the welkin's vicegerent and
sole dominator of Navarre, my soul's earth's god, and
body's fost'ring patron--

COSTARD  Not a word of Costard yet.

KING [reads]  So it is--

COSTARD  It may be so, but if he say it is so, he is, in
telling true, but so.

KING  Peace.

COSTARD  Be to me, and every man that dares not fight.

KING  No words.

COSTARD  Of other men's secrets, I beseech you.

KING [reads]  So it is, besieged with sable-colored melancholy,
I did commend the black oppressing humor
to the most wholesome physic of thy health-giving air;
and, as I am a gentleman, betook myself to walk. The
time when? About the sixth hour, when beasts most
graze, birds best peck, and men sit down to that
nourishment which is called supper. So much for the
time when. Now for the ground which--which, I
mean, I walked upon. It is yclept thy park. Then for the
place where--where, I mean, I did encounter that
obscene and most prepost'rous event that draweth
from my snow-white pen the ebon-colored ink, which
here thou viewest, beholdest, surveyest, or seest. But to
the place where. It standeth north-north-east and by
east from the west corner of thy curious-knotted
garden. There did I see that low-spirited swain, that
base minnow of thy mirth,--

COSTARD  Me?

KING [reads]  that unlettered, small-knowing soul,--

COSTARD  Me?

KING [reads]  that shallow vassal,--

COSTARD  Still me?

KING [reads]  which, as I remember, hight Costard,--

COSTARD  O, me!

KING [reads]  sorted and consorted, contrary to thy
established proclaimed edict and continent canon,
which with--O with--but with this I passion to say
wherewith--

COSTARD  With a wench.

KING [reads]  with a child of our grandmother Eve, a
female; or, for thy more sweet understanding, a
woman: him, I, as my ever-esteemed duty pricks
me on, have sent to thee, to receive the meed of
punishment by thy sweet Grace's officer, Anthony
Dull, a man of good repute, carriage, bearing, and
estimation.

DULL  Me, an 't shall please you. I am Anthony Dull.

KING [reads]  For Jaquenetta--so is the weaker vessel
called which I apprehended with the aforesaid
swain--I keep her as a vessel of thy law's fury, and
shall, at the least of thy sweet notice, bring her to trial.
Thine, in all compliments of devoted and heartburning
heat of duty,
Don Adriano de Armado.

BEROWNE  This is not so well as I looked for, but the
best that ever I heard.

KING  Ay, the best, for the worst. [To Costard.] But,
sirrah, what say you to this?

COSTARD  Sir, I confess the wench.

KING  Did you hear the proclamation?

COSTARD  I do confess much of the hearing it, but little
of the marking of it.

KING  It was proclaimed a year's imprisonment to be
taken with a wench.

COSTARD  I was taken with none, sir. I was taken with a
damsel.

KING  Well, it was proclaimed "damsel."

COSTARD  This was no damsel neither, sir. She was a
virgin.

BEROWNE  It is so varied too, for it was proclaimed
"virgin."

COSTARD  If it were, I deny her virginity. I was taken
with a maid.

KING  This "maid" will not serve your turn, sir.

COSTARD  This maid will serve my turn, sir.

KING  Sir, I will pronounce your sentence: you shall
fast a week with bran and water.

COSTARD  I had rather pray a month with mutton and
porridge.

KING  And Don Armado shall be your keeper.
My Lord Berowne, see him delivered o'er,
And go we, lords, to put in practice that
   Which each to other hath so strongly sworn.
[King, Longaville, and Dumaine exit.]

BEROWNE
I'll lay my head to any goodman's hat,
   These oaths and laws will prove an idle scorn.
Sirrah, come on.

COSTARD  I suffer for the truth, sir; for true it is I was
taken with Jaquenetta, and Jaquenetta is a true
girl. And therefore welcome the sour cup of prosperity.
Affliction may one day smile again, and till
then, sit thee down, sorrow.
[They exit.]

Scene 2
=======
[Enter Armado and Mote, his page.]


ARMADO  Boy, what sign is it when a man of great spirit
grows melancholy?

BOY  A great sign, sir, that he will look sad.

ARMADO  Why, sadness is one and the selfsame thing,
dear imp.

BOY  No, no. O Lord, sir, no!

ARMADO  How canst thou part sadness and melancholy,
my tender juvenal?

BOY  By a familiar demonstration of the working, my
tough signior.

ARMADO  Why "tough signior"? Why "tough signior"?

BOY  Why "tender juvenal"? Why "tender juvenal"?

ARMADO  I spoke it "tender juvenal" as a congruent
epitheton appertaining to thy young days, which
we may nominate "tender."

BOY  And I "tough signior" as an appurtenant title to
your old time, which we may name "tough."

ARMADO  Pretty and apt.

BOY  How mean you, sir? I pretty and my saying apt, or
I apt and my saying pretty?

ARMADO  Thou pretty because little.

BOY  Little pretty, because little. Wherefore apt?

ARMADO  And therefore apt, because quick.

BOY  Speak you this in my praise, master?

ARMADO  In thy condign praise.

BOY  I will praise an eel with the same praise.

ARMADO  What, that an eel is ingenious?

BOY  That an eel is quick.

ARMADO  I do say thou art quick in answers. Thou
heat'st my blood.

BOY  I am answered, sir.

ARMADO  I love not to be crossed.

BOY, [aside]  He speaks the mere contrary; crosses love
not him.

ARMADO  I have promised to study three years with the
Duke.

BOY  You may do it in an hour, sir.

ARMADO  Impossible.

BOY  How many is one thrice told?

ARMADO  I am ill at reckoning. It fitteth the spirit of a
tapster.

BOY  You are a gentleman and a gamester, sir.

ARMADO  I confess both. They are both the varnish of a
complete man.

BOY  Then I am sure you know how much the gross
sum of deuce-ace amounts to.

ARMADO  It doth amount to one more than two.

BOY  Which the base vulgar do call "three."

ARMADO  True.

BOY  Why, sir, is this such a piece of study? Now here is
"three" studied ere you'll thrice wink. And how
easy it is to put "years" to the word "three" and
study "three years" in two words, the dancing horse
will tell you.

ARMADO  A most fine figure.

BOY, [aside]  To prove you a cipher.

ARMADO  I will hereupon confess I am in love; and as it
is base for a soldier to love, so am I in love with a
base wench. If drawing my sword against the
humor of affection would deliver me from the
reprobate thought of it, I would take desire prisoner
and ransom him to any French courtier for a
new-devised curtsy. I think scorn to sigh; methinks
I should outswear Cupid. Comfort me, boy. What
great men have been in love?

BOY  Hercules, master.

ARMADO  Most sweet Hercules! More authority, dear
boy, name more; and, sweet my child, let them be
men of good repute and carriage.

BOY  Samson, master; he was a man of good carriage,
great carriage, for he carried the town gates on his
back like a porter, and he was in love.

ARMADO  O, well-knit Samson, strong-jointed Samson;
I do excel thee in my rapier as much as thou didst
me in carrying gates. I am in love too. Who was
Samson's love, my dear Mote?

BOY  A woman, master.

ARMADO  Of what complexion?

BOY  Of all the four, or the three, or the two, or one of
the four.

ARMADO  Tell me precisely of what complexion.

BOY  Of the sea-water green, sir.

ARMADO  Is that one of the four complexions?

BOY  As I have read, sir, and the best of them too.

ARMADO  Green indeed is the color of lovers. But to
have a love of that color, methinks Samson had
small reason for it. He surely affected her for her
wit.

BOY  It was so, sir, for she had a green wit.

ARMADO  My love is most immaculate white and red.

BOY  Most maculate thoughts, master, are masked
under such colors.

ARMADO  Define, define, well-educated infant.

BOY  My father's wit and my mother's tongue, assist
me.

ARMADO  Sweet invocation of a child, most pretty and
pathetical.

BOY
	If she be made of white and red,
	   Her faults will ne'er be known,
	For blushing cheeks by faults are bred,
	   And fears by pale white shown.
	Then if she fear, or be to blame,
	   By this you shall not know,
	For still her cheeks possess the same
	   Which native she doth owe.
A dangerous rhyme, master, against the reason of
white and red.

ARMADO  Is there not a ballad, boy, of "The King and
the Beggar"?

BOY  The world was very guilty of such a ballad some
three ages since, but I think now 'tis not to be found;
or if it were, it would neither serve for the writing
nor the tune.

ARMADO  I will have that subject newly writ o'er, that I
may example my digression by some mighty precedent.
Boy, I do love that country girl that I took in
the park with the rational hind Costard. She deserves
well.

BOY, [aside]  To be whipped--and yet a better love than
my master.

ARMADO  Sing, boy. My spirit grows heavy in love.

BOY, [aside]  And that's great marvel, loving a light
wench.

ARMADO  I say sing.

BOY  Forbear till this company be past.

[Enter Clown (Costard,) Constable (Dull,) and Wench
(Jaquenetta.)]


DULL, [to Armado]  Sir, the Duke's pleasure is that you
keep Costard safe, and you must suffer him to take
no delight, nor no penance, but he must fast three
days a week. For this damsel, I must keep her at the
park. She is allowed for the dey-woman. Fare you
well.

ARMADO, [aside]  I do betray myself with blushing.--
Maid.

JAQUENETTA  Man.

ARMADO  I will visit thee at the lodge.

JAQUENETTA  That's hereby.

ARMADO  I know where it is situate.

JAQUENETTA  Lord, how wise you are.

ARMADO  I will tell thee wonders.

JAQUENETTA  With that face?

ARMADO  I love thee.

JAQUENETTA  So I heard you say.

ARMADO  And so, farewell.

JAQUENETTA  Fair weather after you.

DULL  Come, Jaquenetta, away.
[Dull and Jaquenetta exit.]

ARMADO, [to Costard]  Villain, thou shalt fast for thy
offenses ere thou be pardoned.

COSTARD  Well, sir, I hope when I do it I shall do it on
a full stomach.

ARMADO  Thou shalt be heavily punished.

COSTARD  I am more bound to you than your fellows,
for they are but lightly rewarded.

ARMADO, [to Boy]  Take away this villain. Shut him up.

BOY  Come, you transgressing slave, away.

COSTARD, [to Armado]  Let me not be pent up, sir. I will
fast being loose.

BOY  No, sir, that were fast and loose. Thou shalt to
prison.

COSTARD  Well, if ever I do see the merry days of
desolation that I have seen, some shall see.

BOY  What shall some see?

COSTARD  Nay, nothing, Master Mote, but what they
look upon. It is not for prisoners to be too silent in
their words, and therefore I will say nothing. I thank
God I have as little patience as another man, and
therefore I can be quiet.
[Costard and Boy exit.]

ARMADO  I do affect the very ground (which is base)
where her shoe (which is baser) guided by her foot
(which is basest) doth tread. I shall be forsworn
(which is a great argument of falsehood) if I love.
And how can that be true love which is falsely
attempted? Love is a familiar; love is a devil. There is
no evil angel but love, yet was Samson so tempted,
and he had an excellent strength; yet was Solomon
so seduced, and he had a very good wit. Cupid's
butt-shaft is too hard for Hercules' club, and therefore
too much odds for a Spaniard's rapier. The first
and second cause will not serve my turn; the
passado he respects not, the duello he regards not.
His disgrace is to be called "boy," but his glory is to
subdue men. Adieu, valor; rust, rapier; be still,
drum, for your manager is in love. Yea, he loveth.
Assist me, some extemporal god of rhyme, for I am
sure I shall turn sonnet. Devise wit, write pen, for I
am for whole volumes in folio.
[He exits.]


ACT 2
=====

Scene 1
=======
[Enter the Princess of France, with three attending
Ladies (Rosaline, Maria, and Katherine), Boyet
and other Lords.]


BOYET
Now, madam, summon up your dearest spirits.
Consider who the King your father sends,
To whom he sends, and what's his embassy.
Yourself, held precious in the world's esteem,
To parley with the sole inheritor
Of all perfections that a man may owe,
Matchless Navarre; the plea of no less weight
Than Aquitaine, a dowry for a queen.
Be now as prodigal of all dear grace
As nature was in making graces dear
When she did starve the general world besides
And prodigally gave them all to you.

PRINCESS
Good Lord Boyet, my beauty, though but mean,
Needs not the painted flourish of your praise.
Beauty is bought by judgment of the eye,
Not uttered by base sale of chapmen's tongues.
I am less proud to hear you tell my worth
Than you much willing to be counted wise
In spending your wit in the praise of mine.
But now to task the tasker: good Boyet,
You are not ignorant all-telling fame
Doth noise abroad Navarre hath made a vow,
Till painful study shall outwear three years,
No woman may approach his silent court.
Therefore to 's seemeth it a needful course,
Before we enter his forbidden gates,
To know his pleasure, and in that behalf,
Bold of your worthiness, we single you
As our best-moving fair solicitor.
Tell him the daughter of the King of France
On serious business craving quick dispatch,
Importunes personal conference with his Grace.
Haste, signify so much, while we attend,
Like humble-visaged suitors, his high will.

BOYET
Proud of employment, willingly I go.

PRINCESS
All pride is willing pride, and yours is so.
[Boyet exits.]
Who are the votaries, my loving lords,
That are vow-fellows with this virtuous duke?

A LORD
Lord Longaville is one.

PRINCESS  Know you the man?

MARIA
I know him, madam. At a marriage feast
Between Lord Perigort and the beauteous heir
Of Jaques Falconbridge, solemnized
In Normandy, saw I this Longaville.
A man of sovereign parts he is esteemed,
Well fitted in arts, glorious in arms.
Nothing becomes him ill that he would well.
The only soil of his fair virtue's gloss,
If virtue's gloss will stain with any soil,
Is a sharp wit matched with too blunt a will,
Whose edge hath power to cut, whose will still wills
It should none spare that come within his power.

PRINCESS
Some merry mocking lord, belike. Is 't so?

MARIA
They say so most that most his humors know.

PRINCESS
Such short-lived wits do wither as they grow.
Who are the rest?

KATHERINE
The young Dumaine, a well-accomplished youth,
Of all that virtue love for virtue loved.
Most power to do most harm, least knowing ill;
For he hath wit to make an ill shape good,
And shape to win grace though he had no wit.
I saw him at the Duke Alanson's once,
And much too little of that good I saw
Is my report to his great worthiness.

ROSALINE
Another of these students at that time
Was there with him, if I have heard a truth.
Berowne they call him, but a merrier man,
Within the limit of becoming mirth,
I never spent an hour's talk withal.
His eye begets occasion for his wit,
For every object that the one doth catch
The other turns to a mirth-moving jest,
Which his fair tongue, conceit's expositor,
Delivers in such apt and gracious words
That aged ears play truant at his tales,
And younger hearings are quite ravished,
So sweet and voluble is his discourse.

PRINCESS
God bless my ladies, are they all in love,
That every one her own hath garnished
With such bedecking ornaments of praise?

A LORD
Here comes Boyet.

[Enter Boyet.]


PRINCESS  Now, what admittance, lord?

BOYET
Navarre had notice of your fair approach,
And he and his competitors in oath
Were all addressed to meet you, gentle lady,
Before I came. Marry, thus much I have learned:
He rather means to lodge you in the field,
Like one that comes here to besiege his court,
Than seek a dispensation for his oath
To let you enter his unpeopled house.

[Enter King of Navarre, Longaville, Dumaine, and
Berowne.]

Here comes Navarre.

KING  Fair Princess, welcome to the court of Navarre.

PRINCESS  "Fair" I give you back again, and "welcome"
I have not yet. The roof of this court is too
high to be yours, and welcome to the wide fields too
base to be mine.

KING
You shall be welcome, madam, to my court.

PRINCESS
I will be welcome, then. Conduct me thither.

KING
Hear me, dear lady. I have sworn an oath.

PRINCESS
Our Lady help my lord! He'll be forsworn.

KING
Not for the world, fair madam, by my will.

PRINCESS
Why, will shall break it, will and nothing else.

KING
Your Ladyship is ignorant what it is.

PRINCESS
Were my lord so, his ignorance were wise,
Where now his knowledge must prove ignorance.
I hear your Grace hath sworn out housekeeping.
'Tis deadly sin to keep that oath, my lord,
And sin to break it.
But pardon me, I am too sudden bold.
To teach a teacher ill beseemeth me.
Vouchsafe to read the purpose of my coming,
And suddenly resolve me in my suit.
[She gives him a paper.]

KING
Madam, I will, if suddenly I may.

PRINCESS
You will the sooner that I were away,
For you'll prove perjured if you make me stay.
[They walk aside while the King reads the paper.]

BEROWNE, [to Rosaline]
Did not I dance with you in Brabant once?

ROSALINE
Did not I dance with you in Brabant once?

BEROWNE
I know you did.

ROSALINE  How needless was it then
To ask the question.

BEROWNE  You must not be so quick.

ROSALINE
'Tis long of you that spur me with such questions.

BEROWNE
Your wit's too hot, it speeds too fast; 'twill tire.

ROSALINE
Not till it leave the rider in the mire.

BEROWNE
What time o' day?

ROSALINE  The hour that fools should ask.

BEROWNE  Now fair befall your mask.

ROSALINE  Fair fall the face it covers.

BEROWNE  And send you many lovers.

ROSALINE  Amen, so you be none.

BEROWNE  Nay, then, will I be gone.

KING, [coming forward with the Princess]
Madam, your father here doth intimate
The payment of a hundred thousand crowns,
Being but the one half of an entire sum
Disbursed by my father in his wars.
But say that he or we, as neither have,
Received that sum, yet there remains unpaid
A hundred thousand more, in surety of the which
One part of Aquitaine is bound to us,
Although not valued to the money's worth.
If then the King your father will restore
But that one half which is unsatisfied,
We will give up our right in Aquitaine,
And hold fair friendship with his Majesty.
But that, it seems, he little purposeth;
For here he doth demand to have repaid
A hundred thousand crowns, and not demands,
On payment of a hundred thousand crowns,
To have his title live in Aquitaine--
Which we much rather had depart withal,
And have the money by our father lent,
Than Aquitaine, so gelded as it is.
Dear Princess, were not his requests so far
From reason's yielding, your fair self should make
A yielding 'gainst some reason in my breast,
And go well satisfied to France again.

PRINCESS
You do the King my father too much wrong,
And wrong the reputation of your name,
In so unseeming to confess receipt
Of that which hath so faithfully been paid.

KING
I do protest I never heard of it;
And if you prove it, I'll repay it back
Or yield up Aquitaine.

PRINCESS  We arrest your word.--
Boyet, you can produce acquittances
For such a sum from special officers
Of Charles his father.

KING  Satisfy me so.

BOYET
So please your Grace, the packet is not come
Where that and other specialties are bound.
Tomorrow you shall have a sight of them.

KING
It shall suffice me; at which interview
All liberal reason I will yield unto.
Meantime receive such welcome at my hand
As honor (without breach of honor) may
Make tender of to thy true worthiness.
You may not come, fair princess, within my gates,
But here without you shall be so received
As you shall deem yourself lodged in my heart,
Though so denied fair harbor in my house.
Your own good thoughts excuse me, and farewell.
Tomorrow shall we visit you again.

PRINCESS
Sweet health and fair desires consort your Grace.

KING
Thy own wish wish I thee in every place.
[He exits with Dumaine,
Longaville, and Attendants.]

BEROWNE, [to Rosaline]  Lady, I will commend you to
my own heart.

ROSALINE  Pray you, do my commendations. I would
be glad to see it.

BEROWNE  I would you heard it groan.

ROSALINE  Is the fool sick?

BEROWNE  Sick at the heart.

ROSALINE  Alack, let it blood.

BEROWNE  Would that do it good?

ROSALINE  My physic says "ay."

BEROWNE  Will you prick 't with your eye?

ROSALINE  No point, with my knife.

BEROWNE  Now God save thy life.

ROSALINE  And yours from long living.

BEROWNE  I cannot stay thanksgiving.	[He exits.]

[Enter Dumaine.]


DUMAINE, [to Boyet]
Sir, I pray you, a word. What lady is that same?

BOYET
The heir of Alanson, Katherine her name.

DUMAINE
A gallant lady, monsieur. Fare you well.	[He exits.]

[Enter Longaville.]


LONGAVILLE, [to Boyet]
I beseech you, a word. What is she in the white?

BOYET
A woman sometimes, an you saw her in the light.

LONGAVILLE
Perchance light in the light. I desire her name.

BOYET
She hath but one for herself; to desire that were a
shame.

LONGAVILLE  Pray you, sir, whose daughter?

BOYET  Her mother's, I have heard.

LONGAVILLE  God's blessing on your beard!

BOYET  Good sir, be not offended. She is an heir of
Falconbridge.

LONGAVILLE  Nay, my choler is ended. She is a most
sweet lady.

BOYET  Not unlike, sir, that may be.
[Longaville exits.]

[Enter Berowne.]


BEROWNE, [to Boyet]  What's her name in the cap?

BOYET  Rosaline, by good hap.

BEROWNE  Is she wedded or no?

BOYET  To her will, sir, or so.

BEROWNE  You are welcome, sir. Adieu.

BOYET  Farewell to me, sir, and welcome to you.
[Berowne exits.]

MARIA
That last is Berowne, the merry madcap lord.
Not a word with him but a jest.

BOYET  And every jest but
a word.

PRINCESS
It was well done of you to take him at his word.

BOYET
I was as willing to grapple as he was to board.

KATHERINE
Two hot sheeps, marry.

BOYET  And wherefore not ships?
No sheep, sweet lamb, unless we feed on your lips.

KATHERINE
You sheep and I pasture. Shall that finish the jest?

BOYET
So you grant pasture for me.	[He tries to kiss her.]

KATHERINE  Not so, gentle beast,
My lips are no common, though several they be.

BOYET
Belonging to whom?

KATHERINE  To my fortunes and me.

PRINCESS
Good wits will be jangling; but, gentles, agree,
This civil war of wits were much better used
On Navarre and his bookmen, for here 'tis abused.

BOYET
If my observation, which very seldom lies,
By the heart's still rhetoric, disclosed wi' th' eyes,
Deceive me not now, Navarre is infected.

PRINCESS  With what?

BOYET
With that which we lovers entitle "affected."

PRINCESS  Your reason?

BOYET
Why, all his behaviors did make their retire
To the court of his eye, peeping thorough desire.
His heart like an agate with your print impressed,
Proud with his form, in his eye pride expressed.
His tongue, all impatient to speak and not see,
Did stumble with haste in his eyesight to be;
All senses to that sense did make their repair,
To feel only looking on fairest of fair.
Methought all his senses were locked in his eye,
As jewels in crystal for some prince to buy,
Who, tend'ring their own worth from where they
were glassed,
Did point you to buy them along as you passed.
His face's own margent did quote such amazes
That all eyes saw his eyes enchanted with gazes.
I'll give you Aquitaine, and all that is his,
An you give him for my sake but one loving kiss.

PRINCESS, [to her Ladies]
Come, to our pavilion. Boyet is disposed.

BOYET
But to speak that in words which his eye hath
disclosed.
I only have made a mouth of his eye
By adding a tongue which I know will not lie.

MARIA
Thou art an old lovemonger and speakest skillfully.

KATHERINE
He is Cupid's grandfather, and learns news of him.

ROSALINE
Then was Venus like her mother, for her father is
but grim.

BOYET
Do you hear, my mad wenches?

MARIA  No.

BOYET  What then, do
you see?

MARIA
Ay, our way to be gone.

BOYET  You are too hard for me.
[They all exit.]


ACT 3
=====

Scene 1
=======
[Enter Braggart Armado and his Boy.]


ARMADO  Warble, child, make passionate my sense of
hearing.

BOY [sings]  Concolinel.

ARMADO  Sweet air. Go, tenderness of years. [He hands
over a key.] Take this key, give enlargement to the
swain, bring him festinately hither. I must employ
him in a letter to my love.

BOY  Master, will you win your love with a French
brawl?

ARMADO  How meanest thou? Brawling in French?

BOY  No, my complete master, but to jig off a tune at the
tongue's end, canary to it with your feet, humor it
with turning up your eyelids, sigh a note and sing a
note, sometimes through the throat as if you
swallowed love with singing love, sometimes
through the nose as if you snuffed up love by
smelling love; with your hat penthouse-like o'er the
shop of your eyes, with your arms crossed on your
thin-belly doublet like a rabbit on a spit; or your
hands in your pocket like a man after the old
painting; and keep not too long in one tune, but a
snip and away. These are compliments, these are
humors; these betray nice wenches that would be
betrayed without these, and make them men of
note--do you note me?--that most are affected
to these.

ARMADO  How hast thou purchased this experience?

BOY  By my penny of observation.

ARMADO  But O-- but O--.

BOY  "The hobby-horse is forgot."

ARMADO  Call'st thou my love "hobby-horse"?

BOY  No, master. The hobby-horse is but a colt, [aside]
and your love perhaps a hackney.--But have you
forgot your love?

ARMADO  Almost I had.

BOY  Negligent student, learn her by heart.

ARMADO  By heart and in heart, boy.

BOY  And out of heart, master. All those three I will
prove.

ARMADO  What wilt thou prove?

BOY  A man, if I live; and this "by, in, and without,"
upon the instant: "by" heart you love her, because
your heart cannot come by her; "in" heart you love
her, because your heart is in love with her; and
"out" of heart you love her, being out of heart that
you cannot enjoy her.

ARMADO  I am all these three.

BOY  And three times as much more, [aside] and yet
nothing at all.

ARMADO  Fetch hither the swain. He must carry me a
letter.

BOY  A message well sympathized--a horse to be ambassador
for an ass.

ARMADO  Ha? Ha? What sayest thou?

BOY  Marry, sir, you must send the ass upon the horse,
for he is very slow-gaited. But I go.

ARMADO  The way is but short. Away!

BOY  As swift as lead, sir.

ARMADO  Thy meaning, pretty ingenious?
Is not lead a metal heavy, dull, and slow?

BOY
Minime, honest master, or rather, master, no.

ARMADO
I say lead is slow.

BOY  You are too swift, sir, to say so.
Is that lead slow which is fired from a gun?

ARMADO  Sweet smoke of rhetoric!
He reputes me a cannon, and the bullet, that's
he.--
I shoot thee at the swain.

BOY  Thump, then, and I flee.
[He exits.]

ARMADO
A most acute juvenal, voluble and free of grace.
By thy favor, sweet welkin, I must sigh in thy face.
Most rude melancholy, valor gives thee place.
My herald is returned.

[Enter Boy and Clown Costard.]


BOY  A wonder, master!
Here's a costard broken in a shin.

ARMADO
Some enigma, some riddle. Come, thy l'envoi begin.

COSTARD  No egma, no riddle, no l'envoi, no salve in
the mail, sir. O, sir, plantain, a plain plantain! No
l'envoi, no l'envoi, no salve, sir, but a plantain.

ARMADO  By virtue, thou enforcest laughter; thy silly
thought, my spleen. The heaving of my lungs
provokes me to ridiculous smiling. O pardon me,
my stars! Doth the inconsiderate take salve for
l'envoi, and the word l'envoi for a salve?

BOY
Do the wise think them other? Is not l'envoi a salve?

ARMADO
No, page, it is an epilogue or discourse to make plain
Some obscure precedence that hath tofore been sain.
I will example it:
	The fox, the ape, and the humble-bee
	Were still at odds, being but three.
There's the moral. Now the l'envoi.

BOY  I will add the l'envoi. Say the moral again.

ARMADO
	The fox, the ape, and the humble-bee
	Were still at odds, being but three.

BOY
	Until the goose came out of door
	And stayed the odds by adding four.
Now will I begin your moral, and do you follow with
my l'envoi.
	The fox, the ape, and the humble-bee
	Were still at odds, being but three.

ARMADO
	Until the goose came out of door,
	Staying the odds by adding four.

BOY  A good l'envoi, ending in the goose. Would you
desire more?

COSTARD
The boy hath sold him a bargain--a goose, that's
flat.--
Sir, your pennyworth is good, an your goose be fat.
To sell a bargain well is as cunning as fast and
loose.
Let me see: a fat l'envoi--ay, that's a fat goose.

ARMADO
Come hither, come hither. How did this argument
begin?

BOY
By saying that a costard was broken in a shin.
Then called you for the l'envoi.

COSTARD  True, and I for a plantain. Thus came your
argument in. Then the boy's fat l'envoi, the goose
that you bought; and he ended the market.

ARMADO  But tell me, how was there a costard broken
in a shin?

BOY  I will tell you sensibly.

COSTARD  Thou hast no feeling of it, Mote. I will speak
that l'envoi.
	I, Costard, running out, that was safely within,
	Fell over the threshold and broke my shin.

ARMADO  We will talk no more of this matter.

COSTARD  Till there be more matter in the shin.

ARMADO  Sirrah Costard, I will enfranchise thee.

COSTARD  O, marry me to one Frances! I smell some
l'envoi, some goose, in this.

ARMADO  By my sweet soul, I mean, setting thee at
liberty, enfreedoming thy person. Thou wert immured,
restrained, captivated, bound.

COSTARD  True, true; and now you will be my purgation,
and let me loose.

ARMADO  I give thee thy liberty, set thee from durance,
and, in lieu thereof, impose on thee nothing but
this: bear this significant to the country maid
Jaquenetta. [(He gives him a paper.)] There is remuneration
[(giving him a coin,)] for the best ward of
mine honor is rewarding my dependents.--Mote,
follow.	[He exits.]

BOY  Like the sequel, I. Signior Costard, adieu.
[He exits.]

COSTARD
My sweet ounce of man's flesh, my incony Jew!
Now will I look to his remuneration. [He looks at the
coin.] "Remuneration"! O, that's the Latin word for
three farthings. Three farthings--remuneration.
"What's the price of this inkle?" "One penny." "No,
I'll give you a remuneration." Why, it carries it!
Remuneration. Why, it is a fairer name than "French
crown." I will never buy and sell out of this word.

[Enter Berowne.]


BEROWNE  My good knave Costard, exceedingly well
met.

COSTARD  Pray you, sir, how much carnation ribbon
may a man buy for a remuneration?

BEROWNE  What is a remuneration?

COSTARD  Marry, sir, halfpenny farthing.

BEROWNE  Why then, three farthing worth of silk.

COSTARD  I thank your Worship. God be wi' you.
[He begins to exit.]

BEROWNE  Stay, slave, I must employ thee.
As thou wilt win my favor, good my knave,
Do one thing for me that I shall entreat.

COSTARD  When would you have it done, sir?

BEROWNE  This afternoon.

COSTARD  Well, I will do it, sir. Fare you well.

BEROWNE  Thou knowest not what it is.

COSTARD  I shall know, sir, when I have done it.

BEROWNE  Why, villain, thou must know first.

COSTARD  I will come to your Worship tomorrow
morning.

BEROWNE  It must be done this afternoon. Hark, slave,
it is but this:
The Princess comes to hunt here in the park,
And in her train there is a gentle lady.
When tongues speak sweetly, then they name her
name,
And Rosaline they call her. Ask for her,
And to her white hand see thou do commend
This sealed-up counsel. There's thy guerdon. [He
gives him money.] Go.

COSTARD  Gardon. [He looks at the money.] O sweet
gardon! Better than remuneration, a 'levenpence
farthing better! Most sweet gardon. I will do it, sir,
in print. Gardon! Remuneration!	[He exits.]

BEROWNE
And I forsooth in love! I that have been love's whip,
A very beadle to a humorous sigh,
A critic, nay, a nightwatch constable,
A domineering pedant o'er the boy,
Than whom no mortal so magnificent.
This wimpled, whining, purblind, wayward boy,
This Signior Junior, giant dwarf, Dan Cupid,
Regent of love rhymes, lord of folded arms,
Th' anointed sovereign of sighs and groans,
Liege of all loiterers and malcontents,
Dread prince of plackets, king of codpieces,
Sole imperator and great general
Of trotting paritors--O my little heart!
And I to be a corporal of his field
And wear his colors like a tumbler's hoop!
What? I love, I sue, I seek a wife?
A woman, that is like a German clock,
Still a-repairing, ever out of frame,
And never going aright, being a watch,
But being watched that it may still go right.
Nay, to be perjured, which is worst of all.
And, among three, to love the worst of all,
A whitely wanton with a velvet brow,
With two pitch-balls stuck in her face for eyes.
Ay, and by heaven, one that will do the deed
Though Argus were her eunuch and her guard.
And I to sigh for her, to watch for her,
To pray for her! Go to. It is a plague
That Cupid will impose for my neglect
Of his almighty dreadful little might.
Well, I will love, write, sigh, pray, sue, groan.
Some men must love my lady, and some Joan.
[He exits.]


ACT 4
=====

Scene 1
=======
[Enter the Princess, a Forester, her Ladies, Boyet and
her other Lords.]


PRINCESS
Was that the King that spurred his horse so hard
Against the steep uprising of the hill?

FORESTER
I know not, but I think it was not he.

PRINCESS
Whoe'er he was, he showed a mounting mind.--
Well, lords, today we shall have our dispatch.
Or Saturday we will return to France.--
Then, forester, my friend, where is the bush
That we must stand and play the murderer in?

FORESTER
Hereby, upon the edge of yonder coppice,
A stand where you may make the fairest shoot.

PRINCESS
I thank my beauty, I am fair that shoot,
And thereupon thou speakst "the fairest shoot."

FORESTER
Pardon me, madam, for I meant not so.

PRINCESS
What, what? First praise me, and again say no?
O short-lived pride. Not fair? Alack, for woe!

FORESTER
Yes, madam, fair.

PRINCESS  Nay, never paint me now.
Where fair is not, praise cannot mend the brow.
Here, good my glass, take this for telling true.
[She gives him money.]
Fair payment for foul words is more than due.

FORESTER
Nothing but fair is that which you inherit.

PRINCESS
See, see, my beauty will be saved by merit.
O heresy in fair, fit for these days!
A giving hand, though foul, shall have fair praise.
But come, the bow. [He hands her a bow.] Now
mercy goes to kill,
And shooting well is then accounted ill.
Thus will I save my credit in the shoot:
Not wounding, pity would not let me do 't;
If wounding, then it was to show my skill,
That more for praise than purpose meant to kill.
And out of question so it is sometimes:
Glory grows guilty of detested crimes,
When for fame's sake, for praise, an outward part,
We bend to that the working of the heart;
As I for praise alone now seek to spill
The poor deer's blood, that my heart means no ill.

BOYET
Do not curst wives hold that self sovereignty
Only for praise' sake when they strive to be
Lords o'er their lords?

PRINCESS
Only for praise; and praise we may afford
To any lady that subdues a lord.

[Enter Clown Costard.]


BOYET
Here comes a member of the commonwealth.

COSTARD  God dig-you-den all! Pray you, which is the
head lady?

PRINCESS  Thou shalt know her, fellow, by the rest that
have no heads.

COSTARD  Which is the greatest lady, the highest?

PRINCESS  The thickest and the tallest.

COSTARD
The thickest and the tallest: it is so, truth is
truth.
An your waist, mistress, were as slender as my wit,
One o' these maids' girdles for your waist should be
fit.
Are not you the chief woman? You are the thickest
here.

PRINCESS  What's your will, sir? What's your will?

COSTARD  I have a letter from Monsieur Berowne to
one Lady Rosaline.

PRINCESS
O, thy letter, thy letter! He's a good friend of mine.
Stand aside, good bearer.--Boyet, you can carve.
Break up this capon.

BOYET, [taking the letter]  I am bound to serve.
This letter is mistook; it importeth none here.
It is writ to Jaquenetta.

PRINCESS  We will read it, I swear.
Break the neck of the wax, and everyone give ear.

BOYET [reads.]  By heaven, that thou art fair is most
infallible, true that thou art beauteous, truth itself
that thou art lovely. More fairer than fair, beautiful
than beauteous, truer than truth itself, have commiseration
on thy heroical vassal. The magnanimous and
most illustrate King Cophetua set eye upon the pernicious
and indubitate beggar Zenelophon; and he it
was that might rightly say "Veni, vidi, vici," which to
annothanize in the vulgar (O base and obscure vulgar!)
videlicet, "He came, see, and overcame": He
came, one; see, two; overcame, three. Who came? The
King. Why did he come? To see. Why did he see? To
overcome. To whom came he? To the beggar. What
saw he? The beggar. Who overcame he? The beggar.
The conclusion is victory. On whose side? The
King's. The captive is enriched. On whose side? The
beggar's. The catastrophe is a nuptial. On whose side?
The King's--no, on both in one, or one in both. I am
the King, for so stands the comparison; thou the
beggar, for so witnesseth thy lowliness. Shall I command
thy love? I may. Shall I enforce thy love? I could.
Shall I entreat thy love? I will. What shalt thou
exchange for rags? Robes. For tittles? Titles. For thyself?
Me. Thus expecting thy reply, I profane my lips on thy
foot, my eyes on thy picture, and my heart on thy every
part.
	Thine, in the dearest design of industry,
Don Adriano de Armado.
	Thus dost thou hear the Nemean lion roar
	   'Gainst thee, thou lamb, that standest as his prey.
	Submissive fall his princely feet before,
	   And he from forage will incline to play.
	But if thou strive, poor soul, what art thou then?
	Food for his rage, repasture for his den.

PRINCESS
What plume of feathers is he that indited this letter?
What vane? What weathercock? Did you ever hear
better?

BOYET
I am much deceived but I remember the style.

PRINCESS
Else your memory is bad, going o'er it erewhile.

BOYET
This Armado is a Spaniard that keeps here in court,
A phantasime, a Monarcho, and one that makes
sport
To the Prince and his bookmates.

PRINCESS, [to Costard]  Thou, fellow, a word.
Who gave thee this letter?

COSTARD  I told you: my lord.

PRINCESS
To whom shouldst thou give it?

COSTARD  From my lord to my
lady.

PRINCESS  From which lord to which lady?

COSTARD
From my Lord Berowne, a good master of mine,
To a lady of France that he called Rosaline.

PRINCESS
Thou hast mistaken his letter. Come, lords, away.
[To Rosaline.] Here, sweet, put up this; 'twill be
thine another day.
[The Princess, Katherine, Lords, and
Forester exit. Boyet, Rosaline, Maria,
and Costard remain.]

BOYET
Who is the shooter? Who is the shooter?

ROSALINE  Shall I
teach you to know?

BOYET
Ay, my continent of beauty.

ROSALINE  Why, she that bears the bow.
Finely put off.

BOYET
My lady goes to kill horns, but if thou marry,
Hang me by the neck if horns that year miscarry.
Finely put on.

ROSALINE
Well, then, I am the shooter.

BOYET  And who is your deer?

ROSALINE
If we choose by the horns, yourself come not near.
Finely put on, indeed.

MARIA
You still wrangle with her, Boyet, and she strikes at
the brow.

BOYET
But she herself is hit lower. Have I hit her now?

ROSALINE  Shall I come upon thee with an old saying,
that was a man when King Pippen of France was a
little boy, as touching the hit it?

BOYET  So I may answer thee with one as old, that was a
woman when Queen Guinover of Britain was a little
wench, as touching the hit it.

ROSALINE [sings]
	Thou canst not hit it, hit it, hit it,
	Thou canst not hit it, my good man.

BOYET [sings]
	An I cannot, cannot, cannot,
	An I cannot, another can.
[Rosaline exits.]

COSTARD
By my troth, most pleasant. How both did fit it!

MARIA
A mark marvelous well shot, for they both did hit
it.

BOYET
A mark! O, mark but that mark. "A mark," says my
lady.
Let the mark have a prick in 't to mete at, if it may
be.

MARIA
Wide o' the bow hand! I' faith, your hand is out.

COSTARD
Indeed, he must shoot nearer, or he'll ne'er hit the
clout.

BOYET, [to Maria]
An if my hand be out, then belike your hand is in.

COSTARD
Then will she get the upshoot by cleaving the pin.

MARIA
Come, come, you talk greasily. Your lips grow foul.

COSTARD, [to Boyet]
She's too hard for you at pricks, sir. Challenge her
to bowl.

BOYET
I fear too much rubbing. Good night, my good owl.
[Boyet and Maria exit.]

COSTARD
By my soul, a swain, a most simple clown.
Lord, Lord, how the ladies and I have put him
down.
O' my troth, most sweet jests, most incony vulgar
wit,
When it comes so smoothly off, so obscenely, as it
were, so fit.
Armado o' th' one side, O, a most dainty man!
To see him walk before a lady and to bear her fan.
To see him kiss his hand, and how most sweetly he
will swear.
And his page o' t' other side, that handful of wit!
Ah heavens, it is a most pathetical nit.
[Shout within.]
Sola, sola!
[He exits.]

Scene 2
=======
[Enter Dull the Constable, Holofernes the Pedant, and
Nathaniel the Curate.]


NATHANIEL  Very reverend sport, truly, and done in the
testimony of a good conscience.

HOLOFERNES  The deer was, as you know, sanguis, in
blood, ripe as the pomewater, who now hangeth
like a jewel in the ear of caelo, the sky, the welkin,
the heaven, and anon falleth like a crab on the face
of terra, the soil, the land, the earth.

NATHANIEL  Truly, Master Holofernes, the epithets are
sweetly varied, like a scholar at the least. But, sir, I
assure you, it was a buck of the first head.

HOLOFERNES  Sir Nathaniel, haud credo.

DULL  'Twas not a haud credo, 'twas a pricket.

HOLOFERNES  Most barbarous intimation! Yet a kind of
insinuation, as it were, in via, in way, of explication;
facere, as it were, replication, or rather, ostentare, to
show, as it were, his inclination, after his undressed,
unpolished, uneducated, unpruned, untrained, or
rather unlettered, or ratherest, unconfirmed fashion,
to insert again my haud credo for a deer.

DULL  I said the deer was not a haud credo, 'twas a
pricket.

HOLOFERNES  Twice-sod simplicity, bis coctus!
O thou monster ignorance, how deformed dost thou
look!

NATHANIEL
Sir, he hath never fed of the dainties that are bred
in a book.
He hath not eat paper, as it were; he hath not drunk
ink. His intellect is not replenished. He is only an
animal, only sensible in the duller parts.
And such barren plants are set before us that we
thankful should be--
Which we of taste and feeling are--for those parts
that do fructify in us more than he.
For as it would ill become me to be vain, indiscreet,
or a fool,
So were there a patch set on learning, to see him in
a school.
But omne bene, say I, being of an old father's mind:
Many can brook the weather that love not the wind.

DULL
You two are bookmen. Can you tell me by your wit
What was a month old at Cain's birth that's not
five weeks old as yet?

HOLOFERNES  Dictynna, goodman Dull, Dictynna,
goodman Dull.

DULL  What is "dictima"?

NATHANIEL
A title to Phoebe, to Luna, to the moon.

HOLOFERNES
The moon was a month old when Adam was no
more.
And raught not to five weeks when he came to
fivescore.
Th' allusion holds in the exchange.

DULL  'Tis true indeed. The collusion holds in the
exchange.

HOLOFERNES  God comfort thy capacity! I say, th' allusion
holds in the exchange.

DULL  And I say the pollution holds in the exchange, for
the moon is never but a month old. And I say besides
that, 'twas a pricket that the Princess killed.

HOLOFERNES  Sir Nathaniel, will you hear an extemporal
epitaph on the death of the deer? And, to humor
the ignorant, call I the deer the Princess killed a
pricket.

NATHANIEL  Perge, good Master Holofernes, perge, so it
shall please you to abrogate scurrility.

HOLOFERNES  I will something affect the letter, for it
argues facility.
The preyful princess pierced and pricked
a pretty pleasing pricket,
   Some say a sore, but not a sore till now made
   sore with shooting.
The dogs did yell. Put "l" to "sore," then sorel
jumps from thicket,
   Or pricket sore, or else sorel. The people fall
   a-hooting.
If sore be sore, then "L" to "sore" makes fifty
sores o' sorel.
Of one sore I an hundred make by adding but one
more "L."

NATHANIEL  A rare talent.

DULL, [aside]  If a talent be a claw, look how he claws
him with a talent.

HOLOFERNES  This is a gift that I have, simple, simple--
a foolish extravagant spirit, full of forms,
figures, shapes, objects, ideas, apprehensions, motions,
revolutions. These are begot in the ventricle
of memory, nourished in the womb of pia mater,
and delivered upon the mellowing of occasion. But
the gift is good in those in whom it is acute, and I
am thankful for it.

NATHANIEL  Sir, I praise the Lord for you, and so may
my parishioners, for their sons are well tutored by
you, and their daughters profit very greatly under
you. You are a good member of the
commonwealth.

HOLOFERNES  Mehercle, if their sons be ingenious,
they shall want no instruction; if their daughters be
capable, I will put it to them. But Vir sapis qui pauca
loquitur. A soul feminine saluteth us.

[Enter Jaquenetta and the Clown Costard.]


JAQUENETTA, [to Nathaniel]  God give you good morrow,
Master Person.

HOLOFERNES  Master Person, quasi pierce one. And
if one should be pierced, which is the one?

COSTARD  Marry, Master Schoolmaster, he that is likeliest
to a hogshead.

HOLOFERNES  Of piercing a hogshead! A good luster
of conceit in a turf of earth; fire enough for a flint,
pearl enough for a swine. 'Tis pretty, it is well.

JAQUENETTA, [to Nathaniel]  Good Master Parson, be so
good as read me this letter. It was given me by
Costard, and sent me from Don Armado. I beseech
you, read it.
[She hands Nathaniel a paper, which he looks at.]

HOLOFERNES
	Facile precor gelida quando peccas omnia sub umbra.
Ruminat--
and so forth. Ah, good old Mantuan! I may speak of
thee as the traveler doth of Venice:
	Venetia, Venetia,
	Chi non ti vede, non ti pretia.
Old Mantuan, old Mantuan! Who understandeth
thee not, loves thee not. [(He sings.)] Ut, re, sol, la,
mi, fa. [(To Nathaniel.)] Under pardon, sir, what are
the contents? Or rather, as Horace says in his--
[(Looking at the letter.)] What, my soul, verses?

NATHANIEL  Ay, sir, and very learned.

HOLOFERNES   Let me hear a staff, a stanza, a verse,
Lege, domine.

NATHANIEL, [reads]
	If love make me forsworn, how shall I swear to love?
	   Ah, never faith could hold, if not to beauty vowed!
	Though to myself forsworn, to thee I'll faithful prove.
	   Those thoughts to me were oaks, to thee like osiers
   bowed.
	Study his bias leaves and makes his book thine eyes,
	   Where all those pleasures live that art would
   comprehend.
	If knowledge be the mark, to know thee shall suffice.
	   Well-learned is that tongue that well can thee
   commend.
	All ignorant that soul that sees thee without wonder;
	   Which is to me some praise that I thy parts admire.
	Thy eye Jove's lightning bears, thy voice his dreadful
thunder,
	   Which, not to anger bent, is music and sweet fire.
	Celestial as thou art, O, pardon love this wrong,
	That sings heaven's praise with such an earthly tongue.

HOLOFERNES  You find not the apostrophus, and so
miss the accent. Let me supervise the canzonet.
[He takes the paper.] Here are only numbers ratified,
but, for the elegancy, facility, and golden cadence of
poesy--caret. Ovidius Naso was the man. And why
indeed "Naso," but for smelling out the odoriferous
flowers of fancy, the jerks of invention? Imitari is
nothing: so doth the hound his master, the ape his
keeper, the tired horse his rider.--But damosella
virgin, was this directed to you?

JAQUENETTA  Ay, sir, from one Monsieur Berowne, one
of the strange queen's lords.

HOLOFERNES  I will overglance the superscript: "To
the snow-white hand of the most beauteous Lady
Rosaline." I will look again on the intellect of the
letter for the nomination of the party writing to
the person written unto: "Your Ladyship's in all
desired employment, Berowne." Sir Nathaniel, this
Berowne is one of the votaries with the King, and
here he hath framed a letter to a sequent of the
stranger queen's: which accidentally, or by the way
of progression, hath miscarried. [To Jaquenetta.]
Trip and go, my sweet. Deliver this paper into the
royal hand of the King. It may concern much. Stay
not thy compliment. I forgive thy duty. Adieu.

JAQUENETTA  Good Costard, go with me.--Sir, God
save your life.

COSTARD  Have with thee, my girl.
[Costard and Jaquenetta exit.]

NATHANIEL  Sir, you have done this in the fear of God
very religiously; and, as a certain Father saith--

HOLOFERNES  Sir, tell not me of the Father. I do fear
colorable colors. But to return to the verses: did
they please you, Sir Nathaniel?

NATHANIEL  Marvelous well for the pen.

HOLOFERNES  I do dine today at the father's of a certain
pupil of mine, where if, before repast, it shall
please you to gratify the table with a grace, I will,
on my privilege I have with the parents of the
foresaid child or pupil, undertake your ben venuto;
where I will prove those verses to be very unlearned,
neither savoring of poetry, wit, nor invention.
I beseech your society.

NATHANIEL  And thank you too; for society, saith the
text, is the happiness of life.

HOLOFERNES  And certes the text most infallibly concludes
it. [To Dull.] Sir, I do invite you too. You shall
not say me nay. Pauca verba. Away! The gentles are
at their game, and we will to our recreation.
[They exit.]

Scene 3
=======
[Enter Berowne with a paper in his hand, alone.]


BEROWNE  The King, he is hunting the deer; I am
coursing myself. They have pitched a toil; I am
toiling in a pitch--pitch that defiles. Defile! A foul
word. Well, "set thee down, sorrow"; for so they
say the fool said, and so say I, and I the fool. Well
proved, wit. By the Lord, this love is as mad as Ajax.
It kills sheep, it kills me, I a sheep. Well proved
again, o' my side. I will not love. If I do, hang me. I'
faith, I will not. O, but her eye! By this light, but for
her eye I would not love her; yes, for her two eyes.
Well, I do nothing in the world but lie, and lie in my
throat. By heaven, I do love, and it hath taught me to
rhyme, and to be melancholy. And here is part of my
rhyme, and here my melancholy. Well, she hath one
o' my sonnets already. The clown bore it, the fool
sent it, and the lady hath it. Sweet clown, sweeter
fool, sweetest lady. By the world, I would not care a
pin, if the other three were in. Here comes one with
a paper. God give him grace to groan.
[He stands aside.]

[The King entereth with a paper.]


KING  Ay me!

BEROWNE, [aside]  Shot, by heaven! Proceed, sweet
Cupid. Thou hast thumped him with thy birdbolt
under the left pap. In faith, secrets!

KING [reads]
	So sweet a kiss the golden sun gives not
	   To those fresh morning drops upon the rose
	As thy eyebeams, when their fresh rays have smote
	   The night of dew that on my cheeks down flows.
	Nor shines the silver moon one-half so bright
	   Through the transparent bosom of the deep
	As doth thy face, through tears of mine, give light.
	   Thou shin'st in every tear that I do weep.
	No drop but as a coach doth carry thee;
	   So ridest thou triumphing in my woe.
	Do but behold the tears that swell in me,
	   And they thy glory through my grief will show.
	But do not love thyself; then thou wilt keep
	My tears for glasses, and still make me weep.
	O queen of queens, how far dost thou excel
	No thought can think, nor tongue of mortal tell.

How shall she know my griefs? I'll drop the paper.
Sweet leaves, shade folly. Who is he comes here?

[Enter Longaville, with papers. The King steps aside.]

What, Longaville, and reading! Listen, ear.

BEROWNE, [aside]
Now, in thy likeness, one more fool appear!

LONGAVILLE  Ay me! I am forsworn.

BEROWNE, [aside]
Why, he comes in like a perjure, wearing papers!

KING, [aside]
In love, I hope! Sweet fellowship in shame.

BEROWNE, [aside]
One drunkard loves another of the name.

LONGAVILLE
Am I the first that have been perjured so?

BEROWNE, [aside]
I could put thee in comfort: not by two that I know.
Thou makest the triumviry, the corner-cap of
society,
The shape of love's Tyburn, that hangs up simplicity.

LONGAVILLE
I fear these stubborn lines lack power to move.
[Reads.] O sweet Maria, empress of my love--
These numbers will I tear and write in prose.
[He tears the paper.]

BEROWNE, [aside]
O, rhymes are guards on wanton Cupid's hose.
Disfigure not his shop!

LONGAVILLE, [taking another paper]  This same shall go.
[He reads the sonnet.]
	Did not the heavenly rhetoric of thine eye,
	   'Gainst whom the world cannot hold argument,
	Persuade my heart to this false perjury?
	   Vows for thee broke deserve not punishment.
	A woman I forswore, but I will prove,
	   Thou being a goddess, I forswore not thee.
	My vow was earthly, thou a heavenly love.
	   Thy grace being gained cures all disgrace in me.
	Vows are but breath, and breath a vapor is.
	   Then thou, fair sun, which on my Earth dost
   shine,
	Exhal'st this vapor-vow; in thee it is.
	   If broken, then, it is no fault of mine.
	If by me broke, what fool is not so wise
	To lose an oath to win a paradise?

BEROWNE, [aside]
This is the liver vein, which makes flesh a deity,
A green goose a goddess. Pure, pure idolatry.
God amend us, God amend. We are much out o' th'
way.

LONGAVILLE
By whom shall I send this?--Company? Stay.
[He steps aside.]

[Enter Dumaine, with a paper.]


BEROWNE, [aside]
All hid, all hid--an old infant play.
Like a demigod here sit I in the sky,
And wretched fools' secrets heedfully o'ereye.
More sacks to the mill. O heavens, I have my wish.
Dumaine transformed! Four woodcocks in a dish.

DUMAINE  O most divine Kate!

BEROWNE, [aside]  O most profane coxcomb!

DUMAINE
By heaven, the wonder in a mortal eye!

BEROWNE, [aside]
By Earth, she is not, corporal. There you lie.

DUMAINE
Her amber hairs for foul hath amber quoted.

BEROWNE, [aside]
An amber-colored raven was well noted.

DUMAINE
As upright as the cedar.

BEROWNE, [aside]  Stoop, I say.
Her shoulder is with child.

DUMAINE  As fair as day.

BEROWNE, [aside]
Ay, as some days, but then no sun must shine.

DUMAINE
O, that I had my wish!

LONGAVILLE, [aside]  And I had mine!

KING, [aside]  And mine too, good Lord!

BEROWNE, [aside]
Amen, so I had mine. Is not that a good word?

DUMAINE
I would forget her, but a fever she
Reigns in my blood, and will remembered be.

BEROWNE, [aside]
A fever in your blood? Why, then incision
Would let her out in saucers! Sweet misprision.

DUMAINE
Once more I'll read the ode that I have writ.

BEROWNE, [aside]
Once more I'll mark how love can vary wit.

DUMAINE [reads his sonnet.]
	On a day--alack the day!--
	Love, whose month is ever May,
	Spied a blossom passing fair,
	Playing in the wanton air.
	Through the velvet leaves the wind,
	All unseen, can passage find;
	That the lover, sick to death,
	Wished himself the heaven's breath.
	"Air," quoth he, "thy cheeks may blow.
	Air, would I might triumph so!"
	But, alack, my hand is sworn
	Ne'er to pluck thee from thy thorn.
	Vow, alack, for youth unmeet,
	Youth so apt to pluck a sweet.
	Do not call it sin in me
	That I am forsworn for thee--
	Thou for whom Jove would swear
	Juno but an Ethiope were,
	And deny himself for Jove,
	Turning mortal for thy love.
This will I send, and something else more plain
That shall express my true love's fasting pain.
O, would the King, Berowne, and Longaville
Were lovers too! Ill to example ill
Would from my forehead wipe a perjured note,
For none offend where all alike do dote.

LONGAVILLE, [coming forward]
Dumaine, thy love is far from charity,
That in love's grief desir'st society.
You may look pale, but I should blush, I know,
To be o'er-heard and taken napping so.

KING, [coming forward]
[To Longaville.] Come, sir, you blush! As his, your
case is such.
You chide at him, offending twice as much.
You do not love Maria? Longaville
Did never sonnet for her sake compile,
Nor never lay his wreathed arms athwart
His loving bosom to keep down his heart?
I have been closely shrouded in this bush
And marked you both, and for you both did blush.
I heard your guilty rhymes, observed your fashion,
Saw sighs reek from you, noted well your passion.
"Ay, me!" says one. "O Jove!" the other cries.
One, her hairs were gold, crystal the other's eyes.
[To Longaville.] You would for paradise break faith
and troth,
[To Dumaine.] And Jove, for your love, would
infringe an oath.
What will Berowne say when that he shall hear
Faith infringed, which such zeal did swear?
How will he scorn, how will he spend his wit!
How will he triumph, leap, and laugh at it!
For all the wealth that ever I did see,
I would not have him know so much by me.

BEROWNE, [coming forward]
Now step I forth to whip hypocrisy.
Ah, good my liege, I pray thee pardon me.
Good heart, what grace hast thou thus to reprove
These worms for loving, that art most in love?
Your eyes do make no coaches; in your tears
There is no certain princess that appears.
You'll not be perjured, 'tis a hateful thing!
Tush, none but minstrels like of sonneting!
But are you not ashamed? Nay, are you not,
All three of you, to be thus much o'ershot?
[To Longaville.] You found his mote, the King your
mote did see,
But I a beam do find in each of three.
O, what a scene of fool'ry have I seen,
Of sighs, of groans, of sorrow, and of teen!
O me, with what strict patience have I sat,
To see a king transformed to a gnat!
To see great Hercules whipping a gig,
And profound Solomon to tune a jig,
And Nestor play at pushpin with the boys,
And critic Timon laugh at idle toys.
Where lies thy grief, O tell me, good Dumaine?
And gentle Longaville, where lies thy pain?
And where my liege's? All about the breast!
A caudle, ho!

KING  Too bitter is thy jest.
Are we betrayed thus to thy overview?

BEROWNE
Not you to me, but I betrayed by you.
I, that am honest, I, that hold it sin
To break the vow I am engaged in.
I am betrayed by keeping company
With men like you, men of inconstancy.
When shall you see me write a thing in rhyme?
Or groan for Joan? or spend a minute's time
In pruning me? When shall you hear that I
Will praise a hand, a foot, a face, an eye,
A gait, a state, a brow, a breast, a waist,
A leg, a limb--

[Enter Jaquenetta, with a paper, and Clown Costard.]
[Berowne begins to exit.]


KING  Soft, whither away so fast?
A true man, or a thief, that gallops so?

BEROWNE
I post from love. Good lover, let me go.

JAQUENETTA
God bless the King.

KING  What present hast thou there?

COSTARD
Some certain treason.

KING  What makes treason here?

COSTARD
Nay, it makes nothing, sir.

KING  If it mar nothing neither,
The treason and you go in peace away together.

JAQUENETTA
I beseech your Grace, let this letter be read.
Our person misdoubts it. 'Twas treason, he said.

KING
Berowne, read it over.
[Berowne reads the letter.]
[To Jaquenetta.] Where hadst thou it?

JAQUENETTA  Of Costard.

KING, [to Costard]  Where hadst thou it?

COSTARD  Of Dun Adramadio, Dun Adramadio.
[Berowne tears the paper.]

KING, [to Berowne]
How now, what is in you? Why dost thou tear it?

BEROWNE
A toy, my liege, a toy. Your Grace needs not fear it.

LONGAVILLE
It did move him to passion, and therefore let's hear
it.

DUMAINE, [picking up the papers]
It is Berowne's writing, and here is his name.

BEROWNE, [to Costard]
Ah, you whoreson loggerhead, you were born to do
me shame.--
Guilty, my lord, guilty. I confess, I confess.

KING  What?

BEROWNE
That you three fools lacked me fool to make up
the mess.
He, he, and you--and you, my liege--and I
Are pickpurses in love, and we deserve to die.
O, dismiss this audience, and I shall tell you more.

DUMAINE
Now the number is even.

BEROWNE  True, true, we are four.
[Pointing to Jaquenetta and Costard.] Will these
turtles be gone?

KING  Hence, sirs. Away.

COSTARD
Walk aside the true folk, and let the traitors stay.
[Jaquenetta and Costard exit.]

BEROWNE
Sweet lords, sweet lovers, O, let us embrace.
   As true we are as flesh and blood can be.
The sea will ebb and flow, heaven show his face;
   Young blood doth not obey an old decree.
We cannot cross the cause why we were born;
Therefore of all hands must we be forsworn.

KING
What, did these rent lines show some love of thine?

BEROWNE
Did they, quoth you? Who sees the heavenly
Rosaline
That, like a rude and savage man of Ind
   At the first op'ning of the gorgeous East,
Bows not his vassal head and, strucken blind,
   Kisses the base ground with obedient breast?
What peremptory eagle-sighted eye
   Dares look upon the heaven of her brow
That is not blinded by her majesty?

KING
   What zeal, what fury, hath inspired thee now?
My love, her mistress, is a gracious moon,
   She an attending star scarce seen a light.

BEROWNE
My eyes are then no eyes, nor I Berowne.
   O, but for my love, day would turn to night!
Of all complexions the culled sovereignty
   Do meet as at a fair in her fair cheek.
Where several worthies make one dignity,
   Where nothing wants that want itself doth seek.
Lend me the flourish of all gentle tongues--
   Fie, painted rhetoric! O, she needs it not!
To things of sale a seller's praise belongs.
   She passes praise. Then praise too short doth blot.
A withered hermit, fivescore winters worn,
   Might shake off fifty, looking in her eye.
Beauty doth varnish age, as if newborn,
   And gives the crutch the cradle's infancy.
O, 'tis the sun that maketh all things shine!

KING
   By heaven, thy love is black as ebony.

BEROWNE
Is ebony like her? O word divine!
   A wife of such wood were felicity.
O, who can give an oath? Where is a book,
   That I may swear beauty doth beauty lack
If that she learn not of her eye to look?
   No face is fair that is not full so black.

KING
O, paradox! Black is the badge of hell,
   The hue of dungeons and the school of night,
And beauty's crest becomes the heavens well.

BEROWNE
   Devils soonest tempt, resembling spirits of light.
O, if in black my lady's brows be decked,
   It mourns that painting and usurping hair
Should ravish doters with a false aspect:
   And therefore is she born to make black fair.
Her favor turns the fashion of the days,
   For native blood is counted painting now.
And therefore red, that would avoid dispraise,
   Paints itself black to imitate her brow.

DUMAINE
To look like her are chimney-sweepers black.

LONGAVILLE
   And since her time are colliers counted bright.

KING
And Ethiopes of their sweet complexion crack.

DUMAINE
   Dark needs no candles now, for dark is light.

BEROWNE
Your mistresses dare never come in rain,
   For fear their colors should be washed away.

KING
'Twere good yours did, for, sir, to tell you plain,
   I'll find a fairer face not washed today.

BEROWNE
I'll prove her fair, or talk till doomsday here.

KING
   No devil will fright thee then so much as she.

DUMAINE
I never knew man hold vile stuff so dear.

LONGAVILLE, [showing his shoe]
   Look, here's thy love; my foot and her face see.

BEROWNE
O, if the streets were paved with thine eyes.
   Her feet were much too dainty for such tread.

DUMAINE
O vile! Then as she goes, what upward lies
   The street should see as she walked overhead.

KING
But what of this? Are we not all in love?

BEROWNE
   Nothing so sure, and thereby all forsworn.

KING
Then leave this chat, and, good Berowne, now prove
   Our loving lawful, and our faith not torn.

DUMAINE
Ay, marry, there, some flattery for this evil.

LONGAVILLE
   O, some authority how to proceed,
Some tricks, some quillets, how to cheat the devil.

DUMAINE
   Some salve for perjury.

BEROWNE     O, 'tis more than need.
Have at you, then, affection's men-at-arms!
O, we have made a vow to study, lords,
And in that vow we have forsworn our books.
For when would you, my liege, or you, or you,
In leaden contemplation have found out
Such fiery numbers as the prompting eyes
Of beauty's tutors have enriched you with?
Other slow arts entirely keep the brain
And therefore, finding barren practicers,
Scarce show a harvest of their heavy toil.
But love, first learned in a lady's eyes,
Lives not alone immured in the brain,
But with the motion of all elements
Courses as swift as thought in every power,
And gives to every power a double power,
Above their functions and their offices.
It adds a precious seeing to the eye.
A lover's eyes will gaze an eagle blind.
A lover's ear will hear the lowest sound,
When the suspicious head of theft is stopped.
Love's feeling is more soft and sensible
Than are the tender horns of cockled snails.
Love's tongue proves dainty Bacchus gross in taste.
For valor, is not love a Hercules,
Still climbing trees in the Hesperides?
Subtle as Sphinx, as sweet and musical
As bright Apollo's lute strung with his hair.
And when love speaks, the voice of all the gods
Make heaven drowsy with the harmony.
Never durst poet touch a pen to write
Until his ink were tempered with love's sighs.
O, then his lines would ravish savage ears
And plant in tyrants mild humility.
From women's eyes this doctrine I derive.
They sparkle still the right Promethean fire.
They are the books, the arts, the academes
That show, contain, and nourish all the world.
Else none at all in ought proves excellent.
Then fools you were these women to forswear,
Or, keeping what is sworn, you will prove fools.
For wisdom's sake, a word that all men love,
Or for love's sake, a word that loves all men,
Or for men's sake, the authors of these women,
Or women's sake, by whom we men are men,
Let us once lose our oaths to find ourselves,
Or else we lose ourselves to keep our oaths.
It is religion to be thus forsworn,
For charity itself fulfills the law,
And who can sever love from charity?

KING
Saint Cupid, then, and, soldiers, to the field!

BEROWNE
Advance your standards, and upon them, lords.
Pell-mell, down with them. But be first advised
In conflict that you get the sun of them.

LONGAVILLE
Now to plain dealing. Lay these glozes by.
Shall we resolve to woo these girls of France?

KING
And win them, too. Therefore let us devise
Some entertainment for them in their tents.

BEROWNE
First, from the park let us conduct them thither.
Then homeward every man attach the hand
Of his fair mistress. In the afternoon
We will with some strange pastime solace them,
Such as the shortness of the time can shape;
For revels, dances, masques, and merry hours
Forerun fair love, strewing her way with flowers.

KING
Away, away! No time shall be omitted
That will betime and may by us be fitted.

BEROWNE
Allons! Allons! Sowed cockle reaped no corn,
   And justice always whirls in equal measure.
Light wenches may prove plagues to men forsworn;
   If so, our copper buys no better treasure.
[They exit.]


ACT 5
=====

Scene 1
=======
[Enter Holofernes the Pedant, Nathaniel the Curate,
and Dull the Constable.]


HOLOFERNES  Satis quid sufficit.

NATHANIEL  I praise God for you, sir. Your reasons at
dinner have been sharp and sententious, pleasant
without scurrility, witty without affection, audacious
without impudency, learned without opinion,
and strange without heresy. I did converse this
quondam day with a companion of the King's, who
is intituled, nominated, or called Don Adriano de
Armado.

HOLOFERNES  Novi hominem tanquam te. His humor
is lofty, his discourse peremptory, his tongue filed,
his eye ambitious, his gait majestical, and his general
behavior vain, ridiculous, and thrasonical. He is
too picked, too spruce, too affected, too odd, as it
were, too peregrinate, as I may call it.

NATHANIEL  A most singular and choice epithet.
[Draw out his table book.]

HOLOFERNES  He draweth out the thread of his verbosity
finer than the staple of his argument. I abhor
such fanatical phantasimes, such insociable and
point-devise companions, such rackers of orthography,
as to speak "dout," fine, when he should
say "doubt"; "det" when he should pronounce
"debt"--d, e, b, t, not d, e, t. He clepeth a calf
"cauf," half "hauf," neighbor vocatur "nebor";
neigh abbreviated ne. This is abhominable--which
he would call "abominable." It insinuateth me of
insanie. Ne intelligis, domine? To make frantic,
lunatic.

NATHANIEL  Laus Deo, bone intelligo.

HOLOFERNES  Bone? Bone for bene? Priscian a little
scratched; 'twill serve.

[Enter Armado the Braggart, Boy, and Costard.]


NATHANIEL  Videsne quis venit?

HOLOFERNES  Video, et gaudeo.

ARMADO  Chirrah.

HOLOFERNES  Quare "chirrah," not "sirrah"?

ARMADO  Men of peace, well encountered.

HOLOFERNES  Most military sir, salutation.

BOY, [aside to Costard]  They have been at a great feast
of languages and stolen the scraps.

COSTARD, [aside to Boy]  O, they have lived long on the
almsbasket of words. I marvel thy master hath not
eaten thee for a word, for thou art not so long by the
head as honorificabilitudinitatibus. Thou art easier
swallowed than a flapdragon.

BOY, [aside to Costard]  Peace, the peal begins.

ARMADO, [to Holofernes]  Monsieur, are you not
lettered?

BOY  Yes, yes, he teaches boys the hornbook.--What is
a, b spelled backward, with the horn on his head?

HOLOFERNES  Ba, pueritia, with a horn added.

BOY  Ba, most silly sheep, with a horn.--You hear his
learning.

HOLOFERNES  Quis, quis, thou consonant?

BOY  The last of the five vowels, if you repeat them; or
the fifth, if I.

HOLOFERNES  I will repeat them: a, e, i--

BOY  The sheep. The other two concludes it: o, u.

ARMADO  Now by the salt wave of the Mediterraneum,
a sweet touch, a quick venue of wit! Snip, snap,
quick and home. It rejoiceth my intellect. True
wit.

BOY  Offered by a child to an old man--which is
wit-old.

HOLOFERNES  What is the figure? What is the figure?

BOY  Horns.

HOLOFERNES  Thou disputes like an infant. Go whip thy
gig.

BOY  Lend me your horn to make one, and I will whip
about your infamy--unum cita--a gig of a cuckold's
horn.

COSTARD  An I had but one penny in the world, thou
shouldst have it to buy gingerbread! Hold, there is
the very remuneration I had of thy master, thou
halfpenny purse of wit, thou pigeon egg of discretion.
[He gives him money.] O, an the heavens were
so pleased that thou wert but my bastard, what a
joyful father wouldest thou make me! Go to, thou
hast it ad dunghill, at the fingers' ends, as they say.

HOLOFERNES  Oh, I smell false Latin! Dunghill for
unguem.

ARMADO  Arts-man, preambulate. We will be singuled
from the barbarous. Do you not educate youth at
the charge-house on the top of the mountain?

HOLOFERNES  Or mons, the hill.

ARMADO  At your sweet pleasure, for the mountain.

HOLOFERNES  I do, sans question.

ARMADO  Sir, it is the King's most sweet pleasure and
affection to congratulate the Princess at her pavilion
in the posteriors of this day, which the rude
multitude call the afternoon.

HOLOFERNES  "The posterior of the day," most generous
sir, is liable, congruent, and measurable for
"the afternoon"; the word is well culled, chose,
sweet, and apt, I do assure you, sir, I do assure.

ARMADO  Sir, the King is a noble gentleman, and my
familiar, I do assure you, very good friend. For
what is inward between us, let it pass. I do beseech
thee, remember thy courtesy; I beseech thee apparel
thy head. And among other important and most
serious designs, and of great import indeed, too--
but let that pass; for I must tell thee, it will please his
Grace, by the world, sometimes to lean upon my
poor shoulder and with his royal finger thus dally
with my excrement, with my mustachio--but,
sweetheart, let that pass. By the world, I recount no
fable! Some certain special honors it pleaseth his
Greatness to impart to Armado, a soldier, a man of
travel, that hath seen the world--but let that pass.
The very all of all is--but sweetheart, I do implore
secrecy--that the King would have me present the
Princess, sweet chuck, with some delightful ostentation,
or show, or pageant, or antic, or firework.
Now, understanding that the curate and your sweet
self are good at such eruptions and sudden breaking
out of mirth, as it were, I have acquainted you
withal to the end to crave your assistance.

HOLOFERNES  Sir, you shall present before her the Nine
Worthies.--Sir Nathaniel, as concerning some
entertainment of time, some show in the posterior
of this day, to be rendered by our assistance, the
King's command, and this most gallant, illustrate,
and learned gentleman, before the Princess--I say,
none so fit as to present the Nine Worthies.

NATHANIEL  Where will you find men worthy enough to
present them?

HOLOFERNES  Joshua, yourself; myself; and this gallant
gentleman, Judas Maccabaeus. This swain, because
of his great limb or joint, shall pass Pompey
the Great; the page, Hercules--

ARMADO  Pardon, sir--error. He is not quantity
enough for that Worthy's thumb; he is not so big as
the end of his club!

HOLOFERNES  Shall I have audience? He shall present
Hercules in minority. His enter and exit shall be
strangling a snake; and I will have an apology for
that purpose.

BOY  An excellent device. So, if any of the audience
hiss, you may cry "Well done, Hercules, now thou
crushest the snake." That is the way to make an
offense gracious, though few have the grace to do it.

ARMADO  For the rest of the Worthies?

HOLOFERNES  I will play three myself.

BOY  Thrice-worthy gentleman!

ARMADO, [to Holofernes]  Shall I tell you a thing?

HOLOFERNES  We attend.

ARMADO  We will have, if this fadge not, an antic. I
beseech you, follow.

HOLOFERNES  Via, goodman Dull. Thou hast spoken no
word all this while.

DULL  Nor understood none neither, sir.

HOLOFERNES  Allons! We will employ thee.

DULL  I'll make one in a dance, or so; or I will play on
the tabor to the Worthies and let them dance the
hay.

HOLOFERNES  Most dull, honest Dull. To our sport!
Away.
[They exit.]

Scene 2
=======
[Enter the Ladies (the Princess, Rosaline,
Katherine, and Maria.)]


PRINCESS
Sweethearts, we shall be rich ere we depart,
If fairings come thus plentifully in.
A lady walled about with diamonds!
Look you what I have from the loving king.
[She shows a jewel.]

ROSALINE
Madam, came nothing else along with that?

PRINCESS
Nothing but this? Yes, as much love in rhyme
As would be crammed up in a sheet of paper
Writ o' both sides the leaf, margent and all,
That he was fain to seal on Cupid's name.

ROSALINE
That was the way to make his godhead wax,
For he hath been five thousand year a boy.

KATHERINE
Ay, and a shrewd unhappy gallows, too.

ROSALINE
You'll ne'er be friends with him. He killed your
sister.

KATHERINE
He made her melancholy, sad, and heavy,
And so she died. Had she been light like you,
Of such a merry, nimble, stirring spirit,
She might ha' been a grandam ere she died.
And so may you, for a light heart lives long.

ROSALINE
What's your dark meaning, mouse, of this light
word?

KATHERINE
A light condition in a beauty dark.

ROSALINE
We need more light to find your meaning out.

KATHERINE
You'll mar the light by taking it in snuff;
Therefore I'll darkly end the argument.

ROSALINE
Look what you do, you do it still i' th' dark.

KATHERINE
So do not you, for you are a light wench.

ROSALINE
Indeed, I weigh not you, and therefore light.

KATHERINE
You weigh me not? O, that's you care not for me.

ROSALINE
Great reason: for past care is still past cure.

PRINCESS
Well bandied both; a set of wit well played.
But, Rosaline, you have a favor too.
Who sent it? And what is it?

ROSALINE  I would you knew.
An if my face were but as fair as yours,
My favor were as great. Be witness this.
[She shows a gift.]
Nay, I have verses too, I thank Berowne;
The numbers true; and were the numb'ring too,
I were the fairest goddess on the ground.
I am compared to twenty thousand fairs.
O, he hath drawn my picture in his letter.

PRINCESS  Anything like?

ROSALINE
Much in the letters, nothing in the praise.

PRINCESS
Beauteous as ink: a good conclusion.

KATHERINE
Fair as a text B in a copybook.

ROSALINE
Ware pencils, ho! Let me not die your debtor,
My red dominical, my golden letter.
O, that your face were not so full of O's!

PRINCESS
A pox of that jest! And I beshrew all shrows.
But, Katherine, what was sent to you
From fair Dumaine?

KATHERINE
Madam, this glove.	[She shows the glove.]

PRINCESS  Did he not send you twain?

KATHERINE  Yes, madam, and moreover,
Some thousand verses of a faithful lover,
A huge translation of hypocrisy,
Vilely compiled, profound simplicity.

MARIA
This, and these pearls, to me sent Longaville.
[She shows a paper and pearls.]
The letter is too long by half a mile.

PRINCESS
I think no less. Dost thou not wish in heart
The chain were longer and the letter short?

MARIA
Ay, or I would these hands might never part.

PRINCESS
We are wise girls to mock our lovers so.

ROSALINE
They are worse fools to purchase mocking so.
That same Berowne I'll torture ere I go.
O, that I knew he were but in by th' week,
How I would make him fawn, and beg, and seek,
And wait the season, and observe the times,
And spend his prodigal wits in bootless rhymes,
And shape his service wholly to my hests,
And make him proud to make me proud that jests!
So pair-taunt-like would I o'ersway his state,
That he should be my fool, and I his fate.

PRINCESS
None are so surely caught, when they are catched,
As wit turned fool. Folly in wisdom hatched
Hath wisdom's warrant and the help of school,
And wit's own grace to grace a learned fool.

ROSALINE
The blood of youth burns not with such excess
As gravity's revolt to wantonness.

MARIA
Folly in fools bears not so strong a note
As fool'ry in the wise, when wit doth dote,
Since all the power thereof it doth apply
To prove, by wit, worth in simplicity.

[Enter Boyet.]


PRINCESS
Here comes Boyet, and mirth is in his face.

BOYET
O, I am stabbed with laughter. Where's her Grace?

PRINCESS
Thy news, Boyet?

BOYET  Prepare, madam, prepare.
Arm, wenches, arm. Encounters mounted are
Against your peace. Love doth approach, disguised,
Armed in arguments. You'll be surprised.
Muster your wits, stand in your own defense,
Or hide your heads like cowards, and fly hence.

PRINCESS
Saint Denis to Saint Cupid! What are they
That charge their breath against us? Say, scout, say.

BOYET
Under the cool shade of a sycamore,
I thought to close mine eyes some half an hour.
When, lo, to interrupt my purposed rest,
Toward that shade I might behold addressed
The King and his companions. Warily
I stole into a neighbor thicket by,
And overheard what you shall overhear:
That, by and by, disguised, they will be here.
Their herald is a pretty knavish page
That well by heart hath conned his embassage.
Action and accent did they teach him there:
"Thus must thou speak," and "thus thy body bear."
And ever and anon they made a doubt
Presence majestical would put him out;
"For," quoth the King, "an angel shalt thou see;
Yet fear not thou, but speak audaciously."
The boy replied "An angel is not evil.
I should have feared her had she been a devil."
With that, all laughed and clapped him on the
shoulder,
Making the bold wag by their praises bolder.
One rubbed his elbow thus, and fleered, and swore
A better speech was never spoke before.
Another with his finger and his thumb,
Cried "Via! We will do 't, come what will come."
The third he capered and cried "All goes well!"
The fourth turned on the toe, and down he fell.
With that, they all did tumble on the ground
With such a zealous laughter so profound
That in this spleen ridiculous appears,
To check their folly, passion's solemn tears.

PRINCESS
But what, but what? Come they to visit us?

BOYET
They do, they do; and are appareled thus,
Like Muscovites, or Russians, as I guess.
Their purpose is to parley, to court, and dance,
And every one his love-feat will advance
Unto his several mistress--which they'll know
By favors several which they did bestow.

PRINCESS
And will they so? The gallants shall be tasked,
For, ladies, we will every one be masked,
And not a man of them shall have the grace,
Despite of suit, to see a lady's face.
Hold, Rosaline, this favor thou shalt wear,
And then the King will court thee for his dear.
Hold, take thou this, my sweet, and give me thine.
So shall Berowne take me for Rosaline.
[Princess and Rosaline exchange favors.]
And change you favors too. So shall your loves
Woo contrary, deceived by these removes.
[Katherine and Maria exchange favors.]

ROSALINE
Come on, then, wear the favors most in sight.

KATHERINE, [to Princess]
But in this changing, what is your intent?

PRINCESS
The effect of my intent is to cross theirs.
They do it but in mockery merriment,
And mock for mock is only my intent.
Their several counsels they unbosom shall
To loves mistook, and so be mocked withal
Upon the next occasion that we meet,
With visages displayed, to talk and greet.

ROSALINE
But shall we dance, if they desire us to 't?

PRINCESS
No, to the death we will not move a foot,
Nor to their penned speech render we no grace,
But while 'tis spoke each turn away her face.

BOYET
Why, that contempt will kill the speaker's heart,
And quite divorce his memory from his part.

PRINCESS
Therefore I do it, and I make no doubt
The rest will ne'er come in if he be out.
There's no such sport as sport by sport o'erthrown,
To make theirs ours and ours none but our own.
So shall we stay, mocking intended game,
And they, well mocked, depart away with shame.
[Sound trumpet, within.]

BOYET
The trumpet sounds. Be masked; the maskers come.
[The Ladies mask.]

[Enter Blackamoors with music, the Boy with a speech,
the King, Berowne, and the rest of the Lords disguised.]


BOY
All hail, the richest beauties on the Earth!

BOYET
Beauties no richer than rich taffeta.

BOY
A holy parcel of the fairest dames
[The Ladies turn their backs to him.]
That ever turned their--backs--to mortal views.

BEROWNE  Their eyes, villain, their eyes!

BOY
That ever turned their eyes to mortal views.
Out--

BOYET  True; out indeed.

BOY
Out of your favors, heavenly spirits, vouchsafe
Not to behold--

BEROWNE  Once to behold, rogue!

BOY
Once to behold with your sun-beamed eyes--
With your sun-beamed eyes--

BOYET
They will not answer to that epithet.
You were best call it "daughter-beamed eyes."

BOY
They do not mark me, and that brings me out.

BEROWNE
Is this your perfectness? Begone, you rogue!
[Boy exits.]

ROSALINE, [speaking as the Princess]
What would these strangers? Know their minds,
Boyet.
If they do speak our language, 'tis our will
That some plain man recount their purposes.
Know what they would.

BOYET  What would you with the
Princess?

BEROWNE
Nothing but peace and gentle visitation.

ROSALINE  What would they, say they?

BOYET
Nothing but peace and gentle visitation.

ROSALINE
Why, that they have, and bid them so be gone.

BOYET
She says you have it, and you may be gone.

KING
Say to her we have measured many miles
To tread a measure with her on this grass.

BOYET
They say that they have measured many a mile
To tread a measure with you on this grass.

ROSALINE
It is not so. Ask them how many inches
Is in one mile. If they have measured many,
The measure then of one is eas'ly told.

BOYET
If to come hither you have measured miles,
And many miles, the Princess bids you tell
How many inches doth fill up one mile.

BEROWNE
Tell her we measure them by weary steps.

BOYET
She hears herself.

ROSALINE  How many weary steps
Of many weary miles you have o'ergone
Are numbered in the travel of one mile?

BEROWNE
We number nothing that we spend for you.
Our duty is so rich, so infinite,
That we may do it still without account.
Vouchsafe to show the sunshine of your face
That we, like savages, may worship it.

ROSALINE
My face is but a moon, and clouded too.

KING
Blessed are clouds, to do as such clouds do!
Vouchsafe, bright moon, and these thy stars, to
shine,
Those clouds removed, upon our watery eyne.

ROSALINE
O vain petitioner, beg a greater matter!
Thou now requests but moonshine in the water.

KING
Then in our measure do but vouchsafe one change.
Thou bidd'st me beg; this begging is not strange.

ROSALINE
Play music, then. Nay, you must do it soon.
[Music begins.]
Not yet? No dance! Thus change I like the moon.

KING
Will you not dance? How come you thus estranged?

ROSALINE
You took the moon at full, but now she's changed.

KING
Yet still she is the moon, and I the man.
The music plays. Vouchsafe some motion to it.

ROSALINE
Our ears vouchsafe it.

KING  But your legs should do it.

ROSALINE
Since you are strangers and come here by chance,
We'll not be nice. Take hands. We will not dance.
[She offers her hand.]

KING
Why take we hands then?

ROSALINE  Only to part friends.--
Curtsy, sweethearts--and so the measure ends.

KING
More measure of this measure! Be not nice.

ROSALINE
We can afford no more at such a price.

KING
Prize you yourselves. What buys your company?

ROSALINE
Your absence only.

KING  That can never be.

ROSALINE
Then cannot we be bought. And so adieu--
Twice to your visor, and half once to you.

KING
If you deny to dance, let's hold more chat.

ROSALINE
In private, then.

KING  I am best pleased with that.
[They move aside.]

BEROWNE, [to the Princess]
White-handed mistress, one sweet word with thee.

PRINCESS, [speaking as Rosaline]
Honey, and milk, and sugar--there is three.

BEROWNE
Nay then, two treys, an if you grow so nice,
Metheglin, wort, and malmsey. Well run, dice!
There's half a dozen sweets.

PRINCESS  Seventh sweet, adieu.
Since you can cog, I'll play no more with you.

BEROWNE
One word in secret.

PRINCESS  Let it not be sweet.

BEROWNE
Thou grievest my gall.

PRINCESS  Gall! Bitter.

BEROWNE  Therefore meet.
[They move aside.]

DUMAINE, [to Maria]
Will you vouchsafe with me to change a word?

MARIA, [speaking as Katherine]
Name it.

DUMAINE  Fair lady--

MARIA  Say you so? Fair lord!
Take that for your "fair lady."

DUMAINE  Please it you
As much in private, and I'll bid adieu.
[They move aside.]

KATHERINE, [speaking as Maria]
What, was your vizard made without a tongue?

LONGAVILLE
I know the reason, lady, why you ask.

KATHERINE
O, for your reason! Quickly, sir, I long.

LONGAVILLE
You have a double tongue within your mask,
And would afford my speechless vizard half.

KATHERINE
Veal, quoth the Dutchman. Is not veal a calf?

LONGAVILLE
A calf, fair lady?

KATHERINE  No, a fair Lord Calf.

LONGAVILLE
Let's part the word.

KATHERINE  No, I'll not be your half.
Take all and wean it. It may prove an ox.

LONGAVILLE
Look how you butt yourself in these sharp mocks.
Will you give horns, chaste lady? Do not so.

KATHERINE
Then die a calf before your horns do grow.

LONGAVILLE
One word in private with you ere I die.

KATHERINE
Bleat softly, then. The butcher hears you cry.
[They move aside.]

BOYET
The tongues of mocking wenches are as keen
   As is the razor's edge invisible,
Cutting a smaller hair than may be seen;
   Above the sense of sense, so sensible
Seemeth their conference. Their conceits have
wings
Fleeter than arrows, bullets, wind, thought, swifter
things.

ROSALINE
Not one word more, my maids. Break off, break off!
[The Ladies move away from the Lords.]

BEROWNE
By heaven, all dry-beaten with pure scoff!

KING
Farewell, mad wenches. You have simple wits.
[King, Lords, and Blackamoors exit.]
[The Ladies unmask.]

PRINCESS
Twenty adieus, my frozen Muskovits.--
Are these the breed of wits so wondered at?

BOYET
   Tapers they are, with your sweet breaths puffed
   out.

ROSALINE
Well-liking wits they have; gross, gross; fat, fat.

PRINCESS
   O poverty in wit, kingly-poor flout!
Will they not, think you, hang themselves tonight?
   Or ever but in vizards show their faces?
This pert Berowne was out of count'nance quite.

ROSALINE
   They were all in lamentable cases.
The King was weeping ripe for a good word.

PRINCESS
   Berowne did swear himself out of all suit.

MARIA
Dumaine was at my service, and his sword.
   "No point," quoth I. My servant straight was
   mute.

KATHERINE
Lord Longaville said I came o'er his heart.
   And trow you what he called me?

PRINCESS     Qualm, perhaps.

KATHERINE
Yes, in good faith.

PRINCESS  Go, sickness as thou art!

ROSALINE
   Well, better wits have worn plain statute-caps.
But will you hear? The King is my love sworn.

PRINCESS
   And quick Berowne hath plighted faith to me.

KATHERINE
And Longaville was for my service born.

MARIA
   Dumaine is mine as sure as bark on tree.

BOYET
Madam, and pretty mistresses, give ear.
Immediately they will again be here
In their own shapes, for it can never be
They will digest this harsh indignity.

PRINCESS
Will they return?

BOYET  They will, they will, God knows,
And leap for joy, though they are lame with blows.
Therefore change favors, and when they repair,
Blow like sweet roses in this summer air.

PRINCESS
How "blow"? How "blow"? Speak to be understood.

BOYET
Fair ladies masked are roses in their bud.
Dismasked, their damask sweet commixture shown,
Are angels vailing clouds, or roses blown.

PRINCESS
Avaunt, perplexity!--What shall we do
If they return in their own shapes to woo?

ROSALINE
Good madam, if by me you'll be advised,
Let's mock them still, as well known as disguised.
Let us complain to them what fools were here,
Disguised like Muscovites in shapeless gear,
And wonder what they were, and to what end
Their shallow shows and prologue vilely penned,
And their rough carriage so ridiculous,
Should be presented at our tent to us.

BOYET
Ladies, withdraw. The gallants are at hand.

PRINCESS
Whip to our tents, as roes runs o'er land.
[The Princess and the Ladies exit.]

[Enter the King and the rest, as themselves.]


KING, [to Boyet]
Fair sir, God save you. Where's the Princess?

BOYET
Gone to her tent. Please it your Majesty
Command me any service to her thither?

KING
That she vouchsafe me audience for one word.

BOYET
I will, and so will she, I know, my lord.	[He exits.]

BEROWNE
This fellow pecks up wit as pigeons peas,
And utters it again when God doth please.
He is wit's peddler, and retails his wares
At wakes and wassails, meetings, markets, fairs.
And we that sell by gross, the Lord doth know,
Have not the grace to grace it with such show.
This gallant pins the wenches on his sleeve.
Had he been Adam, he had tempted Eve.
He can carve too, and lisp. Why, this is he
That kissed his hand away in courtesy.
This is the ape of form, Monsieur the Nice,
That, when he plays at tables, chides the dice
In honorable terms. Nay, he can sing
A mean most meanly; and in ushering
Mend him who can. The ladies call him sweet.
The stairs, as he treads on them, kiss his feet.
This is the flower that smiles on everyone
To show his teeth as white as whale's bone;
And consciences that will not die in debt
Pay him the due of "honey-tongued Boyet."

KING
A blister on his sweet tongue, with my heart,
That put Armado's page out of his part!

[Enter the Ladies, with Boyet.]


BEROWNE
See where it comes! Behavior, what wert thou
Till this madman showed thee? And what art thou
now?

KING, [to Princess]
All hail, sweet madam, and fair time of day.

PRINCESS
   "Fair" in "all hail" is foul, as I conceive.

KING
Construe my speeches better, if you may.

PRINCESS
   Then wish me better. I will give you leave.

KING
We came to visit you, and purpose now
   To lead you to our court. Vouchsafe it, then.

PRINCESS
This field shall hold me, and so hold your vow.
   Nor God nor I delights in perjured men.

KING
Rebuke me not for that which you provoke.
   The virtue of your eye must break my oath.

PRINCESS
You nickname virtue; "vice" you should have spoke,
   For virtue's office never breaks men's troth.
Now by my maiden honor, yet as pure
   As the unsullied lily, I protest,
A world of torments though I should endure,
   I would not yield to be your house's guest,
So much I hate a breaking cause to be
Of heavenly oaths vowed with integrity.

KING
O, you have lived in desolation here,
   Unseen, unvisited, much to our shame.

PRINCESS
Not so, my lord. It is not so, I swear.
   We have had pastimes here and pleasant game.
A mess of Russians left us but of late.

KING
   How, madam? Russians?

PRINCESS     Ay, in truth, my lord.
Trim gallants, full of courtship and of state.

ROSALINE
   Madam, speak true.--It is not so, my lord.
My lady, to the manner of the days,
In courtesy gives undeserving praise.
We four indeed confronted were with four
In Russian habit. Here they stayed an hour
And talked apace; and in that hour, my lord,
They did not bless us with one happy word.
I dare not call them fools; but this I think:
When they are thirsty, fools would fain have drink.

BEROWNE
This jest is dry to me. Gentle sweet,
Your wits makes wise things foolish. When we greet,
With eyes' best seeing, heaven's fiery eye,
By light we lose light. Your capacity
Is of that nature that to your huge store
Wise things seem foolish and rich things but poor.

ROSALINE
This proves you wise and rich, for in my eye--

BEROWNE
I am a fool, and full of poverty.

ROSALINE
But that you take what doth to you belong,
It were a fault to snatch words from my tongue.

BEROWNE
O, I am yours, and all that I possess!

ROSALINE
All the fool mine?

BEROWNE  I cannot give you less.

ROSALINE
Which of the vizards was it that you wore?

BEROWNE
Where? When? What vizard? Why demand you this?

ROSALINE
There; then; that vizard; that superfluous case
That hid the worse and showed the better face.

KING, [aside to Dumaine]
We were descried. They'll mock us now downright.

DUMAINE, [aside to King]
Let us confess and turn it to a jest.

PRINCESS, [to King]
Amazed, my lord? Why looks your Highness sad?

ROSALINE
Help, hold his brows! He'll swoon!--Why look you
pale?
Seasick, I think, coming from Muscovy.

BEROWNE
Thus pour the stars down plagues for perjury.
   Can any face of brass hold longer out?
Here stand I, lady. Dart thy skill at me.
   Bruise me with scorn, confound me with a flout.
Thrust thy sharp wit quite through my ignorance.
   Cut me to pieces with thy keen conceit,
And I will wish thee nevermore to dance,
   Nor nevermore in Russian habit wait.
O, never will I trust to speeches penned,
   Nor to the motion of a schoolboy's tongue,
Nor never come in vizard to my friend,
   Nor woo in rhyme like a blind harper's song.
Taffeta phrases, silken terms precise,
   Three-piled hyperboles, spruce affectation,
Figures pedantical--these summer flies
   Have blown me full of maggot ostentation.
I do forswear them, and I here protest
   By this white glove--how white the hand, God
   knows!--
Henceforth my wooing mind shall be expressed
   In russet yeas and honest kersey noes.
And to begin: Wench, so God help me, law,
My love to thee is sound, sans crack or flaw.

ROSALINE
Sans "sans," I pray you.

BEROWNE  Yet I have a trick
Of the old rage. Bear with me, I am sick;
I'll leave it by degrees. Soft, let us see:
Write "Lord have mercy on us" on those three.
They are infected; in their hearts it lies.
They have the plague, and caught it of your eyes.
These lords are visited. You are not free,
For the Lord's tokens on you do I see.

PRINCESS
No, they are free that gave these tokens to us.

BEROWNE
Our states are forfeit. Seek not to undo us.

ROSALINE
It is not so, for how can this be true,
That you stand forfeit, being those that sue?

BEROWNE
Peace, for I will not have to do with you.

ROSALINE
Nor shall not, if I do as I intend.

BEROWNE, [to King, Longaville, and Dumaine]
Speak for yourselves. My wit is at an end.

KING, [to Princess]
Teach us, sweet madam, for our rude transgression
Some fair excuse.

PRINCESS  The fairest is confession.
Were not you here but even now, disguised?

KING
Madam, I was.

PRINCESS  And were you well advised?

KING
I was, fair madam.

PRINCESS  When you then were here,
What did you whisper in your lady's ear?

KING
That more than all the world I did respect her.

PRINCESS
When she shall challenge this, you will reject her.

KING
Upon mine honor, no.

PRINCESS  Peace, peace, forbear!
Your oath once broke, you force not to forswear.

KING
Despise me when I break this oath of mine.

PRINCESS
I will, and therefore keep it.--Rosaline,
What did the Russian whisper in your ear?

ROSALINE
Madam, he swore that he did hold me dear
As precious eyesight, and did value me
Above this world, adding thereto moreover
That he would wed me or else die my lover.

PRINCESS
God give thee joy of him! The noble lord
Most honorably doth uphold his word.

KING
What mean you, madam? By my life, my troth,
I never swore this lady such an oath.

ROSALINE
By heaven, you did! And to confirm it plain,
You gave me this. [She shows a token.] But take it,
sir, again.

KING
My faith and this the Princess I did give.
I knew her by this jewel on her sleeve.

PRINCESS
Pardon me, sir. This jewel did she wear.
[She points to Rosaline.]
And Lord Berowne, I thank him, is my dear.
[To Berowne.] What, will you have me, or your pearl
again?	[She shows the token.]

BEROWNE
Neither of either. I remit both twain.
I see the trick on 't. Here was a consent,
Knowing aforehand of our merriment,
To dash it like a Christmas comedy.
Some carry-tale, some please-man, some slight
zany,
Some mumble-news, some trencher-knight, some
Dick,
That smiles his cheek in years and knows the trick
To make my lady laugh when she's disposed,
Told our intents before; which once disclosed,
The ladies did change favors; and then we,
Following the signs, wooed but the sign of she.
Now, to our perjury to add more terror,
We are again forsworn in will and error.
Much upon this 'tis. [To Boyet.] And might not you
Forestall our sport, to make us thus untrue?
Do not you know my lady's foot by th' squier?
   And laugh upon the apple of her eye?
And stand between her back, sir, and the fire,
   Holding a trencher, jesting merrily?
You put our page out. Go, you are allowed.
Die when you will, a smock shall be your shroud.
You leer upon me, do you? There's an eye
Wounds like a leaden sword.

BOYET  Full merrily
Hath this brave manage, this career been run.

BEROWNE
Lo, he is tilting straight! Peace, I have done.

[Enter Clown Costard.]

Welcome, pure wit. Thou part'st a fair fray.

COSTARD  O Lord, sir, they would know
Whether the three Worthies shall come in or no.

BEROWNE
What, are there but three?

COSTARD  No, sir; but it is vara fine,
For every one pursents three.

BEROWNE  And three times thrice
is nine.

COSTARD
Not so, sir, under correction, sir, I hope it is not so.
You cannot beg us, sir, I can assure you, sir; we
know what we know.
I hope, sir, three times thrice, sir--

BEROWNE  Is not nine?

COSTARD  Under correction, sir, we know whereuntil it
doth amount.

BEROWNE
By Jove, I always took three threes for nine.

COSTARD  O Lord, sir, it were pity you should get your
living by reckoning, sir.

BEROWNE  How much is it?

COSTARD  O Lord, sir, the parties themselves, the actors,
sir, will show whereuntil it doth amount. For
mine own part, I am, as they say, but to parfect one
man in one poor man--Pompion the Great, sir.

BEROWNE  Art thou one of the Worthies?

COSTARD  It pleased them to think me worthy of Pompey
the Great. For mine own part, I know not the
degree of the Worthy, but I am to stand for him.

BEROWNE  Go bid them prepare.

COSTARD
We will turn it finely off, sir. We will take some
care.	[He exits.]

KING
Berowne, they will shame us. Let them not
approach.

BEROWNE
We are shame-proof, my lord; and 'tis some policy
To have one show worse than the King's and his
company.

KING  I say they shall not come.

PRINCESS
Nay, my good lord, let me o'errule you now.
That sport best pleases that doth least know how,
Where zeal strives to content, and the contents
Dies in the zeal of that which it presents.
Their form confounded makes most form in mirth,
When great things laboring perish in their birth.

BEROWNE
A right description of our sport, my lord.

[Enter Braggart Armado.]


ARMADO, [to King]  Anointed, I implore so much expense
of thy royal sweet breath as will utter a brace
of words.	[Armado and King step aside, and
Armado gives King a paper.]

PRINCESS  Doth this man serve God?

BEROWNE  Why ask you?

PRINCESS
He speaks not like a man of God his making.

ARMADO, [to King]  That is all one, my fair sweet honey
monarch, for, I protest, the schoolmaster is exceeding
fantastical, too, too vain, too, too vain. But
we will put it, as they say, to fortuna de la guerra.--I
wish you the peace of mind, most royal
couplement!	[He exits.]

KING, [reading the paper]  Here is like to be a good
presence of Worthies. He presents Hector of Troy,
the swain Pompey the Great, the parish curate
Alexander, Armado's page Hercules, the pedant
Judas Maccabaeus.
And if these four Worthies in their first show thrive,
These four will change habits and present the other
five.

BEROWNE  There is five in the first show.

KING  You are deceived. 'Tis not so.

BEROWNE  The pedant, the braggart, the hedge
priest, the fool, and the boy.
Abate throw at novum, and the whole world again
Cannot pick out five such, take each one in his vein.

KING
The ship is under sail, and here she comes amain.

[Enter Costard as Pompey.]


COSTARD
I Pompey am--

BEROWNE  You lie; you are not he.

COSTARD
I Pompey am--

BOYET  With leopard's head on knee.

BEROWNE
Well said, old mocker. I must needs be friends with
thee.

COSTARD
I Pompey am, Pompey, surnamed the Big--

DUMAINE  "The Great."

COSTARD
It is "Great," sir.--Pompey, surnamed the
Great,
That oft in field, with targe and shield, did make my
foe to sweat.
And traveling along this coast, I here am come by
chance,
And lay my arms before the legs of this sweet lass of
France.
[He places his weapons at the feet of the Princess.]
If your Ladyship would say "Thanks, Pompey," I
had done.

PRINCESS  Great thanks, great Pompey.

COSTARD  'Tis not so much worth, but I hope I was
perfect. I made a little fault in "Great."

BEROWNE  My hat to a halfpenny, Pompey proves the
best Worthy.	[Costard stands aside.]

[Enter Curate Nathaniel for Alexander.]


NATHANIEL
When in the world I lived, I was the world's
commander.
By east, west, north, and south, I spread my
conquering might.
My scutcheon plain declares that I am Alisander--

BOYET
Your nose says no, you are not, for it stands too
right.

BEROWNE, [to Boyet]
Your nose smells "no" in this, most tender-smelling
knight.

PRINCESS
The conqueror is dismayed.--Proceed, good
Alexander.

NATHANIEL
When in the world I lived, I was the world's
commander--

BOYET
Most true; 'tis right. You were so, Alisander.

BEROWNE, [to Costard]  Pompey the Great--

COSTARD  Your servant, and Costard.

BEROWNE  Take away the conqueror. Take away
Alisander.

COSTARD, [to Nathaniel]  O sir, you have overthrown
Alisander the Conqueror. You will be scraped out of
the painted cloth for this. Your lion, that holds his
polax sitting on a close-stool, will be given to Ajax.
He will be the ninth Worthy. A conqueror, and
afeard to speak? Run away for shame, Alisander.
[Nathaniel exits.]
There, an 't shall please you, a foolish mild man, an
honest man, look you, and soon dashed. He is a
marvelous good neighbor, faith, and a very good
bowler. But, for Alisander--alas, you see how 'tis--
a little o'erparted. But there are Worthies a-coming
will speak their mind in some other sort.

[Enter Pedant Holofernes for Judas, and the Boy
for Hercules.]


PRINCESS, [to Costard]  Stand aside, good Pompey.

HOLOFERNES
Great Hercules is presented by this imp,
   Whose club killed Cerberus, that three-headed canus,
And when he was a babe, a child, a shrimp,
   Thus did he strangle serpents in his manus.
Quoniam he seemeth in minority,
Ergo I come with this apology.
[To Boy.] Keep some state in thy exit, and vanish.
[Boy steps aside.]

HOLOFERNES
Judas I am--

DUMAINE  A Judas!

HOLOFERNES  Not Iscariot, sir.
Judas I am, yclept Maccabaeus.

DUMAINE  Judas Maccabaeus clipped is plain Judas.

BEROWNE  A kissing traitor.--How art thou proved
Judas?

HOLOFERNES
Judas I am--

DUMAINE  The more shame for you, Judas.

HOLOFERNES  What mean you, sir?

BOYET  To make Judas hang himself.

HOLOFERNES  Begin, sir, you are my elder.

BEROWNE  Well followed. Judas was hanged on an
elder.

HOLOFERNES  I will not be put out of countenance.

BEROWNE  Because thou hast no face.

HOLOFERNES  What is this?	[He points to his own face.]

BOYET  A cittern-head.

DUMAINE  The head of a bodkin.

BEROWNE  A death's face in a ring.

LONGAVILLE  The face of an old Roman coin, scarce
seen.

BOYET  The pommel of Caesar's falchion.

DUMAINE  The carved-bone face on a flask.

BEROWNE  Saint George's half-cheek in a brooch.

DUMAINE  Ay, and in a brooch of lead.

BEROWNE  Ay, and worn in the cap of a tooth-drawer.
And now forward, for we have put thee in
countenance.

HOLOFERNES  You have put me out of countenance.

BEROWNE  False. We have given thee faces.

HOLOFERNES  But you have outfaced them all.

BEROWNE
An thou wert a lion, we would do so.

BOYET
Therefore, as he is an ass, let him go.--
And so adieu, sweet Jude. Nay, why dost thou stay?

DUMAINE  For the latter end of his name.

BEROWNE
For the "ass" to the "Jude"? Give it him.--Jud-as,
away!

HOLOFERNES
This is not generous, not gentle, not humble.

BOYET
A light for Monsieur Judas! It grows dark; he may
stumble.	[Holofernes exits.]

PRINCESS
Alas, poor Maccabaeus, how hath he been baited!

[Enter Braggart Armado as Hector.]


BEROWNE  Hide thy head, Achilles. Here comes Hector
in arms.

DUMAINE  Though my mocks come home by me, I will
now be merry.

KING  Hector was but a Troyan in respect of this.

BOYET  But is this Hector?

KING  I think Hector was not so clean-timbered.

LONGAVILLE  His leg is too big for Hector's.

DUMAINE  More calf, certain.

BOYET  No, he is best endued in the small.

BEROWNE  This cannot be Hector.

DUMAINE  He's a god or a painter, for he makes faces.

ARMADO
The armipotent Mars, of lances the almighty,
   Gave Hector a gift--

DUMAINE  A gilt nutmeg.

BEROWNE  A lemon.

LONGAVILLE  Stuck with cloves.

DUMAINE  No, cloven.

ARMADO  Peace!
The armipotent Mars, of lances the almighty,
   Gave Hector a gift, the heir of Ilion,
A man so breathed, that certain he would fight, yea,
   From morn till night, out of his pavilion.
I am that flower--

DUMAINE  That mint.

LONGAVILLE  That columbine.

ARMADO  Sweet Lord Longaville, rein thy tongue.

LONGAVILLE  I must rather give it the rein, for it runs
against Hector.

DUMAINE  Ay, and Hector's a greyhound.

ARMADO  The sweet warman is dead and rotten. Sweet
chucks, beat not the bones of the buried. When he
breathed, he was a man. But I will forward with my
device. [To Princess.] Sweet royalty, bestow on me
the sense of hearing.
[Berowne steps forth.]

PRINCESS
Speak, brave Hector. We are much delighted.

ARMADO  I do adore thy sweet Grace's slipper.

BOYET  Loves her by the foot.

DUMAINE  He may not by the yard.

ARMADO
This Hector far surmounted Hannibal.
The party is gone--

COSTARD  Fellow Hector, she is gone; she is two
months on her way.

ARMADO  What meanest thou?

COSTARD  Faith, unless you play the honest Troyan, the
poor wench is cast away. She's quick; the child
brags in her belly already. 'Tis yours.

ARMADO  Dost thou infamonize me among potentates?
Thou shalt die!

COSTARD  Then shall Hector be whipped for Jaquenetta,
that is quick by him, and hanged for Pompey,
that is dead by him.

DUMAINE  Most rare Pompey!

BOYET  Renowned Pompey!

BEROWNE  Greater than "Great"! Great, great, great
Pompey. Pompey the Huge!

DUMAINE  Hector trembles.

BEROWNE  Pompey is moved. More Ates, more Ates!
Stir them on, stir them on.

DUMAINE  Hector will challenge him.

BEROWNE  Ay, if he have no more man's blood in his
belly than will sup a flea.

ARMADO, [to Costard]  By the North Pole, I do challenge
thee!

COSTARD  I will not fight with a pole like a northern
man! I'll slash. I'll do it by the sword.--I bepray
you, let me borrow my arms again.

DUMAINE  Room for the incensed Worthies!

COSTARD  I'll do it in my shirt.	[He removes his doublet.]

DUMAINE  Most resolute Pompey!

BOY, [to Armado]  Master, let me take you a buttonhole
lower. Do you not see Pompey is uncasing for the
combat? What mean you? You will lose your
reputation.

ARMADO  Gentlemen and soldiers, pardon me. I will
not combat in my shirt.

DUMAINE  You may not deny it. Pompey hath made the
challenge.

ARMADO  Sweet bloods, I both may and will.

BEROWNE  What reason have you for 't?

ARMADO  The naked truth of it is, I have no shirt. I go
woolward for penance.

BOYET  True, and it was enjoined him in Rome for want
of linen; since when, I'll be sworn, he wore none
but a dishclout of Jaquenetta's, and that he wears
next his heart for a favor.

[Enter a Messenger, Monsieur Marcade.]


MARCADE, [to Princess]  God save you, madam.

PRINCESS  Welcome, Marcade,
But that thou interruptest our merriment.

MARCADE
I am sorry, madam, for the news I bring
Is heavy in my tongue. The King your father--

PRINCESS
Dead, for my life.

MARCADE  Even so. My tale is told.

BEROWNE
Worthies, away! The scene begins to cloud.

ARMADO  For mine own part, I breathe free breath. I
have seen the day of wrong through the little hole
of discretion, and I will right myself like a soldier.
[Worthies exit.]

KING, [to Princess]  How fares your Majesty?

PRINCESS
Boyet, prepare. I will away tonight.

KING
Madam, not so. I do beseech you stay.

PRINCESS, [to Boyet]
Prepare, I say.--I thank you, gracious lords,
For all your fair endeavors, and entreat,
Out of a new-sad soul, that you vouchsafe
In your rich wisdom to excuse or hide
The liberal opposition of our spirits,
If overboldly we have borne ourselves
In the converse of breath; your gentleness
Was guilty of it. Farewell, worthy lord.
A heavy heart bears not a humble tongue.
Excuse me so, coming too short of thanks
For my great suit so easily obtained.

KING
The extreme parts of time extremely forms
All causes to the purpose of his speed,
And often at his very loose decides
That which long process could not arbitrate.
And though the mourning brow of progeny
Forbid the smiling courtesy of love
The holy suit which fain it would convince,
Yet since love's argument was first on foot,
Let not the cloud of sorrow jostle it
From what it purposed, since to wail friends lost
Is not by much so wholesome-profitable
As to rejoice at friends but newly found.

PRINCESS
I understand you not. My griefs are double.

BEROWNE
Honest plain words best pierce the ear of grief,
And by these badges understand the King:
For your fair sakes have we neglected time,
Played foul play with our oaths. Your beauty, ladies,
Hath much deformed us, fashioning our humors
Even to the opposed end of our intents.
And what in us hath seemed ridiculous--
As love is full of unbefitting strains,
All wanton as a child, skipping and vain,
Formed by the eye and therefore, like the eye,
Full of strange shapes, of habits, and of forms,
Varying in subjects as the eye doth roll
To every varied object in his glance;
Which parti-coated presence of loose love
Put on by us, if, in your heavenly eyes,
Have misbecomed our oaths and gravities,
Those heavenly eyes, that look into these faults,
Suggested us to make. Therefore, ladies,
Our love being yours, the error that love makes
Is likewise yours. We to ourselves prove false
By being once false forever to be true
To those that make us both--fair ladies, you.
And even that falsehood, in itself a sin,
Thus purifies itself and turns to grace.

PRINCESS
We have received your letters full of love;
Your favors, the ambassadors of love;
And in our maiden council rated them
At courtship, pleasant jest, and courtesy,
As bombast and as lining to the time.
But more devout than this in our respects
Have we not been, and therefore met your loves
In their own fashion, like a merriment.

DUMAINE
Our letters, madam, showed much more than jest.

LONGAVILLE
So did our looks.

ROSALINE  We did not quote them so.

KING
Now, at the latest minute of the hour,
Grant us your loves.

PRINCESS  A time, methinks, too short
To make a world-without-end bargain in.
No, no, my lord, your Grace is perjured much,
Full of dear guiltiness, and therefore this:
If for my love--as there is no such cause--
You will do aught, this shall you do for me:
Your oath I will not trust, but go with speed
To some forlorn and naked hermitage,
Remote from all the pleasures of the world.
There stay until the twelve celestial signs
Have brought about the annual reckoning.
If this austere insociable life
Change not your offer made in heat of blood;
If frosts and fasts, hard lodging, and thin weeds
Nip not the gaudy blossoms of your love,
But that it bear this trial, and last love;
Then, at the expiration of the year,
Come challenge me, challenge me by these deserts,
[She takes his hand.]
And by this virgin palm now kissing thine,
I will be thine. And till that instant shut
My woeful self up in a mourning house,
Raining the tears of lamentation
For the remembrance of my father's death.
If this thou do deny, let our hands part,
Neither entitled in the other's heart.

KING
If this, or more than this, I would deny,
   To flatter up these powers of mine with rest,
The sudden hand of death close up mine eye!
   Hence hermit, then. My heart is in thy breast.
[They step aside.]

DUMAINE, [to Katherine]
But what to me, my love? But what to me?
A wife?

KATHERINE  A beard, fair health, and honesty.
With threefold love I wish you all these three.

DUMAINE
O, shall I say "I thank you, gentle wife"?

KATHERINE
Not so, my lord. A twelvemonth and a day
I'll mark no words that smooth-faced wooers say.
Come when the King doth to my lady come;
Then, if I have much love, I'll give you some.

DUMAINE
I'll serve thee true and faithfully till then.

KATHERINE
Yet swear not, lest you be forsworn again.
[They step aside.]

LONGAVILLE
What says Maria?

MARIA  At the twelvemonth's end
I'll change my black gown for a faithful friend.

LONGAVILLE
I'll stay with patience, but the time is long.

MARIA
The liker you; few taller are so young.
[They step aside.]

BEROWNE, [to Rosaline]
Studies my lady? Mistress, look on me.
Behold the window of my heart, mine eye,
What humble suit attends thy answer there.
Impose some service on me for thy love.

ROSALINE
Oft have I heard of you, my Lord Berowne,
Before I saw you; and the world's large tongue
Proclaims you for a man replete with mocks,
Full of comparisons and wounding flouts,
Which you on all estates will execute
That lie within the mercy of your wit.
To weed this wormwood from your fruitful brain,
And therewithal to win me, if you please,
Without the which I am not to be won,
You shall this twelvemonth term from day to day
Visit the speechless sick, and still converse
With groaning wretches; and your task shall be,
With all the fierce endeavor of your wit,
To enforce the pained impotent to smile.

BEROWNE
To move wild laughter in the throat of death?
It cannot be, it is impossible.
Mirth cannot move a soul in agony.

ROSALINE
Why, that's the way to choke a gibing spirit,
Whose influence is begot of that loose grace
Which shallow laughing hearers give to fools.
A jest's prosperity lies in the ear
Of him that hears it, never in the tongue
Of him that makes it. Then if sickly ears,
Deafed with the clamors of their own dear groans
Will hear your idle scorns, continue then,
And I will have you and that fault withal.
But if they will not, throw away that spirit,
And I shall find you empty of that fault,
Right joyful of your reformation.

BEROWNE
A twelvemonth? Well, befall what will befall,
I'll jest a twelvemonth in an hospital.

PRINCESS, [to King]
Ay, sweet my lord, and so I take my leave.

KING
No, madam, we will bring you on your way.

BEROWNE
Our wooing doth not end like an old play.
Jack hath not Jill. These ladies' courtesy
Might well have made our sport a comedy.

KING
Come, sir, it wants a twelvemonth and a day,
And then 'twill end.

BEROWNE  That's too long for a play.

[Enter Braggart Armado.]


ARMADO  Sweet Majesty, vouchsafe me--

PRINCESS
Was not that Hector?

DUMAINE  The worthy knight of Troy.

ARMADO  I will kiss thy royal finger, and take leave. I
am a votary; I have vowed to Jaquenetta to hold the
plow for her sweet love three year. But, most
esteemed Greatness, will you hear the dialogue that
the two learned men have compiled in praise of the
owl and the cuckoo? It should have followed in the
end of our show.

KING  Call them forth quickly. We will do so.

ARMADO  Holla! Approach.

[Enter all.]

This side is Hiems, Winter; this Ver, the Spring; the
one maintained by the owl, th' other by the cuckoo.
Ver, begin.
The Song.

SPRING
	When daisies pied and violets blue,
	   And lady-smocks all silver-white,
	And cuckoo-buds of yellow hue
	   Do paint the meadows with delight,
	The cuckoo then on every tree
	Mocks married men; for thus sings he:
		"Cuckoo!
	Cuckoo, cuckoo!" O word of fear,
	Unpleasing to a married ear.

	When shepherds pipe on oaten straws,
	   And merry larks are plowmen's clocks;
	When turtles tread, and rooks and daws,
	   And maidens bleach their summer smocks;
	The cuckoo then on every tree
	Mocks married men, for thus sings he:
		"Cuckoo!
	Cuckoo, cuckoo!" O word of fear,
	Unpleasing to a married ear.


WINTER
	When icicles hang by the wall,
	   And Dick the shepherd blows his nail,
	And Tom bears logs into the hall,
	   And milk comes frozen home in pail;
	When blood is nipped, and ways be foul,
	Then nightly sings the staring owl
	"Tu-whit to-who." A merry note,
	While greasy Joan doth keel the pot.

	When all aloud the wind doth blow,
	   And coughing drowns the parson's saw,
	And birds sit brooding in the snow,
	   And Marian's nose looks red and raw;
	When roasted crabs hiss in the bowl,
	Then nightly sings the staring owl
	"Tu-whit to-who." A merry note,
	While greasy Joan doth keel the pot.


ARMADO  The words of Mercury are harsh after the
songs of Apollo. You that way; we this way.
[They all exit.]
`

	b["richard-iii_TXT_FolgerShakespeare.txt"] = `
Richard III
by William Shakespeare
Edited by Barbara A. Mowat and Paul Werstine
  with Michael Poston and Rebecca Niles
Folger Shakespeare Library
https://shakespeare.folger.edu/shakespeares-works/richard-iii/
Created on Jul 31, 2015, from FDT version 0.9.2

Characters in the Play
======================
RICHARD, Duke of Gloucester, later King Richard III
LADY ANNE, widow of Edward, son to the late King Henry VI; later wife to Richard
KING EDWARD IV, brother to Richard
QUEEN ELIZABETH, Edward's wife, formerly the Lady Grey
Their sons:
  PRINCE EDWARD
  RICHARD, DUKE OF YORK
GEORGE, DUKE OF CLARENCE, brother to Edward and Richard
Clarence's BOY
Clarence's DAUGHTER
DUCHESS OF YORK, mother of Richard, Edward, and Clarence
QUEEN MARGARET, widow of King Henry VI
DUKE OF BUCKINGHAM
WILLIAM, LORD HASTINGS, Lord Chamberlain
LORD STANLEY, Earl of Derby
EARL RIVERS, brother to Queen Elizabeth
Sons of Queen Elizabeth by her former marriage:
  LORD GREY
  MARQUESS OF DORSET
SIR THOMAS VAUGHAN
Richard's supporters:
  SIR WILLIAM CATESBY
  SIR RICHARD RATCLIFFE
  LORD LOVELL
  DUKE OF NORFOLK
  EARL OF SURREY
EARL OF RICHMOND, Henry Tudor, later King Henry VII
Richmond's supporters:
  EARL OF OXFORD
  SIR JAMES BLUNT
  SIR WALTER HERBERT
  SIR WILLIAM BRANDON
  SIR CHRISTOPHER, a priest
ARCHBISHOP
CARDINAL
JOHN MORTON, BISHOP OF ELY
SIR ROBERT BRAKENBURY, Lieutenant of the Tower in London
JAMES TYRREL, gentleman
GENTLEMAN, attending Lady Anne
Two MURDERERS
KEEPER in the Tower
Three CITIZENS
LORD MAYOR of London
PURSUIVANT
SIR JOHN, a priest
SCRIVENER
PAGE
SHERIFF
Seven MESSENGERS
GHOSTS of King Henry VI, his son Prince Edward, Clarence, Rivers, Grey, Vaughan, the two Princes, Hastings, Lady Anne, and Buckingham
Guards, Tressel, Berkeley, Halberds, Gentlemen, Anthony Woodeville and Lord Scales (brothers to Queen Elizabeth), Two Bishops, Sir William Brandon, Lords, Attendants, Citizens, Aldermen, Councillors, Soldiers


ACT 1
=====

Scene 1
=======
[Enter Richard, Duke of Gloucester, alone.]


RICHARD
Now is the winter of our discontent
Made glorious summer by this son of York,
And all the clouds that loured upon our house
In the deep bosom of the ocean buried.
Now are our brows bound with victorious wreaths,
Our bruised arms hung up for monuments,
Our stern alarums changed to merry meetings,
Our dreadful marches to delightful measures.
Grim-visaged war hath smoothed his wrinkled front;
And now, instead of mounting barbed steeds
To fright the souls of fearful adversaries,
He capers nimbly in a lady's chamber
To the lascivious pleasing of a lute.
But I, that am not shaped for sportive tricks,
Nor made to court an amorous looking glass;
I, that am rudely stamped and want love's majesty
To strut before a wanton ambling nymph;
I, that am curtailed of this fair proportion,
Cheated of feature by dissembling nature,
Deformed, unfinished, sent before my time
Into this breathing world scarce half made up,
And that so lamely and unfashionable
That dogs bark at me as I halt by them--
Why, I, in this weak piping time of peace,
Have no delight to pass away the time,
Unless to see my shadow in the sun
And descant on mine own deformity.
And therefore, since I cannot prove a lover
To entertain these fair well-spoken days,
I am determined to prove a villain
And hate the idle pleasures of these days.
Plots have I laid, inductions dangerous,
By drunken prophecies, libels, and dreams,
To set my brother Clarence and the King
In deadly hate, the one against the other;
And if King Edward be as true and just
As I am subtle, false, and treacherous,
This day should Clarence closely be mewed up
About a prophecy which says that "G"
Of Edward's heirs the murderer shall be.
Dive, thoughts, down to my soul. Here Clarence
comes.

[Enter Clarence, guarded, and Brakenbury.]

Brother, good day. What means this armed guard
That waits upon your Grace?

CLARENCE  His Majesty,
Tend'ring my person's safety, hath appointed
This conduct to convey me to the Tower.

RICHARD
Upon what cause?

CLARENCE  Because my name is
George.

RICHARD
Alack, my lord, that fault is none of yours.
He should, for that, commit your godfathers.
O, belike his Majesty hath some intent
That you should be new christened in the Tower.
But what's the matter, Clarence? May I know?

CLARENCE
Yea, Richard, when I know, for I protest
As yet I do not. But, as I can learn,
He hearkens after prophecies and dreams,
And from the crossrow plucks the letter G,
And says a wizard told him that by "G"
His issue disinherited should be.
And for my name of George begins with G,
It follows in his thought that I am he.
These, as I learn, and such like toys as these
Hath moved his Highness to commit me now.

RICHARD
Why, this it is when men are ruled by women.
'Tis not the King that sends you to the Tower.
My Lady Grey his wife, Clarence, 'tis she
That tempers him to this extremity.
Was it not she and that good man of worship,
Anthony Woodeville, her brother there,
That made him send Lord Hastings to the Tower,
From whence this present day he is delivered?
We are not safe, Clarence; we are not safe.

CLARENCE
By heaven, I think there is no man secure
But the Queen's kindred and night-walking heralds
That trudge betwixt the King and Mistress Shore.
Heard you not what an humble suppliant
Lord Hastings was to her for his delivery?

RICHARD
Humbly complaining to her Deity
Got my Lord Chamberlain his liberty.
I'll tell you what: I think it is our way,
If we will keep in favor with the King,
To be her men and wear her livery.
The jealous o'erworn widow and herself,
Since that our brother dubbed them gentlewomen,
Are mighty gossips in our monarchy.

BRAKENBURY
I beseech your Graces both to pardon me.
His Majesty hath straitly given in charge
That no man shall have private conference,
Of what degree soever, with your brother.

RICHARD
Even so. An please your Worship, Brakenbury,
You may partake of anything we say.
We speak no treason, man. We say the King
Is wise and virtuous, and his noble queen
Well struck in years, fair, and not jealous.
We say that Shore's wife hath a pretty foot,
A cherry lip, a bonny eye, a passing pleasing tongue,
And that the Queen's kindred are made gentlefolks.
How say you, sir? Can you deny all this?

BRAKENBURY
With this, my lord, myself have naught to do.

RICHARD
Naught to do with Mistress Shore? I tell thee,
fellow,
He that doth naught with her, excepting one,
Were best to do it secretly, alone.

BRAKENBURY
I do beseech your Grace to pardon me, and withal
Forbear your conference with the noble duke.

CLARENCE
We know thy charge, Brakenbury, and will obey.

RICHARD
We are the Queen's abjects and must obey.--
Brother, farewell. I will unto the King,
And whatsoe'er you will employ me in,
Were it to call King Edward's widow "sister,"
I will perform it to enfranchise you.
Meantime, this deep disgrace in brotherhood
Touches me deeper than you can imagine.

CLARENCE
I know it pleaseth neither of us well.

RICHARD
Well, your imprisonment shall not be long.
I will deliver you or else lie for you.
Meantime, have patience.

CLARENCE  I must, perforce. Farewell.
[Exit Clarence, Brakenbury, and guard.]

RICHARD
Go tread the path that thou shalt ne'er return.
Simple, plain Clarence, I do love thee so
That I will shortly send thy soul to heaven,
If heaven will take the present at our hands.
But who comes here? The new-delivered Hastings?

[Enter Lord Hastings.]


HASTINGS
Good time of day unto my gracious lord.

RICHARD
As much unto my good Lord Chamberlain.
Well are you welcome to the open air.
How hath your Lordship brooked imprisonment?

HASTINGS
With patience, noble lord, as prisoners must.
But I shall live, my lord, to give them thanks
That were the cause of my imprisonment.

RICHARD
No doubt, no doubt; and so shall Clarence too,
For they that were your enemies are his
And have prevailed as much on him as you.

HASTINGS
More pity that the eagles should be mewed,
Whiles kites and buzzards prey at liberty.

RICHARD  What news abroad?

HASTINGS
No news so bad abroad as this at home:
The King is sickly, weak, and melancholy,
And his physicians fear him mightily.

RICHARD
Now, by Saint John, that news is bad indeed.
O, he hath kept an evil diet long,
And overmuch consumed his royal person.
'Tis very grievous to be thought upon.
Where is he, in his bed?

HASTINGS  He is.

RICHARD
Go you before, and I will follow you.
[Exit Hastings.]
He cannot live, I hope, and must not die
Till George be packed with post-horse up to heaven.
I'll in to urge his hatred more to Clarence
With lies well steeled with weighty arguments,
And, if I fail not in my deep intent,
Clarence hath not another day to live;
Which done, God take King Edward to His mercy,
And leave the world for me to bustle in.
For then I'll marry Warwick's youngest daughter.
What though I killed her husband and her father?
The readiest way to make the wench amends
Is to become her husband and her father;
The which will I, not all so much for love
As for another secret close intent
By marrying her which I must reach unto.
But yet I run before my horse to market.
Clarence still breathes; Edward still lives and reigns.
When they are gone, then must I count my gains.
[He exits.]

Scene 2
=======
[Enter the corse of Henry the Sixth on a bier, with
Halberds to guard it, Lady Anne being the mourner,
accompanied by Gentlemen.]


ANNE
Set down, set down your honorable load,
If honor may be shrouded in a hearse,
Whilst I awhile obsequiously lament
Th' untimely fall of virtuous Lancaster.
[They set down the bier.]
Poor key-cold figure of a holy king,
Pale ashes of the house of Lancaster,
Thou bloodless remnant of that royal blood,
Be it lawful that I invocate thy ghost
To hear the lamentations of poor Anne,
Wife to thy Edward, to thy slaughtered son,
Stabbed by the selfsame hand that made these
wounds.
Lo, in these windows that let forth thy life
I pour the helpless balm of my poor eyes.
O, cursed be the hand that made these holes;
Cursed the heart that had the heart to do it;
Cursed the blood that let this blood from hence.
More direful hap betide that hated wretch
That makes us wretched by the death of thee
Than I can wish to wolves, to spiders, toads,
Or any creeping venomed thing that lives.
If ever he have child, abortive be it,
Prodigious, and untimely brought to light,
Whose ugly and unnatural aspect
May fright the hopeful mother at the view,
And that be heir to his unhappiness.
If ever he have wife, let her be made
More miserable by the death of him
Than I am made by my young lord and thee.--
Come now towards Chertsey with your holy load,
Taken from Paul's to be interred there.
[They take up the bier.]
And still, as you are weary of this weight,
Rest you, whiles I lament King Henry's corse.

[Enter Richard, Duke of Gloucester.]


RICHARD
Stay, you that bear the corse, and set it down.

ANNE
What black magician conjures up this fiend
To stop devoted charitable deeds?

RICHARD
Villains, set down the corse or, by Saint Paul,
I'll make a corse of him that disobeys.

GENTLEMAN
My lord, stand back and let the coffin pass.

RICHARD
Unmannered dog, stand thou when I command!--
Advance thy halberd higher than my breast,
Or by Saint Paul I'll strike thee to my foot
And spurn upon thee, beggar, for thy boldness.
[They set down the bier.]

ANNE, [to the Gentlemen and Halberds]
What, do you tremble? Are you all afraid?
Alas, I blame you not, for you are mortal,
And mortal eyes cannot endure the devil.--
Avaunt, thou dreadful minister of hell.
Thou hadst but power over his mortal body;
His soul thou canst not have. Therefore begone.

RICHARD
Sweet saint, for charity, be not so curst.

ANNE
Foul devil, for God's sake, hence, and trouble us
not,
For thou hast made the happy Earth thy hell,
Filled it with cursing cries and deep exclaims.
If thou delight to view thy heinous deeds,
Behold this pattern of thy butcheries.
[She points to the corpse.]
O, gentlemen, see, see dead Henry's wounds
Open their congealed mouths and bleed afresh!--
Blush, blush, thou lump of foul deformity,
For 'tis thy presence that exhales this blood
From cold and empty veins where no blood dwells.
Thy deeds, inhuman and unnatural,
Provokes this deluge most unnatural.--
O God, which this blood mad'st, revenge his death!
O Earth, which this blood drink'st, revenge his
death!
Either heaven with lightning strike the murderer
dead,
Or Earth gape open wide and eat him quick,
As thou dost swallow up this good king's blood,
Which his hell-governed arm hath butchered.

RICHARD
Lady, you know no rules of charity,
Which renders good for bad, blessings for curses.

ANNE
Villain, thou know'st nor law of God nor man.
No beast so fierce but knows some touch of pity.

RICHARD
But I know none, and therefore am no beast.

ANNE
O, wonderful, when devils tell the truth!

RICHARD
More wonderful, when angels are so angry.
Vouchsafe, divine perfection of a woman,
Of these supposed crimes to give me leave
By circumstance but to acquit myself.

ANNE
Vouchsafe, defused infection of a man,
Of these known evils but to give me leave
By circumstance to curse thy cursed self.

RICHARD
Fairer than tongue can name thee, let me have
Some patient leisure to excuse myself.

ANNE
Fouler than heart can think thee, thou canst make
No excuse current but to hang thyself.

RICHARD
By such despair I should accuse myself.

ANNE
And by despairing shalt thou stand excused
For doing worthy vengeance on thyself
That didst unworthy slaughter upon others.

RICHARD  Say that I slew them not.

ANNE  Then say they were not slain.
But dead they are, and, devilish slave, by thee.

RICHARD  I did not kill your husband.

ANNE  Why then, he is alive.

RICHARD
Nay, he is dead, and slain by Edward's hands.

ANNE
In thy foul throat thou liest. Queen Margaret saw
Thy murd'rous falchion smoking in his blood,
The which thou once didst bend against her breast,
But that thy brothers beat aside the point.

RICHARD
I was provoked by her sland'rous tongue,
That laid their guilt upon my guiltless shoulders.

ANNE
Thou wast provoked by thy bloody mind,
That never dream'st on aught but butcheries.
Didst thou not kill this king?

RICHARD  I grant you.

ANNE
Dost grant me, hedgehog? Then, God grant me too
Thou mayst be damned for that wicked deed.
O, he was gentle, mild, and virtuous.

RICHARD
The better for the King of heaven that hath him.

ANNE
He is in heaven, where thou shalt never come.

RICHARD
Let him thank me, that holp to send him thither,
For he was fitter for that place than Earth.

ANNE
And thou unfit for any place but hell.

RICHARD
Yes, one place else, if you will hear me name it.

ANNE  Some dungeon.

RICHARD  Your bedchamber.

ANNE
Ill rest betide the chamber where thou liest!

RICHARD
So will it, madam, till I lie with you.

ANNE
I hope so.

RICHARD  I know so. But, gentle Lady Anne,
To leave this keen encounter of our wits
And fall something into a slower method:
Is not the causer of the timeless deaths
Of these Plantagenets, Henry and Edward,
As blameful as the executioner?

ANNE
Thou wast the cause and most accursed effect.

RICHARD
Your beauty was the cause of that effect--
Your beauty, that did haunt me in my sleep
To undertake the death of all the world,
So I might live one hour in your sweet bosom.

ANNE
If I thought that, I tell thee, homicide,
These nails should rend that beauty from my
cheeks.

RICHARD
These eyes could not endure that beauty's wrack.
You should not blemish it, if I stood by.
As all the world is cheered by the sun,
So I by that. It is my day, my life.

ANNE
Black night o'ershade thy day, and death thy life.

RICHARD
Curse not thyself, fair creature; thou art both.

ANNE
I would I were, to be revenged on thee.

RICHARD
It is a quarrel most unnatural
To be revenged on him that loveth thee.

ANNE
It is a quarrel just and reasonable
To be revenged on him that killed my husband.

RICHARD
He that bereft thee, lady, of thy husband
Did it to help thee to a better husband.

ANNE
His better doth not breathe upon the earth.

RICHARD
He lives that loves thee better than he could.

ANNE
Name him.

RICHARD  Plantagenet.

ANNE  Why, that was he.

RICHARD
The selfsame name, but one of better nature.

ANNE
Where is he?

RICHARD  Here. [(She spits at him.)] Why dost
thou spit at me?

ANNE
Would it were mortal poison for thy sake.

RICHARD
Never came poison from so sweet a place.

ANNE
Never hung poison on a fouler toad.
Out of my sight! Thou dost infect mine eyes.

RICHARD
Thine eyes, sweet lady, have infected mine.

ANNE
Would they were basilisks' to strike thee dead.

RICHARD
I would they were, that I might die at once,
For now they kill me with a living death.
Those eyes of thine from mine have drawn salt
tears,
Shamed their aspects with store of childish drops.
These eyes, which never shed remorseful tear--
No, when my father York and Edward wept
To hear the piteous moan that Rutland made
When black-faced Clifford shook his sword at him;
Nor when thy warlike father, like a child,
Told the sad story of my father's death
And twenty times made pause to sob and weep,
That all the standers-by had wet their cheeks
Like trees bedashed with rain--in that sad time,
My manly eyes did scorn an humble tear;
And what these sorrows could not thence exhale
Thy beauty hath, and made them blind with
weeping.
I never sued to friend nor enemy;
My tongue could never learn sweet smoothing word.
But now thy beauty is proposed my fee,
My proud heart sues and prompts my tongue to
speak.	[She looks scornfully at him.]
Teach not thy lip such scorn, for it was made
For kissing, lady, not for such contempt.
If thy revengeful heart cannot forgive,
Lo, here I lend thee this sharp-pointed sword,
Which if thou please to hide in this true breast
And let the soul forth that adoreth thee,
I lay it naked to the deadly stroke
And humbly beg the death upon my knee.
[He kneels and lays his breast open;
she offers at it with his sword.]
Nay, do not pause, for I did kill King Henry--
But 'twas thy beauty that provoked me.
Nay, now dispatch; 'twas I that stabbed young
Edward--
But 'twas thy heavenly face that set me on.
[She falls the sword.]
Take up the sword again, or take up me.

ANNE
Arise, dissembler. Though I wish thy death,
I will not be thy executioner.

RICHARD, [rising]
Then bid me kill myself, and I will do it.

ANNE
I have already.

RICHARD  That was in thy rage.
Speak it again and, even with the word,
This hand, which for thy love did kill thy love,
Shall for thy love kill a far truer love.
To both their deaths shalt thou be accessory.

ANNE  I would I knew thy heart.

RICHARD  'Tis figured in my tongue.

ANNE  I fear me both are false.

RICHARD  Then never was man true.

ANNE  Well, well, put up your sword.

RICHARD  Say then my peace is made.

ANNE  That shalt thou know hereafter.

RICHARD  But shall I live in hope?

ANNE  All men I hope live so.

RICHARD  Vouchsafe to wear this ring.

ANNE  To take is not to give.
[He places the ring on her hand.]

RICHARD
Look how my ring encompasseth thy finger;
Even so thy breast encloseth my poor heart.
Wear both of them, for both of them are thine.
And if thy poor devoted servant may
But beg one favor at thy gracious hand,
Thou dost confirm his happiness forever.

ANNE  What is it?

RICHARD
That it may please you leave these sad designs
To him that hath most cause to be a mourner,
And presently repair to Crosby House,
Where, after I have solemnly interred
At Chertsey monast'ry this noble king
And wet his grave with my repentant tears,
I will with all expedient duty see you.
For divers unknown reasons, I beseech you,
Grant me this boon.

ANNE
With all my heart, and much it joys me too
To see you are become so penitent.--
Tressel and Berkeley, go along with me.

RICHARD
Bid me farewell.

ANNE  'Tis more than you deserve;
But since you teach me how to flatter you,
Imagine I have said "farewell" already.
[Two exit with Anne. The bier is taken up.]

GENTLEMAN  Towards Chertsey, noble lord?

RICHARD
No, to Whitefriars. There attend my coming.
[Halberds and gentlemen exit with corse.]
Was ever woman in this humor wooed?
Was ever woman in this humor won?
I'll have her, but I will not keep her long.
What, I that killed her husband and his father,
To take her in her heart's extremest hate,
With curses in her mouth, tears in her eyes,
The bleeding witness of my hatred by,
Having God, her conscience, and these bars against
me,
And I no friends to back my suit at all
But the plain devil and dissembling looks?
And yet to win her, all the world to nothing!
Ha!
Hath she forgot already that brave prince,
Edward, her lord, whom I some three months since
Stabbed in my angry mood at Tewkesbury?
A sweeter and a lovelier gentleman,
Framed in the prodigality of nature,
Young, valiant, wise, and, no doubt, right royal,
The spacious world cannot again afford.
And will she yet abase her eyes on me,
That cropped the golden prime of this sweet prince
And made her widow to a woeful bed?
On me, whose all not equals Edward's moiety?
On me, that halts and am misshapen thus?
My dukedom to a beggarly denier,
I do mistake my person all this while!
Upon my life, she finds, although I cannot,
Myself to be a marv'lous proper man.
I'll be at charges for a looking glass
And entertain a score or two of tailors
To study fashions to adorn my body.
Since I am crept in favor with myself,
I will maintain it with some little cost.
But first I'll turn yon fellow in his grave
And then return lamenting to my love.
Shine out, fair sun, till I have bought a glass,
That I may see my shadow as I pass.
[He exits.]

Scene 3
=======
[Enter Queen Elizabeth, the Lord Marquess of Dorset,
Lord Rivers, and Lord Grey.]


RIVERS
Have patience, madam. There's no doubt his
Majesty
Will soon recover his accustomed health.

GREY
In that you brook it ill, it makes him worse.
Therefore, for God's sake, entertain good comfort
And cheer his Grace with quick and merry eyes.

QUEEN ELIZABETH
If he were dead, what would betide on me?

GREY
No other harm but loss of such a lord.

QUEEN ELIZABETH
The loss of such a lord includes all harms.

GREY
The heavens have blessed you with a goodly son
To be your comforter when he is gone.

QUEEN ELIZABETH
Ah, he is young, and his minority
Is put unto the trust of Richard Gloucester,
A man that loves not me nor none of you.

RIVERS
Is it concluded he shall be Protector?

QUEEN ELIZABETH
It is determined, not concluded yet;
But so it must be if the King miscarry.

[Enter Buckingham and Lord Stanley, Earl of Derby.]


GREY
Here comes the lord of Buckingham, and Derby.

BUCKINGHAM, [to Queen Elizabeth]
Good time of day unto your royal Grace.

STANLEY
God make your Majesty joyful, as you have been.

QUEEN ELIZABETH
The Countess Richmond, good my lord of Derby,
To your good prayer will scarcely say amen.
Yet, Derby, notwithstanding she's your wife
And loves not me, be you, good lord, assured
I hate not you for her proud arrogance.

STANLEY
I do beseech you either not believe
The envious slanders of her false accusers,
Or if she be accused on true report,
Bear with her weakness, which I think proceeds
From wayward sickness and no grounded malice.

QUEEN ELIZABETH
Saw you the King today, my lord of Derby?

STANLEY
But now the Duke of Buckingham and I
Are come from visiting his Majesty.

QUEEN ELIZABETH
What likelihood of his amendment, lords?

BUCKINGHAM
Madam, good hope. His Grace speaks cheerfully.

QUEEN ELIZABETH
God grant him health. Did you confer with him?

BUCKINGHAM
Ay, madam. He desires to make atonement
Between the Duke of Gloucester and your brothers,
And between them and my Lord Chamberlain,
And sent to warn them to his royal presence.

QUEEN ELIZABETH
Would all were well--but that will never be.
I fear our happiness is at the height.

[Enter Richard, Duke of Gloucester, and Hastings.]


RICHARD
They do me wrong, and I will not endure it!
Who is it that complains unto the King
That I, forsooth, am stern and love them not?
By holy Paul, they love his Grace but lightly
That fill his ears with such dissentious rumors.
Because I cannot flatter and look fair,
Smile in men's faces, smooth, deceive, and cog,
Duck with French nods and apish courtesy,
I must be held a rancorous enemy.
Cannot a plain man live and think no harm,
But thus his simple truth must be abused
With silken, sly, insinuating Jacks?

GREY
To who in all this presence speaks your Grace?

RICHARD
To thee, that hast nor honesty nor grace.
When have I injured thee? When done thee
wrong?--
Or thee?--Or thee? Or any of your faction?
A plague upon you all! His royal Grace,
Whom God preserve better than you would wish,
Cannot be quiet scarce a breathing while
But you must trouble him with lewd complaints.

QUEEN ELIZABETH
Brother of Gloucester, you mistake the matter.
The King, on his own royal disposition,
And not provoked by any suitor else,
Aiming belike at your interior hatred
That in your outward action shows itself
Against my children, brothers, and myself,
Makes him to send, that he may learn the ground.

RICHARD
I cannot tell. The world is grown so bad
That wrens make prey where eagles dare not perch.
Since every Jack became a gentleman,
There's many a gentle person made a Jack.

QUEEN ELIZABETH
Come, come, we know your meaning, brother
Gloucester.
You envy my advancement, and my friends'.
God grant we never may have need of you.

RICHARD
Meantime God grants that we have need of
you.
Our brother is imprisoned by your means,
Myself disgraced, and the nobility
Held in contempt, while great promotions
Are daily given to ennoble those
That scarce some two days since were worth a
noble.

QUEEN ELIZABETH
By Him that raised me to this careful height
From that contented hap which I enjoyed,
I never did incense his Majesty
Against the Duke of Clarence, but have been
An earnest advocate to plead for him.
My lord, you do me shameful injury
Falsely to draw me in these vile suspects.

RICHARD
You may deny that you were not the mean
Of my Lord Hastings' late imprisonment.

RIVERS  She may, my lord, for--

RICHARD
She may, Lord Rivers. Why, who knows not so?
She may do more, sir, than denying that.
She may help you to many fair preferments
And then deny her aiding hand therein,
And lay those honors on your high desert.
What may she not? She may, ay, marry, may she--

RIVERS  What, marry, may she?

RICHARD
What, marry, may she? Marry with a king,
A bachelor, and a handsome stripling too.
Iwis, your grandam had a worser match.

QUEEN ELIZABETH
My lord of Gloucester, I have too long borne
Your blunt upbraidings and your bitter scoffs.
By heaven, I will acquaint his Majesty
Of those gross taunts that oft I have endured.
I had rather be a country servant-maid
Than a great queen with this condition,
To be so baited, scorned, and stormed at.

[Enter old Queen Margaret, apart from the others.]

Small joy have I in being England's queen.

QUEEN MARGARET, [aside]
And lessened be that small, God I beseech Him!
Thy honor, state, and seat is due to me.

RICHARD, [to Queen Elizabeth]
What, threat you me with telling of the King?
Tell him and spare not. Look, what I have said,
I will avouch 't in presence of the King;
I dare adventure to be sent to th' Tower.
'Tis time to speak. My pains are quite forgot.

QUEEN MARGARET, [aside]
Out, devil! I do remember them too well:
Thou killed'st my husband Henry in the Tower,
And Edward, my poor son, at Tewkesbury.

RICHARD, [to Queen Elizabeth]
Ere you were queen, ay, or your husband king,
I was a packhorse in his great affairs,
A weeder-out of his proud adversaries,
A liberal rewarder of his friends.
To royalize his blood, I spent mine own.

QUEEN MARGARET, [aside]
Ay, and much better blood than his or thine.

RICHARD, [to Queen Elizabeth]
In all which time, you and your husband Grey
Were factious for the House of Lancaster.--
And, Rivers, so were you.--Was not your husband
In Margaret's battle at Saint Albans slain?
Let me put in your minds, if you forget,
What you have been ere this, and what you are;
Withal, what I have been, and what I am.

QUEEN MARGARET, [aside]
A murd'rous villain, and so still thou art.

RICHARD, [to Queen Elizabeth]
Poor Clarence did forsake his father Warwick,
Ay, and forswore himself--which Jesu pardon!--

QUEEN MARGARET, [aside]  Which God revenge!

RICHARD
To fight on Edward's party for the crown;
And for his meed, poor lord, he is mewed up.
I would to God my heart were flint, like Edward's,
Or Edward's soft and pitiful, like mine.
I am too childish-foolish for this world.

QUEEN MARGARET, [aside]
Hie thee to hell for shame, and leave this world,
Thou cacodemon! There thy kingdom is.

RIVERS
My lord of Gloucester, in those busy days
Which here you urge to prove us enemies,
We followed then our lord, our sovereign king.
So should we you, if you should be our king.

RICHARD
If I should be? I had rather be a peddler.
Far be it from my heart, the thought thereof.

QUEEN ELIZABETH
As little joy, my lord, as you suppose
You should enjoy were you this country's king,
As little joy you may suppose in me
That I enjoy, being the queen thereof.

QUEEN MARGARET, [aside]
As little joy enjoys the queen thereof,
For I am she, and altogether joyless.
I can no longer hold me patient.
[She steps forward.]
Hear me, you wrangling pirates, that fall out
In sharing that which you have pilled from me!
Which of you trembles not that looks on me?
If not, that I am queen, you bow like subjects,
Yet that, by you deposed, you quake like rebels.--
Ah, gentle villain, do not turn away.

RICHARD
Foul, wrinkled witch, what mak'st thou in my
sight?

QUEEN MARGARET
But repetition of what thou hast marred.
That will I make before I let thee go.

RICHARD
Wert thou not banished on pain of death?

QUEEN MARGARET
I was, but I do find more pain in banishment
Than death can yield me here by my abode.
A husband and a son thou ow'st to me;
[To Queen Elizabeth.] And thou a kingdom;--all
of you, allegiance.
This sorrow that I have by right is yours,
And all the pleasures you usurp are mine.

RICHARD
The curse my noble father laid on thee
When thou didst crown his warlike brows with
paper,
And with thy scorns drew'st rivers from his eyes,
And then, to dry them, gav'st the Duke a clout
Steeped in the faultless blood of pretty Rutland--
His curses then, from bitterness of soul
Denounced against thee, are all fall'n upon thee,
And God, not we, hath plagued thy bloody deed.

QUEEN ELIZABETH
So just is God to right the innocent.

HASTINGS
O, 'twas the foulest deed to slay that babe,
And the most merciless that e'er was heard of!

RIVERS
Tyrants themselves wept when it was reported.

DORSET
No man but prophesied revenge for it.

BUCKINGHAM
Northumberland, then present, wept to see it.

QUEEN MARGARET
What, were you snarling all before I came,
Ready to catch each other by the throat,
And turn you all your hatred now on me?
Did York's dread curse prevail so much with
heaven
That Henry's death, my lovely Edward's death,
Their kingdom's loss, my woeful banishment,
Should all but answer for that peevish brat?
Can curses pierce the clouds and enter heaven?
Why then, give way, dull clouds, to my quick
curses!
Though not by war, by surfeit die your king,
As ours by murder to make him a king.
[To Queen Elizabeth.] Edward thy son, that now is
Prince of Wales,
For Edward our son, that was Prince of Wales,
Die in his youth by like untimely violence.
Thyself a queen, for me that was a queen,
Outlive thy glory, like my wretched self.
Long mayst thou live to wail thy children's death
And see another, as I see thee now,
Decked in thy rights, as thou art stalled in mine.
Long die thy happy days before thy death,
And, after many lengthened hours of grief,
Die neither mother, wife, nor England's queen.--
Rivers and Dorset, you were standers-by,
And so wast thou, Lord Hastings, when my son
Was stabbed with bloody daggers. God I pray Him
That none of you may live his natural age,
But by some unlooked accident cut off.

RICHARD
Have done thy charm, thou hateful, withered hag.

QUEEN MARGARET
And leave out thee? Stay, dog, for thou shalt hear
me.
If heaven have any grievous plague in store
Exceeding those that I can wish upon thee,
O, let them keep it till thy sins be ripe
And then hurl down their indignation
On thee, the troubler of the poor world's peace.
The worm of conscience still begnaw thy soul.
Thy friends suspect for traitors while thou liv'st,
And take deep traitors for thy dearest friends.
No sleep close up that deadly eye of thine,
Unless it be while some tormenting dream
Affrights thee with a hell of ugly devils.
Thou elvish-marked, abortive, rooting hog,
Thou that wast sealed in thy nativity
The slave of nature and the son of hell,
Thou slander of thy heavy mother's womb,
Thou loathed issue of thy father's loins,
Thou rag of honor, thou detested--

RICHARD  Margaret.

QUEEN MARGARET  Richard!

RICHARD  Ha?

QUEEN MARGARET  I call thee not.

RICHARD
I cry thee mercy, then, for I did think
That thou hadst called me all these bitter names.

QUEEN MARGARET
Why, so I did, but looked for no reply.
O, let me make the period to my curse!

RICHARD
'Tis done by me and ends in "Margaret."

QUEEN ELIZABETH, [to Queen Margaret]
Thus have you breathed your curse against yourself.

QUEEN MARGARET
Poor painted queen, vain flourish of my fortune,
Why strew'st thou sugar on that bottled spider,
Whose deadly web ensnareth thee about?
Fool, fool, thou whet'st a knife to kill thyself.
The day will come that thou shalt wish for me
To help thee curse this poisonous bunch-backed
toad.

HASTINGS
False-boding woman, end thy frantic curse,
Lest to thy harm thou move our patience.

QUEEN MARGARET
Foul shame upon you, you have all moved mine.

RIVERS
Were you well served, you would be taught your
duty.

QUEEN MARGARET
To serve me well, you all should do me duty:
Teach me to be your queen, and you my subjects.
O, serve me well, and teach yourselves that duty!

DORSET, [to Rivers]
Dispute not with her; she is lunatic.

QUEEN MARGARET
Peace, Master Marquess, you are malapert.
Your fire-new stamp of honor is scarce current.
O, that your young nobility could judge
What 'twere to lose it and be miserable!
They that stand high have many blasts to shake
them,
And if they fall, they dash themselves to pieces.

RICHARD
Good counsel, marry.--Learn it, learn it, marquess.

DORSET
It touches you, my lord, as much as me.

RICHARD
Ay, and much more; but I was born so high.
Our aerie buildeth in the cedar's top,
And dallies with the wind and scorns the sun.

QUEEN MARGARET
And turns the sun to shade. Alas, alas,
Witness my son, now in the shade of death,
Whose bright out-shining beams thy cloudy wrath
Hath in eternal darkness folded up.
Your aerie buildeth in our aerie's nest.
O God, that seest it, do not suffer it!
As it is won with blood, lost be it so.

BUCKINGHAM
Peace, peace, for shame, if not for charity.

QUEEN MARGARET
Urge neither charity nor shame to me.
[Addressing the others.] Uncharitably with me have
you dealt,
And shamefully my hopes by you are butchered.
My charity is outrage, life my shame,
And in that shame still live my sorrows' rage.

BUCKINGHAM  Have done, have done.

QUEEN MARGARET
O princely Buckingham, I'll kiss thy hand
In sign of league and amity with thee.
Now fair befall thee and thy noble house!
Thy garments are not spotted with our blood,
Nor thou within the compass of my curse.

BUCKINGHAM
Nor no one here, for curses never pass
The lips of those that breathe them in the air.

QUEEN MARGARET
I will not think but they ascend the sky,
And there awake God's gentle sleeping peace.
[Aside to Buckingham.] O Buckingham, take heed of
yonder dog!
Look when he fawns, he bites; and when he bites,
His venom tooth will rankle to the death.
Have not to do with him. Beware of him.
Sin, death, and hell have set their marks on him,
And all their ministers attend on him.

RICHARD
What doth she say, my lord of Buckingham?

BUCKINGHAM
Nothing that I respect, my gracious lord.

QUEEN MARGARET
What, dost thou scorn me for my gentle counsel,
And soothe the devil that I warn thee from?
O, but remember this another day,
When he shall split thy very heart with sorrow,
And say poor Margaret was a prophetess.--
Live each of you the subjects to his hate,
And he to yours, and all of you to God's.	[She exits.]

BUCKINGHAM
My hair doth stand an end to hear her curses.

RIVERS
And so doth mine. I muse why she's at liberty.

RICHARD
I cannot blame her. By God's holy mother,
She hath had too much wrong, and I repent
My part thereof that I have done to her.

QUEEN ELIZABETH
I never did her any, to my knowledge.

RICHARD
Yet you have all the vantage of her wrong.
I was too hot to do somebody good
That is too cold in thinking of it now.
Marry, as for Clarence, he is well repaid;
He is franked up to fatting for his pains.
God pardon them that are the cause thereof.

RIVERS
A virtuous and a Christian-like conclusion
To pray for them that have done scathe to us.

RICHARD
So do I ever--[(speaks to himself)] being well advised,
For had I cursed now, I had cursed myself.

[Enter Catesby.]


CATESBY
Madam, his Majesty doth call for you,--
And for your Grace,--and yours, my gracious
lords.

QUEEN ELIZABETH
Catesby, I come.--Lords, will you go with me?

RIVERS  We wait upon your Grace.
[All but Richard, Duke of Gloucester exit.]

RICHARD
I do the wrong and first begin to brawl.
The secret mischiefs that I set abroach
I lay unto the grievous charge of others.
Clarence, who I indeed have cast in darkness,
I do beweep to many simple gulls,
Namely, to Derby, Hastings, Buckingham,
And tell them 'tis the Queen and her allies
That stir the King against the Duke my brother.
Now they believe it and withal whet me
To be revenged on Rivers, Dorset, Grey;
But then I sigh and, with a piece of scripture,
Tell them that God bids us do good for evil;
And thus I clothe my naked villainy
With odd old ends stol'n forth of Holy Writ,
And seem a saint when most I play the devil.

[Enter two Murderers.]

But soft, here come my executioners.--
How now, my hardy, stout, resolved mates?
Are you now going to dispatch this thing?

MURDERER
We are, my lord, and come to have the warrant
That we may be admitted where he is.

RICHARD
Well thought upon. I have it here about me.
[He gives a paper.]
When you have done, repair to Crosby Place.
But, sirs, be sudden in the execution,
Withal obdurate; do not hear him plead,
For Clarence is well-spoken and perhaps
May move your hearts to pity if you mark him.

MURDERER
Tut, tut, my lord, we will not stand to prate.
Talkers are no good doers. Be assured
We go to use our hands and not our tongues.

RICHARD
Your eyes drop millstones when fools' eyes fall
tears.
I like you lads. About your business straight.
Go, go, dispatch.

MURDERERS  We will, my noble lord.
[They exit.]

Scene 4
=======
[Enter Clarence and Keeper.]


KEEPER
Why looks your Grace so heavily today?

CLARENCE
O, I have passed a miserable night,
So full of fearful dreams, of ugly sights,
That, as I am a Christian faithful man,
I would not spend another such a night
Though 'twere to buy a world of happy days,
So full of dismal terror was the time.

KEEPER
What was your dream, my lord? I pray you tell me.

CLARENCE
Methoughts that I had broken from the Tower
And was embarked to cross to Burgundy,
And in my company my brother Gloucester,
Who from my cabin tempted me to walk
Upon the hatches. Thence we looked toward
England
And cited up a thousand heavy times,
During the wars of York and Lancaster,
That had befall'n us. As we paced along
Upon the giddy footing of the hatches,
Methought that Gloucester stumbled, and in falling
Struck me, that thought to stay him, overboard
Into the tumbling billows of the main.
O Lord, methought what pain it was to drown,
What dreadful noise of waters in my ears,
What sights of ugly death within my eyes.
Methoughts I saw a thousand fearful wracks,
A thousand men that fishes gnawed upon,
Wedges of gold, great anchors, heaps of pearl,
Inestimable stones, unvalued jewels,
All scattered in the bottom of the sea.
Some lay in dead men's skulls, and in the holes
Where eyes did once inhabit, there were crept--
As 'twere in scorn of eyes--reflecting gems,
That wooed the slimy bottom of the deep
And mocked the dead bones that lay scattered by.

KEEPER
Had you such leisure in the time of death
To gaze upon these secrets of the deep?

CLARENCE
Methought I had, and often did I strive
To yield the ghost, but still the envious flood
Stopped in my soul and would not let it forth
To find the empty, vast, and wand'ring air,
But smothered it within my panting bulk,
Who almost burst to belch it in the sea.

KEEPER
Awaked you not in this sore agony?

CLARENCE
No, no, my dream was lengthened after life.
O, then began the tempest to my soul.
I passed, methought, the melancholy flood,
With that sour ferryman which poets write of,
Unto the kingdom of perpetual night.
The first that there did greet my stranger-soul
Was my great father-in-law, renowned Warwick,
Who spake aloud "What scourge for perjury
Can this dark monarchy afford false Clarence?"
And so he vanished. Then came wand'ring by
A shadow like an angel, with bright hair
Dabbled in blood, and he shrieked out aloud
"Clarence is come--false, fleeting, perjured
Clarence,
That stabbed me in the field by Tewkesbury.
Seize on him, furies. Take him unto torment."
With that, methoughts, a legion of foul fiends
Environed me and howled in mine ears
Such hideous cries that with the very noise
I trembling waked, and for a season after
Could not believe but that I was in hell,
Such terrible impression made my dream.

KEEPER
No marvel, lord, though it affrighted you.
I am afraid, methinks, to hear you tell it.

CLARENCE
Ah keeper, keeper, I have done these things,
That now give evidence against my soul,
For Edward's sake, and see how he requites me.--
O God, if my deep prayers cannot appease thee,
But thou wilt be avenged on my misdeeds,
Yet execute thy wrath in me alone!
O, spare my guiltless wife and my poor children!--
Keeper, I prithee sit by me awhile.
My soul is heavy, and I fain would sleep.

KEEPER
I will, my lord. God give your Grace good rest.
[Clarence sleeps.]

[Enter Brakenbury the Lieutenant.]


BRAKENBURY
Sorrow breaks seasons and reposing hours,
Makes the night morning, and the noontide night.
Princes have but their titles for their glories,
An outward honor for an inward toil,
And, for unfelt imaginations,
They often feel a world of restless cares,
So that between their titles and low name
There's nothing differs but the outward fame.

[Enter two Murderers.]


FIRST MURDERER  Ho, who's here?

BRAKENBURY
What wouldst thou, fellow? And how cam'st thou
hither?

SECOND MURDERER  I would speak with Clarence, and I
came hither on my legs.

BRAKENBURY  What, so brief?

FIRST MURDERER  'Tis better, sir, than to be tedious.--
Let him see our commission, and talk no more.
[Brakenbury reads the commission.]

BRAKENBURY
I am in this commanded to deliver
The noble Duke of Clarence to your hands.
I will not reason what is meant hereby
Because I will be guiltless from the meaning.
There lies the Duke asleep, and there the keys.
[He hands them keys.]
I'll to the King and signify to him
That thus I have resigned to you my charge.

FIRST MURDERER  You may, sir. 'Tis a point of wisdom.
Fare you well.
[Brakenbury and the Keeper exit.]

SECOND MURDERER  What, shall I stab him as he
sleeps?

FIRST MURDERER  No. He'll say 'twas done cowardly,
when he wakes.

SECOND MURDERER  Why, he shall never wake until the
great Judgment Day.

FIRST MURDERER  Why, then he'll say we stabbed him
sleeping.

SECOND MURDERER  The urging of that word "judgment"
hath bred a kind of remorse in me.

FIRST MURDERER  What, art thou afraid?

SECOND MURDERER  Not to kill him, having a warrant,
but to be damned for killing him, from the which
no warrant can defend me.

FIRST MURDERER  I thought thou hadst been resolute.

SECOND MURDERER  So I am--to let him live.

FIRST MURDERER  I'll back to the Duke of Gloucester
and tell him so.

SECOND MURDERER  Nay, I prithee stay a little. I hope
this passionate humor of mine will change. It was
wont to hold me but while one tells twenty.

FIRST MURDERER  How dost thou feel thyself now?

SECOND MURDERER  Faith, some certain dregs of conscience
are yet within me.

FIRST MURDERER  Remember our reward when the
deed's done.

SECOND MURDERER  Zounds, he dies! I had forgot the
reward.

FIRST MURDERER  Where's thy conscience now?

SECOND MURDERER  O, in the Duke of Gloucester's
purse.

FIRST MURDERER  When he opens his purse to give us
our reward, thy conscience flies out.

SECOND MURDERER  'Tis no matter. Let it go. There's
few or none will entertain it.

FIRST MURDERER  What if it come to thee again?

SECOND MURDERER  I'll not meddle with it. It makes a
man a coward: a man cannot steal but it accuseth
him; a man cannot swear but it checks him; a man
cannot lie with his neighbor's wife but it detects
him. 'Tis a blushing, shamefaced spirit that mutinies
in a man's bosom. It fills a man full of
obstacles. It made me once restore a purse of gold
that by chance I found. It beggars any man that
keeps it. It is turned out of towns and cities for a
dangerous thing, and every man that means to live
well endeavors to trust to himself and live without it.

FIRST MURDERER  Zounds, 'tis even now at my elbow,
persuading me not to kill the Duke.

SECOND MURDERER  Take the devil in thy mind, and
believe him not. He would insinuate with thee but
to make thee sigh.

FIRST MURDERER  I am strong-framed. He cannot prevail
with me.

SECOND MURDERER  Spoke like a tall man that respects
thy reputation. Come, shall we fall to work?

FIRST MURDERER  Take him on the costard with the
hilts of thy sword, and then throw him into the
malmsey butt in the next room.

SECOND MURDERER  O, excellent device--and make a
sop of him!

FIRST MURDERER  Soft, he wakes.

SECOND MURDERER  Strike!

FIRST MURDERER  No, we'll reason with him.
[Clarence wakes.]

CLARENCE
Where art thou, keeper? Give me a cup of wine.

SECOND MURDERER
You shall have wine enough, my lord, anon.

CLARENCE
In God's name, what art thou?

FIRST MURDERER  A man, as you are.

CLARENCE  But not, as I am, royal.

FIRST MURDERER  Nor you, as we are, loyal.

CLARENCE
Thy voice is thunder, but thy looks are humble.

FIRST MURDERER
My voice is now the King's, my looks mine own.

CLARENCE
How darkly and how deadly dost thou speak!
Your eyes do menace me. Why look you pale?
Who sent you hither? Wherefore do you come?

SECOND MURDERER  To, to, to--

CLARENCE  To murder me?

BOTH  Ay, ay.

CLARENCE
You scarcely have the hearts to tell me so
And therefore cannot have the hearts to do it.
Wherein, my friends, have I offended you?

FIRST MURDERER
Offended us you have not, but the King.

CLARENCE
I shall be reconciled to him again.

SECOND MURDERER
Never, my lord. Therefore prepare to die.

CLARENCE
Are you drawn forth among a world of men
To slay the innocent? What is my offense?
Where is the evidence that doth accuse me?
What lawful quest have given their verdict up
Unto the frowning judge? Or who pronounced
The bitter sentence of poor Clarence' death
Before I be convict by course of law?
To threaten me with death is most unlawful.
I charge you, as you hope to have redemption,
By Christ's dear blood shed for our grievous sins,
That you depart, and lay no hands on me.
The deed you undertake is damnable.

FIRST MURDERER
What we will do, we do upon command.

SECOND MURDERER
And he that hath commanded is our king.

CLARENCE
Erroneous vassals, the great King of kings
Hath in the table of His law commanded
That thou shalt do no murder. Will you then
Spurn at His edict and fulfill a man's?
Take heed, for He holds vengeance in His hand
To hurl upon their heads that break His law.

SECOND MURDERER
And that same vengeance doth He hurl on thee
For false forswearing and for murder too.
Thou didst receive the sacrament to fight
In quarrel of the House of Lancaster.

FIRST MURDERER
And, like a traitor to the name of God,
Didst break that vow, and with thy treacherous
blade
Unrippedst the bowels of thy sovereign's son.

SECOND MURDERER
Whom thou wast sworn to cherish and defend.

FIRST MURDERER
How canst thou urge God's dreadful law to us
When thou hast broke it in such dear degree?

CLARENCE
Alas! For whose sake did I that ill deed?
For Edward, for my brother, for his sake.
He sends you not to murder me for this,
For in that sin he is as deep as I.
If God will be avenged for the deed,
O, know you yet He doth it publicly!
Take not the quarrel from His powerful arm;
He needs no indirect or lawless course
To cut off those that have offended Him.

FIRST MURDERER
Who made thee then a bloody minister
When gallant-springing, brave Plantagenet,
That princely novice, was struck dead by thee?

CLARENCE
My brother's love, the devil, and my rage.

FIRST MURDERER
Thy brother's love, our duty, and thy faults
Provoke us hither now to slaughter thee.

CLARENCE
If you do love my brother, hate not me.
I am his brother, and I love him well.
If you are hired for meed, go back again,
And I will send you to my brother Gloucester,
Who shall reward you better for my life
Than Edward will for tidings of my death.

SECOND MURDERER
You are deceived. Your brother Gloucester hates
you.

CLARENCE
O no, he loves me, and he holds me dear.
Go you to him from me.

FIRST MURDERER  Ay, so we will.

CLARENCE
Tell him, when that our princely father York
Blessed his three sons with his victorious arm,
He little thought of this divided friendship.
Bid Gloucester think of this, and he will weep.

FIRST MURDERER
Ay, millstones, as he lessoned us to weep.

CLARENCE
O, do not slander him, for he is kind.

FIRST MURDERER
Right, as snow in harvest. Come, you deceive
yourself.
'Tis he that sends us to destroy you here.

CLARENCE
It cannot be, for he bewept my fortune,
And hugged me in his arms, and swore with sobs
That he would labor my delivery.

FIRST MURDERER
Why, so he doth, when he delivers you
From this Earth's thralldom to the joys of heaven.

SECOND MURDERER
Make peace with God, for you must die, my lord.

CLARENCE
Have you that holy feeling in your souls
To counsel me to make my peace with God,
And are you yet to your own souls so blind
That you will war with God by murd'ring me?
O sirs, consider: they that set you on
To do this deed will hate you for the deed.

SECOND MURDERER, [to First Murderer]
What shall we do?

CLARENCE  Relent, and save your souls.
Which of you--if you were a prince's son
Being pent from liberty, as I am now--
If two such murderers as yourselves came to you,
Would not entreat for life? Ay, you would beg,
Were you in my distress.

FIRST MURDERER
Relent? No. 'Tis cowardly and womanish.

CLARENCE
Not to relent is beastly, savage, devilish.
[To Second Murderer.] My friend, I spy some pity
in thy looks.
O, if thine eye be not a flatterer,
Come thou on my side and entreat for me.
A begging prince what beggar pities not?

SECOND MURDERER  Look behind you, my lord.

FIRST MURDERER
Take that, and that. [(Stabs him.)] If all this will not
do,
I'll drown you in the malmsey butt within.
[He exits with the body.]

SECOND MURDERER
A bloody deed, and desperately dispatched.
How fain, like Pilate, would I wash my hands
Of this most grievous murder.

[Enter First Murderer.]


FIRST MURDERER
How now? What mean'st thou that thou help'st me
not?
By heavens, the Duke shall know how slack you
have been.

SECOND MURDERER
I would he knew that I had saved his brother.
Take thou the fee, and tell him what I say,
For I repent me that the Duke is slain.	[He exits.]

FIRST MURDERER
So do not I. Go, coward as thou art.
Well, I'll go hide the body in some hole
Till that the Duke give order for his burial.
And when I have my meed, I will away,
For this will out, and then I must not stay.
[He exits.]


ACT 2
=====

Scene 1
=======
[Flourish. Enter King Edward, sick, Queen Elizabeth,
Lord Marquess Dorset, Rivers, Hastings, Buckingham,
Woodeville, Grey, and Scales.]


KING EDWARD
Why, so. Now have I done a good day's work.
You peers, continue this united league.
I every day expect an embassage
From my Redeemer to redeem me hence,
And more in peace my soul shall part to heaven
Since I have made my friends at peace on Earth.
Rivers and Hastings, take each other's hand.
Dissemble not your hatred. Swear your love.

RIVERS, [taking Hastings' hand]
By heaven, my soul is purged from grudging hate,
And with my hand I seal my true heart's love.

HASTINGS
So thrive I as I truly swear the like.

KING EDWARD
Take heed you dally not before your king,
Lest He that is the supreme King of kings
Confound your hidden falsehood and award
Either of you to be the other's end.

HASTINGS
So prosper I as I swear perfect love.

RIVERS
And I as I love Hastings with my heart.

KING EDWARD, [to Queen Elizabeth]
Madam, yourself is not exempt from this,--
Nor you, son Dorset,--Buckingham, nor you.
You have been factious one against the other.--
Wife, love Lord Hastings. Let him kiss your hand,
And what you do, do it unfeignedly.

QUEEN ELIZABETH
There, Hastings, I will never more remember
Our former hatred, so thrive I and mine.
[Hastings kisses her hand.]

KING EDWARD
Dorset, embrace him.--Hastings, love Lord
Marquess.

DORSET
This interchange of love, I here protest,
Upon my part shall be inviolable.

HASTINGS  And so swear I.	[They embrace.]

KING EDWARD
Now, princely Buckingham, seal thou this league
With thy embracements to my wife's allies
And make me happy in your unity.

BUCKINGHAM, [to Queen Elizabeth]
Whenever Buckingham doth turn his hate
Upon your Grace, but with all duteous love
Doth cherish you and yours, God punish me
With hate in those where I expect most love.
When I have most need to employ a friend,
And most assured that he is a friend,
Deep, hollow, treacherous, and full of guile
Be he unto me: this do I beg of God,
When I am cold in love to you or yours.
[Queen Elizabeth and Buckingham embrace.]

KING EDWARD
A pleasing cordial, princely Buckingham,
Is this thy vow unto my sickly heart.
There wanteth now our brother Gloucester here
To make the blessed period of this peace.

BUCKINGHAM  And in good time
Here comes Sir Richard Ratcliffe and the Duke.

[Enter Ratcliffe, and Richard, Duke of Gloucester.]


RICHARD
Good morrow to my sovereign king and queen,
And, princely peers, a happy time of day.

KING EDWARD
Happy indeed, as we have spent the day.
Gloucester, we have done deeds of charity,
Made peace of enmity, fair love of hate,
Between these swelling, wrong-incensed peers.

RICHARD
A blessed labor, my most sovereign lord.
Among this princely heap, if any here
By false intelligence or wrong surmise
Hold me a foe,
If I unwittingly, or in my rage,
Have aught committed that is hardly borne
By any in this presence, I desire
To reconcile me to his friendly peace.
'Tis death to me to be at enmity;
I hate it, and desire all good men's love.
First, madam, I entreat true peace of you,
Which I will purchase with my duteous service;--
Of you, my noble cousin Buckingham,
If ever any grudge were lodged between us;--
Of you and you, Lord Rivers and of Dorset,
That all without desert have frowned on me;--
Of you, Lord Woodeville and Lord Scales;--of you,
Dukes, earls, lords, gentlemen; indeed, of all.
I do not know that Englishman alive
With whom my soul is any jot at odds
More than the infant that is born tonight.
I thank my God for my humility.

QUEEN ELIZABETH
A holy day shall this be kept hereafter.
I would to God all strifes were well compounded.
My sovereign lord, I do beseech your Highness
To take our brother Clarence to your grace.

RICHARD
Why, madam, have I offered love for this,
To be so flouted in this royal presence?
Who knows not that the gentle duke is dead?
[They all start.]
You do him injury to scorn his corse.

KING EDWARD
Who knows not he is dead! Who knows he is?

QUEEN ELIZABETH
All-seeing heaven, what a world is this!

BUCKINGHAM
Look I so pale, Lord Dorset, as the rest?

DORSET
Ay, my good lord, and no man in the presence
But his red color hath forsook his cheeks.

KING EDWARD
Is Clarence dead? The order was reversed.

RICHARD
But he, poor man, by your first order died,
And that a winged Mercury did bear.
Some tardy cripple bare the countermand,
That came too lag to see him buried.
God grant that some, less noble and less loyal,
Nearer in bloody thoughts, and not in blood,
Deserve not worse than wretched Clarence did,
And yet go current from suspicion.

[Enter Lord Stanley, Earl of Derby.]


STANLEY, [kneeling]
A boon, my sovereign, for my service done.

KING EDWARD
I prithee, peace. My soul is full of sorrow.

STANLEY
I will not rise unless your Highness hear me.

KING EDWARD
Then say at once what is it thou requests.

STANLEY
The forfeit, sovereign, of my servant's life,
Who slew today a riotous gentleman
Lately attendant on the Duke of Norfolk.

KING EDWARD
Have I a tongue to doom my brother's death,
And shall that tongue give pardon to a slave?
My brother killed no man; his fault was thought,
And yet his punishment was bitter death.
Who sued to me for him? Who, in my wrath,
Kneeled at my feet, and bade me be advised?
Who spoke of brotherhood? Who spoke of love?
Who told me how the poor soul did forsake
The mighty Warwick and did fight for me?
Who told me, in the field at Tewkesbury,
When Oxford had me down, he rescued me,
And said "Dear brother, live, and be a king"?
Who told me, when we both lay in the field
Frozen almost to death, how he did lap me
Even in his garments and did give himself,
All thin and naked, to the numb-cold night?
All this from my remembrance brutish wrath
Sinfully plucked, and not a man of you
Had so much grace to put it in my mind.
But when your carters or your waiting vassals
Have done a drunken slaughter and defaced
The precious image of our dear Redeemer,
You straight are on your knees for pardon, pardon,
And I, unjustly too, must grant it you.
[Stanley rises.]
But for my brother, not a man would speak,
Nor I, ungracious, speak unto myself
For him, poor soul. The proudest of you all
Have been beholding to him in his life,
Yet none of you would once beg for his life.
O God, I fear Thy justice will take hold
On me and you, and mine and yours for this!--
Come, Hastings, help me to my closet.--
Ah, poor Clarence.
[Some exit with King and Queen.]

RICHARD
This is the fruits of rashness. Marked you not
How that the guilty kindred of the Queen
Looked pale when they did hear of Clarence' death?
O, they did urge it still unto the King.
God will revenge it. Come, lords, will you go
To comfort Edward with our company?

BUCKINGHAM  We wait upon your Grace.
[They exit.]

Scene 2
=======
[Enter the old Duchess of York with the two
children of Clarence.]


BOY
Good grandam, tell us, is our father dead?

DUCHESS  No, boy.

DAUGHTER
Why do you weep so oft, and beat your breast,
And cry "O Clarence, my unhappy son"?

BOY
Why do you look on us and shake your head,
And call us orphans, wretches, castaways,
If that our noble father were alive?

DUCHESS
My pretty cousins, you mistake me both.
I do lament the sickness of the King,
As loath to lose him, not your father's death.
It were lost sorrow to wail one that's lost.

BOY
Then, you conclude, my grandam, he is dead.
The King mine uncle is to blame for it.
God will revenge it, whom I will importune
With earnest prayers, all to that effect.

DAUGHTER  And so will I.

DUCHESS
Peace, children, peace. The King doth love you
well.
Incapable and shallow innocents,
You cannot guess who caused your father's death.

BOY
Grandam, we can, for my good uncle Gloucester
Told me the King, provoked to it by the Queen,
Devised impeachments to imprison him;
And when my uncle told me so, he wept,
And pitied me, and kindly kissed my cheek,
Bade me rely on him as on my father,
And he would love me dearly as a child.

DUCHESS
Ah, that deceit should steal such gentle shape,
And with a virtuous visor hide deep vice.
He is my son, ay, and therein my shame,
Yet from my dugs he drew not this deceit.

BOY
Think you my uncle did dissemble, grandam?

DUCHESS  Ay, boy.

BOY
I cannot think it. Hark, what noise is this?

[Enter Queen Elizabeth with her hair about her ears,
Rivers and Dorset after her.]


QUEEN ELIZABETH
Ah, who shall hinder me to wail and weep,
To chide my fortune and torment myself?
I'll join with black despair against my soul
And to myself become an enemy.

DUCHESS
What means this scene of rude impatience?

QUEEN ELIZABETH
To make an act of tragic violence.
Edward, my lord, thy son, our king, is dead.
Why grow the branches when the root is gone?
Why wither not the leaves that want their sap?
If you will live, lament. If die, be brief,
That our swift-winged souls may catch the King's,
Or, like obedient subjects, follow him
To his new kingdom of ne'er-changing night.

DUCHESS
Ah, so much interest have I in thy sorrow
As I had title in thy noble husband.
I have bewept a worthy husband's death
And lived with looking on his images;
But now two mirrors of his princely semblance
Are cracked in pieces by malignant death,
And I, for comfort, have but one false glass
That grieves me when I see my shame in him.
Thou art a widow, yet thou art a mother,
And hast the comfort of thy children left,
But death hath snatched my husband from mine
arms
And plucked two crutches from my feeble hands,
Clarence and Edward. O, what cause have I,
Thine being but a moiety of my moan,
To overgo thy woes and drown thy cries!

BOY, [to Queen Elizabeth]
Ah, aunt, you wept not for our father's death.
How can we aid you with our kindred tears?

DAUGHTER, [to Queen Elizabeth]
Our fatherless distress was left unmoaned.
Your widow-dolor likewise be unwept!

QUEEN ELIZABETH
Give me no help in lamentation.
I am not barren to bring forth complaints.
All springs reduce their currents to mine eyes,
That I, being governed by the watery moon,
May send forth plenteous tears to drown the world.
Ah, for my husband, for my dear lord Edward!

CHILDREN
Ah, for our father, for our dear lord Clarence!

DUCHESS
Alas for both, both mine, Edward and Clarence!

QUEEN ELIZABETH
What stay had I but Edward? And he's gone.

CHILDREN
What stay had we but Clarence? And he's gone.

DUCHESS
What stays had I but they? And they are gone.

QUEEN ELIZABETH
Was never widow had so dear a loss.

CHILDREN
Were never orphans had so dear a loss.

DUCHESS
Was never mother had so dear a loss.
Alas, I am the mother of these griefs.
Their woes are parceled; mine is general.
She for an Edward weeps, and so do I;
I for a Clarence weep; so doth not she.
These babes for Clarence weep, and so do I;
I for an Edward weep; so do not they.
Alas, you three, on me, threefold distressed,
Pour all your tears. I am your sorrow's nurse,
And I will pamper it with lamentation.

DORSET, [to Queen Elizabeth]
Comfort, dear mother. God is much displeased
That you take with unthankfulness His doing.
In common worldly things, 'tis called ungrateful
With dull unwillingness to repay a debt
Which with a bounteous hand was kindly lent;
Much more to be thus opposite with heaven,
For it requires the royal debt it lent you.

RIVERS
Madam, bethink you, like a careful mother,
Of the young prince your son. Send straight for
him.
Let him be crowned. In him your comfort lives.
Drown desperate sorrow in dead Edward's grave
And plant your joys in living Edward's throne.

[Enter Richard, Duke of Gloucester, Buckingham, Lord
Stanley, Earl of Derby, Hastings, and Ratcliffe.]


RICHARD, [to Queen Elizabeth]
Sister, have comfort. All of us have cause
To wail the dimming of our shining star,
But none can help our harms by wailing them.--
Madam my mother, I do cry you mercy;
I did not see your Grace. Humbly on my knee
I crave your blessing.	[He kneels.]

DUCHESS
God bless thee, and put meekness in thy breast,
Love, charity, obedience, and true duty.

RICHARD, [standing]
Amen. [Aside.] And make me die a good old man!
That is the butt end of a mother's blessing;
I marvel that her Grace did leave it out.

BUCKINGHAM
You cloudy princes and heart-sorrowing peers
That bear this heavy mutual load of moan,
Now cheer each other in each other's love.
Though we have spent our harvest of this king,
We are to reap the harvest of his son.
The broken rancor of your high-swoll'n hates,
But lately splintered, knit, and joined together,
Must gently be preserved, cherished, and kept.
Meseemeth good that with some little train
Forthwith from Ludlow the young prince be fet
Hither to London, to be crowned our king.

RIVERS
Why "with some little train," my lord of
Buckingham?

BUCKINGHAM
Marry, my lord, lest by a multitude
The new-healed wound of malice should break out,
Which would be so much the more dangerous
By how much the estate is green and yet
ungoverned.
Where every horse bears his commanding rein
And may direct his course as please himself,
As well the fear of harm as harm apparent,
In my opinion, ought to be prevented.

RICHARD
I hope the King made peace with all of us;
And the compact is firm and true in me.

RIVERS
And so in me, and so, I think, in all.
Yet since it is but green, it should be put
To no apparent likelihood of breach,
Which haply by much company might be urged.
Therefore I say with noble Buckingham
That it is meet so few should fetch the Prince.

HASTINGS  And so say I.

RICHARD
Then be it so, and go we to determine
Who they shall be that straight shall post to
Ludlow.--
Madam, and you, my sister, will you go
To give your censures in this business?
[All but Buckingham and Richard exit.]

BUCKINGHAM
My lord, whoever journeys to the Prince,
For God's sake let not us two stay at home.
For by the way I'll sort occasion,
As index to the story we late talked of,
To part the Queen's proud kindred from the Prince.

RICHARD
My other self, my council's consistory,
My oracle, my prophet, my dear cousin,
I, as a child, will go by thy direction.
Toward Ludlow then, for we'll not stay behind.
[They exit.]

Scene 3
=======
[Enter one Citizen at one door, and another at the other.]


FIRST CITIZEN
Good morrow, neighbor, whither away so fast?

SECOND CITIZEN
I promise you I scarcely know myself.
Hear you the news abroad?

FIRST CITIZEN  Yes, that the King is dead.

SECOND CITIZEN
Ill news, by 'r Lady. Seldom comes the better.
I fear, I fear, 'twill prove a giddy world.

[Enter another Citizen.]


THIRD CITIZEN
Neighbors, God speed.

FIRST CITIZEN  Give you good morrow, sir.

THIRD CITIZEN
Doth the news hold of good King Edward's death?

SECOND CITIZEN
Ay, sir, it is too true, God help the while.

THIRD CITIZEN
Then, masters, look to see a troublous world.

FIRST CITIZEN
No, no, by God's good grace, his son shall reign.

THIRD CITIZEN
Woe to that land that's governed by a child.

SECOND CITIZEN
In him there is a hope of government,
Which, in his nonage, council under him,
And, in his full and ripened years, himself,
No doubt shall then, and till then, govern well.

FIRST CITIZEN
So stood the state when Henry the Sixth
Was crowned in Paris but at nine months old.

THIRD CITIZEN
Stood the state so? No, no, good friends, God wot,
For then this land was famously enriched
With politic grave counsel; then the King
Had virtuous uncles to protect his Grace.

FIRST CITIZEN
Why, so hath this, both by his father and mother.

THIRD CITIZEN
Better it were they all came by his father,
Or by his father there were none at all,
For emulation who shall now be nearest
Will touch us all too near if God prevent not.
O, full of danger is the Duke of Gloucester,
And the Queen's sons and brothers haught and
proud,
And were they to be ruled, and not to rule,
This sickly land might solace as before.

FIRST CITIZEN
Come, come, we fear the worst. All will be well.

THIRD CITIZEN
When clouds are seen, wise men put on their
cloaks;
When great leaves fall, then winter is at hand;
When the sun sets, who doth not look for night?
Untimely storms makes men expect a dearth.
All may be well; but if God sort it so,
'Tis more than we deserve or I expect.

SECOND CITIZEN
Truly, the hearts of men are full of fear.
You cannot reason almost with a man
That looks not heavily and full of dread.

THIRD CITIZEN
Before the days of change, still is it so.
By a divine instinct, men's minds mistrust
Ensuing danger, as by proof we see
The water swell before a boist'rous storm.
But leave it all to God. Whither away?

SECOND CITIZEN
Marry, we were sent for to the Justices.

THIRD CITIZEN
And so was I. I'll bear you company.
[They exit.]

Scene 4
=======
[Enter Archbishop, the young Duke of York,
Queen Elizabeth, and the Duchess of York.]


ARCHBISHOP
Last night, I hear, they lay at Stony Stratford,
And at Northampton they do rest tonight.
Tomorrow or next day they will be here.

DUCHESS
I long with all my heart to see the Prince.
I hope he is much grown since last I saw him.

QUEEN ELIZABETH
But I hear no; they say my son of York
Has almost overta'en him in his growth.

YORK
Ay, mother, but I would not have it so.

DUCHESS
Why, my good cousin? It is good to grow.

YORK
Grandam, one night as we did sit at supper,
My uncle Rivers talked how I did grow
More than my brother. "Ay," quoth my uncle
Gloucester,
"Small herbs have grace; great weeds do grow
apace."
And since, methinks I would not grow so fast
Because sweet flowers are slow and weeds make
haste.

DUCHESS
Good faith, good faith, the saying did not hold
In him that did object the same to thee!
He was the wretched'st thing when he was young,
So long a-growing and so leisurely,
That if his rule were true, he should be gracious.

YORK
And so no doubt he is, my gracious madam.

DUCHESS
I hope he is, but yet let mothers doubt.

YORK
Now, by my troth, if I had been remembered,
I could have given my uncle's Grace a flout
To touch his growth nearer than he touched mine.

DUCHESS
How, my young York? I prithee let me hear it.

YORK
Marry, they say my uncle grew so fast
That he could gnaw a crust at two hours old.
'Twas full two years ere I could get a tooth.
Grandam, this would have been a biting jest.

DUCHESS
I prithee, pretty York, who told thee this?

YORK  Grandam, his nurse.

DUCHESS
His nurse? Why, she was dead ere thou wast born.

YORK
If 'twere not she, I cannot tell who told me.

QUEEN ELIZABETH
A parlous boy! Go to, you are too shrewd.

DUCHESS
Good madam, be not angry with the child.

QUEEN ELIZABETH  Pitchers have ears.

[Enter a Messenger.]


ARCHBISHOP  Here comes a messenger.--What news?

MESSENGER
Such news, my lord, as grieves me to report.

QUEEN ELIZABETH  How doth the Prince?

MESSENGER  Well, madam, and in health.

DUCHESS  What is thy news?

MESSENGER
Lord Rivers and Lord Grey are sent to Pomfret,
And, with them, Sir Thomas Vaughan, prisoners.

DUCHESS  Who hath committed them?

MESSENGER
The mighty dukes, Gloucester and Buckingham.

ARCHBISHOP  For what offense?

MESSENGER
The sum of all I can, I have disclosed.
Why, or for what, the nobles were committed
Is all unknown to me, my gracious lord.

QUEEN ELIZABETH
Ay me! I see the ruin of my house.
The tiger now hath seized the gentle hind.
Insulting tyranny begins to jut
Upon the innocent and aweless throne.
Welcome, destruction, blood, and massacre.
I see, as in a map, the end of all.

DUCHESS
Accursed and unquiet wrangling days,
How many of you have mine eyes beheld?
My husband lost his life to get the crown,
And often up and down my sons were tossed
For me to joy, and weep, their gain and loss.
And being seated, and domestic broils
Clean overblown, themselves the conquerors
Make war upon themselves, brother to brother,
Blood to blood, self against self. O, preposterous
And frantic outrage, end thy damned spleen,
Or let me die, to look on Earth no more.

QUEEN ELIZABETH, [to York]
Come, come, my boy. We will to sanctuary.--
Madam, farewell.

DUCHESS  Stay, I will go with you.

QUEEN ELIZABETH
You have no cause.

ARCHBISHOP, [to Queen Elizabeth]  My gracious lady, go,
And thither bear your treasure and your goods.
For my part, I'll resign unto your Grace
The seal I keep; and so betide to me
As well I tender you and all of yours.
Go. I'll conduct you to the sanctuary.
[They exit.]


ACT 3
=====

Scene 1
=======
[The trumpets sound. Enter young Prince Edward,
Richard Duke of Gloucester, Buckingham,
the Cardinal, Catesby, and others.]


BUCKINGHAM
Welcome, sweet prince, to London, to your chamber.

RICHARD, [to Prince]
Welcome, dear cousin, my thoughts' sovereign.
The weary way hath made you melancholy.

PRINCE
No, uncle, but our crosses on the way
Have made it tedious, wearisome, and heavy.
I want more uncles here to welcome me.

RICHARD
Sweet prince, the untainted virtue of your years
Hath not yet dived into the world's deceit;
Nor more can you distinguish of a man
Than of his outward show, which, God He knows,
Seldom or never jumpeth with the heart.
Those uncles which you want were dangerous.
Your Grace attended to their sugared words
But looked not on the poison of their hearts.
God keep you from them, and from such false
friends.

PRINCE
God keep me from false friends, but they were none.

RICHARD
My lord, the Mayor of London comes to greet you.

[Enter Lord Mayor with others.]


MAYOR
God bless your Grace with health and happy days.

PRINCE
I thank you, good my lord, and thank you all.--
I thought my mother and my brother York
Would long ere this have met us on the way.
Fie, what a slug is Hastings that he comes not
To tell us whether they will come or no!

[Enter Lord Hastings.]


BUCKINGHAM
And in good time here comes the sweating lord.

PRINCE
Welcome, my lord. What, will our mother come?

HASTINGS
On what occasion God He knows, not I,
The Queen your mother and your brother York
Have taken sanctuary. The tender prince
Would fain have come with me to meet your Grace,
But by his mother was perforce withheld.

BUCKINGHAM
Fie, what an indirect and peevish course
Is this of hers!--Lord Cardinal, will your Grace
Persuade the Queen to send the Duke of York
Unto his princely brother presently?--
If she deny, Lord Hastings, go with him,
And from her jealous arms pluck him perforce.

CARDINAL
My lord of Buckingham, if my weak oratory
Can from his mother win the Duke of York,
Anon expect him here; but if she be obdurate
To mild entreaties, God in heaven forbid
We should infringe the holy privilege
Of blessed sanctuary! Not for all this land
Would I be guilty of so deep a sin.

BUCKINGHAM
You are too senseless obstinate, my lord,
Too ceremonious and traditional.
Weigh it but with the grossness of this age,
You break not sanctuary in seizing him.
The benefit thereof is always granted
To those whose dealings have deserved the place
And those who have the wit to claim the place.
This prince hath neither claimed it nor deserved it
And therefore, in mine opinion, cannot have it.
Then taking him from thence that is not there,
You break no privilege nor charter there.
Oft have I heard of sanctuary men,
But sanctuary children, never till now.

CARDINAL
My lord, you shall o'errule my mind for once.--
Come on, Lord Hastings, will you go with me?

HASTINGS  I go, my lord.

PRINCE
Good lords, make all the speedy haste you may.
[The Cardinal and Hastings exit.]
Say, uncle Gloucester, if our brother come,
Where shall we sojourn till our coronation?

RICHARD
Where it seems best unto your royal self.
If I may counsel you, some day or two
Your Highness shall repose you at the Tower;
Then where you please and shall be thought most fit
For your best health and recreation.

PRINCE
I do not like the Tower, of any place.--
Did Julius Caesar build that place, my lord?

BUCKINGHAM
He did, my gracious lord, begin that place,
Which, since, succeeding ages have re-edified.

PRINCE
Is it upon record, or else reported
Successively from age to age, he built it?

BUCKINGHAM  Upon record, my gracious lord.

PRINCE
But say, my lord, it were not registered,
Methinks the truth should live from age to age,
As 'twere retailed to all posterity,
Even to the general all-ending day.

RICHARD, [aside]
So wise so young, they say, do never live long.

PRINCE  What say you, uncle?

RICHARD
I say, without characters fame lives long.
[Aside.] Thus, like the formal Vice, Iniquity,
I moralize two meanings in one word.

PRINCE
That Julius Caesar was a famous man.
With what his valor did enrich his wit,
His wit set down to make his valor live.
Death makes no conquest of this conqueror,
For now he lives in fame, though not in life.
I'll tell you what, my cousin Buckingham--

BUCKINGHAM  What, my gracious lord?

PRINCE
An if I live until I be a man,
I'll win our ancient right in France again
Or die a soldier, as I lived a king.

RICHARD, [aside]
Short summers lightly have a forward spring.

[Enter young Duke of York, Hastings, and the
Cardinal.]


BUCKINGHAM
Now in good time here comes the Duke of York.

PRINCE
Richard of York, how fares our loving brother?

YORK
Well, my dread lord--so must I call you now.

PRINCE
Ay, brother, to our grief, as it is yours.
Too late he died that might have kept that title,
Which by his death hath lost much majesty.

RICHARD
How fares our cousin, noble lord of York?

YORK
I thank you, gentle uncle. O my lord,
You said that idle weeds are fast in growth.
The Prince my brother hath outgrown me far.

RICHARD
He hath, my lord.

YORK  And therefore is he idle?

RICHARD
O my fair cousin, I must not say so.

YORK
Then he is more beholding to you than I.

RICHARD
He may command me as my sovereign,
But you have power in me as in a kinsman.

YORK
I pray you, uncle, give me this dagger.

RICHARD
My dagger, little cousin? With all my heart.

PRINCE  A beggar, brother?

YORK
Of my kind uncle, that I know will give,
And being but a toy, which is no grief to give.

RICHARD
A greater gift than that I'll give my cousin.

YORK
A greater gift? O, that's the sword to it.

RICHARD
Ay, gentle cousin, were it light enough.

YORK
O, then I see you will part but with light gifts.
In weightier things you'll say a beggar nay.

RICHARD
It is too heavy for your Grace to wear.

YORK
I weigh it lightly, were it heavier.

RICHARD
What, would you have my weapon, little lord?

YORK
I would, that I might thank you as you call me.

RICHARD  How?

YORK  Little.

PRINCE
My lord of York will still be cross in talk.
Uncle, your Grace knows how to bear with him.

YORK
You mean, to bear me, not to bear with me.--
Uncle, my brother mocks both you and me.
Because that I am little, like an ape,
He thinks that you should bear me on your
shoulders.

BUCKINGHAM, [aside]
With what a sharp-provided wit he reasons!
To mitigate the scorn he gives his uncle,
He prettily and aptly taunts himself.
So cunning and so young is wonderful.

RICHARD, [to Prince]
My lord, will 't please you pass along?
Myself and my good cousin Buckingham
Will to your mother, to entreat of her
To meet you at the Tower and welcome you.

YORK, [to Prince]
What, will you go unto the Tower, my lord?

PRINCE
My Lord Protector needs will have it so.

YORK
I shall not sleep in quiet at the Tower.

RICHARD  Why, what should you fear?

YORK
Marry, my uncle Clarence' angry ghost.
My grandam told me he was murdered there.

PRINCE  I fear no uncles dead.

RICHARD  Nor none that live, I hope.

PRINCE
An if they live, I hope I need not fear.
[To York.] But come, my lord. With a heavy heart,
Thinking on them, go I unto the Tower.
[A sennet. Prince Edward, the Duke of York,
and Hastings exit. Richard, Buckingham,
and Catesby remain.]

BUCKINGHAM, [to Richard]
Think you, my lord, this little prating York
Was not incensed by his subtle mother
To taunt and scorn you thus opprobriously?

RICHARD
No doubt, no doubt. O, 'tis a parlous boy,
Bold, quick, ingenious, forward, capable.
He is all the mother's, from the top to toe.

BUCKINGHAM
Well, let them rest.--Come hither, Catesby.
Thou art sworn as deeply to effect what we intend
As closely to conceal what we impart.
Thou knowest our reasons, urged upon the way.
What thinkest thou? Is it not an easy matter
To make William Lord Hastings of our mind
For the installment of this noble duke
In the seat royal of this famous isle?

CATESBY
He, for his father's sake, so loves the Prince
That he will not be won to aught against him.

BUCKINGHAM
What think'st thou then of Stanley? Will not he?

CATESBY
He will do all in all as Hastings doth.

BUCKINGHAM
Well then, no more but this: go, gentle Catesby,
And, as it were far off, sound thou Lord Hastings
How he doth stand affected to our purpose
And summon him tomorrow to the Tower
To sit about the coronation.
If thou dost find him tractable to us,
Encourage him and tell him all our reasons.
If he be leaden, icy, cold, unwilling,
Be thou so too, and so break off the talk,
And give us notice of his inclination;
For we tomorrow hold divided councils,
Wherein thyself shalt highly be employed.

RICHARD
Commend me to Lord William. Tell him, Catesby,
His ancient knot of dangerous adversaries
Tomorrow are let blood at Pomfret Castle,
And bid my lord, for joy of this good news,
Give Mistress Shore one gentle kiss the more.

BUCKINGHAM
Good Catesby, go effect this business soundly.

CATESBY
My good lords both, with all the heed I can.

RICHARD
Shall we hear from you, Catesby, ere we sleep?

CATESBY  You shall, my lord.

RICHARD
At Crosby House, there shall you find us both.
[Catesby exits.]

BUCKINGHAM
Now, my lord, what shall we do if we perceive
Lord Hastings will not yield to our complots?

RICHARD
Chop off his head. Something we will determine.
And look when I am king, claim thou of me
The earldom of Hereford, and all the movables
Whereof the King my brother was possessed.

BUCKINGHAM
I'll claim that promise at your Grace's hand.

RICHARD
And look to have it yielded with all kindness.
Come, let us sup betimes, that afterwards
We may digest our complots in some form.
[They exit.]

Scene 2
=======
[Enter a Messenger to the door of Hastings.]


MESSENGER, [knocking]  My lord, my lord.

HASTINGS, [within]  Who knocks?

MESSENGER  One from the Lord Stanley.

HASTINGS, [within]  What is 't o'clock?

MESSENGER  Upon the stroke of four.

[Enter Lord Hastings.]


HASTINGS
Cannot my Lord Stanley sleep these tedious nights?

MESSENGER
So it appears by that I have to say.
First, he commends him to your noble self.

HASTINGS  What then?

MESSENGER
Then certifies your Lordship that this night
He dreamt the boar had razed off his helm.
Besides, he says there are two councils kept,
And that may be determined at the one
Which may make you and him to rue at th' other.
Therefore he sends to know your Lordship's
pleasure,
If you will presently take horse with him
And with all speed post with him toward the north
To shun the danger that his soul divines.

HASTINGS
Go, fellow, go. Return unto thy lord.
Bid him not fear the separated council.
His Honor and myself are at the one,
And at the other is my good friend Catesby,
Where nothing can proceed that toucheth us
Whereof I shall not have intelligence.
Tell him his fears are shallow, without instance.
And for his dreams, I wonder he's so simple
To trust the mock'ry of unquiet slumbers.
To fly the boar before the boar pursues
Were to incense the boar to follow us
And make pursuit where he did mean no chase.
Go, bid thy master rise and come to me,
And we will both together to the Tower,
Where he shall see the boar will use us kindly.

MESSENGER
I'll go, my lord, and tell him what you say.	[He exits.]

[Enter Catesby.]


CATESBY
Many good morrows to my noble lord.

HASTINGS
Good morrow, Catesby. You are early stirring.
What news, what news in this our tott'ring state?

CATESBY
It is a reeling world indeed, my lord,
And I believe will never stand upright
Till Richard wear the garland of the realm.

HASTINGS
How "wear the garland"? Dost thou mean the
crown?

CATESBY  Ay, my good lord.

HASTINGS
I'll have this crown of mine cut from my shoulders
Before I'll see the crown so foul misplaced.
But canst thou guess that he doth aim at it?

CATESBY
Ay, on my life, and hopes to find you forward
Upon his party for the gain thereof;
And thereupon he sends you this good news,
That this same very day your enemies,
The kindred of the Queen, must die at Pomfret.

HASTINGS
Indeed, I am no mourner for that news,
Because they have been still my adversaries.
But that I'll give my voice on Richard's side
To bar my master's heirs in true descent,
God knows I will not do it, to the death.

CATESBY
God keep your Lordship in that gracious mind.

HASTINGS
But I shall laugh at this a twelve-month hence,
That they which brought me in my master's hate,
I live to look upon their tragedy.
Well, Catesby, ere a fortnight make me older
I'll send some packing that yet think not on 't.

CATESBY
'Tis a vile thing to die, my gracious lord,
When men are unprepared and look not for it.

HASTINGS
O monstrous, monstrous! And so falls it out
With Rivers, Vaughan, Grey; and so 'twill do
With some men else that think themselves as safe
As thou and I, who, as thou know'st, are dear
To princely Richard and to Buckingham.

CATESBY
The Princes both make high account of you--
[Aside.] For they account his head upon the Bridge.

HASTINGS
I know they do, and I have well deserved it.

[Enter Lord Stanley.]

Come on, come on. Where is your boar-spear, man?
Fear you the boar and go so unprovided?

STANLEY
My lord, good morrow.--Good morrow, Catesby.--
You may jest on, but, by the Holy Rood,
I do not like these several councils, I.

HASTINGS
My lord, I hold my life as dear as you do yours,
And never in my days, I do protest,
Was it so precious to me as 'tis now.
Think you but that I know our state secure,
I would be so triumphant as I am?

STANLEY
The lords at Pomfret, when they rode from London,
Were jocund and supposed their states were sure,
And they indeed had no cause to mistrust;
But yet you see how soon the day o'ercast.
This sudden stab of rancor I misdoubt.
Pray God, I say, I prove a needless coward!
What, shall we toward the Tower? The day is spent.

HASTINGS
Come, come. Have with you. Wot you what, my lord?
Today the lords you talked of are beheaded.

STANLEY
They, for their truth, might better wear their heads
Than some that have accused them wear their hats.
But come, my lord, let's away.

[Enter a Pursuivant.]


HASTINGS
Go on before. I'll talk with this good fellow.
[Lord Stanley and Catesby exit.]
How now, sirrah? How goes the world with thee?

PURSUIVANT
The better that your Lordship please to ask.

HASTINGS
I tell thee, man, 'tis better with me now
Than when thou met'st me last where now we meet.
Then was I going prisoner to the Tower
By the suggestion of the Queen's allies.
But now, I tell thee--keep it to thyself--
This day those enemies are put to death,
And I in better state than e'er I was.

PURSUIVANT
God hold it, to your Honor's good content!

HASTINGS
Gramercy, fellow. There, drink that for me.
[Throws him his purse.]

PURSUIVANT  I thank your Honor.	[Pursuivant exits.]

[Enter a Priest.]


PRIEST
Well met, my lord. I am glad to see your Honor.

HASTINGS
I thank thee, good Sir John, with all my heart.
I am in your debt for your last exercise.
Come the next sabbath, and I will content you.

PRIEST  I'll wait upon your Lordship.	[Priest exits.]

[Enter Buckingham.]


BUCKINGHAM
What, talking with a priest, Lord Chamberlain?
Your friends at Pomfret, they do need the priest;
Your Honor hath no shriving work in hand.

HASTINGS
Good faith, and when I met this holy man,
The men you talk of came into my mind.
What, go you toward the Tower?

BUCKINGHAM
I do, my lord, but long I cannot stay there.
I shall return before your Lordship thence.

HASTINGS
Nay, like enough, for I stay dinner there.

BUCKINGHAM, [aside]
And supper too, although thou know'st it not.--
Come, will you go?

HASTINGS  I'll wait upon your Lordship.
[They exit.]

Scene 3
=======
[Enter Sir Richard Ratcliffe, with Halberds, carrying the
nobles Rivers, Grey, and Vaughan to death at Pomfret.]


RIVERS
Sir Richard Ratcliffe, let me tell thee this:
Today shalt thou behold a subject die
For truth, for duty, and for loyalty.

GREY, [to Ratcliffe]
God bless the Prince from all the pack of you!
A knot you are of damned bloodsuckers.

VAUGHAN, [to Ratcliffe]
You live that shall cry woe for this hereafter.

RATCLIFFE
Dispatch. The limit of your lives is out.

RIVERS
O Pomfret, Pomfret! O thou bloody prison,
Fatal and ominous to noble peers!
Within the guilty closure of thy walls,
Richard the Second here was hacked to death,
And, for more slander to thy dismal seat,
We give to thee our guiltless blood to drink.

GREY
Now Margaret's curse is fall'n upon our heads,
When she exclaimed on Hastings, you, and I,
For standing by when Richard stabbed her son.

RIVERS
Then cursed she Richard. Then cursed she
Buckingham.
Then cursed she Hastings. O, remember, God,
To hear her prayer for them as now for us!
And for my sister and her princely sons,
Be satisfied, dear God, with our true blood,
Which, as thou know'st, unjustly must be spilt.

RATCLIFFE
Make haste. The hour of death is expiate.

RIVERS
Come, Grey. Come, Vaughan. Let us here embrace.
[They embrace.]
Farewell until we meet again in heaven.
[They exit.]

Scene 4
=======
[Enter Buckingham, Lord Stanley, Earl of Derby,
Hastings, Bishop of Ely, Norfolk, Ratcliffe, Lovell, with
others, at a table.]


HASTINGS
Now, noble peers, the cause why we are met
Is to determine of the coronation.
In God's name, speak. When is the royal day?

BUCKINGHAM
Is all things ready for the royal time?

STANLEY
It is, and wants but nomination.

ELY
Tomorrow, then, I judge a happy day.

BUCKINGHAM
Who knows the Lord Protector's mind herein?
Who is most inward with the noble duke?

ELY
Your Grace, we think, should soonest know his
mind.

BUCKINGHAM
We know each other's faces; for our hearts,
He knows no more of mine than I of yours,
Or I of his, my lord, than you of mine.--
Lord Hastings, you and he are near in love.

HASTINGS
I thank his Grace, I know he loves me well.
But for his purpose in the coronation,
I have not sounded him, nor he delivered
His gracious pleasure any way therein.
But you, my honorable lords, may name the time,
And in the Duke's behalf I'll give my voice,
Which I presume he'll take in gentle part.

[Enter Richard, Duke of Gloucester.]


ELY
In happy time here comes the Duke himself.

RICHARD
My noble lords and cousins all, good morrow.
I have been long a sleeper; but I trust
My absence doth neglect no great design
Which by my presence might have been concluded.

BUCKINGHAM
Had you not come upon your cue, my lord,
William Lord Hastings had pronounced your part--
I mean your voice for crowning of the King.

RICHARD
Than my Lord Hastings no man might be bolder.
His Lordship knows me well and loves me well.--
My lord of Ely, when I was last in Holborn
I saw good strawberries in your garden there;
I do beseech you, send for some of them.

ELY
Marry and will, my lord, with all my heart.
[Exit Bishop of Ely.]

RICHARD
Cousin of Buckingham, a word with you.
[They move aside.]
Catesby hath sounded Hastings in our business
And finds the testy gentleman so hot
That he will lose his head ere give consent
His master's child, as worshipfully he terms it,
Shall lose the royalty of England's throne.

BUCKINGHAM
Withdraw yourself awhile. I'll go with you.
[Richard and Buckingham exit.]

STANLEY
We have not yet set down this day of triumph.
Tomorrow, in my judgment, is too sudden,
For I myself am not so well provided
As else I would be, were the day prolonged.

[Enter the Bishop of Ely.]


ELY
Where is my lord the Duke of Gloucester?
I have sent for these strawberries.

HASTINGS
His Grace looks cheerfully and smooth this
morning.
There's some conceit or other likes him well
When that he bids good morrow with such spirit.
I think there's never a man in Christendom
Can lesser hide his love or hate than he,
For by his face straight shall you know his heart.

STANLEY
What of his heart perceive you in his face
By any livelihood he showed today?

HASTINGS
Marry, that with no man here he is offended,
For were he, he had shown it in his looks.

[Enter Richard and Buckingham.]


RICHARD
I pray you all, tell me what they deserve
That do conspire my death with devilish plots
Of damned witchcraft, and that have prevailed
Upon my body with their hellish charms?

HASTINGS
The tender love I bear your Grace, my lord,
Makes me most forward in this princely presence
To doom th' offenders, whosoe'er they be.
I say, my lord, they have deserved death.

RICHARD
Then be your eyes the witness of their evil.
[He shows his arm.]
Look how I am bewitched! Behold mine arm
Is like a blasted sapling withered up;
And this is Edward's wife, that monstrous witch,
Consorted with that harlot, strumpet Shore,
That by their witchcraft thus have marked me.

HASTINGS
If they have done this deed, my noble lord--

RICHARD
If? Thou protector of this damned strumpet,
Talk'st thou to me of "ifs"? Thou art a traitor.--
Off with his head. Now by Saint Paul I swear
I will not dine until I see the same.--
Lovell and Ratcliffe, look that it be done.--
The rest that love me, rise and follow me.
[They exit. Lovell and Ratcliffe remain,
with the Lord Hastings.]

HASTINGS
Woe, woe for England! Not a whit for me,
For I, too fond, might have prevented this.
Stanley did dream the boar did raze his helm,
And I did scorn it and disdain to fly.
Three times today my foot-cloth horse did stumble,
And started when he looked upon the Tower,
As loath to bear me to the slaughterhouse.
O, now I need the priest that spake to me!
I now repent I told the pursuivant,
As too triumphing, how mine enemies
Today at Pomfret bloodily were butchered,
And I myself secure in grace and favor.
O Margaret, Margaret, now thy heavy curse
Is lighted on poor Hastings' wretched head.

RATCLIFFE
Come, come, dispatch. The Duke would be at
dinner.
Make a short shrift. He longs to see your head.

HASTINGS
O momentary grace of mortal men,
Which we more hunt for than the grace of God!
Who builds his hope in air of your good looks
Lives like a drunken sailor on a mast,
Ready with every nod to tumble down
Into the fatal bowels of the deep.

LOVELL
Come, come, dispatch. 'Tis bootless to exclaim.

HASTINGS
O bloody Richard! Miserable England,
I prophesy the fearfull'st time to thee
That ever wretched age hath looked upon.--
Come, lead me to the block. Bear him my head.
They smile at me who shortly shall be dead.
[They exit.]

Scene 5
=======
[Enter Richard and Buckingham, in rotten armor,
marvelous ill-favored.]


RICHARD
Come, cousin, canst thou quake and change thy
color,
Murder thy breath in middle of a word,
And then again begin, and stop again,
As if thou were distraught and mad with terror?

BUCKINGHAM
Tut, I can counterfeit the deep tragedian,
Speak, and look back, and pry on every side,
Tremble and start at wagging of a straw,
Intending deep suspicion. Ghastly looks
Are at my service, like enforced smiles,
And both are ready, in their offices,
At any time to grace my stratagems.
But what, is Catesby gone?

RICHARD
He is; and see he brings the Mayor along.

[Enter the Mayor and Catesby.]


BUCKINGHAM  Lord Mayor--

RICHARD  Look to the drawbridge there!

BUCKINGHAM  Hark, a drum!

RICHARD  Catesby, o'erlook the walls.
[Catesby exits.]

BUCKINGHAM  Lord Mayor, the reason we have sent--

RICHARD
Look back! Defend thee! Here are enemies.

BUCKINGHAM
God and our innocence defend and guard us!

[Enter Lovell and Ratcliffe, with Hastings' head.]


RICHARD
Be patient. They are friends, Ratcliffe and Lovell.

LOVELL
Here is the head of that ignoble traitor,
The dangerous and unsuspected Hastings.

RICHARD
So dear I loved the man that I must weep.
I took him for the plainest harmless creature
That breathed upon the Earth a Christian;
Made him my book, wherein my soul recorded
The history of all her secret thoughts.
So smooth he daubed his vice with show of virtue
That, his apparent open guilt omitted--
I mean his conversation with Shore's wife--
He lived from all attainder of suspects.

BUCKINGHAM
Well, well, he was the covert'st sheltered traitor
That ever lived.--
Would you imagine, or almost believe,
Were 't not that by great preservation
We live to tell it, that the subtle traitor
This day had plotted, in the council house,
To murder me and my good lord of Gloucester?

MAYOR  Had he done so?

RICHARD
What, think you we are Turks or infidels?
Or that we would, against the form of law,
Proceed thus rashly in the villain's death,
But that the extreme peril of the case,
The peace of England, and our persons' safety
Enforced us to this execution?

MAYOR
Now fair befall you! He deserved his death,
And your good Graces both have well proceeded
To warn false traitors from the like attempts.

BUCKINGHAM
I never looked for better at his hands
After he once fell in with Mistress Shore.
Yet had we not determined he should die
Until your Lordship came to see his end
(Which now the loving haste of these our friends,
Something against our meanings, have prevented),
Because, my lord, I would have had you heard
The traitor speak and timorously confess
The manner and the purpose of his treasons,
That you might well have signified the same
Unto the citizens, who haply may
Misconster us in him, and wail his death.

MAYOR
But, my good lord, your Graces' words shall serve
As well as I had seen and heard him speak;
And do not doubt, right noble princes both,
But I'll acquaint our duteous citizens
With all your just proceedings in this case.

RICHARD
And to that end we wished your Lordship here,
T' avoid the censures of the carping world.

BUCKINGHAM
Which since you come too late of our intent,
Yet witness what you hear we did intend.
And so, my good Lord Mayor, we bid farewell.
[Mayor exits.]

RICHARD
Go after, after, cousin Buckingham.
The Mayor towards Guildhall hies him in all post.
There, at your meetest vantage of the time,
Infer the bastardy of Edward's children.
Tell them how Edward put to death a citizen
Only for saying he would make his son
Heir to the Crown--meaning indeed his house,
Which, by the sign thereof, was termed so.
Moreover, urge his hateful luxury
And bestial appetite in change of lust,
Which stretched unto their servants, daughters,
wives,
Even where his raging eye or savage heart,
Without control, lusted to make a prey.
Nay, for a need, thus far come near my person:
Tell them when that my mother went with child
Of that insatiate Edward, noble York
My princely father then had wars in France,
And, by true computation of the time,
Found that the issue was not his begot,
Which well appeared in his lineaments,
Being nothing like the noble duke my father.
Yet touch this sparingly, as 'twere far off,
Because, my lord, you know my mother lives.

BUCKINGHAM
Doubt not, my lord. I'll play the orator
As if the golden fee for which I plead
Were for myself. And so, my lord, adieu.

RICHARD
If you thrive well, bring them to Baynard's Castle,
Where you shall find me well accompanied
With reverend fathers and well-learned bishops.

BUCKINGHAM
I go; and towards three or four o'clock
Look for the news that the Guildhall affords.
[Buckingham exits.]

RICHARD
Go, Lovell, with all speed to Doctor Shaa.
[To Ratcliffe.] Go thou to Friar Penker. Bid them
both
Meet me within this hour at Baynard's Castle.
[Ratcliffe and Lovell exit.]
Now will I go to take some privy order
To draw the brats of Clarence out of sight,
And to give order that no manner person
Have any time recourse unto the Princes.
[He exits.]

Scene 6
=======
[Enter a Scrivener.]


SCRIVENER
Here is the indictment of the good Lord Hastings,
Which in a set hand fairly is engrossed,
That it may be today read o'er in Paul's.
And mark how well the sequel hangs together:
Eleven hours I have spent to write it over,
For yesternight by Catesby was it sent me;
The precedent was full as long a-doing,
And yet within these five hours Hastings lived,
Untainted, unexamined, free, at liberty.
Here's a good world the while! Who is so gross
That cannot see this palpable device?
Yet who so bold but says he sees it not?
Bad is the world, and all will come to naught
When such ill dealing must be seen in thought.
[He exits.]

Scene 7
=======
[Enter Richard and Buckingham at several doors.]


RICHARD
How now, how now? What say the citizens?

BUCKINGHAM
Now, by the holy mother of our Lord,
The citizens are mum, say not a word.

RICHARD
Touched you the bastardy of Edward's children?

BUCKINGHAM
I did; with his contract with Lady Lucy
And his contract by deputy in France;
Th' unsatiate greediness of his desire
And his enforcement of the city wives;
His tyranny for trifles; his own bastardy,
As being got, your father then in France,
And his resemblance being not like the Duke.
Withal, I did infer your lineaments,
Being the right idea of your father,
Both in your form and nobleness of mind;
Laid open all your victories in Scotland,
Your discipline in war, wisdom in peace,
Your bounty, virtue, fair humility;
Indeed, left nothing fitting for your purpose
Untouched or slightly handled in discourse.
And when mine oratory drew toward end,
I bid them that did love their country's good
Cry "God save Richard, England's royal king!"

RICHARD  And did they so?

BUCKINGHAM
No. So God help me, they spake not a word
But, like dumb statues or breathing stones,
Stared each on other and looked deadly pale;
Which when I saw, I reprehended them
And asked the Mayor what meant this willful silence.
His answer was, the people were not used
To be spoke to but by the Recorder.
Then he was urged to tell my tale again:
"Thus saith the Duke. Thus hath the Duke
inferred"--
But nothing spoke in warrant from himself.
When he had done, some followers of mine own,
At lower end of the hall, hurled up their caps,
And some ten voices cried "God save King Richard!"
And thus I took the vantage of those few.
"Thanks, gentle citizens and friends," quoth I.
"This general applause and cheerful shout
Argues your wisdoms and your love to Richard"--
And even here brake off and came away.

RICHARD
What tongueless blocks were they! Would they not
speak?
Will not the Mayor then and his brethren come?

BUCKINGHAM
The Mayor is here at hand. Intend some fear;
Be not you spoke with but by mighty suit.
And look you get a prayer book in your hand
And stand between two churchmen, good my lord,
For on that ground I'll make a holy descant.
And be not easily won to our requests.
Play the maid's part: still answer "nay," and take it.

RICHARD
I go. An if you plead as well for them
As I can say "nay" to thee for myself,
No doubt we bring it to a happy issue.
[Knocking within.]

BUCKINGHAM
Go, go, up to the leads. The Lord Mayor knocks.
[Richard exits.]

[Enter the Mayor and Citizens.]

Welcome, my lord. I dance attendance here.
I think the Duke will not be spoke withal.

[Enter Catesby.]

Now, Catesby, what says your lord to my request?

CATESBY
He doth entreat your Grace, my noble lord,
To visit him tomorrow or next day.
He is within, with two right reverend fathers,
Divinely bent to meditation,
And in no worldly suits would he be moved
To draw him from his holy exercise.

BUCKINGHAM
Return, good Catesby, to the gracious duke.
Tell him myself, the Mayor, and aldermen,
In deep designs, in matter of great moment
No less importing than our general good,
Are come to have some conference with his Grace.

CATESBY
I'll signify so much unto him straight.	[He exits.]

BUCKINGHAM
Ah ha, my lord, this prince is not an Edward!
He is not lolling on a lewd love-bed,
But on his knees at meditation;
Not dallying with a brace of courtesans,
But meditating with two deep divines;
Not sleeping, to engross his idle body,
But praying, to enrich his watchful soul.
Happy were England would this virtuous prince
Take on his Grace the sovereignty thereof.
But sure I fear we shall not win him to it.

MAYOR
Marry, God defend his Grace should say us nay.

BUCKINGHAM
I fear he will. Here Catesby comes again.

[Enter Catesby.]

Now, Catesby, what says his Grace?

CATESBY
He wonders to what end you have assembled
Such troops of citizens to come to him,
His Grace not being warned thereof before.
He fears, my lord, you mean no good to him.

BUCKINGHAM
Sorry I am my noble cousin should
Suspect me that I mean no good to him.
By heaven, we come to him in perfect love,
And so once more return and tell his Grace.
[Catesby exits.]
When holy and devout religious men
Are at their beads, 'tis much to draw them thence,
So sweet is zealous contemplation.

[Enter Richard aloft, between two Bishops.]
[Catesby reenters.]


MAYOR
See where his Grace stands, 'tween two clergymen.

BUCKINGHAM
Two props of virtue for a Christian prince,
To stay him from the fall of vanity;
And, see, a book of prayer in his hand,
True ornaments to know a holy man.--
Famous Plantagenet, most gracious prince,
Lend favorable ear to our requests,
And pardon us the interruption
Of thy devotion and right Christian zeal.

RICHARD
My lord, there needs no such apology.
I do beseech your Grace to pardon me,
Who, earnest in the service of my God,
Deferred the visitation of my friends.
But, leaving this, what is your Grace's pleasure?

BUCKINGHAM
Even that, I hope, which pleaseth God above
And all good men of this ungoverned isle.

RICHARD
I do suspect I have done some offense
That seems disgracious in the city's eye,
And that you come to reprehend my ignorance.

BUCKINGHAM
You have, my lord. Would it might please your
Grace,
On our entreaties, to amend your fault.

RICHARD
Else wherefore breathe I in a Christian land?

BUCKINGHAM
Know, then, it is your fault that you resign
The supreme seat, the throne majestical,
The sceptered office of your ancestors,
Your state of fortune, and your due of birth,
The lineal glory of your royal house,
To the corruption of a blemished stock,
Whiles in the mildness of your sleepy thoughts,
Which here we waken to our country's good,
The noble isle doth want her proper limbs--
Her face defaced with scars of infamy,
Her royal stock graft with ignoble plants,
And almost shouldered in the swallowing gulf
Of dark forgetfulness and deep oblivion;
Which to recure, we heartily solicit
Your gracious self to take on you the charge
And kingly government of this your land,
Not as Protector, steward, substitute,
Or lowly factor for another's gain,
But as successively, from blood to blood,
Your right of birth, your empery, your own.
For this, consorted with the citizens,
Your very worshipful and loving friends,
And by their vehement instigation,
In this just cause come I to move your Grace.

RICHARD
I cannot tell if to depart in silence
Or bitterly to speak in your reproof
Best fitteth my degree or your condition.
If not to answer, you might haply think
Tongue-tied ambition, not replying, yielded
To bear the golden yoke of sovereignty,
Which fondly you would here impose on me.
If to reprove you for this suit of yours,
So seasoned with your faithful love to me,
Then on the other side I checked my friends.
Therefore, to speak, and to avoid the first,
And then, in speaking, not to incur the last,
Definitively thus I answer you:
Your love deserves my thanks, but my desert
Unmeritable shuns your high request.
First, if all obstacles were cut away
And that my path were even to the crown
As the ripe revenue and due of birth,
Yet so much is my poverty of spirit,
So mighty and so many my defects,
That I would rather hide me from my greatness,
Being a bark to brook no mighty sea,
Than in my greatness covet to be hid
And in the vapor of my glory smothered.
But, God be thanked, there is no need of me,
And much I need to help you, were there need.
The royal tree hath left us royal fruit,
Which, mellowed by the stealing hours of time,
Will well become the seat of majesty,
And make, no doubt, us happy by his reign.
On him I lay that you would lay on me,
The right and fortune of his happy stars,
Which God defend that I should wring from him.

BUCKINGHAM
My lord, this argues conscience in your Grace,
But the respects thereof are nice and trivial,
All circumstances well considered.
You say that Edward is your brother's son;
So say we too, but not by Edward's wife.
For first was he contract to Lady Lucy--
Your mother lives a witness to his vow--
And afterward by substitute betrothed
To Bona, sister to the King of France.
These both put off, a poor petitioner,
A care-crazed mother to a many sons,
A beauty-waning and distressed widow,
Even in the afternoon of her best days,
Made prize and purchase of his wanton eye,
Seduced the pitch and height of his degree
To base declension and loathed bigamy.
By her in his unlawful bed he got
This Edward, whom our manners call "the Prince."
More bitterly could I expostulate,
Save that, for reverence to some alive,
I give a sparing limit to my tongue.
Then, good my lord, take to your royal self
This proffered benefit of dignity,
If not to bless us and the land withal,
Yet to draw forth your noble ancestry
From the corruption of abusing times
Unto a lineal, true-derived course.

MAYOR
Do, good my lord. Your citizens entreat you.

BUCKINGHAM
Refuse not, mighty lord, this proffered love.

CATESBY
O, make them joyful. Grant their lawful suit.

RICHARD
Alas, why would you heap this care on me?
I am unfit for state and majesty.
I do beseech you, take it not amiss;
I cannot, nor I will not, yield to you.

BUCKINGHAM
If you refuse it, as in love and zeal
Loath to depose the child, your brother's son--
As well we know your tenderness of heart
And gentle, kind, effeminate remorse,
Which we have noted in you to your kindred
And equally indeed to all estates--
Yet know, whe'er you accept our suit or no,
Your brother's son shall never reign our king,
But we will plant some other in the throne,
To the disgrace and downfall of your house.
And in this resolution here we leave you.--
Come, citizens. Zounds, I'll entreat no more.

RICHARD
O, do not swear, my lord of Buckingham!
[Buckingham and some others exit.]

CATESBY
Call him again, sweet prince. Accept their suit.
If you deny them, all the land will rue it.

RICHARD
Will you enforce me to a world of cares?
Call them again. I am not made of stones,
But penetrable to your kind entreaties,
Albeit against my conscience and my soul.

[Enter Buckingham and the rest.]

Cousin of Buckingham and sage, grave men,
Since you will buckle Fortune on my back,
To bear her burden, whe'er I will or no,
I must have patience to endure the load;
But if black scandal or foul-faced reproach
Attend the sequel of your imposition,
Your mere enforcement shall acquittance me
From all the impure blots and stains thereof,
For God doth know, and you may partly see,
How far I am from the desire of this.

MAYOR
God bless your Grace! We see it and will say it.

RICHARD
In saying so, you shall but say the truth.

BUCKINGHAM
Then I salute you with this royal title:
Long live Richard, England's worthy king!

ALL  Amen.

BUCKINGHAM
Tomorrow may it please you to be crowned?

RICHARD
Even when you please, for you will have it so.

BUCKINGHAM
Tomorrow, then, we will attend your Grace,
And so most joyfully we take our leave.

RICHARD, [to the Bishops]
Come, let us to our holy work again.--
Farewell, my cousin. Farewell, gentle friends.
[They exit.]


ACT 4
=====

Scene 1
=======
[Enter Queen Elizabeth, with the Duchess of York, and
the Lord Marquess of Dorset, at one door; Anne,
Duchess of Gloucester with Clarence's daughter, at
another door.]


DUCHESS
Who meets us here? My niece Plantagenet
Led in the hand of her kind aunt of Gloucester?
Now, for my life, she's wandering to the Tower,
On pure heart's love, to greet the tender prince.--
Daughter, well met.

ANNE  God give your Graces both
A happy and a joyful time of day.

QUEEN ELIZABETH
As much to you, good sister. Whither away?

ANNE
No farther than the Tower, and, as I guess,
Upon the like devotion as yourselves,
To gratulate the gentle princes there.

QUEEN ELIZABETH
Kind sister, thanks. We'll enter all together.

[Enter Brakenbury, the Lieutenant.]

And in good time here the Lieutenant comes.--
Master Lieutenant, pray you, by your leave,
How doth the Prince and my young son of York?

BRAKENBURY
Right well, dear madam. By your patience,
I may not suffer you to visit them.
The King hath strictly charged the contrary.

QUEEN ELIZABETH
The King? Who's that?

BRAKENBURY  I mean, the Lord Protector.

QUEEN ELIZABETH
The Lord protect him from that kingly title!
Hath he set bounds between their love and me?
I am their mother. Who shall bar me from them?

DUCHESS
I am their father's mother. I will see them.

ANNE
Their aunt I am in law, in love their mother.
Then bring me to their sights. I'll bear thy blame
And take thy office from thee, on my peril.

BRAKENBURY
No, madam, no. I may not leave it so.
I am bound by oath, and therefore pardon me.
[Brakenbury the Lieutenant exits.]

[Enter Stanley.]


STANLEY
Let me but meet you ladies one hour hence,
And I'll salute your Grace of York as mother
And reverend looker-on of two fair queens.
[To Anne.] Come, madam, you must straight to
Westminster,
There to be crowned Richard's royal queen.

QUEEN ELIZABETH  Ah, cut my lace asunder
That my pent heart may have some scope to beat,
Or else I swoon with this dead-killing news!

ANNE
Despiteful tidings! O, unpleasing news!

DORSET, [to Queen Elizabeth]
Be of good cheer, mother. How fares your Grace?

QUEEN ELIZABETH
O Dorset, speak not to me. Get thee gone.
Death and destruction dogs thee at thy heels.
Thy mother's name is ominous to children.
If thou wilt outstrip death, go, cross the seas,
And live with Richmond, from the reach of hell.
Go, hie thee, hie thee from this slaughterhouse,
Lest thou increase the number of the dead
And make me die the thrall of Margaret's curse,
Nor mother, wife, nor England's counted queen.

STANLEY
Full of wise care is this your counsel, madam.
[To Dorset.] Take all the swift advantage of the
hours.
You shall have letters from me to my son
In your behalf, to meet you on the way.
Be not ta'en tardy by unwise delay.

DUCHESS
O ill-dispersing wind of misery!
O my accursed womb, the bed of death!
A cockatrice hast thou hatched to the world,
Whose unavoided eye is murderous.

STANLEY, [to Anne]
Come, madam, come. I in all haste was sent.

ANNE
And I with all unwillingness will go.
O, would to God that the inclusive verge
Of golden metal that must round my brow
Were red-hot steel to sear me to the brains!
Anointed let me be with deadly venom,
And die ere men can say "God save the Queen."

QUEEN ELIZABETH
Go, go, poor soul, I envy not thy glory.
To feed my humor, wish thyself no harm.

ANNE
No? Why? When he that is my husband now
Came to me as I followed Henry's corse,
When scarce the blood was well washed from his
hands
Which issued from my other angel husband
And that dear saint which then I weeping followed--
O, when, I say, I looked on Richard's face,
This was my wish: be thou, quoth I, accursed
For making me, so young, so old a widow;
And, when thou wedd'st, let sorrow haunt thy bed;
And be thy wife, if any be so mad,
More miserable by the life of thee
Than thou hast made me by my dear lord's death.
Lo, ere I can repeat this curse again,
Within so small a time my woman's heart
Grossly grew captive to his honey words
And proved the subject of mine own soul's curse,
Which hitherto hath held my eyes from rest,
For never yet one hour in his bed
Did I enjoy the golden dew of sleep,
But with his timorous dreams was still awaked.
Besides, he hates me for my father Warwick,
And will, no doubt, shortly be rid of me.

QUEEN ELIZABETH
Poor heart, adieu. I pity thy complaining.

ANNE
No more than with my soul I mourn for yours.

DORSET
Farewell, thou woeful welcomer of glory.

ANNE
Adieu, poor soul that tak'st thy leave of it.

DUCHESS, [to Dorset]
Go thou to Richmond, and good fortune guide thee.
[To Anne.] Go thou to Richard, and good angels
tend thee.
[To Queen Elizabeth.] Go thou to sanctuary, and
good thoughts possess thee.
I to my grave, where peace and rest lie with me.
Eighty-odd years of sorrow have I seen,
And each hour's joy wracked with a week of teen.

QUEEN ELIZABETH
Stay, yet look back with me unto the Tower.--
Pity, you ancient stones, those tender babes
Whom envy hath immured within your walls--
Rough cradle for such little pretty ones.
Rude ragged nurse, old sullen playfellow
For tender princes, use my babies well.
So foolish sorrows bids your stones farewell.
[They exit.]

Scene 2
=======
[Sound a sennet. Enter Richard in pomp; Buckingham,
Catesby, Ratcliffe, Lovell, and others, including a Page.]


RICHARD
Stand all apart.--Cousin of Buckingham.
[The others move aside.]

BUCKINGHAM  My gracious sovereign.

RICHARD
Give me thy hand.
[Here he ascendeth the throne. Sound trumpets.]
Thus high, by thy advice
And thy assistance is King Richard seated.
But shall we wear these glories for a day,
Or shall they last and we rejoice in them?

BUCKINGHAM
Still live they, and forever let them last.

RICHARD
Ah, Buckingham, now do I play the touch,
To try if thou be current gold indeed:
Young Edward lives; think now what I would speak.

BUCKINGHAM  Say on, my loving lord.

RICHARD
Why, Buckingham, I say I would be king.

BUCKINGHAM
Why so you are, my thrice-renowned lord.

RICHARD
Ha! Am I king? 'Tis so--but Edward lives.

BUCKINGHAM
True, noble prince.

RICHARD  O bitter consequence
That Edward still should live "true noble prince"!
Cousin, thou wast not wont to be so dull.
Shall I be plain? I wish the bastards dead,
And I would have it suddenly performed.
What sayst thou now? Speak suddenly. Be brief.

BUCKINGHAM  Your Grace may do your pleasure.

RICHARD
Tut, tut, thou art all ice; thy kindness freezes.
Say, have I thy consent that they shall die?

BUCKINGHAM
Give me some little breath, some pause, dear lord,
Before I positively speak in this.
I will resolve you herein presently.
[Buckingham exits.]

CATESBY, [aside to the other Attendants]
The King is angry. See, he gnaws his lip.

RICHARD, [aside]
I will converse with iron-witted fools
And unrespective boys. None are for me
That look into me with considerate eyes.
High-reaching Buckingham grows circumspect.--
Boy!

PAGE, [coming forward]  My lord?

RICHARD
Know'st thou not any whom corrupting gold
Will tempt unto a close exploit of death?

PAGE
I know a discontented gentleman
Whose humble means match not his haughty spirit.
Gold were as good as twenty orators,
And will, no doubt, tempt him to anything.

RICHARD
What is his name?

PAGE  His name, my lord, is Tyrrel.

RICHARD
I partly know the man. Go, call him hither, boy.
[Page exits.]
[Aside.] The deep-revolving witty Buckingham
No more shall be the neighbor to my counsels.
Hath he so long held out with me, untired,
And stops he now for breath? Well, be it so.

[Enter Stanley.]

How now, Lord Stanley, what's the news?

STANLEY  Know, my loving lord,
The Marquess Dorset, as I hear, is fled
To Richmond, in the parts where he abides.
[He walks aside.]

RICHARD
Come hither, Catesby. Rumor it abroad
That Anne my wife is very grievous sick.
I will take order for her keeping close.
Inquire me out some mean poor gentleman,
Whom I will marry straight to Clarence' daughter.
The boy is foolish, and I fear not him.
Look how thou dream'st! I say again, give out
That Anne my queen is sick and like to die.
About it, for it stands me much upon
To stop all hopes whose growth may damage me.
[Catesby exits.]
[Aside.] I must be married to my brother's daughter,
Or else my kingdom stands on brittle glass.
Murder her brothers, and then marry her--
Uncertain way of gain. But I am in
So far in blood that sin will pluck on sin.
Tear-falling pity dwells not in this eye.

[Enter Tyrrel.]

Is thy name Tyrrel?

TYRREL
James Tyrrel, and your most obedient subject.

RICHARD
Art thou indeed?

TYRREL  Prove me, my gracious lord.

RICHARD
Dar'st thou resolve to kill a friend of mine?

TYRREL
Please you. But I had rather kill two enemies.

RICHARD
Why then, thou hast it. Two deep enemies,
Foes to my rest, and my sweet sleep's disturbers,
Are they that I would have thee deal upon.
Tyrrel, I mean those bastards in the Tower.

TYRREL
Let me have open means to come to them,
And soon I'll rid you from the fear of them.

RICHARD
Thou sing'st sweet music. Hark, come hither, Tyrrel.
[Tyrrel approaches Richard and kneels.]
Go, by this token. Rise, and lend thine ear.
[Tyrrel rises, and Richard whispers
to him. Then Tyrrel steps back.]
There is no more but so. Say it is done,
And I will love thee and prefer thee for it.

TYRREL  I will dispatch it straight.	[He exits.]

[Enter Buckingham.]


BUCKINGHAM
My lord, I have considered in my mind
The late request that you did sound me in.

RICHARD
Well, let that rest. Dorset is fled to Richmond.

BUCKINGHAM  I hear the news, my lord.

RICHARD
Stanley, he is your wife's son. Well, look unto it.

BUCKINGHAM
My lord, I claim the gift, my due by promise,
For which your honor and your faith is pawned--
Th' earldom of Hereford and the movables
Which you have promised I shall possess.

RICHARD
Stanley, look to your wife. If she convey
Letters to Richmond, you shall answer it.

BUCKINGHAM
What says your Highness to my just request?

RICHARD
I do remember me, Henry the Sixth
Did prophesy that Richmond should be king,
When Richmond was a little peevish boy.
A king perhaps--

BUCKINGHAM  My lord--

RICHARD
How chance the prophet could not at that time
Have told me, I being by, that I should kill him?

BUCKINGHAM
My lord, your promise for the earldom--

RICHARD
Richmond! When last I was at Exeter,
The Mayor in courtesy showed me the castle
And called it Rougemont, at which name I started,
Because a bard of Ireland told me once
I should not live long after I saw Richmond.

BUCKINGHAM  My lord--

RICHARD  Ay, what's o'clock?

BUCKINGHAM
I am thus bold to put your Grace in mind
Of what you promised me.

RICHARD  Well, but what's o'clock?

BUCKINGHAM  Upon the stroke of ten.

RICHARD  Well, let it strike.

BUCKINGHAM  Why let it strike?

RICHARD
Because that, like a jack, thou keep'st the stroke
Betwixt thy begging and my meditation.
I am not in the giving vein today.

BUCKINGHAM
Why then, resolve me whether you will or no.

RICHARD
Thou troublest me; I am not in the vein.
[He exits, and is followed by all but Buckingham.]

BUCKINGHAM
And is it thus? Repays he my deep service
With such contempt? Made I him king for this?
O, let me think on Hastings and be gone
To Brecknock, while my fearful head is on!
[He exits.]

Scene 3
=======
[Enter Tyrrel.]


TYRREL
The tyrannous and bloody act is done,
The most arch deed of piteous massacre
That ever yet this land was guilty of.
Dighton and Forrest, who I did suborn
To do this piece of ruthless butchery,
Albeit they were fleshed villains, bloody dogs,
Melted with tenderness and mild compassion,
Wept like two children in their deaths' sad story.
"O thus," quoth Dighton, "lay the gentle babes."
"Thus, thus," quoth Forrest, "girdling one another
Within their alabaster innocent arms.
Their lips were four red roses on a stalk,
And in their summer beauty kissed each other.
A book of prayers on their pillow lay,
Which once," quoth Forrest, "almost changed my
mind,
But, O, the devil--" There the villain stopped;
When Dighton thus told on: "We smothered
The most replenished sweet work of nature
That from the prime creation e'er she framed."
Hence both are gone with conscience and remorse;
They could not speak; and so I left them both
To bear this tidings to the bloody king.

[Enter Richard.]

And here he comes.--All health, my sovereign lord.

RICHARD
Kind Tyrrel, am I happy in thy news?

TYRREL
If to have done the thing you gave in charge
Beget your happiness, be happy then,
For it is done.

RICHARD  But did'st thou see them dead?

TYRREL
I did, my lord.

RICHARD  And buried, gentle Tyrrel?

TYRREL
The chaplain of the Tower hath buried them,
But where, to say the truth, I do not know.

RICHARD
Come to me, Tyrrel, soon at after-supper,
When thou shalt tell the process of their death.
Meantime, but think how I may do thee good,
And be inheritor of thy desire.
Farewell till then.

TYRREL  I humbly take my leave.
[Tyrrel exits.]

RICHARD
The son of Clarence have I pent up close,
His daughter meanly have I matched in marriage,
The sons of Edward sleep in Abraham's bosom,
And Anne my wife hath bid this world goodnight.
Now, for I know the Breton Richmond aims
At young Elizabeth, my brother's daughter,
And by that knot looks proudly on the crown,
To her go I, a jolly thriving wooer.

[Enter Ratcliffe.]


RATCLIFFE  My lord.

RICHARD
Good or bad news, that thou com'st in so bluntly?

RATCLIFFE
Bad news, my lord. Morton is fled to Richmond,
And Buckingham, backed with the hardy Welshmen,
Is in the field, and still his power increaseth.

RICHARD
Ely with Richmond troubles me more near
Than Buckingham and his rash-levied strength.
Come, I have learned that fearful commenting
Is leaden servitor to dull delay;
Delay leads impotent and snail-paced beggary;
Then fiery expedition be my wing,
Jove's Mercury, and herald for a king.
Go, muster men. My counsel is my shield.
We must be brief when traitors brave the field.
[They exit.]

Scene 4
=======
[Enter old Queen Margaret.]


QUEEN MARGARET
So now prosperity begins to mellow
And drop into the rotten mouth of death.
Here in these confines slyly have I lurked
To watch the waning of mine enemies.
A dire induction am I witness to,
And will to France, hoping the consequence
Will prove as bitter, black, and tragical.
Withdraw thee, wretched Margaret. Who comes
here?	[She steps aside.]

[Enter Duchess of York and Queen Elizabeth.]


QUEEN ELIZABETH
Ah, my poor princes! Ah, my tender babes,
My unblown flowers, new-appearing sweets,
If yet your gentle souls fly in the air
And be not fixed in doom perpetual,
Hover about me with your airy wings
And hear your mother's lamentation.

QUEEN MARGARET, [aside]
Hover about her; say that right for right
Hath dimmed your infant morn to aged night.

DUCHESS
So many miseries have crazed my voice
That my woe-wearied tongue is still and mute.
Edward Plantagenet, why art thou dead?

QUEEN MARGARET, [aside]
Plantagenet doth quit Plantagenet;
Edward for Edward pays a dying debt.

QUEEN ELIZABETH
Wilt thou, O God, fly from such gentle lambs
And throw them in the entrails of the wolf?
When didst thou sleep when such a deed was done?

QUEEN MARGARET, [aside]
When holy Harry died, and my sweet son.

DUCHESS, [to Queen Elizabeth]
Dead life, blind sight, poor mortal living ghost,
Woe's scene, world's shame, grave's due by life
usurped,
Brief abstract and record of tedious days,
Rest thy unrest on England's lawful earth,
Unlawfully made drunk with innocent blood.

QUEEN ELIZABETH, [as they both sit down]
Ah, that thou wouldst as soon afford a grave
As thou canst yield a melancholy seat,
Then would I hide my bones, not rest them here.
Ah, who hath any cause to mourn but we?

QUEEN MARGARET, [coming forward]
If ancient sorrow be most reverend,
Give mine the benefit of seigniory,
And let my griefs frown on the upper hand.
If sorrow can admit society,
Tell over your woes again by viewing mine.
I had an Edward till a Richard killed him;
I had a husband till a Richard killed him.
Thou hadst an Edward till a Richard killed him;
Thou hadst a Richard till a Richard killed him.

DUCHESS
I had a Richard too, and thou did'st kill him;
I had a Rutland too; thou holp'st to kill him.

QUEEN MARGARET
Thou hadst a Clarence too, and Richard killed him.
From forth the kennel of thy womb hath crept
A hellhound that doth hunt us all to death--
That dog, that had his teeth before his eyes,
To worry lambs and lap their gentle blood;
That excellent grand tyrant of the Earth,
That reigns in galled eyes of weeping souls;
That foul defacer of God's handiwork
Thy womb let loose to chase us to our graves.
O upright, just, and true-disposing God,
How do I thank thee that this carnal cur
Preys on the issue of his mother's body
And makes her pew-fellow with others' moan!

DUCHESS, [standing]
O Harry's wife, triumph not in my woes!
God witness with me, I have wept for thine.

QUEEN MARGARET
Bear with me. I am hungry for revenge,
And now I cloy me with beholding it.
Thy Edward he is dead, that killed my Edward,
Thy other Edward dead, to quit my Edward;
Young York, he is but boot, because both they
Matched not the high perfection of my loss.
Thy Clarence he is dead that stabbed my Edward,
And the beholders of this frantic play,
Th' adulterate Hastings, Rivers, Vaughan, Grey,
Untimely smothered in their dusky graves.
Richard yet lives, hell's black intelligencer,
Only reserved their factor to buy souls
And send them thither. But at hand, at hand
Ensues his piteous and unpitied end.
Earth gapes, hell burns, fiends roar, saints pray,
To have him suddenly conveyed from hence.
Cancel his bond of life, dear God I pray,
That I may live and say "The dog is dead."

QUEEN ELIZABETH, [standing]
O, thou didst prophesy the time would come
That I should wish for thee to help me curse
That bottled spider, that foul bunch-backed toad!

QUEEN MARGARET
I called thee then "vain flourish of my fortune."
I called thee then poor shadow, "painted queen,"
The presentation of but what I was,
The flattering index of a direful pageant,
One heaved a-high to be hurled down below,
A mother only mocked with two fair babes,
A dream of what thou wast, a garish flag
To be the aim of every dangerous shot,
A sign of dignity, a breath, a bubble,
A queen in jest, only to fill the scene.
Where is thy husband now? Where be thy brothers?
Where are thy two sons? Wherein dost thou joy?
Who sues and kneels and says "God save the
Queen?"
Where be the bending peers that flattered thee?
Where be the thronging troops that followed thee?
Decline all this, and see what now thou art:
For happy wife, a most distressed widow;
For joyful mother, one that wails the name;
For one being sued to, one that humbly sues;
For queen, a very caitiff crowned with care;
For she that scorned at me, now scorned of me;
For she being feared of all, now fearing one;
For she commanding all, obeyed of none.
Thus hath the course of justice whirled about
And left thee but a very prey to time,
Having no more but thought of what thou wast
To torture thee the more, being what thou art.
Thou didst usurp my place, and dost thou not
Usurp the just proportion of my sorrow?
Now thy proud neck bears half my burdened yoke,
From which even here I slip my weary head
And leave the burden of it all on thee.
Farewell, York's wife, and queen of sad mischance.
These English woes shall make me smile in France.
[She begins to exit.]

QUEEN ELIZABETH
O, thou well-skilled in curses, stay awhile,
And teach me how to curse mine enemies.

QUEEN MARGARET
Forbear to sleep the nights, and fast the days;
Compare dead happiness with living woe;
Think that thy babes were sweeter than they were,
And he that slew them fouler than he is.
Bettering thy loss makes the bad causer worse.
Revolving this will teach thee how to curse.

QUEEN ELIZABETH
My words are dull. O, quicken them with thine!

QUEEN MARGARET
Thy woes will make them sharp and pierce like
mine.	[Margaret exits.]

DUCHESS
Why should calamity be full of words?

QUEEN ELIZABETH
Windy attorneys to their clients' woes,
Airy succeeders of intestate joys,
Poor breathing orators of miseries,
Let them have scope; though what they will impart
Help nothing else, yet do they ease the heart.

DUCHESS
If so, then be not tongue-tied. Go with me,
And in the breath of bitter words let's smother
My damned son that thy two sweet sons smothered.
[A trumpet sounds.]
[The trumpet sounds.] Be copious in exclaims.

	[Enter King Richard and his train, including Catesby.]


RICHARD
Who intercepts me in my expedition?

DUCHESS
O, she that might have intercepted thee,
By strangling thee in her accursed womb,
From all the slaughters, wretch, that thou hast done.

QUEEN ELIZABETH, [to Richard]
Hid'st thou that forehead with a golden crown
Where should be branded, if that right were right,
The slaughter of the prince that owed that crown
And the dire death of my poor sons and brothers?
Tell me, thou villain-slave, where are my children?

DUCHESS, [to Richard]
Thou toad, thou toad, where is thy brother Clarence,
And little Ned Plantagenet his son?

QUEEN ELIZABETH, [to Richard]
Where is the gentle Rivers, Vaughan, Grey?

DUCHESS, [to Richard]  Where is kind Hastings?

RICHARD
A flourish, trumpets! Strike alarum, drums!
Let not the heavens hear these telltale women
Rail on the Lord's anointed. Strike, I say!
[Flourish. Alarums.]
Either be patient and entreat me fair,
Or with the clamorous report of war
Thus will I drown your exclamations.

DUCHESS  Art thou my son?

RICHARD
Ay, I thank God, my father, and yourself.

DUCHESS
Then patiently hear my impatience.

RICHARD
Madam, I have a touch of your condition,
That cannot brook the accent of reproof.

DUCHESS
O, let me speak!

RICHARD  Do then, but I'll not hear.

DUCHESS
I will be mild and gentle in my words.

RICHARD
And brief, good mother, for I am in haste.

DUCHESS
Art thou so hasty? I have stayed for thee,
God knows, in torment and in agony.

RICHARD
And came I not at last to comfort you?

DUCHESS
No, by the Holy Rood, thou know'st it well.
Thou cam'st on Earth to make the Earth my hell.
A grievous burden was thy birth to me;
Tetchy and wayward was thy infancy;
Thy school days frightful, desp'rate, wild, and
furious;
Thy prime of manhood daring, bold, and venturous;
Thy age confirmed, proud, subtle, sly, and bloody,
More mild, but yet more harmful, kind in hatred.
What comfortable hour canst thou name,
That ever graced me with thy company?

RICHARD
Faith, none but Humfrey Hower, that called your
Grace
To breakfast once, forth of my company.
If I be so disgracious in your eye,
Let me march on and not offend you, madam.--
Strike up the drum.

DUCHESS  I prithee, hear me speak.

RICHARD
You speak too bitterly.

DUCHESS  Hear me a word,
For I shall never speak to thee again.

RICHARD  So.

DUCHESS
Either thou wilt die by God's just ordinance
Ere from this war thou turn a conqueror,
Or I with grief and extreme age shall perish
And nevermore behold thy face again.
Therefore take with thee my most grievous curse,
Which in the day of battle tire thee more
Than all the complete armor that thou wear'st.
My prayers on the adverse party fight,
And there the little souls of Edward's children
Whisper the spirits of thine enemies
And promise them success and victory.
Bloody thou art; bloody will be thy end.
Shame serves thy life and doth thy death attend.
[She exits.]

QUEEN ELIZABETH
Though far more cause, yet much less spirit to
curse
Abides in me. I say amen to her.

RICHARD
Stay, madam. I must talk a word with you.

QUEEN ELIZABETH
I have no more sons of the royal blood
For thee to slaughter. For my daughters, Richard,
They shall be praying nuns, not weeping queens,
And therefore level not to hit their lives.

RICHARD
You have a daughter called Elizabeth,
Virtuous and fair, royal and gracious.

QUEEN ELIZABETH
And must she die for this? O, let her live,
And I'll corrupt her manners, stain her beauty,
Slander myself as false to Edward's bed,
Throw over her the veil of infamy.
So she may live unscarred of bleeding slaughter,
I will confess she was not Edward's daughter.

RICHARD
Wrong not her birth. She is a royal princess.

QUEEN ELIZABETH
To save her life, I'll say she is not so.

RICHARD
Her life is safest only in her birth.

QUEEN ELIZABETH
And only in that safety died her brothers.

RICHARD
Lo, at their birth good stars were opposite.

QUEEN ELIZABETH
No, to their lives ill friends were contrary.

RICHARD
All unavoided is the doom of destiny.

QUEEN ELIZABETH
True, when avoided grace makes destiny.
My babes were destined to a fairer death
If grace had blessed thee with a fairer life.

RICHARD
You speak as if that I had slain my cousins.

QUEEN ELIZABETH
Cousins, indeed, and by their uncle cozened
Of comfort, kingdom, kindred, freedom, life.
Whose hand soever launched their tender hearts,
Thy head, all indirectly, gave direction.
No doubt the murd'rous knife was dull and blunt
Till it was whetted on thy stone-hard heart,
To revel in the entrails of my lambs.
But that still use of grief makes wild grief tame,
My tongue should to thy ears not name my boys
Till that my nails were anchored in thine eyes,
And I, in such a desp'rate bay of death,
Like a poor bark of sails and tackling reft,
Rush all to pieces on thy rocky bosom.

RICHARD
Madam, so thrive I in my enterprise
And dangerous success of bloody wars
As I intend more good to you and yours
Than ever you or yours by me were harmed!

QUEEN ELIZABETH
What good is covered with the face of heaven,
To be discovered, that can do me good?

RICHARD
Th' advancement of your children, gentle lady.

QUEEN ELIZABETH
Up to some scaffold, there to lose their heads.

RICHARD
Unto the dignity and height of fortune,
The high imperial type of this Earth's glory.

QUEEN ELIZABETH
Flatter my sorrow with report of it.
Tell me what state, what dignity, what honor,
Canst thou demise to any child of mine?

RICHARD
Even all I have--ay, and myself and all--
Will I withal endow a child of thine;
So in the Lethe of thy angry soul
Thou drown the sad remembrance of those wrongs
Which thou supposest I have done to thee.

QUEEN ELIZABETH
Be brief, lest that the process of thy kindness
Last longer telling than thy kindness' date.

RICHARD
Then know that from my soul I love thy daughter.

QUEEN ELIZABETH
My daughter's mother thinks it with her soul.

RICHARD  What do you think?

QUEEN ELIZABETH
That thou dost love my daughter from thy soul.
So from thy soul's love didst thou love her brothers,
And from my heart's love I do thank thee for it.

RICHARD
Be not so hasty to confound my meaning.
I mean that with my soul I love thy daughter
And do intend to make her Queen of England.

QUEEN ELIZABETH
Well then, who dost thou mean shall be her king?

RICHARD
Even he that makes her queen. Who else should be?

QUEEN ELIZABETH
What, thou?

RICHARD  Even so. How think you of it?

QUEEN ELIZABETH
How canst thou woo her?

RICHARD  That would I learn of you,
As one being best acquainted with her humor.

QUEEN ELIZABETH  And wilt thou learn of me?

RICHARD  Madam, with all my heart.

QUEEN ELIZABETH
Send to her, by the man that slew her brothers,
A pair of bleeding hearts; thereon engrave
"Edward" and "York." Then haply will she weep.
Therefore present to her--as sometime Margaret
Did to thy father, steeped in Rutland's blood--
A handkerchief, which say to her did drain
The purple sap from her sweet brother's body,
And bid her wipe her weeping eyes withal.
If this inducement move her not to love,
Send her a letter of thy noble deeds;
Tell her thou mad'st away her uncle Clarence,
Her uncle Rivers, ay, and for her sake
Mad'st quick conveyance with her good aunt Anne.

RICHARD
You mock me, madam. This is not the way
To win your daughter.

QUEEN ELIZABETH  There is no other way,
Unless thou couldst put on some other shape
And not be Richard, that hath done all this.

RICHARD
Say that I did all this for love of her.

QUEEN ELIZABETH
Nay, then indeed she cannot choose but hate thee,
Having bought love with such a bloody spoil.

RICHARD
Look what is done cannot be now amended.
Men shall deal unadvisedly sometimes,
Which after-hours gives leisure to repent.
If I did take the kingdom from your sons,
To make amends I'll give it to your daughter.
If I have killed the issue of your womb,
To quicken your increase I will beget
Mine issue of your blood upon your daughter.
A grandam's name is little less in love
Than is the doting title of a mother.
They are as children but one step below,
Even of your metal, of your very blood,
Of all one pain, save for a night of groans
Endured of her for whom you bid like sorrow.
Your children were vexation to your youth,
But mine shall be a comfort to your age.
The loss you have is but a son being king,
And by that loss your daughter is made queen.
I cannot make you what amends I would;
Therefore accept such kindness as I can.
Dorset your son, that with a fearful soul
Leads discontented steps in foreign soil,
This fair alliance quickly shall call home
To high promotions and great dignity.
The king that calls your beauteous daughter wife
Familiarly shall call thy Dorset brother.
Again shall you be mother to a king,
And all the ruins of distressful times
Repaired with double riches of content.
What, we have many goodly days to see!
The liquid drops of tears that you have shed
Shall come again, transformed to orient pearl,
Advantaging their love with interest
Of ten times double gain of happiness.
Go then, my mother; to thy daughter go.
Make bold her bashful years with your experience;
Prepare her ears to hear a wooer's tale;
Put in her tender heart th' aspiring flame
Of golden sovereignty; acquaint the Princess
With the sweet silent hours of marriage joys;
And when this arm of mine hath chastised
The petty rebel, dull-brained Buckingham,
Bound with triumphant garlands will I come
And lead thy daughter to a conqueror's bed,
To whom I will retail my conquest won,
And she shall be sole victoress, Caesar's Caesar.

QUEEN ELIZABETH
What were I best to say? Her father's brother
Would be her lord? Or shall I say her uncle?
Or he that slew her brothers and her uncles?
Under what title shall I woo for thee,
That God, the law, my honor, and her love
Can make seem pleasing to her tender years?

RICHARD
Infer fair England's peace by this alliance.

QUEEN ELIZABETH
Which she shall purchase with still-lasting war.

RICHARD
Tell her the King, that may command, entreats--

QUEEN ELIZABETH
That, at her hands, which the King's King forbids.

RICHARD
Say she shall be a high and mighty queen.

QUEEN ELIZABETH
To vail the title, as her mother doth.

RICHARD
Say I will love her everlastingly.

QUEEN ELIZABETH
But how long shall that title "ever" last?

RICHARD
Sweetly in force unto her fair life's end.

QUEEN ELIZABETH
But how long fairly shall her sweet life last?

RICHARD
As long as heaven and nature lengthens it.

QUEEN ELIZABETH
As long as hell and Richard likes of it.

RICHARD
Say I, her sovereign, am her subject low.

QUEEN ELIZABETH
But she, your subject, loathes such sovereignty.

RICHARD
Be eloquent in my behalf to her.

QUEEN ELIZABETH
An honest tale speeds best being plainly told.

RICHARD
Then plainly to her tell my loving tale.

QUEEN ELIZABETH
Plain and not honest is too harsh a style.

RICHARD
Your reasons are too shallow and too quick.

QUEEN ELIZABETH
O no, my reasons are too deep and dead--
Too deep and dead, poor infants, in their graves.

RICHARD
Harp not on that string, madam; that is past.

QUEEN ELIZABETH
Harp on it still shall I till heart-strings break.

RICHARD
Now by my George, my Garter, and my crown--

QUEEN ELIZABETH
Profaned, dishonored, and the third usurped.

RICHARD
I swear--

QUEEN ELIZABETH  By nothing, for this is no oath.
Thy George, profaned, hath lost his lordly honor;
Thy Garter, blemished, pawned his knightly virtue;
Thy crown, usurped, disgraced his kingly glory.
If something thou wouldst swear to be believed,
Swear then by something that thou hast not
wronged.

RICHARD
Then, by myself--

QUEEN ELIZABETH  Thyself is self-misused.

RICHARD
Now, by the world--

QUEEN ELIZABETH  'Tis full of thy foul wrongs.

RICHARD
My father's death--

QUEEN ELIZABETH  Thy life hath it dishonored.

RICHARD
Why then, by God.

QUEEN ELIZABETH  God's wrong is most of all.
If thou didst fear to break an oath with Him,
The unity the King my husband made
Thou hadst not broken, nor my brothers died.
If thou hadst feared to break an oath by Him,
Th' imperial metal circling now thy head
Had graced the tender temples of my child,
And both the Princes had been breathing here,
Which now, two tender bedfellows for dust,
Thy broken faith hath made the prey for worms.
What canst thou swear by now?

RICHARD  The time to come.

QUEEN ELIZABETH
That thou hast wronged in the time o'erpast;
For I myself have many tears to wash
Hereafter time, for time past wronged by thee.
The children live whose fathers thou hast
slaughtered,
Ungoverned youth, to wail it in their age;
The parents live whose children thou hast
butchered,
Old barren plants, to wail it with their age.
Swear not by time to come, for that thou hast
Misused ere used, by times ill-used o'erpast.

RICHARD
As I intend to prosper and repent,
So thrive I in my dangerous affairs
Of hostile arms! Myself myself confound,
Heaven and fortune bar me happy hours,
Day, yield me not thy light, nor night thy rest,
Be opposite all planets of good luck
To my proceeding if, with dear heart's love,
Immaculate devotion, holy thoughts,
I tender not thy beauteous princely daughter.
In her consists my happiness and thine.
Without her follows to myself and thee,
Herself, the land, and many a Christian soul,
Death, desolation, ruin, and decay.
It cannot be avoided but by this;
It will not be avoided but by this.
Therefore, dear mother--I must call you so--
Be the attorney of my love to her;
Plead what I will be, not what I have been;
Not my deserts, but what I will deserve.
Urge the necessity and state of times,
And be not peevish found in great designs.

QUEEN ELIZABETH
Shall I be tempted of the devil thus?

RICHARD
Ay, if the devil tempt you to do good.

QUEEN ELIZABETH
Shall I forget myself to be myself?

RICHARD
Ay, if your self's remembrance wrong yourself.

QUEEN ELIZABETH  Yet thou didst kill my children.

RICHARD
But in your daughter's womb I bury them,
Where, in that nest of spicery, they will breed
Selves of themselves, to your recomforture.

QUEEN ELIZABETH
Shall I go win my daughter to thy will?

RICHARD
And be a happy mother by the deed.

QUEEN ELIZABETH  I go. Write to me very shortly,
And you shall understand from me her mind.

RICHARD
Bear her my true love's kiss; and so, farewell.
[Queen exits.]
Relenting fool and shallow, changing woman!

[Enter Ratcliffe.]

How now, what news?

RATCLIFFE
Most mighty sovereign, on the western coast
Rideth a puissant navy. To our shores
Throng many doubtful hollow-hearted friends,
Unarmed and unresolved to beat them back.
'Tis thought that Richmond is their admiral;
And there they hull, expecting but the aid
Of Buckingham to welcome them ashore.

RICHARD
Some light-foot friend post to the Duke of
Norfolk--
Ratcliffe thyself, or Catesby. Where is he?

CATESBY
Here, my good lord.

RICHARD  Catesby, fly to the Duke.

CATESBY
I will, my lord, with all convenient haste.

RICHARD
Ratcliffe, come hither. Post to Salisbury.When thou com'st thither--[To Catesby.] Dull,
unmindful villain,
Why stay'st thou here and go'st not to the Duke?

CATESBY
First, mighty liege, tell me your Highness' pleasure,
What from your Grace I shall deliver to him.

RICHARD
O true, good Catesby. Bid him levy straight
The greatest strength and power that he can make
And meet me suddenly at Salisbury.

CATESBY  I go.	[He exits.]

RATCLIFFE
What, may it please you, shall I do at Salisbury?

RICHARD
Why, what wouldst thou do there before I go?

RATCLIFFE
Your Highness told me I should post before.

RICHARD
My mind is changed.

[Enter Lord Stanley.]

Stanley, what news with you?

STANLEY
None good, my liege, to please you with the hearing,
Nor none so bad but well may be reported.

RICHARD
Hoyday, a riddle! Neither good nor bad.
What need'st thou run so many miles about
When thou mayst tell thy tale the nearest way?
Once more, what news?

STANLEY  Richmond is on the seas.

RICHARD
There let him sink, and be the seas on him!
White-livered runagate, what doth he there?

STANLEY
I know not, mighty sovereign, but by guess.

RICHARD  Well, as you guess?

STANLEY
Stirred up by Dorset, Buckingham, and Morton,
He makes for England, here to claim the crown.

RICHARD
Is the chair empty? Is the sword unswayed?
Is the King dead, the empire unpossessed?
What heir of York is there alive but we?
And who is England's king but great York's heir?
Then tell me, what makes he upon the seas?

STANLEY
Unless for that, my liege, I cannot guess.

RICHARD
Unless for that he comes to be your liege,
You cannot guess wherefore the Welshman comes.
Thou wilt revolt and fly to him, I fear.

STANLEY
No, my good lord. Therefore mistrust me not.

RICHARD
Where is thy power, then, to beat him back?
Where be thy tenants and thy followers?
Are they not now upon the western shore,
Safe-conducting the rebels from their ships?

STANLEY
No, my good lord. My friends are in the north.

RICHARD
Cold friends to me. What do they in the north
When they should serve their sovereign in the west?

STANLEY
They have not been commanded, mighty king.
Pleaseth your Majesty to give me leave,
I'll muster up my friends and meet your Grace
Where and what time your Majesty shall please.

RICHARD
Ay, thou wouldst be gone to join with Richmond,
But I'll not trust thee.

STANLEY  Most mighty sovereign,
You have no cause to hold my friendship doubtful.
I never was nor never will be false.

RICHARD
Go then and muster men, but leave behind
Your son George Stanley. Look your heart be firm,
Or else his head's assurance is but frail.

STANLEY
So deal with him as I prove true to you.
[Stanley exits.]

[Enter a Messenger.]


FIRST MESSENGER
My gracious sovereign, now in Devonshire,
As I by friends am well advertised,
Sir Edward Courtney and the haughty prelate,
Bishop of Exeter, his elder brother,
With many more confederates are in arms.

[Enter another Messenger.]


SECOND MESSENGER
In Kent, my liege, the Guilfords are in arms,
And every hour more competitors
Flock to the rebels, and their power grows strong.

[Enter another Messenger.]


THIRD MESSENGER
My lord, the army of great Buckingham--

RICHARD
Out on you, owls! Nothing but songs of death.
[He striketh him.]
There, take thou that till thou bring better news.

THIRD MESSENGER
The news I have to tell your Majesty
Is that by sudden floods and fall of waters
Buckingham's army is dispersed and scattered,
And he himself wandered away alone,
No man knows whither.

RICHARD  I cry thee mercy.
There is my purse to cure that blow of thine.
[He gives money.]
Hath any well-advised friend proclaimed
Reward to him that brings the traitor in?

THIRD MESSENGER
Such proclamation hath been made, my lord.

[Enter another Messenger.]


FOURTH MESSENGER
Sir Thomas Lovell and Lord Marquess Dorset,
'Tis said, my liege, in Yorkshire are in arms.
But this good comfort bring I to your Highness:
The Breton navy is dispersed by tempest.
Richmond, in Dorsetshire, sent out a boat
Unto the shore to ask those on the banks
If they were his assistants, yea, or no--
Who answered him they came from Buckingham
Upon his party. He, mistrusting them,
Hoised sail and made his course again for Brittany.

RICHARD
March on, march on, since we are up in arms,
If not to fight with foreign enemies,
Yet to beat down these rebels here at home.

[Enter Catesby.]


CATESBY
My liege, the Duke of Buckingham is taken.
That is the best news. That the Earl of Richmond
Is with a mighty power landed at Milford
Is colder tidings, yet they must be told.

RICHARD
Away towards Salisbury! While we reason here,
A royal battle might be won and lost.
Someone take order Buckingham be brought
To Salisbury. The rest march on with me.
[Flourish. They exit.]

Scene 5
=======
[Enter Stanley, Earl of Derby, and Sir Christopher.]


STANLEY
Sir Christopher, tell Richmond this from me:
That in the sty of the most deadly boar
My son George Stanley is franked up in hold;
If I revolt, off goes young George's head;
The fear of that holds off my present aid.
So get thee gone. Commend me to thy lord.
Withal, say that the Queen hath heartily consented
He should espouse Elizabeth her daughter.
But tell me, where is princely Richmond now?

CHRISTOPHER
At Pembroke, or at Ha'rfordwest in Wales.

STANLEY  What men of name resort to him?

CHRISTOPHER
Sir Walter Herbert, a renowned soldier;
Sir Gilbert Talbot, Sir William Stanley,
Oxford, redoubted Pembroke, Sir James Blunt,
And Rice ap Thomas, with a valiant crew,
And many other of great name and worth;
And towards London do they bend their power,
If by the way they be not fought withal.

STANLEY, [giving Sir Christopher a paper]
Well, hie thee to thy lord. I kiss his hand.
My letter will resolve him of my mind.
Farewell.
[They exit.]


ACT 5
=====

Scene 1
=======
[Enter Buckingham, with Sheriff and Halberds, led to
execution.]


BUCKINGHAM
Will not King Richard let me speak with him?

SHERIFF
No, my good lord. Therefore be patient.

BUCKINGHAM
Hastings and Edward's children, Grey and Rivers,
Holy King Henry and thy fair son Edward,
Vaughan, and all that have miscarried
By underhand, corrupted, foul injustice,
If that your moody, discontented souls
Do through the clouds behold this present hour,
Even for revenge mock my destruction.--
This is All Souls' Day, fellow, is it not?

SHERIFF  It is.

BUCKINGHAM
Why, then, All Souls' Day is my body's doomsday.
This is the day which, in King Edward's time,
I wished might fall on me when I was found
False to his children and his wife's allies.
This is the day wherein I wished to fall
By the false faith of him whom most I trusted.
This, this All Souls' Day to my fearful soul
Is the determined respite of my wrongs.
That high All-seer which I dallied with
Hath turned my feigned prayer on my head
And given in earnest what I begged in jest.
Thus doth he force the swords of wicked men
To turn their own points in their masters' bosoms.
Thus Margaret's curse falls heavy on my neck:
"When he," quoth she, "shall split thy heart with
sorrow,
Remember Margaret was a prophetess."--
Come, lead me, officers, to the block of shame.
Wrong hath but wrong, and blame the due of blame.
[Buckingham exits with Officers.]

Scene 2
=======
[Enter Richmond, Oxford, Blunt, Herbert, and others,
with Drum and Colors.]


RICHMOND
Fellows in arms, and my most loving friends,
Bruised underneath the yoke of tyranny,
Thus far into the bowels of the land
Have we marched on without impediment,
And here receive we from our father Stanley
Lines of fair comfort and encouragement.
The wretched, bloody, and usurping boar,
That spoiled your summer fields and fruitful vines,
Swills your warm blood like wash, and makes his
trough
In your embowelled bosoms--this foul swine
Is now even in the center of this isle,
Near to the town of Leicester, as we learn.
From Tamworth thither is but one day's march.
In God's name, cheerly on, courageous friends,
To reap the harvest of perpetual peace
By this one bloody trial of sharp war.

OXFORD
Every man's conscience is a thousand men
To fight against this guilty homicide.

HERBERT
I doubt not but his friends will turn to us.

BLUNT
He hath no friends but what are friends for fear,
Which in his dearest need will fly from him.

RICHMOND
All for our vantage. Then, in God's name, march.
True hope is swift, and flies with swallow's wings;
Kings it makes gods, and meaner creatures kings.
[All exit.]

Scene 3
=======
[Enter King Richard, in arms, with Norfolk, Ratcliffe, and
the Earl of Surrey, with Soldiers.]


RICHARD
Here pitch our tent, even here in Bosworth field.
[Soldiers begin to pitch the tent.]
My lord of Surrey, why look you so sad?

SURREY
My heart is ten times lighter than my looks.

RICHARD
My lord of Norfolk--

NORFOLK  Here, most gracious liege.

RICHARD
Norfolk, we must have knocks, ha, must we not?

NORFOLK
We must both give and take, my loving lord.

RICHARD
Up with my tent!--Here will I lie tonight.
But where tomorrow? Well, all's one for that.
Who hath descried the number of the traitors?

NORFOLK
Six or seven thousand is their utmost power.

RICHARD
Why, our battalia trebles that account.
Besides, the King's name is a tower of strength
Which they upon the adverse faction want.--
Up with the tent!--Come, noble gentlemen,
Let us survey the vantage of the ground.
Call for some men of sound direction;
Let's lack no discipline, make no delay,
For, lords, tomorrow is a busy day.
[The tent now in place, they exit.]

[Enter Richmond, Sir William Brandon, Oxford,
Dorset, Herbert, Blunt, and others who set up
Richmond's tent.]


RICHMOND
The weary sun hath made a golden set,
And by the bright track of his fiery car
Gives token of a goodly day tomorrow.--
Sir William Brandon, you shall bear my standard.--
Give me some ink and paper in my tent;
I'll draw the form and model of our battle,
Limit each leader to his several charge,
And part in just proportion our small power.--
My Lord of Oxford, you, Sir William Brandon,
And you, Sir Walter Herbert, stay with me.
The Earl of Pembroke keeps his regiment.--
Good Captain Blunt, bear my goodnight to him,
And by the second hour in the morning
Desire the Earl to see me in my tent.
Yet one thing more, good captain, do for me.
Where is Lord Stanley quartered, do you know?

BLUNT
Unless I have mista'en his colors much,
Which well I am assured I have not done,
His regiment lies half a mile, at least,
South from the mighty power of the King.

RICHMOND
If without peril it be possible,
Sweet Blunt, make some good means to speak with
him,
And give him from me this most needful note.
[He gives a paper.]

BLUNT
Upon my life, my lord, I'll undertake it,
And so God give you quiet rest tonight.

RICHMOND
Good night, good Captain Blunt.	[Blunt exits.]
Come, gentlemen,
Let us consult upon tomorrow's business.
Into my tent. The dew is raw and cold.
[Richmond, Brandon, Dorset, Herbert, and Oxford
withdraw into the tent. The others exit.]

[Enter to his tent Richard, Ratcliffe, Norfolk, and
Catesby, with Soldiers.]


RICHARD  What is 't o'clock?

CATESBY
It's suppertime, my lord. It's nine o'clock.

RICHARD
I will not sup tonight. Give me some ink and paper.
What, is my beaver easier than it was,
And all my armor laid into my tent?

CATESBY
It is, my liege, and all things are in readiness.

RICHARD
Good Norfolk, hie thee to thy charge.
Use careful watch. Choose trusty sentinels.

NORFOLK  I go, my lord.

RICHARD
Stir with the lark tomorrow, gentle Norfolk.

NORFOLK  I warrant you, my lord.	[He exits.]

RICHARD  Catesby.

CATESBY  My lord.

RICHARD  Send out a pursuivant-at-arms
To Stanley's regiment. Bid him bring his power
Before sunrising, lest his son George fall
Into the blind cave of eternal night.	[Catesby exits.]
[To Soldiers.] Fill me a bowl of wine. Give me a
watch.
Saddle white Surrey for the field tomorrow.
Look that my staves be sound and not too heavy.--
Ratcliffe.

RATCLIFFE  My lord.

RICHARD
Sawst thou the melancholy Lord Northumberland?

RATCLIFFE
Thomas the Earl of Surrey and himself,
Much about cockshut time, from troop to troop
Went through the army cheering up the soldiers.

RICHARD
So, I am satisfied. Give me a bowl of wine.
I have not that alacrity of spirit
Nor cheer of mind that I was wont to have.
[Wine is brought.]
Set it down. Is ink and paper ready?

RATCLIFFE
It is, my lord.

RICHARD  Bid my guard watch. Leave me.
Ratcliffe, about the mid of night come to my tent
And help to arm me. Leave me, I say.
[Ratcliffe exits. Richard sleeps in his tent,
which is guarded by Soldiers.]

[Enter Stanley, Earl of Derby to Richmond in his tent.]


STANLEY
Fortune and victory sit on thy helm!

RICHMOND
All comfort that the dark night can afford
Be to thy person, noble father-in-law.
Tell me, how fares our loving mother?

STANLEY
I, by attorney, bless thee from thy mother,
Who prays continually for Richmond's good.
So much for that. The silent hours steal on,
And flaky darkness breaks within the east.
In brief, for so the season bids us be,
Prepare thy battle early in the morning,
And put thy fortune to the arbitrament
Of bloody strokes and mortal-staring war.
I, as I may--that which I would I cannot--
With best advantage will deceive the time
And aid thee in this doubtful shock of arms.
But on thy side I may not be too forward,
Lest, being seen, thy brother, tender George,
Be executed in his father's sight.
Farewell. The leisure and the fearful time
Cuts off the ceremonious vows of love
And ample interchange of sweet discourse,
Which so-long-sundered friends should dwell upon.
God give us leisure for these rites of love!
Once more, adieu. Be valiant and speed well.

RICHMOND
Good lords, conduct him to his regiment.
I'll strive with troubled thoughts to take a nap,
Lest leaden slumber peise me down tomorrow
When I should mount with wings of victory.
Once more, good night, kind lords and gentlemen.

[All but Richmond leave his tent and exit.]
[Richmond kneels.]
O Thou, whose captain I account myself,
Look on my forces with a gracious eye.
Put in their hands Thy bruising irons of wrath,
That they may crush down with a heavy fall
The usurping helmets of our adversaries.
Make us Thy ministers of chastisement,
That we may praise Thee in the victory.
To Thee I do commend my watchful soul,
Ere I let fall the windows of mine eyes.
Sleeping and waking, O, defend me still!	[Sleeps.]

[Enter the Ghost of young Prince Edward, son to Harry
the Sixth.]


GHOST OF EDWARD, [to Richard]
Let me sit heavy on thy soul tomorrow.
Think how thou stabbed'st me in my prime of
youth
At Tewkesbury. Despair therefore, and die!
[(To Richmond.)] Be cheerful, Richmond, for the
wronged souls
Of butchered princes fight in thy behalf.
King Henry's issue, Richmond, comforts thee.
[He exits.]

[Enter the Ghost of Henry the Sixth.]


GHOST OF HENRY, [to Richard]
When I was mortal, my anointed body
By thee was punched full of deadly holes.
Think on the Tower and me. Despair and die!
Harry the Sixth bids thee despair and die.
[(To Richmond.)] Virtuous and holy, be thou conqueror.
Harry, that prophesied thou shouldst be king,
Doth comfort thee in thy sleep. Live and flourish.
[He exits.]

[Enter the Ghost of Clarence.]


GHOST OF CLARENCE, [to Richard]
Let me sit heavy in thy soul tomorrow,
I, that was washed to death with fulsome wine,
Poor Clarence, by thy guile betrayed to death.
Tomorrow in the battle think on me,
And fall thy edgeless sword. Despair and die!
[(To Richmond.)] Thou offspring of the house of
Lancaster,
The wronged heirs of York do pray for thee.
Good angels guard thy battle. Live and flourish.
[He exits.]

[Enter the Ghosts of Rivers, Grey, and Vaughan.]


GHOST OF RIVERS, [to Richard]
Let me sit heavy in thy soul tomorrow,
Rivers, that died at Pomfret. Despair and die!

GHOST OF GREY, [to Richard]
Think upon Grey, and let thy soul despair!

GHOST OF VAUGHAN, [to Richard]
Think upon Vaughan, and with guilty fear
Let fall thy lance. Despair and die!

ALL, [to Richmond]
Awake, and think our wrongs in Richard's bosom
Will conquer him. Awake, and win the day.
[They exit.]

[Enter the Ghosts of the two young Princes.]


GHOSTS OF PRINCES, [to Richard]
Dream on thy cousins smothered in the Tower.
Let us be lead within thy bosom, Richard,
And weigh thee down to ruin, shame, and death.
Thy nephews' souls bid thee despair and die.
[(To Richmond.)] Sleep, Richmond, sleep in peace
and wake in joy.
Good angels guard thee from the boar's annoy.
Live, and beget a happy race of kings.
Edward's unhappy sons do bid thee flourish.
[They exit.]

[Enter the Ghost of Hastings.]


GHOST OF HASTINGS, [to Richard]
Bloody and guilty, guiltily awake,
And in a bloody battle end thy days.
Think on Lord Hastings. Despair and die!
[(To Richmond.)] Quiet, untroubled soul, awake, awake.
Arm, fight, and conquer for fair England's sake.
[He exits.]

[Enter the Ghost of Lady Anne his wife.]


GHOST OF ANNE, [to Richard]
Richard, thy wife, that wretched Anne thy wife,
That never slept a quiet hour with thee,
Now fills thy sleep with perturbations.
Tomorrow, in the battle, think on me,
And fall thy edgeless sword. Despair and die!
[(To Richmond.)] Thou quiet soul, sleep thou a quiet
sleep.
Dream of success and happy victory.
Thy adversary's wife doth pray for thee.	[She exits.]

[Enter the Ghost of Buckingham.]


GHOST OF BUCKINGHAM, [to Richard]
The first was I that helped thee to the crown;
The last was I that felt thy tyranny.
O, in the battle think on Buckingham,
And die in terror of thy guiltiness.
Dream on, dream on, of bloody deeds and death.
Fainting, despair; despairing, yield thy breath.
[(To Richmond.)] I died for hope ere I could lend
thee aid,
But cheer thy heart, and be thou not dismayed.
God and good angels fight on Richmond's side,
And Richard fall in height of all his pride.
[He exits.]
[Richard starteth up out of a dream.]

RICHARD
Give me another horse! Bind up my wounds!
Have mercy, Jesu!--Soft, I did but dream.
O coward conscience, how dost thou afflict me!
The lights burn blue; it is now dead midnight.
Cold fearful drops stand on my trembling flesh.
What do I fear? Myself? There's none else by.
Richard loves Richard, that is, I am I.
Is there a murderer here? No. Yes, I am.
Then fly! What, from myself? Great reason why:
Lest I revenge. What, myself upon myself?
Alack, I love myself. Wherefore? For any good
That I myself have done unto myself?
O, no. Alas, I rather hate myself
For hateful deeds committed by myself.
I am a villain. Yet I lie; I am not.
Fool, of thyself speak well. Fool, do not flatter.
My conscience hath a thousand several tongues,
And every tongue brings in a several tale,
And every tale condemns me for a villain.
Perjury, perjury, in the highest degree;
Murder, stern murder, in the direst degree;
All several sins, all used in each degree,
Throng to the bar, crying all "Guilty, guilty!"
I shall despair. There is no creature loves me,
And if I die no soul will pity me.
And wherefore should they, since that I myself
Find in myself no pity to myself?
Methought the souls of all that I had murdered
Came to my tent, and every one did threat
Tomorrow's vengeance on the head of Richard.

[Enter Ratcliffe.]


RATCLIFFE  My lord.

RICHARD  Zounds, who is there?

RATCLIFFE
Ratcliffe, my lord, 'tis I. The early village cock
Hath twice done salutation to the morn.
Your friends are up and buckle on their armor.

RICHARD
O Ratcliffe, I have dreamed a fearful dream!
What think'st thou, will our friends prove all true?

RATCLIFFE
No doubt, my lord.

RICHARD  O Ratcliffe, I fear, I fear.

RATCLIFFE
Nay, good my lord, be not afraid of shadows.

RICHARD
By the apostle Paul, shadows tonight
Have struck more terror to the soul of Richard
Than can the substance of ten thousand soldiers
Armed in proof and led by shallow Richmond.
'Tis not yet near day. Come, go with me.
Under our tents I'll play the eavesdropper
To see if any mean to shrink from me.
[Richard and Ratcliffe exit.]

[Enter the Lords to Richmond, in his tent.]


LORDS  Good morrow, Richmond.

RICHMOND
Cry mercy, lords and watchful gentlemen,
That you have ta'en a tardy sluggard here.

A LORD  How have you slept, my lord?

RICHMOND
The sweetest sleep and fairest-boding dreams
That ever entered in a drowsy head
Have I since your departure had, my lords.
Methought their souls whose bodies Richard
murdered
Came to my tent and cried on victory.
I promise you, my soul is very jocund
In the remembrance of so fair a dream.
How far into the morning is it, lords?

A LORD  Upon the stroke of four.

RICHMOND, [leaving the tent]
Why, then 'tis time to arm and give direction.

His oration to his soldiers.

More than I have said, loving countrymen,
The leisure and enforcement of the time
Forbids to dwell upon. Yet remember this:
God, and our good cause, fight upon our side.
The prayers of holy saints and wronged souls,
Like high-reared bulwarks, stand before our faces.
Richard except, those whom we fight against
Had rather have us win than him they follow.
For what is he they follow? Truly, gentlemen,
A bloody tyrant and a homicide;
One raised in blood, and one in blood established;
One that made means to come by what he hath,
And slaughtered those that were the means to help
him;
A base foul stone, made precious by the foil
Of England's chair, where he is falsely set;
One that hath ever been God's enemy.
Then if you fight against God's enemy,
God will, in justice, ward you as his soldiers.
If you do sweat to put a tyrant down,
You sleep in peace, the tyrant being slain.
If you do fight against your country's foes,
Your country's fat shall pay your pains the hire.
If you do fight in safeguard of your wives,
Your wives shall welcome home the conquerors.
If you do free your children from the sword,
Your children's children quits it in your age.
Then, in the name of God and all these rights,
Advance your standards; draw your willing swords.
For me, the ransom of my bold attempt
Shall be this cold corpse on the Earth's cold face,
But if I thrive, the gain of my attempt
The least of you shall share his part thereof.
Sound drums and trumpets boldly and cheerfully.
God, and Saint George, Richmond, and victory!
[They exit.]

[Enter King Richard, Ratcliffe, and Soldiers.]


RICHARD
What said Northumberland as touching Richmond?

RATCLIFFE
That he was never trained up in arms.

RICHARD
He said the truth. And what said Surrey then?

RATCLIFFE
He smiled and said "The better for our purpose."

RICHARD
He was in the right, and so indeed it is.
[The clock striketh.]
Tell the clock there. Give me a calendar.
[He looks in an almanac.]
Who saw the sun today?

RATCLIFFE  Not I, my lord.

RICHARD
Then he disdains to shine, for by the book
He should have braved the east an hour ago.
A black day will it be to somebody.
Ratcliffe!

RATCLIFFE
My lord.

RICHARD  The sun will not be seen today.
The sky doth frown and lour upon our army.
I would these dewy tears were from the ground.
Not shine today? Why, what is that to me
More than to Richmond, for the selfsame heaven
That frowns on me looks sadly upon him.

[Enter Norfolk.]


NORFOLK
Arm, arm, my lord. The foe vaunts in the field.

RICHARD
Come, bustle, bustle. Caparison my horse.--
Call up Lord Stanley; bid him bring his power.--
I will lead forth my soldiers to the plain,
And thus my battle shall be ordered:
My foreward shall be drawn out all in length,
Consisting equally of horse and foot;
Our archers shall be placed in the midst.
John Duke of Norfolk, Thomas Earl of Surrey,
Shall have the leading of this foot and horse.
They thus directed, we will follow
In the main battle, whose puissance on either side
Shall be well winged with our chiefest horse.
This, and Saint George to boot!--What think'st
thou, Norfolk?

NORFOLK
A good direction, warlike sovereign.
[He sheweth him a paper.]
This found I on my tent this morning.

RICHARD [reads]
	Jockey of Norfolk, be not so bold.
	For Dickon thy master is bought and sold.
A thing devised by the enemy.--
Go, gentlemen, every man unto his charge.
Let not our babbling dreams affright our souls.
Conscience is but a word that cowards use,
Devised at first to keep the strong in awe.
Our strong arms be our conscience, swords our law.
March on. Join bravely. Let us to it pell mell,
If not to heaven, then hand in hand to hell.

His oration to his army.

What shall I say more than I have inferred?
Remember whom you are to cope withal,
A sort of vagabonds, rascals, and runaways,
A scum of Bretons and base lackey peasants,
Whom their o'ercloyed country vomits forth
To desperate adventures and assured destruction.
You sleeping safe, they bring to you unrest;
You having lands and blessed with beauteous wives,
They would restrain the one, distain the other.
And who doth lead them but a paltry fellow,
Long kept in Brittany at our mother's cost,
A milksop, one that never in his life
Felt so much cold as overshoes in snow?
Let's whip these stragglers o'er the seas again,
Lash hence these overweening rags of France,
These famished beggars weary of their lives,
Who, but for dreaming on this fond exploit,
For want of means, poor rats, had hanged
themselves.
If we be conquered, let men conquer us,
And not these bastard Bretons, whom our fathers
Have in their own land beaten, bobbed, and
thumped,
And in record left them the heirs of shame.
Shall these enjoy our lands, lie with our wives,
Ravish our daughters?	[Drum afar off.]
Hark, I hear their drum.
Fight, gentlemen of England.--Fight, bold
yeomen.--
Draw, archers; draw your arrows to the head.--
Spur your proud horses hard, and ride in blood.
Amaze the welkin with your broken staves.--

[Enter a Messenger.]

What says Lord Stanley? Will he bring his power?

MESSENGER  My lord, he doth deny to come.

RICHARD  Off with his son George's head!

NORFOLK
My lord, the enemy is past the marsh.
After the battle let George Stanley die.

RICHARD
A thousand hearts are great within my bosom.
Advance our standards. Set upon our foes.
Our ancient word of courage, fair Saint George,
Inspire us with the spleen of fiery dragons.
Upon them! Victory sits on our helms.
[They exit.]

Scene 4
=======
[Alarum. Excursions. Enter Norfolk, with Soldiers, and
Catesby.]


CATESBY
Rescue, my lord of Norfolk, rescue, rescue!
The King enacts more wonders than a man,
Daring an opposite to every danger.
His horse is slain, and all on foot he fights,
Seeking for Richmond in the throat of death.
Rescue, fair lord, or else the day is lost.
[Norfolk exits with Soldiers.]

[Alarums. Enter Richard.]


RICHARD
A horse, a horse, my kingdom for a horse!

CATESBY
Withdraw, my lord. I'll help you to a horse.

RICHARD
Slave, I have set my life upon a cast,
And I will stand the hazard of the die.
I think there be six Richmonds in the field;
Five have I slain today instead of him.
A horse, a horse, my kingdom for a horse!
[They exit.]

Scene 5
=======
[Alarum. Enter Richard and Richmond. They fight.
Richard is slain. Then retreat being sounded, Richmond
exits, and Richard's body is removed. Flourish. Enter
Richmond, Stanley, Earl of Derby, bearing the crown,
with other Lords, and Soldiers.]


RICHMOND
God and your arms be praised, victorious friends!
The day is ours; the bloody dog is dead.

STANLEY, [offering him the crown]
Courageous Richmond, well hast thou acquit thee.
Lo, here this long-usurped royalty
From the dead temples of this bloody wretch
Have I plucked off, to grace thy brows withal.
Wear it, enjoy it, and make much of it.

RICHMOND
Great God of heaven, say amen to all!
But tell me, is young George Stanley living?

STANLEY
He is, my lord, and safe in Leicester town,
Whither, if it please you, we may now withdraw us.

RICHMOND
What men of name are slain on either side?

STANLEY
John, Duke of Norfolk, Walter, Lord Ferrers,
Sir Robert Brakenbury, and Sir William Brandon.

RICHMOND
Inter their bodies as becomes their births.
Proclaim a pardon to the soldiers fled
That in submission will return to us.
And then, as we have ta'en the sacrament,
We will unite the white rose and the red;
Smile heaven upon this fair conjunction,
That long have frowned upon their enmity.
What traitor hears me and says not "Amen"?
England hath long been mad and scarred herself:
The brother blindly shed the brother's blood;
The father rashly slaughtered his own son;
The son, compelled, been butcher to the sire.
All this divided York and Lancaster,
Divided in their dire division.
O, now let Richmond and Elizabeth,
The true succeeders of each royal house,
By God's fair ordinance conjoin together,
And let their heirs, God, if Thy will be so,
Enrich the time to come with smooth-faced peace,
With smiling plenty and fair prosperous days.
Abate the edge of traitors, gracious Lord,
That would reduce these bloody days again
And make poor England weep in streams of blood.
Let them not live to taste this land's increase,
That would with treason wound this fair land's peace.
Now civil wounds are stopped, peace lives again.
That she may long live here, God say amen.
[They exit.]
`

	b["the-tempest_TXT_FolgerShakespeare.txt"] = `
The Tempest
by William Shakespeare
Edited by Barbara A. Mowat and Paul Werstine
  with Michael Poston and Rebecca Niles
Folger Shakespeare Library
https://shakespeare.folger.edu/shakespeares-works/the-tempest/
Created on Jul 31, 2015, from FDT version 0.9.2

Characters in the Play
======================
PROSPERO, the former duke of Milan, now a magician on a Mediterranean island
MIRANDA, Prospero's daughter
ARIEL, a spirit, servant to Prospero
CALIBAN, an inhabitant of the island, servant to Prospero
FERDINAND, prince of Naples
ALONSO, king of Naples
ANTONIO, duke of Milan and Prospero's brother
SEBASTIAN, Alonso's brother
GONZALO, councillor to Alonso and friend to Prospero
Courtiers in attendance on Alonso:
  ADRIAN
  FRANCISCO
TRINCULO, servant to Alonso
STEPHANO, Alonso's butler
SHIPMASTER
BOATSWAIN
MARINERS
Players who, as spirits, take the roles of Iris, Ceres, Juno, Nymphs, and Reapers in Prospero's masque, and who, in other scenes, take the roles of "islanders" and of hunting dogs


ACT 1
=====

Scene 1
=======
[A tempestuous noise of thunder and lightning heard.
Enter a Shipmaster and a Boatswain.]


MASTER  Boatswain!

BOATSWAIN  Here, master. What cheer?

MASTER  Good, speak to th' mariners. Fall to 't yarely,
or we run ourselves aground. Bestir, bestir!
[He exits.]

[Enter Mariners.]


BOATSWAIN  Heigh, my hearts! Cheerly, cheerly, my
hearts! Yare, yare! Take in the topsail. Tend to th'
Master's whistle.--Blow till thou burst thy wind, if
room enough!

[Enter Alonso, Sebastian, Antonio, Ferdinand, Gonzalo,
and others.]


ALONSO  Good boatswain, have care. Where's the Master?
Play the men.

BOATSWAIN  I pray now, keep below.

ANTONIO  Where is the Master, boatswain?

BOATSWAIN  Do you not hear him? You mar our labor.
Keep your cabins. You do assist the storm.

GONZALO  Nay, good, be patient.

BOATSWAIN  When the sea is. Hence! What cares these
roarers for the name of king? To cabin! Silence!
Trouble us not.

GONZALO  Good, yet remember whom thou hast
aboard.

BOATSWAIN  None that I more love than myself. You are
a councillor; if you can command these elements
to silence, and work the peace of the present, we
will not hand a rope more. Use your authority. If
you cannot, give thanks you have lived so long, and
make yourself ready in your cabin for the mischance
of the hour, if it so hap.--Cheerly, good
hearts!--Out of our way, I say!	[He exits.]

GONZALO  I have great comfort from this fellow. Methinks
he hath no drowning mark upon him. His
complexion is perfect gallows. Stand fast, good
Fate, to his hanging. Make the rope of his destiny
our cable, for our own doth little advantage. If he be
not born to be hanged, our case is miserable.
[He exits with Alonso, Sebastian,
and the other courtiers.]

[Enter Boatswain.]


BOATSWAIN  Down with the topmast! Yare! Lower, lower!
Bring her to try wi' th' main course. [(A cry
within.)] A plague upon this howling! They are
louder than the weather or our office.

[Enter Sebastian, Antonio, and Gonzalo.]

Yet again? What do you here? Shall we give o'er and
drown? Have you a mind to sink?

SEBASTIAN  A pox o' your throat, you bawling, blasphemous,
incharitable dog!

BOATSWAIN  Work you, then.

ANTONIO  Hang, cur, hang, you whoreson, insolent
noisemaker! We are less afraid to be drowned than
thou art.

GONZALO  I'll warrant him for drowning, though the
ship were no stronger than a nutshell and as leaky
as an unstanched wench.

BOATSWAIN  Lay her ahold, ahold! Set her two courses.
Off to sea again! Lay her off!

[Enter more Mariners, wet.]


MARINERS  All lost! To prayers, to prayers! All lost!
[Mariners exit.]

BOATSWAIN  What, must our mouths be cold?

GONZALO  The King and Prince at prayers. Let's assist
them, for our case is as theirs.

SEBASTIAN  I am out of patience.

ANTONIO  We are merely cheated of our lives by drunkards.
This wide-chopped rascal--would thou
mightst lie drowning the washing of ten tides!
[Boatswain exits.]

GONZALO  He'll be hanged yet, though every drop of
water swear against it and gape at wid'st to glut him.


[A confused noise within:]  "Mercy on us!"--"We split, we
split!"--"Farewell, my wife and children!"--
"Farewell, brother!"--"We split, we split, we
split!"


ANTONIO  Let's all sink wi' th' King.

SEBASTIAN  Let's take leave of him.
[He exits with Antonio.]

GONZALO  Now would I give a thousand furlongs of sea
for an acre of barren ground: long heath, brown
furze, anything. The wills above be done, but I
would fain die a dry death.
[He exits.]

Scene 2
=======
[Enter Prospero and Miranda.]


MIRANDA
If by your art, my dearest father, you have
Put the wild waters in this roar, allay them.
The sky, it seems, would pour down stinking pitch,
But that the sea, mounting to th' welkin's cheek,
Dashes the fire out. O, I have suffered
With those that I saw suffer! A brave vessel,
Who had, no doubt, some noble creature in her,
Dashed all to pieces. O, the cry did knock
Against my very heart! Poor souls, they perished.
Had I been any god of power, I would
Have sunk the sea within the earth or ere
It should the good ship so have swallowed, and
The fraughting souls within her.

PROSPERO  Be collected.
No more amazement. Tell your piteous heart
There's no harm done.

MIRANDA  O, woe the day!

PROSPERO  No harm.
I have done nothing but in care of thee,
Of thee, my dear one, thee, my daughter, who
Art ignorant of what thou art, naught knowing
Of whence I am, nor that I am more better
Than Prospero, master of a full poor cell,
And thy no greater father.

MIRANDA  More to know
Did never meddle with my thoughts.

PROSPERO  'Tis time
I should inform thee farther. Lend thy hand
And pluck my magic garment from me.
[Putting aside his cloak.]
So,
Lie there, my art.--Wipe thou thine eyes. Have
comfort.
The direful spectacle of the wrack, which touched
The very virtue of compassion in thee,
I have with such provision in mine art
So safely ordered that there is no soul--
No, not so much perdition as an hair,
Betid to any creature in the vessel
Which thou heard'st cry, which thou saw'st sink. Sit
down,
For thou must now know farther.	[They sit.]

MIRANDA  You have often
Begun to tell me what I am, but stopped
And left me to a bootless inquisition,
Concluding "Stay. Not yet."

PROSPERO  The hour's now come.
The very minute bids thee ope thine ear.
Obey, and be attentive. Canst thou remember
A time before we came unto this cell?
I do not think thou canst, for then thou wast not
Out three years old.

MIRANDA  Certainly, sir, I can.

PROSPERO
By what? By any other house or person?
Of anything the image tell me that
Hath kept with thy remembrance.

MIRANDA  'Tis far off
And rather like a dream than an assurance
That my remembrance warrants. Had I not
Four or five women once that tended me?

PROSPERO
Thou hadst, and more, Miranda. But how is it
That this lives in thy mind? What seest thou else
In the dark backward and abysm of time?
If thou rememb'rest aught ere thou cam'st here,
How thou cam'st here thou mayst.

MIRANDA  But that I do not.

PROSPERO
Twelve year since, Miranda, twelve year since,
Thy father was the Duke of Milan and
A prince of power.

MIRANDA  Sir, are not you my father?

PROSPERO
Thy mother was a piece of virtue, and
She said thou wast my daughter. And thy father
Was Duke of Milan, and his only heir
And princess no worse issued.

MIRANDA  O, the heavens!
What foul play had we that we came from thence?
Or blessed was 't we did?

PROSPERO  Both, both, my girl.
By foul play, as thou sayst, were we heaved thence,
But blessedly holp hither.

MIRANDA  O, my heart bleeds
To think o' th' teen that I have turned you to,
Which is from my remembrance. Please you,
farther.

PROSPERO
My brother and thy uncle, called Antonio--
I pray thee, mark me--that a brother should
Be so perfidious!--he whom next thyself
Of all the world I loved, and to him put
The manage of my state, as at that time
Through all the signories it was the first,
And Prospero the prime duke, being so reputed
In dignity, and for the liberal arts
Without a parallel. Those being all my study,
The government I cast upon my brother
And to my state grew stranger, being transported
And rapt in secret studies. Thy false uncle--
Dost thou attend me?

MIRANDA  Sir, most heedfully.

PROSPERO
Being once perfected how to grant suits,
How to deny them, who t' advance, and who
To trash for overtopping, new created
The creatures that were mine, I say, or changed 'em,
Or else new formed 'em, having both the key
Of officer and office, set all hearts i' th' state
To what tune pleased his ear, that now he was
The ivy which had hid my princely trunk
And sucked my verdure out on 't. Thou attend'st not.

MIRANDA
O, good sir, I do.

PROSPERO  I pray thee, mark me.
I, thus neglecting worldly ends, all dedicated
To closeness and the bettering of my mind
With that which, but by being so retired,
O'erprized all popular rate, in my false brother
Awaked an evil nature, and my trust,
Like a good parent, did beget of him
A falsehood in its contrary as great
As my trust was, which had indeed no limit,
A confidence sans bound. He being thus lorded,
Not only with what my revenue yielded
But what my power might else exact, like one
Who, having into truth by telling of it,
Made such a sinner of his memory
To credit his own lie, he did believe
He was indeed the Duke, out o' th' substitution
And executing th' outward face of royalty
With all prerogative. Hence, his ambition growing--
Dost thou hear?

MIRANDA
Your tale, sir, would cure deafness.

PROSPERO
To have no screen between this part he played
And him he played it for, he needs will be
Absolute Milan. Me, poor man, my library
Was dukedom large enough. Of temporal royalties
He thinks me now incapable; confederates,
So dry he was for sway, wi' th' King of Naples
To give him annual tribute, do him homage,
Subject his coronet to his crown, and bend
The dukedom, yet unbowed--alas, poor Milan!--
To most ignoble stooping.

MIRANDA  O, the heavens!

PROSPERO
Mark his condition and th' event. Then tell me
If this might be a brother.

MIRANDA  I should sin
To think but nobly of my grandmother.
Good wombs have borne bad sons.

PROSPERO  Now the condition.
This King of Naples, being an enemy
To me inveterate, hearkens my brother's suit,
Which was that he, in lieu o' th' premises
Of homage and I know not how much tribute,
Should presently extirpate me and mine
Out of the dukedom, and confer fair Milan,
With all the honors, on my brother; whereon,
A treacherous army levied, one midnight
Fated to th' purpose did Antonio open
The gates of Milan, and i' th' dead of darkness
The ministers for th' purpose hurried thence
Me and thy crying self.

MIRANDA  Alack, for pity!
I, not rememb'ring how I cried out then,
Will cry it o'er again. It is a hint
That wrings mine eyes to 't.

PROSPERO  Hear a little further,
And then I'll bring thee to the present business
Which now 's upon 's, without the which this story
Were most impertinent.

MIRANDA  Wherefore did they not
That hour destroy us?

PROSPERO  Well demanded, wench.
My tale provokes that question. Dear, they durst not,
So dear the love my people bore me, nor set
A mark so bloody on the business, but
With colors fairer painted their foul ends.
In few, they hurried us aboard a bark,
Bore us some leagues to sea, where they prepared
A rotten carcass of a butt, not rigged,
Nor tackle, sail, nor mast; the very rats
Instinctively have quit it. There they hoist us
To cry to th' sea that roared to us, to sigh
To th' winds, whose pity, sighing back again,
Did us but loving wrong.

MIRANDA  Alack, what trouble
Was I then to you!

PROSPERO  O, a cherubin
Thou wast that did preserve me. Thou didst smile,
Infused with a fortitude from heaven,
When I have decked the sea with drops full salt,
Under my burden groaned, which raised in me
An undergoing stomach to bear up
Against what should ensue.

MIRANDA  How came we ashore?

PROSPERO  By providence divine.
Some food we had, and some fresh water, that
A noble Neapolitan, Gonzalo,
Out of his charity, who being then appointed
Master of this design, did give us, with
Rich garments, linens, stuffs, and necessaries,
Which since have steaded much. So, of his
gentleness,
Knowing I loved my books, he furnished me
From mine own library with volumes that
I prize above my dukedom.

MIRANDA  Would I might
But ever see that man.

PROSPERO, [standing]  Now I arise.
Sit still, and hear the last of our sea-sorrow.
Here in this island we arrived, and here
Have I, thy schoolmaster, made thee more profit
Than other princes can, that have more time
For vainer hours and tutors not so careful.

MIRANDA
Heavens thank you for 't. And now I pray you, sir--
For still 'tis beating in my mind--your reason
For raising this sea storm?

PROSPERO  Know thus far forth:
By accident most strange, bountiful Fortune,
Now my dear lady, hath mine enemies
Brought to this shore; and by my prescience
I find my zenith doth depend upon
A most auspicious star, whose influence
If now I court not, but omit, my fortunes
Will ever after droop. Here cease more questions.
Thou art inclined to sleep. 'Tis a good dullness,
And give it way. I know thou canst not choose.
[Miranda falls asleep.]
[Prospero puts on his cloak.]
Come away, servant, come. I am ready now.
Approach, my Ariel. Come.

[Enter Ariel.]


ARIEL
All hail, great master! Grave sir, hail! I come
To answer thy best pleasure. Be 't to fly,
To swim, to dive into the fire, to ride
On the curled clouds, to thy strong bidding task
Ariel and all his quality.

PROSPERO  Hast thou, spirit,
Performed to point the tempest that I bade thee?

ARIEL  To every article.
I boarded the King's ship; now on the beak,
Now in the waist, the deck, in every cabin,
I flamed amazement. Sometimes I'd divide
And burn in many places. On the topmast,
The yards, and bowsprit would I flame distinctly,
Then meet and join. Jove's lightning, the precursors
O' th' dreadful thunderclaps, more momentary
And sight-outrunning were not. The fire and cracks
Of sulfurous roaring the most mighty Neptune
Seem to besiege and make his bold waves tremble,
Yea, his dread trident shake.

PROSPERO  My brave spirit!
Who was so firm, so constant, that this coil
Would not infect his reason?

ARIEL  Not a soul
But felt a fever of the mad, and played
Some tricks of desperation. All but mariners
Plunged in the foaming brine and quit the vessel,
Then all afire with me. The King's son, Ferdinand,
With hair up-staring--then like reeds, not hair--
Was the first man that leaped; cried "Hell is empty,
And all the devils are here."

PROSPERO  Why, that's my spirit!
But was not this nigh shore?

ARIEL  Close by, my master.

PROSPERO
But are they, Ariel, safe?

ARIEL  Not a hair perished.
On their sustaining garments not a blemish,
But fresher than before; and, as thou bad'st me,
In troops I have dispersed them 'bout the isle.
The King's son have I landed by himself,
Whom I left cooling of the air with sighs
In an odd angle of the isle, and sitting,
His arms in this sad knot.	[He folds his arms.]

PROSPERO  Of the King's ship,
The mariners say how thou hast disposed,
And all the rest o' th' fleet.

ARIEL  Safely in harbor
Is the King's ship. In the deep nook, where once
Thou called'st me up at midnight to fetch dew
From the still-vexed Bermoothes, there she's hid;
The mariners all under hatches stowed,
Who, with a charm joined to their suffered labor,
I have left asleep. And for the rest o' th' fleet,
Which I dispersed, they all have met again
And are upon the Mediterranean float,
Bound sadly home for Naples,
Supposing that they saw the King's ship wracked
And his great person perish.

PROSPERO  Ariel, thy charge
Exactly is performed. But there's more work.
What is the time o' th' day?

ARIEL  Past the mid season.

PROSPERO
At least two glasses. The time 'twixt six and now
Must by us both be spent most preciously.

ARIEL
Is there more toil? Since thou dost give me pains,
Let me remember thee what thou hast promised,
Which is not yet performed me.

PROSPERO  How now? Moody?
What is 't thou canst demand?

ARIEL  My liberty.

PROSPERO
Before the time be out? No more.

ARIEL  I prithee,
Remember I have done thee worthy service,
Told thee no lies, made no mistakings, served
Without or grudge or grumblings. Thou did promise
To bate me a full year.

PROSPERO  Dost thou forget
From what a torment I did free thee?

ARIEL  No.

PROSPERO
Thou dost, and think'st it much to tread the ooze
Of the salt deep,
To run upon the sharp wind of the North,
To do me business in the veins o' th' Earth
When it is baked with frost.

ARIEL  I do not, sir.

PROSPERO
Thou liest, malignant thing. Hast thou forgot
The foul witch Sycorax, who with age and envy
Was grown into a hoop? Hast thou forgot her?

ARIEL  No, sir.

PROSPERO
Thou hast. Where was she born? Speak. Tell me.

ARIEL
Sir, in Argier.

PROSPERO  O, was she so? I must
Once in a month recount what thou hast been,
Which thou forget'st. This damned witch Sycorax,
For mischiefs manifold, and sorceries terrible
To enter human hearing, from Argier,
Thou know'st, was banished. For one thing she did
They would not take her life. Is not this true?

ARIEL  Ay, sir.

PROSPERO
This blue-eyed hag was hither brought with child
And here was left by th' sailors. Thou, my slave,
As thou report'st thyself, was then her servant,
And for thou wast a spirit too delicate
To act her earthy and abhorred commands,
Refusing her grand hests, she did confine thee,
By help of her more potent ministers
And in her most unmitigable rage,
Into a cloven pine, within which rift
Imprisoned thou didst painfully remain
A dozen years; within which space she died
And left thee there, where thou didst vent thy groans
As fast as mill wheels strike. Then was this island
(Save for the son that she did litter here,
A freckled whelp, hag-born) not honored with
A human shape.

ARIEL  Yes, Caliban, her son.

PROSPERO
Dull thing, I say so; he, that Caliban
Whom now I keep in service. Thou best know'st
What torment I did find thee in. Thy groans
Did make wolves howl, and penetrate the breasts
Of ever-angry bears. It was a torment
To lay upon the damned, which Sycorax
Could not again undo. It was mine art,
When I arrived and heard thee, that made gape
The pine and let thee out.

ARIEL  I thank thee, master.

PROSPERO
If thou more murmur'st, I will rend an oak
And peg thee in his knotty entrails till
Thou hast howled away twelve winters.

ARIEL  Pardon, master.
I will be correspondent to command
And do my spriting gently.

PROSPERO  Do so, and after two days
I will discharge thee.

ARIEL  That's my noble master.
What shall I do? Say, what? What shall I do?

PROSPERO
Go make thyself like a nymph o' th' sea. Be subject
To no sight but thine and mine, invisible
To every eyeball else. Go, take this shape,
And hither come in 't. Go, hence with diligence!
[Ariel exits.]
Awake, dear heart, awake. Thou hast slept well.
Awake.	[Miranda wakes.]

MIRANDA  The strangeness of your story put
Heaviness in me.

PROSPERO  Shake it off. Come on,
We'll visit Caliban, my slave, who never
Yields us kind answer.

MIRANDA, [rising]  'Tis a villain, sir,
I do not love to look on.

PROSPERO  But, as 'tis,
We cannot miss him. He does make our fire,
Fetch in our wood, and serves in offices
That profit us.--What ho, slave, Caliban!
Thou earth, thou, speak!

CALIBAN, [within]  There's wood enough within.

PROSPERO
Come forth, I say. There's other business for thee.
Come, thou tortoise. When?

[Enter Ariel like a water nymph.]

Fine apparition! My quaint Ariel,
Hark in thine ear.	[He whispers to Ariel.]

ARIEL  My lord, it shall be done.	[He exits.]

PROSPERO, [to Caliban]
Thou poisonous slave, got by the devil himself
Upon thy wicked dam, come forth!

[Enter Caliban.]


CALIBAN
As wicked dew as e'er my mother brushed
With raven's feather from unwholesome fen
Drop on you both. A southwest blow on you
And blister you all o'er.

PROSPERO
For this, be sure, tonight thou shalt have cramps,
Side-stitches that shall pen thy breath up. Urchins
Shall forth at vast of night that they may work
All exercise on thee. Thou shalt be pinched
As thick as honeycomb, each pinch more stinging
Than bees that made 'em.

CALIBAN  I must eat my dinner.
This island's mine by Sycorax, my mother,
Which thou tak'st from me. When thou cam'st first,
Thou strok'st me and made much of me, wouldst
give me
Water with berries in 't, and teach me how
To name the bigger light and how the less,
That burn by day and night. And then I loved thee,
And showed thee all the qualities o' th' isle,
The fresh springs, brine pits, barren place and
fertile.
Cursed be I that did so! All the charms
Of Sycorax, toads, beetles, bats, light on you,
For I am all the subjects that you have,
Which first was mine own king; and here you sty me
In this hard rock, whiles you do keep from me
The rest o' th' island.

PROSPERO  Thou most lying slave,
Whom stripes may move, not kindness, I have used
thee,
Filth as thou art, with humane care, and lodged
thee
In mine own cell, till thou didst seek to violate
The honor of my child.

CALIBAN
O ho, O ho! Would 't had been done!
Thou didst prevent me. I had peopled else
This isle with Calibans.

MIRANDA  Abhorred slave,
Which any print of goodness wilt not take,
Being capable of all ill! I pitied thee,
Took pains to make thee speak, taught thee each
hour
One thing or other. When thou didst not, savage,
Know thine own meaning, but wouldst gabble like
A thing most brutish, I endowed thy purposes
With words that made them known. But thy vile
race,
Though thou didst learn, had that in 't which good
natures
Could not abide to be with. Therefore wast thou
Deservedly confined into this rock,
Who hadst deserved more than a prison.

CALIBAN
You taught me language, and my profit on 't
Is I know how to curse. The red plague rid you
For learning me your language!

PROSPERO  Hagseed, hence!
Fetch us in fuel; and be quick, thou 'rt best,
To answer other business. Shrugg'st thou, malice?
If thou neglect'st or dost unwillingly
What I command, I'll rack thee with old cramps,
Fill all thy bones with aches, make thee roar
That beasts shall tremble at thy din.

CALIBAN  No, pray thee.
[Aside.] I must obey. His art is of such power
It would control my dam's god, Setebos,
And make a vassal of him.

PROSPERO  So, slave, hence.
[Caliban exits.]

[Enter Ferdinand; and Ariel, invisible,
playing and singing.]

Song.

ARIEL
	Come unto these yellow sands,
	   And then take hands.
	Curtsied when you have, and kissed
	   The wild waves whist.
	Foot it featly here and there,
	   And sweet sprites bear
	The burden. Hark, hark!
	      [Burden dispersedly, within:] Bow-wow.
	   The watchdogs bark.
	      [Burden dispersedly, within:] Bow-wow.
	Hark, hark! I hear
	The strain of strutting chanticleer
	   Cry cock-a-diddle-dow.

FERDINAND
Where should this music be? I' th' air, or th' earth?
It sounds no more; and sure it waits upon
Some god o' th' island. Sitting on a bank,
Weeping again the King my father's wrack,
This music crept by me upon the waters,
Allaying both their fury and my passion
With its sweet air. Thence I have followed it,
Or it hath drawn me rather. But 'tis gone.
No, it begins again.

Song.

ARIEL
	Full fathom five thy father lies.
	   Of his bones are coral made.
	Those are pearls that were his eyes.
	   Nothing of him that doth fade
	But doth suffer a sea change
	Into something rich and strange.
	Sea nymphs hourly ring his knell.
		[Burden, within:] Ding dong.
	Hark, now I hear them: ding dong bell.

FERDINAND
The ditty does remember my drowned father.
This is no mortal business, nor no sound
That the Earth owes. I hear it now above me.

PROSPERO, [to Miranda]
The fringed curtains of thine eye advance
And say what thou seest yond.

MIRANDA  What is 't? A spirit?
Lord, how it looks about! Believe me, sir,
It carries a brave form. But 'tis a spirit.

PROSPERO
No, wench, it eats and sleeps and hath such senses
As we have, such. This gallant which thou seest
Was in the wrack; and, but he's something stained
With grief--that's beauty's canker--thou might'st
call him
A goodly person. He hath lost his fellows
And strays about to find 'em.

MIRANDA  I might call him
A thing divine, for nothing natural
I ever saw so noble.

PROSPERO, [aside]  It goes on, I see,
As my soul prompts it. [To Ariel.] Spirit, fine spirit,
I'll free thee
Within two days for this.

FERDINAND, [seeing Miranda]  Most sure, the goddess
On whom these airs attend!--Vouchsafe my prayer
May know if you remain upon this island,
And that you will some good instruction give
How I may bear me here. My prime request,
Which I do last pronounce, is--O you wonder!--
If you be maid or no.

MIRANDA  No wonder, sir,
But certainly a maid.

FERDINAND  My language! Heavens!
I am the best of them that speak this speech,
Were I but where 'tis spoken.

PROSPERO  How? The best?
What wert thou if the King of Naples heard thee?

FERDINAND
A single thing, as I am now, that wonders
To hear thee speak of Naples. He does hear me,
And that he does I weep. Myself am Naples,
Who with mine eyes, never since at ebb, beheld
The King my father wracked.

MIRANDA  Alack, for mercy!

FERDINAND
Yes, faith, and all his lords, the Duke of Milan
And his brave son being twain.

PROSPERO, [aside]  The Duke of Milan
And his more braver daughter could control thee,
If now 'twere fit to do 't. At the first sight
They have changed eyes.--Delicate Ariel,
I'll set thee free for this. [To Ferdinand.] A word,
good sir.
I fear you have done yourself some wrong. A word.

MIRANDA
Why speaks my father so ungently? This
Is the third man that e'er I saw, the first
That e'er I sighed for. Pity move my father
To be inclined my way.

FERDINAND  O, if a virgin,
And your affection not gone forth, I'll make you
The Queen of Naples.

PROSPERO  Soft, sir, one word more.
[Aside.] They are both in either's powers. But this
swift business
I must uneasy make, lest too light winning
Make the prize light. [To Ferdinand.] One word
more. I charge thee
That thou attend me. Thou dost here usurp
The name thou ow'st not, and hast put thyself
Upon this island as a spy, to win it
From me, the lord on 't.

FERDINAND  No, as I am a man!

MIRANDA
There's nothing ill can dwell in such a temple.
If the ill spirit have so fair a house,
Good things will strive to dwell with 't.

PROSPERO, [to Ferdinand]  Follow me.
[To Miranda.] Speak not you for him. He's a traitor.
[To Ferdinand.] Come,
I'll manacle thy neck and feet together.
Sea water shalt thou drink. Thy food shall be
The fresh-brook mussels, withered roots, and husks
Wherein the acorn cradled. Follow.

FERDINAND  No,
I will resist such entertainment till
Mine enemy has more power.
[He draws, and is charmed from moving.]

MIRANDA  O dear father,
Make not too rash a trial of him, for
He's gentle and not fearful.

PROSPERO  What, I say,
My foot my tutor?--Put thy sword up, traitor,
Who mak'st a show, but dar'st not strike, thy
conscience
Is so possessed with guilt. Come from thy ward,
For I can here disarm thee with this stick
And make thy weapon drop.

MIRANDA  Beseech you, father--

PROSPERO
Hence! Hang not on my garments.

MIRANDA  Sir, have pity.
I'll be his surety.

PROSPERO  Silence! One word more
Shall make me chide thee, if not hate thee. What,
An advocate for an impostor? Hush.
Thou think'st there is no more such shapes as he,
Having seen but him and Caliban. Foolish wench,
To th' most of men this is a Caliban,
And they to him are angels.

MIRANDA  My affections
Are then most humble. I have no ambition
To see a goodlier man.

PROSPERO, [to Ferdinand]  Come on, obey.
Thy nerves are in their infancy again
And have no vigor in them.

FERDINAND  So they are.
My spirits, as in a dream, are all bound up.
My father's loss, the weakness which I feel,
The wrack of all my friends, nor this man's threats
To whom I am subdued, are but light to me,
Might I but through my prison once a day
Behold this maid. All corners else o' th' Earth
Let liberty make use of. Space enough
Have I in such a prison.

PROSPERO, [aside]  It works.--Come on.--
Thou hast done well, fine Ariel.--Follow me.
[To Ariel.] Hark what thou else shalt do me.

MIRANDA, [to Ferdinand]  Be of
comfort.
My father's of a better nature, sir,
Than he appears by speech. This is unwonted
Which now came from him.

PROSPERO, [to Ariel]  Thou shalt be as free
As mountain winds; but then exactly do
All points of my command.

ARIEL  To th' syllable.

PROSPERO, [to Ferdinand]
Come follow. [To Miranda.] Speak not for him.
[They exit.]


ACT 2
=====

Scene 1
=======
[Enter Alonso, Sebastian, Antonio, Gonzalo, Adrian,
Francisco, and others.]


GONZALO, [to Alonso]
Beseech you, sir, be merry. You have cause--
So have we all--of joy, for our escape
Is much beyond our loss. Our hint of woe
Is common; every day some sailor's wife,
The masters of some merchant, and the merchant
Have just our theme of woe. But for the miracle--
I mean our preservation--few in millions
Can speak like us. Then wisely, good sir, weigh
Our sorrow with our comfort.

ALONSO  Prithee, peace.

SEBASTIAN, [aside to Antonio]  He receives comfort like
cold porridge.

ANTONIO  The visitor will not give him o'er so.

SEBASTIAN  Look, he's winding up the watch of his wit.
By and by it will strike.

GONZALO, [to Alonso]  Sir--

SEBASTIAN  One. Tell.

GONZALO  When every grief is entertained that's offered,
comes to th' entertainer--

SEBASTIAN  A dollar.

GONZALO  Dolor comes to him indeed. You have spoken
truer than you purposed.

SEBASTIAN  You have taken it wiselier than I meant you
should.

GONZALO, [to Alonso]  Therefore, my lord--

ANTONIO  Fie, what a spendthrift is he of his tongue.

ALONSO, [to Gonzalo]  I prithee, spare.

GONZALO  Well, I have done. But yet--

SEBASTIAN, [aside to Antonio]  He will be talking.

ANTONIO, [aside to Sebastian]  Which, of he or Adrian,
for a good wager, first begins to crow?

SEBASTIAN  The old cock.

ANTONIO  The cockerel.

SEBASTIAN  Done. The wager?

ANTONIO  A laughter.

SEBASTIAN  A match!

ADRIAN  Though this island seem to be desert--

ANTONIO  Ha, ha, ha.

SEBASTIAN  So. You're paid.

ADRIAN  Uninhabitable and almost inaccessible--

SEBASTIAN  Yet--

ADRIAN  Yet--

ANTONIO  He could not miss 't.

ADRIAN  It must needs be of subtle, tender, and delicate
temperance.

ANTONIO  Temperance was a delicate wench.

SEBASTIAN  Ay, and a subtle, as he most learnedly
delivered.

ADRIAN  The air breathes upon us here most sweetly.

SEBASTIAN  As if it had lungs, and rotten ones.

ANTONIO  Or as 'twere perfumed by a fen.

GONZALO  Here is everything advantageous to life.

ANTONIO  True, save means to live.

SEBASTIAN  Of that there's none, or little.

GONZALO  How lush and lusty the grass looks! How
green!

ANTONIO  The ground indeed is tawny.

SEBASTIAN  With an eye of green in 't.

ANTONIO  He misses not much.

SEBASTIAN  No, he doth but mistake the truth totally.

GONZALO  But the rarity of it is, which is indeed almost
beyond credit--

SEBASTIAN  As many vouched rarities are.

GONZALO  That our garments, being, as they were,
drenched in the sea, hold notwithstanding their
freshness and gloss, being rather new-dyed than
stained with salt water.

ANTONIO  If but one of his pockets could speak, would
it not say he lies?

SEBASTIAN  Ay, or very falsely pocket up his report.

GONZALO  Methinks our garments are now as fresh as
when we put them on first in Afric, at the marriage
of the King's fair daughter Claribel to the King of
Tunis.

SEBASTIAN  'Twas a sweet marriage, and we prosper
well in our return.

ADRIAN  Tunis was never graced before with such a
paragon to their queen.

GONZALO  Not since widow Dido's time.

ANTONIO  Widow? A pox o' that! How came that "widow"
in? Widow Dido!

SEBASTIAN  What if he had said "widower Aeneas" too?
Good Lord, how you take it!

ADRIAN, [to Gonzalo]  "Widow Dido," said you? You
make me study of that. She was of Carthage, not of
Tunis.

GONZALO  This Tunis, sir, was Carthage.

ADRIAN  Carthage?

GONZALO  I assure you, Carthage.

ANTONIO  His word is more than the miraculous harp.

SEBASTIAN  He hath raised the wall, and houses too.

ANTONIO  What impossible matter will he make easy
next?

SEBASTIAN  I think he will carry this island home in his
pocket and give it his son for an apple.

ANTONIO  And sowing the kernels of it in the sea, bring
forth more islands.

GONZALO  Ay.

ANTONIO  Why, in good time.

GONZALO, [to Alonso]  Sir, we were talking that our
garments seem now as fresh as when we were at
Tunis at the marriage of your daughter, who is now
queen.

ANTONIO  And the rarest that e'er came there.

SEBASTIAN  Bate, I beseech you, widow Dido.

ANTONIO  O, widow Dido? Ay, widow Dido.

GONZALO, [to Alonso]  Is not, sir, my doublet as fresh as
the first day I wore it? I mean, in a sort.

ANTONIO  That "sort" was well fished for.

GONZALO, [to Alonso]  When I wore it at your daughter's
marriage.

ALONSO
You cram these words into mine ears against
The stomach of my sense. Would I had never
Married my daughter there, for coming thence
My son is lost, and, in my rate, she too,
Who is so far from Italy removed
I ne'er again shall see her.--O, thou mine heir
Of Naples and of Milan, what strange fish
Hath made his meal on thee?

FRANCISCO  Sir, he may live.
I saw him beat the surges under him
And ride upon their backs. He trod the water,
Whose enmity he flung aside, and breasted
The surge most swoll'n that met him. His bold head
'Bove the contentious waves he kept, and oared
Himself with his good arms in lusty stroke
To th' shore, that o'er his wave-worn basis bowed,
As stooping to relieve him. I not doubt
He came alive to land.

ALONSO  No, no, he's gone.

SEBASTIAN
Sir, you may thank yourself for this great loss,
That would not bless our Europe with your daughter,
But rather lose her to an African,
Where she at least is banished from your eye,
Who hath cause to wet the grief on 't.

ALONSO  Prithee, peace.

SEBASTIAN
You were kneeled to and importuned otherwise
By all of us; and the fair soul herself
Weighed between loathness and obedience at
Which end o' th' beam should bow. We have lost
your son,
I fear, forever. Milan and Naples have
More widows in them of this business' making
Than we bring men to comfort them.
The fault's your own.

ALONSO  So is the dear'st o' th' loss.

GONZALO  My lord Sebastian,
The truth you speak doth lack some gentleness
And time to speak it in. You rub the sore
When you should bring the plaster.

SEBASTIAN  Very well.

ANTONIO  And most chirurgeonly.

GONZALO, [to Alonso]
It is foul weather in us all, good sir,
When you are cloudy.

SEBASTIAN  Foul weather?

ANTONIO  Very foul.

GONZALO
Had I plantation of this isle, my lord--

ANTONIO
He'd sow 't with nettle seed.

SEBASTIAN  Or docks, or mallows.

GONZALO
And were the king on 't, what would I do?

SEBASTIAN  Scape being drunk, for want of wine.

GONZALO
I' th' commonwealth I would by contraries
Execute all things, for no kind of traffic
Would I admit; no name of magistrate;
Letters should not be known; riches, poverty,
And use of service, none; contract, succession,
Bourn, bound of land, tilth, vineyard, none;
No use of metal, corn, or wine, or oil;
No occupation; all men idle, all,
And women too, but innocent and pure;
No sovereignty--

SEBASTIAN  Yet he would be king on 't.

ANTONIO  The latter end of his commonwealth forgets
the beginning.

GONZALO
All things in common nature should produce
Without sweat or endeavor; treason, felony,
Sword, pike, knife, gun, or need of any engine
Would I not have; but nature should bring forth
Of its own kind all foison, all abundance,
To feed my innocent people.

SEBASTIAN  No marrying 'mong his subjects?

ANTONIO  None, man, all idle: whores and knaves.

GONZALO
I would with such perfection govern, sir,
T' excel the Golden Age.

SEBASTIAN  'Save his Majesty!

ANTONIO
Long live Gonzalo!

GONZALO  And do you mark me, sir?

ALONSO
Prithee, no more. Thou dost talk nothing to me.

GONZALO  I do well believe your Highness, and did it to
minister occasion to these gentlemen, who are of
such sensible and nimble lungs that they always use
to laugh at nothing.

ANTONIO  'Twas you we laughed at.

GONZALO  Who in this kind of merry fooling am
nothing to you. So you may continue, and laugh at
nothing still.

ANTONIO  What a blow was there given!

SEBASTIAN  An it had not fallen flatlong.

GONZALO  You are gentlemen of brave mettle. You
would lift the moon out of her sphere if she would
continue in it five weeks without changing.

[Enter Ariel invisible, playing solemn music.]


SEBASTIAN  We would so, and then go a-batfowling.

ANTONIO, [to Gonzalo]  Nay, good my lord, be not angry.

GONZALO  No, I warrant you, I will not adventure my
discretion so weakly. Will you laugh me asleep?
For I am very heavy.

ANTONIO  Go sleep, and hear us.
[All sink down asleep except Alonso,
Antonio, and Sebastian.]

ALONSO
What, all so soon asleep? I wish mine eyes
Would, with themselves, shut up my thoughts. I find
They are inclined to do so.

SEBASTIAN  Please you, sir,
Do not omit the heavy offer of it.
It seldom visits sorrow; when it doth,
It is a comforter.

ANTONIO  We two, my lord,
Will guard your person while you take your rest,
And watch your safety.

ALONSO  Thank you. Wondrous heavy.
[Alonso sleeps. Ariel exits.]

SEBASTIAN
What a strange drowsiness possesses them!

ANTONIO
It is the quality o' th' climate.

SEBASTIAN  Why
Doth it not then our eyelids sink? I find
Not myself disposed to sleep.

ANTONIO  Nor I. My spirits are nimble.
They fell together all, as by consent.
They dropped as by a thunderstroke. What might,
Worthy Sebastian, O, what might--? No more.
And yet methinks I see it in thy face
What thou shouldst be. Th' occasion speaks thee, and
My strong imagination sees a crown
Dropping upon thy head.

SEBASTIAN  What, art thou waking?

ANTONIO
Do you not hear me speak?

SEBASTIAN  I do, and surely
It is a sleepy language, and thou speak'st
Out of thy sleep. What is it thou didst say?
This is a strange repose, to be asleep
With eyes wide open--standing, speaking, moving--
And yet so fast asleep.

ANTONIO  Noble Sebastian,
Thou let'st thy fortune sleep, die rather, wink'st
Whiles thou art waking.

SEBASTIAN  Thou dost snore distinctly.
There's meaning in thy snores.

ANTONIO
I am more serious than my custom. You
Must be so too, if heed me; which to do
Trebles thee o'er.

SEBASTIAN  Well, I am standing water.

ANTONIO
I'll teach you how to flow.

SEBASTIAN  Do so. To ebb
Hereditary sloth instructs me.

ANTONIO  O,
If you but knew how you the purpose cherish
Whiles thus you mock it, how in stripping it
You more invest it. Ebbing men indeed
Most often do so near the bottom run
By their own fear or sloth.

SEBASTIAN  Prithee, say on.
The setting of thine eye and cheek proclaim
A matter from thee, and a birth indeed
Which throes thee much to yield.

ANTONIO  Thus, sir:
Although this lord of weak remembrance--this,
Who shall be of as little memory
When he is earthed--hath here almost persuaded--
For he's a spirit of persuasion, only
Professes to persuade--the King his son's alive,
'Tis as impossible that he's undrowned
As he that sleeps here swims.

SEBASTIAN  I have no hope
That he's undrowned.

ANTONIO  O, out of that no hope
What great hope have you! No hope that way is
Another way so high a hope that even
Ambition cannot pierce a wink beyond,
But doubt discovery there. Will you grant with me
That Ferdinand is drowned?

SEBASTIAN  He's gone.

ANTONIO  Then tell me,
Who's the next heir of Naples?

SEBASTIAN  Claribel.

ANTONIO
She that is Queen of Tunis; she that dwells
Ten leagues beyond man's life; she that from Naples
Can have no note, unless the sun were post--
The man i' th' moon's too slow--till newborn chins
Be rough and razorable; she that from whom
We all were sea-swallowed, though some cast again,
And by that destiny to perform an act
Whereof what's past is prologue, what to come
In yours and my discharge.

SEBASTIAN  What stuff is this? How say you?
'Tis true my brother's daughter's Queen of Tunis,
So is she heir of Naples, 'twixt which regions
There is some space.

ANTONIO  A space whose ev'ry cubit
Seems to cry out "How shall that Claribel
Measure us back to Naples? Keep in Tunis
And let Sebastian wake." Say this were death
That now hath seized them, why, they were no worse
Than now they are. There be that can rule Naples
As well as he that sleeps, lords that can prate
As amply and unnecessarily
As this Gonzalo. I myself could make
A chough of as deep chat. O, that you bore
The mind that I do, what a sleep were this
For your advancement! Do you understand me?

SEBASTIAN
Methinks I do.

ANTONIO  And how does your content
Tender your own good fortune?

SEBASTIAN  I remember
You did supplant your brother Prospero.

ANTONIO  True,
And look how well my garments sit upon me,
Much feater than before. My brother's servants
Were then my fellows; now they are my men.

SEBASTIAN  But, for your conscience?

ANTONIO
Ay, sir, where lies that? If 'twere a kibe,
'Twould put me to my slipper, but I feel not
This deity in my bosom. Twenty consciences
That stand 'twixt me and Milan, candied be they
And melt ere they molest! Here lies your brother,
No better than the earth he lies upon.
If he were that which now he's like--that's dead--
Whom I with this obedient steel, three inches of it,
Can lay to bed forever; whiles you, doing thus,
To the perpetual wink for aye might put
This ancient morsel, this Sir Prudence, who
Should not upbraid our course. For all the rest,
They'll take suggestion as a cat laps milk.
They'll tell the clock to any business that
We say befits the hour.

SEBASTIAN  Thy case, dear friend,
Shall be my precedent: as thou got'st Milan,
I'll come by Naples. Draw thy sword. One stroke
Shall free thee from the tribute which thou payest,
And I the King shall love thee.

ANTONIO  Draw together,
And when I rear my hand, do you the like
To fall it on Gonzalo.	[They draw their swords.]

SEBASTIAN  O, but one word.
[They talk apart.]

[Enter Ariel, invisible, with music and song.]


ARIEL, [to the sleeping Gonzalo]
My master through his art foresees the danger
That you, his friend, are in, and sends me forth--
For else his project dies--to keep them living.
[Sings in Gonzalo's ear:]
	While you here do snoring lie,
	Open-eyed conspiracy
	   His time doth take.
	If of life you keep a care,
	Shake off slumber and beware.
	   Awake, awake!

ANTONIO, [to Sebastian]  Then let us both be sudden.

GONZALO, [waking]  Now, good angels preserve the
King!	[He wakes Alonso.]

ALONSO, [to Sebastian]
Why, how now, ho! Awake? Why are you drawn?
Wherefore this ghastly looking?

GONZALO, [to Sebastian]  What's the matter?

SEBASTIAN
Whiles we stood here securing your repose,
Even now, we heard a hollow burst of bellowing
Like bulls, or rather lions. Did 't not wake you?
It struck mine ear most terribly.

ALONSO  I heard nothing.

ANTONIO
O, 'twas a din to fright a monster's ear,
To make an earthquake. Sure, it was the roar
Of a whole herd of lions.

ALONSO  Heard you this, Gonzalo?

GONZALO
Upon mine honor, sir, I heard a humming,
And that a strange one too, which did awake me.
I shaked you, sir, and cried. As mine eyes opened,
I saw their weapons drawn. There was a noise,
That's verily. 'Tis best we stand upon our guard,
Or that we quit this place. Let's draw our weapons.

ALONSO
Lead off this ground, and let's make further search
For my poor son.

GONZALO  Heavens keep him from these beasts,
For he is, sure, i' th' island.

ALONSO  Lead away.

ARIEL, [aside]
Prospero my lord shall know what I have done.
So, king, go safely on to seek thy son.
[They exit.]

Scene 2
=======
[Enter Caliban with a burden of wood. A noise of
thunder heard.]


CALIBAN
All the infections that the sun sucks up
From bogs, fens, flats, on Prosper fall and make him
By inchmeal a disease! His spirits hear me,
And yet I needs must curse. But they'll nor pinch,
Fright me with urchin-shows, pitch me i' th' mire,
Nor lead me like a firebrand in the dark
Out of my way, unless he bid 'em. But
For every trifle are they set upon me,
Sometimes like apes, that mow and chatter at me
And after bite me; then like hedgehogs, which
Lie tumbling in my barefoot way and mount
Their pricks at my footfall. Sometime am I
All wound with adders, who with cloven tongues
Do hiss me into madness. Lo, now, lo!
Here comes a spirit of his, and to torment me
For bringing wood in slowly. I'll fall flat.
Perchance he will not mind me.
[He lies down and covers himself with a cloak.]

[Enter Trinculo.]


TRINCULO  Here's neither bush nor shrub to bear off
any weather at all. And another storm brewing; I
hear it sing i' th' wind. Yond same black cloud, yond
huge one, looks like a foul bombard that would shed
his liquor. If it should thunder as it did before, I
know not where to hide my head. Yond same cloud
cannot choose but fall by pailfuls. [Noticing Caliban.]
What have we here, a man or a fish? Dead or
alive? A fish, he smells like a fish--a very ancient
and fishlike smell, a kind of not-of-the-newest poor-John.
A strange fish. Were I in England now, as once
I was, and had but this fish painted, not a holiday
fool there but would give a piece of silver. There
would this monster make a man. Any strange beast
there makes a man. When they will not give a doit to
relieve a lame beggar, they will lay out ten to see a
dead Indian. Legged like a man, and his fins like
arms! Warm, o' my troth! I do now let loose my
opinion, hold it no longer: this is no fish, but an
islander that hath lately suffered by a thunderbolt.
[Thunder.] Alas, the storm is come again. My best
way is to creep under his gaberdine. There is no
other shelter hereabout. Misery acquaints a man
with strange bedfellows. I will here shroud till the
dregs of the storm be past.
[He crawls under Caliban's cloak.]

[Enter Stephano singing.]


STEPHANO
	I shall no more to sea, to sea.
	Here shall I die ashore--
This is a very scurvy tune to sing at a man's funeral.
Well, here's my comfort.	[Drinks.]
[Sings.]

	The master, the swabber, the boatswain, and I,
	   The gunner and his mate,
	Loved Mall, Meg, and Marian, and Margery,
	   But none of us cared for Kate.
	      For she had a tongue with a tang,
	      Would cry to a sailor "Go hang!"
	She loved not the savor of tar nor of pitch,
	Yet a tailor might scratch her where'er she did itch.
	      Then to sea, boys, and let her go hang!
This is a scurvy tune too. But here's my comfort.
[Drinks.]

CALIBAN  Do not torment me! O!

STEPHANO  What's the matter? Have we devils here? Do
you put tricks upon 's with savages and men of Ind?
Ha? I have not scaped drowning to be afeard now
of your four legs, for it hath been said "As proper a
man as ever went on four legs cannot make him
give ground," and it shall be said so again while
Stephano breathes at' nostrils.

CALIBAN  The spirit torments me. O!

STEPHANO  This is some monster of the isle with four
legs, who hath got, as I take it, an ague. Where the
devil should he learn our language? I will give him
some relief, if it be but for that. If I can recover him
and keep him tame and get to Naples with him,
he's a present for any emperor that ever trod on
neat's leather.

CALIBAN  Do not torment me, prithee. I'll bring my
wood home faster.

STEPHANO  He's in his fit now, and does not talk after
the wisest. He shall taste of my bottle. If he have
never drunk wine afore, it will go near to remove
his fit. If I can recover him and keep him tame, I will
not take too much for him. He shall pay for him that
hath him, and that soundly.

CALIBAN  Thou dost me yet but little hurt. Thou wilt
anon; I know it by thy trembling. Now Prosper
works upon thee.

STEPHANO  Come on your ways. Open your mouth.
Here is that which will give language to you, cat.
Open your mouth. This will shake your shaking, I
can tell you, and that soundly. [Caliban drinks.] You
cannot tell who's your friend. Open your chaps
again.

TRINCULO  I should know that voice. It should be--but
he is drowned, and these are devils. O, defend me!

STEPHANO  Four legs and two voices--a most delicate
monster! His forward voice now is to speak well of
his friend. His backward voice is to utter foul
speeches and to detract. If all the wine in my bottle
will recover him, I will help his ague. Come.
[Caliban drinks.] Amen! I will pour some in thy
other mouth.

TRINCULO  Stephano!

STEPHANO  Doth thy other mouth call me? Mercy, mercy,
this is a devil, and no monster! I will leave him; I
have no long spoon.

TRINCULO  Stephano! If thou be'st Stephano, touch me
and speak to me, for I am Trinculo--be not
afeard--thy good friend Trinculo.

STEPHANO  If thou be'st Trinculo, come forth. I'll pull
thee by the lesser legs. If any be Trinculo's legs,
these are they. [He pulls him out from under Caliban's
cloak.] Thou art very Trinculo indeed. How
cam'st thou to be the siege of this mooncalf? Can
he vent Trinculos?

TRINCULO  I took him to be killed with a thunderstroke.
But art thou not drowned, Stephano? I
hope now thou art not drowned. Is the storm
overblown? I hid me under the dead mooncalf's
gaberdine for fear of the storm. And art thou living,
Stephano? O Stephano, two Neapolitans scaped!

STEPHANO  Prithee, do not turn me about. My stomach
is not constant.

CALIBAN, [aside]  These be fine things, an if they be not
sprites. That's a brave god and bears celestial liquor.
I will kneel to him.
[He crawls out from under the cloak.]

STEPHANO, [to Trinculo]  How didst thou scape? How
cam'st thou hither? Swear by this bottle how thou
cam'st hither--I escaped upon a butt of sack, which
the sailors heaved o'erboard--by this bottle, which
I made of the bark of a tree with mine own hands,
since I was cast ashore.

CALIBAN  I'll swear upon that bottle to be thy true
subject, for the liquor is not earthly.

STEPHANO, [to Trinculo]  Here. Swear then how thou
escapedst.

TRINCULO  Swum ashore, man, like a duck. I can swim
like a duck, I'll be sworn.

STEPHANO  Here, kiss the book.	[Trinculo drinks.]
Though thou canst swim like a duck, thou art made
like a goose.

TRINCULO  O Stephano, hast any more of this?

STEPHANO  The whole butt, man. My cellar is in a rock
by th' seaside, where my wine is hid.--How now,
mooncalf, how does thine ague?

CALIBAN  Hast thou not dropped from heaven?

STEPHANO  Out o' th' moon, I do assure thee. I was the
man i' th' moon when time was.

CALIBAN  I have seen thee in her, and I do adore thee.
My mistress showed me thee, and thy dog, and thy
bush.

STEPHANO  Come, swear to that. Kiss the book. I will
furnish it anon with new contents. Swear.
[Caliban drinks.]

TRINCULO  By this good light, this is a very shallow
monster. I afeard of him? A very weak monster. The
man i' th' moon? A most poor, credulous monster!
--Well drawn, monster, in good sooth!

CALIBAN  I'll show thee every fertile inch o' th' island,
and I will kiss thy foot. I prithee, be my god.

TRINCULO  By this light, a most perfidious and drunken
monster. When 's god's asleep, he'll rob his bottle.

CALIBAN  I'll kiss thy foot. I'll swear myself thy subject.

STEPHANO  Come on, then. Down, and swear.
[Caliban kneels.]

TRINCULO  I shall laugh myself to death at this puppy-headed
monster. A most scurvy monster. I could
find in my heart to beat him--

STEPHANO  Come, kiss.

TRINCULO  --but that the poor monster's in drink. An
abominable monster.

CALIBAN
I'll show thee the best springs. I'll pluck thee berries.
I'll fish for thee and get thee wood enough.
A plague upon the tyrant that I serve.
I'll bear him no more sticks, but follow thee,
Thou wondrous man.

TRINCULO  A most ridiculous monster, to make a wonder
of a poor drunkard.

CALIBAN, [standing]
I prithee, let me bring thee where crabs grow,
And I with my long nails will dig thee pignuts,
Show thee a jay's nest, and instruct thee how
To snare the nimble marmoset. I'll bring thee
To clustering filberts, and sometimes I'll get thee
Young scamels from the rock. Wilt thou go with me?

STEPHANO  I prithee now, lead the way without any
more talking.--Trinculo, the King and all our
company else being drowned, we will inherit here.
--Here, bear my bottle.--Fellow Trinculo, we'll
fill him by and by again.

CALIBAN [sings drunkenly]
	Farewell, master, farewell, farewell.

TRINCULO  A howling monster, a drunken monster.

CALIBAN [sings]
	No more dams I'll make for fish,
	   Nor fetch in firing
	   At requiring,
	Nor scrape trenchering, nor wash dish.
	   'Ban, 'ban, Ca-caliban
	   Has a new master. Get a new man.
Freedom, high-day! High-day, freedom! Freedom,
high-day, freedom!

STEPHANO  O brave monster! Lead the way.
[They exit.]


ACT 3
=====

Scene 1
=======
[Enter Ferdinand bearing a log.]


FERDINAND
There be some sports are painful, and their labor
Delight in them sets off; some kinds of baseness
Are nobly undergone; and most poor matters
Point to rich ends. This my mean task
Would be as heavy to me as odious, but
The mistress which I serve quickens what's dead
And makes my labors pleasures. O, she is
Ten times more gentle than her father's crabbed,
And he's composed of harshness. I must remove
Some thousands of these logs and pile them up,
Upon a sore injunction. My sweet mistress
Weeps when she sees me work, and says such
baseness
Had never like executor. I forget;
But these sweet thoughts do even refresh my labors,
Most busiest when I do it.

[Enter Miranda; and Prospero at a distance, unobserved.]


MIRANDA  Alas now, pray you,
Work not so hard. I would the lightning had
Burnt up those logs that you are enjoined to pile.
Pray, set it down and rest you. When this burns
'Twill weep for having wearied you. My father
Is hard at study. Pray now, rest yourself.
He's safe for these three hours.

FERDINAND  O most dear mistress,
The sun will set before I shall discharge
What I must strive to do.

MIRANDA  If you'll sit down,
I'll bear your logs the while. Pray, give me that.
I'll carry it to the pile.

FERDINAND  No, precious creature,
I had rather crack my sinews, break my back,
Than you should such dishonor undergo
While I sit lazy by.

MIRANDA  It would become me
As well as it does you, and I should do it
With much more ease, for my good will is to it,
And yours it is against.

PROSPERO, [aside]  Poor worm, thou art infected.
This visitation shows it.

MIRANDA  You look wearily.
FERDINAND
No, noble mistress, 'tis fresh morning with me
When you are by at night. I do beseech you,
Chiefly that I might set it in my prayers,
What is your name?

MIRANDA  Miranda.--O my father,
I have broke your hest to say so!

FERDINAND  Admired Miranda!
Indeed the top of admiration, worth
What's dearest to the world! Full many a lady
I have eyed with best regard, and many a time
Th' harmony of their tongues hath into bondage
Brought my too diligent ear. For several virtues
Have I liked several women, never any
With so full soul but some defect in her
Did quarrel with the noblest grace she owed,
And put it to the foil. But you, O you,
So perfect and so peerless, are created
Of every creature's best.

MIRANDA  I do not know
One of my sex, no woman's face remember,
Save, from my glass, mine own. Nor have I seen
More that I may call men than you, good friend,
And my dear father. How features are abroad
I am skilless of, but by my modesty,
The jewel in my dower, I would not wish
Any companion in the world but you,
Nor can imagination form a shape
Besides yourself to like of. But I prattle
Something too wildly, and my father's precepts
I therein do forget.

FERDINAND  I am in my condition
A prince, Miranda; I do think a king--
I would, not so!--and would no more endure
This wooden slavery than to suffer
The flesh-fly blow my mouth. Hear my soul speak:
The very instant that I saw you did
My heart fly to your service, there resides
To make me slave to it, and for your sake
Am I this patient log-man.

MIRANDA  Do you love me?

FERDINAND
O heaven, O Earth, bear witness to this sound,
And crown what I profess with kind event
If I speak true; if hollowly, invert
What best is boded me to mischief. I,
Beyond all limit of what else i' th' world,
Do love, prize, honor you.

MIRANDA  I am a fool
To weep at what I am glad of.

PROSPERO, [aside]  Fair encounter
Of two most rare affections. Heavens rain grace
On that which breeds between 'em!

FERDINAND  Wherefore
weep you?

MIRANDA
At mine unworthiness, that dare not offer
What I desire to give, and much less take
What I shall die to want. But this is trifling,
And all the more it seeks to hide itself,
The bigger bulk it shows. Hence, bashful cunning,
And prompt me, plain and holy innocence.
I am your wife if you will marry me.
If not, I'll die your maid. To be your fellow
You may deny me, but I'll be your servant
Whether you will or no.

FERDINAND
My mistress, dearest, and I thus humble ever.

MIRANDA
My husband, then?

FERDINAND  Ay, with a heart as willing
As bondage e'er of freedom. Here's my hand.

MIRANDA, [clasping his hand]
And mine, with my heart in 't. And now farewell
Till half an hour hence.

FERDINAND  A thousand thousand.
[They exit.]

PROSPERO
So glad of this as they I cannot be,
Who are surprised withal; but my rejoicing
At nothing can be more. I'll to my book,
For yet ere suppertime must I perform
Much business appertaining.
[He exits.]

Scene 2
=======
[Enter Caliban, Stephano, and Trinculo.]


STEPHANO, [to Trinculo]  Tell not me. When the butt is
out, we will drink water; not a drop before. Therefore
bear up and board 'em.--Servant monster,
drink to me.

TRINCULO  Servant monster? The folly of this island!
They say there's but five upon this isle; we are three
of them. If th' other two be brained like us, the state
totters.

STEPHANO  Drink, servant monster, when I bid thee.
Thy eyes are almost set in thy head.
[Caliban drinks.]

TRINCULO  Where should they be set else? He were a
brave monster indeed if they were set in his tail.

STEPHANO  My man-monster hath drowned his tongue
in sack. For my part, the sea cannot drown me. I
swam, ere I could recover the shore, five-and-thirty
leagues off and on, by this light.--Thou shalt be my
lieutenant, monster, or my standard.

TRINCULO  Your lieutenant, if you list. He's no
standard.

STEPHANO  We'll not run, Monsieur Monster.

TRINCULO  Nor go neither. But you'll lie like dogs, and
yet say nothing neither.

STEPHANO  Mooncalf, speak once in thy life, if thou
be'st a good mooncalf.

CALIBAN  How does thy Honor? Let me lick thy shoe. I'll
not serve him; he is not valiant.

TRINCULO  Thou liest, most ignorant monster. I am in
case to justle a constable. Why, thou debauched
fish, thou! Was there ever man a coward that hath
drunk so much sack as I today? Wilt thou tell a
monstrous lie, being but half a fish and half a
monster?

CALIBAN  Lo, how he mocks me! Wilt thou let him, my
lord?

TRINCULO  "Lord," quoth he? That a monster should be
such a natural!

CALIBAN  Lo, lo again! Bite him to death, I prithee.

STEPHANO  Trinculo, keep a good tongue in your head.
If you prove a mutineer, the next tree. The poor
monster's my subject, and he shall not suffer
indignity.

CALIBAN  I thank my noble lord. Wilt thou be pleased
to harken once again to the suit I made to thee?

STEPHANO  Marry, will I. Kneel and repeat it. I will
stand, and so shall Trinculo.

[Enter Ariel, invisible.]


CALIBAN, [kneeling]  As I told thee before, I am subject
to a tyrant, a sorcerer, that by his cunning hath
cheated me of the island.

ARIEL, [in Trinculo's voice]  Thou liest.

CALIBAN, [to Trinculo]  Thou liest, thou jesting monkey,
thou. [He stands.] I would my valiant master would
destroy thee. I do not lie.

STEPHANO  Trinculo, if you trouble him any more in 's
tale, by this hand, I will supplant some of your
teeth.

TRINCULO  Why, I said nothing.

STEPHANO  Mum then, and no more. [Trinculo stands
aside.] Proceed.

CALIBAN
I say by sorcery he got this isle;
From me he got it. If thy Greatness will,
Revenge it on him, for I know thou dar'st,
But this thing dare not.

STEPHANO  That's most certain.

CALIBAN
Thou shalt be lord of it, and I'll serve thee.

STEPHANO  How now shall this be compassed? Canst
thou bring me to the party?

CALIBAN
Yea, yea, my lord. I'll yield him thee asleep,
Where thou mayst knock a nail into his head.

ARIEL, [in Trinculo's voice]  Thou liest. Thou canst not.

CALIBAN
What a pied ninny's this!--Thou scurvy patch!--
I do beseech thy Greatness, give him blows
And take his bottle from him. When that's gone,
He shall drink naught but brine, for I'll not show him
Where the quick freshes are.

STEPHANO  Trinculo, run into no further danger. Interrupt
the monster one word further, and by this
hand, I'll turn my mercy out o' doors and make a
stockfish of thee.

TRINCULO  Why, what did I? I did nothing. I'll go
farther off.

STEPHANO  Didst thou not say he lied?

ARIEL, [in Trinculo's voice]  Thou liest.

STEPHANO  Do I so? Take thou that.	[He beats Trinculo.]
As you like this, give me the lie another time.

TRINCULO  I did not give the lie! Out o' your wits and
hearing too? A pox o' your bottle! This can sack and
drinking do. A murrain on your monster, and the
devil take your fingers!

CALIBAN  Ha, ha, ha!

STEPHANO  Now forward with your tale. [To Trinculo.]
Prithee, stand further off.

CALIBAN
Beat him enough. After a little time
I'll beat him too.

STEPHANO  Stand farther. [Trinculo moves farther
away.] Come, proceed.

CALIBAN
Why, as I told thee, 'tis a custom with him
I' th' afternoon to sleep. There thou mayst brain him,
Having first seized his books, or with a log
Batter his skull, or paunch him with a stake,
Or cut his weasand with thy knife. Remember
First to possess his books, for without them
He's but a sot, as I am, nor hath not
One spirit to command. They all do hate him
As rootedly as I. Burn but his books.
He has brave utensils--for so he calls them--
Which, when he has a house, he'll deck withal.
And that most deeply to consider is
The beauty of his daughter. He himself
Calls her a nonpareil. I never saw a woman
But only Sycorax my dam and she;
But she as far surpasseth Sycorax
As great'st does least.

STEPHANO  Is it so brave a lass?

CALIBAN
Ay, lord, she will become thy bed, I warrant,
And bring thee forth brave brood.

STEPHANO  Monster, I will kill this man. His daughter
and I will be king and queen--save our Graces!--
and Trinculo and thyself shall be viceroys.--Dost
thou like the plot, Trinculo?

TRINCULO  Excellent.

STEPHANO  Give me thy hand. I am sorry I beat thee.
But while thou liv'st, keep a good tongue in thy
head.

CALIBAN
Within this half hour will he be asleep.
Wilt thou destroy him then?

STEPHANO  Ay, on mine honor.

ARIEL, [aside]  This will I tell my master.

CALIBAN
Thou mak'st me merry. I am full of pleasure.
Let us be jocund. Will you troll the catch
You taught me but whilere?

STEPHANO  At thy request, monster, I will do reason,
any reason.--Come on, Trinculo, let us sing.
[Sings.]
	Flout 'em and cout 'em
	And scout 'em and flout 'em!
	   Thought is free.

CALIBAN  That's not the tune.
[Ariel plays the tune on a tabor and pipe.]

STEPHANO  What is this same?

TRINCULO  This is the tune of our catch played by the
picture of Nobody.

STEPHANO, [to the invisible musician]  If thou be'st a
man, show thyself in thy likeness. If thou be'st a
devil, take 't as thou list.

TRINCULO  O, forgive me my sins!

STEPHANO  He that dies pays all debts.--I defy thee!--
Mercy upon us!

CALIBAN  Art thou afeard?

STEPHANO  No, monster, not I.

CALIBAN
Be not afeard. The isle is full of noises,
Sounds and sweet airs that give delight and hurt not.
Sometimes a thousand twangling instruments
Will hum about mine ears, and sometimes voices
That, if I then had waked after long sleep,
Will make me sleep again; and then, in dreaming,
The clouds methought would open, and show riches
Ready to drop upon me, that when I waked
I cried to dream again.

STEPHANO  This will prove a brave kingdom to me,
where I shall have my music for nothing.

CALIBAN  When Prospero is destroyed.

STEPHANO  That shall be by and by. I remember the
story.

TRINCULO  The sound is going away. Let's follow it, and
after do our work.

STEPHANO  Lead, monster. We'll follow.--I would I
could see this taborer. He lays it on. Wilt come?

TRINCULO  I'll follow, Stephano.
[They exit.]

Scene 3
=======
[Enter Alonso, Sebastian, Antonio, Gonzalo, Adrian,
Francisco, etc.]


GONZALO
By 'r lakin, I can go no further, sir.
My old bones aches. Here's a maze trod indeed
Through forthrights and meanders. By your
patience,
I needs must rest me.

ALONSO  Old lord, I cannot blame thee.
Who am myself attached with weariness
To th' dulling of my spirits. Sit down and rest.
Even here I will put off my hope and keep it
No longer for my flatterer. He is drowned
Whom thus we stray to find, and the sea mocks
Our frustrate search on land. Well, let him go.

ANTONIO, [aside to Sebastian]
I am right glad that he's so out of hope.
Do not, for one repulse, forgo the purpose
That you resolved t' effect.

SEBASTIAN, [aside to Antonio]  The next advantage
Will we take throughly.

ANTONIO, [aside to Sebastian]  Let it be tonight;
For now they are oppressed with travel, they
Will not nor cannot use such vigilance
As when they are fresh.

SEBASTIAN, [aside to Antonio]  I say tonight. No more.

[Solemn and strange music, and enter Prospero on the
top invisible.]


ALONSO
What harmony is this? My good friends, hark.

GONZALO  Marvelous sweet music!

[Enter several strange shapes, bringing in a banquet, and
dance about it with gentle actions of salutations.]


ALONSO
Give us kind keepers, heavens! What were these?

SEBASTIAN
A living drollery! Now I will believe
That there are unicorns, that in Arabia
There is one tree, the phoenix' throne, one phoenix
At this hour reigning there.

ANTONIO  I'll believe both;
And what does else want credit, come to me
And I'll be sworn 'tis true. Travelers ne'er did lie,
Though fools at home condemn 'em.

GONZALO  If in Naples
I should report this now, would they believe me?
If I should say I saw such islanders--
For, certes, these are people of the island--
Who, though they are of monstrous shape, yet note
Their manners are more gentle, kind, than of
Our human generation you shall find
Many, nay, almost any.

PROSPERO, [aside]  Honest lord,
Thou hast said well, for some of you there present
Are worse than devils.

ALONSO  I cannot too much muse
Such shapes, such gesture, and such sound,
expressing--
Although they want the use of tongue--a kind
Of excellent dumb discourse.

PROSPERO, [aside]  Praise in departing.
[Inviting the King, etc., to eat, the shapes depart.]

FRANCISCO  They vanished strangely.

SEBASTIAN  No matter, since
They have left their viands behind, for we have
stomachs.
Will 't please you taste of what is here?

ALONSO  Not I.

GONZALO
Faith, sir, you need not fear. When we were boys,
Who would believe that there were mountaineers
Dewlapped like bulls, whose throats had hanging at
'em
Wallets of flesh? Or that there were such men
Whose heads stood in their breasts? Which now we
find
Each putter-out of five for one will bring us
Good warrant of.

ALONSO  I will stand to and feed.
Although my last, no matter, since I feel
The best is past. Brother, my lord the Duke,
Stand to, and do as we.
[Alonso, Sebastian, and Antonio
move toward the table.]

[Thunder and lightning. Enter Ariel, like a Harpy, claps
his wings upon the table, and with a quaint device the
banquet vanishes.]


ARIEL [as Harpy]
You are three men of sin, whom Destiny,
That hath to instrument this lower world
And what is in 't, the never-surfeited sea
Hath caused to belch up you, and on this island,
Where man doth not inhabit, you 'mongst men
Being most unfit to live. I have made you mad;
And even with such-like valor, men hang and drown
Their proper selves.
[Alonso, Sebastian, and Antonio draw their swords.]
You fools, I and my fellows
Are ministers of Fate. The elements
Of whom your swords are tempered may as well
Wound the loud winds or with bemocked-at stabs
Kill the still-closing waters as diminish
One dowl that's in my plume. My fellow ministers
Are like invulnerable. If you could hurt,
Your swords are now too massy for your strengths
And will not be uplifted. But remember--
For that's my business to you--that you three
From Milan did supplant good Prospero,
Exposed unto the sea, which hath requit it,
Him and his innocent child, for which foul deed,
The powers--delaying, not forgetting--have
Incensed the seas and shores, yea, all the creatures
Against your peace. Thee of thy son, Alonso,
They have bereft; and do pronounce by me
Ling'ring perdition, worse than any death
Can be at once, shall step by step attend
You and your ways, whose wraths to guard you
from--
Which here, in this most desolate isle, else falls
Upon your heads--is nothing but heart's sorrow
And a clear life ensuing.	[He vanishes in thunder.]

[Then, to soft music, enter the shapes again, and dance,
with mocks and mows, and carrying out the table.]



PROSPERO, [aside]
Bravely the figure of this Harpy hast thou
Performed, my Ariel. A grace it had, devouring.
Of my instruction hast thou nothing bated
In what thou hadst to say. So, with good life
And observation strange, my meaner ministers
Their several kinds have done. My high charms
work,
And these mine enemies are all knit up
In their distractions. They now are in my power;
And in these fits I leave them while I visit
Young Ferdinand, whom they suppose is drowned,
And his and mine loved darling.	[He exits, above.]

GONZALO, [to Alonso]
I' th' name of something holy, sir, why stand you
In this strange stare?

ALONSO  O, it is monstrous, monstrous!
Methought the billows spoke and told me of it;
The winds did sing it to me, and the thunder,
That deep and dreadful organ pipe, pronounced
The name of Prosper. It did bass my trespass.
Therefor my son i' th' ooze is bedded, and
I'll seek him deeper than e'er plummet sounded,
And with him there lie mudded.	[He exits.]

SEBASTIAN  But one fiend at a time,
I'll fight their legions o'er.

ANTONIO  I'll be thy second.
[They exit.]

GONZALO
All three of them are desperate. Their great guilt,
Like poison given to work a great time after,
Now 'gins to bite the spirits. I do beseech you
That are of suppler joints, follow them swiftly
And hinder them from what this ecstasy
May now provoke them to.

ADRIAN  Follow, I pray you.
[They all exit.]


ACT 4
=====

Scene 1
=======
[Enter Prospero, Ferdinand, and Miranda.]


PROSPERO, [to Ferdinand]
If I have too austerely punished you,
Your compensation makes amends, for I
Have given you here a third of mine own life,
Or that for which I live; who once again
I tender to thy hand. All thy vexations
Were but my trials of thy love, and thou
Hast strangely stood the test. Here afore heaven
I ratify this my rich gift. O Ferdinand,
Do not smile at me that I boast of her,
For thou shalt find she will outstrip all praise
And make it halt behind her.

FERDINAND  I do believe it
Against an oracle.

PROSPERO
Then, as my gift and thine own acquisition
Worthily purchased, take my daughter. But
If thou dost break her virgin-knot before
All sanctimonious ceremonies may
With full and holy rite be ministered,
No sweet aspersion shall the heavens let fall
To make this contract grow; but barren hate,
Sour-eyed disdain, and discord shall bestrew
The union of your bed with weeds so loathly
That you shall hate it both. Therefore take heed,
As Hymen's lamps shall light you.

FERDINAND  As I hope
For quiet days, fair issue, and long life,
With such love as 'tis now, the murkiest den,
The most opportune place, the strong'st suggestion
Our worser genius can shall never melt
Mine honor into lust to take away
The edge of that day's celebration
When I shall think or Phoebus' steeds are foundered
Or night kept chained below.

PROSPERO  Fairly spoke.
Sit then and talk with her. She is thine own.
[Ferdinand and Miranda move aside.]
What, Ariel, my industrious servant, Ariel!

[Enter Ariel.]


ARIEL
What would my potent master? Here I am.

PROSPERO
Thou and thy meaner fellows your last service
Did worthily perform, and I must use you
In such another trick. Go bring the rabble,
O'er whom I give thee power, here to this place.
Incite them to quick motion, for I must
Bestow upon the eyes of this young couple
Some vanity of mine art. It is my promise,
And they expect it from me.

ARIEL  Presently?

PROSPERO  Ay, with a twink.

ARIEL
	Before you can say "Come" and "Go,"
	And breathe twice, and cry "So, so,"
	Each one, tripping on his toe,
	Will be here with mop and mow.
	Do you love me, master? No?

PROSPERO
Dearly, my delicate Ariel. Do not approach
Till thou dost hear me call.

ARIEL  Well; I conceive.
[He exits.]

PROSPERO, [to Ferdinand]
Look thou be true; do not give dalliance
Too much the rein. The strongest oaths are straw
To th' fire i' th' blood. Be more abstemious,
Or else goodnight your vow.

FERDINAND  I warrant you, sir,
The white cold virgin snow upon my heart
Abates the ardor of my liver.

PROSPERO  Well.--
Now come, my Ariel. Bring a corollary
Rather than want a spirit. Appear, and pertly.
[Soft music.]
No tongue. All eyes. Be silent.

[Enter Iris.]


IRIS
Ceres, most bounteous lady, thy rich leas
Of wheat, rye, barley, vetches, oats, and peas;
Thy turfy mountains, where live nibbling sheep,
And flat meads thatched with stover, them to keep;
Thy banks with pioned and twilled brims,
Which spongy April at thy hest betrims
To make cold nymphs chaste crowns; and thy
broom groves,
Whose shadow the dismissed bachelor loves,
Being lass-lorn; thy poll-clipped vineyard,
And thy sea marge, sterile and rocky hard,
Where thou thyself dost air--the Queen o' th' sky,
Whose wat'ry arch and messenger am I,
Bids thee leave these, and with her sovereign grace,
Here on this grass-plot, in this very place,
To come and sport. Her peacocks fly amain.
Approach, rich Ceres, her to entertain.

[Enter Ceres.]


CERES
Hail, many-colored messenger, that ne'er
Dost disobey the wife of Jupiter;
Who with thy saffron wings upon my flowers
Diffusest honey drops, refreshing showers;
And with each end of thy blue bow dost crown
My bosky acres and my unshrubbed down,
Rich scarf to my proud Earth. Why hath thy queen
Summoned me hither to this short-grassed green?

IRIS
A contract of true love to celebrate,
And some donation freely to estate
On the blest lovers.

CERES  Tell me, heavenly bow,
If Venus or her son, as thou dost know,
Do now attend the Queen? Since they did plot
The means that dusky Dis my daughter got,
Her and her blind boy's scandaled company
I have forsworn.

IRIS  Of her society
Be not afraid. I met her deity
Cutting the clouds towards Paphos, and her son
Dove-drawn with her. Here thought they to have
done
Some wanton charm upon this man and maid,
Whose vows are that no bed-right shall be paid
Till Hymen's torch be lighted--but in vain.
Mars's hot minion is returned again;
Her waspish-headed son has broke his arrows,
Swears he will shoot no more, but play with
sparrows,
And be a boy right out.

[Juno descends.]


CERES  Highest queen of state,
Great Juno, comes. I know her by her gait.

JUNO
How does my bounteous sister? Go with me
To bless this twain, that they may prosperous be
And honored in their issue.
[They sing.]

JUNO
	Honor, riches, marriage-blessing,
	Long continuance and increasing,
	Hourly joys be still upon you.
	Juno sings her blessings on you.

CERES
	Earth's increase, foison plenty,
	Barns and garners never empty,
	Vines with clust'ring bunches growing,
	Plants with goodly burden bowing;
	Spring come to you at the farthest
	In the very end of harvest.
	Scarcity and want shall shun you.
	Ceres' blessing so is on you.

FERDINAND
This is a most majestic vision, and
Harmonious charmingly. May I be bold
To think these spirits?

PROSPERO  Spirits, which by mine art
I have from their confines called to enact
My present fancies.

FERDINAND  Let me live here ever.
So rare a wondered father and a wise
Makes this place paradise.
[Juno and Ceres whisper,
and send Iris on employment.]

PROSPERO  Sweet now, silence.
Juno and Ceres whisper seriously.
There's something else to do. Hush, and be mute,
Or else our spell is marred.

IRIS
You nymphs, called naiads of the windring brooks,
With your sedged crowns and ever-harmless looks,
Leave your crisp channels and on this green land
Answer your summons, Juno does command.
Come, temperate nymphs, and help to celebrate
A contract of true love. Be not too late.

[Enter certain Nymphs.]

You sunburned sicklemen, of August weary,
Come hither from the furrow and be merry.
Make holiday: your rye-straw hats put on,
And these fresh nymphs encounter every one
In country footing.

[Enter certain Reapers, properly habited. They join with
the Nymphs in a graceful dance, towards the end
whereof Prospero starts suddenly and speaks.]


PROSPERO
I had forgot that foul conspiracy
Of the beast Caliban and his confederates
Against my life. The minute of their plot
Is almost come.--Well done. Avoid. No more.
[To a strange, hollow, and confused noise,
the spirits heavily vanish.]

FERDINAND, [to Miranda]
This is strange. Your father's in some passion
That works him strongly.

MIRANDA  Never till this day
Saw I him touched with anger, so distempered.

PROSPERO, [to Ferdinand]
You do look, my son, in a moved sort,
As if you were dismayed. Be cheerful, sir.
Our revels now are ended. These our actors,
As I foretold you, were all spirits and
Are melted into air, into thin air;
And like the baseless fabric of this vision,
The cloud-capped towers, the gorgeous palaces,
The solemn temples, the great globe itself,
Yea, all which it inherit, shall dissolve,
And, like this insubstantial pageant faded,
Leave not a rack behind. We are such stuff
As dreams are made on, and our little life
Is rounded with a sleep. Sir, I am vexed.
Bear with my weakness. My old brain is troubled.
Be not disturbed with my infirmity.
If you be pleased, retire into my cell
And there repose. A turn or two I'll walk
To still my beating mind.

FERDINAND/MIRANDA  We wish your peace.
[They exit.]

[Enter Ariel.]


PROSPERO
Come with a thought. I thank thee, Ariel. Come.

ARIEL
Thy thoughts I cleave to. What's thy pleasure?

PROSPERO  Spirit,
We must prepare to meet with Caliban.

ARIEL
Ay, my commander. When I presented Ceres,
I thought to have told thee of it, but I feared
Lest I might anger thee.

PROSPERO
Say again, where didst thou leave these varlets?

ARIEL
I told you, sir, they were red-hot with drinking,
So full of valor that they smote the air
For breathing in their faces, beat the ground
For kissing of their feet; yet always bending
Towards their project. Then I beat my tabor,
At which, like unbacked colts, they pricked their
ears,
Advanced their eyelids, lifted up their noses
As they smelt music. So I charmed their ears
That, calf-like, they my lowing followed through
Toothed briers, sharp furzes, pricking gorse, and
thorns,
Which entered their frail shins. At last I left them
I' th' filthy-mantled pool beyond your cell,
There dancing up to th' chins, that the foul lake
O'erstunk their feet.

PROSPERO  This was well done, my bird.
Thy shape invisible retain thou still.
The trumpery in my house, go bring it hither
For stale to catch these thieves.

ARIEL  I go, I go.	[He exits.]

PROSPERO
A devil, a born devil, on whose nature
Nurture can never stick; on whom my pains,
Humanely taken, all, all lost, quite lost;
And as with age his body uglier grows,
So his mind cankers. I will plague them all
Even to roaring.

[Enter Ariel, loaden with glistering apparel, etc.]

Come, hang them on this line.

[Enter Caliban, Stephano, and Trinculo, all wet, as
Prospero and Ariel look on.]


CALIBAN  Pray you, tread softly, that the blind mole
may not hear a footfall. We now are near his cell.

STEPHANO  Monster, your fairy, which you say is a
harmless fairy, has done little better than played the
jack with us.

TRINCULO  Monster, I do smell all horse piss, at which
my nose is in great indignation.

STEPHANO  So is mine.--Do you hear, monster. If I
should take a displeasure against you, look you--

TRINCULO  Thou wert but a lost monster.

CALIBAN
Good my lord, give me thy favor still.
Be patient, for the prize I'll bring thee to
Shall hoodwink this mischance. Therefore speak
softly.
All's hushed as midnight yet.

TRINCULO  Ay, but to lose our bottles in the pool!

STEPHANO  There is not only disgrace and dishonor in
that, monster, but an infinite loss.

TRINCULO  That's more to me than my wetting. Yet this
is your harmless fairy, monster!

STEPHANO  I will fetch off my bottle, though I be o'er
ears for my labor.

CALIBAN
Prithee, my king, be quiet. Seest thou here,
This is the mouth o' th' cell. No noise, and enter.
Do that good mischief which may make this island
Thine own forever, and I, thy Caliban,
For aye thy foot-licker.

STEPHANO  Give me thy hand. I do begin to have bloody
thoughts.

TRINCULO, [seeing the apparel]  O King Stephano, O
peer, O worthy Stephano, look what a wardrobe
here is for thee!

CALIBAN
Let it alone, thou fool. It is but trash.

TRINCULO  Oho, monster, we know what belongs to a
frippery. [He puts on one of the gowns.] O King
Stephano!

STEPHANO  Put off that gown, Trinculo. By this hand,
I'll have that gown.

TRINCULO  Thy Grace shall have it.

CALIBAN
The dropsy drown this fool! What do you mean
To dote thus on such luggage? Let 't alone,
And do the murder first. If he awake,
From toe to crown he'll fill our skins with pinches,
Make us strange stuff.

STEPHANO  Be you quiet, monster.--Mistress Line, is
not this my jerkin?	[He takes a jacket from the tree.]
Now is the jerkin under the line.--Now, jerkin, you
are like to lose your hair and prove a bald jerkin.

TRINCULO  Do, do. We steal by line and level, an 't like
your Grace.

STEPHANO  I thank thee for that jest. Here's a garment
for 't. Wit shall not go unrewarded while I am king
of this country. "Steal by line and level" is an excellent
pass of pate. There's another garment for 't.

TRINCULO  Monster, come, put some lime upon your
fingers, and away with the rest.

CALIBAN
I will have none on 't. We shall lose our time
And all be turned to barnacles or to apes
With foreheads villainous low.

STEPHANO  Monster, lay to your fingers. Help to bear
this away where my hogshead of wine is, or I'll turn
you out of my kingdom. Go to, carry this.

TRINCULO  And this.

STEPHANO  Ay, and this.
[A noise of hunters heard.]

[Enter divers spirits in shape of dogs and hounds,
hunting them about, Prospero and Ariel setting them on.]


PROSPERO  Hey, Mountain, hey!

ARIEL  Silver! There it goes, Silver!

PROSPERO
Fury, Fury! There, Tyrant, there! Hark, hark!
[Caliban, Stephano, and Trinculo are driven off.]
Go, charge my goblins that they grind their joints
With dry convulsions, shorten up their sinews
With aged cramps, and more pinch-spotted make
them
Than pard or cat o' mountain.

ARIEL  Hark, they roar.

PROSPERO
Let them be hunted soundly. At this hour
Lies at my mercy all mine enemies.
Shortly shall all my labors end, and thou
Shalt have the air at freedom. For a little
Follow and do me service.
[They exit.]



ACT 5
=====

Scene 1
=======
[Enter Prospero in his magic robes, and Ariel.]


PROSPERO
Now does my project gather to a head.
My charms crack not, my spirits obey, and time
Goes upright with his carriage.--How's the day?

ARIEL
On the sixth hour, at which time, my lord,
You said our work should cease.

PROSPERO  I did say so
When first I raised the tempest. Say, my spirit,
How fares the King and 's followers?

ARIEL  Confined
together
In the same fashion as you gave in charge,
Just as you left them; all prisoners, sir,
In the line grove which weather-fends your cell.
They cannot budge till your release. The King,
His brother, and yours abide all three distracted,
And the remainder mourning over them,
Brimful of sorrow and dismay; but chiefly
Him that you termed, sir, the good old Lord
Gonzalo.
His tears runs down his beard like winter's drops
From eaves of reeds. Your charm so strongly works
'em
That if you now beheld them, your affections
Would become tender.

PROSPERO  Dost thou think so, spirit?

ARIEL
Mine would, sir, were I human.

PROSPERO  And mine shall.
Hast thou, which art but air, a touch, a feeling
Of their afflictions, and shall not myself,
One of their kind, that relish all as sharply
Passion as they, be kindlier moved than thou art?
Though with their high wrongs I am struck to th'
quick,
Yet with my nobler reason 'gainst my fury
Do I take part. The rarer action is
In virtue than in vengeance. They being penitent,
The sole drift of my purpose doth extend
Not a frown further. Go, release them, Ariel.
My charms I'll break, their senses I'll restore,
And they shall be themselves.

ARIEL  I'll fetch them, sir.
[He exits.]

[Prospero draws a large circle on the stage with his staff.]


PROSPERO
You elves of hills, brooks, standing lakes, and groves,
And you that on the sands with printless foot
Do chase the ebbing Neptune, and do fly him
When he comes back; you demi-puppets that
By moonshine do the green sour ringlets make,
Whereof the ewe not bites; and you whose pastime
Is to make midnight mushrumps, that rejoice
To hear the solemn curfew; by whose aid,
Weak masters though you be, I have bedimmed
The noontide sun, called forth the mutinous winds,
And 'twixt the green sea and the azured vault
Set roaring war; to the dread rattling thunder
Have I given fire, and rifted Jove's stout oak
With his own bolt; the strong-based promontory
Have I made shake, and by the spurs plucked up
The pine and cedar; graves at my command
Have waked their sleepers, oped, and let 'em forth
By my so potent art. But this rough magic
I here abjure, and when I have required
Some heavenly music, which even now I do,
[Prospero gestures with his staff.]
To work mine end upon their senses that
This airy charm is for, I'll break my staff,
Bury it certain fathoms in the earth,
And deeper than did ever plummet sound
I'll drown my book.	[Solemn music.]

[Here enters Ariel before; then Alonso with a frantic
gesture, attended by Gonzalo; Sebastian and Antonio in
like manner attended by Adrian and Francisco. They all
enter the circle which Prospero had made, and there
stand charmed; which Prospero observing, speaks.]

A solemn air, and the best comforter
To an unsettled fancy, cure thy brains,
Now useless, boiled within thy skull. There stand,
For you are spell-stopped.--
Holy Gonzalo, honorable man,
Mine eyes, e'en sociable to the show of thine,
Fall fellowly drops.--The charm dissolves apace,
And as the morning steals upon the night,
Melting the darkness, so their rising senses
Begin to chase the ignorant fumes that mantle
Their clearer reason.--O good Gonzalo,
My true preserver and a loyal sir
To him thou follow'st, I will pay thy graces
Home, both in word and deed.--Most cruelly
Didst thou, Alonso, use me and my daughter.
Thy brother was a furtherer in the act.--
Thou art pinched for 't now, Sebastian.--Flesh and
blood,
You, brother mine, that entertained ambition,
Expelled remorse and nature, whom, with Sebastian,
Whose inward pinches therefore are most strong,
Would here have killed your king, I do forgive thee,
Unnatural though thou art.--Their understanding
Begins to swell, and the approaching tide
Will shortly fill the reasonable shore
That now lies foul and muddy. Not one of them
That yet looks on me or would know me.--Ariel,
Fetch me the hat and rapier in my cell.
[Ariel exits and at once returns
with Prospero's ducal robes.]
I will discase me and myself present
As I was sometime Milan.--Quickly, spirit,
Thou shalt ere long be free.

ARIEL [sings, and helps to attire him.]

	   Where the bee sucks, there suck I.
	   In a cowslip's bell I lie.
	   There I couch when owls do cry.
	   On the bat's back I do fly
	   After summer merrily.
	Merrily, merrily shall I live now
	Under the blossom that hangs on the bow.


PROSPERO
Why, that's my dainty Ariel. I shall miss
Thee, but yet thou shalt have freedom. So, so, so.
To the King's ship, invisible as thou art.
There shalt thou find the mariners asleep
Under the hatches. The master and the boatswain
Being awake, enforce them to this place,
And presently, I prithee.

ARIEL
I drink the air before me, and return
Or ere your pulse twice beat.	[He exits.]

GONZALO
All torment, trouble, wonder, and amazement
Inhabits here. Some heavenly power guide us
Out of this fearful country!

PROSPERO, [to Alonso]  Behold, sir king,
The wronged Duke of Milan, Prospero.
For more assurance that a living prince
Does now speak to thee, I embrace thy body,
[He embraces Alonso.]
And to thee and thy company I bid
A hearty welcome.

ALONSO  Whe'er thou be'st he or no,
Or some enchanted trifle to abuse me
(As late I have been) I not know. Thy pulse
Beats as of flesh and blood; and since I saw thee,
Th' affliction of my mind amends, with which
I fear a madness held me. This must crave,
An if this be at all, a most strange story.
Thy dukedom I resign, and do entreat
Thou pardon me my wrongs. But how should
Prospero
Be living and be here?

PROSPERO, [to Gonzalo]  First, noble friend,
Let me embrace thine age, whose honor cannot
Be measured or confined.

GONZALO  Whether this be
Or be not, I'll not swear.

PROSPERO  You do yet taste
Some subtleties o' th' isle, that will not let you
Believe things certain. Welcome, my friends all.
[Aside to Sebastian and Antonio.] But you, my brace
of lords, were I so minded,
I here could pluck his Highness' frown upon you
And justify you traitors. At this time
I will tell no tales.

SEBASTIAN, [aside]  The devil speaks in him.

PROSPERO, [aside to Sebastian]  No.
[To Antonio.] For you, most wicked sir, whom to
call brother
Would even infect my mouth, I do forgive
Thy rankest fault, all of them, and require
My dukedom of thee, which perforce I know
Thou must restore.

ALONSO  If thou be'st Prospero,
Give us particulars of thy preservation,
How thou hast met us here, whom three hours since
Were wracked upon this shore, where I have lost--
How sharp the point of this remembrance is!--
My dear son Ferdinand.

PROSPERO  I am woe for 't, sir.

ALONSO
Irreparable is the loss, and patience
Says it is past her cure.

PROSPERO  I rather think
You have not sought her help, of whose soft grace,
For the like loss, I have her sovereign aid
And rest myself content.

ALONSO  You the like loss?

PROSPERO
As great to me as late, and supportable
To make the dear loss have I means much weaker
Than you may call to comfort you, for I
Have lost my daughter.

ALONSO  A daughter?
O heavens, that they were living both in Naples,
The King and Queen there! That they were, I wish
Myself were mudded in that oozy bed
Where my son lies!--When did you lose your
daughter?

PROSPERO
In this last tempest. I perceive these lords
At this encounter do so much admire
That they devour their reason, and scarce think
Their eyes do offices of truth, their words
Are natural breath.--But howsoe'er you have
Been justled from your senses, know for certain
That I am Prospero and that very duke
Which was thrust forth of Milan, who most
strangely
Upon this shore, where you were wracked, was
landed
To be the lord on 't. No more yet of this.
For 'tis a chronicle of day by day,
Not a relation for a breakfast, nor
Befitting this first meeting. [To Alonso.] Welcome, sir.
This cell's my court. Here have I few attendants,
And subjects none abroad. Pray you, look in.
My dukedom since you have given me again,
I will requite you with as good a thing,
At least bring forth a wonder to content you
As much as me my dukedom.
[Here Prospero discovers Ferdinand and Miranda,
playing at chess.]

MIRANDA, [to Ferdinand]
Sweet lord, you play me false.

FERDINAND  No, my dearest love,
I would not for the world.

MIRANDA
Yes, for a score of kingdoms you should wrangle,
And I would call it fair play.

ALONSO  If this prove
A vision of the island, one dear son
Shall I twice lose.

SEBASTIAN  A most high miracle!

FERDINAND, [seeing Alonso and coming forward]
Though the seas threaten, they are merciful.
I have cursed them without cause.	[He kneels.]

ALONSO  Now, all the
blessings
Of a glad father compass thee about!
Arise, and say how thou cam'st here.
[Ferdinand stands.]

MIRANDA, [rising and coming forward]  O wonder!
How many goodly creatures are there here!
How beauteous mankind is! O, brave new world
That has such people in 't!

PROSPERO  'Tis new to thee.

ALONSO, [to Ferdinand]
What is this maid with whom thou wast at play?
Your eld'st acquaintance cannot be three hours.
Is she the goddess that hath severed us
And brought us thus together?

FERDINAND  Sir, she is mortal,
But by immortal providence she's mine.
I chose her when I could not ask my father
For his advice, nor thought I had one. She
Is daughter to this famous Duke of Milan,
Of whom so often I have heard renown,
But never saw before, of whom I have
Received a second life; and second father
This lady makes him to me.

ALONSO  I am hers.
But, O, how oddly will it sound that I
Must ask my child forgiveness!

PROSPERO  There, sir, stop.
Let us not burden our remembrances with
A heaviness that's gone.

GONZALO  I have inly wept
Or should have spoke ere this. Look down, you
gods,
And on this couple drop a blessed crown,
For it is you that have chalked forth the way
Which brought us hither.

ALONSO  I say "Amen," Gonzalo.

GONZALO
Was Milan thrust from Milan, that his issue
Should become kings of Naples? O, rejoice
Beyond a common joy, and set it down
With gold on lasting pillars: in one voyage
Did Claribel her husband find at Tunis,
And Ferdinand, her brother, found a wife
Where he himself was lost; Prospero his dukedom
In a poor isle; and all of us ourselves
When no man was his own.

ALONSO, [to Ferdinand and Miranda]  Give me your
hands.
Let grief and sorrow still embrace his heart
That doth not wish you joy!

GONZALO  Be it so. Amen.

[Enter Ariel, with the Master and Boatswain
amazedly following.]

O, look, sir, look, sir, here is more of us.
I prophesied if a gallows were on land,
This fellow could not drown. Now, blasphemy,
That swear'st grace o'erboard, not an oath on
shore?
Hast thou no mouth by land? What is the news?

BOATSWAIN
The best news is that we have safely found
Our king and company. The next: our ship,
Which, but three glasses since, we gave out split,
Is tight and yare and bravely rigged as when
We first put out to sea.

ARIEL, [aside to Prospero]  Sir, all this service
Have I done since I went.

PROSPERO, [aside to Ariel]  My tricksy spirit!

ALONSO
These are not natural events. They strengthen
From strange to stranger.--Say, how came you
hither?

BOATSWAIN
If I did think, sir, I were well awake,
I'd strive to tell you. We were dead of sleep
And--how, we know not--all clapped under
hatches,
Where, but even now, with strange and several
noises
Of roaring, shrieking, howling, jingling chains,
And more diversity of sounds, all horrible,
We were awaked, straightway at liberty,
Where we, in all her trim, freshly beheld
Our royal, good, and gallant ship, our master
Cap'ring to eye her. On a trice, so please you,
Even in a dream were we divided from them
And were brought moping hither.

ARIEL, [aside to Prospero]  Was 't well done?

PROSPERO, [aside to Ariel]
Bravely, my diligence. Thou shalt be free.

ALONSO
This is as strange a maze as e'er men trod,
And there is in this business more than nature
Was ever conduct of. Some oracle
Must rectify our knowledge.

PROSPERO  Sir, my liege,
Do not infest your mind with beating on
The strangeness of this business. At picked leisure,
Which shall be shortly, single I'll resolve you,
Which to you shall seem probable, of every
These happened accidents; till when, be cheerful
And think of each thing well. [Aside to Ariel.]
Come hither, spirit;
Set Caliban and his companions free.
Untie the spell. [Ariel exits.] How fares my gracious
sir?
There are yet missing of your company
Some few odd lads that you remember not.

[Enter Ariel, driving in Caliban, Stephano, and Trinculo
in their stolen apparel.]


STEPHANO  Every man shift for all the rest, and let no
man take care for himself, for all is but fortune.
Coraggio, bully monster, coraggio.

TRINCULO  If these be true spies which I wear in my
head, here's a goodly sight.

CALIBAN  O Setebos, these be brave spirits indeed! How
fine my master is! I am afraid he will chastise me.

SEBASTIAN  Ha, ha!
What things are these, my Lord Antonio?
Will money buy 'em?

ANTONIO  Very like. One of them
Is a plain fish and no doubt marketable.

PROSPERO
Mark but the badges of these men, my lords,
Then say if they be true. This misshapen knave,
His mother was a witch, and one so strong
That could control the moon, make flows and ebbs,
And deal in her command without her power.
These three have robbed me, and this demi-devil,
For he's a bastard one, had plotted with them
To take my life. Two of these fellows you
Must know and own. This thing of darkness I
Acknowledge mine.

CALIBAN  I shall be pinched to death.

ALONSO
Is not this Stephano, my drunken butler?

SEBASTIAN  He is drunk now. Where had he wine?

ALONSO
And Trinculo is reeling ripe. Where should they
Find this grand liquor that hath gilded 'em?
[To Trinculo.] How cam'st thou in this pickle?

TRINCULO  I have been in such a pickle since I saw you
last that I fear me will never out of my bones. I
shall not fear flyblowing.

SEBASTIAN  Why, how now, Stephano?

STEPHANO  O, touch me not! I am not Stephano, but a
cramp.

PROSPERO  You'd be king o' the isle, sirrah?

STEPHANO  I should have been a sore one, then.

ALONSO, [indicating Caliban]
This is as strange a thing as e'er I looked on.

PROSPERO
He is as disproportioned in his manners
As in his shape. [To Caliban.] Go, sirrah, to my cell.
Take with you your companions. As you look
To have my pardon, trim it handsomely.

CALIBAN
Ay, that I will, and I'll be wise hereafter
And seek for grace. What a thrice-double ass
Was I to take this drunkard for a god,
And worship this dull fool!

PROSPERO  Go to, away!

ALONSO, [to Stephano and Trinculo]
Hence, and bestow your luggage where you found it.

SEBASTIAN  Or stole it, rather.
[Caliban, Stephano, and Trinculo exit.]

PROSPERO
Sir, I invite your Highness and your train
To my poor cell, where you shall take your rest
For this one night, which part of it I'll waste
With such discourse as, I not doubt, shall make it
Go quick away: the story of my life
And the particular accidents gone by
Since I came to this isle. And in the morn
I'll bring you to your ship, and so to Naples,
Where I have hope to see the nuptial
Of these our dear-beloved solemnized,
And thence retire me to my Milan, where
Every third thought shall be my grave.

ALONSO  I long
To hear the story of your life, which must
Take the ear strangely.

PROSPERO  I'll deliver all,
And promise you calm seas, auspicious gales,
And sail so expeditious that shall catch
Your royal fleet far off. [Aside to Ariel.] My Ariel,
chick,
That is thy charge. Then to the elements
Be free, and fare thou well.--Please you, draw near.
[They all exit.]

EPILOGUE,
[spoken by Prospero.]
============================


Now my charms are all o'erthrown,
And what strength I have 's mine own,
Which is most faint. Now 'tis true
I must be here confined by you,
Or sent to Naples. Let me not,
Since I have my dukedom got
And pardoned the deceiver, dwell
In this bare island by your spell,
But release me from my bands
With the help of your good hands.
Gentle breath of yours my sails
Must fill, or else my project fails,
Which was to please. Now I want
Spirits to enforce, art to enchant,
And my ending is despair,
Unless I be relieved by prayer,
Which pierces so that it assaults
Mercy itself, and frees all faults.
   As you from crimes would pardoned be,
   Let your indulgence set me free.
[He exits.]
`
return b
}
