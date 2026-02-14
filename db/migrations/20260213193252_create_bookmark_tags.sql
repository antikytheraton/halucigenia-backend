-- migrate:up

CREATE TABLE bookmark_tags (
    bookmark_id uuid NOT NULL,
    tag_id uuid NOT NULL,
    PRIMARY KEY (bookmark_id, tag_id),
    FOREIGN KEY (bookmark_id) REFERENCES bookmarks(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

-- migrate:down

DROP TABLE IF EXISTS bookmark_tags;
