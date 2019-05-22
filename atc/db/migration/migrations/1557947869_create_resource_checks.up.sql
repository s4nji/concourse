BEGIN;

  CREATE TABLE resource_checks (
      id integer,
      resource_config_scope_id integer REFERENCES resource_config_scopes(id) ON DELETE CASCADE,
      start_time timestamp without time zone,
      end_time timestamp without time zone,
      create_time timestamp without time zone DEFAULT now() NOT NULL,
      from_version jsonb,
      modified_time timestamp without time zone DEFAULT now() NOT NULL
  );

COMMIT;
