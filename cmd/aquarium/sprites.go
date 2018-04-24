package main

import "github.com/blinsay/aquarium/termdraw"

var anglerL = termdraw.NewAnimatedSprite([]string{`
 /
/\/
\/\
 \
	`})

var anglerR = termdraw.NewAnimatedSprite([]string{
	`
 \
\/\
/\/
 /
`})

var seaweed = termdraw.NewAnimatedSprite([]string{
	`
)
(
)
(
 `,
	`
)
(
)
(
 `,
	`
(
)
(
)
`,
	`
(
)
(
)
`,
})

var minnowL = termdraw.NewAnimatedSprite([]string{`
<><
`})

var minnowSchoolSmallL = termdraw.NewAnimatedSprite([]string{`
<><
<>< <><
  <>< <><
`})

var minnowSchoolMedL = termdraw.NewAnimatedSprite([]string{`
<><   <><
<>< <><   <><
  <>< <><    <><
      <>< <><
`})

var minnowUnionL = termdraw.NewAnimatedSprite([]string{`
                <><
            <>< <>< <><
        <>< <>< <>< <>< <><              <>< <><
    <>< <>< <>< <>< <>< <>< <><        <>< <><
  <>< <>< <>< <>< <>< <>< <>< <><     <>< <><
 <><     <>< <>< <>< <>< <>< <>< <>< <>< <><
 <><     <>< <>< <>< <>< <>< <>< <>< <>< <><
 <>< <>< <>< <>< <>< <>< <>< <>< <>< <>< <><
   <>< <>< <>< <>< <>< <>< <>< <><    <>< <><
     <>< <>< <>< <>< <>< <>< <><        <>< <><
         <>< <>< <>< <>< <><              <>< <><
              <>< <>< <><
                 <><

`})

var mackrelR = termdraw.NewAnimatedSprite([]string{
	`~~.~~.~~. ><(((°>`,
	`..~..~..~ ><(((°>`,
})
