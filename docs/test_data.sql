INSERT INTO projects(name, logo) VALUES ('Todu', '/pub/uploads/project_logos/1.png');
INSERT INTO projects(name, logo) VALUES ('Kudos', '/no_image_available.png');
INSERT INTO projects(name, logo) VALUES ('Payment', '/pub/uploads/project_logos/3.png');
INSERT INTO projects(name, logo) VALUES ('QA', '/no_image_available.png');

INSERT INTO backlogs("number", "type", "text", "order", "status", "point", "project_id") VALUES 
(1, 'user_story', 'Product owner, төслөө сонгож орохын тулд, төслүүдээ хармаар байна.', 1, 'done', 1, 1),
(2, 'user_story', 'Product owner, төслөө сонгоод доторх product backlog-оо хянахын тулд, төслийн product backlog-оо хармаар байна.', 1, 'done', 3, 1),
(3, 'user_story', 'Product owner, хэрэгцээндээ тааруулж, backlog-оо User story форматаар бичдэг баймаар байна.', 1, 'in_progress', 2, 1),
(4, 'user_story', 'Product owner, хөгжүүлэгчид ажлын дарааллыг тодорхой болгохын тулд, backlog-оо эрэмбэлдэг баймаар байна.', 1, 'pending', 64, 1),
(5, 'user_story', 'Scrum master хөгжүүлэлт хэзээ дуусахыг тооцолохын тулд, backlog-д оноо тавьдаг баймаар байна.', 1, 'todo', 1, 1),
(6, 'user_story', 'Scrum master backlog-ыг хөгжүүлэгчид хуваахын тулд, backlog-г дотор нь subtask болгон задалмаар байна.', 1, 'rejected', 5, 2),
(7, 'user_story', 'Scrum master backlog-ыг хөгжүүлэгчид хуваарьлахын тулд, backlog бүр дээр хүн хуваарьлаж болдог баймаар байна.', 1, 'done', 1, 2),
(8, 'user_story', 'Scrum master subtask-ыг хэн хийх нь тодорхой байлгахын тулд, subtask-ийг хуваарьлаж болдог баймаар байна.', 1, 'done', 8, 2),
(9, 'user_story', 'Scrum master subtask-ыг үнэлэхийн тулд, subtask-д оноо өгдөг баймаар байна.', 1, 'done', 2, 2),
(10, 'technical_task', 'Бүх оролтыг Validate хийх', 1, 'done', 2, 1),
(11, 'bug', 'Бүх юм яагаад саарал байгаад байгаан :(', 1, 'done', 2, 1);

INSERT INTO subtasks ("text", "order", "status", "point", "backlog_id") VALUES
('Subtask 1', 1, 'done', 1, 1),
('Subtask 2', 1, 'todo', 1, 1),
('Subtask 3', 1, 'in_progress', 1, 1),
('Subtask 4', 1, 'done', 1, 2),
('Subtask 5', 1, 'todo', 1, 2),
('Subtask 6', 1, 'in_progress', 1, 2),
('Subtask 7', 1, 'done', 1, 3),
('Subtask 8', 1, 'todo', 1, 3),
('Subtask 9', 1, 'in_progress', 1, 3);