/* Insert user */
INSERT INTO `user`(`name`, `password`) VALUES
    ('asuka', 'asuka'),
    ('foo', 'foo'),
    ('bar', 'bar'),
    ('aaa', 'aaa'),
    ('bbb', 'bbb'),
    ('ccc', 'ccc'),
    ('ddd', 'ddd'),
    ('eee', 'eee'),
    ('fff', 'fff'),
    ('ggg', 'ggg'),
    ('hhh', 'hhh'),
    ('iii', 'iii'),
    ('jjj', 'jjj'),
    ('kkk', 'kkk'),
    ('lll', 'lll'),
    ('mmm', 'mmm'),
    ('nnn', 'nnn'),
    ('ooo', 'ooo'),
    ('ppp', 'ppp'),
    ('qqq', 'qqq'),
    ('rrr', 'rrr'),
    ('sss', 'sss'),
    ('ttt', 'ttt'),
    ('uuu', 'uuu'),
    ('vvv', 'vvv'),
    ('www', 'www'),
    ('xxx', 'xxx'),
    ('yyy', 'yyy'),
    ('zzz', 'zzz');

/* Insert monitor */
INSERT INTO `monitor`(`observer`, `target`, `rate`, `review`) VALUES
    (2, 1, 4, 'とても良かった'),
    (3, 1, 2, '想像してたのと違う'),
    (4, 1, 4, '期待通りでした'),
    (5, 1, 1, 'なにこれ'),
    (6, 1, 0, '高すぎ'),
    (7, 1, 5, 'よかったです'),
    (8, 1, 4, 'よかったです'),
    (9, 1, 3, 'いいね'),
    (10, 1, 4, 'とても良かった'),
    (11, 1, 5, 'とても良かった'),
    (12, 1, 5, 'とても良かった'),
    (13, 1, 4, 'とても良かった');
INSERT INTO `monitor`(`observer`, `target`, `rate`, `review`) VALUES
    (1, 2, 4, '理解のある人でした'),
    (1, 3, 0, '後から追加で色々と条件を追加されました'),
    (1, 4, 4, '理解のある人でした'),
    (1, 5, 0, 'お金払うからなんでも言って言い訳無いでしょう'),
    (1, 6, 0, 'めっちゃ値切られました'),
    (1, 7, 5, '理解のある人でした'),
    (1, 8, 4, 'よかったです'),
    (1, 9, 3, 'いいね'),
    (1, 10, 4, 'またよろしくお願いします'),
    (1, 11, 5, 'またよろしくお願いします'),
    (1, 12, 5, 'またよろしくお願いします'),
    (1, 13, 4, 'またよろしくお願いします');

/* Insert category */
INSERT INTO `category`(`name`) VALUES
    ('Real'),
    ('Comic'),
    ('Gothic'),
    ('Fantasy'),
    ('Science fiction'),
    ('Mechanic');

/* Insert gallery */
INSERT INTO `gallery`(`user`, `image`) VALUES
    (1, 'sample.png');

/* Insert works */
INSERT INTO `work_wanted`(`user`, `title`, `description`, `price`) VALUES
    (1, 'Twitter用アイコンイラスト', '１枚あたり平均３日ほどお時間いただきます。', 10000),
    (1, 'キャラクターイラスト', 'A４サイズのキャラクターイラストを描きます。１枚あたり、平均１週間ほどお時間いただきます', 30000);
INSERT INTO `work_wanted`(`user`, `title`, `description`, `price`, `alive`) VALUES
    (1, 'キャラクター原案', 'コンセプトに基づいてキャラクターを作成します。', 100000, false),
    (1, 'コンセプト原案', 'コンセプトイラストを作成します', 150000, false);

INSERT INTO `work_request`(`user`, `wanted`, `requester`, `title`, `description`, `price`, `establish`, `alive`, `check`) VALUES
    (1, 1, 2, 'ツイッターのアイコンを書いてほしいです', '猫耳の黒髪ロング女の子のイラストを描いてほしいです。', 10000, true, false, true),
    (1, 1, 3, '作成依頼', '男の子のイラストをお願いします。デザインなどはおまかせします。', 10000, true, false, true),
    (1, 1, 4, 'アイコン作成依頼', '初音ミクのアイコンを描いてほしいです。', 10000, true, false, true),
    (1, 1, 5, 'アイコン依頼', '5000円でアイコン描いてください', 10000, false, false, true),
    (1, 1, 6, 'アイコンを描いてほしいです', '学生でお金無いので、1000円で描いてくれませんか', 10000, false, false, true),
    (1, 1, 7, 'Twitter用アイコンかいてほしいです。', 'お願いします。', 10000, true, false, true),
    (1, 1, 8, 'アイコン作成依頼', 'おねしょたきぼんぬ', 10000, true, false, true),
    (1, 1, 9, 'きぼんぬ', 'アイコン描いてください。お願いします。', 10000, true, false, true),
    (1, 1, 10, 'ツイッターのアイコンを書いてほしいです', '猫耳の黒髪ロング女の子のイラストを描いてほしいです。', 10000, true, false, true),
    (1, 1, 11, 'ツイッターのアイコンを書いてほしいです', '猫耳の黒髪ロング女の子のイラストを描いてほしいです。', 10000, true, false, true),
    (1, 1, 12, 'ツイッターのアイコンを書いてほしいです', '猫耳の黒髪ロング女の子のイラストを描いてほしいです。', 10000, true, false, true),
    (1, 1, 13, 'ツイッターのアイコンを書いてほしいです', '猫耳の黒髪ロング女の子のイラストを描いてほしいです。', 10000, true, false, true);

