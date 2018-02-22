# Migrations

All migrations should be written with a `up` and `down` scripts that will
apply and remove the schema changes respectively.

From [`mattes/migrate`](https://github.com/mattes/migrate)
[best practices](https://github.com/mattes/migrate/blob/master/MIGRATIONS.md)
docs:
> The ordering and direction of the migration files is determined by the filenames used for them. migrate expects the filenames of migrations to have the format:
>
> {version}_{title}.up.{extension}
> {version}_{title}.down.{extension}
>
> The title of each migration is unused, and is only for readability. Similarly, the extension of the migration files is not checked by the library, and should be an appropriate format for the database in use (.sql for SQL variants, for instance).

Please read the [best practices](https://github.com/mattes/migrate/blob/master/MIGRATIONS.md)
for more information and details.
