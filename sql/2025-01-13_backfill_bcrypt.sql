-- Idempotent bcrypt backfill for existing users.
-- Rehash any password_hash values that are not already bcrypt ($2a/$2b/$2y).
-- Assumes pgcrypto is installed (schema.sql already creates the extension).

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM information_schema.routines WHERE routine_name = 'gen_salt') THEN
    RAISE NOTICE 'pgcrypto not installed; skipping bcrypt backfill';
    RETURN;
  END IF;

  UPDATE users
  SET password_hash = crypt(password_hash, gen_salt('bf', 12)),
      updated_at = NOW()
  WHERE password_hash NOT LIKE '$2a$%%'
    AND password_hash NOT LIKE '$2b$%%'
    AND password_hash NOT LIKE '$2y$%%'
    AND username <> 'guest';
END$$;
