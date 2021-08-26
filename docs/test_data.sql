INSERT INTO projects(name, logo) VALUES ('Todu', '/pub/uploads/project_logos/1.png');
INSERT INTO projects(name, logo) VALUES ('Kudos', '/no_image_available.png');
INSERT INTO projects(name, logo) VALUES ('Payment', '/pub/uploads/project_logos/3.png');
INSERT INTO projects(name, logo) VALUES ('QA', '/no_image_available.png');

INSERT INTO user_stories("number", "as", "because", "want", "order", "status", "point", "project_id") VALUES 
(1, 'Product owner', 'Төслөө сонгож орохын тулд', 'Төслүүдээ хармаар байна.', 1, 'done', 1, 1),
(2, 'Product owner', 'Төслөө сонгоод доторх product backlog-оо хянахын тулд', 'Төслийн product backlog-оо хармаар байна.', 1, 'done', 3, 1),
(3, 'Product owner', 'Хэрэгцээндээ тааруулж', 'Backlog-оо User story форматаар бичдэг баймаар байна.', 1, 'in_progress', 2, 1),
(4, 'Product owner', 'Хөгжүүлэгчид ажлын дарааллыг тодорхой болгохын тулд', 'Backlog-оо эрэмбэлдэг баймаар байна.', 1, 'pending', 64, 1),
(5, 'Scrum master', 'Хөгжүүлэлт хэзээ дуусахыг тооцолохын тулд', 'Backlog-д оноо тавьдаг баймаар байна.', 1, 'todo', 1, 1),
(6, 'Scrum master', 'Backlog-ыг хөгжүүлэгчид хуваахын тулд', 'Backlog-г дотор нь subtask болгон задалмаар байна.', 1, 'rejected', 5, 2),
(7, 'Scrum master', 'Backlog-ыг хөгжүүлэгчид хуваарьлахын тулд', 'Backlog бүр дээр хүн хуваарьлаж болдог баймаар байна.', 1, 'done', 1, 2),
(8, 'Scrum master', 'Subtask-ыг хэн хийх нь тодорхой байлгахын тулд', 'Subtask-ийг хуваарьлаж болдог баймаар байна.', 1, 'done', 8, 2),
(9, 'Scrum master', 'Subtask-ыг үнэлэхийн тулд', 'Subtask-д оноо өгдөг баймаар байна.', 1, 'done', 2, 2);

INSERT INTO subtasks ("text","order","status","point","user_story_id") VALUES
('shdflaksj', 1, 'done', 1, 1),
('shdflaksj', 1, 'todo', 1, 1),
('shdflaksj', 1, 'in_progress', 1, 1);