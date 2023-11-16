INSERT INTO users(user_id,github_id) VALUES('hoge','hoge')
INSERT INTO users(user_id,github_id) VALUES('fuga','fuga');

INSERT INTO skills(skill_id,skill_name,required_bp,skilltype,"value",description) VALUES(1,'攻撃',0,'attack',1,'通常攻撃');
INSERT INTO skills(skill_id,skill_name,required_bp,skilltype,"value",description) VALUES(2,'防御',0,'defence',1,'防御');


INSERT INTO gitmons(gitmon_id,owner_id,gitmon_name,exp,base_hp,current_hp,base_attack,current_attack,base_defence,current_defence,base_speed,current_speed) VALUES('hogehoge','hoge','ほげもん',100,1500,1500,100,100,100,100,100,100);
INSERT INTO gitmons(gitmon_id,owner_id,gitmon_name,exp,base_hp,current_hp,base_attack,current_attack,base_defence,current_defence,base_speed,current_speed) VALUES('fugafuga','fuga','ふがもん',100,1500,1500,100,100,100,100,100,100);
INSERT INTO gitmon_skills(gitmon_id,skill_id,is_active) VALUES('hogehoge',1,true);

INSERT INTO gitmon_skills(gitmon_id,skill_id,is_active) VALUES('hogehoge',2,true);
INSERT INTO gitmon_skills(gitmon_id,skill_id,is_active) VALUES('fugafuga',1,true);
INSERT INTO gitmon_skills(gitmon_id,skill_id,is_active) VALUES('fugafuga',2,true);
