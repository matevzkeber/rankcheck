package commands

import (
	"RankCheck/embeds"
	"RankCheck/globals"

	"github.com/bwmarrin/discordgo"
)

var (
	LinksCommand = Command{
		Name: "links",
		Command: discordgo.ApplicationCommand{
			Name:        "links",
			Description: "All the important links for RankCheck",
		},
		Execute: linksCommand,
	}
)

func linksCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := embeds.EmbedWithPFP(s)
	embed.Title = "RankCheck Links"
	embed.Description = "RankCheck is an ever growing project, and we have a couple of links we believe you'll find useful!"
	embed.Fields = []*discordgo.MessageEmbedField{
		{
			Name:   "Invite RankCheck",
			Value:  "Invite RankCheck to your server!",
			Inline: false,
		},
		{
			Name:   "Website",
			Value:  "Check out our website for more information about RankCheck!",
			Inline: false,
		},
		{
			Name:   "Matevž's website",
			Value:  "Check out my personal website!",
			Inline: false,
		},
		{
			Name:   "Faith's Instagram",
			Value:  "Faith is the designer of the RankCheck profile picture and a valued member of the team!\nCheck out her Instagram, but you can also reach her on Discord `@faith.art28`.",
			Inline: false,
		},
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed,
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "📨",
							},
							Label: "Invite RankCheck",
							Style: discordgo.LinkButton,
							URL:   globals.BotInvite,
						},
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "🌍",
							},
							Label: "Website",
							Style: discordgo.LinkButton,
							URL:   globals.Website,
						},
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "🌍",
							},
							Label: "Matevž's website",
							Style: discordgo.LinkButton,
							URL:   "https://matevzkeber.si",
						},
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "🎨",
							},
							Label: "Faith's Instagram",
							Style: discordgo.LinkButton,
							URL:   "https://www.instagram.com/faith.art28",
						},
					},
				},
			},
			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
}
