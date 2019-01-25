/* Insert user */
INSERT INTO `user`(`name`, `password`) VALUES
    ('asuka', 'asuka'),
    ('foo', 'foo'),
    ('bar', 'bar');

/* Insert monitor */
INSERT INTO `monitor`(`observer`, `target`, `rate`, `review`) VALUES
    (2, 1, 4, 'foooooooo'),
    (3, 1, 3, 'baaaaaaar');

/* Insert category */
INSERT INTO `category`(`name`) VALUES
    ('Real'),
    ('Comic'),
    ('Gothic'),
    ('Fantasy'),
    ('Science fiction'),
    ('Mechanic');

/* Insert works */
INSERT INTO `work_wanted`(`user`, `title`, `description`, `price`) VALUES
    (1, 'aaaaa', 'zzzzzzzzzz', 10000);

INSERT INTO `work_request`(`user`, `wanted`, `requester`, `title`, `description`, `price`) VALUES
    (1, 1, 2, 'foo', 'foooooo', 12000);

