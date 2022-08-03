% 2022-08-03

# Week 31
## Naveen.md on 2022-08-02

fly in on friday

prioritisation happening
roadmap
u planning

susan shared slides
compliance is the priority
DSA/DMA

RMG - something we are reactive to on the store
the number of the things we want to add protections to

for U timeframe: how much do we invest into doing SC in house vs adding 3P, incl these capabilities under this umbrella

have meeting with Eugene right afterwards
Krish asked if Naveen can become primary PM for PbA

Not just improved, but allow others to build SC design of their own

endgoal is to have a PbA moment. Can we drive it by UI we are pushing on everyone? Or do we nmeed a different strategy.




idea permissions

## William.md on 2022-08-02

the only prepandemic person on the team is Aaron
let's cancel programme review and do email instead.

Long discussion about what we should present at programme review and how, which ended up with cancelling said review altogether.

## 2022-08-02.md on 2022-08-02

There was some discussions, meeting with DaveK / Sameer. DaveK for a while .. every time there’s some thing about malware, Sameer is like how did it get in? Dave tried to communicate this is just expected, we have millions of apps, we not gonna catch everything. Part of this discussion - why isn’t some of this behaviour possible on iOS? 

From Dave - this was a brief conversation, but led to many people panicking. Dave’s request we want to get back to Sameer this are the top abuse vectors on Android. That led to lots of discussions about this. Made a little bit harder by - there’re few people who are coming up to speed on this for the first time. Doing a lot of internal education. Sebastian was deeply involved.

Two things stand out:

-   Dave talks about API abuse, some stuff is not platform specific - e.g. obfuscation. DCL to same extent. 
    
-   For lot of abuse types, the view from Sebastian, backed up by PHA trends is that quite a lot of these types of malware, while abuse Android APIs, they still work without abusing APIs – e.g. tollfraud. 100% social engineering.
    


Was a bit worried last week when saw we are behind
Did good cleanup of go/angel-truth
Met android performance team


We don't call it AiAi anymore. This was the name research team used, and this is prod. We decoupled conceptual part of it - PCC - from the APKs we ship. 

Android System Intelligence.  PCC is the new name.

PM (Shan) doesn't like name PCC, he pushes Android System Intelligence.

Two teams:
- volnov@ - Spina team, shipping ASI APK, leads of governance part, do approvals, talk to system health team, security, privacy etc. Now own QA process. AiAi council approval. 1 LGTM from Android, 1 from Cerebra (Wei). Eugenio delegates to Sergey.
- Wei - Cerebra team weih@

Idea: production systems owned by PCC, research owned by Cerebra.

Feature side: 150 SWEs submitting code into aiai directory/package. SysUI, transformer, ppl who work on Cerebra, Android Audio framework, TLV team. 20+ features that we approved for Android T launch. 

For production systems we work directly with Sergey's team. If something doesn't exist today, Wei's team. Cerebra is research team and sometimes they work on projects that would take forever and never launch. E.g. Astrea plugins - is likely to never launch. Not realistic, Samsung n their devices have their own app doing content capture. Even if we make it work on Pixel, won't scale to ecosystem. 

Chronicle - we are pretty invested into it. One of the big features in ASI right now doesn't use Chronicle, so they decided to build their own infra. Chronicle - annotation that allow you to explain how your data will be persisted, so you can say that this can be stored with this TTL, this only in memory, etc. 



Most of the team in MTV, staying home to focus

Was supposed to go but visa expired. Waiting for it. Appointment in October.

Bought an e-bike. Exploring Cambridge a bit. Bought an inflatable kayak, going over river Cam.

Mum - Geology, retired, conference in Glasgow. 

Working on dialogues. Final phase - working on some logging. Once done will send to Play Testing. Metrics still not started. Sketching phase, collecting feedback. As soon as done with dialogues will start writing some docs. 

Gauarav joined the team. Today Inara is in the office. Situation not ideal because most team in the U.S.

Team is doing well. Last time feedback was - happy with the team, this hasn’t changed. Something new happened - Bessie left, but she’s seating just next to us. New people doing quite well, Luke ramped up quite fast. Haven’t spoken to Gaurav much, but do have a feeling he’s gonna be good. 

Niko - schedules career conversations. Always has initiative to listen and propose new things. Made explicit eventually want to get promoted to L5.



Charmaine and Shan are enthusiastic. 

Discussion in U planning on perm rationales.

CMU study: will start in September. 


Comitted to figure out what next APP, APP v2
We learnt a lot about being able to protect users automatically through Carter. Want to take best of it and move to APP. 

Question of: how do the different policies change given auto-enrollment. E.g. SK won't be hard requirement. 

We should also be looking at what we do with AP4A. Team presented concept 1 or 2 months ago, concept approved but they need to come with solution now. They due to come with it in August - good time to get together and decide.

At the same time - enhanced protection mode. THis one came out from protected computing effort. ProtComp helps to do e2ee. What's the high level opportunity? Enhanced Protection mode - similar to APP, one opt-in, but covers you across privacy/security/safety. Some overlap with APP. Going through scoping. Shuvo working with them. joshi@ Josef key person, in PSS. PM Ane Fabou (?) privacy org, 60% capacity. Shuvo is also putting attention to this.

OneSafety - slower than I like. Will submit as 2023 planning project. 


Angel Droidfood Go!!!
PbA dead, long live PbA
Roadmap exercise
FMD fire drill
ODAD conversations


Angel Droidfood Go!!!
PbA dead, long live PbA
Roadmap exercise
FMD fire drill
ODAD conversations

wife was diagnosed with a tumour, benign
embellisation - they go in and cut blood supply to tumour, did this, painful process
recovering from it, hoping this helps

she wasn't adjusting to pain medication
starting to feel better now, at least from pain mgmt perspective

daughter finished her high school, wife wanted to spend time with her

insightful team - the team has grown significantly, people do great job in their circles, but not necessarily aware of what's going on.

H2 HC on-hold. Re-examining, guidance from finance. Worried about last minute HC given for CY for Ads, Policy stuff ... all of this up in the air. Guy's team had 6-7 HC for variety of things, trying to help Zhao - she had open heads. There are some L3s who are hired but they don't have home right now, that's one way to save some HC. Bad situation.

backfills are also on hold. 

anyone we can save by keeping an eye on what's they are looking for - we should try to save. 

U.S. GDP came - 0.9%, second consequtive decline in GDP. We are definitely in recession. 

As a leader we need to do everything that the team understands the big picture and the team feels psychologically safe. Responsible thing for Google to put pause at the moment, better than layoffs later. 

FMD - give heads up to PR that this is happening. 

Really liking this team

Right now working on:
1. Driving the content capture further - from our demo to release
2. Looking into releasing stalkerware in October, having heuristic Hades protected
3. Helping with internship
4. Longer-term -- looking into DCL protection for in-memory dex loader.
**


- Notifications are ready to go -- someone done something in Photopicker
- 4 notifications privacy - noitificaiton listener, a11y, bg location and permission auto-revoke, HUN sent by Permission Controller. They have top little string from which app came from, this app would PbA. That's all done, behind flags
- Remaining two pieces:
	- Rebranding GPP notifications to do the same thing, not that hard, 5 SWE-days + a lot of talking to Play leads. No product alignment with Play, Play leads are pushing. 
	- SC itself - have to put branding in the right places on cards, some graphics are wrong, 5 more SWE/days.
- Unanswered questions:
	- When is it launching?
	- Given launching requires flipping flags in system image, which devices this launches on?
	- might launch in QPR1. 
	- Android Platform Phenotype flag flip - factored in into above. 


## Gabriel.md on 2022-08-02

Most of the team in MTV, staying home to focus

Was supposed to go but visa expired. Waiting for it. Appointment in October.

Bought an e-bike. Exploring Cambridge a bit. Bought an inflatable kayak, going over river Cam.

Mum - Geology, retired, conference in Glasgow. 

Working on dialogues. Final phase - working on some logging. Once done will send to Play Testing. Metrics still not started. Sketching phase, collecting feedback. As soon as done with dialogues will start writing some docs. 

Gauarav joined the team. Today Inara is in the office. Situation not ideal because most team in the U.S.

Team is doing well. Last time feedback was - happy with the team, this hasn’t changed. Something new happened - Bessie left, but she’s seating just next to us. New people doing quite well, Luke ramped up quite fast. Haven’t spoken to Gaurav much, but do have a feeling he’s gonna be good. 

Niko - schedules career conversations. Always has initiative to listen and propose new things. Made explicit eventually want to get promoted to L5.


## Ed.md on 2022-08-02

There was some discussions, meeting with DaveK / Sameer. DaveK for a while .. every time there’s some thing about malware, Sameer is like how did it get in? Dave tried to communicate this is just expected, we have millions of apps, we not gonna catch everything. Part of this discussion - why isn’t some of this behaviour possible on iOS? 

From Dave - this was a brief conversation, but led to many people panicking. Dave’s request we want to get back to Sameer this are the top abuse vectors on Android. That led to lots of discussions about this. Made a little bit harder by - there’re few people who are coming up to speed on this for the first time. Doing a lot of internal education. Sebastian was deeply involved.

Two things stand out:

-   Dave talks about API abuse, some stuff is not platform specific - e.g. obfuscation. DCL to same extent. 
    
-   For lot of abuse types, the view from Sebastian, backed up by PHA trends is that quite a lot of these types of malware, while abuse Android APIs, they still work without abusing APIs – e.g. tollfraud. 100% social engineering.
    

# Week 30
## Elliot.md on 2022-07-28

Was a bit worried last week when saw we are behind
Did good cleanup of go/angel-truth
Met android performance team

## Nino.md on 2022-07-28
Really liking this team

Right now working on:
1. Driving the content capture further - from our demo to release
2. Looking into releasing stalkerware in October, having heuristic Hades protected
3. Helping with internship
4. Longer-term -- looking into DCL protection for in-memory dex loader.
**

## Guemmy.md on 2022-07-28

Comitted to figure out what next APP, APP v2
We learnt a lot about being able to protect users automatically through Carter. Want to take best of it and move to APP. 

Question of: how do the different policies change given auto-enrollment. E.g. SK won't be hard requirement. 

We should also be looking at what we do with AP4A. Team presented concept 1 or 2 months ago, concept approved but they need to come with solution now. They due to come with it in August - good time to get together and decide.

At the same time - enhanced protection mode. THis one came out from protected computing effort. ProtComp helps to do e2ee. What's the high level opportunity? Enhanced Protection mode - similar to APP, one opt-in, but covers you across privacy/security/safety. Some overlap with APP. Going through scoping. Shuvo working with them. joshi@ Josef key person, in PSS. PM Ane Fabou (?) privacy org, 60% capacity. Shuvo is also putting attention to this.

OneSafety - slower than I like. Will submit as 2023 planning project. 

## Guy.md on 2022-07-28

Angel Droidfood Go!!!
PbA dead, long live PbA
Roadmap exercise
FMD fire drill
ODAD conversations

## Khawaja.md on 2022-07-28

Angel Droidfood Go!!!
PbA dead, long live PbA
Roadmap exercise
FMD fire drill
ODAD conversations

wife was diagnosed with a tumour, benign
embellisation - they go in and cut blood supply to tumour, did this, painful process
recovering from it, hoping this helps

she wasn't adjusting to pain medication
starting to feel better now, at least from pain mgmt perspective

daughter finished her high school, wife wanted to spend time with her

insightful team - the team has grown significantly, people do great job in their circles, but not necessarily aware of what's going on.

H2 HC on-hold. Re-examining, guidance from finance. Worried about last minute HC given for CY for Ads, Policy stuff ... all of this up in the air. Guy's team had 6-7 HC for variety of things, trying to help Zhao - she had open heads. There are some L3s who are hired but they don't have home right now, that's one way to save some HC. Bad situation.

backfills are also on hold. 

anyone we can save by keeping an eye on what's they are looking for - we should try to save. 

U.S. GDP came - 0.9%, second consequtive decline in GDP. We are definitely in recession. 

As a leader we need to do everything that the team understands the big picture and the team feels psychologically safe. Responsible thing for Google to put pause at the moment, better than layoffs later. 

FMD - give heads up to PR that this is happening. 

## Eugenio.md on 2022-07-27

We don't call it AiAi anymore. This was the name research team used, and this is prod. We decoupled conceptual part of it - PCC - from the APKs we ship. 

Android System Intelligence.  PCC is the new name.

PM (Shan) doesn't like name PCC, he pushes Android System Intelligence.

Two teams:
- volnov@ - Spina team, shipping ASI APK, leads of governance part, do approvals, talk to system health team, security, privacy etc. Now own QA process. AiAi council approval. 1 LGTM from Android, 1 from Cerebra (Wei). Eugenio delegates to Sergey.
- Wei - Cerebra team weih@

Idea: production systems owned by PCC, research owned by Cerebra.

Feature side: 150 SWEs submitting code into aiai directory/package. SysUI, transformer, ppl who work on Cerebra, Android Audio framework, TLV team. 20+ features that we approved for Android T launch. 

For production systems we work directly with Sergey's team. If something doesn't exist today, Wei's team. Cerebra is research team and sometimes they work on projects that would take forever and never launch. E.g. Astrea plugins - is likely to never launch. Not realistic, Samsung n their devices have their own app doing content capture. Even if we make it work on Pixel, won't scale to ecosystem. 

Chronicle - we are pretty invested into it. One of the big features in ASI right now doesn't use Chronicle, so they decided to build their own infra. Chronicle - annotation that allow you to explain how your data will be persisted, so you can say that this can be stored with this TTL, this only in memory, etc. 


## Simon.md on 2022-07-27

- Notifications are ready to go -- someone done something in Photopicker
- 4 notifications privacy - noitificaiton listener, a11y, bg location and permission auto-revoke, HUN sent by Permission Controller. They have top little string from which app came from, this app would PbA. That's all done, behind flags
- Remaining two pieces:
	- Rebranding GPP notifications to do the same thing, not that hard, 5 SWE-days + a lot of talking to Play leads. No product alignment with Play, Play leads are pushing. 
	- SC itself - have to put branding in the right places on cards, some graphics are wrong, 5 more SWE/days.
- Unanswered questions:
	- When is it launching?
	- Given launching requires flipping flags in system image, which devices this launches on?
	- might launch in QPR1. 
	- Android Platform Phenotype flag flip - factored in into above. 

## Giles.md on 2022-07-27

Charmaine and Shan are enthusiastic. 

Discussion in U planning on perm rationales.

CMU study: will start in September. 

## Anurag.md on 2022-07-26

code freeze for beta next week

team has 8-9 parallel work streams going on
our job is to make sure the ads sdk like admob can fully move into sandbox
for that they need lots of features
have different levels of involvement into them

longer term - planning for U, policies, partners ...

team has recently increased, some folks started report to A
9 directs, lot of time just talking to them

Two L5s, one of them TLM
both very reliable
can count on one of them to get almost anything done

other person has lots of strengths driving decisions across group but don't land it
start discussion, then get bored and move on

L4s are pretty reliable. sometimes few people get lost in the weeds

Sorin is probably the best L5 A worked with. L5s in current team are in the same zone but aren't as reliable. 

lots of folks in MTV, lots of chats, lots of emails, docs. No way to be on top of everything
Almost everyone on the team is new to Android


## William.md on 2022-07-26

Let's do a mini-review of where we are with everything. Things that I am interested in:

- Stalkerware - are we good to go? What are the current test numbers if any? When are we launching?

Plan on dogfooding in August. The original MVP we ship the heuristic inside of ODAD app directly, without model protection. We are still going ahead with that. If we can get Hades support, it's a bonus. 

We are in a good shape. In dogfood want to see what FPs look like. If we are not flagging randomg apps, we can launch. One blocker is people having GPP turned off. 

Allowlist ODAD notification in QPR2.

- Various content capture tasks:
	- What's our level of confidence that this will work?
	- What are our test numbers?
	- What are the blockers? (I know about Astrea)
- What's the latest on OEM adoption?
- What other areas are we considering?  What are our mid-term plans?


## Jim.md on 2022-07-25
Making progress towards initial dogfood of the new web ui
Long tail of style cleanups and that
All the basic functionality should be done next week / 2
Original date was start of august, will be second week.

Private dogfood, will start FMD team, then Spot

Meeting Finder UX person this afternoon
Finder app made some UX changes
Have to decide if we should launch with exact same UI or should update to match

## Niko.md on 2022-07-25
Have a chat re EPIC / PSIC litigation

Crazy install flow diagramme

Tried to explain Sunny how things work and why, hence diagramme.

Gaurav starts next


# Week 29
## Saagar.md on 2022-07-22

Google is very slow at ramp-up. Doing starter bugs.

3 bugs:
- Issue ODAD din't have certain permissions and would crash without these permissions. The bug was to check permissions. Seemed like it was pretty easy but ended up changing a lot of things.
- We have something unused, should we delete it? 
- For one of the tests we want to make sure we are doing some reporting. Not started

Def not something I'd want to work a whole lot. Want to do reverse engineering / malware research. 

## Shriya.md on 2022-07-22

Everybody on the team said that -- about burnout

Current plan feels feasible

Everyone is very helpful and welcoming

Trying to work a bit more on UI part now, previously Phonesky

Notting Hill

## Gaurav.md on 2022-07-22

6-7 years of experience
worked at small, medium and large companies
coming from Amazon
self-evaluated as strong L4, ready for next step
got a little bit experience working with applications
masters in machine learning, hoping to contribute in that area

thesis - reinforcement learning, intrinsic reward, more complicated prisoner's dilemma
bsc - biology & computer science, wanted to be a doctor but changed my mind
too much memorisation

ramp-up stuff
ODAD client. starter project: make us PCC compliant. Encrypting some of the data fields for cached features.

plan  - ramp up on client, and get on the server as well. 

product: from codebase PoV pretty good. no spiderweb code and logic. on macro-level still getting all the context. 

bus factor on the team is very very low.  as a consequence team feels very siloed right now

haven't seen a lot of "here's my design doc, let me present it to the team".

William keeps an eye on CLs and calls out things that go really well.

First time having a technical manager. 

All managers in the past used to be technical but not anymore. 


## Mark.md on 2022-07-21

Last week before holidays

Going to CO with girlfriend, going to be married there next year

Her family and Mark are both from Michigan

SHe's first her residency

fear as time goes on (a) it's not susteinable, I'll be pigeon-holed (b) ability to advance will be limited

challenges are becoming less technical and more organisational


## Narayan.md on 2022-07-20


> **from bdc@ update**
> Angel update
> - Reached M-09, confident we'll make it to M-11. 
> - Rollout will have to be postponed because M-11 will ship too late for flag flips, practically rollout in January, but eng deadlines remain unchanged
> 
> - Two fires, one related to broken pending intent sent by settings, another related to location toggle, both contained, will be addressed in QPR + OEMs will pick up relevant changes.
> - Code complete for M-11 mid-August, tools dow on the initial release 14/09
> - Working on plan for U / 2023. Would really love to use Mainline SDK, but seeing some pushback from the team - they are not very interested in supporting us.


Live in Pimlico, 100m from the river

Dogfooding modules. Not doing dogfood is not a bad thing necessarily. But we have to have experiments framework. We should either have dogfood programme that works, or we have to have run these studies on production population. 

## Guy.md on 2022-07-20
Haven't heard anything concrete about Sundar's email, was optimistic
Sundar said there will be priority towards engineering

Had a meeting with Anita. 
She said about launch of LE devices with Finder. She asked if this will move forward, debate opt-in/opt-out. She mentioned she has some concerns about people safety - people putting these tags in other people bags and tracking them. She head there may be some effort to mitigate this.


## Guanhuan.md on 2022-07-19

Want to get U req done

Categories:
- Apps
- Device
- Accounts, and privacy
- System

System is scary thing you don't want to touch

Concept of safety sources/categories as seats

Explanation is somewhat tedious!


## bdc@.md on 2022-07-19

Angel update
- Reached M-09, confident we'll make it to M-11. 
- Rollout will have to be postponed because M-11 will ship too late for flag flips, practically rollout in January, but eng deadlines remain unchanged

New strategy it's not only about having UI in Mainline, but API as well.

- Two fires, one related to broken pending intent sent by settings, another related to location toggle, both contained, will be addressed in QPR + OEMs will pick up relevant changes.
- Code complete for M-11 mid-August, tools dow on the initial release 14/09
- Working on plan for U / 2023. Would really love to use Mainline SDK, but seeing some pushback from the team - they are not very interested in supporting us.


Talk to robertogil@google.com

Google Play Protect update:
- Working on just-in-time verification. Big deal, we are seeing lots of polymorphic malware in the ecosystem. Likely to improve our AV-Test scores significantly
- In parallel we are thinking about GPP migration to Angel - something that we wanted to do from the very beginning.


## Giulio.md on 2022-07-18
Currently in Italy, no A/C
Peruggia - middle 
2 weeks

Finalised MVP
Working on finalising some of the strings, some of the strings in  MVP not ready
Working mostly on system server side
QPR change is on Elliot
Going to drop our module to QPR

Last week found out there are GMS Chinese devices. We decided to disable our GTS tests for these devices. Now some failures with Go devices, need to investigate this.

Go/Wembley devices - we are functional but slow. Karishma and David were doing some tests. 

An OEM raised a request to waive GTS test, because Go. Go devices don't have Mainline right now. 

Simon: haven't had that much time together, holidays, then covid, but we did the expectations talk, checked them in. 


## Yuri.md on 2022-07-18

Leyton

Animations in Preferences

expand and collapse for groups in SC

preferences render in Recycler View, and wrapped in Fragment. 

remaining: logging, motion design for everything

worst case: postpone animation, leave uncollapsed

another way: change rendering, make all groups/safety entries components of a single preference

we want to get rid of preferences altogether. Jetpack compose etc

QA generally positive

Simon - 2 1:1s, 1 intro. Too concentrated on work right now. 

# Week 28
## Inara.md on 2022-07-15

Finishing up autodisable and update. Probably few more weeks of work. Plan is outsourcing some stuff to Gaurav.


## Guy.md on 2022-07-15
- Pending Intent fire in Angel
- Location toggle fire in Angel
- Conversation with Product and UX. Jeroen.

G heard from Guanhuan that we are taking complete ownership of Angel on product side.

Chatted with Simon, was pretty good. Will have recurring 1:1s. Mentioned some of the thoughts about Angel. 

- JIT update
- FMD update

Talked to Naveen. Mentioned about Finder team in Israel, Ronald, Product team etc. The only thing they are interested in is launching new low energy devices. 

- Feedback about visits


## Nikolay.md on 2022-07-14

Usage Stats give better foreground info

## Tyler.md on 2022-07-14

Team feels low energy for the last couple of weeks, people feel more tired. Completed this milestone but hasn't come to fruition yet. 

You can feel the difference in atmosphere. 

In general feel like broader meerkat team is doing better.

Something that people enjoy in person would be more appropriate. For team it would be good to take some time to stop and acknowledge what we've done so far. Might be more satisfying when product gets out to people.

Was around for Khawaja visit. Not sure i understood our place any better in the org. Personally didn't find tons of values in this. Feel like Meerkat org is really tight, where above is just a chain of directors, unsure what Khawaja's impact on what we do or expertise he provides. 

## Guanhuan.md on 2022-07-14

Recovering from 01:30 AM landing


## Jeroen.md on 2022-07-14
How to organise ourselves?
Working with all the different teams etc
Last couple of weeks: trying to start creating some of our narrative how can we be talking about all of these projects in order to streamline all these different things

https://docs.google.com/presentation/d/1dkGkSFWhDl1DmWeAw5bS7Zs6I1TdoQvncx5if7AlOLQ/edit?resourcekey=0-JaxMXeI7-PXWItgFkomSoA#slide=id.g130ff56195f_0_0

other random projects: Rubidium, PCC, Permissions projects... how to make sense of all of it? There's no single framework that allows us to make decisions in consistent way

want to align more with ASAP given this is where security and privacy defined for Android


## Naveen.md on 2022-07-14

in LON on 6th of August


## William.md on 2022-07-12
Spoke to Shivan AiAi , repots to Sergey who reports to Eugenio https://moma.corp.google.com/person/shivanker

Content Capture as contingency plan. Wanted to do it via Astrea, but there's pushback. If we can'd to it via Astrea we do it via AiAi, which will send verdict over to ODAD.

Astrea is the open source component of AiAi. The only thing it does currently is federated analytics. For ODAD v2 they add model downloading. 

The downloading part is fine, on its way except for more and more issues.

Content capture - push back from system health team. They thing that spending 10mb for transparency is justified. AiAi team needs to find other clients for this, such as us. We need to get exec support to show that this is important. 

Steve is aware. 

Contingency plan to do it in AiAi. 


## RichardN.md on 2022-07-12

Worried about Blackhat talk

Egor passed his UK driving test
Lena is joining ASA Red Team


## Isabelle.md on 2022-07-11
Mtg on Friday was very helpful.

Still working on boqweb.
Goes at reasonable pace. Aiming end of July dogfood. Finder hasn't launched app dogfood either. Plan is to enable for personal accounts.

FMD killed in work profile - because work profile admins can turn it on and off, too much power. 

What about DO? Need to check that. Expected behaviour is the same.

GRAD. Not sure how it's gonna work. Expectations mostly based on OKRs. Didn't put P3 OKRs.

Team: since Anurag left, we have less engineering capacity, Sorin is quite stretched. Other than that feel we work well as a team. Sorin is doing great as a tech lead. 

Have regular 1:1s with Sorin. Mostly talking about tech stuff, also about how team can work better. Each sprint there's a certain amount of stuff which is maintenance / mandate. So everyone is working only 2/3 on actual stuff, 1/3 on toil. In future it would be great to have a rotation, but need more engineers, so less context switching for everyone else. 

Sorin started doing some more boqweb stuff, discussing/complaining about how it works. 

Muhammad: main thing has been - he doesn't necessarily have the time to get up to speed with everything we are doing. He started coming to a decent number of the meetings. There's a certain amount of knowledge gap there. 

Jim  made a high-level diagram and put it onto docs page. 

M interaction - unsure yet, kinda fine. 

It's great to have a tech lead manager rather than just a manager. 


# Week 27
## Shuvo.md on 2022-07-07

Safety summit in New York, 2nd and 3rd.




## Peter.md on 2022-07-07
Still waiting for approval transfer.

Assuming that working at L6 now.
Do feel right now it is more stressful than it should be because we are hiring two people to work on the verifier, but right now I'm doing a lot of this work on my own. Should be leading people, but there are no people to lead.

PV need more people. Have agreement with Niko these people will work on PV. Luke and Gabriel are also contributing. 

## Nikolay.md on 2022-07-06

Spoke to William about usage stats, can't use them in production
We have the permission, but this was only for network stats, and usage stats privacy team said we shouldn't use them for privacy reasons. 

Was not part of that discussion, we should make the case for that

Even without usage stats we probably have enough of a signal. Not convinced that usage stats will add much else. Fg/bg will be the only thing that'll help.


## Muhammad.md on 2022-07-06

Guy came a bit helpless and quiet in meetings. Max used a word that he almost didn't want to be in the room with them.

With Khawaja they found his answers all on point, but a bit too high level. 

Goals and Priorities of AASE, but what didn't happen was just that - how does this map to their function specifically. 

Niko gave the opposite. His people were more engaged with Guy than Khawaja.

Niko set up an OKR around convergence of GPP and Angel. 

Slight concern with Guanhuan and product roadmap for Angel.



## William.md on 2022-07-06
Had first real vacation

Wife got pregnant then nearly got a miscarriage

Saagar heard someone saying (Chad?) he hopes that on-device won't be successful. 


## Anita.md on 2022-07-05
Today's the meetings day
Priority #1 task - beginning a design doc
LKL for wearables

Muhammad is doing a good job
Meeting with Khawaja went well

Expectations!

## Gabriel.md on 2022-07-05

1,5hr commute, but can work on a train
Main limiting factor commute price
Renting place, Cambridge
Looking to buy something
Partner is flexible, she needs to go to the office, currently goes only 1 every two weeks

Conversation with kshams - the big one
Also had a smaller meeting 
possibility of implementing func client that allows server to tweak configuration

With Guy talked more about abstract things

Planning to go to MTV, but U.S. expired
Renewing US visa
Expect to have one by October


# Week 26
## Elliot.md on 2022-07-01

SysServer side is going well, lot of stuff left over to do
Some items can be punted
Shriya/Max/Jan are going to start looking for tasks to do on Android side, at least Jan is almost done with GCore and Shriya/Max will be done next few weeks
Looking pretty confident

MVP state is OK, could be better

Rescan button - sends broadcasts to all the sources and waits for them to come back
Rescan vs Page Open refresh - we let them know if this is a high-cost operation vs just page open

Practically only GPP does use this

AcctSec always makes RPC in the background

Presubmit takes hours to run, on average 3hr if it doesn't fail. Too easy to bypass presubmit. Improve presubmit - only run stuff that needs to run. Fixing typo in the code will require rebuilding entire source code


## Muhammad.md on 2022-07-01
David has been diagnosed ADHD and dysgraphia. Took him 6 months to tell us. He confided in Simon and then Simon moved off, and then he joined us again.

Asks for consideration in terms of types of work we are allocating to him. 

He written a document, shared with Muhammad.

He'll get ADHD coaching and support.

It's about managing his attention.

kshams didn't have any experience with this, but he told go and consult HR

David is asking for interesting work. He calls out specifically boring work such as documentation. This is the reason he didn't take on OEM integration guide - not stimulating, boring, etc. 

Prabal lost his grandfather. Simon knows more. M knows through Elliot. He'll be less available and likely working remotely. His first few months out of homeland, double-tough

GRAD expectations: M was trying to align expectations, the level of it, across David/Elliot/Simon and then bubbling it up to M level.

Goes to Boston for holidays, all the way to Charleston.


## Barak.md on 2022-06-30
Release trains
If we are looking to include cards/advice from Chrome
We can use GmsCore as a fallback because they have integration with Chrome
Use go/angel-fr


## Fit_Calls.md on 2022-06-28

Background in electrical engineering
MSFT 13 years
Most - identity & security
MSFT first boxed AV product (Defender), trustworthy computing initiative
Ran account security for MSFT account
One of the world's first ML cloud anti-malware service
London from 2013
Worked from Skype
Left MSFT, doing series of start-ups
EdTech - teaching kids how to code
Connected vehicle, PE backed
Last couple of years - Feedo, identity verification


## William.md on 2022-06-28
Working on GRAD with people on the team


## Jason-Cornwell.md on 2022-06-27
Dismissal - should we change copy to more concrete ack, maybe user doesn't understand the risks?

# Week 25
## Sudhi.md on 2022-06-24

Very large effort that Ed is the PM for. 
Responsible APIs.
Looking at all the framework APIs holistically and finding which of them have potential for abuse
Fairly well defined roadmap 
Problem is S see it is: APIs Ed and Chad picked primarily based on intuition.
What is missing is data
Independent of API work, we know that there're these categories of malware. We don't know which APIs are abusing.
There's no analysis doc of existing malware and mapping to API that are potentially harmful
Asked Ed to see if he can get Sebastian to spend some cycles on identifying such APIs / malware

In Rubidium we have SDKs living in the sandbox
But this SDK is also house for 3P anti-fraud
but why doesn't it come as Java SDK?
Because it will be reversed really quickly, it can be introspected

The way admob works they use DCL via GmsCore dynamite, to obfuscate signals they collect. Even if they need 1 signal, they collect 10 of them, and they rotate them every 15 days.

Talking to ZacharyLF@ from AdFraud. Chatting about standartising on Lua. VM is lightweight, ~kilobytes. 

## Nikolay.md on 2022-06-24

Went to Bulgaria for Easter


## Sebastian.md on 2022-06-24

<a/c priv>


## David_Coffin.md on 2022-06-23

GRAD

## Sorin.md on 2022-06-23

2SV being launched
BoqWeb have fully working framework
We have
- Initial page load
- Rendering devices
- Actions you can perform on the devices
Working on:
- Special UI for rich actions 

Jim created a list of all remaining stuff we have to do for dogfood, current ETA early August (might be skewed)

Changed the HTML completely layout. So it looks pretty good on mobile. 

Currently in expectations process
Decided on Q2 OKRs

Finder: nothing changed on eng side. Trying to get more contact and be more in touch with the product side (Guy/Sham/Ronald), 

M plans for Finder to own the roadmap

Get along well with Muhammad. He's very involved into the team, he's excited about FMD

He handled expectation setting very well. Initially had sense from the team, mixed feeling about GRAD, but he managed to change the mind of folks, everyone is willing to give GRAD a chance. 

## Jeff.md on 2022-06-22

Companion Device Manager, CDM. 
ability to sync permissions
Long term plans for U+, Mainline module on the way
Samsung got them to backport not just to T and S, but also R
New Mainline module!!

Samsung expects this yesterday

bdc@ - slight adjustments to quarterly planning, because Mainline

Dianne also wants to get to quarterly API bumps
QPRs can be release vehicles for API releases


Hansson - need to speak to him. TLing mainline updatable API. 
Rubidium can not fail. 

Mainline SDK APIs are likely to go through Jetpack

DCL

developers have a use case for this.
Split APKs should solve this in theory

Are there feature splits in Split APKs?

ODAD concerns - if we create mechanisms to allow on-device detection, this might be a venue for abuse

advertise split apks
then push really hard for DCL on policy end

## Aaron.md on 2022-06-22

Starting to get Westworld data from dogfood
Still need approval for prod launch

Saagar hasn't quite started yet, he's just about to start his first bug
Gaurav started and doing noogler bugs
Nino got right off

Ishtiaq - going OK. RAB effort, the way it's working, Aaron is leading the effort, going to mtgs with Samurai team, and intervene if necessary, and we are making progress on RAB. The most recent CL, ODAD was writing a file sdcard, and Samurai team was responsible for retreiving it - that one was OK. We needed some unit tests for it, asked to implement. Some struggles in implementing unit tests. Talked about this with William. Sometimes I don't know what level of answering questions.

## Yuri.md on 2022-06-21

Finishing UI for QA deadline tomorrow

On time, whatever is ready is ready for QA.


## HMD meeting.md on 2022-06-21

Anderson - HMD, experience manager, Google-related topics, Dialer, Enterprise
Lian - HMD SWE
Luo Jun - HMD application team, SWE
Eric - Google TAM for HMD

Anderson: is this Android S or T settings UI? Security & Privacy or Safety?
  

## William.md on 2022-06-21

The team said they are fine with coming back
Ishtiaq wants 2 days a week

Spoke to Nikolay about progress. He showed his work. Feel less concerned now. He does have everything that he needs to have incl network usage, but he's not very good with giving updates.


## Guanhuan.md on 2022-06-21

Scalable UI, scalable integrations, API


## Giulio.md on 2022-06-20

Working on system server side, Permission APEX
Migrating google3 logic
Need persistent storage
e.g. need very simple for loops
Using XML file for storage
Using single lock for everything

Tests!
We needed a volounteer / PoC for QA team, Giulio did
We should get very close to QA start testing
Two weeks from now
Giulio is out next week but doing all the preparation this week
Targeting parity with Angel on S
Coordinating test plans, cloned.
The approach to do it from sources perspective
Jan reviewed all GmsCore tests

Tablet UI?


# Week 24
## Dianne.md on 2022-06-14

Status Update
OEM update
Mainline APIs


Robolectric in platform - complete chaos abstraction between testing env and platform. 

## Dave.md on 2022-06-14

Ash to lead responsible APIs

Chad might be on the way out?? Vish will be moved from under Chad


## Naveen.md on 2022-06-14
wanted to get better level of clarity on our relationship with OEMs

e.g. Samsung and the fact that we agreed to do API only version for them.

BD team has closer relationship with OEMs than TAM

Two things Android can mandate:
- CDD
- GMS

we were previously going to put Angel to CDD. Sameer pushed back, saying this should be GMS, because we have more leverage there. Samsung escalated all the way to the exec level. There was discussion with Hiroshi where they said we are doing nothing. They conceded on Rubidium. 

**FMD**



# Week 23
## Sebastian.md on 2022-06-10

In the office July 12

28 devices with compromised platform signing keys

they are planning to double sign everything, custom Samsung signed process

Lukasz pushing to publish it


## Narayan.md on 2022-06-10

Nate, Hai stepping up to manage permissions team. Dan is helping to keep things ticking along. 

Looking at hiring a backfill for Eric. 

Taking over QS by LON is OK

go/join-api-council


## Guy.md on 2022-06-09

**Angel**

Went well -- RTO policy, hybrid is going well. The fact that some people are fully remote like Marie in Canada is working well. 

Team spirit / attitude. 

Fast growth of the team and how they manage to onboard people, growth done in the correct way.

Technical strengths / skills. 

To keep in mind -- they are working in a high pressure environment, crazy deadlines. Makes it hard to make it in sustained fashion. Hopeful things will change once they finish with the current launch. 

One person was vocal about it is David. He even mentioned at some point (1:1) that he's trying to assess what's right for him. He wants to work towards promotion.

Elliot did seem to be more thriving in high-pressure env. 

Other issues brought up - collaboration with US teams is a hit and miss. Some go very well, some not so much. API council. Performance team. Haven't been approaching them as partners - more like guardians / blockers. They didn't have enough bandwidth to work with us. Weren't as aware of urgency / importance of the project -- potentially something we can do about.

Timezone difference wrt US and Taiwan. Holidays they weren't aware of. 

Complaints about Android workflow. 

**FMD**

Good comments on the transition to Muhammad. They think that while it gonna take M some time to get to details of diff FMD components, they think he's doing a good job, learning and they feel good about transition. 

They had good feedback on the team and team work. 

They did say they think the reorg went well and they were happy to give away app responsibility, not upset about that. 

Two concerns:
- They are working hard towards launch in Jan. But after launch what will happen if Finder team decides to focus their energy mostly on low energy devices and less on phones, they'll be left in the situation they'll be responsible for all phones - a lot of work. If we can get clarity and buy-in from Finder team to ack that. 
- They spend relatively high amount of time to keep lights on / tech debt and issues that come up, and they have long list of outstanding improvements. Team size is relatively small and smaller by one because Anu HC was lost. Given the team size they think it will be challenging to do a lot of additional work. That's not great for the team if they mostly focus on keeping lights on. 

We need to understand what are the problems of FMD and alert BeTo guys. 

Roadmap / list of features/fixes for FMD.

Let Muhammad drive this?

**GPP**

A lot of compliments for Niko, his management, his leadership, team atmosphere. They feel good about it.

No major concerns from the team. 

They did ask what G thinks strategy-wise what the team should work on. 

GPP report?


## Richard Gaywoord.md on 2022-06-09

Been doing Telemetry

Over time taking over some of Pauline's portfolio for LON engineering teams. Will be LON PwG mobile contact. Enterprise, Telemetry, Wellbeing, AGSA.

## Dave.md on 2022-06-09

Angel OEM update
ASL team meeting
OneSafety meeting at xfn Sec Leads
ODAD maybe

Google v Android safety. Our thing need to include G accounts, reversing is painful. Gut tell me it's hurting us in the long run. We keep shipping our org chat. 

How can we present the data driven process for these things?

GRAD - will do team forum next week on setting expectations.


## Prabal.md on 2022-06-09

Ramped up nicely on .. feeling more involved, been 3.5 months. 

Helping Elliott with migration / system service. We are figuring out a way - when hit Refresh - when there's no work profile then we should not send signals for both profiles. Broadcasting to both profiles creates timeouts if there's no work profile / disabled.

Also working on how things look in Quiet mode. Earlier someone said we shouldn't show it at all - no we should but it should be greyed out.

Simon.

Team is really nice, feel like everyone is working a lot. 

Doing well: everyone is very helpful, whenever post something on the chat, someone always replies. Don't think I've ever been in the state when don't know something and have no one to ask. 

Can do better: I should seek out more what people are doing. Information sharing within the team. 

Having quick chats at lunch - very useful. 
Have a rough idea about the projects, when in sync ups. 



## Naveen.md on 2022-06-08

Okta, then Cloud Identity, Firebase Identification
Big change - change of strategy for identity, instead of requiring Google account, natively support external identity. 
Google distributed cloud hosting - private cloud.

Areas responsible:
- Malware initiatives in general
- Work across Meerkat


## Peter.md on 2022-06-07
Add Peter to review the Geist doc but not the meeting

Back to the regular routine

Focussing a bit less on individual IC work and more on leadership and trying to think ahead. Talking to Steve Kafka, William and other leadership to understand where their priorities are. 

Trying to learn a bit more about related technology. SimHash. 


## Niko.md on 2022-06-07
Things worried about right now.

Peter wrote a plan on his 70%. Had conversation. Trying to formalise what Peter wants to do. Niko's concern - he wants to go TL-y role, but of a smaller team.


## Giles.md on 2022-06-07

Giles to come to London w/o 25/06

Richard Gaywood to focus on Enterprise and Wear stuff

We want to have well defined Pixel bar for S&P and expose other OEMs that don't meet this bar

Idea, talking to Bennet. Washington Post - op-ed. PSL useless. No one reads privacy policies.

Even PSL - too much information, not actionable. Can we act on behalf of user, "Helpful Google" style. PSL machine readable, allow user to set some preferences on the practices they will never accept. 

Critical mass of opinionated features for Enhanced Protection opt-in

WWDC - something for abusive partner situation. 

Shouldn't runtime permission work? No, this is about access to location, not about sharing with 3P. But filtering out apps is probably infeasible. 


## Gabriel.md on 2022-06-07
Trying to get tickets to Austria but staying here for now

NExt week hopefully Austria

Wrapping up the dialogue work, strings, dimensions, etc

Next - Metrics. Will start owning this. 

Have weekly sync with Sunny, BA team. Trying to get product req from her side.

Plan is:

- Get the requirements from many different people, product, BA, me... - get a rough list of what we need to do. 
- Once we have that, sketch the bare minimum of what we can implement in the short term
- Prioritise the requirements so we can enhance the system in the future.

Def underestimated dialogues work.  Initial answer was weeks, because created PoC, took half a day, assumed it'll take short time. 

Complexity of the solution - extracting UI into a different module etc. Using new UI stack - Univision - on Phonesky. Added complexity. Took way longer than expecting.

Best / worst things -- been around for 8 months. People we work is unusually good people to work around, shouldn't take for granted. Everyone is very motivated and positive. Everyone is very technically sound and smart. 

Downsides -- feels like processes are slowing us down more than they should. Not unique for us. More layers than when it should be to go from point A to point B. Worked on S3 previously, but it feels we might be above Amazon in terms of process.

We have software component which is part of Phonesky. We need to go through multiple Phonesky processes, which might feel cumbersome. We need to adhere to all UI guidelines of Phonesky. 

Don't think current config is bad, but think certain things can be optimised. 

Niko - very positive feedback. Excellent manager and person. "Work friend".

Example: whenever I send him a message, he always answers very fast. His response time is O(1 hr) always and he goes an extra mile to reply fast and fully. Whenever I ask him something, his answers are always like if he tried to solve problem himself, but give me enough space to solve problem on my own - he's very deliberate about it.

Last time spoke a bit about commute. Two times instead of 3 times, haven't tried because enjoying the office a lot.


# Week 21
## Running_notes.md on 2022-05-26

Part of AiAi, scans incoming messages. Have on-device model which compares to phishing messages, we show warnings. Phishing / social engineering attack, etc

EU legislation about child safety. Expanding safety features, starting with phishing detection. More advanced features - grooming when adult tried to contact a kid, or explicit images. 

Enhanced Safety - let Hen know

## Guanhuan.md on 2022-05-26

Landscape of Angel beyond T - how updates will work, integration, etc etc


## Muhammad.md on 2022-05-26

Postmortem Done!

Spoke to Bridgett. Got on really well. She way saying about push from Laura about retrospective. She's going to work with M and PC folks to have some sort of thing in place going forward. 

M met Charmaine. She spoke about Shan driving this. 

Ronald wants to rename the whole team Find My Device.


## Dave.md on 2022-05-26

Angel postmortem
Samsung / GPP SC
Peter remote work arrangement


## Wa.md on 2022-05-25


## Guy.md on 2022-05-25
L3 hires

KS asked to look at the spreadsheet and asked if we want to use existing HC for any of the candidates, Stefan decided to take one. 

Yesterday 3PM KS said "I have to take 3 L3s from the list"


## Niko.md on 2022-05-25

Concerns:

- Caught by surprise
- Not quite what discussed with William
- Had chat with Guanhuan and not sure if he's on board
- Minor: if we start aggregating GPP alongside other solutions this could work against messaging that your device is secure by default, small incentive.
- Major: Samsung situation


## Running_notes.md on 2022-05-24
Best place to go is go/urpayment

Unrestricted gift - Money G is giving to Uni without any expectations Google is getting anything back. 

Maggie Johnson is VP of UR

Level 8 

q 14 - N/A

66001200 - non-U.S. Charitable Grant, Gift or Donation (unrestricted gift)


# Week 20
## Elliot.md on 2022-05-19

ex mgr on Fit moved to DUB, they lowered salary. Went to Stripe. 

Hedge funds did offer a lot more than Google.

Was getting a bit concerned didn't get much time to code past few weeks. Next week have summit. This week wrking from home to make more progress. 

Was talking to Muhammad, we need to clarify for 09 if it needs to be perfect. 

RK: will be shared with OEMs but not public

Remaining items - refresh status, the status of the loading bar (UI doesn't know when refresh is happening), persistent storage of dismissal, logging -- 3 big remaining items

Polishing, CTS testing, etc

Team - doing good. Prabal & Shriya ramped up pretty well. E is involved with Prabhal, Shriya - mentoring, but on tech front she's working with Max. Also Giulio, Karishma ramping down. 

Prabhal worked on a difficult CTS test, reached out to the right people, made progress with less guidance than needed before. 

Don't have visibility on the UI side - don't have visibility on what's they doing.


## Henry.md on 2022-05-19

Bridgette is good. Getting to know her, she's very experienced. Quite hands off.  Better fit than Laura. 

DaveK / Laura have aspiration to have across ASAP, bring cohesion - she's doing it across ASE. 

Will be doing Aleksandr Day as well, to ensure team has enough attention, planning for later in the year -- later in June.

Trying out some tracking stuff with FMD as well, better task load.

Eng postmortem will be complete shortly. Laura asked to do something more broad looking. Original plan was to engage with more senior exec types like DaveK - what was your part in that? 

RK: involve Krish, Narayan, Charmaine

Feedback during repro session - mainly frustration with not enough time. Less of a vision of what we are trying to build in terms of ASL footprint. At jnr level they think they just get thrown a lot of work. 

## RichardN.md on 2022-05-19

Windows installing in a month

Nikita applied to his job back, HC looking today

RomanU WhatsApp hacking

Lukasz found Triada modules supporting injection into WhatsApp


## Ishtiaq.md on 2022-05-19

Working on RAB

running different classifiers / heuristics before launching into production, joint between ODAD and Dynasaur team.

Dyn will be giving us access to the devices. We want to make sure that performance of newly developed models / heuristics is WAI. 

https://docs.google.com/document/d/19N1vszhEYJCU7taqJ3PnuPGnk3dEfytC84rTq0NSokU/edit?resourcekey=0-kQkWpsozSAuwdAmZEeI-RA

Will eventually introduce precision/recall to ensure model / heuristics acts properly. Will not flag benign apps, malignant apps only. 

Allows for potential experimentation.

Only working for ODAD so far. Classifier will run every 15 minutes. Eventually classification results will be stored on internal SD cards. 

RAB stands for Rapid Action Batallion

At the very beginning had a lot of dependency on other team members, didn't have enough understanding of Android / build system. 

Noticed that William gave Ishtiaq some tasks he started micromanaging Ish. Gave tasks said you have to do it by yourself, "this is the way you should do it". He gave me freedom to work independently now. 


## Nikolay.md on 2022-05-18

Working on stalkerware.  Submitting requests on more stalkerware subscriptions

SW WG achieved their goals - defined a lot of scanning scripts, but a lot of these things are off-market. 

Detection for existing samples - dumping everything the app is doing and trying to figure how to properly detect it.

bg/fg - not getting a lot of luck with this. Ask Hai to help?


## Aaron.md on 2022-05-17

Historical app ops part is mostly done. It's in datastore. Nikolay is now using it for his stalkerware heuristics. 

From A perspective, he'll be helping with stalkerware stuff. Rewrite scanning algorithm. 

1. We don't scan offmarket apps - for SW we want to scan offmarket 
2. We want to run more often - heuristic is pretty cheap compare to ML models

Have to start workmanager task for any type of work we do, we'll decide how many packages we will scan, retrieve some appops data. Heuristic itself is basically free. Then we write our results back to datastore.

Did a little work with Perfetto and Hawkeye. Plan is when Saagar joins (he has performance testing bg in Twitter) - Aaron wrote a system health vision doc, things he wanted to do but never had bandwidth. 

RK: OEMs will want system health numbers.

This stuff is not super hard to get through Perfetto. What we don't have is a really good automation / automatic finding regressions. 

Feedback to William: nothing much, interested to see now that we are getting more and more people, how he'll handle it. William is really good at it. When Nino came to the team one of the complaints about his old team that mgr had too many people and he didn't get much 1/1 time. Our team is very very small and William is doing quite a lot 1:1s. Once it gets past 10 it starts getting interesting. 


## Steve_Kafka.md on 2022-05-17

Two big new things:
- JIT
- [Project Honeybadger](https://docs.google.com/document/d/1vTFZnKDW8f57dLlxYK9axDuZYCw5k65MHhx99spaxlc/edit?resourcekey=0-Z05dQUeW7Em90ZOd27gCKQ)

Idea for that is that we've mostly been focussed on detection and throwing it over the wall. Apps are getting flagged from intel desk / escalated, we pat ourselves on the back - we found it, but Sameer is like - it is still live!

We need to take more ownership over e2e process. Let's get more aggressive.

IDV rollout - stages. Combining some of these ideas with devrisk is the right direction.

Raz - they've done pilots asking devs for more info and half of them abandoned their account. 

Be proactive, lockdown APIs that are causing us problems in the first place. 

Can we figure out APIs that are problematic and fence them off (full dev verification?) to reduce the risk?

We never will review same % of apps off-Play as Play. 

Enhanced Security mode on a per-OEM basis?

RK: Let's review most popular most dangerous off-Play apps.

SK: honeybadger is about risk-based approach to reviewing apps, and this can apply off-Play as well. 

Aggressiveness, risk management. 

We should say that JIT isn't enough. Only chasing detection won't get us everything we need. We want to get a little more radical for off-Play. 

RK: off-play limit risk

we consider all off-Play equal. Samsung store and random website are not equal.

stores can come in and register with us, and show xyz for keeping the store clean.

What volume do we need for 3P app stores?

app raising no flags whatsoever vs app raising some serious flags


## Yuri.md on 2022-05-16

Finished integration wtih Settings (Safety Sources), bugs fixed etc - this part is done!

Now working on UI with Tyler, implemented cards, rounded corners.

Postmortem session - super useful. Consensus on the problems. Same root cause - committed to a too large of a project. 

UI wise deadlines are manageable with reduced scope, if we cut most complex animations/design. 

Y doesn't feel like he can improve much in terms of technical debt / eng prod work.


## William.md on 2022-05-16


W described CUJ, agreed it doesn't make sense
AiAi is on board

> -   Good news, AiAi Privacy council will consider a one-time opt in to share ODAD verdicts with GPP:
> -   Lots of brainstorming of ideas that didn’t lead anywhere, so not going to note them here
> -   We landed on the following design I proposed:
> -   If user opts in, we send verdict through “framework API” to the “verifier”, which on GMS devices is only Google Play Protect
> -   I’m thinking of using Angel as the “framework API” - this API exists in T to report verdicts to Angel (that way we don’t need to build another API in Android U and wait more years):
> -   The Angel API/config file has a boolean about whether Angel can share this data outside of Angel
> -   If user opts in, ODAD sends verdict to Angel with share bit true (i.e. it can share it back to GPP)
> -   Action item:
> -   jgillula@ to come up with policy on one-time opt-in, which currently doesn’t exist - the fear is that every team will want a one-time opt-in
> -   jgillula@ thinks one-time opt-in makes sense for ODAD, and needs to come up with some rules to stop other features from doing this

Vish tried to bring some legal concerns again. His concern is he doesn't want ODAD to have any more signals. His concern is that other AVs will ask for it as well. Did follow up with Jessica, she said BS - we have rules, if they follow rules to the t it's fine but they won't follow the rules.

But what are the rules?

We can work this out with AiAi privacy council. Once we move plugins to Chronicle - once we move them there we don't even have access, as Astrea handles everything.


He wasn't overly enthusiastic about it. He brought up some good feedback:
1. How easy it is for someone to work around our detection?
2. What impact is this likely to have?

D not worried about security review. His concern is that we are tipping off malware authors.

Maybe he wasn't excited because we are moving too slow, and people are losing belief that we'll show impact.


Stalkerware - code complete on platform
Heuristic - getting close to code complete as well. Have some dependencies, some of them are running a little too close. 
Removing INTERNET permission - one of those. Someone is working on this with Hades team. This should be transparent to us. Hades team should use what Nir is implementing in Astrea. 

PRIMES integration for Astrea. They came up with a design with PRIMES TLM. They now need to present the design to AiAi privacy council, haven't done that. We don't have droidguard so we need hardware attestation.

Most of the team incl Eugenio expressed doubts about hardware attestation. For Pixel it's strong enough but for non-Pixels YMMV. Hardware attestation only verifies state at boot. 


## Khawaja.md on 2022-05-16

- Demo, demo!!
- Simon is hired, start 13/06
- Angel is in teamfood, have it on my test device

Got pulled in into jury duty. Mom got sick. Sister has brain haemorrhage. Convinced mum to go to New York, she catched COVID other there.

A lot of chatter on malware right now. 

HC coming soon. Supposed to get 6-7 more HC in Q3. 


## Nandini.md on 2022-05-16

- Announced at I/O and in Google Blog, only one sentence so far along the lines of "coming later this year"
- The plan is to launch via Mainline using November train. Will launch roughly at the same time as T-QPR, so they will be droidfooded together, starting approximately August
- Lowlights: OEM situation, Samsung. Solution: API in Android U

- New interest in securing off-Play. 
- Also AV-Test
- JIT

- ODAD v1 was about testing round trip, also caught some malware
- Launching stalkerware detection with Android T on Pixel
- Conversations about shipping to other OEMs, we depend on AiAi which will get GMS optional in T, and so (hopefully) will the ODAD
- Looking at content capture for ODAD v3

- Anurag left the team, now they all report to Muhammad who's also leading Angel team
- Working on BoqWeb migration of the front end
- Want to bring more work to LON


Samsung is also pushing on Nearby in Mainline.

Every team is making localised decision on Samsung

Other OEMs are way more ready to play ball, less work for them. 

Jen Chai would know about PCC on OEM. 

Kids 5 and 9

## Niko.md on 2022-05-16

Want to go MTV w/o July 11th


# Week 19
## Krish.md on 2022-05-13

Topics for discussion
----------------------

- No single product leadership for Angel. This results in many decisions being late, or not made at all. OEMs as one example - should we prioritised speaking to OEMs in September, we would have a chance to launch with Samsung. PbA branding is another one - right now we can't even decide whether we want to translate it. 
- Conversations with OEM are handled predominantly by privacy folks. As a result, certain things slip through the cracks - as one example, at one moment in time no one could answer a question of whether there will be GPP in Samsung's implementation of Safety Centre - Samsung provided mocks and PMs in MTV have largely blessed it, even though there was no mention of GPP (it was replaced with McAffee). We had to bring this up (luckily, we figured that little known Verify Apps API actually gives most of the information OEMs will need, suboptimal as it is.)

Notes
------

Giving an offer to 7.5 who waits for L8 promo. For PMs theres a mapping process, 7.5 will get mapped to 8 in fall. LON candidate commitee said hire for 7, NH for 8. Talking to them today. Going to ask them if they will consider L7 offer. If yes talk to me. If he says no then we'll continue looking. One more candidate, based off U.S. but has been working with Israeli team so his hours are EU hours. 

7.5 - will get date today. He's giving his notice today. 

https://www.linkedin.com/in/luke-abrams/?original_referer=https%3A%2F%2Fwww.google.com%2F

https://www.linkedin.com/in/zach-howes-b20366/

Krish to talk to Charmaine. GPP is one example. Guanhuan forwarded thread, in T we are not making it mandatory for other OEMs. 

We are asking Tina to move to Charmains team and take over Sara's responsibilities. 

Apple is drumming the beat off-store is bad. We need to do the right thing. There will be additional work for GPP, how we sample off market. 

The goal. Asking team uplevel malware strategy. It was always about detection, each year "we detected more". Asking to do 3 things:
1. Our APIs surface area. The biggest reason we have more malware is because we have more open APIs. Value prop we have for users, but also reason why it's incredibly hard (speak to [[Ed]]). How do I protect these high value things better as opposed to permission by permission whack-a-mole
2. Think about not just malware through Play, but also holistic approach about preloads, offmarket and Play. We have some of these things, but come up with cohesive Android story and not just Play story. 
3. Lot of malware is not what we think is malware (e.g. stealing your money). Lot of malware is a low quality mw, broken functionality, apps trying to do only ads, web wrappers - low quality stuff. Want malware team to start working with quality effort we spin up with Play and AppSafety. Go and remove all these things that were never meant to be on Play to begin with. Helps us to focus on the things that matter. Start aggressively removing things off users devices. Adding friction every time you use bad/low quality apps. 

Start thinking what role GPP can play in that - not only as GPP client owner, but GPP-on-steroids. GPP as a security product.

Need to come up with key metrics. 

Based on what K see, he feels these things will be important. Maybe for sideloading we can say we are as effective on sideloading as we are on Play (which we are not). Put friction for all apps we removed from Play. 

## Sebastian.md on 2022-05-13

Comes to London in July

WhatsApp. Lukasz and Roman found new Triada modules that target WA and FB. Sebastian found Triada module targets GMail. C2 registered in 2012. 


## Alan.md on 2022-05-12

Isolated compilation - mostly done, some bugs
Heavily into U planning
Had mtg with partner team - most of the pKVM team there. Worked really well, more junior people didn't contribute that much but got the context, also teams got to know each other. 

## Guy.md on 2022-05-12

- Dogfood status - looks like there's way forward
- Simon joining mid-June, unfortunately week after Guy comes over and when I'm in  MTV
- Product situation. Not malicious, just very MTV-centric - which also shows up in the vision of the product. 

- JIT protection - converging on UX, the key will be what sort of consent do we need from the user. Need to re-think consent also in light of ODAD verdict sharing. Ideally we have to push for single Protect++ consent which will take care of many of these things.

- Stalkerware for T seems to be in good shape
- Phishing demo is really cool, and so is the toll fraud.
- AiAi team is supportive of sharing verdict with GPP provided there's some consent from the user. Asking UX to work with Leandro on figuring this out.

- Continuing to work on BoqWeb migration, everything is according to Plan
- Muhammad will start diving deeper into FMD as soon as Simon ramps up

Guy sent intro email for Muhammad to Ronald and Erik Kay. Both were nice. Muhammad will try to meet them in person. 


## Rodrigo.md on 2022-05-12

still a lot of confusion bout the brand
the shield icon is pretty generic, wouldn't be an issue
as long as we are not showing PbA branding

No clear product strategy yet. 

In future iterations everyone have to be more involved

blog.google on SC

## Prabal.md on 2022-05-12

When had last 1:1, didn't have too much context. After 1 month feel have a lot more context, feel a lot more confident. Learned more in the last month than the previous 2. 

Completed Work Policy Info.

Speak to M every week. 

Hindi / Urdu. Can have 1:1s in native language, mostly when speaking about non-directly-work related things.

Having same background helps when connecting with someone. 

Usually tell about what he's done in the 1:1s. Last week asked to complete / wrap up project. 

Will be helping Elliot with migration. Also will help writing some CTS.

Elliot very quick to respond and will make time for you. 


## Leandro.md on 2022-05-12

North Star Metrics -- still refining the document. 

Leandro to work with Guanhuan and Rodrigo on Protect++ approach.

Shall we notify only off-Play for beginning?

L pushing team to sandbox testing - being able to measure quality of heuristic, TP/FP.  L wants to have a place where we can test this heuristics which will give us a piece of mind. Also helps with system health footprint. 

We can't put real malware to dynamic analysis, we don't have proxy etc etc. They are migrating to Cuttlefish.  Why not use Dynasaur for this?

Do we have external company we can pay?  Speak to intel desk.

In v1 - near impossible, have very high number of FPs. 

 

## Krish.md on 2022-05-11

GPP v Angel - what happens now?
PbA
Hiring


## P&E_leads.md on 2022-05-11

Tangor launch June 2023

## Guanhuan.md on 2022-05-11

We need to plug into OEM conversations, maybe even lead these conversations completely. 

Guanhuan was in call with Samsung Tuesday 1am. 

Guanhuan spoke to Krish about it. Krish was surprised to find out we are in this situation (Pixel only).

Did they offer TC? They didn't accept new round of offer yet.

GPP <-> Angel: Krish thinks that Angel should subsume GPP.


## Shan.md on 2022-05-11


Tone in the meetings is distrustful
Us vs Them attitude



- Situation with OEMs
- Situation with PbA - or is it really? 

Real concern: decisions are made in MTV, and communicated down to LON team. Should I really tell that? Probably not. 

When Eng discussions going into, Eng should get involved.
Many other discussions not about tech, still aligning on some concepts

## William.md on 2022-05-10

Who is our new PWG person for ODAD?

Nino is checking in the demo, Rahul to schedule something in a few weeks?  Use `scrcpy` dammit!

Ishtiaq: CL he sent out, Aaron reviewing it, makes no sense what he's doing

RAB = ODAD version for Samurai


## Muhammad.md on 2022-05-10

- Dogfooding - see email from Mark Harpin
- Simon - what's the start date?

Which Mainline release will TM-QPR1 release with?

Simon - 13/06


# Week 18
## Running_notes.md on 2022-05-06

UX Milestone 1 - what's it about?
Get progress report next sync


- Are we saying that PbA is Pixel first, only or what?
- What is our messaging about GPP brand? PbA shield looks an awful lot like GPP shield
- No PbA branding on permission dialogue


# Week 17
## Jim.md on 2022-04-25

Who is doing Spot integration with the BoqWeb? Jim met Rashmi, TPM for Finder. 

David L he has a new starter on his team, Jim bringing them up to speed, working
on replacement for location history. He's not working on Spot integration for
Web UI. 


# Week 16
## Henry.md on 2022-04-21

New manager - Bridgette, 


## Muhammad.md on 2022-04-21

- Simon
- Angel syncs - move to offline?
- Milestones


## Sham.md on 2022-04-20

Review with Sameer last Thursday. Creating churn on the team. Sameer is asking
why do we need an opt-in, how can we achieve scale?

Takes us back to the drawing board with PWG/Legal. 

Eng of Aug / eaerly Sep launchn date.

New PM reporting to Ronald, all integration related features. What are we doing
with Chromebooks, Wear, 1P uwb, Eos - new Wear/Fitbit watch release I/O 23.

grogu - 1P UWB tag

Sham to focus on phone and v1 launch

dogfood starting soon. Dogfooding phone features and chippolo tag and Sony
earbuds.

Mike wants to work on enterprenerial projects, change of pace. 

## Khawaja.md on 2022-04-20

Write up mgmt charter Muhammad/Simon come calbiration time to understand if
there's space for 2 EMs


# Week 15
## William.md on 2022-04-14

Aaron is complete owner of ODAD app. He syncs with Nikolay, who doesn't know the
app to help him get his work done. Sending Aaron off to all kinds of meetings W
wants to be in. Trying to focus myself on early-stage meetings that are very
important.

Also responsible for integrations with other teams, Hades, Astrea plugin. System
health of the app.

Ishtiaq has been doing work for Aaron. Pulled out because want him to work on
Samurai project. Gaurav is coming. W wants Wa to be mentor for Gaurav. G will a
bunch of different things. Wa is responsible for everything happening cloud
side.

Nino is working on demo.

Nikolay: Leandro gave him EE for leadership. Participating a lot, coming up with
designs, helping Ishtiaq a little. He's in charge of stalkerware.

Nino - his manager says he's awesome engineer, what is amazing, how's not so
good. On his old team he'd jump in and criticise design from more senior people,
even when he was right, it was very rough. 

Starter project - using content capture, detect GriftHorse phishing.


## Prabal.md on 2022-04-14

Was surprised, thought the team was working more on security, but we are SWE.
Was assuming there will be ISE in the team.

Elliott is the mentor. Very responsive. Every week talk to Muhammad. 

Working on work policy info. 

Taking longer than expected, don't know the code properly. Something as simple
as importing the class takes time.

Like the most on project: had to read the code, all the entries. Had to go
through all of them, after doing it understand it a lot better. 

Graduated last year, pretty much first job.

People are sometimes working long hours. 

Speak about manager next time.


## Nino.md on 2022-04-14

Android Auto connectivity




## Barak.md on 2022-04-14

Transition. Moving 3 engineers and some assets are moving to the new team. GaNT
team. 

Smallest group from security centre moving with the assets themselves, account
notifier assets, iReach, push notifications on iOS. The rest of the team people
are moving to GanT but assets going elsewhere.

Hiring in Brazil - about 6 months. Interim - sunnyvale, Diana Smetters' team.

Looking to make a lot of progress this week but Guanhuan covid.

Talk a lot about various exceptions. 


## Isabelle.md on 2022-04-13

Aiming to have dogfood ready boqweb end of may.

https://find-my-device-autopush.corp.google.com/

Launch some time in August.


Hybrid, flexibility!

B


## Muhammad.md on 2022-04-12

shallow depth of understanding


## Dave.md on 2022-04-12

SecEngMgr -> SWE - not Dave's rule

GPP / Angel - why should we force tech integration? So many things Angel has to
integrate which are not part of Angel. 

https://moma.corp.google.com/person/menakashroff@google.com



# Week 14
## Krish.md on 2022-04-08

- We have a green light for September hitting code complete for December launch. 
- Suzanne will be asking us for post mortem, need to make it happen. 
- We will be asked to escalate faster if something's wrong.
- Need regular milestones and hygene

- Branding is figuring out what's included into I/O. SC is part of I/O. 
- There might be asks for mocks, etc. 
- Need to support I/O.

- We will go back to Samsung with an eye towards U. If they launch rogue
  experience, let them - we are not going to take this fight. 
- We asked Sara / Guanhuan to write down what are Samsung's requirements. They
  want customised UX - non-negotiable. Non-negotiable for us - not drop GPP.
- They will ask to include things into data layer. 
- We gotta write down what are negotiables and what are not. We have 18 months. 

Some of our decisions are binary. But everything with Samsung is a negotiation. 

Need help to keep team focussed on these 3. 

No UX - not the end of the world. 

Segmentation - OK for U. 

Finished all 3 L8 candidates. LON candidate is still #1. Krish out next week.
Meeting after to decide who to extend offer. Still hiring James' replacement.


## Max.md on 2022-04-06

M sometimes put people too much on the spot. Asking for volounteers. Maybe it
would be more appropriate to do this in a 1:1 setting? Sometimes people feel
they can't say no. 

What if Max goes ill - what Phonesky?

# Week 13
## Nandini.md on 2022-03-30

Re-org moved to Tristan. New charter - working with Asisstant, stalling
progress. 


## David_Coffin.md on 2022-03-30

Uncomfortabnel with the way working with Muhammad. He questions through me a
lot, rather than talking to team members individually. Feels like he wants David
to be a manager for a team, doing it with Elliot as well. Can you please ask so
and so do it, and ask so and so for status update?

I just don't have enough time to do all that. Average 13 hr meetings a week. 

a: you either TL or not


# Week 12
## Fit_Calls.md on 2022-03-23

PhD Paris, distriobuted alg. Assistant professor. 
Few years wanted to switch to industry.
Research is amazing and exciting. 
But academic jobs is really diffrerent from our dreams. Solitary work. WHen you
are at uni which is heavily teaching oriented after couple years you feel not
advancing even tho you have impact - love students etc. You feel can do other
things as well, building stuff!

Growth opportunities scarce. 

Dream job: my first wish would be to learn from the technical side and practical
side. Haven't had a job at software company, never worked real life project.
Have very good theoretical knowledge, but only familiar with coding at toy
projects.

Really looking forward to learn more about real life projects and how to work on
them. 

Motivated with tch challanges.

Societe Generale in HGK, fresh out of school, didn't know what to do. Was
contacted by lots of recruiters, especially finance. Gig in HKG, because why
not? Not really liked very much my experience there. Enjoyed the experience but
did not see myself continuing this line of work. Tech challenge was there but it
lacked sense of purpose.

Joining as junior engineer: I know that I'll have downgrade in seniority and I'm
not very concerned about it since the way I see it is to take few steps back in
order to advance faster. And I know I'll be advancing that I'll be advancing a
lot faster than many other juniors, just because I learnt a lot about work in
general, working with people, etc etc - this will help me to prgress quickly.
Hoping not to start from scarch, from zero.

Her questions:
Have discussed Niko/Bessie, they answered lots of questions about the team. 
- THis team wants on GPP from cleint r


## Nikolay.md on 2022-03-21

have app running, dashboard. Positive change. 

Westworld - although takes a lot of time, the process isnt' super
time-consuming. Have to see how we do things in paralell. Long time to do
something.

Not platform release required.

If one use counter they will approve it. If you want to send strings it's more
difficult.

Mostly working on stalkerware.


## Niko.md on 2022-03-21

Hiring: doing fit calls. 
**



## Aaron.md on 2022-03-21

Historical App Ops work. Hopefully later today planning to meet Nikolay.
Responsible for extracting the historical app ops as a feature. He's gonna work
on the heuristics. Discussing prototype of the feature.

historical - 15 min - hour -- in memory, after than - disk.
discreete - also keeps in memory for some time.

Understand people on mgmt chain to have good idea on what we are working on. My
only concern is - giving estimates.


# Week 11
## Luke.md on 2022-03-16

Very different to old team. Moved from Android Studio. Organised very
differently. One of the main reasons moved - learn new things, been there for 3
years. Niko seems like a very good manager, welcoming, helpful. Team in general
is very nice - one of the reasons. 

They don't do iterations / Kanban boards. There's a lot less process when it
comes to releasing things, they only release 2-3 times a year. Despite that
there's a lot more craziness when it goes to releasing stuff, often have lots of
bugs coming in and fixes only availble next release... That's negative aspect.
Positive is things done very quick and people have a lot of autonomy and
continue working on project they think important. Downside - not much
preparations for OKRs, usually happens later and not very detailed. Android
studio's demograpohics is developers - people who can use workarounds and won't
be bother by that.

Internship - Dublin, SRE team, step intern. Monitoring tool before IRM. 
Second - MTV, AdWords platform team. Mid transition from very-very old ad
platform and whole new Material design one, worked on report downloading
Third - NYC, ads platforms team as well, isolated team that created experiments
framework, APEX (like Mendel, but ads-specific). Helped them to find what to do
with bots on the platform. 

DUB - not a lot of SWEs. MTV - fine, very nice, sunny, but don't really like
area (lived in Sunnyvale). NYC - favorite outside of Europe, milling things to
do. Now they bought Chelsea Market, and one more building to the pier. LON is
more mature as an office, but smaller. 

2 starter bugs - in Android R new API to pick out installer package name, and
it's sent in a big report of installed packages. Previously installer could
reset it to anything else. Second - there are incremental apps released in T,
instead of scanning them as they are being downloaded, if it's incremental app,
schedule autoscan next day.


## Anurag.md on 2022-03-16

Looking at Rubidium opportunity. Nandana. Sorin TL for FMD, let team continue
working on projects till Q3. At that time can start pulling majority of team
into 1S. q


## Muhammad.md on 2022-03-15

Settings exception done
A11Y is with permission controller, they are implementing it inside permission
controller module itself so it's not subject to deadline. Will roll out as part
of Mainline release.

Concern - need extra help with getting APIs nailed for deadline, incorporating
feedback provided by API council. Main concern is about ODAD - it will
potentially impact XML config, but effectively API council might consider this
new feature. Conversations are happening this week, working with William & Vish
on this.

Rogue tags: not letting the team lose focus. P2 feature for Muhammad. We can
even do it for T timeline, because code is in GmsCore. However, Guanhuan needs
to get a bit of discipline on how he interacts / does stuff. 

Team: Anurag is settled and wants to move. Haven't got a TLM. 


## Krish.md on 2022-03-14

PM lead
Possible frictions with Safety Centre PMs
Mention update on FtM and how it influences AV Test

3 L8 candidates, one Googler, 2 ext, 1 based in London. 


## Angel.md on 2022-03-14

Team is not in a really good situation. Went through reorg. Jenny is departing.
Currently fairly insecure situation. 

Discussing with Guy additional resources. Guy is very supportive. Have internal
candidate, senior, inclined to join in early April. 50% malware, 50% GPP+Angel.
Will revisit this plan with Guy and keep updated.

# Week 10
## Tyler.md on 2022-03-09

Excited about UI stuff. Getting good feedback, breakdown & task mgmt stuff.
Exited to get into details and start implementing stuff with Yuri. New person
TBD. 

Managing changes in Android is still PITA. Worried about work/life balance when
in the UK? See my co-workers staying online pretty late, sometimes beyond even
Tyler's working hours.


## Giles.md on 2022-03-09

ODAD mostly?

WebView - a year ago realised that WV is sending package name in their header by
default. Something been working for several months to stop. 

Can OEM replace ODAD?

Phishing: getting unphishable auth widely used is the most impactful way.


## Khawaja.md on 2022-03-09

- ODAD headcount
- ODAD - stalkerware is all clear, we have all we need
- ODAD - talking to AiAi, will be able to access the screen and everything in
  it. This may allow us in-app impersonation and phishing
- ODAD - conversation with WebView team, injection into non-developer owned
  websites

- Angel - latest Muhammad's update summarised it well.

<!-- - GPP - csc@ asked to transfer bessiej@ with HC, can't allow that right now.
  New starter this week (HC on loan from William), but need more because
  expanded remit. -->

- FMD - Anurag is looking to leave, still will go ahead with re-orging the team
  under Muhammad and starting to support Angel.

- Googlegeist


Anurag - can he be used in Wei's team? Is there a role? They do a lot of infra work.

9Q2, 9Q3, 9q4


## Anurag.md on 2022-03-08

Concerns MikeT: Anurag spoke to David if there any any concerns. David's
understanding of where we are is exact same as Anurag's. Explained where we are,
timelines, and where we'll need help. D said a perosn coming on their team who'd
start helping with BoqWeb next month. Long pole - identity requirements. He is
somewhat annoyed by new PgM - Rashmi, new to Google, constant requirement of
status updates is bothering him, he's escalating this to his manager. He said
that concerns are possibly because of misunderstanding b/w Wei and Mike on who's
doing what on BoqWeb project, if Spot is involved and in what capacity, etc. A
to have 1:1 with Wei. 

Are we still on schedule? We should be, schedule has been communicated to Finder
team, they have visibility into this process.

Bigger thing: there will be no I/O announcement. The launch date is some time in
August, b/c if we announce at I/O and there are delays that won't look good.
Also because Apple released app on Android to detect rogue tags, possible that
Android privacy team will force Finder to do the same, to release iOS app that
can detect rogue Android tags. It takes a _long_ time to get approval from
policy / Apple / everything. If this ends up a requirement, no way we can do
this before August. 

Primary concern remains the same: problem very much like FMD, and I don't feel
too excited about that problem or FMD. I'm still not clear on what I want to do,
but I'll start looking for opportinities. Spoke to immigration lawyers,
relocation looks really hard, not happening in short term. 


# Week 9
## Sebastian.md on 2022-03-04

What are we thinking of doing in ODAD, and which categories of malware it can
make the most difference?

* Historical app ops - already happening in  T
* Network traffic stats maybe?
* Access same things as AiAi does including content on the screen

Stalkerware is #1 goal on Pixel right now.

What is hard?

* Cryptocurrency fake wallets - in the malware exec review that will come out
  next week, 11K fake wallets in Google Play that steal cryptocurrency. They are
  flooding Google Play every day; just tracking which crypto exists in the first
  place is hard. Which are official apps? Which are unofficial legitimate apps?
  Which are fakes? We have to review every single app in crypto space :(

* WebView based phishing. What happens is that very often legitimate websites
  are loaded by malicious apps (e.g. facebook.com), and then JavaScript is
  injected into WebView. And that even of JS injection is extremely interesting.
  If we had some support from WebView team, they can make it available as a
  single for us. Just trying to understand all the applications that try to
  inject legitimate websites with their JS? 

Is this something we can detect cloud side? In a perfect world, you can detect
everything cloud side. But in real world, I would reframe this as "on the cloud
side we have to work very hard to defeat anti-analysis techniques, but on client
side we do not".

In general, hope to move more and more signals client side.

Digital asset links?

More fine-grained signals? Phishing in particular, e.g. injected JS has accessed
password field. Any application that reads password field is interesting. 

Get a hash of the javascript?

* Stalkerware is a very good one. If you can't detect stalkerware, you can't
  detect anything at all, this category is the noisiest. But what do you do once
  you detect it? 

Malware is not very text based. Wouldn't know what to get out of text-based
analysis. Wondering if you have more luck with image-based analysis (like SB
team uses icon detection to understand if phishing page or not). Ransomware
really comes to mind - if something asks you for bitcoin, doesn't look good. 

Person to work on this: roman or nikita.

* Tollfraud - still the biggest pain in Google Play. There are malware families
  which are social eng based, but then there are apps which use notifications +
  JS to paste into websites. Both need very different approaches. Codes from SMS
  being entered in app - very interesting signal, no way around it. Does the
  Android operating system / security model allows for such very specific
  observation? Also helps with phishing applications.

* Proxy malware - unclear if there's a way to use metadata available to Android
  OS. Talk to Nikita?

KFC/ODAD - not the first time I hear about that. Very interesting, having the
ability to essentially query the behavioural properties of apps across 2B
devices is interesting. Have to add a lot of safeguards not to DDoS. Also what
signals might we use? If we go down that path, we need to work closely with
reverse engineers, let's look closely to 10-20 malware families, how would we
discover these apps in the wild? Lukasz might be the right contact person for
this. 


## Guy.md on 2022-03-03

Notes optional, but encouraged for more difficult cases, or at least a sentence
or two for candidates you'd like to discuss.

How to make ODAD more marketable, sell it better? 
We don't say are we gonna detect it on the device, or send it to the server?


## Fit_Calls.md on 2022-03-01

https://hiring.corp.google.com/hiring/hms#/hms/reviewcandidate/a1e9c4ef-5efa-407f-9a37-079e5182c210/candidates/google.com-378526541/applications/453021090

William spoke about multiple levels of malware detection, codebase, just
launched v1, have long term vision, thinking about remote. 

Makes sense why we building detection without privileged access. 

Browsers - are we talking about malware sneaking through the browser

SWE at Amazon. Firehose team - data streaming platform at AWS. Before sso team
at AWS. Worked at vairous companies, done full stack and Android app development
- joined startup and rewrote their app from ground up. 

Master in ML. 

Ideal role - right now (5-6 years of experience), want to work on skillset
missing for L5. Less experienced with system design. Looking for projects that
could help me to get to a point of a senior engineer. 

Within 5 years - not really sure. Soft skills, presentation skills - considered
mgmt role but definitely not ready to make the leap. Maybe executing through
others in 5 years.

Soft skills - debates at school, love presenting, can do presentations and
workshops all day. Noticed on my team I'm happy with my tech ability but what
makes me stand out is my docs, building consensus etc. Caught up in performance.


## Narayan.md on 2022-02-28

Bad bad connectivity booooo

Things are going quite well. Rather complicated project, execution has been
pretty good, team pulling together well. Had some bumps - odd comms with VPs.
Muhammad is doing a good job in getting started out. He is know understanding
how Android works, release process. Impressed so far, Erik likes working with
him too, smooth collab eng wise. Might be a few hiccups ahead with partner
feedback. 


# Week 8
## Yuri.md on 2022-02-23

Safety Sources in Settings framework.


# Week 6
## Inara.md on 2022-02-09

Unknown installers - rephrased, removed safety concerns, trying to get it in

Latency issues on Phonesky start up

Waiting for metrics, should have more metrics by next week. Testing shows it's
better.

100% mid-end Feb

Group MuWS notification to launch. Changes re restoring notifications. At 1%,
but early this week realised we show 1M -> 30M MuWS notifications a day, big
increase, rethought a bit, going to descope this restoring notif changes from
this feature.

A lot of scope for improvement in verifier. Not just improving the health of our
codebase, but also lots of small wins. We still only warn against 1 type of MuWS
at autoscan, should warn all. ODML. 


## Khawaja.md on 2022-02-09

What is he trying to achieve with reliability numbers? What is the deeper
question he tries to get answered? 

FMD -> Infra (Angel)

There are certain things that help me see how reliable service is. What's
uptime, latency distribution, throughput, error rate? Trying to understand -
what's missing in our ability to measure it? 

Improving 99th percentile is a difficult problem. That's what you use in getting
HC and making your point across - by computing how many people might be dropping
out.

First, engineering metrics. Then, product numbers. If we see odd thigns in
product numbers we now might have an answer why we see this - e.g. service is
not 99% available.


Roman: no end-to-end latency, it's a gap. KS: it's OK to have missing pieces, I
just want to have them all in one place.


## Tyler.md on 2022-02-09

Re-doing UI and checking in for prototype.

Very basic UI.

Muhammad: weekly 1:1s. Feel that I get positive feedback from him. No
constructive feedback just yet. 

Visa March 28, relo around end of March.


## Guy.md on 2022-02-09

Same 1:1 freq with Anurag
Add a new person or HC
Monitor how Muhammad treats him wrt perf etc - protect him
Another L7 strong supporter. His interest aligned with your success


## Running_notes.md on 2022-02-09

MikeT is leaving Google.
Wei to work with me to sort it out, have agreement on things (?)

Collab between FMD and David's team.

BT got fraction of the HC. Goes to very urgent projects. Could get more HC, next
batch. Can try to secure 1 HC.

https://docs.google.com/document/d/1q6IWlGMP63c7u6sC1bo-yH1Qf3qeXfb7aymaSCVPT8A/edit#

App UX

Everything that has to be rapidly changed based on product requirements should
be in MTV. Security/privacy/infra in LON. David might have more clear answer on
this.

Next HC batch - time unknown.


## Jan.md on 2022-02-09

Works on GmsCore integration. Demo works fine for FMD. Next step - pushing
Ariane for approvals. Preparing implementation in the background. Impl can only
be submitted once Ariane approved.

Added gOTA source to GmsCore. Non-gOTA OEMs will need to add their own static
source. We will read system property when was the OTA last applied. If OEMs are
using same system property for their OTAs it will also work for them. 

There exists an API which no one uses. Connecting to this API is more involved
than parity, copying what we have. Let's track this.

Settings integration is done, Yuri submitted his last CLs.

Muhammad: talking what Jan's doing, he read GmsCore design doc. Started talking
about perf. Spoke to M about L5, probably autumn cycle is too early, but next
spring cycle.

## Narayan.md on 2022-02-08

Two options - Android and Mainline

No difference from the user perspective. T devices not reaching user till end
October. From the user impact - no real difference. 


## Fit_Calls.md on 2022-02-08

Joined Dialer when 3 engineers. Now 30+. Back then doing PM'ing, UX'ing,
PlayStore'ing and everything. Miss that part - when team is small and you do
everything. As part of call screen - first unbundled feature. Unbundling process
was complex. How not to be dependent on Android release cycle? This has never
been done before - logistics, but also technical limitations. 

Prototyping this was amazing. All smart audio features live in Assistant. Love
process of working with lots of different teams, coming up with hacks to work
around technical limitations, onboarding to OEMs. Real nice thing I liked was
challenging problem - not solved as such, technically possible but how? Going
to Seung on saturday to LGTM CL! Having to convince different people, different
stakeholders something Usman really enjoyed. Now the scope is much more narrow
because there are PMs, PgMs etc etc etc.

Natural trajectory - expanding impact and scope; convincing through data and any
means possible, getting people onboard. Really enjoyed interfacing with
different teams. That is the part I missed, but really want to re-live this
experience.

L5. 

Questions:
- Very long term would love to go to VP. Or Fellow? No, VP. reason for that -
  interacted with a lot of VPs, but not a lot of fellows. 
- What can I do? If I join the team and hit ground running, what would I be
  doing?


## Leandro.md on 2022-02-07

Integration with Angel. Will share first version of the document. 
We have an option - unified interface. 


## Anita.md on 2022-02-07

Career conversations - once a quarter.

Weekly 1:1s. He pops up in stand ups / team meetings. Helpful when issue he's
expert in. Some monarch configuration / gmon alerting config, he's been very
helpful.

Mgr: stepped up well and he's very good at listening. I feel like am able to
talk to him honestly about my concerns about the team. 

Finder: good grass-roots effort, collaboration. Still don't have a sense of
project direction / who's the PM. We have set up weekly mtgs with eng team, we
can bring up what's happening in terms of releases, app, how it affect what we
are doing.

FMD team is not strictly hands-off, did some pre-work to get rid of 2SV, which
would require extra login when there's no lock screen set. We no longer monitor
FMD app buganizer component, goes all to them. Lollipop bug was our fault
becuase September.

Spot is doing big project to finish re-architecture app, doing it with Sorin's
help.

We could be better with overall project direction - need better overview of
where we are going. 

2 more TVCs - they have desks on 7th floor.


## Niko.md on 2022-02-07

Questions for kshams

Pottery class. 

Consent default on. Normally for really bad PHA we'd autodisable them. There
were a bug introduced where we'd automatically remove them.


# Week 5
## Monir.md on 2022-02-04

Mubdi 9 months old on 3rd Feb

## Ahmed_Radwan.md on 2022-02-04

Just joined. Come from Apple, spent last 8 years. Before AAPL finished grad in
the US, PhD Distributed Computing / data mgmt. IBM.

https://www.linkedin.com/in/ahmedradwan/

AAPL - developer connections, clusters of developers. 

Focus - risk management. 

Apple is trying to push a lot of stuff on device, because of privacy et al. 

3 kids. 17 (girl) applying for colleges, art. Really talented artist. Applying
for colleges with strong art programme. Meaningful art. 

boy 12 y.o. jiu-jitsu since 4 year old. 

girl 11 y.o. gymnastics.
 saad

## Max.md on 2022-02-02

Fourth week. 
Doing GTI.

Starter project - safety source, charged with trying to implement mainline
safety source, Mainline Updates. Phonesky stack. 

Worked for Guardian. Office is near KGX as well. A lot of in-house software,
tools for producing content etc. Was on Android team there. 

A lot of cultural things. Ownership of work / features by developers. 

Team: everyone is really nice. Feel a little it's a difficult time to be
joining, deadline in Feb. Onboarding people it's work! 

Muhammad: for me as a new starter him having just recently joining is a
double-edged sword. On one hand he went through similar experience recently.
OTOH he doesn't have a lot of experience at the company in general. His
communication is good on one hand, I like the way he speaks relatively plainly,
he and I can talk honestly with one another.

Sometimes think there are maybe we could try to be a bit more - I have more
things I need to speak to him about, Need more time with him. What Muhammad's
remit? Why am I reporting to him? What is his objective? What's his relationship
with other senior developers in Angel team? How does his role and David Coffin's
role is different? 


## Peter.md on 2022-02-02

Documenting verifier, production playbook. 
Made a lot of progress on understanding verifier features. Have an idea of what
everything does and why it exists.

Niko's working on metrics trying to come up with metrics actually measuring
product and reliable. We have tons which are outdated. Tries to simplify.

Auto-reset - we had 10%. We are making Phonesky start up slower. We think we
know why but don't have fix just yet. Have another experiment which just
launched, and they are conflicting with each other. Submitted some fixes to
improve the situation.

Other exp - security status API, reads security information from several sources
and passes it to GPP Home / My Apps. To do the auto-disable-and-update they
needed to load this data and cache it on Phonesky startup. By launching
permission revocation we are adding more data. We are not blocking launch, it's
in background. We don't need to load permission revocation data on start up.
Having chat with Niko today about options we have. We made data load more
efficient, submitted fix, should be faster already.


## Dave.md on 2022-02-01

Angel update
 - https://drive.google.com/file/d/16ymalZkyiAoGz6Vkz4G0bltRAeIBGaKI/view
Acct Security
GPP update
ODAD

- Verified that we can receive data now, back to 0.95 threshold
- Telemetry


Warnings on app installers card in Angel

GPP

2022 focus on malware. Malware Agility!!

We can't detect things at scale, and won't be, but we can react faster to
protect users. If there's big campaign that we can prevent and can never prevent
we can act faster to protect user.

Wondering how we manage a definition of green/yellow/red state over time. 

We got the number we asked for. Bad news: we didn't ask for everything
EVERYTHING. We had to cut it arbitrarily. Now that we got heads we need to go
back and decide who gets what. 


## David_Coffin.md on 2022-02-01

Jacob plays piano

Detached last week. 
H
API stuff coming together. Finding people are reopening design decisions we've
made some time ago, takes time.


## Gabriel.md on 2022-01-31

GPP dialogue, started implementation. 

Peter taking lead on understanding features we lack knowledge
Once he's away Gabriel will start working on PV with Bessie

Niko still good. Very available. Feel like I gained a bit more independence. 

Career convo - he's been very open to it. Was L5 at Amazon, ~L4 here. Midway
through getting promoted. 

G documented where he wants to be in a few years, specific spreadsheet. 


# Week 4
## Karishma.md on 2022-01-27

Stuck on API review. 
Making long chain of changes is difficult
We are writing CTS tests. Elliot is working on e2e prototype.



## Elliot.md on 2022-01-27

Concerned about making changes to API after 15/02

Worked closely with Giulio (XML config)


Biggest risk - not making to deadline

Muhammad - really good. Trying to enable the team and he's really good as a
manager on a people level. 

One thing to change - amount of meetings. He seems to be quite process heavy.
Some of the meetings are necessary. Might be better if we can defer it to
Henry/Guanhuan? 

A lot of process on Marie/Elliot - Elliot suggested when we do daily sync
updates, take notes as we are doing them so we don't have to do this extra step
of process. This has worked well. 

He's good at receiving feedback, very receptive to any comments. As a manager,
he's clearly coming with experience. Feels very genuine, ramped up very quickly. 


## Guy.md on 2022-01-27

GPP - impression from kshams "I was asking for reliability numbers and better
understanding a while ago, and I would've preferred to have meeting to review
what we have right now". 


## Wa.md on 2022-01-27

Working on how get ODAD online. Wa's part is collecting results. 

We build a parsing pipeline to collect the data from Brella. We also monitor
the data flow. Fixed few bugs with Astrea. 

ODAD is sending data to Astrea. Astrea is on the device. Astrea communicates to
Brella. Core function is federated analytics.

We send malware digest and module whether it's from rep module or from heuristic
module. 

Brella receives the data and saves to CNS. After that we can start collecting
the data from CNS (after secure aggregation). We can see how many times each
digest appears in checkpoint files.

Moira - periodic classification worker, connected to Hades team.

Team - mood is OK, Nikolay and Ishtiaq are getting more and more knowledge about
the team, they are active in meetings. Wa feels we need more help, applies for
interns.

William - every time I need help, every time he's there, very detailed, taught
me many things, has lots of suggestions and guidance. Explains a lot, provides a
lot of help, if he doesn't know finds someone who knows.

He may need help because after Alice left training part we are not quite good at
what's going on even though Alice did a knowledge transfer session, sometimes we
still don't know what's going on. 


## Muhammad.md on 2022-01-25

Pretty much everything is green - go/angel-eng-progress

HC filled up

Prabal is also coming rouhgly at the same time, so Shriya and Prabal both on the
same level.

Challenge for this week - we are looking at how inline actions / rescan, how can
we do it more intelligently. If we try for everything to have inline action,
it'll take long. How can we put dynamic sources in first iteration, and then add
inline actions later on? 

Mid-Feb - we are ahead of the curve. We'll have XML config and settings sources,
Phonesky will be progressed as well as GMS stuff. The challenge might be that
some sources are not integrated at all and placeholders will be there, and UI
looks very clunky. 

Team morale - fantastic last week. 

We are doing glueing end-to-end, all the way from adding 1 source, static
entries all the way to XML. We'll try to have it by 7th-ish.

Team wants faster turnaround. Settings - can they turn around stuff faster? (Edgar)
Likewise Karishma - Rubin.


## Fit_Calls.md on 2022-01-25

William told: you do must server side, but things might sneak through. Scan on
the device. 

Ryan's thoughts: really interesting. Lots of things jumped out to me. Lots of
things I try to look for when discussing with team. He brought up some
interesting constraints, clever with solutions, careful, not straighforward.
Constant comms with privacy security teams. Somehting I've been doing a lot at
my current team, something I enjoy.

Most recent team: TLM of data mgmt team. From a high level, greater org is big
data extraction/processing platform. Provide platform that allows data
scientists. Data mgmt team managers all the data to make sure it's compliant by
external/intenral policies (GDPR etc). WHen you make a lot of changes, onboard
new data you want to understand how we classify it, not really on me to do it,
so we'd have strong comm channels with privacy/security teams. Discussed a
little with William, what my aspirations are.

Definitely more extroverted side of things, enjoy meetings a little bit. Like to
have these open discussions. 

Super long term, see myself growing in leadership area. Like technical breadth
vs depth, being able to oversee several projects, not just deep diving into one.
Really enjoy: ability to influence others. Mentoring, affecting their work in a
positive way. 

We had a big celebration - had reliability of one of our APIs from 99.5 to
99.9, that was cool! 

Questions:

1. Domain/problems: it's also the way you described it, resonates very strongly
   with me. The test that I do: am I excited about it -- what if I told others
   what I did? Would they get excited? --- but that wasn't really a question
   [RK].
2. What do you find the biggest challenge?
3. What was the most catastrophic thing that happened? 
4. What traits one must possess to succeed in this team/org?


## Anurag.md on 2022-01-25

Anurag spoke to Guy re OKR this year
Asked how things going with Finder
Came to an idea that might be good. Goal for 2022 - establish FMD as useful
product users know about. For the secure and private part: 2 concrete goals. We
know vulnerabilities we have to fix. For the part how many users find FMD
helpful, we should do a survey how users know about FMD. At the end of the year
we can do survey again. 


## Bernardo.md on 2022-01-25

Finished all features for T. Ricky is working - Ed came up - enhanced
confirmation mode. For apps not from store, for those apps if you try to enable
any restricted permissions, like a11y / notification you'll see this setting
will be disabled and user will have to go through another setting, restricted settings.

# Week 3
## Ronald.md on 2022-01-20

Two kids 5 and 2

At I/O planning to announce tags support.  Support for anti-tags.

Stalking protection. Rogue tags still part of the agenda.

Headphones in essense are tags. 


## Fit_Calls.md on 2022-01-20

Currently work @ MSFT SWE. First project - part of Teams - Android mobile code.
The area my team works in. Communication Platform. Similar to Google Meet.

That's the time I got introduced to Java and mobile. One of the reasons
specifically applied to Google / Android.

Next project - also Teams, part of IT Pro org, goal support Teams adm management
for large scale enterprises. Configuring/mgmt. Got experience with BE (C#/.net).
Exciting was to work on product used by millions of customers. That's what
looking for going forward, large scale impact. ASAP is perfect in that aspects.

MSFT Turing (?) team - very large collaborative team across MSFT. Goal introduce
ML capabilities into many existing software products at MSFT. Collab with Bing,
XBox, O365 et al. Get to work xfn. 

Dir or Principal? IC track smth I enjoy. Would prefer mgmt aspect. Getting to
define what products we work on. Mgr is about working with people, interacting.
Do enjoy people aspect as well because really like collaborative aspect. I enjoy
contributing to proudct definition and providing guidance along that lines, as
well as interacting with other people.


- Project and the task I will be working on. What is the main goal of ASAP team,
  what is the direction the team will be taking on?


## Khawaja.md on 2022-01-20

big thing - there is HC. haven't received it yet. Couple of buckets.

- On-device / ML
- Rubidium
- General App Safety

Mobile one should have something for Angel and ODAD.


## Aaron.md on 2022-01-20

Metrics - not really enough.

Nikolay is still onboarding, but trying to include him into Westworld stuff.
Have to put it all through AiAi first. 

AiAi didn't say no or anything. e.g. if you log any type of counter, they have
very specific way of doing that, regardless of how you are logging it.

Little less clear, eligibility eval task (?) tasks deciding whether to
participate. 

go/odad-westworld

Nik - talking to him, he seems to be ramping up pretty fast. Assigned to him a
little bit of Kotlin conversion work, he's doing good at it.

Kotlin conversion -- many Android teams are transitioning into this, async code
is handled a lot better. ListenableFuture vs couroutines. Some APIs are not even
supported (in Java?) -- standard KV store doesn't have a very good Java API at
all. A lot of things are preferred in Kotlin, get first class citizen; prod
excellence. 

Not a ton left, so makes sense to finish it.

Removing tons of boilerplate is nice. Since this is the direction Android is
going, been noticing more and more new developed things don't treat Java as a
first class citizen. Jetpack Datastore - if you do it in Java, they don't even
have async APIs anymore. 

Ishtiaq - still improving. He's helping with few things, Kotlin conversion.
That's been going pretty much fine. Almost 100% converted, just a few more parts
convering ListnFutures to coroutines.

Ishtiaq is also working on a test app which simulates malicious apps. Just
started a week ago, he showed some progress, goes in the right direction, but
too soon to 100% tell. 

The number of questions about things he should know about has reduced.

Now we are in Googlefood, sending digests back. We intentionally lowered
treshold, these digests are FPs, but at least we confirmed that the flow is
working. Processing features and processing results etc. Once we in prod we'll
move threshold to 0.95, happens in a week or so. 

ODAD v2 - basically done right now. Ben did everything for that. Knowledge
transfer to Aaron/William but maybe also Nikolay.


## Yuri.md on 2022-01-20

5 year old

helping Jan with his project, all in settings for Angel

Team. Just 1:1s.

Muhammad. Doesn't micro manage. Knows what I'm doing, interested in how are
things.

Most experienced of all managers. Super-ability to find solutions to problems no
one can solve. 


## Fit_Calls.md on 2022-01-19

https://hiring.corp.google.com/hiring/hms#/hms/reviewcandidate/a1e9c4ef-5efa-407f-9a37-079e5182c210/candidates/google.com-446965773/applications/451796471


7 years ago moved to the US
Alkami Tech

Member who joined in early stage, company went IPO. Worked all aspects of
platform. Focussed on contect creation delivery feature. Allows our customers to
create promo content.

Full stack.

See myself to be in top engineer position. My career path was focussing on going
down the path of engineer. Be an expert in a certain area. That's how I position
myself. 

Really like technology. Really like using my ideas to create new product
features. Principal Engineer.

Want to be a team lead, who will lead a team, which will create excitin feature
/ product, ultimate goal.

Sometimes need to launch / debug mobile / webview issues. Sometimes need to run
an emulator, and to run our app in emulator, track the traffic. 

Have android tab and iPhone.


team is 3 SrSWE incl TL. Ming is the most senior member on the team, not a TL.
60% is designing / coding, other 40% - mentoring junior developers, maintaining
internal wikis, etc. Jumping in on product issue, working with devops. We have
PMs, we talk to them and figure out new feature they want to implement and we'll
give them ideas and estimation.

I'm not a TL. I have been a TL on another team. But after we create content
delivery (below) there were certain changes between teams and they needed an
engineer very familiar with feature to join another team and take feature with
them. That's how I ended up on that team.

How does he feel about not being TL? No hard feelings. What thinking is, Im
familiar with this feature, doing knowledge transfer, to other team members,
continuing doing my work. Eventually the TL is going to leave as a manager, we
are in a talk that Ming will take the role of TL or we find another senior
member. He really likes me and wanted me to be the TL. Don't want to jeopardize
the team because moving company.


Content creation and delivery. Mgmt of company had idea they want to improve
content creation and delivery feature, which our customers always complain
legacy one. 

We were assigned to do this work in a very short period. Didn't follow normal
process. 3 senior SWEs and PM, brainstormed together. We see what users should
use, vis workflow. I contributed a lot of my ideas on how feature should be impl
so users will love to use it. We discuss this. PM would document it. We'd start
creating a roadmap for it, for short time period, how it'll be implemented who
will work on what. 

I'm proud of image picker feature, I'm the main contributor of this feature,
contributed many new ideas and implemented PoC in one week, shown PM, he really
likes it. Eventually 3 of us implement all that and create roadmap for Phase 2. 

We had it released back in 2020, user response - how we measure it? We put
metrics inside the feature, we are seeing usage is growing since it released,
rapidly. We asked PM to talk to end users, and their feedback is really good.


Williams team is working on a confidential project. Feature is like windows
defender, that's embedded into Android platform. -- this feature really excited
me, very interested in implementing some very new ideas. 

How large is the team.

What are the expectations - that it will be released, etc?


1. I will talk to as many people as I can to have some knowledge transferred, or
   will find some documents on the internet / google everything to learn as much
   as I can on this subject.
2. While doing learning will work in parallel on at least brainstorming /
   talking to other team members to have an idea of what my roadmap will be
3. Will need to make plan for impl of the feature
4. Then will start working days and nights on it.

Communication is very important. Will find anyone acessible to me, be that TL or
other people, PM, etc ... Talk to as many people as I can to gather as muchn
information as I can. That way at least I'll have clear image of what I'm doing.

That's what I've been doing on this team by the way, all these features are new
to me (except content creation). As soon as I joined the team I set up meetings
with TL, PM - they didn't have any documents. During my learning I modified the
documents for future members. 


## Nikolay.md on 2022-01-19

Continues to push telemetry.

The app was released before I joined. Without proper information about what's
happening in production, it was more or less not working in prod. The missing
telemetry - still the biggest problem.

When talking to the team, (someone) wrote some proposal last year for WestWorld
integration. Hopefully we can put something in place so we can have our feedback
on what's happening in prod. Still #1 problem for me.

Once we have this system in place, everything else depends on that.

Can PCC app send back any telemetry data? Still not clear on all the rules,
tomorrow meeting with WW will hopefully know more. That's something new for me.

Current understanding - just app release, but not 100% sure.

Aaron discovered (problem) by looking through source code, someone else found
something. 

Ops review - look at metrics in production to see what's working. Have to talk
to Wa and rest of team members of what we can monitor.

We are running on 14M devices right now. Respectable number. 1% becomes really
big number.

AiAi is using WW. The door closed after them and the rules became more strict.

People are really good with filing improvements, fixing here and there, but big
picture is missing. William promised to schedule a big brainstorming for the
future. N asked this to be a regular thing, not a one-off, just having a 1hr
brainstorming not going to do it. 

With the current team size it's def understaffed. Codebase isn't very big, but
we are missing a lot of capabilities around testing/reporting so we can say with
confidence we are delivering good solutions.

Unit test are in a good shape, but above unit tests, reliable infra is missing.

No regula release cadence. Only when incidents.


## William.md on 2022-01-19

Lack of telemetry - agree with Nikolay. Wanted to do WW in v1. Taking this
seriously, both Nik and Aaron working on it.

William spoke to AiAi and they said aggregated counters is OK. 

Above-unit-test-coverage is missing. Working with Hades which doesn't have e2e
tests. No e2e for fedanalytics. Hades need engprod help.

Downloading the models - from Astrea. Promised to be done by EOQ2.

Teaming up with messages team to escalate.

Content capture - all text, near instant. Right now allowlist but want to have
all apps.


## David_Coffin.md on 2022-01-19

Feel more positive now, have to drop some requirements / simplify. Still think it's high risk.
Elliot reasonably confident that if we keep team focussed, we can get it done Marie is pretty
pessimistic. Once we go past IC complete deadline we have 8 weeks before Mainline train departs.
Given the pace of elements in Mainline, it's a very short time. Concerned we might get some last
minute feedback from OEM, which might or might not be launch blocking. Concerned we are integrating
so late in the process, that increases the risks - forced on the team by  API approval process. Not
possible to add hacky API to master to integrate end-to-end as PoC. Test population we get will be
smaller than we'd like. Concerned this means we might not find out about serious issues until too
late. Amount of work vs time is the main source of risk.

We had similar amount of time for S. API was completely rebuilt.

We doubled the size of the team. Mainline is not twice as slow, its 10x-20x slow. We can run unit
test maybe 3-4 times a day, previously it would take just a few minutes. We knew working in Mainline
will take longer, we didn't anticipate it will take so long. A lot of lost engineering productivity.

mb Elliot can work on something else in background, e.g. prototyping ODAD integration. 

Majority of our code is in the system server, logic, storage etc. This means we have to compile
entire Mainline module => slow refresh cycle. We expect UI stuff proceed relatively quickly. 

The code we are doing in system server in some ways is similar to what we done before, but we are
changing many things.

Past IC complete we do whole spring where we integrate everything and try to make it all work
end-to-end.

In terms of developer prod - not a lot we can do about that. Checking with PC /
engprod teams?


## Muhammad.md on 2022-01-18

David - not very organised. Very worried time is running away. 


## Anurag.md on 2022-01-18

Tags no longer the main focus of Finder launch


## Guanhuan.md on 2022-01-18

GPP -- definitely oversight. Difficult time ramping up on Angel. 

Notifications should adapt to be better fit with Angel.


## Angel_Cheng.md on 2022-01-18

2 kids, went to Hawaii for Christmas

older one good swimmer, she swims really far

7 and 10 girls




## Guy.md on 2022-01-17

Dave - Angel will be one of the bigger things in new Android. Will get a big chunk of I/O. 
A lot of people are going to be tracking it and looking at it.

Some influential reviewer on YT - marques brownlee - did a very sceptical review, caused some rep damage wrt quality/stability

Many people are very nervous about quality/stability. A lot of scrutiny. We need to do a better job, careful, more explicit job in communicating risks, where we are, what do we need. 

Both Guy and Roman need more time with Suzanne, would be good for our career growth. Opportunity for me to work on exec briefing skills. Sometimes we let other people (PMs/PgMs) do that.


## Giulio.md on 2022-01-17

Blocked on Rubin - biggest problem
OEM requirements - how do we enforce OEMs do the things the way we like?

Say we make a req - how do we enforce it? 


# Week 1
## Guy.md on 2022-01-06

- Angel hiring
- Situation with Aleksandr
- ODAD backfill


Big strategic vision for ODAD - think 5 years from now.


# Week 51
## Narayan.md on 2021-12-21

External API is likely to receive most push back. We have to finish API by API complete.
We can backload actual implementation and GTS/documentation

Need to ensure that the scope is extremely limited - doesn't affect apps,
doesn't affect internal system stability, can turn off functionality. No changes
in common code - i.e. all functionality is Angel specific.

API is between system server and PC. 

We need a flag to control entire thing which we can flip internally / externally
to get droidfood feedback.

Code complete, goes into mainline module, it goes to all devices - T, S, R ...
and then we separtely have flag rollout to droidfood the feature. For that flag
rollout we need population big enough for this feature. That population Narayan
is suggesting to use QPR (but we can also use Googlefood etc).

If we find a bug in API in March we will still be fixing it, so it's mainly
about semantic.

Narayan - we can pull this off, we have worked on far riskier features. Anything
that affects app behavior is risky, this is opposite of that. Narayan is not
worried at all. 


## Peter.md on 2021-12-21

PV: Current allocation - hard to deal with
Two OEM reported issues. One was release blocker for Samsung. Other fixing now,
as verifier needs to be updated to match changes in the OS. Then there were also
incremental installs. 


## Niko.md on 2021-12-20

Integration - not looking at a wider picture.

Wrt notifications being proxied from GPP to Angel. eg. 3x CTA


# Week 49
## Mike_Tsao.md on 2021-12-09

Overall positive with how things are going
Don't know whether we will hit the deadlines
Anurag has been an especially valuable contributor

Not so much in a high-bandwidth sense, but he provides a real positive energy to
the team. Got feedback from couple of people really enjoy interacting with
optimistic TL.

HC req to CC


## Guy.md on 2021-12-09

Situation with GPP team. Forced mandates from OEMs. Time running out.

Angel - naming, current status.


## Muhammad.md on 2021-12-09

Notifications:
 * The situation was: there were req for notifications, we weren't sure if we have
   contradict req from API council vs Product, weren't sure how to resolve it.
 * Behaviour: a number of long meetings/conversations, spending a lot of time on the doc,
   little resolution. We would've resolved quicker if we would've got right people in the room.
 * Current situation: depending on what OEM say. We are good for notifications. 

 # 2021-01-11

 Challenges:

  * Definitions of the APIs -- in a normal world we'd do some prototyping. We
    are trying to fast track this. E.g. integration with biometrics and
    screenlock is a big concern. Marie leading it, she's only got 2 months. With
    OEMs we might even have bespoke integration.

  * XML config. Big challenge. We are missing the tests - don't have XML
    definitions, don't have anything for IC. Big concern. Not an interface per
    se. Need not just CTS, but also GTS tests.

  * Integration with the sources


## Guanhuan.md on 2021-12-08

Angel vs Safety Centre



## Gabriel.md on 2021-12-07

Dialogues still going on

Niko is phenomenal manager

knows how to delegate things

Tech proficiency. Very available

bessiej good person not only with G but in general

had a chance to work with Peter - v knowledgeable

P different kind of knowledge. V tech proficient. Right-a-lot. Good ideas.
Nothing against him at all, funny guy. Whenever ask question, gives short but
accurate answer.


## Muhammad.md on 2021-12-06

Did we send a welcome email for Tyler?


# Week 48
## Guy.md on 2021-12-02

Finder - build credibility, smaller things

Do a review of what's happening?

Become a part of a leadership team


## Running_notes.md on 2021-12-02

Linus is unable to complete the work he committed to - can't do actual
development. Push back on Ash?

Concerns from GPP team - fewer people will be driven to GPP to solve problems
there. Need review with Niko and Muhammad.

Play Permission - dragging on and on. Involve someone from GPP team?


Trying to enrol more people into APP, some people even automatically. Dropping security key requirement.

Carter. If we could put protections automatically, how would that look?


Remind how we went a long way
Success of Angel in Android S
Even bigger thing.
One of the most exciting features in T
IO time
Time is hard but things look plausible, Narayan also says so
Remove obstacles
Execute


what we are too late with Samsung - with partner feedback? DO we push it to U?

Sara: feature will land late. We won't hit beta.

Creative ways to test feature on OEM devices? What are the options?

RyanBarnett: if we are making API changes they have to be in T. For Mainline
QPR1/2 we can have requirements to enforce.

We can give GMS waiver for OEMs for a while until e.g. any build coming in in Q1
2023.

Suzanne: Someone to put together next steps on how we are going ahead with
Samsung



## Notificaitons_and_Pull.md on 2021-11-30

Problem with warnings cards - it's the same thing as notifications really as
OEMs can technically try to upsell users on paid add-ons. Visibility will not be
significantly different.

We can allow notifications when they are coming from the OS itself as well as
system components such as package verifier. This means we can show privacy
reminders and GPP notifications.

We can allow notifications from components which are signed either by Platform
key or Google (this will include mcaffee on Samsung devices, but they are not
trying to upsell) with GMS language prohibiting upselling user on any paid
upgrades. 

Pull on launch - we have to work with perf team and escalate at the same time.
Trust into Google safety is too important thing for the company + existing
precedent. 

Ideal scenario: push back with some compromises as suggested in the document.

## Khawaja.md on 2021-11-30

What to do with FMD? Current situation.


## Narayan.md on 2021-11-29

Google-only notifications only - worth asking

Send Narayan go/angel-notifications

It's API council's job to assign sponsors.

There's a training program for API council. Large time commitment. Narayan and
Wale had API council privileges revoked because time. It's community service /
citizenship, non-trivial time commitment. 


## Nandini.md on 2021-11-29

lot of things in Android are cross-org

collabs in Android are fluid

guide team to be succesful in cross-team world, challenges etc

or come to mgmt and say this is not a path for me

nandini worked with MikeT a bit, shared admin
few people worked on safety led largely by N

Erik makes most decision


# Week 46
## Aaron.md on 2021-11-17

Replacing Hades would be difficult

Ishtiaq -- improving. Mentored Sam when he came on. Mentored intern. Don't have
tons and tons of experience, but to Aaron feels that by now he should've been
more independent. Still asking lots of questions. Def not there yet for L4.
Should dig more on his own. 


## Khawaja.md on 2021-11-17

HC for 2022 - ugrad

lot of pressure clear negatives in the org


## Nikolay.md on 2021-11-16

Still learning context. For ODAD v2 some discussions are going in circle. 


## William.md on 2021-11-16

For next one-to-one -- team doesn't feel like a team


## Dave.md on 2021-11-15

Strategy session when we've last spoken - how did it go? Project Cake?


## Giulio.md on 2021-11-15

Working with Elliot and Karishma on migration

Tomorrow have a whiteboarding session with Dario et al - where the code should
go etc. Had meeting with Hai, pros/cons putting lots of logic into system
server. Have a general understanding.


## Muhammad.md on 2021-11-15


3 subteams -

David - unlocking API definition, possibly Giulio too (data store)
Elliot and Karishma, maybe Yuri, bulk of the work to port
Jan/Marie - UI element, blocked only on our subteams

=== for next 1:1 ====
some agile-y changes were introduced maybe too early, and team is now back to
what it used to do before - thoughts?


# Week 45
## Krish.md on 2021-11-09


2024-2027 - a bit far out, unclear.
Changes to landing page for 2022 when user is coming from Angel
Existing GPP can stay existing GPP minus branding changes

Making sure that flows and hand-offs consistent between Angel and GPP -- the
2022 plan.

Post-2022: backporting is humongous project, don't want to sign team for that.

When you unlock Angel in 2022. A lot of good credit, OEMs, ecosystem. After a
while - can we get this for older ecosystem? Then we can talk about what does
this mean.

One way: build Angel for older devices. Another way: make GPP look more like
Angel <-- thats what Krish proposes for 2023.

Krish asked Guanhuan - nail PRD and plan which talks about how does Android T
integration and handshakes work with GPP?


Dave's view point is not actionable. There's no simple solution.

https://docs.google.com/presentation/d/1L07-vTQdQhqiwU2Su8plTnTZc4xu58M7yALAcWxWkKY/edit#slide=id.gecf194057f_1_58


Project Cake
Waiting for central plan for ap4a from AP team
Stay in maintenance mode for now.
What are the things we need to fix to go to 100%?

What existing CUJs will change? What immediate investment we need to put?


## Running_notes.md on 2021-11-09

niu@

people on abuse side
- sconsolvo@google.com
- kurtthomas@google.com

go/mspy_analysis_report
go/mobistealth_analysis
go/highster-mobile-stalkerware
This drive contains the mobistealth APK 
https://drive.google.com/corp/drive/folders/1RBgE4BILOEXhS7gLol0dyOGOl4l42X6Z


The team has grown a lot in the last year
30 FTE right now!

Changing the format - fit for better purpose

* New joiners:
* Angel team
 - Yuri
 - Giulio
 - Tyler
* FMD team
 - Varsha
 - Yusef
* Aleksandr
 - Noriko


Wins of the week - in future someone else right now:

- ODAD v2 fully approved by security, privacy and legal
- Angel API CL is submitted - yay!
- Aleksandr - permission auto reset will start ramping up in Jan, Google blog, video!

Permission Controller 


## Monir.md on 2021-11-08

We know where infra gaps. Dave: we need to do better in AV Test without emphasis
on off-market. Are we trying to do too much? He wants us to focus on short-term
shortcuts to make us better.

Khawaja is maybe we should do off-market properly.

DEX, manifest hashes


# Week 44
## Muhammad.md on 2021-11-05

50% GTI + Courses
GTI - this week it's done, just once a week half day



## Rebecca_Salois.md on 2021-11-03

Moving away from focus on Pixel/Premium, focussing on Trust and aging
gracefully, Trust and NBU, etc.

## Gabriel.md on 2021-11-02

Starter project - refactoring GPP dialogues. Writing design doc. 
Niko - good manager. Like him as a person, lots of matches.

Manager: very available. G asked lots of questions, N was always answering them
- e.g. Robolectric. Rather than solving problems for G he was pointing to the
  right direction.



## William.md on 2021-11-01

Giles approval painful. Could've gone better

Launched at 2% right now, 20% today

If we can get stuff from federated analytics, it's success

Stretch - find any malware we haven't found



Next 1:1 -- areas of delegation


# Week 43
## Comp_Planning_2021.md on 2021-10-29


William - bumping up equity to $200K to send a message, last year was $190K

David - bumping up equity to $145K because shows promise

Marie - bumping up equity to $116K due to exceptional work in the cycle, near-EE
and generally very promising engineer

Niko - bumping up equity to $110K because upward trend

Aaron - bumped to $100K 


## Anwar.md on 2021-10-28

make a plan 
how much customisation OEMs can do?
talk to partners, talk to Samsung, Xiaomi, Oppo - send a note Anwar, TT, Ryan
Barnett (reports to TT). 

when MSFT built av, plug in different AV it will give status.

Let user choose the AV provider?

something that doesn't confuse the user. 

Reconcile: get user's choice (McAffee/built-in one).

## Jeroen.md on 2021-10-28



Rodrigo to report from next cycle

Set up recurrent mtg


## Guy.md on 2021-10-28

Mum in Jerusalem, want to move tlV

Hiroshi: demand for Pixel phones higher than expected

Khawaja/Monir/Wei -- AV Test


## Dave.md on 2021-10-28

- Top tier security, above GMS
- We need transparency that shows diff in implementation, XIaomi
- ASE side - we know we can't get abuse rates down to the level we want if we
  look to the whole store. Project Cake. Can we get high severity abuse down for
  curated section? AP4a is a guinea pig on how we curate at scale.


# Week 42
## Running_notes.md on 2021-10-21

Launching something at I/O, not scoped yet


https://screenshot.googleplex.com/3E8DmsMxgts9uhf






## Khawaja.md on 2021-10-19

Hiring


- Public launch tonight - actually, right now.


- Most design documents are complete
- Going to Mainline review this week
- Programme review with Framework (incl Dianne and Narayan) went well
- First set of CLs with Mainline API is out, in review with PC team

- Discussing OEM strategy this week

< to be updated >

Launching this week. Press spotted ODAD Play app, mystery!

- On a good track to add what we need from Framework. 
- Had a good conversation with Narayan about what we are doing, he's fully on
  board, added few suggestions.
- They were incorporated in meeting with Svet & Erik, they are fairly supportive.
- One risk is communication between ODAD (PCC app) and Angel (it's not). Meeting
  to discuss this with many folks (incl PlatSec) tomorrow. They may push back.

Niko -- open HC, cancelled interviews


# Week 41
## Mike_Tsao.md on 2021-10-15

set up 121 with A


## William.md on 2021-10-14

None of Mainline modules fall into private compute core. Only AiAi and ODAD. 


## Inara.md on 2021-10-14

Reliable usage stats

1% at beginning of December / when we get QA approval
After 1 week we'll start revoking permissions beginning of December, but for 1%
of users

Assuming 1% goes fine we wait out the holiday period and then in January start
ramping up to 5%

Niko - doing well, really helped with quite a few different things going on.
Been very helpful.

Did address a "lack of vision thing". Made more clear to Inara. 


## Elliot.md on 2021-10-14

Prototype - right now have a PoC CL that seems to highlight that design doc from
David seems appropriate, planning to submit it to Mainline.

cr/403055293 - PoC in google3, moving to AOSP.

Writing it right now.

Giulio - working on his starter project on S. Got it working partially, not
triggering for some reason. By next week he'll have some CLs to send out.
Talking about different approaches for S. Keen to contribute to T.



## Anurag.md on 2021-10-14

BT - had conversations with Ronald and Mike separately, priorities for 2022.
They re aligned with plans that we have, as in continue to fix privacy and
security model for FMD. WebUI: he thinks we shouldn't invest a lot into it,
especially until after I/O. Second round of prioritisation after I/O, what's the
longer term we can do. 

Ronald asked: he can involve us in other tracks on Finder project if we have
capacity. 

Wei is still hiring a team for the app. Till they hire a team they can't start
on UX revamp. The refactor of the app is on the way, sponsored by Spot team.

Web UI - continuing to move to Boq Web. Have to decide if after migration we'll
revamp the UX. Right now it seems not do until I/O, as there might be higher
priority things.

Mtg Mike went pretty good. He was super happy Anurag helped with denisty
analysis. Talked long term collaborations. Doesn't see this as one team owning
something, but team being responsible for something. He seemed not to be too
involved into short term picture. 

He was asking about other ideas, A said about detecting the device is lost
before find out, he said interesting from security point as well, ASAP might be
interested in it.


## Giles.md on 2021-10-12

One big diff security / privacy warnings - priv warning can protect you from
yourself. There's not always an attacker in privacy space.

At Google level - Safety is marketing privacy and security umbrella worlds.

Stalkerware - preloaded or not. Preloaded to filter out.

Really like the focus on at-risk users. You get much bigger bang for your buck
when you solve something when people risk physical violence. Because there's
such a high risk, if we get it wrong, people might die. Do some serious
adversarial testing on how this can go horribly wrong based on our design. We
should do a red team excercise on this, maybe Giles can volounteer some people
on the team. UXR of a sort.

On the topic of at-risk users - wanted to run proposal. Liz/Luis brought issue
of Afghan refuges trying to get through the check points. Liz was contacted by
several NGOs/friends to say what can an Android user do to get through
rubberhose Taliban checkpoint, with the document that proves they collaborate
with U.S. government. Complete deniability. Another problem of if we get it
wrong, people will die.

Idea - as a part of OS / Files we provide something that is on every users'
device, you can't use its presence to identify at-risk actions.

https://docs.google.com/document/d/12EK_-nBCTIeFOaNbua6PnTmZLXAwQlD__18YqPqr8v8/edit?resourcekey=0-vih4uV3pr-S58mbXJeLmOw#

Re-working the read logs permission. Incident: priv preloaded apps with
READ_LOGS permission. EN was logging a bunch of stuff to logcat. In T we are
adding a mandatory dialogue to user every time you access logcat with this permission.


## David_Coffin.md on 2021-10-12

Concern - we may not have time for radical change of the planning process,
something Muhammad might want.

Elliot - very slow, not much happening.


## Jan.md on 2021-10-11

Working on foldables. Most bugs fixed.

On foldables, 2 pane screen. Everything that has same UI design as Settings that
opens from Settings should open in 2-pane mode. Some things in Angel (FMD,
lockscreen, Security Checkup) open in full screen instead. Make activity
launching from the same task as Settings. 

Have to get a local build of this foldable thing running. Doesn't open in
two-pane mode. They are on holidays.


## Running_notes.md on 2021-10-11

AppRisk scaling up 4x

Big change in strategy moving towards pre-publish world. Move the model much
upstream.

Post pub model - after the app has been classified and reviewed, goes to store.
Post-pub risk model - in depth review. More self contained space. For last 3/4
had success post pub risk model, among 10 flagged ups 6 high severity
suspencions. 4x of what we can do in routine procedures.

Synthesize signals from all sources, including on-device teams. Assemble them
together into a risk score, (0,1). We take top X apps and send them for in-depth
review. IDR is a tier 2 review, more thorough, reviewers are looking for
specific areas.

Example signals: massaged user feedback, 10 topics users are discussing for this
app, e.g. ads or disruptive images, or never got paid from the app etc

We show them distribution pattern. If the app is only distributed via deep
link.

To bring risks to submission time:

- Type 1: new submission we know nothing
- Type 2: updated to an existing app. Combine knowledge about version-1 with the
  new APK

T1 - generate risk score based on existing signals.

Low quality apps are being rejected right now. Then there's pQuality - it is
taken as a feature of a risk model. Quality and abuse do not have 1:1 mapping.

Let's say they have very good ads account, getting revenue from Google => most
likely they are OK. If there are many accounts and they associated with 1 app,
then app is more likely to be dangerous.

Web Property ID - self declared, we are talking with ad spam team. We want to
take their ads account risk score and apply to app risk score.

1st version OK, 2nd is dodgy - very often.

GAIA risk score. PG's team is doing developer risk score. Mission has some
similarity, but 1 year goals are very different. 

Dev Risk Score is #1 feature in App Risk Score.

Prioritisation Infrastructure. We have a lot of infra for
review/enforcement/classification. Building general purpose infra. More
intelligent infra effort, when we don't require people to do some things.
Provide intelligent infra to allow teams do their things. Monir team: you have
Subrah, do you want to plug in this into our infra? Not replacing anything, new
infra.


## Narayan.md on 2021-10-11

Some sort of sandbox runs it, verdict

Don't need to ask explicit permission for this things, but having clear design
doc would help, so at least we have someone from AiAi team which validates that
approach is sane.

AI: make sure we have a high quiality design doc
This is FYI for framework, and Eugenio

Role - might not work in the medium term because other app can fill this role on
non-Pixel device.


# Week 40
## Muhammad.md on 2021-10-08

https://www.google.com/maps/place/Sixth+Cross+Rd,+Twickenham+TW2+5PD/@51.4365052,-0.3611637,15.71z/data=!4m5!3m4!1s0x48760ca3bfbeb7bb:0x897aea65f3ed57cd!8m2!3d51.4371259!4d-0.3565104

Son loves cycling

Hi there, 
 
My name is Muhammad Burhanuddin. I originally hail from Karachi, Pakistan, but
have lived in London for the past 20 years. I am 46 years old and married to
Najma. Together we have two teenage boys, Omar (18) and Ahmed (15), who are our
pride and joy.
 
Before joining Google I worked at Sky for 11 years where I contributed to the
development of some of their major flagship products & features, such as the
content discovery experiences on Sky Q &  Now TV and Glass which has literally
just been launched in the UK!
 
I am passionate about technology as a means of bringing joy and entertainment to
people.
 
I like taking long lazy walks and bike rides. And I love good food 😋 (who
doesn't!). I've recently started learning how to cook properly 😋
 
And I am super excited to have joined Google and really keen to help take
Android Security to the next level, to provide our customers peace of mind, safe
in the knowledge that we've got their back.


## Rebecca_Salois.md on 2021-10-06

started at Geo, then Research People+AI. Running research programs Trust in AI,
etc. 

PgM, working very closely with Jeroen build out Trust UX team. 

We have a set of thigns we are doign for T, big step forward, but we can do a
lot more and we can go a lot further. 


## Khawaja.md on 2021-10-06

Hiring

Angel

 - Engineering update
 - UX update
 - Product update

FMD
 

Angel reviews - keep bi-weekly. 

Talk to Dianne


## RichardN.md on 2021-10-05

Lots of talking about on-device detection. TAG wants better instrumentation and
logging. Sebastian is interested in doing some rootkit detection.

Privileged AV increases attack surface. Windows AV is an example in this. 

either platform or priv security agent

platform => anticomp
priv sec => platsec dont want


## Dave.md on 2021-10-04

Update on Angel v1
Angel v2 progress
Safety Centre conversation
Update on hiring
Current situation with FMD.


# Week 38
## William.md on 2021-09-21

Nikolay starter project - get activities time in foreground


## Erik_W.md on 2021-09-21

Notification Permission (SysUI) - mgmt UI





## David_Coffin.md on 2021-09-21

Team energy levels went down a little bit no no particular concerns just yet


## Mike_Tsao.md on 2021-09-21

Model of work
Who owns what
OWNERS and all


# Week 37
## David_Coffin.md on 2021-09-14

Cutting Angel APK now

One more CL wanted to get in this morning

Few more reviews. Need to spend some time this week making sure launch process
is document, monitoring, rollback process. 

Elliot starts looking at prototyping today. Was planning Friday but had an
urgent issue.

Assuming one urgent bugfix will be needed.

We have logs now. 

David wants the chart the day we launch 70% of users were in amber state and
after interacting with Angel, 50% are in amber state.


## Peter.md on 2021-09-14

2w in Hungary and then Malta for few weeks


## Guanhuan.md on 2021-09-14

birthday yesterday, 13/09



## Krish.md on 2021-09-14

https://docs.google.com/presentation/d/1DIqmiclMH8J0fM-bQYuJslL4F8EeytUzg0eDxZMHSKA/edit?resourcekey=0-e7T3PAgF5n88QLMsJRFBMA#slide=id.geecc7e1182_1_231
-- discuss this

https://docs.google.com/document/d/1PuXKqf89597O0zWk1mX2D4Bjz_xkn4rsnV2a-KPhhIk/edit
-- get LGTM on this

Matt not taking it


# Week 36
## Inara.md on 2021-09-09

reliable usage stats - sorted, having conversation with msnider@, want to
provide reliable data to them but don't want to commit to too much work because
we need some very specific data.

Finishing in next week or so. 

Launched ODML warnings to 100% couple of weeks ago

high CTR for that, 20%



## Khawaja.md on 2021-09-09

Short term argument - we are still responsible for FMD, can't let someone go and
blindly change things. 

Use Angel / Permission Controller argument. 

FMD small team, uptime is critical. 

we are working together give us headcount

another option - change Angel meeting today


## Karishma.md on 2021-09-09

Feedback for David
- Great with managing many different things happening in Angel, delegating quite
  effectively.
- Also was very clear on how much autonomy you have, makes me very comfortable
  as I know which things he wants to greenstamp and which things he's OK for
  people to make decision.


## William.md on 2021-09-08

Rubidium situation

Sudhi is running it but not been doing a good job keeping us in a loop

Rubidium conference, we weren't invited till last minute

Reason why not well communicated because they don't intend to support our use
case, even though Guy was trying to get them to support it.

What do we want? To allow ODAD to get system permissions without being system
privileged app. Something Sudhi claimed is possible, James Bender says it's not.
This way we can't be in a system image so we can get deployed to any device. 

Sudhi said we can get this privilege without being in system partition.
Anti-abuse signals. Can we also get these signals? Jeff Bender someone they say
it's not possible.


## Guy.md on 2021-09-08


Rubidium - 18 HC for AASE
ad fraud detection on device
maybe something ODAD can be done?

Monir is AASE lead for Rubidium collaboration
Sudhi - ASAP representative for Rubidium
William might know better who's who in Rubidium eng.


## Niko.md on 2021-09-06

Do you know the muffin man?
Desk for Noriko
HC doc
Pre-callibration: upload to tool or share doc?
Packet
Meerkat weekly?

Desks, desks, desks!!!


# Week 35
## Jessica_Lunney.md on 2021-09-01

Doc in a good shape
Definitions good

We need to do testing of this condition or subset to understand how apps will be
affect.

Need to test appops for definition 2.

Need some form of check of existing Play catalogue just before launch.

OK to use data we already have to use this best we can

start socialising the rest of proosal

was talking to Guanhuan about settings state and not ideal

when ready to enforce we have to have a story why we ignore user settings

should we involve some outside experts on how we do this in a way that
increasing net user safety as opposed to creating crisis

Jessica to talk to few teams who might know who we can reach out to

unconsented turn off is not enough for us to turn it on, but informing user
might be alright

Jessica wants to see rough analysis plan, if we need to add more info

https://quokka.corp.google.com/q/#/app/overview/0819a8170ed641f535bc0e05b518523794f22b8f9aa9c967cff166bf1a499988/1630422466147289

Ask Steve to grab some time with Jessica

Maybe add code to snet to simulate odad?

Sync Guanhuan settings update

Jessica will ask around for expert advisors, to see if there are any red lines
we can avoid.

## William.md on 2021-08-31

Nikolay joined


# Week 34
## Karishma.md on 2021-08-27

Used to interview not now

Also GTI. Used to be in ZRH, but now it's virtual. Lots of volunteering
positions. 

Didn't have any inclination at start. After working with more senior people: for
the next couple of years really build on technical skills and then eventually in
a couple of years move out of just being a coding engineer, go into tech lead or
engineering manager at some point. Don't see myself continuing coding forever,
but wherever I go into my leadership postion, want to be fully confident with
tech skills.

ideal job: for now def enjoying and interested in writing a lot of code. I know
at some point I'll have it enough. 

Was looking at L5 requirements, design of larger things. 

Least exciting: with the work doing on Angel, struggled a little bit with the
pace. Would prefer slightly slower pace, so we can put some time in quality and
stuff. Some times was moving too fast to do the best job I can.

* For all of the design and implementation the time it takes set by the time we
  have rather than how much it takes.
* Did some work on GmsCore, had to do some changes and launch it. Totally
  unfamiliar with the process. 

Always felt supported by everyone on the team. 

W/L balance: issues. Have felt overloaded quite a few times through the project.
Being home helped. Covid makes difficult. 

Saying no - still learning. Can do a bit better.

After work - was a lot easier before Covid. Things have improved now as well.
Yoga. 

Things about the team. Fostered a lot of culture from Aleksandr. Tried to hang
on to that. Alks strong team culture, but it doesn't happen naturally, created
a lot of psychological safety. 

Angel tried to adopt some of the Aleksandr things. Started having some process,
sprint planning. A little bit more of that would help. eg in Aleks all work was
organised in sprints and we followed it religiously, in Angel never had chance
to do that. 

It would be nice for Angel if we can spend more time than we did before on
things like design. More documentation and stuff.

Started taking a look at ladder for L5.

https://docs.google.com/spreadsheets/d/1Bs15SgzD92VJ4j3kGuYj41V4rFOjBO6F0Xkq6x4aQtA/edit#gid=0

AIs:

* Try to take slightly fewer things on. This will allow me to do things a bit better.
* Restart doing some citizenship contributions.


## Guy.md on 2021-08-27

Speak to Mike please


## Andrei.md on 2021-08-24

Talking points for Andrei:

 - Angel launch going as planned
- Pivot to Mainline/GmsCore for Android T. Want to do a techtalk on either old
  or new Angel, TBD
 - Google Play Protect: permission revocation
 - Find My Device - conversations with Better Together.


# Week 33
## Peter.md on 2021-08-18

Niko is overwhelmed.

Reliable usage stats - Inara. Peter steps in, I will do.


## Aaron.md on 2021-08-18

v1 app is being finished
Couple of changes

Two models, one submitted, one not
- Heuristic model
- Need to add the ML model

Now they are finally both submitted. 

Both are based on app ops.

Heuristic model is targeting fishing, reputation model not sure. 
Better to ask Alice. 

Getting data back through federated analytics

Will start coming together next few weeks. 

The group is doing malware research
Aaron is doing priv escalation

William seems to be on top of things
Some cycles he considers more important, asks Aaron to write a Googledoc he can
use. Not a particularly stressful time for him. 


## David_Coffin.md on 2021-08-18

Worried about the mainline of the team

Elliot said with slower platform dev might not be a team for me

Concern about compilation time tooling etc.

Just making sure we can actually deliver everything we are committed too. Marie
is saying it'll take 4-8 weeks to get productive working in the framework.


## James.md on 2021-08-18



Barbet should get Angel

Angel 0day - what exactly should be in there?
Are marketing taking the right visuals? Video production is happening


Need a vision deck for purely security features. Krish asked before James leaves
to put together a story for the Safety Centre. Jeroen is reponsible on UX side. 

Everything should work with Safety Centre. They are going to work on go/bethere
vision (privacy team / Charmaine team). Sarah and some new PM. 

1. Some new top-level thing in the OS that looks like quick settings


Give users something to do there!!

Samsung - sync up with Unsuk.\

Permission Revocation!

## Niko.md on 2021-08-17

New manager orientation

Course is valuable, but a lot is lost because self paced


No doubt talented eng

Performing work at L4

Clear way forward to achieve that

Goes for promo. Packet looks solid.

Couple of particular issues:

1. Def performs as L4, but doesn't necessarily projects L4 attitude/mindset.
   Working with her on this since took over. She has gotten better. 



Tricky one. UXe.


## Steve_Kafka.md on 2021-08-17

kids 12 and  son/dughter


# Week 32
## Narayan.md on 2021-08-12

WHen talks about auditable. WHen we build things from on-device signals (e.g.
privacy dashboard) we want to be open source and auditable. 

What we are talking about (account compromised etc) not based on on-device
signals. 

Narayan's concern is that having a separate APK is a bad idea. Releasing is
hard, especially on OEM devices etc. All of these integration points are more
complicated.

Open sourceness and auditability - not absolute on this point. Some things
should be open source and auditable, some not. To Narayan it feels like -
awkward situation. Some things are open source and have to be kept that way,
e.g. privacy dashboard. Security dashboard should be based on mixture.

GmsCore has its own problems.

We either should fine a way to do in Mainline. Google boundary in mainline.
Non-open source code in Mainline. No technical reason why this can't be the
case. Demarcation that Mainline was meant for open source and gmscore for closed
source. 

Other option would be GmsCore.

Be very clear about what this is. Here's the existing UI, here's what we've
planned, here's what we can't make open source. Maybe have platform APIs backed
by GmsCore - acct mgr, etc. APIs can be in Angel Mainline module. Class of
issues entirely avoidable between versioning of GmsCore, API, Mainline etc.


## Karishma.md on 2021-08-11

WFH until 10th September

Things in flight:
- UX for profile dialogue. 

1PAAC

Coordinating zero-day QA.


## Rodrigo.md on 2021-08-11

Started to work on vision for Angel v2.


## Khawaja.md on 2021-08-11

Updates:
- Hiring
- Angel update - situation with Mainline and the implications.
- Conversation with Mike Tsao. 
- Jeroen

HC stuff

Total Angel - 14. 
5 heads for Angel, still not delivered. 4 for Narayan. 1 for Privacy. 1 Krish. 1
Charmaine. 2 Jeroene.

P0. 

all heads granted in H2, total for - 35. 14 were for Angel.

Bi-weekly quick update for Angel. 


## Inara.md on 2021-08-11

API cut on 11/08

Need reliable app usage stats

Android usage stats mgr is fairly reliable, but Peter did a lot of investigation
and there were quite a few bugs where data is lost.

Phonesky has a wrapper over it.

ODML warnings back at 1%.

Niko did a good job trying to understand what I want from my career plans, and
he made steps to ensure I'm working in the right direction.


## William.md on 2021-08-10

Nikolay's starter project
Stalkerware


## Anurag.md on 2021-08-09

Feedback - do come to standups


## Running_notes.md on 2021-08-09

Holidays in cornwall, swimming in the

BoqWeb migration

Progress been hindrered by distractions

Chatting Anurag every week

Normally are there any work things, blockers, etc. Perf stuff and what I need to
be thinking. Some things he told:

- Figuring out OAuth stuff - write down all tricky stuff as she's figuring out,
  as you'll forget it later.

Anurag is quite overwhelmed at the moment, skipping stand ups, but would be nice
to have him at stand ups.

Will be WFH until in the office full time. Was keen to get to the office but got
used to WFH now. Main reason why not enjoying - lockdown, weren't doing
anything.

Got rejected for laptop upgrade. Goes to NZ end of November.


Dave and Erik still talking to each other on how to continue.

David Lazarov team has expertise on protocol, reasonable amount of familiarity
with the platform. David's team has to do a bit more FE/UX than either of us
think is ideal. 


BA might help to monitor metrics of some project. Then they can run into some
challenges that involve statistics, sampling, etc that might take longer time to
solve. But these things may be beneficial to multiple projects.


Todd: devrel team, exposure notifications in the US, helping partners EMX
(easier version) Boston
Ran: DevRel, working with Appolo, EMEA focussed -- TLV
Michael: ExpNot, Covid-related apps, vaccine passport apps etc govt affairs ppl
-- LON
Aysha: PgM Appolo + Wear, managing multiple, MTV
Ade: run Appolo devrel

https://docs.google.com/document/d/19tHtZsma45pIB0CQnWvSNQ3xfsfTeDvDrwhCwesGUQk/edit?hl=en&resourcekey=0-VUIFXwJ6zXJlvspOj4rC9Q

https://source.corp.google.com/piper///depot/google3/production/config/cdd/android/market/publisher/submission_priority/apollo_government_group.gcl;bpv=0


## Running_notes.md on 2021-08-09

Main focus on ODAD. Close to finishing most of the work. The plan for the rest
of year - start looking for platform functionality for Android.

ANGEL

Do a little bit of 3P maintenance stuff. Bug triaging / bug investigation for
verifier. Some Finsky bug report gets generated etc, usually gets routed to Ben.


ODAD - tying up lose ends, changing app icon. Finalising string resources.
Adding some final safety killswitches. 

Before - did a lot of work on the infra. Very similar architecture as we used
for Gramophone in verifier, ported it to ODAD. Did a lot of work to actually
read the appops. Some complications there wrt how we cash the app ops, did a lot
of work getting the ODAD build set up, work closely with AiAi team, using their
build infra.

Team: we seem to go in cycles of trying to do some project plann-y agile-y
sprints, then this fizzles out, and then we try it again. Ben doesn't really
find any value in that but doesn't mind it too much.

Does this add value to the team? Ben doesn't have good enough optics. From Ben's
perspective it doesn't. 

Maybe it's an indication that it doesn't add value.

Spring planning happens during ODAD weekly.

William: extremely happy with how William managing the team. In terms of all
reviews (bi-annual, promo) William was top-notch. He puts Ben's desires and
well-being as a priority.

For day to day stuff - don't really have any complaints, I've been really happy
with William as a manager. Work really well together. His management style
meshes really nice with how I like to be managed.

Re constructive feedback - in general pretty happy with how it works.

Where I can help?

When we are pushing for new platform functionality, getting buy-in for ODAD. We
are getting a lot of push back on this. Guy helped a lot to work through this.
This will be the key in the near future.


## Bessie.md on 2021-08-09

JEtpack going on

Bessie was very stressed couple of weeks ago (flat move etc). Niko was great about this situation -
suggested to take time off, explaining that Bessie should take care of her
mental health. It was very nice of him to acknowledge that.

# Week 31
## Running_notes.md on 2021-08-04

Mainline - we are effectively changing AOSP code / packaging AOSP code and
taking control away from partner.

There are lots of Google specific hooks (GmsCore, Play, Account Security). gOTA
is also not AOSP. 

Is Mainline really good for us?

Mainline only aims to focus on AOSP code in the name of transparency. 

statsd collects data in platform but writes to Google through GmsCore.

If we use Mainline for Angel UX, we take away control from OEMs. The moment we
try to standartize something in Settings, we remove this from OEMs.

Concept of Optional Module. This is not a steady state. 

DocumentsUI and Permission Controller - only two modules with UX. They have
support for theming and all these things by mechanism called resource overlay.
We still control the UX, but we allow for aesthetic customisation there. 

Permission Controller - Eric who reports to Svet

Replacing Find My Device with Find My Mobile might be a nightmare from Mainline
perspective. 

In the early days of mainline, we found a way to push requirements so they have
to show in their settings Mainline update version.

If we decide to do it in permission controller, it will be easier. Find an
agreement with Permission Controller team. ewol@ zhanghai@ -- find an agreement
with them. 

From Mainline point it's just a feature go/android-feature-request. Drop down
"does this feature affects Mainline module". Then this feature run through 4
dimensions of Mainline approval (Eng, Product, Partner Eng, Testing).

One thing to consider is - since we are putting this feature into permission
controller - permcontroller is already out there. Can we make our feature
available on R/S devices? do we need platform modifications? Mainline always
relies on system APIs.

Perm Controller. APK. Most common type of Mainline module is APEX package, you
can see it as little partition that we materialise on device. In case of
PermController it contains APKs inside. In the APEX package we have some jar
files which are loaded by system server. 

System server is modularised. 

mathewi@ hansson@ while Dario is on holidays


## Sorin.md on 2021-08-03

Like vibrant environment, everyone is motivated, everyone has a shared goal

Defining characteristics:

* There needs to be a goal. Something that we want out of it.
* Next step of getting the team on the same level as yourself

Not a before market scenario.


## William.md on 2021-08-03

For Jessica - we are talking about consent for next year.

Screenshots pulled from Marmot, and then pipeline OCR.


## Peter.md on 2021-08-03

go/revocation-launch


## Running_notes.md on 2021-08-03

We can't ignore settings

If something is too dangerous to control, we must say it so. 

WHo are we cutting out of the market as the result of this action? We are under
pressure.

We need to get the data for user harm they might come to as the result of
stalkerware.

Inherent feature of GMS devices. Need to have a narrative consistent across all
partners.

Stalkerware is a good example. Is there industry standard definition of
stalkerware? We must take a position consistently, these apps shouldn't function
alongside GmsCore. 

Start with the problem statement, scope of impacted users, scope of user harm.
Proposal of several options of how we can solve this in Android T. Jessica and
Roman identify scope of interested parties at Google who must sign off on this. 

Should be very verbatim, specific list of types/behaviours that we are not
allowing. Define them in sort of in/out categories (stalkerware means this but
it doesn't mean this). The app that does surv can only be legitimate if it
notifies user every 24 hr, if it has this type of data but not that type of data
etc.

If we include legitimate use cases into the language, people who build surv
software may use this as a place to hide => we should give apps a way of get out
of scope of stalkerware. 

CSAM also comes to mind but this is very specific / negotiated with consumer
protection regulators.

Backporting is very possible from the legal standpoint. In PSIC it was easier.
Taking things out of user control might be harder. It depends on how we take
this control away.

Overall, worth taking on some risk. It's not about stopping people of suing us,
it's about stopping them winning.


## Narayan.md on 2021-08-03

If ecosystem thing, do not go alone. Angel is really not an app. There will be a
need to customise UI. Long list of safety issues with a standalone apk. Mainline
or gmscore is much safer.

Permission Controller and Documents UI have been updated many times.

There was a docs ui update to tweak some dialogues. Turned out this broke some
UI customisation on Samsung tablets. This resulted in a whole bunch of
changes/safety mechanisms. Mainline objective - quarterly launches. Current
mainline cadence is monthly with some conditions applied.

We have a candidate monthly build, nominated. Part of the mainline process -
there's a stage when OEMs provide feedback. Samsung/other OEMs can block the
launch if it's crashing. 

2 reasons why you wouldn't want to do it:
- All of the stuff is AOSP. We have google3 app but it should probably be in GmsCore.
- Have to find the right Mainline component. Is this a fit for PermissionController?

Doing it only for Pixel is fine, but doing it for entire ecosystem are different
tradeoffs. 

Both GMS and Mainline there are facilities to provide APIs which OEMs can use.

Pluggable UI that can be turned off.

Couple of mainline modules have same question about proprietary code.

Eng best practices are that gcore and Mainline are loosely coupled. No
guarantees device is in consistent state.

Broker -> GmsCore
UI -> Mainline

For UI we should probably go open source. It would be nice to give OEMs
reference implementation to OEMs even if they want to build their own thing.

Digital Wellbeing went through the similar situation. Samsung's DWB is backed by
our implementation.

We might be able to move UI as is to PermissionController because lots of
PermController screens are in SettingsGoogle.

For confidential assessment - reach out to Dario Freni, to get an opinion on
what the problems we should expect.


## Running_notes.md on 2021-08-02

ODAD is present on all OEM devices
ODAD is used to detect multiple types of threats to the users, off- and on-Play
ODAD is deeply integrated into Google's security offering
We can detect threats which are in principle impossible to detect in cloud/lab

SWOT

Strenghts


Why Andrei Popescu?


# Week 30
## Khawaja.md on 2021-07-29

Quick plan in preparation for mtg with Mike
Key deliverables we can achieve by getting more heads

Worst case - I don't have any heads, keep doing what you're doing.

What we can do with X people can't be done independently

What do I want get out of this mtg? What success looks for us? 

-> continue to maintain team and get additional HC

tactical importance: convince and influence them that they can't do without us.

Use Why FMD is difficult slides?

Daughter: research, robots, Boston Dynamics, Robots that can climb/run. 


## Karishma.md on 2021-07-29

Flying out 12/08

Then 4 weeks WFH then 3 weeks holiday

Language Gujarati but mostly English at home, parents hell bent on learning English

Coming back on 04/10

Working on at the moment:
- Building profile disambiguation dialogue, has been working on it for about 1.5
  weeks. Keeping an eye on time as decided to time box.
- Working on QA side - Karishma sent them new test plan, adds more test cases.
  Karishma authored test plan. Coordinating what gets tested.
- General misc Angel stuff. 

A lot of Angel misc work.


## James.md on 2021-07-29

All work on GPP consent punted

Make it an inherent property of a GMS device protecting user

Analogy: certified device / non-certified device. Users can't turn off
capability for us to know if device is certified. Tie this to certified?

Add an API to verify apps to read that state?

`isOdadAvailable` cf `isPlayServicesAvailable`

Special API if user is being stalked.


## Wa.md on 2021-07-28

Joined team last year. Previously had no experience with Android. Most important
thing is to pick up development skills for our team / verifier. Have been
focussing on that for most of the time. 

From this year we ... like malware detection research / patient z. Feels
learning these new skills, understanding how things work here, how to do
research.

Sometimes have to be more creative, to look for stats backing up the hypothesis.
Previous - data integration, from Google infra and our vendors. No research,
basically just implementation.

Focus - enforcement log. Way to collect on device, way autoscan is done. Collect
how many malicious apps uninstalled/disabled/warned and whether this enforcement
action has succeeded. Upload to server. Then can do research on this.

For example: found there's sometimes we fail to enforce every time. We look at
why. 

Currently working on ODAD. Part focussing - federated analytics. Part of the
processing. Fetch malware digests to Astrea, Astrea sends it to Brella. 

We are using AiAi team code for federated learning. Have to talk to their server
to make sure we are using their service correctly. 

Responsible from data entering Astrea all the way to model training, all client side. Brella is
handling training models, we just have to send signals to them. 

Local testing - everything works. Next step - testing on the real devices.

If team in the office, would like to be in the office.

Then Sundar email came in.


## Dave.md on 2021-07-28



- All platform work is done, no significant bugs found. Few issues we'd like to
  do differently though.
- Integrations with Screen Lock, FMD and Security Updates are done.
- In flight: GPP, Mainline, Google Accounts + caching and logging on top of
  that. This is happening as planned as we deliberately deferred this work until
  zero-day update.
- Work profile disambiguation will follow up shortly in Angel 1.1.

Had a meeting with Samsung, working on providing them an APK. They are clearly
very interested. 


Permission Revocation is on track. Blog post for developers really soon.
Roll-out in November, first revocations 3 months later.


Continuing to refine our stalkerware proposal. One of the points of contention
is that if we really want to be effective with stalkerware we should be able to
work when GPP is off, working with legal on that - there are precedents.


Backend migration is done. All traffic goes through new BE. Massive (up to an order of
magnitude) improvements in latency, significantly improved reliability, etc.

We are working on a roadmap which will explain how we are planning to get from
where we are to a future where FMD is a first-class Google product. Multiple
steps there, including ensuring FMD is preloaded on Pixels and can be preloaded
on other OEM devices. Immediate next steps - migration of FMD web app to the new
infra (as a prerequisite for redesign), much better metrics and fire fighting as
usual. Investigating whether we can delay Location History deprecation to go
straight to e2ee lkl.


## William.md on 2021-07-27

App release team. They want us to use PRIMEs metrics. Going back and forth, we
can't use them because can't use GmsCore. They said we can do our initial
launch. Philosophy is: we technically can't bind to GmsCore. 

Primes team - we finally got the on board. Now for some reason they are saying
that crash reports are sent through GmsCore therefore we want use to use
GmsCore. 


## Running_notes.md on 2021-07-27

Safety - device safety, app safety, physical safety.

Press analysis for Find My? 

Find My Alliance

Open API

Leverage ecosystem


London team is a product/user facing branch of Android Security and Privacy. We are
responsible for Google Play Protect, project Angel and Find My Device.



# Week 29
## Running_notes.md on 2021-07-23

Deep Strategy next week. Where are we going to double down on. Angel will get
doubled down on.


## David_Coffin.md on 2021-07-20

Unclear how we'll manage decision making while David is away. Can he pass it
over to Marie and Elliot? 

David will have a meeting with them Thursday afternoon to hand stuff over. 

Going two weeks to Norfolk, mother-in-law?

David: now I understand better how these planning meetings should look like


# Week 28
## Guy.md on 2021-07-15

Initially Krish was "what there's more to do". Guy said lets look at the presentation together. You don't describe what you gonna build, just vision - no specifics?

He sent an email to Erik and Sagar. He told them "what's your priority for this". They didn't respond. He said he'll try to get an answer from them. If they feel this is important, he'll ask to prioritise it.

Guy anticipates they say it's important but not a high priority.


## Running_notes.md on 2021-07-15

Хейвон
Shin-Chul Baik = Shin, based in Korea, part of security team, responsible sec update bug bounty problem, his group responsible for security ops / assesement, governments, privacy, 

Hyewon Park - partner eng team, Banseok's counterpart


## Inara.md on 2021-07-14

JetPack API design is still pending and in review.

Android API council needs a nudge

ODML warnings - we didn't agree on the design, there was no agreed timeline. This was the main mistake. William's team thought it would be much smaller amount of work.

Speak to Niko about it!


## 2021-07-13_08-10_Skip_Levels.md.md on 2021-07-13

device toll fraud detection, based on-device signals

launched heuristic model to detect toll fraud on device

heuristic model based on whether apps turn off wifi, it's a strong signal. if they open webview after that another strong signal. We check if the UI contains mobile country code / network code.

In certain countries network regulations are weaker so malware targets these countries. Thailand / Malaysia

App contacts C2 server to download code
Code might have JS to auto-collect ???
It will contact network to subscribe to the service - through WebView
Network will send a one-time password via SMS
App will try to steal one-time password either through reading SMS (30%) or with bind notification permission

If the app cannot turn off wifi, they can still try to

4-5% - baseline all Play toll fraud - probably not exclusive.


# Week 25
## Running_notes.md on 2021-06-24

For ASL social today:

go/ASAP-FoW


## Running_notes.md on 2021-06-24

Find my Network

Nothing is final until we have more information
Situation was: the was a reorg and they moved to Nest. Not a lot of interest
from anyone in the product until Apple launched Find My. The answer was you
won't get any recognition, if you want recognition you need to talk to Android.

Ronald Ho reached out, a lot of discussions and interest, after Apple/Samsung
launched their networks. David was working with Ronal for the last two month to
transfer all the knowledge. Got Erik Kay involved, all the information was
transferred to him as well. 3 weeks ago there was Android VP forum talking about
Find My network. David joined as a listener, there were only VPs. DaveK
represented us, wasn't just about FMD, but more do you want this project or not,
and if so, how can we fund it. 

Last week David spoke to Erik, nothing was final. Over the weekend - funding was
approved, doesn't know how many HC we will get, they are figuring it out with
the funding.

David is meeting later today with Tsao, there's probably some progress with the
decision. Want update on current state.


## Dave.md on 2021-06-24

FindMyNet: feedback yeah we have to do this. We have an ask for 20 HC total, of
which our team gets 4 or so for this year.


## Running_notes.md on 2021-06-22

Why not disallow getting IMEI?
Why not disallow disabling wifi?



# Week 24
## 2021-07-13_08-10_Skip_Levels.md.md on 2021-06-18

Siloed team. ML part is very very separate

Previously: image detection algorithm, scanning machine airport / subway


## 2021-07-13_08-10_Skip_Levels.md.md on 2021-06-18

ODAD deadline v1 - this Wednesday, didn't make it.

Get a model running for ODAD on daily basis.

Gramophone and other ODML work is getting deprioritised, disappointing

Uninstall signal. Install-time - time between first launch and uninstall

Which signal to create - very ad-hoc

Before ODAD people's work felt a bit siloed. Everyone on the team is working on
a totally different semi-unrelated thing. Right now everyone except 1 person is
working on ODAD.

William to improve: in the past he was micromanaging when Aaron was junior. Now
it's not the case. A lot of the work he's doing Aaron is not necessarily aware
of - maybe only him. But generally pretty happy with him as a manager.

For many parts of ODAD he has plans, e.g. ODAD might be a hub for other abuse
teams.


## Guy.md on 2021-06-17

Krish: he's not optimistic/happy with half-baked plan for 3 people that will not
promise a big shiny impressive Find My Device for Google. He wants something big.


## Khawaja.md on 2021-06-17

Talking points:

- Angel update
- Hiring update
- FMD conversation


# Week 23
## Running_notes.md on 2021-06-11

App Risk

Core model - generating holistical view of the score
Treatment part - how do we apply this score to different risk management

Revive uFBI to some extent. Putting minimum resource there. Better user semantic
score - means it can be used by anyone, reviewer, malware models can user score
etc. New BET - social listening. Want to apply similar language processing model
to social posts / Twitter / Reddit and see what happens. Only public stuff for
now.

Don't have Firehose contract. 

Infra projects. 

Metrics infra. 
Prioritisation infa - helping other teams to prioritise their entity reviews.

Collaboration with William's team:
- Toll fraud on-device signal. Will forward data on toll fraud to William's team.
- Any ideas how to better use these users' signals, we can provide customisations.

Getting a lot of important signals from on-device teams. Install/uninstall/quick
uninstall data. 




## Alan.md on 2021-06-10

Create a role for security hub
Make API available only for this role


Broadcast which reqs to be protected with permission only Angel has



## Dave.md on 2021-06-10

Ask: is ZRH remains a hub for ASAP, maybe a follow-up to Hiroshi's email next
week.


## Simon.md on 2021-06-09

Feedback was given to Krish on J

Karishma, career development, tears. Please can I have less work load.



## Coaching.md on 2021-06-09

Conversation with Khawaja themes:

* Great executive presence
* Good vision, strong relationships, particularly when comfortable with people,
  great vision, explain it well
* Emotional Intelligence
* There's an opportunity for me about how I want to re-brand myself. What brand
  do I want to present? How do I own the scope?
* How do I get two steps ahead, reading audience? Interacting with audience I
  haven't build rapport with? 
* Communication.
* Do I add value? Under-selling myself. How to start living with ambiguity of this
* Situations where you feel more comfortable vs less comfortable. Trigger situations.
* What is my brand? What value to I bring here?

How can we transform challenges into opportunities

Gap analysis. What's the north star / vision, what are the opportunities that'll
get us there?

What is the vision for myself?

80/20 rule: what problem are we trying to solve that will have the biggest
impact? If you want to have a bit of structure for yourself, there is some
structure around it in perf that may cover at least a part of it.

Locus of control

What is my role in this? 

## Giles.md on 2021-06-08

How do we check that the app is preinstalled?

What if the app has disabled the launcher activity?

Play Auto Install / Facebook auto install - is this covered / revoked?


# Week 21
## Guy.md on 2021-05-27

* ODAD update
* Package Verifier
* Hiring update
* Angel

Dave is good with keeping the team


## Dave.md on 2021-05-27

Samsung is willing to integrate, but they want have their own UI and devices.


## Peter.md on 2021-05-25

James: let's finish permission revocation by perf. Peter said no.

Target date is September to launch + 3 months grace period
December we start revoking permissions


## Narayan.md on 2021-05-24

Long term - UX privacy and security need to be merged.

Not sure T is the right release to do it.

In S we are giving partners a lot of leeway to do what they want

In T we'll have to do a lot of work to pull it back

first choice for privacy UX is to be rendered by Google


## Chad.md on 2021-05-24

ODAD

privilege escalation is a tricky one

let's say we have something that indicates consistently bad behaviour

spouse ware

chad loves SMS example

Thoughts: develop high-precision model to detect tracking spyware and ask user
if they know about it / OK with this. 

## Rahul.md on 2021-05-24

Aiming to launch ODAD
Working on Play Store app
Focus: get it out of the door

Risk: this is P21 scope with just one signal
Need to understand how can we expand to other devices

We need a hero case

Even if in 2022 we add more signals, Rahul is not super convinced we'll find any
malware on Pixel. We need to quickly expand to other OEMs. 

Part of Aaron's time was spent on few experiments

Malware on Pixels: no data with Rahul, but William has the numbers
Review with the Pixel team, Rahul to find slides.

Related: how often will we run scan? 

Easiest malware categories would be related to PII leaks and spyware. More
nebulous categories are trojans. Even if we detect something, getting it in
front of reviewer would be a big challenge.

Discussed content abuse really early. But the features needed weren't really
suitable.

Was very hard to get toll fraud approval, so screenshots no go.

go/odad-notes

https://docs.google.com/document/d/15nztqlQZ_wHcANSLOkGwTk8GfPfT4U5u6Lkkc4FwkBE/edit#

# Week 20
## Khawaja.md on 2021-05-21

Angel Android platform part is code complete.

Krish fear:
- if we work on FMD and safety hub lack of focus
- james doesn't want to do this


## Krish.md on 2021-05-18

- Angel update
- ODAD
- FMD


# Week 19
## Running_notes.md on 2021-05-13

https://androidpartner.devsite.corp.google.com/gms/policies/domains/enterprise
There is a UX one
(work profile badging)
Thanks!


https://androidpartner.devsite.corp.google.com/gms/building/settings/security/security-status-settings

https://androidpartner.devsite.corp.google.com/gms/building/settings/security/device-health-services
That may be more relevant
We aldo have to decide if we want to make this mandatory app
and what happens if an OEM re-implements their own thing entirely
What does the mandatort app do?
Notghing? Become a content provider?
Something else?


## Dave.md on 2021-05-13

Plan to land ortholite


## Running_notes.md on 2021-05-13

Talking points:

- Introduction. Been at Google, what was doing. Our failed attempt to do
  on-device malware detection, happiness about what's happening in this team.

- What I am planning to do. In short, nothing - watch, observe, help.

- What I need to know - engineering review of ODAD and other in-flight projects

- How I can help? 


## Jan.md on 2021-05-12

Problem with injected settings

https://googleplex-android-review.git.corp.google.com/c/platform/packages/apps/Settings/+/14468790


## Running_notes.md on 2021-05-12

- Given we have AppOps both in the cloud and on-device, can we train classifiers
  cloud-side and run them client side? If so, what are the benefits of doing it
  in ODAD anyway (with the obvious benefit that we detect bad stuff much
  sooner).


## Peter.md on 2021-05-11

Making some changes into perm revocation UI. Quite close and then go to dogfood
next week. Opt-in dogfood. Haven't built the part when apps can ask users to
grant permissions.

User education will be the key going forward. Show info cards, "here's what we
do for you".\

1:1 with Niko everyone month.


## Anurag.md on 2021-05-11

http://doc/1dnhJSKS3bFvdXS9NmBVj99VJVRXZN_go26HqI79SZlw


http://doc/1F3VC5d8HBLmKjOyQwjEatLuWzG6czw96AFwQGmqEinE

Jim back to FMD on 19th

got side-tracked by Android S settings work. Anita is leading the effort,
Isabelle and Sorin are helping.

Anita to pick up FMD turn on rates. Anita feeling pretty good with GmsCore part.
Took ownership in Feb/March but going in the right direction, growing in
confidence.

Jim got EE, quite happy with his rating. Plan - get him to run larger things in
FMD. Has a tendency of getting involved into small one-off changes, but for his
own development he needs to take on more ambigous things. Took on BoqWeb work,
lot of design, ambiguity. Jim and Isabelle are both working on this project, Jim
is doing more leadership and guidance.

Isabelle looking solid. Ramp-up on FMD already started to exceed expectations.
Can-do attitude. Plan to get her to do more design in BoqWeb project.

Boqweb: Timon is the server-side part. BoqWeb involves changing the JS and HTML
and the Android website to be served by Boq, refactor of JS on the website, 3K
LOC file which is impossible to maintain. We can still use all the Soy template,
some rewrite is involved. No UI refresh per se.

Sorin - paternity in June. Girl.

https://docs.google.com/spreadsheets/d/1qV8Te40c0JG-zrwo-mqWh1ezvc1VP9wgHJq6yg4KjMY/edit?resourcekey=0-eLzivE4_P5kQckLDbpceUg#gid=0


## William.md on 2021-05-11

- Reorg happened. Now what? Let's meet the team? 

We should still keep focussed on ODAD and get v1 out. 

- Maybe add me to the team chat? 
- Branding: meerkat? new brand? 
- kirillov-direct - need new time.

anything that goes into the app sandbox. PlatSec: it has to be an API.

Guy: can we expose some signals via appops? Then this is a sort of official API?


## Chad.md on 2021-05-10

Some stuff in logcat 
Didn't get VRP
went to press, do you know everyone can read logs???!
caused huge hype BS
dutch health minister
activities name show in logs when activities start

ODAD: privileged A/V - many people are not fans of
digital markets act in the EU - requires open stuff up, big deal for Apple. EU
legally enforced platform parity. We can't do things that benefit first party
applications, especially for ODAD will make greatly complicated.

we don't do install time scanning.

Section 230 - Twitter not liable for shit users post as long as they make decent
effort to remove the content.

things to use:

* Powers we have around install-time. Why don't we drive that well?


## Running_notes.md on 2021-05-10

Some stuff in logcat 
Didn't get VRP
went to press, do you know everyone can read logs???!
caused huge hype BS
dutch health minister
activities name show in logs when activities start

ODAD: privileged A/V - many people are not fans of
digital markets act in the EU - requires open stuff up, big deal for Apple. EU
legally enforced platform parity. We can't do things that benefit first party
applications, especially for ODAD will make greatly complicated.

we don't do install time scanning.

Section 230 - Twitter not liable for shit users post as long as they make decent
effort to remove the content.

things to use:

* Powers we have around install-time. Why don't we drive that well?


## Running_notes.md on 2021-05-10

Picking unified brand - Google or Android?


# Week 18
## Running_notes.md on 2021-05-07

Meerkat Weekly Agenda:

- Reorg announcements (Alan, Roman)
- Anita: demo of new FMD GmsCore settings in Android S


Why not platform attestation?
Doesn't solve droidfooders issue. Droidfood use unlocked bootloader, or
pre-release hardware.
Also key attestation - availability and reliability issues.
KEy attestation only addresses one specific way of tampeting
go/key-attestation/sn-dashboard
wouldn't address post boot compromise etc

checked /xbin/su with Narayan
He came back and it doesn't fit with the long term direction


## William.md on 2021-05-06

Flubot - ortholite to trigger FTM and show warning

Family in Ontario

Ben might want to work from another state, Oklahoma etc.


Resources - so many different projects we are working on, have to de- and
re-prioritise things over and over again.

Things without owners:

- Vole. Absolutely no owners at all. Last time William had to do migration when
  service was down when someone pinged him. Being relied on by:
  * Identity team
  * MDMs

- SafetyNet APIs. Droidguard team is 100% owning attestation API. John Ayres

There's a list, William to share.

Resourcing - research skills part is missing. Likely to miss Android T. If we
don't get resources might miss the next one.

Framework signals - actually implementing them. We don't want team to switch
from app to Framework.

Detection side too. We have to build detectors for signals we get back. Damien's
team? They only provide guidance, we have to do the work. Xun does dynamic
analysis, we work with them, but they won't build models for us.

This is likely to be the right model because these two teams report to different
people, so they might not prioritise this as much.

Dianne didn't take any sides in ODAD meeting.

Privacy, Legal. Made a mistake, privacy and legal approved for a while
everything we went them with, but with ODAD we were slow and Chad went to them
first. When we went to PWG they were very different. PWG person - Giles, he
doesn't like ODAD. Pauline is the person now. She signs off if Jessica signs
off.

Someone joins as L4. External. Some Android experience. Did PhD in ML, did
server work for 2 years after PhD. End of June. Ishtiaq.


Enjoy doing various different things. Like doing different things - code,
design, meetings with people, collaborate, etc. Some people want just code
whatever, IC is perfect for them - but not William.

Career path - went 5->6, took 5 years to do. Want to start thinking of moving to
the next level. IC ladder. 

Mgmt style: look at each person and cater to this person's personality. Some
people need more pushing and nudging to get them going. Some people might not
like that at all. 

Guy did skip level convo.


## Khawaja.md on 2021-05-06

https://docs.google.com/document/d/16QAE8Eh8ThriuNmQc1KfkIudKHdZuhq4d-b2G50Oy1M/edit?resourcekey=0-NA-zC8Xonil-DYgGJ7_B4g#

carreer grows comes from: 70 (work experience) /20 (academic) /10 (learning
outside)

6 buckets of skills:

- management (lat. Manus, handle things)
- leadership (mentoring, setting up the vision, building coalition)
- ability to create inspiring visions that attract people, that motivate people
  who work with you and around you towards the goal
- competency, ability to come up with strategies. sharp understanding of what
  makes a good strategy. Strateg for Angel - how to set up right objectives,
  what are hurdles
- how good are your technical skills
- business security. 

where to spend time: 
- 70% -- things you are really good at
- 20% -- grows areas, things you are not competent yet along the 6 buckets above
- 10% -- grunt work. 

Good / Grow / Grunt

RK: grow - stategy, stakeholder mgmt

grunt - e.g. writing perf

grunt work - fixing bugs which can be done by L3/L4 -- setting some examples in
organisation what matters. Uptime matters, code quality matters, etc.

**important**: picking, designing framework and benchmarking yourself according to
that framework. set up some benchmark, as long as you have a way of measuring
progress.

RK: how manye people stakeholders spoke to.


## Running_notes.md on 2021-05-06


* You all know what's happening
* Terrifying that something like that can be possible in 21st century.

If your parents or relatives are affected, there are few things you can use:
 
* Telemedicine option - as far as I understand, this is very much like Doctor
  Care Anywhere we have in the UK, but incl Covid related situations.
* Personal insurance for parents - private medical insurance which covers Covid.

Both options paid.

Finally, go/support-india - Googlers have already raised in excess of £3M and
all donations are 100% matched by Google. Likewise, there's also go/support-brazil.


Lots of speculations, but here's what we have in the end:

* Default option: flexible work week, 3/2. Days in the office are collaboration
  days, decision on which ones are which was pushed down to PAs, so I am hoping
  to hear more updates soon. You can also work from the office more often if you
  want to - but it remains to be seen what will be the situation with the office
  space, as now we can effectively "overbook" our real estate.
* Global work: means you can work from the office different to the one where
  your team is. I think this implies that office must exist in the
  country/location of your choice, because if it doesn't it's
* Option 3 - remote work. Basically, like right now but anywhere else in the world.

Note that if you are moving geographically then your comp will be adjusted per
new location comp.

Hybrid work will begin when LON reaches Level 2, but RTO is voluntary until 1st
of September

In addition to that, with managers approval you can work 4 weeks a year from
anywhere in the world. The way I read it, for example: you can go visit your
family back home, and stay there for extra 4 weeks, working from wherever you
are, so you can spend in total 6 weeks with your family, which doesn't sound all
too bad really. This replaces 14 days policy. You can still apply for exception
if you need to work longer.

Focus time: basically like P&E no meetings weeks, but for entire company
(not immediately clear if they will be aligned acrosss PAs so we'll have to see).

Some really interesting questions on TGIF Dory 
https://dory.corp.google.com/series/101837965?sort=top_resolved
see another email from Sundar


## Guy.md on 2021-05-05

reorg early next week

speak to Sebastian

CAT folks, Sudhi, presentation on Rubidium, ad fraud detection on device.

can we combine it with ODAD? same app / separate app with shared code

triggered conversations between leads

William talk xinzhao@ TLM ad spam PM zacharylf@

Proposal to ML steering committee. Project proposal about ODAD, happened about a
month ago with CAT. Joint project ASE / CAT / PlatSec. 

https://moma.corp.google.com/person/blaisea Cerebra / AiAi

Bunch of teams jumped aboard, generated significant # of HC

csam - child sexual abuse material, child endangerement basically, CP.

proposal: we want to build some common infra that teams can leverage.


# Week 17
## Coaching.md on 2021-04-30

How do we understand the team culture??

Interrogate William?

talk to them how I lead the teams

what is it that I want to hear?


## Khawaja.md on 2021-04-30

Communicating goals - broadly or per-team?

Pass them to prepare presentations for their own teams, let them develop goals.

what do people consider actionable feedback?

97% PMR

how to deliver negative feedback

book: letter to the future exec (??)

TODO: set up mtg for kshams with Frances!!


Security team with making their case. How do they negotiate. They are doing name
dropping (bringing up Dianne etc). 

It's been hard for William to find good talent. Someone who's good with Android.
Need to uplevel the team a little bit.

Retain, attract, external relationship.

Least resistance env.

Bring DeepMind, CAT, AiAi together. Make it big xfn thing.

Who/what/how frequent - network/stakeholder mgmt

what are the attributes people will be looking for at the next level?

write down doc: challenges that come with ODAD, 3 challenges above. Efficacy,
retention, global talent mgmt.

Doc (working doc)

1. Stakeholder Management
2. Leadership/Impact/Tech contributions
3. Results!!!

what would be my strategy to influence Sudhi's team / AiAi team etc?

kernel of strategy is 3 elements:
 - diagnosis
 - guiding policy
 - coherent action

good engineers disagnose the problem.

guiding policy  you know what you can and cannot do

coherent action - execution.

60 min to save the world - 55 min understand the problem

vision - win the war
mission - win the battle

how you solve challenges is a strategy

what Nokia CEO did wrong? Lost the big picture


## Dave.md on 2021-04-29

Convo Erik Kay - he's in exploratory mode, what Android should do, going after
Apple, preventing fragmentation. Apple want to work on Find My
inter-compatibility thing.

birds/dogs - branding exercise. 


## Brandon_Barbello.md on 2021-04-29

Privacy and security for Pixel phones.

Been at Google 7 years
Before research space, then bridge research to product.
Shipped number of on-device ML systems.
Face unlock, Soli, Ai, sensors. 

In the recent reorg came from hardware side Pixel to Sabrina side to form a
bridge. Maintain responsibility for on-device AI with sensors, but was also
asked to lead privacy and security.

Been trying to craft the right scope of this role.

Trying to serve in non-duplicative way. Plenty of folks building great features.
S&P are things that we build for platform are win-win for Pixel.




## Anurag.md on 2021-04-29

Sorin/2FA. Anurag wrote a bunch of options on how to deal with it.
Build it into FMD product as ring tokens. Use Play Family.

Digital Car key revocation. 1 sprint of work.

1P on e2e device tests for FMD.


## William.md on 2021-04-27

No official TL at ODAD
PM is Raz


No desire to go to L6. Happy at 5, happy doing IC work, challenging work,
getting toughest problems. 

For citizenship - very introverted person. Things like interviews freak him out.
He's doing open source things that help other Googlers as well, for citizenship.

Extremely shy. Not very social. Doesn't want to talk a lot during meetings. 

Was ignoring William for the first year despite sitting next to him. Walking
halls looking down. Once you get to know him, you can have conversation. Gets
very stressed talking to people he hasn't met before.

He cares for: all things challenging. Coding something in Assembly language to
solve some puzzle. Rewrite his drivers. Made a keyboard on his own. Only
tech challenging, zero org complexity / leadership.

His vacations are mostly at home, he doesn't go anywhere. Doesn't like
travelling. Likes it very dark as well. Had an incident in the office when
somebody put all the lights on.

He's great with WFH, all lights off, cave.


Doing really well now. Challenging problem. Interested in going to L5. Stepped
up a lot in leadership, most in William's team. Helped to break down all the
tasks for ODAD. Flow diagram of the data. Created spreadsheet and asked people
to sign up for that. Really tries to go to the next level. Volunteers for lots
of things. 

For L5 he can assign some bugs to someone else. 


Assigned her to all the ML stuff. Done that part now. Asked her what she wants
to do next - she wants to learn new things. Put her on ODAD, Android part. 

Taking some of the easier tasks, Ben and Aaron take more complex ones.

Giving her work that is related to ODML stuff. 

ODML is for off-market, which was deprioritised at app safety level
ODAD for cloaked Play malware.


She was working on enforcement latency previously, PatientZ. This has wrapped
up, just a few bugfixes. Put her on ODAD. Works on analytics part with William. 


Not working on ODAD, works on on-device toll fraud detection, very important at
K^2 level. Not high volume detection but found new ways of doing toll fraud (acc
to Unuchek). 

Intel desk numbers going down where we are going up, probably took some of
detection. Previously did some simple heuristic, now working on ML model. One
problem is a lot of false positives.


## Narayan.md on 2021-04-26

MPlatform
Angel

Bluechip fires:

- Permissions Hub, some changes to be made UI-wise, not risky per se, but risks
  that some app/usecase that is doing something bad and will cause a lot of
  negative publicity. Fi is accessing location once every 5 min. But what if
  someone is accessing location once every 5 min and they say because we do
  because of your shitty API? 


## William.md on 2021-04-26

Think longer than P21

2 other approval processes: 
 - Pixel approval process, separate from Android approval process. 
 - Play update approval process.

ODAD and AiAi cannot talk to GmsCore. The challenge is that if we can't talk to
GmsCore, it's very hard to get system health data. They want this for Play
launches.

AiAi got away because they launched with GmsCore support, and then they took it
away later.

They want to see these data for initial launch.

We are trying to get mobile harness data but they are pushing back, saying we
have to get as much as possible from Primes (built into GmsCore, feeds into
Rasta metrics).

They are: Play Launch team.

Pixel launch team they want to track whole bunch of things for the launch. They
thing feature complete supposed to be May/June.

Workstreams:

[x] Finished a lot of streams, mostly focussed on ODAD. ODML card warning. Closed.
[x] DCL analysis - trying to see if there are any anti-cloaking anomalies. An app
  that is malware will DCL a non-malware app or other way around. Closed.
[x] Patient Zero workstream. Based on ODML. Closed.
[ ] Toll fraud workstream. Open. Finding new variants of toll fraud.
[ ] ODAD. Open.

Within ODAD there are multiple workstreams, but mostly focusing on app. Should
also be research into future signals (William working with Xun/Damien), but
needs real owner/TL for this. Then there's implementing these signals in
framework.

If new hired from Google, work on app.
If externally or minimal Android experience, put them on the signals.


## Guy.md on 2021-04-26

MOTORCYCLE KTM DUKE

- Management style
- How often to check in
- What do I need to know?

- Talk about the team
- Key people
- Periodic needs etc

Offer team alternatives

GPP: FE and BE and infra and ML - think more holistically, do better. Set a goal
of having people work together, single roadmap, what would be the most effective
GPP all together.

FMD - break out of migration mode.

William - spend some time understanding what he and his team are working on and
understanding what he wants, his aspirations. There's no need to make any
suggestions / change of plans, and think about it.

Asking what his challenges and how can I help?



# Week 16
## Khawaja.md on 2021-04-22

Topics for today's conversation:

1. Update on hiring
2. Update on Meerkat Platform transfer
   * Transferring lus@, rickywai@ to Ashwini with HC, brufino@ without HC
   * Jan joins Angel

Stefan and Sebastian promoted to L8 and L7 respectively


## William.md on 2021-04-19

Working on ODAD, using AiAi which doesn't help a great deal

Looking for 2 people to work on ODAD




# Week 14
## Simon.md on 2021-04-08

Pending Intent situation

Dave was poked in the eye by child's unicorn. Can't work at the moment. He can
do GVC. We can pull information out of him but we can't expect him to produce
code or documents. 


## Khawaja.md on 2021-04-08

Meeting the team - was brought up multiple times.

Living in an anomalous time. Multiple people who got really sick. 


# Week 13
## Dave.md on 2021-04-01

Was never comfortable about having 2 platform teams.
Limits somewhat the impact of Platform team.


Security Key


## Running_notes.md on 2021-03-30

dbenisch@ 

they are responsible for red teaming the APIs, Derek leading this.

Ashwini = responsible APIs

Long term API foot/gun issues

let's say malware abuse. it depends.


## Jim.md on 2021-03-29

Really liked session at ASL social on mental wellbeing

Asssitant bug for device name - didn't find the bug.

PoC at assistant?

File a bug.

Son off for easter

FMA migration to Boq node and to Timon

Sitreps, ToS and FMA are remaining items, sorted by qps

sitreps mostly coded up, adding some instrumentation to compare old BE sitreps
and Timon sitreps. They both write to same Kansas table.

ASE techtalk on migrating live traffic?


# Week 12
## Sorin.md on 2021-03-24

New daughter due in June, age gap 6 years


## Running_notes.md on 2021-03-22

Ralf and Paul Mcclave, reports to Ralf.

Multi-account must be handled


# Week 11
## Nandini.md on 2021-03-19

go/fmd-lkl


Bypass for ELS

Before P, 50/50 on location enabled
New devices overall 85% have enabled

Overall 70/30

What are the cases when location is not returned

Ask Jen Chai about Folsom, cred on the web thing

Better Together product review series - present Find My Device from Better
Together perspective.


## Monir.md on 2021-03-19

go/swe-l3-to-l4
go/swe-l4-to-l5
go/new-swe-ladder-table


## Simon.md on 2021-03-18

David settled in into the role, feeling more confident
Karishma - been encouraging take a lead a little bit in Karishma/Raissa duo
Raissa is now working more on core GPP stuff
Peter, Inara and Bessie are on auto-revoke


## Running_notes.md on 2021-03-18

- Go outside for walk more often, but I'm too lazy.
- Photography prompts - take a photo of something at 8am and post it to the
  group channel. Random Photography Prompt
- Meeting for a picnic. Totally not work related
- Fake commute. 
- Aleksandr games nights - one Friday evening a month, even over video, really nice.
- Fun events. Helps to mark time. 
- Take a long walk every weekend.
  

At MSFT whenever they have team event, they give £20 Deliveroo voucher so
everyone orders food and eats it together. This might go along with next team
event - not the same as eating with other people but close.

New restriction on international WFH. 


## Coaching.md on 2021-03-17

Mental wellbeing case study

Groups 4-5 people are more likely to share

Talk about wellbeing scores low

Put into employee shoes
Is this solved in one meeting or a part of many conversations

There are many reasons we are feeling in the way we are feeling

First breakout. Something in introduction to this breakout, ground rules, how do
you create safe space for these people?

- Whatever happens in breakout ..
- No wrong answers, 
- Not here to offer any solutions, just here to listen to each other

One person comes back with high-level themes to share with the group

Questions:
 
- What is the biggest challenge you are facing? What is affecting your wellbeing
  at the moment? Get the key thing of their chest.

Keeping things pretty simple.

> What is the biggest challenge you think we are facing at the moment?

1. Intro
2. Breakout
3. Come back / discussion
4. MSFT research
5. Having heard about those themes, what can we do about it? 
6. Assign one topic to breakout group, ask them to discuss it
7. Come back, share findings



## Peter.md on 2021-03-16

Data issues with perm revocation - system can miss some usages. Can't fix it to
100% accuracy but can make it more robust. 

All local.

Works up to Q, then if the user updates to R, we'll migrate things to Settings.
How will we dogfood it? 

go/permission-revocation-timeline

https://docs.google.com/document/d/1y6p-CIA9UYH1d4MZlvZP5voz_swUfn59yqf0SvCRE18/edit#heading=h.c0uts5ftkk58


## Anurag.md on 2021-03-16

To discuss: TVC for EngProd work

* Old backend - lacking automated testing, some unit tests, 0% integration test
  coverage. We have to run a full suite of backend testing, 2-3 hr of work for
  Muz about once a week.
* New backend - good test coverage, ~90% unit tests + integration tests for
  every user visible flow. Do not run manual tests at all, push on green right now.
* GmsCore - have to run through manual tests, requirement for GmsCore, every
  week. Test tracker. 3-4 hr every week.

Then there's app and the web. Web binary releases alongside old backend, so old
backend implicitly tests web. For app, a separate suite of tests, pretty
infrequent, release once a quarter.


## David_Coffin.md on 2021-03-15

More relaxed and less relaxed?

Not at the point when overwhelmed.

Current work streams: David:

1. Get a stub preload AKP generated, TikTok app. Getting confirmation from Linus
   and Karishma and LGTM from Simon about what permissions we need and whether
   we need a particular role.

2. Need to have Rapid build set up for the app. Literally create a stub, won't
   have do anything.

3. Get signing key sorted out. 

All fairly administrative.

Once we have this, it'll make it easier for Marie to work on Settings changes,
once David is done with that he'll join Marie to help with those changes.

Still getting a lot of questions from Karishma and Raissa, dealing with them.
Balancing act. At this point and amount of time we have everyone should just get
on with it.

Interaction with Simon are much better right now.

Prefers independence, getting support he needs (?)

Raissa been talking about trying to develop some UX components we'll be using
within Settings app. I've got an impression that he got mild preferences if
we've got 1 developer from Aleksander it should be Karishma.


## Ed.md on 2021-03-15

T planning - nothing happened yet.


# Week 10
## Khawaja.md on 2021-03-12

areas of Stefan and Monir are further apart from

roadmap of xfn work. Things will change, orgs will change. If you get 50% of it
done increases chance s of going to the next level significantly. 
  
Keep focus on rating but also what manager need to do to help people to develop.

Identify opportunitites for other people. 

build the types of things I do for these ladder attributes over next several
years.

help across the org. Start influencing ASE, outside, etc. 

Citizenship is another area - participating in D&I activities.

What can be achieved in 2 years?

In 5 years too many things will change, 2 years is a right timeframe.

everything to learn. 

Start networking with other orgs in London!!!

Google gives $10Ks a month to non-profits for advertising

How to organize, manage, rally people, converge them on one space / set of
things you'd want to focus.

FACTS - focus, alignment commitment tracking stretching.


## Tadek.md on 2021-03-12

Interesting things:

- Core PA. Multiple pillars. Recently there was a reorg inside the user team
  with the security team joining us. New pillar that called PSS - Privacy,
  Security, Safety. Identity, CAT, Security are all reporting to Rayal, VP. Very
  good thing. It's a good move as these teams will be working even more closely.
  Good people to speak to -- Mark Risher. Mark is now PM lead for all of PSS.
  Fairly good understanding of the scope and what's going on in the PSS. On the
  eng side - Reza, he's not an expert on domain, but Mark is better. Mark

Is any interest in expanding to London. I would poke Mark Risher's brain on
that. He's also product lead for all the security teams. Maybe he has some
thoughts on LON stuff.

Other thing to think about is YouTube. In ZRH it's pretty big. Reto is YouTube
Ads. Maybe someone you might want to talk to is Jeremy Doig (jeremydo@). They
are hiring L7.

# Week 9
## Anurag.md on 2021-03-04

SpO2 normal, seeing cardio.

BC - estimate impact GmsCore, make commitment. Did audit of security module
impact. Medium impact: need to estimate what'll happen with foldables. We
display lockscreen and need to ensure it works fine on large screen devices.
Notification trampolines. GmsCore exempt. Also FindMyDevice module - Spot stuff.

PermHub, FMD req location in bg.

Asking apps to adhere to Material Next. 

GmsCore bug: vulnerability, someone can draw over FMD settings screen in GmsCore
and prevent user from turning FMD on. Alan and Ed feel this is not a real
vulnerability, can be done on many GmsCore screens. Fixing it in R and below is
impossible. In S there's a difference permission, will make a case.


## kshams-leads.md on 2021-03-04

kshams is -1 in Pez but reality is worse because backfill (?)

difference between target and current

target is not the right to look at, right one is target + whatever negative
allowed


## Simon.md on 2021-03-04


Doesn't seem to be doing much
Somewhat similar with Raissa

Simon to speak to David in detail



## Dave.md on 2021-03-02

We haven't received heads but were told assume they are there.

Zuul is using architecture Folsom, B&R. Intended for Chrome sync. Is it generic
enough? 

We have ability with Folsom to escrow secrets. People building solutions around
it, different teams building solutions that use it. We are adding dependency on
another teams' product that might be thrown away. Dave would like to see a plan
of making it really private.

He asked Scott to put someone on trying to look at our previous API safety
plans, just to make sure we have a plan. Dave promised we will have equivalent
plan. That might mean that work is spread out more. We have a lot of platform
security works working on it, not clear if we can put this to London. Important
for me and Sudhi to be talking to Scott / someone in charge, Sean(?) to make
sure we are doing what we can to deliver on theoretical promise of this. We are
doing training, investing into red team, investing number of SWEs. We can
already take credit for it but want to make sure we deliver.



## Narayan.md on 2021-03-01

Advise: content providers are usually terrible choice for any stable APIs.
Narayan's proposal is to use gRPC or AIDL to define API, gives slightly better
structure of API, better to write tests against. 

Worth thinking of what system APIs we want sooner rather than later.


Code red?

No changes in plan for PermHub. Narrow question on whether GmsCore should be
hidden from permhub. Frameworks stands for it should not. 

GmsCore's point it should be considered trusted. There was APPS last week when
they took it to Sameer and he said you should fix it, even if it means toning
down GmsCore functionality.

Problem is that their accesses are not always explainable. There are also
accesses that are not legit / not easy to explain. 

4-5 problematic modules.

We are still hiding location provider etc, not affected by code red. What would
change the work on reworking dangerous permissions will be accelerated.

Went through same exercise with AGSA, it was fixed.

Xiaomi is going ahead with their own permissions hub for R. They are even
willing to take our patches. Xiaomi seem to be very privacy focussed. 


## David_Coffin.md on 2021-03-01

Reached out to Settings folks, had convo on Friday. 

Their TL didn't really seem to care, got an impression we are talking to a
different calibre of team compared to others we interacted at Google.

They were very relaxed.

What if we need to submit code today?
They said submit CLs to me, if I don't understand it I'll redirect to someone on
my team who does. 

What about OWNERS? David will figure this out.

Simon asked if we can add gRPC code to settings, if that's possible. They said
submit a CL and we'll let you know.

Expectations: I'm getting clearer expectations of what's needed from me. I
suppose there is a difference between high-level expectations (I want you to own
this part of project, make it happen) and more low-level stuff, e.g. what kind
of work I prioritise. We are resolving that. Sometimes difficult to understand
what's my responsibility and what's Simon's responsibility.

We gone through shared expectations process. Really interested to see what kind
of perf feedback I get. Don't feel like I'm getting much positive feedback, and
that's been an issue for me. I think I'm lacking steers I'm looking for.  Expect
more career coaching / mentoring. Maybe share move personal examples on "this is
what works for me". More practical guidance on how to make things happen. 

Maybe Simon just doesn't have the time at this point of the project. 


## Jim.md on 2021-03-01

Timon caveats:
 - All actions except wipe and rename are going 100% to Timon
 - Rename tried to ramp up last week but ran into issues
 - Wipe - we ran into problem we are missing monitoring, now we have monitoring
    but tweaking it. Wipe is a very low QPS. 
 - List is also going very soon to 100% (should've been today but resources)

Two directions:
- App/website to device. Once wipe/rename/list done all of this will go through Timon
- Reverse direction we've got some work to do on sitreps - status update we do
  from GmsCore to our backends. We have to know FCM registration ID of the
  device in order to send messages to it. FCM doesn't quite support GmsCore
  properly. 
- Also a bunch of very low traffic things, maybe some have to be killed or
  migrate. Internal logging in web frontend (should be replaced with Clearcut). 
- FMA is going through the old backend too! 


# Week 8
## Giles.md on 2021-02-26

Lock screen stored in trusted hardware only

Giles becomes "an Ads employee" in 6 weeks
Rubidium -- Android version of Chrome's Potassium, we get rid of the AdID,
Chrome gets rid of 3P cookies, we are getting rid of AdID, replacing all ad
functionality which was built on top of AdID with privacy preserving APIs.
Ribcage is happening because of that. 

HC request 600. Sundar said this needs to be solved in 2 years tell us what you
need.

Sensor rate limiting.

Delphi rolls out to the rest of Google, 86% of Clearcut annotated to P&E,
rolling out to the rest.  


## Simon.md on 2021-02-25


David is not doing that much useful work. Not clear if imposter syndrome, or
imposter. Who will do his perf?

Spent a lot of time working on high level design, would've been useful if
produced in first couple of weeks but at this point not useful anymore. 


He's been really useful and helpful as an expert-consultant; he hasn't been that
helpful on answering main question: what API changes do we need to make this
work. Simon to deliver feedback to Alan.


Niko in good shape for promo. Unclear if he gets it, maybe come back next cycle
because not long enough

Marie in very strong shape. Opposite case. Doing for ages, question: she's been
SEE for so long.

James. 


## Alan.md on 2021-02-25

Photos latency down 33% cold start!!!

Fixit!!

no code in data directory


## Coaching.md on 2021-02-24

Speak to Tadek, Reto

WHAT ARE GOALS and aimes, how do we connect to them

Reverse eng google's vision and explain how do we roll up to it.

Users trust their devices

5 whys - we are doing this, why, because that, why ...

how does this tie to the bottom line? rep risks??

posthumous vision without change in synthetic. with change, it's driving.


## Anwar.md on 2021-02-22

abuse potential for new APIs?

there will be update to permission controller this year


## James.md on 2021-02-22

Core metrics

GCS - name Google's security solution


## Khawaja.md on 2021-02-22

HC, vision, what not?

For L8:
- Give it some time
- Impact at Google OKRs, peer promo L8+
- Drive team performance (?). Less about individual projects, more about thought
  leadership. Team improvement. 
- iteration...

How are you removing barriers for your team. 

What is thought leadership??

Vision for the team

HBR what got you here

thought leader team performance 

where do we see role of GPP/Angel/FMD.
User-facing responsibility.

A lot of attention to Google

One big role is information. 

creating wel informed cosnumer giving precise info at the right time


# Week 7
## Simon.md on 2021-02-19

James - do we want to put privacy in Angel? He doesn't really provide a
direction. You ask a question, he'll ask it back to you, direct to Rodrigo, etc.

Core metrics - decided to build ourselves, without asknig James because he
doesn't answer.

He considers GPP day-to-day too junior for him.

SBI - situation behaviour impact


## Alan.md on 2021-02-18

- Good news - one of the ideas of perf team fixes the problem
- Bad news - figuring if it works 100% time is hard
Race conditions.


- [x] if you press home button, it kills all BAL for all apps
- [ ] wanted to do - current problem that if you are visible app and fire
  pending intent, the thing on the other end can launch activity, and sometimes
  when you fire pending intent you don't want this to happen. So if you are on
  settings screen pending intent can launch and make app visible.

exo -- internal, ambient computing something. From their system app they want to
launch e.g. Spotify even though they are in the background.


Ed -- is he somewhat distracted? He always have things in many pies, a lot of
work on 3P appstores, did a quick hack - a little toast that app is accessing
clipboard. But he's offloading a lot of work to other folks. Chatted yesterday,
he's looking forward to get some brainstorming on T. 


Working.


## Andrei.md on 2021-02-18

 
 - Angel go/angel-bc
 - FMD location history
 - Platform work - trusted UI, sdcard package visibility. Plans for T: DCL


- Interesting/new projects happening I can be a part of?
- Intros to directors?

tablets is the new thing
HC is distributed between Nicole and Narayan 

Narayan is hiring ~12 people
Nicole ~10
Media ~3

Andrei made a donation to Chrome team
Funded Grace's team for developer stuff



## Dave.md on 2021-02-18

Resource ask
What is Dave's vision for LON team?

FMD/Spot situation

Angel

API Safety thing

What Dave wants to tell DaveB and Sameer: ...

original proposal: 5 HC. Ash is now working on this full time. Need to add red
team people to really go after these API issues. Program Mgmt. More support in
SWE team as well. Sudhi didn't get any more HC, but decided to prioritise this
in him team.

Dave asked Khawaja/Sudhi/Scott and figure out a preliminary allocation to their
different teams based on OKRs/goals we are committed to. 

Happy that we have LON team, user-facing security. At the same time always
wondering in the long long term if this feature will extend beyond safety alone,
when it goes beyond our core mission, and requires set of product experiences,
goes beyond our mission and capabilities to build it. Spot discussion forces
this to happen.


## Raissa.md on 2021-02-17

Nice working with Karishma

Angel feels like a startup


## Guanhuan.md on 2021-02-16
- Situation with 2FA with FMD


- Android TV - in discussion with them
- Auto
- Battlestar

ATV - going on since Q3 last year, lot of work done. Beyond Mobile team also
ready to execute on that and so is Niko. But ATV walked back on their promise to
implement Q1 this year, but now they say Q2 this year. Discussed with James.

Challenge is that ATV team has record on walking back on their promises.

GPP process isn't really designed for different form factors, not detecting
anything ATV specific.

Let's find out the scale of this problem.




## Peter.md on 2021-02-16

Planning to fly back in 3 weeks time

Bessie still defers decisions to other people, questions she could decide she
says "I don't know what to do about this, wdyt?". She's getting better. 

Permission Rev split into 2 halfs, Peter works on one, Bessie on the other


# Week 6
## Alan.md on 2021-02-11

Ricky's improvements to BAL won't make it probably.

SD card still under risk


## Khawaja.md on 2021-02-11

AP4A - 1%, no visible changes yet.
FMD - location history situation
Angel update
DO we care about other form factors.

He still doesn't have an answer on when HC will materialise


## P&E_leads.md on 2021-02-10

Geoflex: existing intl geoflex extended for risk 4/5, return by May 3rd.

If risk lvl 5 then wait until this level reduced to 4 and then 60 days to return
after.

Split start hire - blanked date to September

Don't require post-promo CME.

Geist results - end of March

Can see WFH insights but not return to work

go/googlegeist-framework


## Inara.md on 2021-02-10

Structure of meetings with Simon:

 - Usually write up some notes on what to bring up before
 - Catching up on different things working on
 - Updates on the projects, issues
 - Talk about anything else unrelated, e.g. recently discussing perf

No specific conversation about expectations yet. In terms of career aspirations
said we'd discuss it after perf. Have been asking him for feedback.


## Coaching.md on 2021-02-09

what work to ask for if want to become director

why not learn more about those areas?


## Running_notes.md on 2021-02-09

He has to figure what's wrong

therapist helps to get to the root cause of the issue

Unpaid leave for few weeks?

speaking with someone impartial - EAP, therapist

GeoFlex extended


TnS in YT highest strategic priority
Very solid funding
Zurich is largest footprint in FTE working in T&S for YT, very serious site
All of the teams have grown
Biggest problem is to find good people fast enough

ZRH TnS is media specific / content understanding
Behavioural interactions, account-level signals, machine learning
Lots of graph mining, clustering
Novel interfaces - dealing with escalations, giving operators powerful tools to
find bad things
Intel Desk - team responsible to understand what's going on in the world
Security Engineering in ZRH
Trust tiers - creating "Credit Scores" for individuals

Hiring a director in ZRH.

tns should be part of fundamental YT infra, set of APIs. 

at any time 60-80 features being built by various YT teams. 

Difficulty is - most of this work, in terms of API definition, and infra
connecting other products, primarily within YT, is mostly happening in Bay Area.
DOesn't mean it's impossible, but it's an anomaly. 

In the perf world - continue on current path, create this set of APIs/processes
how to effectively make TNS set of system APIs, with the rest of Google, provide
this methodology to Maps etc, but most of the people working on this are in the
bay area. In the perfect world this role would exist but be in the Bay Area.

The Director for ZRH will have prevention as the core mandate. Matches to Trust
Tiers, Security, Bad Harvest - understanding what's going on and 1st and 3P API
that go into YT as a whole.

Mario something in ZRH, experienced director who looks after media content
understanding. 

Another person Matthias Conrad, also in ZRH, responsible for content matching,
also some score how much we trust creators. 

Hiring L7 lead for YT Security.

rady@ - he is technical hands-on deep low level security expert, effective TLM
of this team right now.


## Narayan.md on 2021-02-08

Concerns about permission laundering

Explicit permission check

Have to quantify level of privileges

Check in content provider that requestor holds role of X or has special
permission

Role is better - guarantees there's only one of those things

Don't want to set a precedent

Not worried about Pixel but OEM case

API council would tell you not to do it.

Not the sort of things that are stored in the database so no content provider
route.

Partners were pulling out of their version of permission hub

Want to get Permission Hub in time for S


## Monir.md on 2021-02-08

From the detection side - focussing on PII abuse, Spyware is 80% of malware.
Something on static/dynamic analysis for data collection. Want to go more
systematically after API abuse. Xun leading x-team effort, PII abuse WG. 

Login walls - there's work, progressing slowly ... There was a launch last
quarter having some automated login for Google accounts.


## Meerkat_Weekly.md on 2021-02-08

For today's weekly:

 - This is an optional cycle, yay!
 - Perf tool opens on 16/02
 - Self-assessments due 22/03

Shouldn't really affect anyone, but this cycle we are in pilot for in-org promo
up to and including level 7.5. James and Ed, this is your chance :D

Speaking of promo ... feedback.

Today Advanced Protection for Android is going live.

Angel is going nicely

# Week 5
## Simon.md on 2021-02-04

MiL got vaccine, living with them

Linus to handle discussion API council / PlatSec

David still not there yet. Where are technical deficiencies?

Karishma and Raissa working together.


## kshams-leads.md on 2021-02-04

Still don't have heads in pez


## Dave.md on 2021-02-04

Bluechip top of mind, and all the things.
Project Rubidium - ads privacy stuff
Potassium - Chrome privacy strategy

Then Apple did their thing, put Anning on its head.

Suzanne is in the lead right now, Krish is mostly on the hook doing the heavy
lifting for code yellow, looking for new PM lead because Krish doesn't have
time. On Eng side Sudhi is taking leadership role. Looking at bringing new
Android eng leader to look at API surface. 

Narayan/Dianne involved into this discussion. Suzanne and Dave are looking at
finding L9 leads to own this.

Need agreement with Pixel on backporting Angel.

92 heads from Sameer.

Angel, Responsible APIs

2/3 of the heads

Of the 92 heads:
 - 8 donations (3 T&S, 5 various other parts)
 - 84 heads will go to the teams. Of them:
   - 78 will get shipped out in the first tranche
     - 22 (by far largest number) going to Khawaja
   - The rest will come later
     - 2 more to Khawaja
   - Little bit of reserve Dave and Suzanne holding out

Khawaja will get mid-20s.

No anonymous Dory.


## Alan.md on 2021-02-04

Ben got offer from Cambridge!!


## Anurag.md on 2021-02-03

Where are we with location history? Do we need help from Nandini? Update this
thread: https://mail.google.com/mail/u/0/#inbox/FMfcgxwLsKBvSptQbBMwRGMtwWVJshQd


Met on-device location team last year: Chris Snow. He understands our usecase,
and understands it is not covered by on-device, he says that we can try to make
the case, it will be super-hard to keep it because Google started this effort
for a reason (law enforcement etc).

Anurag collected stats, and wrote a doc on how FMD can deal with it. 

If we have to build it, can we build it in a way e2ee. Folsom team - building
e2ee for backup/restore, can sync cryptographic material across devices. Met
Dirk. He said that encrypting geolocation is the usecase interesting to them, OK
to them. Problem is if FMD wants to build it it would be a gigantic effort for
us.

Not engaged with Nandini's team. Spoke to Ravi last year. 


## James.md on 2021-02-02

Sameer thinks we should combine privacy and security.

Abhishek - proposal to combine privacy and security, which looks 95% Angel.

GMS req convo


## Ed.md on 2021-02-01

living on the edge of IC deadlines. Alan is dealing with it. None of
the features are currently at risk.

Started Meerkat T list of ideas:
* DCL and 5 more bulletpoints
* We've blocked BAL. But we've seen reports of apps launching URL into Chrome.
  We prevent background popup, but if you are already using Chrome any app can
  open website.
* Bernardo has ideas untrusted touches.
* and more.

Sameer had a change of heart. 
Comp angle - they are nervous about everything we are doing.

Allowing stores to do updates without confirmation: probably safe-r.

Ed prototyped it, someone from Narayan team building it.


## David_Coffin.md on 2021-02-01


AP4A goes live noon LON time next Monday

Roman to write launch email


One of the biggest concerns - getting settings app and framework changes done in
time. Marie and Simon  were looking at Settings changes and what data do we need
from it. Linus looking at framework changes. 

Plan B - rather than making framework changes to get the data to replicate
existing settings page, we can add content provider to the settings app. 

Stub APK final version is blocked on Settings and Framework work, but David
wants to get at least something checked in into system image.

Angel Business Logic document.

Raissa working on integration with GPP. 

Simon prefers Karishma to lead integration work with GPP, but Raissa is really
very excited about it.


# Week 2
## Simon.md on 2021-01-14

Karishma fed up with lockdowns.

Her mum is in need of non-emergency surgery. Hip replacement surgery.

Karishma wants to at some unspecified time towards the summer when things
looking better, go to India, take 2-3 weeks of carers leave and WF India for
several weeks. 

Reallocate intern headcount.


## Inara.md on 2021-01-13

Granddad 91 old got second vaccine

More interested in working on the UX


## David_Coffin.md on 2021-01-12

Early next week - decision on notifications


## David_Coffin.md on 2021-01-12

Fairly early stage.

Figured out the sequencing of pieces of design work. Written a draft
requirements doc, Simon reviewed it, can't think of anything else, not ready to
share externally yet, captured headlines of what needs to be done for design so
can start breaking out.

Linus will start on Platform changes this week.

Marie looking at Settings changes. 

Have an idea of what will the teamrolls be for the next quarter.

We've got a prototype (shared last year) to prototype an integration with the
settings app / GmsCore.

Requirements in terms of what's in and out of scope of integration between GPP
and Angel, business logic for how settings and traffic lights interact. Started
having conversations with various teams we'll be integrating. 

David is working on prototyping an API to get integration with GPP up and
running.

Pumping more people - 2-4 weeks from now (1-2 sprints time). 

Maybe can have someone work with David on GPP integration?


## Khawaja.md on 2021-01-12

Write up job description in draft mode and send it to people internally.

OPEX was approved.


## Running_notes.md on 2021-01-12

2.5 y.o. she-toddler

Helping parents to make good decisions about apps, doing this minimising the
amount of work they have to do.

Let's produce a library of apps which went through certain amount of scrutiny so
parents can feel safe installing these apps. All apps in the corpus were
manually reviewed. Less about if actor is bad / app is malware, and more about
is the app targeting too much advertising at kids. We have a sense of what most
important apps are, not particularly correlated with popularity. Lots of apps
hyper-popular in kids space that never make it to global popularity, because
space is small -> but dumblestore is not correlated even with that.

Part of the goal is to increase the popularity of good kids' apps. 


## Anurag.md on 2021-01-12

Good retrospective discussion.

Strong focus on everyone being focussed and getting things done.

More ownership even for smaller project parts by individual people. Quite a lot
threads are open and every sprint there's some juggling of bugs.

Have 2 groups / other options. 

Mostly not fixing bugs in the old backend.

TIMON now is just building the new backend, M5 of the original plan.

2 migration to happen after that:
- Migration of web FE
- Migration of the app

Server side platformisation will be done in H1. We'll start to send app traffic
to Timon in H1, but getting it fully launched in Q2 is P2, so it might overspill
to Q3. Web migration is after Q2.

Making app talk to Timon is 75% done. What remains: once backend ready start
sending traffic from Timon to the app, and then revamp the UX.

The app is the majority surface. Maybe lots of people don't understand that the
app is not necessary for finding phone. 

Tracking is less likely because of notification that appears on users' device.

Conversation with Jim about his comp.


## Anita.md on 2021-01-11

Visited parents in Leicester

OKR ownership - spoke to Anurag, but Timon is the big thing right now, after
Timon will be more projects and have Anita leading one of those.

Expectations document - is in agenda.


## James.md on 2021-01-11

Angel - on UX/Product side everything seems to be in a good shape.

Need prototype app ASAP for people to install on Pixels.

DaveK - permission backport.


## Dave.md on 2021-01-11

HC: somewhere on the order of 80% of HC ask.

L6+ go through a different hiring process because of diversity, we have to show
we tried very hard.

Project Cake: fully curated app store. To have a separate app "Google Values
AppStore". Overlay on Google Play. Publish as experimental app.


# Week 1
## Sebastian.md on 2021-01-08

2 new Working Groups:
 - anti-anti dynamic analysis
 - red team / pen test of our defensive systems

present to leadership about native code.


## Alan.md on 2021-01-07

Linus - doesn't want to be in the UK.

Next IC is Jan 20. Only 1 feature done and dusted - Alan's!


## Simon.md on 2021-01-07

Niko stuck in Chile. For at least a week. Flights cancelled.

Peter stuck in Hungary.

Veronica rejoining tomorrow. She wants to switch to UX. Asked if Rodrigo had
something. She might join aroscoe@ team. He was very keen before. She hasn't got
any UX training.

David seems good.

Simon - still trying to workout the childcare. MiL. Maybe have to take 2 days a
week, more likely 1 - 1.5 days a week.

Angel - was trying to work out if realistic. Want to have more realistic plan
for work between now and March. Want to spend tomorrow and much of next week on
that. Might ask me for help. EO next week deadline.


## Guanhuan.md on 2021-01-06

https://en.wikipedia.org/wiki/Wenzhou#Language he's from here.


==== 
for next 1:1 
=====


## Narayan.md on 2021-01-05

Chat at the end of last year with Sagar and PM team, now at a place where we
think it's a right thing to build it in S if we can. Earlier strategy - tell
them how to build it but not build it ourselves, hypocritical.

Sagar has some reservations on how useful it is for the users, which Narayan
partially agrees with. Need to do fair amount of user research.

Resourcing is still not decided. Narayan can scrabble something together in his
team probably. 

If we do this for S, Angel story is not fully complete.


# Week 52
## Titaniums quick research.md on 2020-12-21


```
CREATE TEMP TABLE dumbledore_ap4a AS
SELECT
  distinct(docid.backend_docid) AS package_name,
  IF(docid.backend_docid IN (select package_name from mobile_ap_pipeline_prod_jobs.app_allowlist_ap4a), 1, 0) is_ap4a
FROM
  finsky_offline.playwright_export_prod_prod.kids_corpus_tag.ANDROID_APPS.ANDROID_APP.latest
WHERE
  'approval:kids-corpus' IN UNNEST(annotation_envelope.document_tags.tag);

select count(1),sum(is_ap4a) from dumbledore_ap4a;
```

7328 total dumbledore
572 ap4a approved



```
CREATE TEMP TABLE apps
AS
SELECT
  docid.backend_docid AS docid,
  IF(
    docid.backend_docid IN (
      SELECT package_name FROM mobile_ap_pipeline_prod_jobs.app_allowlist_ap4a
    ),
    1,
    0)
    is_ap4a
FROM
  finsky_offline.playwright_export_prod_prod.ANDROID_APPS.ANDROID_APP.latest
    AS document
WHERE
  EXISTS(
    SELECT *
    FROM document.document.tags.tag AS tags
    WHERE tags = "playpass:playpass_general"
  );

SELECT COUNT(1), sum(is_ap4a) FROM apps;
```

715 playpass
442 ap4a approved


```
CREATE TEMP TABLE apps
AS
SELECT
  a.docid.backend_docid AS package_name,
  IF(
    a.docid.backend_docid IN (
      SELECT package_name FROM mobile_ap_pipeline_prod_jobs.app_allowlist_ap4a
    ),
    1,
    0)
    is_ap4a
FROM
  finsky_offline.playwright_export_prod_prod.ANDROID_APPS.ANDROID_APP.latest a,
  UNNEST(document.decoration) AS _decoration
WHERE _decoration.docid.backend_docid = "editorschoice";

SELECT COUNT(1), sum(is_ap4a) FROM apps;
```

935 total
802 ap4a approved

# Week 51
## James.md on 2020-12-16

FMD/Tile/Spot situation

There was a review last week with SPOT team. Spot are part of Nest now.

Nandini was there for location team, Jen, Jordan.

Topic: crowdsourced location platform, phone/accessories supporting Eddystone.

Very privacy sensitive. Proposal doesn't pass step 1 - needs a privacy design
doc.

Issues on product side. The way privacy design would work: on Apple to find your
lost iPhone they can remotely enable find my iPhone. We don't want anything like
that, privacy-problematic. The way to work around is a concept of trusted
device. 

Seems that Pixel BD baked in their agreement with Tile/set expectations with
Tile, that we'd support finding Tile as part of this. Tile to move to Eddystone
and that'd allow using Android devices base to find their devices, so Tile would
appear in FMD. This means there's hard dependency on this happening by autumn
next year. But we won't get it done by autumn next year. So deal is in jeopardy.

Once Spot deployed their protocol, and we as Android team would help to get it
deployed to all Android devices via GMS. Then we have to support finding Tiles
in FMD.

Tile insisted: basic finding functionality in FMD but no ring.

Spot was just transferred to Nest. 

New location backend coming online, on-device location history.


# Week 50
## Dave.md on 2020-12-10

There was an API call that would give you SIM ID. One API was killed in Q, and
one killed in R. When we took down first API, we never ran into any problems.
Payments apps problems (see Giles' notes).

Play added feature where they can allowlist apps to be treated differently by
the platform using server-side configuration they can push very quickly.

- For now eng team is responsible because low user
- We may hire TVC to do this on a regular basis going forward, if needed.


## Bessie.md on 2020-12-10

Kotlin readability
Auto revoke going on


## Alan.md on 2020-12-10

Delphi finished
Fi collects stacktraces - mentioned to Pauline

Bernardo's feature notification shade - stopping notification trampolines now. 

Ricky is back to HKG.


## Khawaja.md on 2020-12-10

Bonus:  £48.9K
Equity  £180K
Salary: £164K
Total:  £398K 


## Simon.md on 2020-12-10

This is last week of work this year.

- Next 1:1 - not spending enough time with David 

JP in Munich. Manager of PDPO in Munich JPWeber@

taz@ / Tal


## Linus.md on 2020-12-08

Feedback for Alan - he'd probably be IC. Had other managers that had more
interest in being manager. Team is doing things, but don't have clear direction.
Don't have path to grow as a team.

As long as UK is in L5 (at least until Jan 18) will work from Sweden.

## Guanhuan.md on 2020-12-08

Off from this Friday


Latest on the Tile:

- About 3 weeks ago Tile said we are happy with FMD being able to ring our tags.
  However, we want in return to be offered as part of FastPair / Spot offering.
  They wanted to be default part of the offering, whenever we reach out to
  partner Tile is one of the options or default options, which means all future
  partners we bring on board are Tile enabled. 

  Difference between Tile/Spot devices and our Spot devices - they'll be
  onboarded by Tile app. We said no to them, BD team asked to escalate because
  big decision, James and Krish said no as well.

  BD team went back to Tile. Tile came back and said: if we don't have Tile as
  part of FP offering it's fine, but we are happy to embrace whatever you are
  offering, but we also want to put 80% of our product line to implement
  Eddistone protocol. Proposal #2 is they want to be able to take Tile/Spot
  integration to any further of their partners to say - hey, implement this to
  make your device findable. The problem with this is - Tile is probably going
  to ship all of the tags, Bose, etc with new Spot integration, which would
  potentially race us in terms of time -- they can get ahead and be first major
  line of product which supports this. This is not ideal.

  FastPair team is also in favour of not bringing Tile on board, they want us to
  walk away from the deal entirely.

UWB team wants to have a product that would work with Pixel. Also Tile increases
our network significantly.


They are supporting modularisation of our Android app. Near done. Next step -
they gonna work on actual Spot device user journeys in the app. They'll do some
underlying work. Rodrigo hasn't been able to spare any time to work on the
design.

Because Spot belongs to Nest org now, potentially they have other projects to
work on. GH speaks to David periodically. Nest committed 3 engineers from Spot
team for entire 2021 to work on Spot projects, which is basically whatever we
need. Including Material 2.0 compatibility.

For Spot finding network GTM: if we are to enable 2.6B devices to scan, we need
to have adequate number of devices to be findable by this network. 3 options:

1. Tile/Chippolo
2. FastPair partners
3. Enable all Android phones to be findable.

(3) solves offline finding scenario. Not clear how much value it would be. How
many FMD locate requests failed because no connectivity?

Chromebook not listed as option yet, but Nest will speak to one of the PoCs.

Another option (early stage) is Windows laptops - some of the OEMs, 200M per
year. 

If we take option (3) Nest are happy to invest more engineers into project.



## Anurag.md on 2020-12-08

Feedback on review delivered.

Chat with Anita happened on Wednesday. Went through expectations from Anurag,
made sense. Talked to her on Anurag's expectations of her, what ladder says,
went through the ladder together. That was pretty obvious. 

Had reservations about independence. Mostly landed well in the end.

Spoke to Sorin about vacation, he should be empowered to take technical
decisions. Important for execution speed.

Sized OKRs.

Chris Nough ??? dir eng location history who/chsnow

They haven't considered FMD case. Anurag made security trumps privacy argument.
They understood all of that. Their PgM said they want to remove server location
history because law enforcement. Situation from law enforcement won't change if
we build replacement. He did say the very reason for this multi-year multi-PA
project to exist is that they want to stop sending loc from Android to Google
and it is going to be super-unlikely for us to get exception.

Guanhuan and Anurag will write a one-pager on why loc history important for us,
and then we can see if they see the point. From Google's PoV it won't change.
Separate consent in FMD settings.


## Running_notes.md on 2020-12-07

Competition council: they are aware of APP. Things we are doing are justifiable,
user base is very small, protection can be circumvented. Likelyhood of someone
being in position to raise competition concern is close to null.


## Simon.md on 2020-12-07




## David_Coffin.md on 2020-12-07

Bit and pieces of the design doc
Convos with Henry and Simon about it.
What agreed is work split:
 - When Simon and Marie get more availability they can start working on  parts

Working on prototype this week.

Prototype expected: building Android now.

There's some minimal prototype.

Not feeling like spends enough of time with Simon.


# Week 49
## Serban.md on 2020-12-05

Protected KVM

Plan is to get MediaTek to use it. Want to transition some Samsung HV solution
to new HV.

Targeting Note next year for Protect KVM. Similarly on MediaTek.

Pixel roadmap - 4 devices in 22, only some of the devices in 22 will use new
SoC. Low end will ship with old SoC.


## Anurag.md on 2020-12-04

Next week is the last in December. Same for Guanhuan.

What is the plan for team to be productive in that period?

3 week sprint next one. That sprint planning happens next week. Work Profile
will hopefully be fixed next week.

Timon is in the place when they don't need to come to Anurag for answers, partly
because Sorin takes responsibility for design of key pieces. 

External collaborations - people aren't that involved. 

team of 4 engineers, would be more productive if can do more than 1 thing at a
time, instead of 1 thing after another. Thing proposing all Timon still.
Planning to declare server-side platformisation done as soon as M5 complete.

[go/timon-pp Project Plan](go/timon-pp).

Sorin shared some frustration of being slowed down by having explain quite a lot
of things to whoever he's working with. 

People have quite different preferences in FMD. Sorin always wants to stretch
himself. Anita wants to not overload herself or the team. Only way to keep them
both happy is to have individual projects for both of them.

For next 1:1:
 - Development plan for Anita. 
- Feedback on the review: could do with more specific narrative, i.e. what is
  one thing we want execs to remember from that meeting. Having specific ask.

Stressed with the expectations Anita wrote. Key things on project delivery were
missing from the document.


## Simon.md on 2020-12-02

Is Anurag feeling a bit overwhelmed? Not responsive to some Richard's messages.

David - he mentioned did some deep dive discussions with me on exact
expectations.

Angel: David is looking into what APIs we need for device-level stuff. Building
a list, already identified some (PIN/password?) not available in any sensible
way. Linus says "I could turn this into API for you relatively quickly".

Google acct integration - major source of uncertainty. Need to work out how.

James - scattergun, all over the place. Compared to half a year ago - much
better, but if he asked what you want to do  / what your priority never get an
answer. 


## Giles.md on 2020-12-01

Heads down on Delfi annotations. Trying to get all of GmsCore Clearcut annotated
with data collection justifications. Collection-basis verifier. 

3 components:

- all protos annotated with reason
- policy that says which consent you need
- run-time collection-basis verifier that checks that you have consent to
  collect this bit of data.

Failing is an ultimate goal, initially we are just going to log/report it and
file as incident. 

Horrible incident which we fixed - gPay team said there's Indian regulation that
SIM card ID has to be provided by payment provider. Had to do app op flip via
Play.

Finally hired 2 people to work on GmsCore privacy specifically, two new Nooglers
starting on 14th.

Project Amsterdam. Big new preload policy, 7 different areas preload policy. 4
different privileged permissions, incl notification access/listeners, app usage,
a11y, INSTALL_PACKAGES. A lot of new policy around signature permissions and shared UID.
Undisableable apps.


## Anurag.md on 2020-12-01

Server side location history is going away. Sunset in Q3 2022. In the new world,
only on-device store of location data. Exporting this data to Google servers is
not allowed.

For FMD - problem. Options:

1. What if we don't have historical location? How often does users need
   historical location really these days? For FMA it's simpler in the sense that
   if we assume that users lost their accessory, we can make reasonable
   assumption they haven't lost their phone, so we can use device-side
   historical location. Right now it's disallowed to send this location off the
   phone. Maybe we can get an exception?

   Harder problem is for the phones.

2. If we have to replace location history, it's a lot of work. Saving location
   data on the server is part of problem, second problem is optimising it on
   device to use less bandwidth / battery.

3. Get Spot do it for us. They have network of sightings. Right now they can not
   find Android phones because they don't emit BLE, and all of the sightings are
   encrypted and only master phone has encryption key.
   

## Anita.md on 2020-11-30

Uke 140 with 90 without

2160 in total


# Week 48
## James.md on 2020-11-23

Angel for Pixel SW deck
Angel request from PaulB
GPP ATV escalation
FMD platformisation deck


# Week 47
## Running_notes.md on 2020-11-20

Tricky, AP small proportion of users. If we put a blog post - we can't say that.
Worry is that this will really set a flame for developers as they won't be able
to see the relative scope for AP, same press. 

Competition angle - will they spin it?

We don't want to make it look like Play is not secure.

People might say that are you dissuading people from downloading apps which you
said are safe on your store?

Scott's gut - do HC article, have UI in product and reactive comms if that comes
up. if it's only 50K, only 12% apps will see speed bump, not preventing from
being installed. We can have good statement always working on improving
security, assess developers, working to continue to expand the list and if any
developer wants to extend here's link.

Aaron: comp will be tricky part from policy standpoint. We are letting our apps
through but imposing this on developers, and not telling them what specifically
they need to do to make their app safer / get approved.

Objectively good security news.

Proactive: package this up with other announcements?

Scott: difficult to broadcast message that applies to just 50K users.

What will competition council come back with?



## Khawaja.md on 2020-11-20

Kids 16 daughter and 13 son

job apple recommendation ranking relevant search


## Running_notes.md on 2020-11-18

Projects in flight:
 - Angel
 - Permission Revocation backport
 - AP4A; Carter++ want something but this is on hold
 - Find My Device update
 - Platform

Discussion with Suzanne, Krish, Dave Khawaja.

Sameer is on board with Angel - was sceptical couple month ago - now very
excited, want to have broader meeting to discuss investing more in Angel.

Probably no dramatic shift in direction.


## Niko.md on 2020-11-18

ATV escalation:
https://docs.google.com/document/d/1FpERWPDPl_lDS8tn4hycZOWeZwUI8jv7lLebGXPKzoo/edit?pli=1&resourcekey=0-rkzj5Ce1i_pK_ztICpFmCA#

Passport:
https://docs.google.com/document/d/1zMlOlfd_5lregPEhX_XDhTZIoY_n0fDLDEqI7EempHc/edit#heading=h.x9snb54sjlu9


## Anurag.md on 2020-11-17

Takes 15 Jan holidays, also off next week.

Lost a bit of velocity.

Isabelle starting next week. Starter project?

Anita as mentor.

Jim sees himself as a generalist, likes to do many things. Didn't know off the
top of his head projects he'd like to work on. 


## Running_notes.md on 2020-11-17

AP4A biggest concerns:

 - Is the list credible? How did you come up with it?
 - False positives / negatives?
 - Types of apps that are probably OK but not allowlisteD?
 - How will developers react?
 - How will we operationalise the plan to get them allowlisted?

Competition landscape.

Anything that pisses off developers so they go cry to their government, might
require Sameer approval.

Google Patch - bundle of apps, get Google experience on your phone, install
other apps when you like it.

Algo with 40K, how we come up with it etc.


## Dave.md on 2020-11-17

What success look like for Angel?

What % of users is in every state?

Nick won't help with operationalisation.


## Running_notes.md on 2020-11-16

Guardrail - different tactics, no longer innocent until guilty, more important
companies. Fastest lane.

Sameer/execs feel that treating everyone the same is too harsh.  Podcast Addict
story.

There are different violations that require different punishments.


# Week 46
## Titaniums quick research.md on 2020-11-14

I want to know:

 - How many users in AP population will be missing apps?
 - How many apps they will be missing on average?

So we formulate a question as:

    For each user in the AP population, among apps installed from Play, how many apps are not in the final filtered down set? 



## Sebastian.md on 2020-11-13

~40% of apps have native code, so we _do_ need to have native code analysis.

Not yet able to find any bad apps using native code analysis.

We have test results: when native code analysis results were used in ML it
worked pretty well.

Native code analysis is not at risk, but we have to launch all this stuff into
production. Exploits - much harder. If we don't have it, Proj0 disclosing
vulnerabilities in 3 type SDKs.

Creating proposal on how we can get code obfuscation and DRM under control that
allows us to inspect the code ourselves. 

Supposed to tell Shie how we see priorities of policies to be updated. How do we
suppress malicious behaviour with policies.

BenR.


## Monir.md on 2020-11-13

Zhao works on Muws, Damien works on malware ML.


## Anurag.md on 2020-11-13

Sprint planning structure:

- 30 min retrospective
- Team appreciation
- Henry presents on how did the sprint go
- Then discuss it

Then planning poker.

Backlog is pre-groomed.

Also pre-planning meeting, eng team, bullet points of objectives for next
sprint. Which milestones we should focus for in next sprint?


1 sentence goal: get her confidence back and get her more independent.
Does not have motivation problem like Jim. Efficient. She does work 60%, less
time to get things done but she gets them done.

EXPECTATIONS!!!

Right now she has over-reliance on everyone else on the team.


## Jessica_L.md on 2020-11-12

We do have HC link to submit app.

Check our main competitors / antagonists app first and maybe allow them.

Going to be worst if we have first party app for particular function which is
approved but have up and coming apps which are not on the allow list which
compete with our apps.

Reactive comms might help to mitigate the risk.

People don't believe us when we are saying we do things for safety reasons. 

Should we launch appeals process before program going live? Add something to
Play Developer Console.

Treating 1P and 3P devs as equally as we can, apply same objective criteria. 

Maybe we can show notification in PDC?




## Anurag.md on 2020-11-11

working hours 10-17:30, and then couple more hours 10-12.


Need to get more feedback from promo committee on what exactly has happened that
denied promo. 

Asked Sorin to prepare metrics on where he is on L5 ladder. Sorin did all the
work for Timon Android, let's compile one large document.

Sorin taking lead on eng discussion with Spot eng team. They would refactor app
for us, A advised Sorin on taking lead on this refactoring. 

Split Timon into 2-3 major projects. E.g. Remote Instruction is different from
List.

Sorin will possibly look into collecting new metrics - something he can deliver
end to end. My rec: give recommendation and steer team.


Most interesting project for Anurag :)

Shared expectations document.


## Running_notes.md on 2020-11-11

KobiG

GP3 - 80% of all installs.
Using trusted model on this cohort of 5K developers found significant
proportion of developers coming as non-trusted. Need to know.

GP3 = public managed.

This app is because GPlay partner or because reviewed it otherwise.


Chiawei - designer

https://docs.google.com/presentation/d/18X54e91Grl-DIAYrCjojqMyjUN6bJ-OyvwdO96EeRFc/edit#slide=id.g8ac67a6a73_0_0


## Running_notes.md on 2020-11-10

AP pre-review: reduce probability of spearphishing. 


## Guanhuan.md on 2020-11-10

FMD campaign stats:

- On the day, 200K more daily downloads of the app
- Sent to all 365d active play Users, UK/US/DE/BR/JP/CA/AU
- More than 80% user sat, top 3 emails they sent


## Narayan.md on 2020-11-09

Android Go is working on  app hibernation - apps that are not used, preventing
them from running at all. Current step is just revoking runtime permissions.
This has its own UI, notifications etc.

Permission Hub is even more controversial. Leadership group no agreement whether
we should have Perm Hub at all. Had countless issues with previous permission
hub in Q, and they continue being issues. Many things in Perm Hub are not
actionable. Samsung is launching Permission Hub in R - maybe some cut down
version. So is Xiaomi. All info privileged. Charmaine is in the loop, she has
most context.

Chinese government regulation that all devices sold in China have Perm Hub. We
are working with them trying to change definition - a lot of what they want to
show is misleading and pointless.

In S: no perm hub, but some data clean-up. Not building any UX, but if we are
forced to we might be prescriptive about what we permit our partners to build.

Set up some time with Charmaine to talk to her about some plans post S. 


# Week 45
## Khawaja.md on 2020-11-05

Discussed Angel deck
Kshams thinks we should proceed with eng work no matter what, even if Pixel
disagrees. 

Does it give Pixel what we need? Probably it does, and we can iterate on it.


## RichardD.md on 2020-11-05

Anurag handover
Sorin feedback

Very capable engineer. Brilliant as FMD tech lead. Good at detailed programming
as well as system level view, everything expect from L5.

Biggest issue Anurag will have: controlling his workload. He's fairly young.
Does extra work outside of work. Already overworked. His typical day: starts 10
AM, has no worries about going right through the evening with his work. Very
enthusiastic, if he tells he'll get something done, he'll get it done, even if
it needs extra time. 10 till 5-6 but extra hours in the evening, more meetings
with the U.S. Sometimes till 10PM.

Health wise - still gets these sparkles. In Feb it was quite a lot, doctors
couldn't find anything. Then went away for some time, and then came back
recently. 

Anurag has to develop relationships with other people. Not saying it's a
problem, but he didn't have to do it yet. For example, Richard had a long chat
with Pauline who trusts Richard on privacy&security.

He's good at working with people - e.g. other eng to do the work, but building
trusting relationships will be extra strain on him. Very important for FMD. Also
Jessica. 

- Sorin - get along really well with Anurag. See no problems there. He did tell
  R that he's very pleased Anurag is taking over. 
- Anita - equally don't think there will be any issues. 
- Jim - R also struggled with Jim. Working out how to motivate him to work,
  constant battle with him. Maybe some shared expectations document? Scraped at
  CME, let's understand what do you need to do to ensure you're meeting
  expectations of an L4 engineer. Shared expectations document is a preparation
  for NI. 

Couple of times R had to chase him on communication. 


## Anurag.md on 2020-11-02

Better at figuring things out than remembering them.

Send over email is better for important questions.


## Guy.md on 2020-11-02

Webivew tollfraud detection captures 10% of overall tall fraud, want to get to
20%

William got promoted

Get more telemetry on enforcement side, reduce the latency for enforcement,
enforce more types of abuse, potentially MuWS.

How much impact do we have from Warn flags?

How do we inform the user that there's a potentially malicious app?




# Week 44
## Dave.md on 2020-10-29

Not sure why Pixel went quiet for so long.

There was already bluechip is the highest priority, but it was a lipservice -
not that we'll shut down stuff and make hard decisions. Last month they really
started making hard decisions. They shelved Digital Wellbeing for a year.

Staffing will come from ASAP.


## Running_notes.md on 2020-10-29
Dave, Suzanne, Stephan, K2

Privacy Hub tested well, not Angel.

57% which apps are using your data, when and for what purpose
41% get notifications to check if you want allow access to rarely used apps
33% monitor reports and settings


## P&E_leads.md on 2020-10-28

Early stage trial - come to the office for a brainstorm, for an hour or two for
workshop. Let Andrei know, and then Jess or Laura will send a form which we have
to fill in. Most likely some tech talk room will be assigned, and we'll have to
provide feedback to REWS.

Change to wearing a mask policy - up until to this point office priority
required to wear mask except when at desk. Now mask must be worn at all times
except when eat or drink.


## Simon.md on 2020-10-28

Bump up Niko compared to Sorin, and Sorin compared to Jim, and Anurag cmp to
Peter.


## Running_notes.md on 2020-10-28

Settings integration - some OEMs take AOSP settings, and some forked AOSP ages
ago, so integration point question.

We are not mandating digital wellbeing because of competition/partnership
management. Requirement is that OEM should have _a_ feature, and if you don't
have something available, we have a solution for you. It's a MADA / GMS
requirement. 

We can also mandate that e.g. next year Digital Wellbeing must have feature A,
but we are not mandating using our app. About 9 months heads-up.

Samsung and others are very sensitive about biometrics setup. They might see us
as being slow, they've adopted FP 2 years before AOSP, they don't think we can
innovate as fast. Also they want to own app verifier, it first pings Play
verifier, and if it gives green light they kick in their own verifier - they
don't trust Play verifier.

Adding APK to MADA. For headed apps we have some number of slots, and we say
that we will have no more than 10 apps mandated, and raising this number is
no-no. But we have core services / mandatory services, headless apps, we have
something like 15 apps we can preload, and this is mostly frictionless. 

Android Auto wanted to move out of GmsCore, nice to do thing, splitting
permissions, they didn't want to follow candence of GmsCore. This unnecessarily
got BD attention and got to exec level negotiation. It can go either way.

Many governments have good intended but silly rules about preinstalled apps.
Line between app / service is very blurry. From their point, icon is an app, and
you should be able to uninstall every app, taking this to some extreme. Korea,
Taiwan have similar rules.

JFK: pretty frictionless path to move gmscore to headless
UJ: flash size requirements tho.

UJ: you must have a security hub (consuming these events). If you choose Google
Security Hub, this is how you integrate with it.  If the size quite big this
wouldn't fly. 

Benefits of this approach we are flying under radar and not making this obvious
that this is an optional APK they must preload. With digital wellbeing it was
integrated into GmsCore so we didn't have to deal with it.

Current problem as stated here won't make Samsung very excited, as it's very
much focused on Google account.


## Khawaja.md on 2020-10-27

PMR - Roman communicates clear goals for the team - somewhere neutral, some
people on my team think I can do better there.

Giving actionable feedback on regular basis - nobody strongly agree, need to
dive deeper. 

Feedback - more often, more specific, discuss better expectations. Can be
dominant in meetings, both his own as well as others'. Sometimes thinking
whether to invite Roman to meetings.

When directs in meetings - ask open questions, refrain from giving your opinion
until everyone else has spoken.

Areas for development: focusing on doing things better / sometimes saying no to
senior people from the reporting chain.

Potentially getting a coach to help with soft skills. 

Ask yourself - whether your opinion is actually going to help and whether you
feel passionate about it. Pick fewer things to have stronger opinions about.

If you are about to say something that can be offensive but need to be said,
build a rapport with the people before saying that.


## Guanhuan.md on 2020-10-26

Google 1P team for location tags. Advanced Technology Lab, Jacquart (?) smart
apparel. Tile-like tag and embed into shoes and jackets, for motion tracking

https://moma.corp.google.com/team/6786785881

Expressed interested in partnership of Find My Network.

They have couple of existing partnerships with Levi's, Adidas and speaking with
Canada Goose and few others. Thinking of making 100Ks of tags.

Apple - AirTag (if ever going to be released). Crowd-sourced device finding.
Meant to do motion tracking. Will be able to embed it into watch etc.


We kinda blocked Tile partnership for UWB, so there's a risk for them to ship in
P21.:

  - They need a pilot for launch
  - Crowd-sourced location finding spot, close range UWB
  - We said we don't want to proceed without ring, so Spot can't proceed
  - UWB don't care about Ring, we do.
  - One still can produce UWB without Eddistone, but FMD won't be happy with the
    privacy story for crowdsourced location.

Spot != UWB from tech standpoint.

Spot relies on devices implementing e2e Eddistone protocol, basically BLE.

UWB is a new technology, allows to pin-point device in close range, last 50-100m
finding.

UWB on its own, for things finding, isn't adding a lot of value right now. In
future, when there's more adoption/implementation (e.g. identity verification,
ambient experience) - e.g. lights can shut in the room you've left and turned on
in other room, unlock door/car. Prevents man in the middle attack, good for
security.

Find My is one of the top use cases they have in the initial implementation.

They also proposed Find My Friends which is off-track for FMD.


# Week 43
## Khawaja.md on 2020-10-22

- For next 1:1 - email from Bobbie completely excludes anyone not in the U.S.


## Meerkat_Leads.md on 2020-10-21

What success looks like for our products? It's worth having a good understanding
that we can explain to the folks outside of the team.

Privacy is a big topic for 2021. How can we better support this narrative?

Virtual team event

## Guanhuan.md on 2020-10-20

Spot had massively transformed their product, to apply better privacy story.

UWB is 3 part projects, involves all 3 teams:
- FMD - we have an app refresh project, make UI modern and modulize app code so
  other teams can easily contribute to the app. GAR4, material 2 etc. Also reqs
  for Pixel preload
- Spot. They want to onboard bunch of partners to make it real and they want FMD to be the UX.
- UWB: they want to chuck in UWB chip into as many devices as possible, have
  OEMs onboard it and enable close range location pin-pointing. 

In Guanhuan opinion, FMD exists without two other projects. FMD + Eddistone is
another project and it also can exist without UWB. There's collab between
Eddistone and FastPair.

Eddistone brings crowd-sourced location finding to FastPair.

Tile negotiation is stuck - they want to reserve ring for their app only, and
they want to interleave Eddistone with their existing protocol, to make old
Tiles and new Tiles compatible. Not good enough for Google. GH took hard stand
on ring functionality - if you want to have project, FMD happy to provide
surface, but we won't be able to resource it as it doesn't bring any user value
to us.

1. GH told David to speak to FastPair team, FastPair reached out to JBL and
   Sony. 

Retrospective pairing when pairing was done via BT and not FastPair.

There's no place on the internet or device manual to find out if device is
FastPair device. Users won't know that it is FP enabled.

FastPair half sheet. When you bring AirPods close to iPhone it shows little
pop-up bottom of the screen - FP did something similar, just shows certain
information. This can now be used to recommend user to install OEM app.

We also want to nudge user to enable FMA in that half-sheet.


Tag making company called Chippolo - much smaller than Tile. They are super keen
and ready to do whatever we want. They want Spot and UWB. 

That's Spots' partnership negotiation status quo.

From a market perspective, there's an increasing trend for device finding
solution to start finding _things_ (or maybe even friends). Samsung made an app
to find everything Galaxy. Their SmartThings support some kind of Bluetooth.

Obviously Apple's findMy supports findings all Apple things incl tags to be
released.

Huawei out of the big picture, but all other big OEMs have their device finding
solutions.

Raise awareness, make product really good quality, google's 1P solution for the
ecosystem. All other solutions are only OEM-locked in.

In terms of UX, the solution is nearly 10 years old. This is a reason for us to
prioritise app revamp. Work FMD will be doing is the app revamp. UWB AR view
will be done by Spot for us.

Spot team really want to make this project real, and UWB also want. Spot is
freeing up 6 eng to do whatever FMD wants in Q4. 2 SWE to start and then 4 more.

They can start with dark mode. FMD to involve GPP team to get some reviews etc.

FMD dark mode - should start ASAP.

Other side of the project - we want Ryan to do a bit of UX research on FMD and
maybe we can make some simple lo-fi design changes. What the spaces in CUJ we
can improve user experience?


# Week 42
## Rodrigo.md on 2020-10-14

UXR is working on Play Protect Consent https://docs.google.com/presentation/d/1FiUQydjLA4Ve4Jmlz5spOzVNGffNxrthIJkz4KMwyDU/edit

Then FMD

Rodrigo: security checkup.


## Peter.md on 2020-10-13

Permission revokation dirty code complete EOY, full code complete and bugfree
EOQ1.


## Linus.md on 2020-10-13

https://www.google.com/maps/place/312+61+Mellbystrand,+Sweden/@58.2770599,11.051495,6.8z/data=!4m5!3m4!1s0x46519c4e14c51b81:0x26019078bd9b5d72!8m2!3d56.5046038!4d12.9471579

System Alert Window.


## Krish.md on 2020-10-12

go/road-to-angel
Planning update

Project Clear - all of privacy, permissions, ads privacy, etc.

Need at least one use case built on top of the platform.


# Week 41
## Khawaja.md on 2020-10-08

- Richard is leaving. Headcount.
- ASL weekly survey - update, outcome
- Angel update
  
Give resource asks.
Start looking for SWE for FMD.


## Meerkat_Leads.md on 2020-10-07

Updates from leadership:
- Planning process: have eng to be more involved. 

- Have more information about _why_ what the team is doing is difficult - need
  to strike a balance between  going too deep into details and "just trust us
  this is difficult". 

- Expectations of flat HC for 2021. Unclear if this will hold since the same was
  promised for 2020, but that's the general direction of where the conversation
  is going right now.

- Khawaja wants teams to start working on their team charters - why the team
  exists, what it is doing etc. How can we best handle this?

Team health update. What can we do to improve the situation?

go/sre-charter-template


## Anita.md on 2020-10-05

Working on enable device admin action, not needed for new versions of GmsCore
and since we are just reimplementing still need it.

Last few weeks - enjoyed, learning, writing code, debugging. 


## Niko.md on 2020-10-05

What's in flight

- 8 or 9 flags in flight for the moment
- My Apps - preparing for Wonderdome
- Separate launch for Wonderdome not controlled by us
- Snackbars
- Auto-disable who users who have disabled notifications
- Notification behaviour (Marie's work)
- Banner (Karishma)
- #8, #9 - logging changes
- Installer card on stand-by because VP
- Suspended apps - rollout plan for that.

Starting new projects which will take time. 
- Auto-revoke
- Auto-disable apps which have safe update - still in design
- ODML cards, requested for a while, in design

Experiment freeze by EOY, have to get stuff out of the door.

Discussed with Raissa and Karishma, reached out to Henry for help.

And then there's whole TV situation. Ended up pausing both auto-revoke and
update projects, which are P0s.

Too many redundant steps in the launch, too many redundant docs.

Warn on all MuWS on autoscan. Mostly Henry, mostly process.


# Week 40
## RichardD.md on 2020-10-02

Citadel, hedge funds.

Hands notice some time in early November?

Merge Spot and FMD teams. 

Richard recommends Anurag.

UWB - they are asking for commitment to a project, Richard and Guanhuan are
worried about amount of impact this will ever have. Even Spot - no clients.

They are talking about funding this. Most of uwb will be done by David. FMD
needs to support David's team. We need to become a really solid first party app,
preinstalled Pixel app. Support dark mode, be material 2 compliant, etc.
Transformer team is doing it for Google's first party products.

App was rebuilt in 2017, Dagger and stuff. Don't have to rebuild it, practically
UI refresh.


## Dave.md on 2020-10-01

* Update on Advanced Protection
* Update on Angel
  
Planning is starting. 

Files!!

General indication - flat HC. 

We need to know what is the delta for funding, and this delta has to be
believable. There was lack of belief from Brahim. He was saying he didn't feel
convinced we really understood resourced needed.

Files: Dave thinks this should be handled by maybe performance team? Let PMs
talk.


## kshams-leads.md on 2020-10-01

Eng leads to play more active role in presentations

2-3 slides to highlight technical complexity that might not be visible!


## Jim.md on 2020-09-30

At 50% ring, but success rate for Timon dropped.

Lock is in teamfood.

Locate is being worked on at the moment.


## Serban.md on 2020-09-30

pKVM - one of the leading projects for Bluechip, probably not p21 but p22. 

Inherited what Xiaowen used to do, platform hardening features for S.

Those are mostly collaborations with other teams.

MTE - memory-tagging extension.

Eng teams - jeffv partner for platsec, will deacon for Kernel hypervisor.

Isolated compilation is one of the killer usecases. Personal AI.

Phishing detection - team in seattle, Tom Hume.

Wedson is doing remote key provisioning. 

## Guanhuan.md on 2020-09-29

Spot / Tile - biggest requests for FMD. Everything is blocked on Tile at the
moment.

They want to interleave our e2ee with their own protocol. They'll have new tags
they'll create for Google as part of their network. This protocol wouldn't meet
Google's privacy bar.

They want Tile to be the only app ringing the tag.


## Peter.md on 2020-09-29

Marie hasn't owned large project yet.


## Narayan.md on 2020-09-28

We've been very hesitant to add stuff to existing MADA agreements, so need to do
our due diligence. 

Reluctance of adding new APKs. Taking something out of existing MADA APK and
moving to another APKs. Lots of sensitivity around it, need to start having
these convos as soon as possible.

Need to speak to who? Unsuk is CTS person. Someone on TT's team (who's that?).

Audit the APIs!!

Once have something working worth engaging with Samsung.


# Week 39
## Jan.md on 2020-09-23

Came back from Croatia mid August 5 days before restrictions

Working on toast in platform.

Folks are not very familiar with safetynet.

Alan - great as technical support / any type of support. Quick to assist
whenever asking. Giving Jan semi-regular feedback on how Jan's doing. 

Career conv Alan asked platform or server side, Jan said server-side.

I recommended Jan to try his hands in many things.


## Karishma.md on 2020-09-23

Working on rolling out 2 features:

1. Card in details page which asks user to turn on GPP
2. Re-design of MyApps (before Wonderdome)

Marie is working on Wonderdome. Karishma is working on simplifying the code a
bit.

Dealing with some issues. The details page card has - resulted in a little bit
of latency in opening details page, working on the fix for that. 


## Sorin.md on 2020-09-23

Son in Year 1

Wife managing care home

Richard is taking more strict approach to sprint planning this week.

Maybe improved success rate? 10% improvement in success rate. Unclear which
change triggered it. 76% for old FMD backend, 86% for Timon.

Date X for Timon? Milestone 5: 26/12/20

go/timon-pp


## P&E_leads.md on 2020-09-23

No news on return to office due to recent news. Subject to conversations.

Were planning to do a voluntary pilot - if some teams interested in coming to
the office for 1-2 hr for brainstorming etc. Were going to do it in 6PS. 

Right now discussing whether to put pilot on hold / proceed, based on HM
government. Please keep this for P&E leads team.


# Week 38
## Monir.md on 2020-09-18

2021 planning. Bottom up planning, lack of direction.
L2 -> L1 helps to group settings.

Better story about Alan's impact.


## Dave.md on 2020-09-17

- Update on Angel
- Update on Advanced Protection

Good perf is not about eliminating biases, but recognising them and despite the
fact they are there still getting to the most equitable answer.

Pushing people down vs up - that happens every single time, because bias to give
people higher scores. We are consistently higher than rest of Tech.

AP4A
- We've hit the big things, lone wolf, not enough support - makes it harder to
  push very hard on this.
- Before we really do that Dave would advocate to have meeting with Shuvo and
  other folks who work on Carter++, have 30 min chat. If they say it's so
  important we continue working - that would be good for senior sponsorship. 
- Care about high risk users. Also care about curating store. Filter was very
  important test bed for that. If we completely stop working on that part of it,
  what that really means we are 100% dependent on the rest of ASE team to take
  this forward. Personally Dave thinks this is the most important problem in all
  Android & Play.

Instant notifications - Dave doesn't think it will be a big impactful thing.

Protect++

How to get data whether this would work for the users?

Solarwinds done creative UX research projects to try to ask questions.


## kshams-leads.md on 2020-09-17

- Smaller groups
- How to resolve conflicts in borderline situations
- What's is the appropriate time for candidate
- How different people use quick calibration scores
- How should we train our organisation to take quick peer review signals into consideration


Discussion about structured vs unstructred

Apple: 3 main categories ranked against:
- Results
- Innovation
- Team work

only 3 levels: Not meeting, Meeting, Exceeding
3xExceeds = Superb, rarely given, means performing level above

This simplifies the process, people start getting more direct feedback.

Once a year

Cons - if you have a bad manager, not getting any feedback during the year and then
bam. To counter: you can't give surprised to people, whatever you are giving in
feedback must be some trail of how you communicated to people.

Do we have too many ratings? 


## Adam꞉Roman.md on 2020-09-15

Henry is saying: Richard is frustrated and has negative vibe, Henry playing
peacekeeper as he's worried this will impact the team. Henry is not comfortable
giving feedback to Richard. He's worried that when Richard gets feedback he'll
react negatively. Adam: Richard should handle it he's L6, it's conversation for
Roman.

Henry is freaked out that this is a bomb which might blow up. 

Henry feels a lot of pressure to mediate / keep Richard happy, unclear where
this is coming from. Unclear why he decided why this is his remit, he's too
polite but this is not his work. He's in awkward spot and asking for help.

## Narayan.md on 2020-09-14

If we want to push updates to Angel in S.

There are no new mainline module in S.

Team made decision not to support any new modules in S, will change in T.

So shipping updates under Mainline umbrella will be a non-starter.


# Week 37
## Alan.md on 2020-09-10

Incremental changes for SAW:

- Stop granting automatically even to pre-loaded apps
   - Apps have to be able to cope if user turns it off!
- Something for priv-apps which really really need SAW, move them to something
  else that might not even exist. If they absolutely need it they shouldn't use
  permission user can turn off.
- Flag "I don't want to be overlayed". We want any app to be able to say this.
- Bubbles only exist for one very limited use case. We showed the team more
  usecases, but they didn't commit.

Jan:
- Doing couple of toast related things.
- Very very small detail internal fix
- Rate limiting on toast

For next 1:1


## Henry.md on 2020-09-10

Working from Bristol for couple of days

Extending autoscan to cover MuWS - big project, talking to Rahul and other PgMs
in MTV. Org thing. Niko was doing collapsed notifications. A lot of these things
are learning for Henry.

Richard does great job on taking on things team doesn't have to do.

Tends to respond negatively when mgmt chain doesn't do something he'd expect
them to do. E.g. when FMD closing session was cancelled "they don't care about
us anymore" - this was unfortunate. This doesn't seem to affect the team.

Preparing for messaging to prevent negative. Henry puts a lot more effort to
describe things he puts to Richard to ensure he wouldn't be misunderstood. 


## Simon.md on 2020-09-08

Spot Bonus for TV work for Peter and Niko?

Simon: "Discussion around HC was particularly tense, not the right effect on the
team, people uncomfortable, you seem upset in front of the team damaging."

Roman has done:
- Rubberduck
- No HC .. we need to request HC
- Not comparing FMD to Aleksandr


## David_Coffin.md on 2020-09-07

Need to change views for warning ribbon, the actual text displayed with the
little triangle. 

Harder not to be a part of a team


## Anita.md on 2020-09-07

Starting new RPC
couch 2 5K

Last 2 weeks challenging because perf and bug triage

Summit

Tensions about being ambitious (which motivates people on the team) and not
achieving stuff (which de-motivates some other people). Anita on the side of
let's be less ambitious and more focused. Jim is leading this sprint. Several
sprints were over-ambitious. This sprint is about tying things together. Jim is
more on more focused side. 

Anurag and Sorin and speedy at doing things so they expect same pace from
others. Anita is not as accomplished at server side.

Richard is quite supportive during planning. He said in 1:1 good job keeping
team focused. Today he was supporting Jim and Henry on being more focused.


# Week 36
## Anurag.md on 2020-09-03

Feedback for Richard - prioritisation on FMD can be better. Things like Secure
NFC etc. Now Richard spent quite a lot of time thinking of what FMD should
really do. Thanks to Henry for driving the prioritisation  process.

Richard gets a bit confrontational if you disagree with him. On FMA disagreed on
one design, and conversation became quite a lot confrontational. Now Anurag has
to warm Richard up to the idea, if he doesn't like it it gets really hard to get
across => take quite a lot of work.


## Bernardo.md on 2020-09-02

Prevent status bar abuse

Prevent overlay abuse

Phishing != Overlay attack

https://docs.google.com/document/d/1hMljofsTkPRYvJXsb8xCT3Otdgbr7IMKblbNuG2PBNI/edit

SAW is a source of abuse but there are legitimate use cases
For all other use cases no evidence passing touches to other UID is useful


## Simon.md on 2020-09-01

In Plan B, what % is throw-away work?


## Guanhuan.md on 2020-09-01

FMD - PMs were hesitant about it.


# Week 35
## Khawaja.md on 2020-08-27

Angel - goin in the right direction, this is painful. 

https://hbr.org/2018/11/5-ways-smart-people-sabotage-their-success

70/20/10

Emotionally Intelligent Manager
What do you need to do differently to have more productive relationship with
people?

Thinking fast and slow


## Giles.md on 2020-08-25

His mum is in London.

Off-shoot of Delphi - Xiaomi permissions hub. Privacy incident. X took code from
some Q branch that leaked into AOSP and built permission hub and mic/camera
indicators into their latest devices. 1P apps were looking terrible (which is
why PermHub was killed). Code yellow on 1P apps to investigate/fix permission
usage. Planning to put PermHub in QPR1 of R.

Giles has an idea: what if we have a similar thingie for network traffic in
Delphi. We are building network traffic hub, that allows to see which apps are
sending the most data upstream, slicing it by some traffic/time series. Like
Permissions Hub but for data output. Want to launch initially as dogfood app, so
1P apps can see how they stack up against other apps.

If we are launching PermHub and this thing.

Android S: accelerometer rate limiting, limiting data to 200Hz, and put raw
accelerometer under normal permission. Setting that will prevent downgrade of
cellular connection to 2G, because for 2G 10x cheaper surveill equipment than
3G.

Policy on privileged permissions.


## Running_notes.md on 2020-08-25

Apollo / Covid notifications


More links for David, hashtags closer to work


Project Anning. Various directions of privacy safe measurement and targeting on
Android. Interested with everything Android privacy as relates to Ads.


## Jim.md on 2020-08-24

Went to Isle of Wight for a week, could swim. Steam train, zoo, etc.
Went to Norfolk
Bank holiday - seeing in-laws, not staying with them, renting a cottage, central
Wales.

Work's been good lately. Doing 2 weeks sprints, had couple of productive
sprints. Just starting perf right now.

Team is getting some momentum behind Timon. Had some very good and informative
failures. Launched dogfood for ring and then had to roll it back due to a bug in
auth. FMD is shockingly complicated. Was working on impl lock in new backend. In
theory should be simple but not. Example: some bug filed long ago, to not allow
emergency service number as recovery numbers. There's a library which allows to
find those, but requires region code. Getting region code out of apps framework
requires a lot of steps.

3 different paths into FMD: web app, app, LDA/Assistant, and they reuse
surpassingly little code. 

Team had various ups and downs. Doing alright. FMA's been taking a lot of time
from Anurag, and a lot of interrupt time.

Discussing, having debates. Rotating leads - one in charge of setting
goals/sheperding team to get these goals accomplished. The way sprints are going
depends on the person.

This might be bad if we had massive fights over how to deploy etc, but what in
fact happens in Jim's two weeks focussing on getting smth out and deploying, and
then Sorin would want to focus more on the development instead of launching
features arguably ready to go to dogfood.

Richard is pretty comfortable with letting team to hash things out within wide
boundaries. He'll nudge things towards centre if team is going too far in some
direction. 

Feedback: R does a great job of careing about us personally and work with each
person individually to make us more productive. Massive amount of work to
organise bugs into something more manageable of a taxonomy, next steps - take
this work and try to make progress on what needs to be made progress on.

On a good week productivity is higher. Back to normal.


## Running_notes.md on 2020-08-24

https://docs.google.com/document/d/1AzKjTyhRsKvjB4xBkmL2d7xjYQlVGQ4w0N67m1PFLHQ/edit
- sample promo packet L4


# Week 34
## Peter.md on 2020-08-20

Finishing TV UX work

Had a meeting with Guanhuan about longer-term goals of TV/Carsky
Guanhuan is a PM for GPP beyond mobile 

V1: PRD
https://docs.google.com/document/d/1oUcm5uMEd6zViRg88JJsH760ghR3mqnagg0bICNLM5A/edit#
1P
https://docs.google.com/document/d/1uOix24Ah5hn1P2WqU_nLOBFj_5pQCmzBmNWFSv-QDx4/edit

V2:
https://docs.google.com/document/d/11qu3qgxCz2lsf3M3rQKNFwwqR1tg2fhGEgJtrQohD4c/edit
https://docs.google.com/document/d/1Ulbpm4Qp8K0YvQKo-HBoMWf2o9IVmSsOQ56ljuaiTss/edit?ts=5f2ae790


## Marie.md on 2020-08-20

Replacement for MyApps and Games.

## Simon.md on 2020-08-20

Simon to draft an email.

Staffing Angel:

- Marie
- Half or most of Karishma

Aleksandr:
- Peter
- Niko
- Bessie
- Raissa
- November: Inara
- TBH

Dropping auto-revoke (Peter), auto-disable and update (Niko + Karishma)


## Running_notes.md on 2020-08-19

We need to build framework for contribution first.

Everything Pixel must be code complete end of December huh?

Set up contribution framework with dependencies.

There's some meaningful deadline in December but we don't know what this
deadline is.

We need to establish our deadlines better.

What's our MVP?
What deadlines we commit to?
What dependencies?
Resource gaps (PgM, UX)

Need this week.


## Narayan.md on 2020-08-19

GMS Express


## RichardD.md on 2020-08-19

Preparing for FMD summit.

We have our own oauth?


## Running_notes.md on 2020-08-19

Monitoring, reliability - hope someone was taking notes on this?

Krish: 14M users / 10% Android impressive but speak of opportunity.
We want this number to be 3-5 times higher.

How do we know 14M unique users a week? We store records for 29 days keyed by
Gaia.


## Dave.md on 2020-08-18

Pixel loaned head: they said they can't get the loan to us, 2K person team, they
might have backfill somewhere. We loaned to Ramanan and he can't give it back,
Dave said look at entire Pixel team.

Suzanne volunteered to try to continue push for solution. She'll talk to
Sabrian/Seung. She's keen about Angel in London.

She doesn't want GTW implement something for Angel that we'll have to maintain.
She's happy with LON doing Pixel first. Or in GTW - one-off Pixel-only.

Ask Krish.

AP: been talking to Khawaja about the filter, and app/dev risk folks. He'll have
someone in his team working on dev tiering. 

Dave's #1 goal for AppSafety: have an algorithm that can do a good job of
curating the store. AP was plan A in his mind. Dev risk is great but one also
has to look at the apps. Khawaja doesn't have anyone specifically working on
this.

Once filter is out want everyone here teamfooding/using it, tweaking it, making
it better and using it every day.

We are working on getting rid of location history altogether.


## David_Coffin.md on 2020-08-17

Problem with one of the install flow: bottom sheet.
Let's cut it.

Coding: working on UX changes, created AP warning. How to make See Details
green? 

Stalkerware working group: something would like to be part of in some capacity.
Had to chat to Olivier tcn@ about what David might be doing. Won't have 10% time
until this is launched but after that might probably have time. What technical
feature David can do?

I didn't tell David about Angel.

David to start attending their meeting. 
David's proposal: unblockable notification channel into Android Framework.


# Week 33
## Henry.md on 2020-08-13

10 days in Sardinia

FMD summit - Anurag was involved, the rest of the team wasn't. 
Techtalk

FMA goes live today? In-flight.
Timon M1 - complete? Henry doesn't know if we are in production yet.

Aleksandr: not very involved because everyone is doing good job. 


## RichardD.md on 2020-08-12

Dead set against A/C

People are running sprints now, PgM-of-the-week
Anita did good job there, runs meetings.

It was all Jim's idea, but he only had a chance to run half a sprint.


## Khawaja.md on 2020-08-12

The proposal was Pixel team coughs some HC, we also loaned them 2 HC last year.
They only have 1 left, one of them they lost. 

Perf: calibrations 3->4 will not be happening within ASE, as other orgs may not be able to do
it themselves. 3/4 largest. Kshams will be doing 6+. 

Kshams very engaged on TikTok, WeChat. 

FMD to do techtalk.


## Niko.md on 2020-08-12

Pushed to reprioritise things to get things out of the door.
Spending less time:
- Auto-revoke, updating PHAs on Play
- Auto-revoke not time sensitive, lots of work. 


How Niko is doing about his TLship?


## James.md on 2020-08-10

Privacy Shield: Raissa owns it now but def needs some help, maybe someone should
take over it. Deadline 01/09. 

Jessica wants to know:
- Why GPP click logs needs to be gaia tied?
- Wants clear expl what's Gaia, what's logged by snet, and what are we doing to
  not provide any means of joining these identifiers. 

60 day deletion of data. What does this require us to do in order to use this
deletion plan? 

Rumours of HC conversations.
- Rumor 1: Google TW
- Rumor 2: Andrei was given some HC, allocate some to us
- Rumor 3: DaveK pull some Pixel HC
- Rumor 4: Dave/Suzanne get some HC
- Rumor 5: each team underneath Angel gives us 1 HC


## Ed.md on 2020-08-10

Sideloading discussion: another meeting on Tuesday last week with Sameer, he's
having a meeting with Epic Games CEO, no news. On Tuesday sameer got an
impression - he doesn't want to allow third party billing experiences on Play
Store -- which is what some people are pushing for, and rather than risk eroding
Play Store he would rather us take a critical look at 3P store experience, to
make experience more in line with other operating systems. You get lots of
prompts on Windows. 3P app store can't deliver updates seamlessly. 

On one hand Sameer says - let's make it easier for app updates to happen.
Regression on security. The thing that made DaveK happier was that Sameer said:
what if we did something along the lines of - if users enables unknown sources,
3P stores work a bit better, but then our equivalent of Windows S mode,
something that Pixel team wants?


# Week 31
## Dave.md on 2020-07-30

Dave met Brahim and Sagar, security leads meeting. 

Bluechapel is an old name, some religous stuff. Bluechip.

Edict from up above for Bluechip that they should build the best device they
could build, regardless of what we do for the ecosystem. All sorts of UI things
they are building that are not going to AOSP. Prioritise Bluechip above AOSP and
everything else. 

Option 1: build Angel prioritsing Pixel
Option 2: they build whatever they want, throwaway, one-off, as we do Angel
they'll be forced to harmonise their center with Angel. Acceptable from P&E
leadership point of view.
Option 3: "they" go and implement what they think is Angel for Pixel but for the
ecosystem. Dave dead against it. Folks kinda agreed with that.

Brahim thinks LON team is over-sensitive and defensive. Brahim says his #1
priority is option 1. 

Pixel product team hasn't defined what they want. Entirely possible it's not
Angel they really want.

Need to add links to one-pagers / PRDs to the sizings doc so Dave better
understands what's this is about.

Dave: can we still build it in Phonesky? R: no, because reasons.

Apparently there are 2 HC on loan from ASAP to Pixel. Dave thinks what can he do
about it, he's to send email.


## Alan.md on 2020-07-30

Another late fire in R related to public volumes / SD card. Never tested b/c
Pixels don't have one.

Jan's perf:
https://docs.google.com/document/d/176KZRpbUjWJewb4anEqDIzmpCAubjkbdeo8ZGQxndy0/edit?ts=5f187806

Had meeting about scary access permissions.
Everyone agrees it should be done.
Philip says this didn't make the cut for his team for S. 
Abishek very keen on it, cause makes UX much better
Ed: MVP - just change it in settings
Only 3 perm we care about, can just as well move it to permission controller
Philip is OK with us to do this.
Ed's been talking to Zach Koch, trying to persuade him to help us out.
For now this is unstaffed.


## Joanna_Jakobska.md on 2020-07-30

LDAP not generated yet. 

Need to request SkyLab device or two.

For the first week - just invite them via personal email. Keep the list of what
will be shared via personal email and revoke access later.

h.y.abdi@gmail.com


# Week 30
## Karishma.md on 2020-07-23

Apps page card
Simon's work - tweaks to myapps card
WOrking with Niko on auto-disabled


## David_Coffin.md on 2020-07-22

AP4A: 10% by EOQ

Excited about working with team.

21/08 - cut off for strings.


## RichardD.md on 2020-07-22

Still problems with Jim motivation.


# Week 29
## Linus.md on 2020-07-17

Struggling with motivation


## Henry.md on 2020-07-16

Meerkat Band in Daily Insider!
Support two products for long time
Should Henry look at metrics/dashboards?

Anurag is back to Timon and this is great

FMD is much more collaborative than Aleksandr (because different nature of
work). They do nice job of that. Managing work: notable difference, e.g. Sorin
and Anurag vs Jim and Anita.

Last iteration cut half of tasks because we had data that we never achieve it.
Some iteration where performance against plan was very poor.

Grade 0 to 10 - where's FMD = 7. Collab and cohesion in the team is good. Often
about task management/dealing with uncertainty.

To get them to 8: need more time. Some people will just make up their own tasks
and make them happen. Jim and Anita are different, finding way for them to be
independent and driving their own thing.

Met Alan. Doing bi-weekly. Alan asked for some thoughts on what he's doing on
planning for S and OKRs, gave some feedback.


## Marie.md on 2020-07-16

Over the hump of WFH productiveness.

Landing notifications in GPP Home. 20 lines of code from code complete.


## Khawaja.md on 2020-07-16

Notes from James:
- jfkelly was given by Krish right of veto if we don't get staffing
- does kshams fundamentally believe his engineering team should be working on
  this stuff?

Yes it does. Question is - whether we'll get heads to work on this.

Khawaja to go and speak to Dave and Suzanne.

His only concern was how we'll be able to justify to Sameer why we are working on so many things (apparently there's a lot of pressure to put more wood behind fewer arrows).I explained him that one way to look at it is that Angel is essentially GPP v2 - much more expanded, with significantly complex dependencies on platform, but ultimately a necessary thing to continue pushing our message; this naturally involves long transition period of two products co-existing. He thought this is an excellent way of explaining this and agreed with the concept.
He is committed to try to produce headcount for us - he said he'll be speaking to Dave and Suzanne. He asked whether 1 HC can help us (i.e. things would just take longer) - I explained that it wouldn't because hard deadlines, he agreed.
In general he keeps his cards pretty close to his chest, but to me it did sound he's on board and wants to help us
 

## James.md on 2020-07-16

Stephan shouldn't be doing here.

offline? we need Google acct stuff to work offline.


## Nicole.md on 2020-07-15

Had very little time to build this thing. Pitches project end of January,
announced in May and launched in October. Any APIs need to be done by January.

Pros of being in platform - don't have to worry about releases. Digital
Wellbeing has to manage their releases, crash reports, everything. Restricted in
what access you have - if you are in Settings you are Platform signed.

Benefits - Play Store updateable.  Big overhead.

Alternative - Extreme Battery Saver. APK within the system. Injected into
Settings. Gives more flexibility. Pixel only. 

Put into vendor/ directory. If leadership decide this is what they really want,
they can give you extension on your time. Update route - QPR. They don't like
permission changes, don't like new APIs, fairly tightly bound release, but this
is an opportunity to update. December, April, etc.

Digital wellbeing - would keep it where we put it. Wouldn't get anywhere if we
did it in the platform.

If Play updateable - lots and lots permissions will have to be granted.

Would not put into settings APK, the reason: because then you own it, have far
more autonomy on how you are doing this. Settings is AOSP, where this is Pixel only.

## Guanhuan.md on 2020-07-14

privacy shield? all covered

FMD summit: was a bit of pushback from Richard. He still wants my staffig.

Guanhuan expectations: alignment and communication from FMD side.

My stand: traffic light model.

Where is the list?

https://docs.google.com/spreadsheets/d/1A2Q53wtb6jgg0ihR4pHnp45S7-aPxVbL6bnH3OoPfuw/edit#gid=0
- FMD projects


## Ed.md on 2020-07-13

P21 Security Plans slides

If need to free up 1 SWE: SAW will have to go.


## James.md on 2020-07-13

Carter - we want to work on opt-in higher security mode.

One-way road for higher security.

off-Play review has stopped.


# Week 28
## Lisa_Martinez.md on 2020-07-09


## Alan.md on 2020-07-09

Ricky's SD card stuff - roadblocks. Enabled in Master broke a lot of stuff.
Rolled back for now. Ricky is in HKG right now. This will be his priority for a
while. 

Linus is struggling a bit with motivation. He wants few small concrete things he
can achieve so he feels he's done something.

Bluechapel: unknown sources work. Ask for FR!


## RichardD.md on 2020-07-08

FMD summit would be useful, we need an alignment on FMD around ASAP.


## Simon.md on 2020-07-07

Handover:

Simon is back August 12th
Karishma's previous rating was SEE (although not rated)
This one would be SEE
So N/A, EE, (SEE), SEE.

Sorin: L4 to 5. Hasn't really done any promo exercise. Doesn't have many L6
SWEs.
For the last few months running Timon. Mainline might be a good opportunity.

Marie - notifications land in GPP Home.

Raissa - PEP is done, met expectations. Still does have tendency of having
days/weeks when it's not clear what happened.

Timon - Ring requests going through new backend.


## David_Coffin.md on 2020-07-07

Fitting into Q3?

Possible concerns from device health about loading 3-4 Mb data in RAM.


# Week 27
## Alan.md on 2020-07-05

NW - Quokka people are taking seriously.


## Marie.md on 2020-07-05

Dev process complex. Updating behaviour of notifications to go straight to GPP
home.

Delighted with the team. There are things to work on on code quality - package
verifier in particular.

How team is doing - everyone is very supportive and productive. 61 days at home.

Small interaction with Ben and William.


## ASL_weekly.md on 2020-07-05

Main talking points for today:

- Updates from the leadership:
  - When are we going back to the office?
    - PM to speak on Sunday. For what it's worth, I am not expecting him to
      announce that we end lockdown. 
    - Even if we do, it will be most likely a gradual process.
    - Staged return to the office:
      - IRT
      - SREs
      - SWEs who need to be in the office the most
      - After that, likely reduced capacity of the office. Don't know anything
        beyond that, so let's all stay tuned.
  - You've all read email from Andrei about important COVID-19 related efforts
    happening in our office.

> Exposure Notifications (Contact Tracing)
> 
> You are, of course, aware of the Apple - Google collaboration on Exposure
> Notifications API. The Android Context team in London is contributing by
> developing an open reference design for an Android app using the Exposure
> Notifications APIs. You can find the code on the project’s GitHub repo. The team
> is partnering closely with DevRel and EngProd to make this design useful to 3p
> developers.
> 
> 
> NHS Life Lines
> 
> Folks in Android Enterprise Partnerships and Eng worked with BT and Samsung on a
> project that helps the NHS deploy tablets in every single Intensive Care Unit
> (ICU) across the UK. These tablets allow families to connect with their loved
> ones who are being treated in ICUs for COVID-19. Our Googlers provided project
> management and technical expertise (the tablets are managed using Android
> Enterprise technology) as well as funding for the tablets. You can find more
> information on the Life Lines homepage, on Twitter, as well as the BBC.
> 
> 
> COVID-19 apps on Google Play
> 
> The Play Dev team has built a system to help with faster reviews of important
> COVID-19 related apps, such as those from government and health agencies around
> the world. They have also partnered with other teams in Android to assist with
> the review and approval of government apps taking advantage of the new Exposure
> Notifications APIs mentioned above.

Finally, I wanted to mention something important from results of our survey.

  - Thank you!
  - Results are anonymous, even I don't know who wrote what.
  - Will show once have more results.
  - People going through some really difficult things, things no one should -
    ideally - ever go through. This is reflected. 
  - When asking directly everything fine. Survey shows it's not.


## Simon.md on 2020-07-02

Richard seems better, although not 100% significantly improved. As the result of
that he's even less likely to take any time off. Khawaja didn't have a chat with
him. 

Everyone's keen on idea of having a summit. End of July probably. 

Problems of roadmap/planning for future are still there.

HR: refer him to occupational health who will assess whether he's fit to work.

Sorin wants promo
Karishma wants promo

Jim is heading for NI

Some productive conversations with James. 


## Pieter-Paul.md on 2020-07-01

Carter started 2 years ago, built by hijacking team. 

1M accounts, identified based on similarity to accounts we observed to be
targeted.

another 1.5M account - product risk, YouTubers, app developer.

Big thing for 2020 - U.S. elections - 200K accounts.

Users do now know they are in high security group, with exception of Carter++ -
opt in form, we don't advertise it, but through Digital Democracy (NPO US helps
elections get their digital security right), just few 100 submissions.

Medium term: Carter stays as it is, long term not sure, do not anticipate it to
disappear. Carter++ might fall back to the regular carter after elections are
over.

Alexei Czeskis is overall lead of Carter.

## James.md on 2020-06-29

Krish wants Meerkat update

Impressions for treatment group 4% for Oleg, rescan button click 86%
PHA removal rates/suspended/uninstall MuWS rates all are up 30%
GPP disables are down 8%.

Brahim figuring what they want to do for Pixel 21


## Meerkat_Weekly.md on 2020-06-29

Thoughts for the meeting: 

- Shifting to more async model of work. What are your thoughts? Conversations
  with other leads:
  * Having 3 streams to work on: current work, medium-term and long-term work.
  * Async communication. Many things really don't need a meeting. Controversial
    idea: do information gathering via email thread / make decision in the
    meeting. Optionally substitute email thread to a group chat, but the caveat
    of urgency.
  * Interesting conversation about IM: how instant should instant messaging be?
  * Thoughts on remote brainstorming


## Karishma.md on 2020-06-29

Oleg at 100%
walking around town
india cover-ups


# Week 26
## Monir.md on 2020-06-26

Have 3 things running in parallel:
 - Current one
 - Medium-term
 - Long term

Expectation:

You can reach to someone, Chat or Email? What are the expectations?

Should we do a survey on how instant is IM?

Emails vs meetings?

Data gathering happens over email, decision made in a meeting.


## Dave.md on 2020-06-25

Angel + Bluechapel:

1. High level - there was a meeting between ASAP, PlatSec, PJ, Pixel Software.
   We came up with initial first pass of top features / narratives for P21. Has
   to be reviewed, brainstorm data. Informally voted. 2 top issues:
   * Having Pixel recommended store we talked in the past - got a lot of votes,
     very hard to land, odds of getting this into final plan are very slim. One
     of the questions (Brahim) -- based on the latest algo we have, how many
     developers/what % of downloads/revenue are covered, good to get that
     estimate. Dave thinks about 40K apps, ~2%. 5K/1% developers. 80%+
     downloads. A lot of concerns that household names won't be there. App
     Safety is the big problem for Pixel. Suzanne said that marketing won't go
     for this. Dave's answer: the narrative shouldn't be about safety, it should
     be about quality. Dumblestore has 3K educational apps, and all 3K should be
     on  our list.
2. Angel. General idea is - security centre that can be used as security
   check-up, we can help you get to the better state, where all the settings are
   .... There was a link to the Angel deck, but they are not saying let's do
   Angel, saying let's do Angel like. James' answer: we don't have resources.
   Given we are not putting a lot of heads, Pixel can add few people on that.
   All 10 people voted for it.

Ribcage: very hot right now. Apple just announced some stuff about ads privacy
at WWDC, we are behind. Dave is really excited, Dianne loves it, convincing the
rest of Frameworks. What ejc thinks about it?


## Anwar.md on 2020-06-25

Lots of thinking about Pixel 21 devices. Critical for Google. Custom SoC. Expect
issues with SoC, first generation is never good enough.

One thing presented the leads - some data going through with Android leads. We
are losing share in critical markets. Android share in US/Japan is collapsing
right now. How do we reinforce our position in these markets?

In iOS they allow to replace default browser and email.

Dianne - she feels we should be more aggressive in marketing against Apple.
Flexibility and freedom you get in the Android space. Apple is viewed as more
authoritarian parental figure.

Pixel devices margins are negative, not selling enough of volume to justify.
Pixel 4: cheapest is $800 (64Gb).

What is the USP for Pixel? Right now we don't have one.

Last years we focused on hardware features - EdgeSense, Soli, and we lost some
of battery capacity because of some of these decisions.

Privacy/Security stuff is important because it counters the narrative.

Integration with Assistant.


# Week 25
## Sebastian.md on 2020-06-19

Covid tracing apps. Most of them crap from security and privacy perspective,
because they were rushed out.

covid tracing from S Korea, was discovered by Googler who went straight to the
press. An ex-Googler contacted us about app from Bahrain, which is also crap,
code injections etc.

Kylie was tasked to explain execs how to bring up better structure. 50
directors/VPs.


## Krish.md on 2020-06-18

India/China stuff.
Influencers got into fight.
Thrashed TikTok rating.

Project GuardRail - how do you not take down sensitive things.

Krish is OK with continuing to invest into Carter.

Carter will not meaningfully change abuse landscape.


## Simon.md on 2020-06-18

FMD summit in July. James supportive.
Advising Richard to take time off, Khawaja will speak to him.

======================
For our next 1:1
- Situation with Richard. Summit?


## Alan.md on 2020-06-18

Linus: let's add more friction to turn on scary permissions. Ed's ideas: 

1. A11Y screen is quite good now, scary enough. Better than it used to be
   anyway. Different consent screens, let's unify them to make them sufficiently
   scary. Let's add uninstall button. Want to record stats on people saying
   yes/no.
2. Android absolutely hates injecting UX in that way, we can do anything useful
   in the verifier.

Alan: adding more friction will not solve the idea.

Another idea of Ed: rather than dealing with app once it's installed, let's stop
this app from arriving in the first place.

Random website can send you APK via Chrome and lead user through installation
process. Chrome shouldn't have any special handling for APKs - you've downloaded
it, it's in your downloads folder. Remove special logic in Chrome which
instructs to enable unknown apps.

Lack of enthusiasm from Client Software, unclear how to make a difference.

Finished document W^X. Victor and Richard interested. Attaching a permission to
running native code. Nick destroyed it and he's mostly right. Nick's point: all
attacker has to do to force an app to do the dance. Block execution of native
code unless permission - his point you can't even stop native code execution.

Nick also commented on Richard's let's ban this completely. Mentioned to Richard
ODED document. Richard wants to be sure that he can see what code is being run.

One interesting thing to consider: once you can execute native code, we can't
stop you. But to get to that you have to go through our API of opening so file
once.

Victor Hsie - taking over full stack integrity. He wants to kill DCL forever. 

Ricky going to HKG, will be working from there.

Jan is going on holidays to Croatia and then will work from there for a while.


## Khawaja.md on 2020-06-18

FindMy vs FMD numbers

No actual data, just anecdata.

Optimise first on usability and reliability, then on growth.

How to handle some of these crises, choice of language.

active listening
Empathy


## Jason.md on 2020-06-18

Conversation with Dave about future of PgM team.
Disappointed that Dave is not supporting him

https://docs.google.com/presentation/d/1jLb7L35m23N7cUGEERqH5qYvUyhHI_kQp9tCeKo2IR8/edit?ts=5eeba0bb

https://docs.google.com/document/d/12kEor2iYg2_1WAoI-voujQfTnDSkweOnLyalX7gIh84/edit?ts=5eeba05e#heading=h.7occowslfzc0


## Henry.md on 2020-06-18

Aleksandr: making some progress in being more connected with Play. Good lines of
comms with TPMs/UX teams over there. Hoping we don't get into too many messy
situations with our features.

3 -> 1 notifications: we were not allowed to use expedited process.

FMD: collectively managing eng processes, planning sprints, etc. Still changing
and evolving, haven't cracked it yet. Focus was on FMA, Anurag spending a lot of
time on that, losing focus on Timon. Did Q3++ planning.

Richard:

* Internal interaction with the team: really good, strong and clear leader, good
  and letting people speak, good at being decisive.
* Understand his frustrations about feeling of disconnect about FMD and the rest
  of ASE, doesn't feel like FMD has attention / understanding - Henry doesn't
  share this opinion.
* He met up with Khawaja last week. Last week he's been in pretty good form.

Previously he was clearly frustrated that FMD was not being appreciated, now he
feels we are making some progress in being better understand. He takes it very
personally. Found it quite stressful. 

He's good with Guanhuan. Outside of the team: if I was Richard I'd be more
proactive in reaching out. Just reach out and try to smooth it out, he doesn't
find this easy. He prefers to be reached out to. Need to do better work managing
up/sideways.

Try and reach out to Richard a little bit more.

Problem sometimes affects his mood, but doesn't affect his ability to manage the
team and get work done.


## Stefan.md on 2020-06-16

On holidays from July 7th
Need to speak to Fey about backlinks pipeline

Content folks for list of important brands
What brands do you want to make safe?

Speak to Zhao

Mokka = Mentor + Quokka

All review tools are PoS. Managing enormous amount of tech debt. Platform for
creating review tools.

go/mokka

## Niko.md on 2020-06-16

Can't use expedited process, for small changes.


## Simon.md on 2020-06-15

Richard was shouting to Simon in a meeting, not acting reasonable, despite
Simon's ask to calm down.


## GPP_doomsday_brainstorm.md on 2020-06-15

Things that can go wrong:

  * Bad build of GPP causes all installs to fail, incl Play -> huge impact to
    Google/user, difficult to fix.    
  * CSD starts replying "it's malware" to every single install -> high user impact
  * CSD is unresponsive/DDoSed, malware can be freely installed
  * GPP page won't render / FCing Phonesky for some reason
  * Autoscan stops running
  * Snet idle logs stop running



  Groups:

  * Server is down (CSD)
  * Really bad UI regressions 
  * Server sends bad data

## David_Coffin.md on 2020-06-15

Revised timeline? 

No permissions apps can still be harmful.

Make it behind the flag, will enable if we need to. Good enough for MVP.


# Week 24
## Dave.md on 2020-06-11

What if app just sends user to the phishing page.

Bunch of high profile escalations
Execs over-rotate to the stuff going on the media


## Alan.md on 2020-06-11

Getting proper estimates for S planning


## Anu.md on 2020-06-11

Anu, joined team recently ASAP PM. Charter: dev risk, SDK, content abuse,
families & sensitive verticals.

Andrew Raz and Kanika report to Anu. 

Prior to Google - cybersecurity 8 years, ExoBeab, start-up in cybersecurity, 5
years from 5 to 500 people, $100M+. Identifying insider threats, malicious
insiders. Fingerprinting users and devices.

Before that, security, Imperva.

Before that, storage for HPE. 

Not had consumer experience. Complexity of the ecosystem!




## Guanhuan.md on 2020-06-10

Can meet friends now!

Did take one day because didn't sleep very well.

This week/next week - dogfood out for FMA. Dogfood instructions etc.

Couple of flags users potentially need to set.

Thinking of ways to reduce the steps.

presto dogfood group.

estimated 70 responses according to results last time.

spoke to fastpair team - if Assistant headphones, Bisto team can help with
dogfood.

Hoping maybe get around 100 people all together.

Roham - Google LTE watch, next year.

Asked Krish and James, reassuring to hear from Krish level. Q: fmd is both
security and consumer facing product. We are sitting in app safety. For feature
requests such as UWB, Find My Friends which are strongly linked to customer
facing utility, would that sit well with FMD in future, because FMD is in App
Safety? Krish said: as long as the other team thinks this is important enough,
from our side it makes sense to be in FMD, if other team if willing to fund it.

Richard doing better? He's taking some actions now, writing documents, purpose
of FMD etc.


## Rodrigo.md on 2020-06-10

AP with legal, also with Play UX:

  * Alley-Oop - keep it but redirect user to the details page.
  * Play UX: why don't we warn them early on?

Decision: push-back, but maybe David can explore once start working on UI.

Legal review: Jessica wants to know more details about whitelisting. Need a doc
explaining how whitelist will be formed. How will developer do that?

Disable Play malware - user testing starts this week. Reviewed by Play UX.
Malware that was installed from Play but still on users devices, but we are not
warning because safer version is already on Play. We are disabling app and
nudging the user to update. 

Permissions revocation. 

Auto-disable card - project Raissa was working on. All approved, waiting for
Ariane launch approval.

Karishma is working on Enable GPP from details page. Was reviewed by Play Leads
last week, Bankhead had feedback on frequency how often do we show this card.

FMD <-> GPP cards deprioritized, UX wise nothing complex.

Changing confirmation cards to be snackbars. Marie is working on improving
notifications, e.g. instead of opening dialogue will take user to GPP. GPP as
sender, not Play Store, with the icon matching Oleg status.

UX TVC: e.g. FMD starts ringing headphones, I can see Guanhuan would want to
build more FE features, so there will be some stretch looking after FMD. May
limit Rodrigo's ability to do long-term things.

In terms of GPP OK for now.


## Khawaja.md on 2020-06-10

Clarified resourcing with Richard
Need to have a deeper conversation about vision
Very upset about rubber duck comment


## Linus.md on 2020-06-09

Want to go to the office.

The work is a bit slow at the moment: no specific deadlines.

Will be looking at a11y/special permissions.

daily stand up to chat about general stuff.

Find it refreshing to have small breaks to speak to people. WHen on my own get
bored and drift away to check reddit.


## Peter.md on 2020-06-09

a11y education program. Different levels one can reach by doing exercises.

Accessibility Dojo


## Sudhi.md on 2020-06-09

ODAD wrap up this week. Other teams jump in into on-device abuse.
x-Google effort content moderation detection.

Guliz - looking at cleaning scary permissions in framework
https://moma.corp.google.com/person/gulizseray

Dom, John Ayeres working on app registry
if that happens 


## Narayan.md on 2020-06-08

Mi: they decided to bring back permissions hub, something nearly identical
to what we decided not to launch in Q. Brought back permissions dashboard,
sensitive permissions indicator etc... Augumented with some background stat
information, relationship reporing who starts what - e.g. something in Photos
decide to spin up GmsCore. They've done it. Wanted to launch 2 weeks from now.
We had to walk them off the cliff.

AGSA now instituted code yellow on bg location and bg mic. Version of AGSA going
out next week which will address a lot of bg access.

Hotword - they do need access to the mic in the bg. But this is not sent to any
server anywhere.

Some app visibility side channels - disproportionate app compat impact. If not
for scoped storage, external app directories would be side channel. Some small
sidechannels can end up being a lot of work/breakages.


## David_Coffin.md on 2020-06-08

Request for parental leave approved. Planning on taking 1 day off a week,
probably Fridays.

Last week - talked about revised timeline. Haven't done full time table yet,
broken down dev work. Code complete by end of this quarter is still possible,
but tight.


## Simon.md on 2020-06-08

Richard wants HR complaint but happy to speak to Khawaja.

Richard's points:

* Primarily annoyed: discussion after prioritisation exercise, rubber duck comment.
* He feels very lonely/isolated because kids are in NZ, fragile coming out of
  his divorce.
* He said "don't have anyone to talk about this"
* Then I said I can post you a rubber duck.

Then Richard spoke to me again, expected an apology or discussion on how to
improve thing, but dismissive. That I'd told him to see Khawaja if there's
problem - he felt a bit like sent to the headmaster.

He still has problems with James not wanting to commit to anything more than a
day in the future, Richard is an opposite extreme to the James.

Had problems with the way prioritisation was carried out, but maybe didn't
explain the problem that well to Simon. More "how" it was dealt with afterwards.


## Khawaja.md on 2020-06-08

Sync with Richard - did it happen? Need to expedite.

virtual something app from India
Involved passing credentials, some odd behavior
Going to reinstate apps
Ops seen article, didn't find anything wrong
TikTok replacement in India

Intel team was able to replicate
Google login passes jwt to the server not using https

Some embedded code in 350 other apps

Mitron

Richard situation:

Simon said two things:

* What Richard told him, Roman was unpleasant, I'll post you a rubberduck and
  you can talk to that.

??

When do talk to Richard, show some empathy based on his current situation.


# Week 23
## Sai.md on 2020-06-05

NOT A MORNING PERSON!

Two things this Q:

* Shubra - review queue. Jan asked some questions about it. Replace manual queue
  with automatic queue. Just put a scorer now, all scorers are treated equally,
  if precision is good it is automatically started getting used for manual
  review. Moved bunch of Buho scorers to Shubra. Works if detection overlaps
  with existing detection. Doesn't work as good if this is a brand new
  detection. Goal for this quarter: have some kind of sampling, and have these
  apps reviewed to bootstrap scorer.
* Precision -- we use true positives ... not a good ranking metric. Let's say 5
  apps are flagged, all TP => 100% precision. MarmotAI has been flagging apps
  with 90 precision. Moving to more statistically rigorous precision, which
  gives weight to volume.
* Support Whistlepig -- they went their own way. Sai wanted Whistlepig
  effectively another scorer, so if AV is good, it is used for MR.

AVs don't tell us what to look for, just say this is Trojan -- which doesn't
always match us. DaveK suggest negotiating further with one of AV to start
getting more information/context.

WP is not very successful right now. Was not getting many reviews. Precision is
10-20%, recall is 3-4%.

Understanding landscape of preload badness. Initially starting with preloaded
PHAs, but after talking to a bunch of people, including Dave -- we don't know
too much about preloaded PHAs, those which exist tend to stay for years and
years because they are bundled. Marmot is not built for preloads. Looking at
preloaded data exfiltration - scope is more concrete, e.g. Verizon tried to do
it on Pixel, contacts and location. Are users aware of it? Most of the time they
are not.

System apps these days have a proxy API, which is guarded by some other custom
permissions -- apps which it blesses will get access to location without
location permission. Big enough of a problem that GmsCore reached out to mobile
malware to ask if they can do something about it. Working with Sruthi on static
analysis to find code pathes for SMS/CallLog etc without permissions present in
the manifest.


## Simon.md on 2020-06-04

Spoke to kshams. He'll talk to Richard, offer some emotional support. Try to
reassure that no one is out to dismantle his team. Discussed coming up with 2021
plan for FMD. He'll have chat with Krish re having slightly longer term vision
for many projects under GPP UX side of things.

Progress with various Aleksandr projects in flight:

http://go/gpp-home-projects

ASAP logging - writing click event logs into verifier logs.

Better Verifier API - verifier module in Phonesky, owned by William, old and
complicated. As added things over years (auto-disable, privacy cards) needed new
API from verifier to UI. Now we have a unified async API, POJOs etc. Not passing
bundles over Android API calls.

Auto-revoke: apps can declare in manifest I'm emergency app. This is gone. We
will auto-revoke everything. We are not doing anything for pre-R apps. Will ask
pe-xws-process and maybe then ask Hiroshi.


## Alan.md on 2020-06-04

R issue: isolated app processes, used by Chrome. Processes within same sandbox
are highly restricted. Isolated processes can ping /data/data, more restricuted
processes have more access. We think the fix is fairly straightforward.

glapstation wouldn't help

manager roundtable - was interesting. Didn't get much results though.


## David_Coffin.md on 2020-06-02

Config pushing: in ideal world we'd upload it to gstatic with different file
name, so we'd have one per day and we'd use phenotype to push out new URL to
devices. We can't do that because there's no automatic way to push out
Phenotype. Instead of that we'd use Phenotype to put people into experimental
category, beta/prod. They would then download a file, once a day, 1KB file with
the actual URL for the latest version of that file. Or we can put dogfooders
into beta group.

Suggestion by Stadler - push out string rather than boolean ("beta"/"prod").

Flynn: we probably didn't need to use // a lot of complexity in terms of
pipelines written dealing with beta group for file. He's like: you don't need
that. Doesn't make much difference for client side implementation though.
Reduces the cost if we need to do a beta group later.


## Niko.md on 2020-06-02

Need a break


## ASAP_manager_roundtable.md on 2020-06-02

Greg started working with laptop at kitchen table. Realised cannot go like this,
not sustainable. Seeing same pattern in the team, people moving around,
rearranging stuff etc. Some people had setup at home which worked for 2-3 hr of
work a day. Very few had full time WFH setup.

Do you have a setup which supports FULL TIME wfh?

Sudhi: WFH few dimensions:

- Ergonomic setup
- Access to resources freely available in the office but not so easily available
  when WFH, e.g. hardware thats being designed, corp resources.
- People part. People rely on conversations/whiteboarding. Not having this
  impacts project work.

Some people appreciate the fact that they can WFH, because swap home concerns to
work concerns. More time with kids.

OTOH people in shared accomodations.

Move to the country! Tax implications.


## Ed.md on 2020-06-01

Chatting to Zach next Monday.

Sabrina: can we use whitechapel to lock things down sideloading, default
launcher, SMS app etc?


## Sudhi.md on 2020-06-01

Vish wrote a doc, didn't share it.
PlatSec leads have very strong objections against on-device detection.
Dianne is sort of supportive but worried about scope creep / ecosystem abuse of
these privileges. 

e.g. APK itself has no privileges, platform updates module and sends verdict
back. Endpoints are clearly mentioned in the manifest. Protobuf for sending this
information both ways is well documented somewhere. Somebody else can replace
this component with something else.

Will Drewry had suggestion: to overcome introspection without consent, why don't
we inject oded file into the application we are inspecting, so it runs with UID
of the APK.


## Anita.md on 2020-06-01

Trying talkspace therapy


# Week 22
## Anurag.md on 2020-05-29

What's the uptime for FMD? What's the SLO for FMD? 

Health - same situation. Brain started filtering it out.

We don't have an SLO for FMD. Uptime would be pretty good, more than 99. RPCs
don't fail. What does fail - we don't get the right status off the phone, when
something fails on the client we oftentimes won't know.

In Q1 added clearcut logs to the app, need to add to GmsCore, we'll be able to
compute user-visible uptime for the service. GMSCore logging not in this/next
quarter. 

We do have Tier 2 rotation, 30 min response time. Trust into monitoring isn't
great. Lots of breakages are subtle. Timon will solve monitoring problems to a
significant extent, e.g. next sprint looking at user-visible actions alerting
for Timon.

Doomsday scenarios:
- Service goes down, e.g. bug results in RPCs not failing but no requests
  forwarded to the client (so not visible through alerts right now).
- The bug when Jonothan pushed updated so we were at 10% capacity, but because
  we were over-provisioned things were still relatively OK.
- Privacy doomsday - if you are Googler, you can currently do really bad things.
  
Richard was super-stressed last week. When he comes/looks stressed in the
meetings, everyone perceives this. Would like to see more long-term stuff
happen through Guanhuan/Richard, e.g. long email with many important people.
On my previous team everyone would have a sense of what is super-important
right now and what is important going forward. Step down a bit on eng and spend
more time on conversations with other teams?

[Thread on the bug in the prod](https://groups.google.com/a/google.com/g/adm-eng/c/dnqDVrjH-gw/m/t51tmvTdCgAJ)

[Postmortem on the last outage](https://docs.google.com/document/d/1QdjI8aBY1KMe8dW9rmqcoEmTxF5diiuOsibqfbfGmc4/edit)



## Monir.md on 2020-05-29

Curatrix, using backlinks for reputation

Talking about various stuff, how's safety, how are people, news, etc.


## Karishma.md on 2020-05-29

Lockdown lifted in her city, hit quite badly.


## Sorin.md on 2020-05-29

Timon M1 traffic going for dogfooders.
Working on docs.

Never agreed with FMA, waste of time.

Would prefer Simon/Richard to be more straightforward/honest? Would help
career-wise to know where exactly I need to focus.

Need perf support.


## Sebastian.md on 2020-05-29

COVID19 app review, 24hr turnaround, not realistic. vteam first meeting monday.

Patrick Muchler pmutchler@ - runs DeepDive team, has opinion on the topic of
unprotected broadcast receivers and others.

PV in a11y/notifications - only apps which we don't show dialogue should only go
people who actually use a11y support.


## Sai.md on 2020-05-28

What preloads are exposing user data without proper permissions?

go/intercept-design

[Project Amsterdam](https://docs.google.com/document/d/1vcz_qAOYZnfLbUm9SsUT5j8H8STRiuLn72GWTVcKw9w/edit#heading=h.71bwjxl5wwu)


## Bessie.md on 2020-05-28

Wants to take bunch of Mondays off


## Dave.md on 2020-05-28

Backlink scorer -- now used for AdvProt. Can we use it for risk verification?
Spoke to Krish about it and shared with Khawaja.

Are you speaking to Pixel people? Bluechapel. Interested in Angel.

Team temperature. ASAP managers roundtable.

API Abuse v-team.

No United and Avis. Need to use Alexa top? Something KG related?

Ramanan
PJ reports to Ramanan

Directory, we works with that team a lot. Reached out 3-4 months ago, Pixel
wants to focus on digital car key and other keys. Pixel first feature. He said
we might be able to fund it. 

In a middle of trying to figure out what is bluechapel for us. What ASAP does?
What we recommend Pixel team to do?

Bluechapel security PM v-team - Stefan/Brahim. Brahim was appointed by Sabrina
as Pixel PM person looking after safety. Suzanne/Sudhi involved. 

Perf angst. Spoke to Renee and her manager. Will keep pushing.


## Alan.md on 2020-05-28

go/meerkat-s-sheet

https://docs.google.com/document/d/1vcz_qAOYZnfLbUm9SsUT5j8H8STRiuLn72GWTVcKw9w/edit#heading=h.71bwjxl5ww
-- project Amsterdam

SD card work - can turn it on in Master, but Narayan asked to delay.


## Simon.md on 2020-05-28

Simon to speak to kshams and people ops re Richard.


## Ricky.md on 2020-05-27

Want to go back to HKG but afraid for family. 
Previous geoflex came with some restriction

## Bernardo.md on 2020-05-27

Bought a cheap desk from IKEA, now need to return it. Internet is bad.


## Zach_Koch.md on 2020-05-27

Compare Play Protect and Google settings. Hard to reconcile.

What are 4-5 things we need do?


alingli@ - eng director in Taipei.

Bluechapel - name for device we are shipping in 2021 on Whitechapel. 

A lot of Android resources will be focused on Pixel. Settings team is ~10
engineers.

1. What is the exact thing we want to pitch? Concentrated thing that is def
   buildable for Pixel 2021? What is the pitch?

2. Zach goes to Ling; she's the right person to talk to to secure engineering resources.

Let's work with Zach on this.

## Krish.md on 2020-05-26

- Advanced Protection - why we should continue investing

There are things that we should do because they resonate well in press/media.
This is about furthering security narrative.

Make sure that impact is very clear.

What happens if we cut Aleksandr?


## Krish.md on 2020-05-26

- Conversations about Angel

Something special for Pixel. Funding will raise some eyebrows. Explicit guidance
was fund more SDK, elections. If Pixel makes Trust a pillar, Krish wants them to
fund this.

Pixel has heads!

FMD: we should continue to fund FMD. We should keep roadmap. Don't say yes just
because there's a new request. Auto had hundred heads they gave away last year.
Need to make sure Sameer and team are aware that we are funding FMD. Don't want
this to be a side show. 

Re-prioritisation is done. We are getting some heads. 
Spending cycles on project Anning - about ads.


# Week 21
## Henry.md on 2020-05-21

FMA unlocked.

Job signed up to do - yes
Nothing surprises me as per what I'm doing
Doing more software eng mgmt than done in the past


## Simon.md on 2020-05-21

Vision is not a root cause problem.
Not agreeing about long term vision for the project can be perfectly dealt with.

Simon: root cause is relationship based, Richard is a little too fragile, more
fragile than he appears.

Marie is quiet and shy. Very open that she is finding WFH situation very
challenging. General anxiety about the state of the world. Sundar announcement
didn't land well with her. She's coping fine, but anxious. 


## Alan.md on 2020-05-21

Trustworthy UI

BAL - is this being circumvented? unclear if we would get any yet.

Apps can close/open the shade. This is being abused!

Reasons to use: when someone clicks on notification, shade closes. Many apps
start receiver and launch activity, but shade won't close.

Used for locking the device because you can't click notification to close app,
seen it in the field, stops using from accessing device.


## Guanhuan.md on 2020-05-19

Anurag and Sorin are busy on GmsCore launch.

Thursday meeting: the message I delivered is questions of why are we doing
projects such as Timon, I said this was expected.

Beginning of the quarter - aligning our OKRs.

UWB team wanted to do a project with FMD. Find My Friends. Similar to finding
the device, but for friends. Are they ready to fund it?

GPP launched email campaign for AP.


## Renee.md on 2020-05-19

Need to make sure people know about reimbursement program.

Work with Scott on empowering TLMs.

What change will we bring based on our findings? Don't want to give people space
to vent if we don't do anything about it. Let's have few examples of things we
can change.

Manager roundtable.

Manager to reach out to HR to tell them they are concerned about someone.

Well-being checkin.

Maybe one-off managers together.


## Khawaja.md on 2020-05-18

There are many things we are doing about not listed at top level.
One way we are extracting certain things.
If put 20 projects in front of Sameer he's gonna cut all of them.

How can we tie FMD? Android perception probably.

Maybe we should position FMD as part of GPP?

Sameer got pissed off about Whistlepig, 5 VPs talking about project 3 people
working on.

How do we bubble up perception goals?

Goals:
- Platform
- Running current GPP
- Angel
- FMD
- AP4A

End user security. - maybe there's a better way of representing this?

Safety UX.

Theory: he was hearing things from Guanhuan who heard from James.

Let's look at the facts on the ground:
- No changes to FMD.
- Many projects were cut. FMD wasn't at all.
- This says that FMD is important for the company/organisation.

How to build top level goal: for next year OKRs, we need to come up with a story
on how we want to position these things. Abstraction is important in the context
we are.

ASAP OKRs are not specific enough for what we are doing, we need something that
is more focused on what our team is doing

(XX% of users are using safety products??)


## Simon.md on 2020-05-18

Richard: very upset. Isolation is getting to him. Rubber duck comment wasn't good.
Also worried about outcome of what happened?

Doesn't understand the point of the exercise if not to cut people.

Need to make it more clear where FMD sits in terms of priorities.

Get us to kshams' OKRs for next quarter.



# Week 20
## Anwar.md on 2020-05-15

- Situation with Platform. Plans for Android S - still figuring out, working
  with anti-malware team. Trying not to be a substitute of Framework team.
- Situation with Angel. Lack of HC.
- FMD and where do we go from here.

A lot of HC has just evaporated.

Anwar lost 14 heads for Mainline.

A lot of work for Pixel: for Whitechapel, how to make it performant. Foldable to
ship next year, some work was done with Samsung but not there yet. 

Pixel doesn't have much spare bandwidth. Investment into UWB is a good
investment, long-term thing, other OEMs will support it too.

Soli is basically dead. Hard time to deliver on promise. Too power hungry. Can't
afford to have both UWB and Soli.

Angel: what should Pixel Security experience be? We are focusing on Pixel, let's
say we do something on Pixel first -- how do we do this on Pixel? Blue Ocean OS:
the experience we _want_ to deliver on Pixel, could be blueprint for future
Android experience.

Good person to talk - PMs on Pixel side, someone on Sabrina's team. 


## Alan.md on 2020-05-14

Ed started email thread about Samsung - platform signing all apps, granting them
ridiculous permissions, browser can do pretty much everything.

Amsterdam - project to clamp down on this. No op for our team.

Virtual town hall on Monday about S.

We are on go/android-coreos-top3.

Our code for reading app ops is using reflection. Lots of changes to it, some
bits we want to access are available now so we don't need to use reflection, and
app compat stuff says we aren't allowed to run this.

Immediate fix - Moltmann allowed us to read this APIs. Not very helpful.

Ed started a hotlist for app visibility gaps. Pat is looking at package manager
things, we are looking at the rest of them.

R is winding down. SD card did not ship, but there's still a chance of getting
it into QPR. Need to turn it on in Master.

Who's going to do what on S? Bernardo is the biggest problem - he's good and he
wants something big and meaty. Takes interest in 2 big problems:

* Collapsing the status bar - we thought it's trivial, but lots of apps do this,
  because API is badly designed. Figuring how we allow legitimate use cases is
  hard. Might be not the most pressing problem in the ecosystem.
* BAL: we want to do an audit of window manager, find all the dumb things that
  allow you to steal touches etc.

A11Y: hook something like verifier into a11y flow. Convincing framework is hard.

kshams meeting was useful.



## kshams-leads.md on 2020-05-14

Key themes: 


## Simon.md on 2020-05-14

Richard - probably about prioritisation.

Nearby also completed their code.

Anurag was/is ill again, this is still ongoing.
Anita is feeling a little down?

Aleksandrs: 

* Bessie had a full blown breakdown couple of days ago. Feeling extremely
  stressed about the project she's working on. Need a lot more support. Will
  give a project where she's designing things. Package Verifier was done very
  well. Logging will be 10-20 LoC but design is hard.
* Peter is doing much better. Very positive and enthusiastic.
* Karishma still very worried about family, they are in complete and utter
  lockdown, worst part of country.
* Niko is better, jumping into his TL role.
* Raissa is comfortably meeting expectations. PEP will be formally closed next week.
* Marie is still feeling a bit slow, but doing well overall. Struggling with
  lockdown situation.



## Dave.md on 2020-05-14

Talking points:
- Advanced Protection:
  * Unlikely to hit EOQ2 - things are complex, trying to work out the dynamic of
    the manual review process + changes in Phonesky were far more complex than
    we anticipated when originally designing it.
  * App updates are difficult
  * What will happen beyond MVP? Carter?

Richard overreacted at perf too!

We are going to get small subset of T&S HC. API abuse is not part of that HC.
Dave will speak to Sameer about it.

Rotation of volunteers to review apps.


## David_Coffin.md on 2020-05-11

Looking at client side for AP4A.
Is there anything I can do to help?
If anything pops up, will ask.

Discrepancy - 70K 30DA for AP4A, and AP folks report just a shade of that for
iOS.

David to sync with Din.

David to buy and expense a headset.

Android Health team are going to start working on a project to look at how to
ingregrate dexcom data. 10% is OK 20% would be a struggle.


## Sudhi.md on 2020-05-11

Probably WFH till end of year. Should speak to someone to P&E Hiroshi chief of
staff.

Guidance/principles doc

Longer doc list of objections why on-device abuse detection is generally a bad
idea. Process of revising. There are words that need to be reworded, some words
are offensive.



# Week 19
## RichardD.md on 2020-05-06

- Design reviews. Communicating design decisions internally/upwards/sideways
- AASE Eng Review on Find My Device maybe?

Pixel team are going all-in on ultra-wide band and they need an app that's going
to support it, so they need FMD to being able to onboard it.

Jim and Anita are struggling, Jim more so. Anita - pair programming.


Need to store per-device consent. Storing it in FMD server. Used to be able to store this under location history ToS. Whenever connect event, FP need to look if there are relevant ToS for FMD before sending event to FMD. Every time FP contacts us, so we need to change that it doesn't reveal that there was connect or disconnect until ToS is signed. Anurag implemented all of it, badly, before design/review/privacy/security, had to re-do it over the weekend. Learning experience. 

We are going to notify users 3 times that they need to sign ToS, this might be seen as spamming by some people. Pauline prepared to let it go.

The number of notifications we send is the flag on the server. Can change campaign in useful ways.

Trying desperately to get more progress on Timon, trying to focus on Timon. Executing on that project doesn't seem to be there. Lack of motivation to go ahead with Timon.


## Khawaja.md on 2020-05-06

Prioritization exercise.

Simplify bucketing projects to make it easier to present to execs. Connect the
dots between the projects and bigger objectives.

What are the sacred cows?

SDK detection - probably no impact in 2020, but the work must continue
happening.

How are we deploying the resources? How many resources, what purpose, where they
are?

Stefan - move off uFBI.


## Niko.md on 2020-05-05

Droidfood devices - need development devices.


## David_Coffin.md on 2020-05-04

Prototyping dialogues. Spoke to Jan and Alan, will be doing weekly syn now to
stay more in touch with what Jan is doing.

Alan suggested that they write design doc together, Jan thinks he should write
it separately.


# Week 18
## Sebastian.md on 2020-05-01

Discussion kshams/krish: more wood behind fewer arrows. What will get cut? 
Sameer wants to stop investment into advertising fraud, doing too well?
Half way to the annual goal in Q1, should we re-focus?

For election - exploring AP dataset. 


## Peter.md on 2020-04-28
WFH is going OK, better than at work. Feels more calm/relaxes.

Q: How does this reflect on the team?
Peter: don't know how people are feeling but work is going as well or better than before. How they are personally feeling, hard to say.

WFH thing is nicer than expected
More freedom during the day to do other stuff - e.g. go for a walk for an hour or two. North London. Finsbury park. Flexibility.

Bit harder to develop something remotely - slower internet etc. 

GGeist: last topic - stress/well being. 
Integrity: was one of the people who said I don't trust Google with this now.
Not particularly about the team, just more about all these scandals. 


## David_Coffin.md on 2020-04-28

What are the dangerous permissions in question? Want to chat with reverse engineer. A lot of READ_ permissions are dangerous as far as targeted users concerned. A11Y. Camera/Record Audio. 

GGeist - didn't work particularly well, first three points no one really said
anything. Useful to hear my thoughts on this, but no interaction. Question about
integrity -- I would've liked to know what are the broad reasons people are so
negative. Not obvious who wants to speak. 


# Week 17
## Khawaja.md on 2020-04-24

Anything you do to check on the team is not a bad idea.


## Sebastian.md on 2020-04-24

Lack of comms for people reporting to kshams. Reported to Dave.

Will another newsletter save us?

Things that started/happened this week



# Week 14
## Aleksandr.md on 2020-04-05

Auto-disable feature - Raissa will send out a status of what's going with the
feature.

> Autodisable flags: Controller migration: in Googlefood, waiting for it to soak
> for a few more days before releasing to 1%. Autodisable feature [dependent on
> controller migration]: ready for Phonesky QA Enable autodisable feature to all
> devices [dependent on autodisable feature]: releasing to team food

Click log migration was on hold -- was migrating everything to Kotlin. Doing
that part now. Logger would be its own module, and Phonesky is encouraging
everyone with new modules use Kotlin.

Karishma - finishing interview feedback. Working on Oleg PEET tests. Submitted
CL that fixes colour on Oleg shield. Still working on it, figuring out how to
add divider.

Bessie - changing date format changes, i18l, Indonesian/Chinese/Finnish. Design
doc on entry point logging.

Marie - wrapped up passport and design doc on Friday. Looking into how to
measure our metrics.

Rodrigo - preparing for Play UX review, but it got cancelled.

Peter -- Installer card QA approval is very close!

Privacy annotations - project Delphi. CL is out to define the annotations.

## Ed.md on 2020-04-05

Portion of our work we want to outsource to Chad - safer by default APIs.
He's keen, trying to check with him periodically that we are on the same page.

People were coming to him and asking how he helps API abuse team. They work on
making harder to abuse other apps, so API abuse is a second-order effort.

Android S shortlist within a month.

week of 18/05 - hackathon, 2 weeks away.

Love for Alan to carry on being the TLM.

Become a little bit jaded with disparaging comments about Platform Security,
many terrible decisions made in the past. Alan is also guilty at times of the
same thing.

Let's imagine new eng. Not healthy to think that way of Framework / PlatSec.


## Permission_focused_AP_filter.md on 2020-04-05

**David**: we discussed permissions in filter last week. We want to see whether we can
investigate this. Couple of reasons:

* OKR is framed around reduction in number of targeted installs. It's hard to
  create spyware without requesting sensitive permissions, so common sense is
  that we should treat app differently if they request 5-10 sensitive
  permissions.
* We can simplify the messaging to the users/developers - app accesses sensitive
  permissions, hence extra scrutiny. Right now we say that every app is reviewed
  to very high level, but some apps are reviewed more than others ... tricky to
  communicate to the users and developers alike.

For the purposes of comms we can consider many permissions sensitive. 

**Din**: agree that comms is important, most users don't even know about all of
that. Maybe we should call out what this app is requesting which may have them
change the decision.

What are the common permissions that spyware are using?

The list (in the doc) covers almost all of them. 1/3 of the spyware uses certain
permissions, `ACCESS_FINE_LOCATION`, bluetooth admin etc.

RK: what does this permission do? 
DC: not sure why spyware would want to use those.


**Din**: [spreadsheet from Din](https://docs.google.com/spreadsheets/d/1EiGM88Nv4fluEqEVs659EEHqAqVjIB0QwoqLNfAvbNo/edit#gid=2054355730)

76K spyware packages, near 108K versions.

```
app_permissions	cnt_app_versions	% of spyware (app versions) with permission
INTERNET	38,685	36%
ACCESS_NETWORK_STATE	38,203	35%
WRITE_EXTERNAL_STORAGE	36,113	33%
WAKE_LOCK	35,850	33%
ACCESS_FINE_LOCATION	33,827	31%
ACCESS_WIFI_STATE	33,632	31%
READ_PHONE_STATE	33,444	31%
ACCESS_COARSE_LOCATION	33,304	31%
RECEIVE	32,628	30%
RECEIVE_BOOT_COMPLETED	31,262	29%
READ_EXTERNAL_STORAGE	31,258	29%
GET_ACCOUNTS	30,898	29%
BLUETOOTH	27,909	26%
WRITE_SETTINGS	26,868	25%
READ_SETTINGS	26,170	24%
```

DC: how many permissions should we consider too high/suspicious? 

```
package_name	cnt_ap_device	cnt_sensitive_permissions
com.ktcs.whowho	16	31
com.verizon.messaging.vzmsgs	786	27
com.skt.prod.dialer	24	27
com.lbe.parallel.intl.arm64	9	26
com.skt.tmap.ku	8	26
com.sec.android.automotive.drivelink	206	24
com.huawei.health	178	24
com.whatsapp	93	24
me.dingtone.app.im	48	24
com.skype.raider	41	24
```

RK: let's focus on most sensitive permissions.

DL: let's get list of permissions from TAG.

RK: I will go through TAG reports again.

DC: if app has at least one permission - this might be more complex, because
e.g. great many apps are requesting location.

DL: didn't have a chance to run current filter on the general population, too
big autoscan data, will try random subset.

We shouldn't discount any devices where we auto-disabled apps, because they were
not clean, we _made_ them clean.

DC: we are interested in whether you were ever infected.

RK: agreed, if device was ever infected, data exfil has already happened.

DC: for the next steps:

* Up to Din to decide how to proceed with this. I think it might be useful if we
  finalise the list of permissions we consider dangerous. Two categories of
  permissions:
    - Sensitive but not indicative of targeted attack app (e.g. location)
    - Really sensitive (e.g. call log)
* Interesting to see how current filter performs on the general population. 
* We still should overlay permissions on top of that to make comms easier.

DL: do we have a good way determine VIP apps?

RK: no, but think partners or 50M+ installs.

DC: would be useful to see the list of apps which would be excluded based on
these conditions.

DL: if the app has certain number of versions, but only some are asking for
permissions?

RK: look at currently published version.

# Week 1
## Jason.md on 2020-01-05

QBR with Sameer this year. (add deck)

Mostly looked at landings / risks. Have to ruthlessly prioritize stuff. 

What happens with resource spreadsheet? Sent to Dave/Suzanne, will speak at
leads next week.

Ads data transparency/control project -- Anning


## Max.md on 0001-01-01
---
tags:
  - 1:1s
  - angel
---


## Guy.md on 0001-01-01
---
tags:
  - 1:1s
  - bosses
---


## Yuri.md on 0001-01-01
---
tags:
  - 1:1s
  - angel
---


## Guemmy.md on 0001-01-01


## Heger.md on 0001-01-01
---
tags:
  - research
---



## Henry.md on 0001-01-01
---
tags:
    - 1:1s
---


## Inara.md on 0001-01-01
---
tags:
  - 1:1s
  - team
  - aleksandr
---


## Isabelle.md on 0001-01-01
---
tags:
  - 1:1s
  - fmd
---


## Zach_Koch.md on 0001-01-01
---
tags:
    - 1:1s
---


## Ishtiaq.md on 0001-01-01


## bdc@.md on 0001-01-01
----
title: bdc@ 
tags:
 - bosses
 - android
---


## Guanhuan.md on 0001-01-01
---
tags:
    - 1:1s
    - pms
---


## Giulio.md on 0001-01-01
---
tags:
  - 1:1s
  - angel
---


## HMD meeting.md on 0001-01-01
---
title: HMD meeting
tags: 
  - partner
  - oem
---


## James.md on 0001-01-01
---
tags:
  - 1:1s
  - pms
---


## Mainline SDK.md on 0001-01-01
Need to get PM agreement
Mainline SDK is not fully launched
Two use cases:
- new API for R, photo picker
- Rubidium is another use cases, want to add _all_ the APIs after T ships

Limited clients makes it easier
We are not fully integrated with Android Studio, not fully integrated in general 3P

Adding a separate versioning mechanism

In theory this is doable. What are the timelines? Who are end users? 

Do not update google3 with Mainline SDKs yet. Need to figure when we are confident enough to put Mainline SDK to g3. Potential blocker, not the biggest but needs to be done. Jan train is a reasonable time to aim for. In practice Rubidium is P0 for Mainline and Android in general. 

Let's file a Mainline FR, so we can go through design review, engprod, maybe partner eng. 

go/android-fr has a check box if you are affecting Mainline, there will be questinairre




## Running_notes.md on 0001-01-01
---
tags:
  - meeting
---


## Jan.md on 0001-01-01
---
tags:
    - 1:1s
    - angel
---


What is the recommended option? Option (2)

Phonesky user settings. RPC call to update user settings via DFE. Jan to ping
Andy -- update the document and send the doc to Andy.

Work with David - no regular meetings, pinging him regularly with questions,
e.g. his past experience with launching, communicating. Reviewed comments on this
doc.

Alan - involved in AP in a sense that touching on it in 1:1s, chat. Talking to
him about it. Talked to him about option (2). 


## Jason-Cornwell.md on 0001-01-01


## Jason.md on 0001-01-01
---
tags:
    - archive
---


## Jeroen.md on 0001-01-01
---
tags:
  - 1:1s
---


## Giles.md on 0001-01-01
---
tags:
    - 1:1s
    - dk-leads
---


Mainly working on contact tracing -- 60%
It's in EAP. UK - unclear if this is a hard decision.
Mustafa (ex-Deep Mind) is taking to NHS.

Delfi is also taking a lot of time.

Data budgeting -- measure how many bytes each module uploads. Based on current
usage, the goal is not to increase it, work with teams to make collection more
efficient. 

E2E encryption for FMD -- need to have different flow for devices you own.


## Gaurav.md on 0001-01-01


## Gabriel.md on 0001-01-01
---
tags:
  - 1:1s
  - aleksandr
---


## Eugenio.md on 0001-01-01
---
tags:
- 1:1s
---

Private compute backend, dependent on Oak

Google Play Protect Services - name was not approved by AiAi group, as we wouldn't approved this naming. 

What an AV app is and what are implications. Chad/PWG are concerned that in some countries (e.g. Russia) if you have an AV, you have to pre-install Russian version of it. We defined GPP service in a narrow way.

Consent is problematic - it can be forced, unclear, etc. 

What if Samsung decides they want to build their own product? (of ODAD)

What if ODAD team is forced by gov?

What if every time verdict is passed there's a notification

Network Ledge - each time AiAi access network for anything, it shows in the ledger. 


## Erik_W.md on 0001-01-01
---
tags:
  - 1:1s
  - angel
---


## Elliot.md on 0001-01-01
---
tags:
  - 1:1s
  - team
  - angel
---


## Jessica_L.md on 0001-01-01
---
tags:
  - 1:1s
---


## William.md on 0001-01-01
---
tags:
  - 1:1s
  - team
  - odad
  - direct
---


## Jessica_Lunney.md on 0001-01-01
---
tags:
  - 1:1s
  - non-eng
  - legal
---


## Jim.md on 0001-01-01
---
tags:
    - 1:1s
    - fmd
---


Remote development is kinda hard.

FMA - no way to fake an accessory, it's always real accessory / real FMA
backend. Just the way it was built.

How the team doing? Doing OK. Early on people were pretty productive, couple
weeks ago a bit of a slump, team cohesion not as good. Only business, not a lot
of dynamics. Richard recognised it, talked to the team about it, getting a bit
better. One goal of this sprint is to do more collab work.

GVC interrupting - didn't notice this being a problem.

Richard - mostly positive things. Really cares about the team. Both cares about
project but also team members individually. Recognises concerns. Bit of a slip
recently when in general across the team - a bit of frustration that Timon was
not making so much progress as we liked. Tried to shift mid sprint to get more
people to Timon.

Reasons for delays:

- Lockdown, everyone working from home
- FMA taking a lot more cycles than anyone expected
- Should've scoped first milestone to be smaller. A lot of upfront cost in using
  Boq. Fewer reviews - no logs reviews etc, but Boq forces you to make this
  before the start of the project rather than before the launch. A lot of them
  mechanical, but there's a lot of them. M1 - not a lot of code, but a lot of
  set up.

The thing we need is the sense of momentum. More detailed planning might help
but is this a biggest priority? Where do we need to be?

Challenges of rewrite - didn't deliver the benefit of we are done (turn off the
old one). How can we turn it down sooner?


## Ed.md on 0001-01-01
---
tags:
    - 1:1s
---


## Dianne.md on 0001-01-01
---
tags:
 - 1:1s
 - bosses
---


## Wa.md on 0001-01-01
---
tags:
  - 1:1s
  - team
  - odad
---


## Joanna_Jakobska.md on 0001-01-01
---
tags:
  - 1:1s
---


## Karishma.md on 0001-01-01
---
tags:
    - 1:1s
    - team
    - angel
---


## Tyler.md on 0001-01-01
---
tags:
  - 1:1s
  - angel
---


## Sudhi.md on 0001-01-01
---
tags:
    - 1:1s
---


## Steve_Kafka.md on 0001-01-01
---
tags:
  - 1:1s
  - pms
---


## Stefan.md on 0001-01-01
---
tags:
  - 1:1s
---


## Khawaja.md on 0001-01-01
---
tags: 
  - 1:1s
  - bosses
  - dk-leads
---


## Sorin.md on 0001-01-01
---
tags:
    - 1:1s
    - team
    - fmd
---


Working on Timon. Orchestration service, which forwards request to
device-specific service, working on Android version. Working on ring request.
Handling ring from web or app. Managed to ring device from the new backend, can
run infra locally.

For full roll-out -- redirect traffic from GFE directly, but for that we need an
adapter which adapts from old to new format. Current web service is not
OnePlatform compatible, so we are using GFE. 


## Krish.md on 0001-01-01
---
tags:
    - 1:1s
    - bosses
---


---
tags:
    - 1:1s
---


## David_Coffin.md on 0001-01-01
---
tags:
   - 1:1s
   - team
   - angel
title: 'David:Roman'
---


## Leandro.md on 0001-01-01
---
tags:
  - 1:1s
  - odad
---


## Linus.md on 0001-01-01
---
tags:
  - 1:1s
---


## Simon.md on 0001-01-01
---
tags:
    - 1:1s
---


- Lots of interruptions in FMD GVC meetings
- Situation with FMA changes done by Anurag
- Lack of interest in working on Timon

Trying to help James to help with the lock screen. 

talking over each other - Richard is very hard to talk to in meetings, talks
over you, doesn't let you speak, aggressive, confrontational. However, perf/MFS
feedback says exactly opposite. 

Designs were not reviewed communicated.

Aleksandr:
* First signs of some people who might not be coping too well. Niko and Karishma
  are struggling right now. Karishma because family. Other things happening in
  her family apparently. She's really worried, really wants to be there but
  can't.
* Niko -- just a little bit down about the whole thing / fed up with lockdown.

Things to keep an eye on:

  * FMD -- join few dailies, see how the things are going
  * Check in with Marie.


## Shriya.md on 0001-01-01


## Lisa_Martinez.md on 0001-01-01
---
tags:
    - 1:1s
---


## Shan.md on 0001-01-01


## Sham.md on 0001-01-01
---
tags:
  - 1:1s
  - pms
---


## Dave.md on 0001-01-01
---
tags:
    - 1:1s
    - bosses
---



- HC was pulled :(
- What's the latest with platform security?
- Thoughts about Advanced Protection

HC: Dave has no idea. Sameer doesn't know either. The only thing we can do is
for everything that is super-critical, keep that list up to date, see what we
can get approved. There's no even process.

Someone in PlatSec wrote a paper on overlay attacks.

PlatSec on-device: Dave met with Dianne, came with rough idea of what to do,
build AiAi kinda of component. PlatSec tried to kill it again. Been asking Sudhi
to get his team aligned. Had another meeting with Dianne yesterday, aligned on
AiAi model. Vish is on board.

kshams got pentesters Play Store. Khawaja has way of doing it. 


## Chad.md on 0001-01-01
---
tags:
  - 1:1s
---


## Brandon_Barbello.md on 0001-01-01
---
tags:
  - 1:1s
---


## Liza_Ma.md on 0001-01-01
---
tags:
    - 1:1s
---


Liza works on SuW, but used to work on Settings. 

No longer work on settings because team is in Taiwan. Good size team there.
zkoch@

People to talk to:

Position this not as OEM rearchitecture, but we are re-architecting information
architecture we have for settings.

Wait until they send next year comm on whitechapel / bluechapel?

https://moma.corp.google.com/person/gundrum

## Bessie.md on 0001-01-01
---
tags:
    - 1:1s
    - team
    - aleksandr
---


## Luke.md on 0001-01-01
---
tags:
  - 1:1s
  - aleksandr
---


## Serban.md on 0001-01-01
---
tags:
    - 1:1s
---


## Bernardo.md on 0001-01-01
---
tags:
    - 1:1s
    - dotted
---



## Barak.md on 0001-01-01
---
tags:
  - 1:1s
  - pms
---


## Marie.md on 0001-01-01
---
tags:
    - 1:1s
---


## Mike_Tsao.md on 0001-01-01
---
tags:
  - 1:1s
  - fmd
---


## Anwar.md on 0001-01-01
---
tags:
  - 1:1s
  - bosses
---

Next 1:1 - speak about vulnerabilities

https://mail.google.com/mail/u/0/#inbox/FMfcgxwHNVzMTxWrbnPgMxTltRMfbnGc


## ASAP_manager_roundtable.md on 0001-01-01
---
tags:
    - meeting
---


## Monir.md on 0001-01-01
---
tags:
    - 1:1s
    - peers
---


---
tags:
    - 1:1s
---

Son - Mubdi, 3rd of May 2021


## ASL_weekly.md on 0001-01-01
---
tags:
    - meeting
---


## Sebastian.md on 0001-01-01
---
title: 'Sebastian:Roman'
created: '2020-04-24T16:15:43.467Z'
modified: '2020-04-24T16:20:30.515Z'
tags:
 - 1:1s
 - peers
---


## Meerkat_Leads.md on 0001-01-01
---
tags:
  - meeting
---


## Sai.md on 0001-01-01
---
tags:
    - archive
---


## Saagar.md on 0001-01-01



## Meerkat_Weekly.md on 0001-01-01
---
tags:
  - meeting
---


## Ronald.md on 0001-01-01
---
tags:
  - 1:1s
---


## Anurag.md on 0001-01-01
---
tags:
  - team
  - 1:1s
  - fmd
  - direct
---


FMA part is very exciting, also Timon.

In-Maestro FMD flow.

Some tensions in the team? Probably not, but some interruptions on GVCs.

10 min after stand-up just to chit-chat

Myself well. No concerns so far. Some visual noise/tinnitus, but not worse. 


## P&E_leads.md on 0001-01-01


Slowing down hiring significantly. Large number of Googlers who had to be
onboarded remotely. Hiroshi has to approve all the reports. Not a good
experience for candidates.

Sena: there will be yet another prioritization exercise.

Getting back to work: gradual return to work when conditions allow. No
one-size-fits-all approach, each region. All of us are unlikely to go to the
office at the same time, who has absolute need to be in the office and who can
continue to wfh for extended period. Some teams has higher need - IRT, SREs,
maybe Android ...

Matt Brittin leads EMEA IRT. Once EMEA IRT figures out how guidance applies to
our region this will be cascaded to local IRTs. 

Worked with BT and Samsung to provide tablets for every single intensive care
unit worker in the UK. Major challenge - once in the ICU, no way to comm with
the family. Solution was to provide tablets that were configured to use single
app approved by NHS. Google paid for 4.5K tablets to be distributed to every
single ICU in the UK. Configured with AE, MI is EMM, locked down to a single
app. BT provided connectivity. 

Clara: Project Fire: WHO came to Google asking for standardised way to upload the data
from institutions across the world, like Apple Health but not Apple :) 

Riccardo: Project Apollo, BT/Nearby. 

12 apprentices that were meant to be finishing in October, we extended this
period for a year. 

go/android-wfh-bug


## Muhammad.md on 0001-01-01
---
tags:
  - 1:1s
  - angel
  - direct
---



## Nandini.md on 0001-01-01
---
tags:
  - 1:1s
---


## Anu.md on 0001-01-01
---
tags:
  - 1:1s
---


## Rodrigo.md on 0001-01-01
---
tags:
    - 1:1s
    - team
---


Reverted back to the original UX on AP.
Jenny looking at the strings, reviewing with their UX lead. Julia from Play is
involved.

Working remotely absolutely fine. Got used to WFH. Productive. More structure. 

Maybe too much meetings because some things can be said in the corridor.

Some people doing better than others. Haven't seen nothing specifically
negative, but seen some people who may have some personal things going on that
might be affecting them.


## kshams-leads.md on 0001-01-01
---
tags:
    - meeting
---


(deck from Tanner Verhey)

Was review with Sameer yesterday. Assume 0 heads increase, can keep attrition.

1. Yet another "let's assume no HC" exercise.
2. Prioritize for highest-impact effort. Optimise for long-term impact.

Come back with half of OKRs. Pause things which are already green and
concentrate on the red ones.

Jenny: extending her leave. Next planned date is 01/07/2020. 


## Ricky.md on 0001-01-01
---
tags:
    - 1:1s
---


## Narayan.md on 0001-01-01
---
tags:
  - 1:1s
  - bosses
---


## RichardN.md on 0001-01-01
---
tags:
    - 1:1s
    - asap-lon
---


RichardN to set up conversation with Alan and Seb.


## Anita.md on 0001-01-01
---
tags:
    - 1:1s
    - fmd
---


## Angel_Cheng.md on 0001-01-01
---
tags:
  - 1:1s
---


## 2021-07-13_08-10_Skip_Levels.md.md on 0001-01-01
---
tags:
    - 1:1s
---


## Angel.md on 0001-01-01
---
title: Angel Cheng
tags:
  - 1:1s
  - peers
---


## Andrei.md on 0001-01-01
---
tags:
  - 1:1s
---


## RichardD.md on 0001-01-01
---
tags:
 - 1:1s
---


- How's Jim doing?


## Coaching.md on 0001-01-01
---
tags:
  - meeting
---


## Aleksandr.md on 0001-01-01
---
tags:
    - meeting
---


## Naveen.md on 0001-01-01
---
tags:
  - 1:1s
  - pms
---



## Richard Gaywoord.md on 0001-01-01
---
tags:
	- 1:1s
	- privacy
---


## Renee.md on 0001-01-01
---
tags:
    - 1:1s
---


## Nicole.md on 0001-01-01
---
tags:
  - 1:1s
---


## George Roussos.md on 0001-01-01
We have a group of people who deal with this. G will put me in touch with the right people. The way it works there's a gift which is given from Googel to BBK with the purpose of using money to support the scholarship. Standard process of doing that.

Timing - can be done in installments, or a one off things.

Normally give money to the college, this counts as a gift for corporate acct, tax impl. College takes care of everything.

Affiliation - Google / ASAP group.

Timeline - ideally will start advertising in October. 

## Teamfooding instructions.md on 0001-01-01
No flashstation on Mac. 

```
[kirillov@kirillov-macmini:Downloads/angel-teamfood-2022-05-11-experimental]$ cp ~/Downloads/cl_flashstation ~/bin
[kirillov@kirillov-macmini:Downloads/angel-teamfood-2022-05-11-experimental]$ cl_flas
[kirillov@kirillov-macmini:Downloads/angel-teamfood-2022-05-11-experimental]$ export PATH=$PATH:$HOME/bin
[kirillov@kirillov-macmini:Downloads/angel-teamfood-2022-05-11-experimental]$ cl_flashstation
zsh: permission denied: cl_flashstation
[kirillov@kirillov-macmini:Downloads/angel-teamfood-2022-05-11-experimental]$ chmod +x ~/bin/cl_flashstation
```


## safety_hub_linkedin.md on 0001-01-01

The cat is now out of the bag - at Google I/O we have [announced](https://blog.google/products/android/io22-android13beta2/) a unified Security and Privacy centre in Android

> Later this year, we’ll introduce a unified Security & Privacy settings page in Android 13 that brings all your device’s data privacy and security front and center. This will provide a clear, color-coded indicator of your safety status and offer guidance and steps you can take to boost your security.

This is a big deal for the ecosystem safety. Previously in Android 12 (or Android S, as ~~googlers~~ some people are still naming it), we have launched the [Security Hub](https://www.linkedin.com/posts/sgzmd_google-needs-to-bring-the-security-hub-to-activity-6859818722526289920-W4E2?utm_source=linkedin_share&utm_medium=member_desktop_web) which unified most security settings on Pixel devices. Security and Privacy centre (or Safety Centre as we call it internally) is yet another step in the direction of bringing a unified vision of safety to Android users. 

For the vast majority of users the line between security and privacy is blurred at best - and, to be honest, sometimes it's even hard to define it. As one example, app permissions can be considered both a security and a privacy mechanism. There is a lot (and I mean A LOT) articles on the internet where "experts" are showering users with their advise - not always factually correct, and it's easy for a regular user to get lost, and make wrong choices, ultimately hurting their safety. The goal of the Safety Centre is to answer two simple questions many:

1. Am I safe?
2. Is there anything I need to do to get safer?

We strongly believe that following the guidance in Safety Centre is the easiest way for most users to get to the safest state on their device and beyond. We went a long way since [Verify Apps was announced](https://www.androidauthority.com/android-4-2-verify-apps-security-feature-explained-by-google-131514/) in Android 4.2 - it's been 10 years after all! - but there's still more to go. Safety Centre is meant to be an iterative product, closely aligned to the problems and threats Android users are facing, so stay tuned for the full launch later this year as well as any follow-up updates!

## Anurag_Promo.md on 0001-01-01
---
tags:
  - research
---

Solutions to difficult problems easy to maintain - need to reflect this

Timon is difficult, is there proof?

[x] Positive impact on velocity

[x] Informal and formal mentorship

[ ] implementing new practices that address the team’s needs - where?




## App_dashboard_for_AP.md on 0001-01-01

We want to have a dashboard which shows:

1. List of apps in the final whitelist, explaining the reason why it got there
2. List of P0 apps which didn't make the list, ideally also explaining why
3. List of developers affected by (2), with cumulative number of installs.

We start with Din's table
We union it with apps from backlinks which either have a backlink or have exact
email match.

Result:
  * `package_name`
  * `developer_account_id`
  * `num_installs`
  * `is_partner`
  * `reason` (`CURATED_LIST`, `IMPORTANT_APPS`)

We then start with my script which produces P0 apps not in the final list. We
can greatly simplify it by simply finding all partner apps not in the table
above.

This table will have the following information:

  * `package_name`
  * `developer_account_id`
  * `num_installs`
  * `developer_tenure`

## Calibration_MY_2022.md on 0001-01-01
---
tags:
  - research
  - perf
---

Use gdocs: https://docs.google.com/document/d/1nj95nXE5XlfOagUN12BXES1C2aeRmBfdubzcrhMHEFQ/edit?resourcekey=0-JR3eHwBTEObnuVB_6iZyLw#x 


- Oversaw delivery of Permission Autoreset, intervening and unblocking team as
  necessary (although most work was delegated to pvisontay@ and other team
  members). #L6-Strategic-instead-of-tacticals

- Owned an open-ended problem of researching reliability metrics of GPP - while
  didn't quite #L6-Solves-large-open-ended-problems but on his way there.

- [Analysed](https://b/200037740#comment3) performance bottlenecks of previously
  built (by ngbravo@) cache for GPP security status and launched it (e.g.
  b/205532779, cl/411610276)

- Resolved difficult performance mangement case and successfully got report back
  to good standing

- Liaised with HackDroidz team about an informal rotation opportunity for
  interested team member, allowing them to gain useful experience.

- Came up with an idea of ASL-wide "fun committee" and helped kirillov@ to
  organise it #L6-Positive-impact-on-team

- Continued to ramp up as the TLM of the team #L5-May-be-TLM-or-TL

Rating: EE (or high end of CME?)

Why not higher: Niko shows some traits of L6, but they are mostly naiscent.

## Comp_Planning_2021.md on 0001-01-01
---
tags:
  - research
---


## Rebecca_Salois.md on 0001-01-01
---
tags:
  - 1:1s
  - research
---


## DCL.md on 0001-01-01
https://docs.google.com/document/d/1B9I3bqANFBWyH6pa6AoYtp2ArEBJAAJNc-S6PK0pieU/edit?resourcekey=0-bYYI5QlrZm3FjU8gdeQtzQ#heading=h.7olaunfw7vv3 - Ending DCL on Android
go/ending-dcl

Serban is ART PM

## Fit_Calls.md on 0001-01-01
---
tags:
  - research
  - fit-calls
---


## Niko.md on 0001-01-01
---
tags:
    - 1:1s
    - team
    - aleksandr
    - direct
---


## Raissa.md on 0001-01-01
---
tags:
    - 1:1s
---


Wants to take mondays off
Goes to Romania in July. Or not.

Isolation is really getting to some people.
Those who have people around them - it was much easier for them. We should
support these people to go to the office.


## Rahul.md on 0001-01-01
---
tags:
  - 1:1s
---


## Prabal.md on 0001-01-01
---
tags:
  - 1:1s
  - angel
---


## Pieter-Paul.md on 0001-01-01
---
tags:
  - 1:1s
---


## Alan.md on 0001-01-01
---
tags:
    - 1:1s
---


## Ahmed_Radwan.md on 0001-01-01
---
tags:
  - 1:1s
  - peers
---


## Forward_Looking_Promo.md on 0001-01-01

- Contribute to setting technical vision, strategy and roadmap(s) that affect one or more products that directly impact Google's business or overall strategy
- Directly impact success 1+ eng teams
- Identify/own complex, n-dimensional pieces of work x-product
- Understand interaction b/w execution strategy, org structure, talent, external forces
- Drive x-org collaboration
- Key resource, deep understanding of business/user needs
- Community contributions

- Development of product strategy and vision across partner eng, product, UX. Focus: future-focused capabilities
- Execution of 1+ large-scale business-critical software projects
- Anticipate changes
- Own customer experience
- Key decisions w/significat impact on function, drive high-level direction
- Set clear objectives, scope, approach and work plans

- Own complex, n-dimensional, ambiguous pieces of work, planning horizon 1yr+ x-product
- Support senior directors across sub-functions in PA
- Lead role in structuring and solving significant business/org challenges
- Define and address key challenges that have material impact on GOOG work and products

- Drive x-org collab, ensure strategy, direction and decisions of multiple work streams are aligned
- Build and maintain rel with xfn and senior client leadership (VPs, Directors)
- Manage multiple projects/work streams, teams, Mgrs and Sr Mgrs, facilitating talent mobility, growth, development
- Act as a trusted and neutral arbiter

- Recognised leader within a strat business area, deep knowledge and experience
- In-depth SWE knowledge
- "Raise the bar"

## GPP - Angel - ODAD.md on 0001-01-01
Worst case scenario:

We launch GPP migrated to Angel UX
Some OEMs take it
Some do not

As the result those OEMs have to use old GPP UX, and they can't get ODAD

PEP OEMs will _have_ to take Angel - this is a contractual obligation. Tied to revshare. Samsung is not a part of PEP.

So the worst case scenario is likely to be:


## GPP_doomsday_brainstorm.md on 0001-01-01
---
tags:
  - meeting
---


## Nikolay.md on 0001-01-01
---
tags:
  - 1:1s
  - odad
---


## Googlegeist discussion.md on 0001-01-01
---
tags: [management, team]
title: Googlegeist discussion
created: '2020-04-24T12:20:02.170Z'
modified: '2020-04-24T12:58:42.221Z'
---


[Googlegeist report](https://googlegeistreport.googleplex.com/googlegeist-results-jan-2020-7/kirillov/org-1?cat=-1&tab=questions&leaderstype=all)

[Responses](https://docs.google.com/forms/d/1KGEJdHQvbXOGYPRwrYviSwX6NdoLAFuFob6kRj09XA0/edit#responses)


Can't raise concerns without retaliation.
Happened after big scandals.

Andreis misconduct is 46/25/29 vs ours 25/25/50
nmilena 40/30/30

People ambivalent is answering, but now are interested in discussing this.

Maybe we should not come with a plan?

- Find company policies, point people to them
- What would you like us to do?

Not an issue within our workgroup.
How to address in mtg?

So here's the plan: re-iterate everything we've already said. 

There were issues, company ack it, let's explore why people voted red here more than other areas of the org? 


We are between Andrei and Milena. 

Promotion process is fair - people just don't understand it. So we dropped the ball on this, apologise and keep going.

Promo requires visibility - true or false? Not what Android safety required

Part of mgr responsibility.

Promo process is evidence based, reviewed by mgrs within organisation.

RichardD: don't believe it is possible to understand if someone did well or not.


Us - 31/15/54
Andrei - 44/18/38

Simon is doing worst, 72 negative. 


Better than last year but still not good.

Internal resources for dealing with work related stress?
Follow-up in 1:1s?

UK Emotional Wellbeing
Extended Carer leave
Headspace

Establish weekly survey? Might add more stress

This is one of the numbers which is in our power to fix

Share stories about times when I struggled with stress.



## Malware Strategy.md on 0001-01-01
Articles like [this](https://www.phonearena.com/news/millions-android-users-delete-apps-removed-from-google-play-store_id139477) should cease to exist.

Agressively revoke permissions from the apps which we don't feel are safe. Cleanup apps that were suspended on Play for whatever reason.

**but** be smart about detecting if the user is likely using root on purpose, so they will be pissed off if we start fiddling with their device

We need a good view of off-Play ecosystem. At the very least, we need to know which are the most popular off-Play apps, apps rising in popularity etc. 

We can and should crawl 3P app stores.

Tackle DCL -- at the very least, in Enhanced Security mode? 

JIT is a big part of it.

There was an idea of app sandbox.

[go/malware-honeybadger](http://goto.google.com/malware-honeybadger)

## Meerkat_Platform_Transition.md.md on 0001-01-01
---
tags:
  - research
---


Here's what I expect we are doing

  * Jan stays with Angel (keep HC)
  * Bernardo is moved to PlatSec (keep HC)
  * Linus and Ricky moved to PlatSec (move HC)

This way we free up 2 HC + SWE for Angel.

Question is whether Sudhi would go with it. 

What is it we need to find out:

1. Whether Jan stays or goes
2. Whether anyone else on the team would rather stay and do something else (Angel)
3. What's the final tally and what Sudhi can fund.

Here's the plan:

1. Announce to the team that Alan is leaving and inform them that the team is
   likely to be moved to Platform Security, continuing to work on much the same things.
2. Check with Jan (and others really) whether they'd like to stay and work on Angel
3. Pitch final plan to Sudhi/Chad.
4. Set a final transition date.


## Nomination_Review_Notes_Spring_2021.md.md on 0001-01-01
---
tags:
  - research
---




impact above 28% vs 71% strongly above

feedback Mo Yu: not helpful, no examples, no references to ladder

Huaxin mentions that she "mostly owns" ODML and ML using dynamic signals, what
does "mostly" means in this context I'm not fully sure.

Very, very quick ramp-up


L4 evidence:

* Sustained Track Record: perhaps the weakest area here, because Alice is going
  for promo only after 1 EE cycle, and it is unclear if she can sustain this
  level of performance.

* Impact: consistent with L4 requirements. Creating TFLite model on her own
  shows ability to start and finish work artifacts end-to-end. However, there is
  a concern that she only owned the improvements to the model - i.e. basically
  running experiments and tweaking parameters. I do not understand what "Sole
  owner for ML part of ODAD" means. 

* Difficulty: researching ODAD feels like meets the bar. 

* Leadership: collaboration with other groups, providing ML advice to samyou@
  and spearheading ML research on the team meets the bar.

There are concerns about the coding output. Alice averages ~20 CLs per quarter,
which is probably on the low side for an L4. I recognize that Alice is mostly
focussed on ML problems, but for an L4 mastering coding skills is an implied
requirement (Is not expected to have mastered all of these areas [above --
reliability monitoring etc.], but should have mastered at least one such
activity from the full Google "toolkit" outside of core coding).







Creating TFLite model for ODML


Worked with wejin@L6 to deploy models, docteau@L6 on the evaluation pipeline




Sole owner for ML part of ODAD.


ML model improvement


Owned model improvements in all areas


ODML model improvements research?
Owned TFLite evaluation e2e



Presented at the ODML summit



Co-hosted intern


One comment N/A - pattern matches nicely to L5 but not L6 support.

Impact aligned with expectations / independence aligned with expectations from same
reviewer, L7.

No access to Uraniborg doc

BTS Uraniborg Integration Plan doesn't have Billy's name on it at all, although
some comments are definitely present there -- despite being listed as "Started
Uraniborg+BTS".

Uraniborg v2 is in draft, impact unclear.

Impact of Uraniborg itself is fairly clear. It is not clear to me what was the
impact of Billy on Uraniborg (was he the sole person working on it? TL?
contributor?).

P21 antiphishing hasn't landed yet, as such impact is unclear. I

What is GBT? Access denied on both docs P21 and GBT. I


Uraniborg is that sort of problem, but the extent of ownership is unclear (and
so is impact of Uraniborg v2). Although the phrase "He has shown clear
leadership in identifying the preload safety problem before it became
mainstream" implies that it was his project all along.


Some aspects of leadership are found in consulting AiAi team, but unclear how
much that was.


Not immediately evident from the packet.

Past promo committee feedback is not visible (empty??) and not explicitly addressed in the
packet.




Sustained Track Record - evidence of sustained impact with 1xlong cycle SEE + 1
last one SEE
Impact - Owning SmartLock shows impact significantly exceeding what would
be expected from L3 and indeed L4 engineer. Work on single source migration and
security improvements to SmartLock shows substantial impact on the product,
while keeping product alive shows significant impact on the ecosystem.
Difficulty - work on Smartlock in GmsCore is evidently difficult enough for L4.
Leadership - it was clear they were managing their priorities and making
appropriately paced progress, working independently. No immediate evidence of
suggesting areas of future work, but this is not a requirement.
Other (optional) - very strong peer support from senior peers

Notes to Manager (optional) -


Sustained Track Record - shows sustained track record for the last 1.5 years
Impact - It's evident that Jake's work had direct impact on P2B since he
authored most/entire backend and launched it before hard deadline.
Difficulty - 5K LoC for PBC, difficulty unclear re "First launched “new stack”
component." - whether it was truly difficult and why? 
Leadership - demonstrated independent execution for both P2B and Mokka, with
substantial design.

Other (optional) - strong support from senior peers (L5s)
Notes to Manager (optional) - 


Sustained Track Record -  1xlong cycle @ SEE, 1xSEE demonstrate sustained track record
Difficulty - demonstrated by analysing packed code, ultimately resulting in new
dynamic unpacker and anti-cloaking module.
Impact - substantial impact beyond the scope of the team as demonstrated by
unpacking, anti-cloaking, and election protection work.
Leadership - demonstrates sustained leadership by working with multiple
teams/groups as well as driving his own work independently; zTracer.
Community - some community contributions present, in line with level.
Other (optional) - 
Notes to Manager (optional) -

"An L4 SE has expertise in at least two security domains and is either
developing deep technical expertise in a specific focus area or a broader, big
picture view within their specific focus area." -- the work on
reverse-engineering obfuscated/packed/anti-emulation samples, implementing
dynamic unpacker integrated to Marmot AS WELL AS general app review work.

"An L4 typically makes substantial contributions to an existing project or
creates a new solution to an established security problem." - building new
dynamic unpacker which unpacks 18% more apps than existing solution.

"Where called upon to own operational work, an L4 SE should be taking ownership
of more complex operational tasks" - supported other teams with
reversing/reviewing apps, election protection spot bonus, stalkerware WG peer
bonus. 

"The expected impact of an L4 SE is felt at the small team level. " -- review
work + unpacking work + Anti Cloaking Module is substantial for beyond small
team level.



Sustained Track Record -  1xEE long cycle, 1xSEE demonstrate sufficient track record
Difficulty - difficulty is evident from the work on reverse engineering some of
the more prolific apps and areas.
Impact - off-market and stalkerware WG show impact consistent with L5 requirements.
Leadership - starting off and leading multiple workgroups demonstrates
substantial leadership.
Community - very strong community contributions.
Other (optional) - unanimous peer support
Notes to Manager (optional) -



"A Sr. Security Engineer demonstrates independent ownership of a solution to a
security problem or area" - stalkerware and off-market malware were his areas of
ownership (although it is fair to say that neither problem are solved or even
could be solved at this level).

"Sr. SEs contribute to cross-org efforts " -- provided RE support for
go/android-app-visibility-ap-ps to guide Platform changes.

"At L5, Sr. SEs are expected to consistently demonstrate strong leadership" --
ownership of off-market WG, something that is often forgotten. Owner of
stalkerware WG. Strong mentoring efforts.

"Where called upon to own operational work, L5 Sr. SEs are expected to own
complex and difficult operational tasks" -- Olivier's analysis is of
consistently very high quality, e.g. Herole, RedStone and others. They

"The expected impact of an L5 Sr. SE is felt at the larger team or program
level" -- it is basically so, with impact feeling at ASAP level and well beyond.
Go-to person and the recognised expert on stalkerware.

## Notes_for_promo_session_March_11.md.md on 0001-01-01


* Owns multiple (3) projects of very substantial scope; in this cycle
  accomplished multiple launches on all of them materially improving core
  metrics. This is a requirement for L7 therefore it strongly exceeds L6.

* Architecting [Unified Developer Treatment
  Service](https://docs.google.com/document/d/1TO86ZLaavwEQu3OA8M6AL860ouY-ZIBddbFixoPLOTo/edit#heading=h.ljk9srdl7b0u)
  qualifies as very-near "Identify the broad problem worth solving" - to fully
  meet the bar for this the whole Guardrail should've been PG's invention.

  ## Leadership

  * Owning the entire area of DevRisk+Guardrail counts as "Responsibility for a
    key technical area", through technical and thought leadership as proven by
    design documents, as well as the management.
  * x-PA collab with Core counts as  way above L6's "Coordinating (..) multiple teams".


* ""Saved the day" multiple times on very sensitive escalations. Had to
  formulate/deliver solutions rapidly in urgent situations." - this arguably
  counts of L7's "Usually exhibit a high level of creativity and innovation
  ...".
* "Conflicting business objectives like user harm and developer satisfaction."
  is beyond L6's "thrives on ambiguity"


Support SEE.

## Notificaitons_and_Pull.md on 0001-01-01
---
tags:
  - research
---


## Nino.md on 0001-01-01
---
tags:
  - 1:1s
  - odad
---


## ODAD_Research.md on 0001-01-01
---
tags:
  - research
---



- ODAD team
- Someone from dynamic analysis team
- Richard Neal
- Few more folks from malware analysis
- Alan
- Jan


## Package_Verifier_Research.md on 0001-01-01
---
tags:
  - research
---


By the end of the day, my goal is not to explain what needs to be done. My goal
is to create an organisational framework when this can be figured out, have my
engineers come up with recommendations and *then* sell this idea to execs to get
proper funding.

Next steps would be: 

- Agree with William that a brainstorm on package verifier is a good idea
- Send out an email and invite folks to participate in the brainstorm
- Produce meaningful vision of where verifier must be
- Attempt to get this idea funded
- Build a new team in LON and hand over verifier work to it.


The Package Verifier is a component that checks if apps installed on an Android
device are malware (e.g. spy on the user) or pose some other threat to the user
(e.g. were banned from Play for violating a Play policy, or show pop-up ads that
make it difficult to use the device). As such, the package verifier is a
critical component of app safety - arguably, one of the most critical we have. 

On the top of that, package verifier is one of the few components which can
"break the world" as far as Android is concerned, by blocking all app installs,
including Phonesky self-update.

Package verifier has been jointly co-owned by Client Software and GPP teams.
However, with both teams having multiple high priority work to handle, package
verifier was largely uninvested into for a long time. This resulted in a number
of problems:

- Largely monolithic code, that has been modified for many years, becoming more
  and more complex, making it difficult to understand holistically how it works.

- Reduced testability - with multiple features closely coupled with each other,
  testing each individual one becomes progressively more complicated because of
  various side effects.

- No sufficient test infrastructure. For a critical component as that, we must
  ensure that *all* critical user journeys are covered by automatic tests. It is
  difficult/impossible to achieve within the current infra.

Resolving these and other problems, on the other hand, presents us with an
*opportunity* to make the package verifier more robust, allow to release new
features / experiment faster, ...

(what else does this enable?)


## Package_Verifier_Research.md.md on 0001-01-01
---
tags:
  - research
---


The problem as  I understand it that we now have two technologies, each
responsible for detecting makware on the device.

PV can mostly do static detection. ODAD can do both static and dynamic but cant
send anything to Google.

We  have a number of options to consolidate this

If we are happy with sending everything via Astrea then we can consolidate all
analysis in the ODAD

We can't however offload actual package verification to ODAD because it requires
round trip with Marmot.

We can have static analysis always done in package verifier, sending results to
Google but also sharing them with ODAD via unidirectional API. This would
effectively allow us to deliver cloud derived signals to  ODAD without violating
PCC model.


**ODAD can**
 
- Inspect private app data
- request additional signals from the platform (albeit this requires
  platform release)
- Receive new models from Google
- Perform dynamic and static analysis of apps

**ODAD cannot**

- Communicate to Google or indeed anything outside platform and PCC
- Provide (directly) any UX to the user

**GPP can**

- Provide direct UX to the user
- Have bi-directional communication with Google
- Perform static analysis of apps with limited dynamic analysis capabilities


**GPP cannot**

- Access most dynamic analysis signals
- Inspect any app private data


**Angel can**

- Provide UX
- Communicate with Google

**Angel cannot**

- Provide any app scanning capabilities, right now anyhow. Although this
  question is debatable - technically we can shift package verifier to Angel
  just to move it out of Phonesky.

Yet another thing to consider is that there are many UX flows in GPP, and we
don't necessarily want to re-build them in Angel (not yet, anyway).



GPP runs static analysis, checks results with Google,  discovers app is ...
interesting, instructs ODAD to be more diligent about it, ODAD does more
detailed dynamic analysis of what the app is doing.


We probably need both Package Verifier and ODAD going forward. The former is the
first line of defense against bad app installs since it can talk to Google and
ask if it's bad. The latter is the last line of defense since it can actually
inspect the behavior.


With the goal of user being able to do _everything_ in Angel, can we "unbundle"
UX flows from GPP so that the user will not know they've left Angel? Can we
angelise GPP flows? Can we angelise entire GPP?

## Permission_focused_AP_filter.md on 0001-01-01
---
tags:
    - meeting
---


## Peter.md on 0001-01-01
---
tags:
  - 1:1s
  - team
  - aleksandr
title: 'Peter:Roman'
---


## Productionising_VIP_apps_pipeline.md on 0001-01-01
---
tags:
  - research
---

Example CL for creating KG view: 
https://critique-ng.corp.google.com/cl/200641055/depot/google3/knowledge/graph/livegraph/access/batch/query/clients/peopleview/kg-entity-names.sql

Current access problem:

```
  _LegacyStatusNotOK: Not authorized to access table 'kgraph.base.triples.latest': datawarehouse_datascape::PERMISSION_DENIED: permission READER denied for (TABLE) kgraph.base.triples to user 'meerkat-ap4a-backlinks'; Not authorized to access table 'play_console.developers': datawarehouse_datascape::PERMISSION_DENIED: permission READER denied for (TABLE) play_console.developers to user 'meerkat-ap4a-backlinks' [UNKNOWN]
```




## Showing ODAD verdicts.md on 0001-01-01

It sounds like showing verdicts in GPP is simply too difficult.

ODAD is very likely to coincide with our Angel UX

Technically we can even expose it via the API

Why cannot OEM integrate our view of GPP into their implementation of the SC?

Embedding UI into GPP is a no go

Halfway integration between the GPP and Angel?

https://docs.google.com/presentation/d/1URE22UOWEIqaY8sxyO001zwM2oCQ_-OZIelTgIVZKrI/edit?resourcekey=0-_hsHnuPsvrj5TG8VmNcqzg#slide=id.g110fffad98a_7_0 -- the deck slide about GPP -> Angel integration


## Thoughts and Questions.md on 0001-01-01

Promotion to Director. Напрашивается вопрос нахера козе баян но ладно.




This is something I need to discuss with UX people - ultimately, having a strategy for measuring success of our products is something fairly critical.

GPP and Angel should be measured separately. Or not?

Metrics of success for GPP - we have Niko working on metrics for reliability but not success

Synthetic metric representing safety of the users as a whole, by looking at Angel status / warnings / GPP warnings / FMD status etc.



## Thoughts on review.md on 0001-01-01
---
tags:
    - research
---


We assume that the app is requesting at least some sensitive permissions. If no
sensitive permissions are requested, the app is good to go. For many of the app
categories below the fact that they are requesting multiple sensitive
permissions is enough of a reason to suspect some sort of foul play.


The general guidance here: if the app is requesting multiple sensitive
permissions which are clearly not related to the app's main functionality (e.g.
video player requesting location, notifications etc) we should not be including
it. In most cases making this decision should take ~ seconds.


- Probably approve


- Probably not approve


- Can one easily go and buy/subscribe on the website?

	- Yes: probably approve
	- No: e.g. app is required to subscribe: YMMV


- One of the major brands: Firefox, Opera, Brave, etc?

	- Yes: approve
  - No: do not approve. Browsers are inherently risky business, it's a good idea
    only allow the biggest ones which are likely to follow the best security
    practices.


- Reject. Anything that in a nutshell is a book and requests multiple sensitive permissions is suspicious. 


- Reject. Not needed.


- One of the major brands?

	- Yes: probably approve
	- No: do not approve, as what we are doing here is likely to provide more rigorous protection.


- One of the bigger brands?

	- Yes: approve

	  With the caveat of WeChat (add link)

	- No: not approve. If user really needs it they'll install it anyway and we'll count it.


- Difficult to say, need to deal on a case-by-case basis. If comes from a major or well known brand (e.g. Uber, CityMapper, Yandex.*) then we should approve it, otherwise not.


- Probably need to pre-approve top N launchers and leave it at that.


- If comes from a major brand, approve, otherwise there are plenty of alternatives. Either way they shouldn't request multiple sensitive permissions.


- Reject unless major brand. They shouldn't request sensitive permissions.


- Reject. Shouldn't be requesting major permissions.



## Titaniums quick research.md on 0001-01-01
---
title: Titaniums quick research
created: '2020-04-25T08:07:05.166Z'
modified: '2020-04-25T08:07:59.115Z'
tags: ['research']
---


Magisk package name: `com.topjohnwu.magisk`.


## Adam꞉Roman.md on 0001-01-01
---
tags: ['1:1s']
title: 'Adam:Roman'
created: '2020-04-24T14:18:55.996Z'
modified: '2020-04-24T14:19:53.741Z'
---



- Finishing migration, rewards program.
- ASE tour

Feedback about Henry's good work.


## Aaron.md on 0001-01-01
---
tags:
  - 1:1s
  - team
  - odad
---


