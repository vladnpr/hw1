package main

import (
	"fmt"
	TextAnimator "github.com/vladnpr/textAnimator"
)

type person struct {
	firstName   string
	surname     string
	birthDate   string
	dateOfDeath string
	yOld        uint
}

func (p person) biography() string {
	return fmt.Sprintf("%s %s (born Norma Jeane Mortenson; %s – %s) was an American actress, model, and singer. \n "+
		"Known for playing comic \"blonde bombshell\" characters, she became one of the most popular sex symbols of the 1950s and early 1960s, as well as an emblem of the era's sexual revolution. \n"+
		" She was a top-billed actress for a decade, and her films grossed $200 million (equivalent to $2 billion in 2022) by the time of her death in 1962.  \n"+
		"Long after her death, Monroe remains a pop culture icon. In 1999, the American Film Institute ranked her as the sixth-greatest female screen legend from the Golden Age of Hollywood."+
		"\n\nBorn and raised in Los Angeles, Monroe spent most of her childhood in a total of 12 foster homes and an orphanage before marrying James Dougherty at age sixteen. \n"+
		" She was working in a factory during World War II when she met a photographer from the First Motion Picture Unit and began a successful pin-up modeling career, \n"+
		" which led to short-lived film contracts with 20th Century Fox and Columbia Pictures. After a series of minor film roles, she signed a new contract with Fox in late 1950. \n"+
		" Over the next two years, she became a popular actress with roles in several comedies, including As Young as You Feel and Monkey Business, \n"+
		" and in the dramas Clash by Night and Don't Bother to Knock. Monroe faced a scandal when it was revealed that she had posed for nude photographs prior to becoming a star, \n"+
		" but the story did not damage her career and instead resulted in increased interest in her films."+
		"\n\nBy 1953, Monroe was one of the most marketable Hollywood stars. She had leading roles in the film noir Niagara, which overtly relied on her sex appeal, \n"+
		" and the comedies Gentlemen Prefer Blondes and How to Marry a Millionaire, which established her star image as a \"dumb blonde\".  \n"+
		"The same year, her nude images were used as the centerfold and on the cover of the first issue of Playboy.  \n"+
		"Monroe played a significant role in the creation and management of her public image throughout her career, but felt disappointed when typecast and underpaid by the studio. \n "+
		"She was briefly suspended in early 1954 for refusing a film project but returned to star in The Seven Year Itch (1955), one of the biggest box office successes of her career."+
		"\n\nWhen the studio was still reluctant to change Monroe's contract, she founded her own film production company in 1954. \n "+
		"She dedicated 1955 to building the company and began studying method acting under Lee Strasberg at the Actors Studio. \n "+
		"Later that year, Fox awarded her a new contract, which gave her more control and a larger salary. \n"+
		" Her subsequent roles included a critically acclaimed performance in Bus Stop (1956) and her first independent production in The Prince and the Showgirl (1957). \n"+
		"She won a Golden Globe for Best Actress for her role in Some Like It Hot (1959), a critical and commercial success. Her last completed film was the drama The Misfits (1961)."+
		"\n\nMonroe's troubled private life received much attention. She struggled with addiction and mood disorders. \n"+
		" Her marriages to retired baseball star Joe DiMaggio and to playwright Arthur Miller were highly publicized, but ended in divorce. \n "+
		"On August 4, 1962, she died at age %d from an overdose of barbiturates at her Los Angeles home. Her death was ruled a probable suicide.", p.firstName, p.surname, p.birthDate, p.dateOfDeath, p.yOld)
}

func main() {
	p := person{
		firstName:   "Marilyn",
		surname:     "Monroe",
		birthDate:   "June 1, 1926",
		dateOfDeath: "August 4, 1962",
		yOld:        36,
	}
	text := p.biography()

	animator := TextAnimator.NewTextAnimator(text)
	animator.PrintSequential()
}
