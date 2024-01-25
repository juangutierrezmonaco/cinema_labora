INSERT INTO theater (name, capacity, last_row, last_column, created_at) VALUES
('Theater 1', 150, 'Z', 20, 1643548800),
('Theater 2', 200, 'M', 15, 1643548800),
('Theater 3', 120, 'K', 18, 1643548800),
('Theater 4', 180, 'P', 25, 1643548800),
('Theater 5', 100, 'A', 10, 1643548800),
('Theater 6', 160, 'H', 22, 1643548800),
('Theater 7', 130, 'E', 17, 1643548800),
('Theater 8', 220, 'X', 30, 1643548800),
('Theater 9', 90, 'C', 12, 1643548800),
('Theater 10', 170, 'N', 23, 1643548800);

INSERT INTO screening (name, movie_id, theater_id, available_seats, taken_seats, showtime, price, language, views_count, created_at, updated_at) VALUES
('Screening 1', 101, 1, 150, '{A1,A2,B1,B2}', 1643552400, 12.50, 'English', 50, 1643548800, 1643548800),
('Screening 2', 102, 3, 180, '{C3,C4,D1,D2,D3}', 1643556000, 10.00, 'Spanish', 30, 1643548800, 1643548800),
('Screening 3', 103, 5, 120, '{H1,H2,I1,I2}', 1643559600, 15.00, 'French', 40, 1643548800, 1643548800),
('Screening 4', 104, 7, 200, '{E3,E4,F1,F2,F3}', 1643563200, 8.50, 'German', 20, 1643548800, 1643548800),
('Screening 5', 105, 2, 130, '{M1,M2,N1,N2}', 1643566800, 11.50, 'English', 60, 1643548800, 1643548800),
('Screening 6', 106, 4, 160, '{K3,K4,L1,L2,L3}', 1643570400, 14.00, 'Spanish', 25, 1643548800, 1643548800),
('Screening 7', 107, 6, 100, '{X1,X2,Y1,Y2}', 1643574000, 13.00, 'French', 35, 1643548800, 1643548800),
('Screening 8', 108, 8, 220, '{P3,P4,Q1,Q2,Q3}', 1643577600, 9.00, 'German', 15, 1643548800, 1643548800),
('Screening 9', 109, 10, 90, '{Z1,Z2,A3,A4}', 1643581200, 16.00, 'English', 45, 1643548800, 1643548800),
('Screening 10', 110, 9, 170, '{G3,G4,H1,H2,H3}', 1643584800, 12.00, 'Spanish', 55, 1643548800, 1643548800);

INSERT INTO "user" (first_name, last_name, email, password, gender, picture_url, created_at, updated_at) VALUES
('John', 'Doe', 'john@example.com', 'hashed_password', 'M', 'http://example.com/john.jpg', 1643548800, 1643548800),
('Alice', 'Smith', 'alice@example.com', 'hashed_password', 'F', 'http://example.com/alice.jpg', 1643552400, 1643552400),
('Bob', 'Johnson', 'bob@example.com', 'hashed_password', 'M', 'http://example.com/bob.jpg', 1643556000, 1643556000),
('Eva', 'Brown', 'eva@example.com', 'hashed_password', 'F', 'http://example.com/eva.jpg', 1643559600, 1643559600),
('Charlie', 'Davis', 'charlie@example.com', 'hashed_password', 'M', 'http://example.com/charlie.jpg', 1643563200, 1643563200),
('Grace', 'Lee', 'grace@example.com', 'hashed_password', 'F', 'http://example.com/grace.jpg', 1643566800, 1643566800),
('David', 'Wang', 'david@example.com', 'hashed_password', 'M', 'http://example.com/david.jpg', 1643570400, 1643570400),
('Sophie', 'Taylor', 'sophie@example.com', 'hashed_password', 'F', 'http://example.com/sophie.jpg', 1643574000, 1643574000),
('Michael', 'Clark', 'michael@example.com', 'hashed_password', 'M', 'http://example.com/michael.jpg', 1643577600, 1643577600),
('Emma', 'Johnson', 'emma@example.com', 'hashed_password', 'F', 'http://example.com/emma.jpg', 1643581200, 1643581200);

INSERT INTO ticket (pickup_id, user_id, screening_id, created_at) VALUES
('TICKETABC', 1, 1, 1643548800),
('TICKETDEF', 2, 2, 1643552400),
('TICKETGHI', 3, 3, 1643556000),
('TICKETJKL', 4, 4, 1643559600),
('TICKETMNO', 5, 5, 1643563200),
('TICKETPQR', 6, 6, 1643566800),
('TICKETSTU', 7, 7, 1643570400),
('TICKETVWX', 8, 8, 1643574000),
('TICKETYZA', 9, 9, 1643577600),
('TICKETBCD', 10, 10, 1643581200);

INSERT INTO comment (user_id, movie_id, content, created_at, updated_at) VALUES
(1, 101, 'Great movie!', 1643548800, 1643548800),
(2, 102, 'Loved the storyline.', 1643552400, 1643552400),
(3, 103, 'Amazing cinematography!', 1643556000, 1643556000),
(4, 104, 'A must-watch film.', 1643559600, 1643559600),
(5, 105, 'Incredible performances.', 1643563200, 1643563200),
(6, 106, 'Excellent direction!', 1643566800, 1643566800),
(7, 107, 'Captivating plot.', 1643570400, 1643570400),
(8, 108, 'Beautifully shot scenes.', 1643574000, 1643574000),
(9, 109, 'Engaging from start to finish.', 1643577600, 1643577600),
(10, 110, 'A cinematic masterpiece.', 1643581200, 1643581200);
