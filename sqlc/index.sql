CREATE INDEX idx_parent_comments_time
ON ParentComments(PageId, CreatedAt);

CREATE INDEX idx_parent_comments_userid
ON ParentComments(PageId, UserId);

CREATE INDEX idx_child_comments
ON ChildComments(PageId, ParentCommentId);

CREATE INDEX idx_like_history
ON ChildComments(PageId, UserId);