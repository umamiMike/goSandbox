define(
	"main",
	[
		"MessageList"
	],
	function(MessageList) {
		var ws = new WebSocket("ws://7ac69774.ngrok.io/entry");
		var list = new MessageList(ws);
		ko.applyBindings(list);
	}
);
