CREATE TABLE otp (
                              otp_id SERIAL PRIMARY KEY,
                              phonenumber VARCHAR(25),
                              code VARCHAR(10) NOT NULL,
                              purpose VARCHAR(20) NOT NULL DEFAULT 'phone_verification', -- misal: 'phone_verification', 'password_reset'
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              expire_at TIMESTAMP GENERATED ALWAYS AS (created_at + INTERVAL '5 minutes') STORED,
                              CONSTRAINT chk_verified_otp_contact CHECK (phonenumber IS NOT NULL),
                              CONSTRAINT chk_verified_otp_purpose CHECK (purpose IN ('phone_verification', 'password_reset'))
);