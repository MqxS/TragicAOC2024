package day1

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	input := "87501   76559\n70867   16862\n12959   38527\n56898   81917\n80416   13287\n28886   54457\n79252   30354\n47576   88490\n43354   37397\n89248   74846\n39921   24805\n98636   51185\n33277   31605\n45307   13417\n33326   72874\n14449   42023\n64412   40326\n12630   40326\n35665   41197\n35932   59560\n22757   76636\n97387   91997\n83599   74846\n33718   54077\n20879   65995\n42419   35638\n50241   41197\n94123   27231\n82872   65149\n41378   85282\n81233   65415\n98875   21219\n21517   81917\n36314   65845\n64212   43331\n94404   34854\n42166   87444\n13351   12627\n53796   47507\n19837   28551\n59598   71749\n47765   93643\n11282   91997\n71285   69206\n27075   56104\n11470   50196\n75795   17345\n77811   85578\n56347   74690\n54911   35921\n26533   96584\n75314   58859\n49216   30077\n94855   14154\n10775   91997\n58190   81917\n38228   85154\n88321   21470\n99407   38527\n39166   13647\n22369   96563\n61678   29486\n94911   64616\n99565   66640\n64630   86818\n60973   22185\n83684   27341\n44345   22530\n43964   11793\n13207   62913\n20848   30354\n43944   38527\n48992   38527\n44659   10142\n93179   96119\n80123   86215\n46758   75732\n34750   18011\n38136   92652\n25072   58141\n99637   96563\n42591   32117\n48968   60830\n68846   15876\n63257   19695\n20217   69184\n50862   86772\n35900   31605\n61185   10607\n41487   98125\n29962   81917\n44777   73031\n40743   25415\n57518   13407\n89369   36534\n36269   31605\n59657   27902\n72361   51185\n71609   15589\n76578   19417\n35670   18977\n99141   95960\n30841   55313\n91034   30354\n70937   81509\n65910   12959\n62321   74690\n39367   79276\n39883   23467\n15590   54077\n70456   26094\n21706   97640\n91933   26097\n88390   65845\n24955   91997\n87059   71437\n33177   86262\n57578   84284\n24831   21219\n19693   74846\n87500   20041\n88393   50151\n18468   69184\n61548   69184\n75166   74739\n80975   52736\n54909   58854\n85260   61330\n86203   15355\n75868   81154\n57081   81917\n44633   41212\n53395   92203\n38922   84627\n45580   22005\n11492   12959\n13813   23052\n77790   31605\n97416   41197\n60784   40326\n91432   12377\n64293   74468\n53735   54077\n11377   23052\n35848   57154\n29067   57838\n14215   61664\n72068   30354\n23691   58540\n17111   28991\n66651   12959\n35155   25057\n76500   74690\n59018   78562\n51786   64460\n12770   80286\n41212   53868\n85299   46294\n58892   21646\n46842   31347\n80228   82547\n21816   41197\n63096   69071\n55837   74130\n13699   99276\n88678   23052\n17917   65679\n42692   28551\n42536   90759\n29227   74130\n44887   47070\n14869   96563\n62941   87555\n96430   52517\n92100   96563\n69859   65845\n78066   55780\n71937   80440\n79682   46524\n90428   37542\n17489   76606\n12443   48973\n92963   12959\n13845   39796\n27104   19295\n84047   92652\n17026   54077\n16587   14154\n38434   61664\n96992   78286\n60514   74130\n76740   31347\n77040   71639\n57707   69864\n18252   74823\n50362   92774\n87527   54077\n18024   62644\n16185   99146\n15836   31347\n71453   22056\n81519   20539\n92269   41197\n89510   98904\n64230   65845\n44965   69184\n97898   50122\n33394   71749\n51205   31347\n29970   91997\n18703   38527\n95881   85733\n78977   65679\n44185   42353\n59579   79102\n98081   88026\n75176   28551\n23052   25529\n23648   86507\n68646   65149\n81344   94453\n89002   23057\n91635   92652\n83079   54077\n19710   96563\n55515   31857\n94625   45828\n95178   40326\n43069   41940\n32573   48931\n28551   28551\n71918   31605\n64790   25534\n70603   80997\n98971   60973\n26316   45940\n56050   47638\n70043   86262\n89213   96563\n19535   86236\n75037   86262\n56394   30161\n10420   38501\n97325   34627\n79627   38019\n90943   64789\n89009   65149\n41714   74846\n56219   30354\n67414   33426\n70637   91997\n65116   68219\n65081   92652\n30062   28551\n53883   74846\n82201   35608\n30366   36832\n92536   34251\n41242   23208\n28183   95131\n95821   20974\n70104   53868\n92750   81866\n19452   31605\n63808   40326\n48841   74690\n78057   28551\n18349   23052\n49722   69184\n47153   92162\n87301   17834\n62420   60858\n77680   46524\n27690   17204\n59409   92652\n42268   28551\n63330   17663\n36144   77627\n31999   23179\n63097   44128\n32578   30185\n37261   39103\n75864   10428\n49682   64324\n19524   86615\n21013   57384\n28003   14804\n69912   23052\n56391   16437\n55578   63279\n66306   38527\n50485   97228\n17388   69184\n68640   51185\n34002   69184\n53182   74846\n26868   30764\n54480   21219\n78677   30346\n63487   49614\n96563   41212\n63947   86772\n95430   19044\n66914   70185\n17750   12959\n44339   31260\n91997   85956\n19858   81917\n39187   14154\n31250   11742\n81979   11075\n43865   93422\n41577   12959\n59486   26106\n84822   67408\n83663   28563\n53342   88705\n88483   31347\n26040   64385\n59222   22924\n28173   52034\n47265   12959\n98279   23813\n74690   53868\n50359   17936\n24495   97323\n67473   74690\n31605   86772\n47293   69495\n78146   54077\n62042   74130\n22969   81917\n86262   41212\n20100   14154\n31926   51869\n83586   20705\n78067   87533\n15120   80541\n60964   74514\n25983   86772\n22555   85287\n82691   41212\n33704   54077\n60332   65149\n88891   54077\n11798   54077\n33296   95893\n21560   71749\n72107   23992\n99423   16437\n74019   53868\n11793   31347\n24580   57047\n29058   41212\n13516   65325\n45127   65149\n78583   68013\n47357   26483\n74336   54077\n73563   89201\n48046   22056\n99277   40534\n51620   22865\n26126   50197\n39198   46524\n99101   38128\n53868   57894\n60036   79056\n98998   14506\n33596   49268\n96813   97654\n63461   51185\n44286   31605\n32282   65845\n40326   21219\n31739   99276\n86081   64405\n17565   11401\n98499   52794\n54371   51185\n64366   91997\n22558   44849\n80145   80748\n55344   30354\n82390   78198\n67477   52858\n41679   65845\n26759   61664\n25479   24285\n84759   94740\n23922   31734\n69409   32214\n16130   14154\n17075   41197\n26992   88528\n75638   14608\n40799   31367\n70389   65387\n48822   91997\n82741   65149\n53111   75593\n96714   91408\n61824   98553\n80721   68977\n14154   26094\n65521   21219\n62357   65149\n17248   33364\n62195   96563\n31347   69184\n21582   96563\n67200   15744\n32845   70938\n69303   16437\n43616   94424\n70647   32249\n37198   29513\n56883   86314\n55553   74130\n44144   29558\n93899   31605\n94230   96563\n92309   65679\n69517   91305\n30628   21219\n14181   15109\n86696   59832\n92793   79552\n26915   12959\n32013   99276\n81917   11773\n85533   91997\n10291   57895\n47306   86262\n40182   86772\n78380   96563\n14074   31347\n29716   51680\n34074   69184\n83682   79659\n32571   81917\n24215   65845\n14613   30354\n86190   46524\n22056   77073\n47602   77596\n40618   22388\n61777   69973\n83610   86262\n41123   74442\n14249   55123\n49493   94023\n42659   94250\n38055   92652\n49301   72817\n92652   96746\n96637   41212\n63286   53868\n64698   19412\n48939   65149\n77435   42058\n86082   97780\n56906   31605\n18370   69184\n55583   96944\n10163   99796\n57905   99276\n62348   28809\n84711   81543\n14792   31605\n27680   69184\n90559   55495\n56049   23052\n30776   23052\n25298   89829\n72656   86262\n74027   34083\n79517   31376\n70497   36800\n67662   65845\n71206   47116\n14457   17103\n79508   31347\n69184   24527\n18264   85310\n16462   81917\n48031   53868\n85157   82675\n12794   67521\n56636   69184\n63663   23052\n83754   12959\n90766   27743\n48559   15605\n95519   43790\n24578   14154\n13264   29964\n72984   53868\n23859   27144\n52266   14154\n74332   27923\n40908   37044\n45622   65679\n97941   41197\n29959   72533\n75077   76897\n43027   69795\n75223   54077\n46175   41212\n60194   78373\n63100   51133\n41845   19241\n33956   85483\n71994   81917\n32196   30354\n26319   55110\n52113   38527\n45378   17528\n29872   15275\n45386   92652\n50436   51185\n59375   77172\n29944   96563\n80342   40326\n10493   16437\n97587   64996\n59750   65664\n88477   53868\n95607   53868\n81650   60061\n89497   21219\n60828   69184\n18759   41212\n25409   86772\n42395   73365\n42537   71749\n23075   15595\n33176   92903\n61312   77585\n49549   16437\n13037   30354\n22694   27763\n55457   40372\n41517   74846\n62998   69981\n33002   81778\n19731   50866\n37666   72519\n44046   80344\n16308   34728\n39639   35278\n52166   58851\n18885   51185\n52717   16437\n54738   51185\n30354   23052\n71905   83743\n53902   53981\n49202   45117\n29053   53704\n73107   93596\n10814   81917\n47137   81917\n14681   46524\n23659   19157\n41523   19877\n37342   74712\n14677   35820\n80367   30354\n29514   68638\n92987   44546\n53999   13219\n53071   12959\n18116   30354\n41120   21219\n24265   13890\n74236   12238\n16153   40253\n29225   16518\n77153   69767\n11156   15109\n17676   72608\n97762   69184\n14294   74130\n20945   53868\n89054   69842\n50026   92308\n11773   35864\n43398   46524\n81325   76898\n72450   94665\n73264   14831\n11738   45600\n47816   33630\n96440   38527\n49621   26094\n99276   86262\n83302   66242\n44195   82631\n72375   19298\n62769   30354\n24836   41212\n54126   12959\n56172   53521\n40147   84315\n15109   48388\n36959   44227\n72676   50902\n29907   57407\n17101   59056\n27776   92652\n33908   38484\n31997   41212\n60558   60338\n76898   29376\n66530   12210\n55325   74690\n14813   81185\n65149   86262\n64181   42618\n13236   75101\n85280   25972\n23433   27380\n76693   28959\n27598   65845\n85215   16437\n69326   31605\n73527   54077\n83772   50003\n13383   90621\n74846   86772\n76968   51536\n67965   99132\n75669   71623\n11827   13380\n65958   99706\n27705   54077\n29880   83668\n64798   37851\n15552   20138\n81369   86262\n16035   13132\n19559   34135\n58504   41212\n53208   96563\n72798   44868\n63911   54077\n73341   86272\n57970   81361\n34600   40326\n32884   85531\n71749   49989\n16407   39922\n48023   47921\n55208   38527\n87016   88919\n61959   71284\n14403   20794\n61664   74846\n84201   12488\n73175   16437\n75814   65638\n91655   98512\n33948   23052\n80300   69184\n70886   76898\n41410   65149\n78044   16437\n32339   58862\n16859   18026\n10047   30354\n23013   31605\n37757   22056\n28383   39425\n88083   23844\n91191   21148\n21219   53868\n68359   54326\n36728   81111\n83279   55929\n30823   72588\n94846   60851\n45506   40326\n33240   21219\n65845   61445\n45314   95859\n94477   91997\n41623   91997\n18698   81917\n26025   96563\n67139   28719\n70864   20565\n63628   86772\n82479   12959\n56207   23052\n18780   68150\n61220   53868\n94501   47083\n58257   10169\n16437   38161\n89620   44175\n74771   46524\n99455   32934\n77273   91997\n20854   40324\n43825   74130\n89584   41212\n86772   87585\n16169   40326\n23079   91997\n20347   12771\n33162   53321\n77049   61664\n52391   11773\n98777   19105\n86182   77697\n84991   76064\n46587   14887\n13200   76667\n76967   65149\n56719   38542\n77987   41197\n74246   78717\n68857   23052\n92089   65845\n79803   44983\n16391   86083\n85511   82407\n14130   31605\n20211   65845\n51868   35534\n11881   17074\n43073   87808\n74013   95379\n63581   91997\n67718   75679\n95358   38527\n78186   12888\n35563   61664\n63140   92652\n11832   54077\n11698   58297\n22084   11865\n98044   64989\n26094   35964\n45393   92652\n38527   93080\n66819   46524\n31166   52386\n39568   79757\n94063   77614\n16452   83128\n24171   91977\n74980   12959\n47814   38527\n71716   45179\n34263   14154\n43631   53868\n47023   99276\n56699   27179\n13150   41212\n65739   46524\n87085   60735\n67283   21219\n23771   32301\n65647   55281\n50600   45366\n88186   54121\n48164   86262\n34469   21705\n76035   97022\n91266   69184\n55549   65149\n34832   38480\n72693   74690\n66050   65149\n50769   63440\n34690   91929\n58125   37769\n16861   65845\n26369   65845\n31834   74130\n24724   76898\n35677   46524\n79253   16437\n29593   31347\n29695   51028\n18196   96563\n33927   74846\n24378   16437\n37505   47616\n70048   31232\n27708   53750\n14952   92748\n37930   53868\n41197   28551\n77433   38527\n77320   54519\n37592   82995\n73963   65149\n67891   66894\n71199   71879\n13884   73982\n97681   33048\n44494   21083\n56856   62395\n56521   48031\n11077   40326\n98851   49244\n19484   46916\n51325   19011\n74821   65679\n90241   78808\n24897   60176\n40386   30354\n66435   10877\n17691   88972\n17430   96495\n76735   57514\n23936   46524\n86656   86003\n81533   46524\n23473   14374\n44831   17013\n73390   53868\n20456   38527\n47304   43209\n95484   26094\n40033   92652\n49629   65845\n11518   94552\n99642   28659\n46027   38527\n93133   99640\n88453   69184\n52465   87482\n65679   96563\n51133   86262\n82072   24993\n65019   81577\n73181   20846\n25682   74846\n89215   23052\n75313   28895\n59429   99506\n95353   98713\n14199   71449\n28894   21219\n34814   77228\n12462   81917\n77803   10891\n54077   25857\n56105   87001\n49225   91997\n33427   23052\n28144   53868\n44230   74256\n64121   38385\n44223   94238\n91417   28551\n27420   43632\n49221   41212\n37139   31605\n65039   36279\n40544   97794\n70621   62790\n89720   12494\n93652   70171\n99679   77510\n48132   65845\n27910   63372\n61906   54077\n79192   88515\n52071   15109\n68413   40326\n38094   65845\n85325   22914\n34519   60686\n18347   62726\n74130   86262\n22266   70119\n24902   79593\n58897   53868\n67166   94563\n86035   54859\n32151   51640\n19735   28551\n86835   65845\n48502   69184\n98335   65461\n85776   61664\n31060   41212\n51185   25984\n10538   61666\n89716   70800\n93969   46524\n87113   41212\n94835   43501\n31171   61231\n63678   83332\n27983   36115\n31882   66078\n69310   21182\n57748   69184\n34236   61664\n22019   53868\n40783   74204\n11443   65306\n61587   74130\n32107   20837\n42950   75914\n30318   28551\n71868   21219\n11717   29036\n15069   54142\n46368   23052\n19544   53868\n96045   10501\n70708   90926\n66286   14154\n93443   46524\n22988   52299\n59398   87705\n54189   79654\n44636   47669\n34569   12959\n75207   88665\n57213   81917\n17661   76365\n76744   41550\n18525   81729\n44095   89137\n47905   69184\n64998   70782\n94932   14154\n79764   85179\n94708   22056\n97002   11773\n65044   71749\n81636   90224\n67614   31605\n69858   12959\n37652   86650\n59098   35258\n26949   46653\n19239   21219\n46003   58379\n77640   55716\n92129   51133\n46524   60031\n11683   44416\n72664   86772\n80935   81071\n12683   86262\n47561   41319\n97326   32695\n91545   31605\n50527   99276\n78000   61404\n58880   14154\n47352   89525\n64943   54678\n34249   12959\n90703   52069\n42396   21732\n17077   72444\n40369   16956\n17581   39897\n18310   46524\n93377   20361\n21337   41798\n31532   65845\n63298   78248\n39119   46524\n53668   38527\n50927   96563\n57380   47367\n68451   16437\n96725   28551\n49778   86415\n46687   12305\n93651   84334\n17709   46524\n36657   38527\n67669   20743\n18254   55605\n99194   12959\n51889   29126\n11292   35259\n39589   74690\n14716   54077\n71832   67905\n54150   55806\n36451   46524\n14806   96563\n63718   74103\n58557   95632\n98671   95757\n48004   65149\n20103   21592"
	var leftList []int
	var rightList []int

	for _, line := range strings.Split(input, "\n") {
		nums := strings.Split(line, "   ")
		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, leftNum)
		rightList = append(rightList, rightNum)
	}

	sort.IntSlice(leftList).Sort()
	sort.IntSlice(rightList).Sort()

	totalDistance := 0.0
	for i := 0; i < len(leftList); i++ {
		distance := rightList[i] - leftList[i]
		totalDistance += math.Abs(float64(distance))
	}

	fmt.Printf("Part 1: %v\n", int(totalDistance))

	similarityScore := 0
	for _, leftNum := range leftList {
		presentInRight := 0
		for _, rightNum := range rightList {
			if leftNum == rightNum {
				presentInRight++
			}
		}
		similarityScore += leftNum * presentInRight
	}

	fmt.Printf("Part 2: %v\n", similarityScore)
}