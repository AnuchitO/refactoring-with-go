# Prints bill
บริษัททำโรงละคร รับแสดงนอกสถานที่ ลูกค้าสามารถเลือกได้ว่าจะอยากได้การแสดง(play)ประเภทใด
ทางบริษัทจะคิดเงินลูกค้าตามจำนวนผู้ชมและประเภทของการแสดง
ซึ่งตอนนี้ทางบริษัทมีนักแสดงสองแบบ การแสดงโศกนาฏกรรม (tragedies) และ การแสดงตลก (comedies)
ทางบริษัทเองยังต้องออกบิลสำหรับการแสดงแต่ละครั้งให้กับลูกค้า และลูกค้ายังได้แต้ม (volume credits)
สะสมไว้เพื่อใช้เป็นส่วนลดในครั้งต่อไปได้ (นึกถึงว่าอันนี้คือระบบสมาชิกสะสมแต้มประมาณนั้น)


ตอนนี้บริษัทเก็บข้อมูลไว้ในรูปแบบของ JSON

### ข้อมูลนักแสดง
plays.json
```json
{
   "hamlet":{"name":"Hamlet","type":"tragedy"},
   "as-like":{"name":"As You Like It","type":"comedy"},
   "othello":{"name":"Othello","type":"tragedy"}
}
```

### ข้อมูลที่ใช้ในการออกบิลให้ลูกค้า
invoices.json
```json
[
{
   "customer":"BigCo",
   "performances":[
      {
         "playID":"hamlet",
         "audience":55
      },
      {
         "playID":"as-like",
         "audience":35
      },
      {
         "playID":"othello",
         "audience":40
      }
   ]
}
]
```

# Output

## PlainText
```
Statement for Bigco
  Hamlet: $650.00 (55 seats)
  As You Like It: $580.00 (35 seats)
  Othello: $500.00 (40 seats)
Amount owed is $1730.00
you earned 47 credits
```

## HTML
```html
<h1>Statement for BigCo</h1>
<table>
	<tr><th>play</th><th>seats</th><th>cost</th></tr>
	<tr><td>Hamlet</td><td>55</td><td>$650.00</td></tr>
	<tr><td>As You Like It</td><td>35</td><td>$580.00</td></tr>
	<tr><td>Othello</td><td>40</td><td>$500.00</td></tr>
</table>
<p>Amount owed is <em>$1730.00</em></p>
<p>You earned <em>47</em> credits</p>
```
