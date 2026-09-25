CREATE TABLE IF NOT EXISTS users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(80) NOT NULL,
  major VARCHAR(120) NOT NULL,
  credit_score INT NOT NULL DEFAULT 80,
  credit_level VARCHAR(40) NOT NULL
);

CREATE TABLE IF NOT EXISTS skills (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  user_id BIGINT NOT NULL,
  title VARCHAR(120) NOT NULL,
  category VARCHAR(40) NOT NULL,
  level_score INT NOT NULL,
  campus VARCHAR(40) NOT NULL,
  description TEXT NOT NULL,
  portfolio VARCHAR(160) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS needs (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  user_id BIGINT NOT NULL,
  requester_name VARCHAR(80) NOT NULL,
  title VARCHAR(120) NOT NULL,
  category VARCHAR(40) NOT NULL,
  campus VARCHAR(40) NOT NULL,
  expect_time VARCHAR(80) NOT NULL,
  slot_code VARCHAR(40) NOT NULL,
  budget_type VARCHAR(40) NOT NULL,
  description TEXT NOT NULL,
  status VARCHAR(20) NOT NULL DEFAULT 'open',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 同学对需求的响应：交换说明与空闲时段，同一条需求每人仅一条
CREATE TABLE IF NOT EXISTS need_responses (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  need_id BIGINT NOT NULL,
  respondent_name VARCHAR(80) NOT NULL,
  offer_note VARCHAR(200) NOT NULL,
  free_slot_codes VARCHAR(200) NOT NULL COMMENT '逗号分隔的时段编码',
  status VARCHAR(20) NOT NULL DEFAULT 'waiting' COMMENT 'waiting/selected/released',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY uk_need_respondent (need_id, respondent_name),
  INDEX idx_need (need_id)
);

CREATE TABLE IF NOT EXISTS appointments (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  need_id BIGINT NOT NULL DEFAULT 0,
  requester_name VARCHAR(80) NOT NULL,
  respondent_name VARCHAR(80) NOT NULL,
  slot_code VARCHAR(40) NOT NULL,
  exchange_time VARCHAR(80) NOT NULL,
  place VARCHAR(120) NOT NULL,
  agenda TEXT NOT NULL,
  status VARCHAR(20) NOT NULL DEFAULT 'pending' COMMENT 'pending/confirmed/cancelled',
  requester_confirmed TINYINT(1) NOT NULL DEFAULT 0,
  respondent_confirmed TINYINT(1) NOT NULL DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  INDEX idx_slot_person (slot_code, respondent_name),
  INDEX idx_need (need_id)
);

CREATE TABLE IF NOT EXISTS reviews (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  from_user VARCHAR(80) NOT NULL,
  to_user VARCHAR(80) NOT NULL,
  rating INT NOT NULL,
  content TEXT NOT NULL
);

INSERT INTO users(name, major, credit_score, credit_level) VALUES
('林澈', '新闻传播 2023', 91, '黄金导师'),
('孟野', '音乐表演 2022', 88, '白银协作者'),
('周芮', '统计学 2021', 93, '黄金导师');

INSERT INTO skills(user_id, title, category, level_score, campus, description, portfolio) VALUES
(1, '毕业照人像摄影', '摄影', 92, '东校区', '提供构图、修图和毕业季跟拍，可交换吉他入门课。', '12组校园人像作品'),
(2, '民谣吉他陪练', '乐器', 81, '西校区', '节奏型、弹唱和舞台经验分享，想找人拍宣传照。', '校园音乐节演出视频'),
(3, 'Python 数据分析', '编程', 88, '中心校区', 'pandas、可视化、论文数据清洗辅导。', '3份课程项目证书');

INSERT INTO needs(user_id, requester_name, title, category, campus, expect_time, slot_code, budget_type, description, status) VALUES
(2, '孟野', '找人帮忙拍乐队宣传照', '摄影', '西校区', '周六上午', 'sat-morning', '技能交换', '可交换 3 次吉他课，希望会调色和室外构图。', 'open'),
(4, '许安', '求教 Python 数据分析', '编程', '中心校区', '周二晚', 'tue-evening', '小额报酬', '论文问卷数据需要清洗和画图，最好有 pandas 经验。', 'open'),
(1, '林澈', '想学吉他扫弦入门', '乐器', '东校区', '周三晚', 'wed-evening', '技能交换', '用摄影课交换吉他基础，希望同校区或线上。', 'open');

INSERT INTO need_responses(need_id, respondent_name, offer_note, free_slot_codes, status) VALUES
(1, '林澈', '擅长室外自然光和调色，可用毕业照跟拍经验交换吉他课。', 'wed-evening,sat-morning', 'waiting'),
(1, '唐宁', '摄影社成员，带过乐队现场拍摄，设备齐全可出原片。', 'sat-morning,sun-all-day', 'waiting'),
(1, '韩梅', '会后期精修，可先看作品集再决定，接受技能交换。', 'fri-evening,sat-morning', 'waiting'),
(2, '周芮', 'pandas 与论文图表经验丰富，可按次结算报酬。', 'tue-evening,sun-all-day', 'waiting'),
(2, '唐宁', '统计方向研究生，熟悉问卷信效度和回归分析。', 'thu-evening,sun-all-day', 'waiting'),
(3, '孟野', '可从扫弦节奏型讲起，顺带教弹唱配合，想换一次人像拍摄。', 'wed-evening,sat-morning', 'waiting'),
(3, '周芮', '民谣弹唱业余三年，适合陪练入门，时间固定周三晚。', 'wed-evening', 'waiting');

INSERT INTO appointments(need_id, requester_name, respondent_name, slot_code, exchange_time, place, agenda, status, requester_confirmed, respondent_confirmed) VALUES
(0, '顾远', '林澈', 'sat-morning', '周六上午 10:00', '东校区湖边', '社团招新跟拍，约两小时', 'confirmed', 1, 1),
(0, '许安', '周芮', 'tue-evening', '周二晚 19:30', '线上会议室', '导入问卷 CSV 并完成基础可视化', 'pending', 1, 0);
