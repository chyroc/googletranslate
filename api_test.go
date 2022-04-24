package googletranslate

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	type Case struct {
		Name     string
		Input    string
		Expected string
	}
	cases := []Case{
		{
			Name:     "你好",
			Input:    "hello",
			Expected: "你好",
		},
		{
			Name:     "长句子",
			Input:    "Its an app for writing and organizing markdown text documents. While this does not sound like much, it changed my work life a lot and gave me such a positive impact that I wrote this post to share the app with you.",
			Expected: "它是一个用于编写和组织 Markdown 文本文档的应用程序。 虽然这听起来并不多，但它极大地改变了我的工作生活并给了我如此积极的影响，以至于我写了这篇文章与你分享这个应用程序。",
		},
		{
			Name: "段落",
			Input: `The app itself offers a paid sync feature for $4/mo (as the time of this writing). Since the app writes plain text to a folder on your device, you get some other options tough.

I went ahead with [Syncthing], an open source file sync tool for Windows, MacOS, Linux and Android. Since it also works without a cloud or an account, it was the perfect solution for me. Apps like Dropbox or Google Drive dont sync a folder on your mobile and dont work in this case.

Setting it up on android was a bit tricky because of the security constraints but I eventually figured out how to make it work: the "Documents" Folder on my android device can be read/write accessed by any app with permission, so I set the Obsidian Vault location there. I then shared the vault folder through syncthing with my laptop.

Now, whenever I take notes on my mobile, I can pick them up on my laptop. Or when I write down extensive research on my laptop, I can browse and search it on my mobile when I'm on the run.

And since everything is markdown based, I even wrote this blog post in Obsidian. What a great and truly open piece of software!

I just started using obsidian in my everyday work and it already had a big impact on my workflow. I strongly suggest you to give it a try.`,
			Expected: `该应用程序本身提供 4 美元/月的付费同步功能（在撰写本文时）。 由于该应用程序将纯文本写入您设备上的文件夹，因此您会遇到其他一些困难的选择。 我继续使用 [Syncthing]，这是一个适用于 Windows、MacOS、Linux 和 Android 的开源文件同步工具。 因为它也可以在没有云或帐户的情况下工作，所以它对我来说是完美的解决方案。 由于安全限制，在 android 上设置它有点棘手，但我最终想出了如何让它工作：我的 android 设备上的“文档”文件夹可以被任何具有权限的应用程序读/写访问，所以我设置了黑曜石金库的位置在那里。 然后，我通过 syncthing 与我的笔记本电脑共享了保管库文件夹。 现在，每当我在手机上做笔记时，我都可以在笔记本电脑上拿起它们。 或者，当我在笔记本电脑上写下大量研究时，我可以在跑步时在手机上浏览和搜索它。 由于一切都是基于降价的，我什至在 Obsidian 中写了这篇博文。 多么棒的一款真正开放的软件！ 我刚开始在日常工作中使用黑曜石，它已经对我的工作流程产生了重大影响。 我强烈建议你试一试。`,
		},
	}
	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			res, err := Translate(v.Input, En, Zh)
			if err != nil {
				t.Errorf("Translate() error = %v", err)
				return
			}
			if res != v.Expected {
				t.Errorf("Translate() got = %v, want %v", res, v.Expected)
			}
		})
	}
}
