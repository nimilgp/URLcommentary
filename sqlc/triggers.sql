CREATE OR REPLACE FUNCTION set_commentid_parentcomments()
RETURNS TRIGGER AS $$
BEGIN
    SELECT COALESCE(MAX(CommentId), -1) + 1 INTO NEW.CommentId
    FROM ParentComments
    WHERE Pageid = NEW.Pageid;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER parent_comments_insert_trigger
BEFORE INSERT ON ParentComments
FOR EACH ROW
EXECUTE FUNCTION set_commentid_parentcomments();

CREATE OR REPLACE FUNCTION set_commentid_childcomments()
RETURNS TRIGGER AS $$
BEGIN
    SELECT COALESCE(MAX(ChildCommentId), -1) + 1 INTO NEW.ChildCommentId
    FROM ChildComments
    WHERE Pageid = NEW.Pageid;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER child_comments_insert_trigger
BEFORE INSERT ON ChildComments
FOR EACH ROW
EXECUTE FUNCTION set_commentid_childcomments();