BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "genders" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"gender_type"	text,
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "students" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"s_id"	text,
	"password"	text,
	"first_name"	text,
	"last_name"	text,
	"year"	integer,
	"birthday"	datetime,
	"major"	text,
	"gender_id"	integer,
	PRIMARY KEY("id" AUTOINCREMENT),
	CONSTRAINT "fk_genders_students" FOREIGN KEY("gender_id") REFERENCES "genders"("id")
);
CREATE TABLE IF NOT EXISTS "dorms" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"type"	text,
	"gender_id"	integer,
	CONSTRAINT "fk_genders_dorms" FOREIGN KEY("gender_id") REFERENCES "genders"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "rooms" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"room_number"	integer,
	"available"	text,
	"confirmation"	text,
	"dorm_id"	integer,
	CONSTRAINT "fk_dorms_rooms" FOREIGN KEY("dorm_id") REFERENCES "dorms"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "reservations" (
	"id"	integer,
	"created_at"	datetime,
	"updated_at"	datetime,
	"deleted_at"	datetime,
	"reserve_date"	datetime,
	"student_id"	integer,
	"dorm_id"	integer,
	"room_id"	integer,
	CONSTRAINT "fk_dorms_reservations" FOREIGN KEY("dorm_id") REFERENCES "dorms"("id"),
	CONSTRAINT "fk_rooms_reservations" FOREIGN KEY("room_id") REFERENCES "rooms"("id"),
	CONSTRAINT "fk_reservations_student" FOREIGN KEY("student_id") REFERENCES "students"("id"),
	PRIMARY KEY("id" AUTOINCREMENT)
);
CREATE INDEX IF NOT EXISTS "idx_genders_deleted_at" ON "genders" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_students_deleted_at" ON "students" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_dorms_deleted_at" ON "dorms" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_rooms_deleted_at" ON "rooms" (
	"deleted_at"
);
CREATE INDEX IF NOT EXISTS "idx_reservations_deleted_at" ON "reservations" (
	"deleted_at"
);
COMMIT;
