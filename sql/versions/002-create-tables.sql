CREATE TABLE accounts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    company_name TEXT NOT NULL,
    company_country_code NOT NULL,
    company_country_name NOT NULL,
    price_plan_id INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    expires_at INTEGER NOT NULL DEFAULT (0),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX accounts_expires_at ON accounts (expires_at);

CREATE TABLE accounts_absences (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    date INTEGER NOT NULL,
    duration INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX accounts_absences_account_id ON accounts_absences (account_id);
CREATE INDEX accounts_absences_date ON accounts_absences (date);

CREATE TABLE accounts_introductions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX accounts_introductions_account_id ON accounts_introductions (account_id);

CREATE TABLE accounts_invoices (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    price_plan_id INTEGER NOT NULL,
    amount INTEGER NOT NULL,
    due_on INTEGER NOT NULL,
    settled_at INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX accounts_invoices_account_id ON accounts_invoices (account_id);
CREATE INDEX accounts_invoices_due_on ON accounts_invoices (due_on);

CREATE TABLE accounts_settings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX accounts_settings_account_id ON accounts_settings (account_id);

CREATE TABLE accounts_usage (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX accounts_usage_account_id ON accounts_usage (account_id);

CREATE TABLE capabilities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX capabilities_account_id ON capabilities (account_id);

CREATE TABLE enum_durations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    value INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX enum_durations_value ON enum_durations (value);

CREATE TABLE enum_priorities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    value INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX enum_priorities_value ON enum_priorities (value);

CREATE TABLE enum_roles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    value INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX enum_roles_value ON enum_roles (value);

CREATE TABLE enum_statuses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    value INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX enum_statuses_value ON enum_statuses (value);

CREATE TABLE features (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    description_html TEXT NOT NULL,
    priority_name TEXT NOT NULL,
    priority_value INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX features_account_id ON features (account_id);
CREATE INDEX features_product_id ON features (product_id);
CREATE INDEX features_priority_value ON features (priority_value);
CREATE INDEX features_created_by ON features (created_by);

CREATE TABLE features_attachments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    feature_id INTEGER NOT NULL,
    file_id INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX features_attachments_feature_id ON features_attachments (feature_id);
CREATE INDEX features_attachments_file_id ON features_attachments (file_id);
CREATE INDEX features_attachments_created_by ON features_attachments (created_by);

CREATE TABLE features_comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    feature_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    content_html TEXT NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX features_comments_account_id ON features_comments (account_id);
CREATE INDEX features_comments_feature_id ON features_comments (feature_id);
CREATE INDEX features_comments_created_by ON features_comments (created_by);

CREATE TABLE features_dependencies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    feature_id INTEGER NOT NULL,
    depends_on INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX features_dependencies_feature_id ON features_dependencies (feature_id);
CREATE INDEX features_dependencies_depends_on ON features_dependencies (depends_on);

CREATE TABLE files (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    path TEXT NOT NULL,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    size INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX files_account_id ON files (account_id);
CREATE INDEX files_created_by ON files (created_by);

CREATE TABLE intelligence_assistants (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE TABLE intelligence_batches (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE TABLE intelligence_files (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX intelligence_files_account_id ON intelligence_files (account_id);

CREATE TABLE intelligence_vector_files (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX intelligence_vector_files_account_id ON intelligence_vector_files (account_id);

CREATE TABLE intelligence_vector_stores (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX intelligence_vector_stores_account_id ON intelligence_vector_stores (account_id);

CREATE TABLE invitations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    team_id INTEGER NOT NULL,
    role_id INTEGER NOT NULL,
    email TEXT NOT NULL,
    token TEXT NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    expires_at INTEGER NOT NULL DEFAULT (strftime('%s', datetime('now', '+6 hours'))),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX invitations_email ON invitations (email);
CREATE INDEX invitations_token ON invitations (token);
CREATE INDEX invitations_created_by ON invitations (created_by);
CREATE INDEX invitations_expires_at ON invitations (expires_at);

CREATE TABLE milestones (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    description_html TEXT NOT NULL,
    duration INTEGER NOT NULL,
    ends_on INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX milestones_account_id ON milestones (account_id);
CREATE INDEX milestones_ends_on ON milestones (ends_on);

CREATE TABLE milestones_reports (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    milestone_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    content_html TEXT NOT NULL,
    duration_accuracy INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    expires_at INTEGER NOT NULL DEFAULT (strftime('%s', datetime('now', '+1 year'))),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX milestones_reports_account_id ON milestones_reports (account_id);
CREATE INDEX milestones_reports_milestone_id ON milestones_reports (milestone_id);
CREATE INDEX milestones_reports_expires_at ON milestones_reports (expires_at);

CREATE TABLE price_plans (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE TABLE products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    description_html TEXT NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX products_account_id ON products (account_id);
CREATE INDEX products_created_by ON products (created_by);

CREATE TABLE products_attachments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    product_id INTEGER NOT NULL,
    file_id INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX products_attachments_product_id ON products_attachments (product_id);
CREATE INDEX products_attachments_file_id ON products_attachments (file_id);
CREATE INDEX products_attachments_created_by ON products_attachments (created_by);

CREATE TABLE products_notes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    content_html TEXT NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX products_notes_account_id ON products_notes (account_id);
CREATE INDEX products_notes_product_id ON products_notes (product_id);
CREATE INDEX products_notes_created_by ON products_notes (created_by);

CREATE TABLE registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    price_plan_id INTEGER NOT NULL,
    company_name TEXT NOT NULL,
    company_country_code NOT NULL,
    company_country_name NOT NULL,
    primary_contact_name_first TEXT NOT NULL,
    primary_contact_name_last TEXT NOT NULL,
    primary_contact_email TEXT NOT NULL,
    token TEXT NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    expires_at INTEGER NOT NULL DEFAULT (strftime('%s', datetime('now', '+6 hours'))),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE TABLE tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    feature_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    description_html TEXT NOT NULL,
    priority_name TEXT NOT NULL,
    priority_value INTEGER NOT NULL,
    status_name TEXT NOT NULL,
    status_value INTEGER NOT NULL,
    expected_start_on INTEGER NOT NULL,
    expected_duration INTEGER NOT NULL,
    started_on INTEGER NOT NULL,
    ended_on INTEGER NOT NULL,
    duration INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX tasks_feature_id ON tasks (feature_id);
CREATE INDEX tasks_user_id ON tasks (user_id);
CREATE INDEX tasks_priority_value ON tasks (priority_value);
CREATE INDEX tasks_status_value ON tasks (status_value);
CREATE INDEX tasks_started_on ON tasks (started_on);
CREATE INDEX tasks_ended_on ON tasks (ended_on);
CREATE INDEX tasks_created_by ON tasks (created_by);

CREATE TABLE tasks_attachments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id INTEGER NOT NULL,
    file_id INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX tasks_attachments_task_id ON tasks_attachments (task_id);
CREATE INDEX tasks_attachments_file_id ON tasks_attachments (file_id);
CREATE INDEX tasks_attachments_created_by ON tasks_attachments (created_by);

CREATE TABLE tasks_changes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    task_id INTEGER NOT NULL,
    reason TEXT NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX tasks_changes_account_id ON tasks_changes (account_id);
CREATE INDEX tasks_changes_task_id ON tasks_changes (task_id);
CREATE INDEX tasks_changes_created_by ON tasks_changes (created_by);

CREATE TABLE tasks_comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    task_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    content_html TEXT NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX tasks_acocunt_id ON tasks_comments (account_id);
CREATE INDEX tasks_comments_task_id ON tasks_comments (task_id);
CREATE INDEX tasks_comments_created_by ON tasks_comments (created_by);

CREATE TABLE tasks_dependencies (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id INTEGER NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX tasks_dependencies_task_id ON tasks_dependencies (task_id);
CREATE INDEX tasks_dependencies_created_by ON tasks_dependencies (created_by);

CREATE TABLE tasks_reports (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    task_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    content_html TEXT NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    expires_at INTEGER NOT NULL DEFAULT (strftime('%s', datetime('now', '+1 year'))),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX tasks_reports_account_id ON tasks_reports (account_id);
CREATE INDEX tasks_reports_task_id ON tasks_reports (task_id);
CREATE INDEX tasks_reports_expires_at ON tasks_reports (expires_at);

CREATE TABLE teams (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    created_by INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX teams_account_id ON teams (account_id);
CREATE INDEX teams_created_by ON teams (created_by);

CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    team_id INTEGER NOT NULL,
    name_first TEXT NOT NULL,
    name_last TEXT NOT NULL,
    name_initials TEXT NOT NULL,
    contact_email TEXT NOT NULL,
    contact_phone TEXT NOT NULL,
    role_name TEXT NOT NULL,
    role_value INTEGER NOT NULL,
    password_hash TEXT NOT NULL,
    session_token TEXT NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX users_account_id ON users (account_id);
CREATE INDEX users_team_id ON users (team_id);
CREATE INDEX users_contact_email ON users (contact_email);
CREATE INDEX users_session_token ON users (session_token);

CREATE TABLE users_absences (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    date INTEGER NOT NULL,
    duration INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX users_absences_user_id ON users_absences (user_id);
CREATE INDEX users_absences_date ON users_absences (date);

CREATE TABLE users_availability (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    day INTEGER NOT NULL,
    duration INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX users_availability_user_id ON users_availability (user_id);
CREATE INDEX users_availability_day ON users_availability (day);

CREATE TABLE users_capabilities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    capability_name TEXT NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX users_capabilities_user_id ON users_capabilities (user_id);
CREATE INDEX users_capabilities_capability_name ON users_capabilities (capability_name);

CREATE TABLE users_reports (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    account_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    expires_at INTEGER NOT NULL DEFAULT (strftime('%s', datetime('now', '+1 year'))),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX users_reports_account_id ON users_reports (account_id);
CREATE INDEX users_reports_user_id ON users_reports (user_id);
CREATE INDEX users_reports_expires_at ON users_reports (expires_at);

CREATE TABLE users_settings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    created_at INTEGER NOT NULL DEFAULT (strftime('%s','now')),
    updated_at INTEGER NOT NULL DEFAULT (strftime('%s','now'))
);

CREATE INDEX users_settings_user_id ON users_settings (user_id);
