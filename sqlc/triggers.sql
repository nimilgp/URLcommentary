CREATE OR REPLACE FUNCTION set_commentid_parentcomments()
RETURNS TRIGGER AS $$
BEGIN
    SELECT COALESCE(MAX(CommentId), 0) + 1 INTO NEW.CommentId
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
    SELECT COALESCE(MAX(ChildCommentId), 0) + 1 INTO NEW.ChildCommentId
    FROM ChildComments
    WHERE Pageid = NEW.Pageid;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER child_comments_insert_trigger
BEFORE INSERT ON ChildComments
FOR EACH ROW
EXECUTE FUNCTION set_commentid_childcomments();

--  increment Page's comment count on insert in ParentComments
CREATE OR REPLACE FUNCTION increment_page_comments_count()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE Pages
  SET CommentsCount = CommentsCount + 1
  WHERE PageId = NEW.PageId;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER increment_page_comments_count_trigger
AFTER INSERT ON ParentComments
FOR EACH ROW
EXECUTE FUNCTION increment_page_comments_count();

--  increment ParentComment's child comment count on insert in ChildComments
CREATE OR REPLACE FUNCTION increment_parent_comment_child_count()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE ParentComments
  SET ChildCommentCount = ChildCommentCount + 1
  WHERE PageId = NEW.PageId AND CommentId = NEW.ParentCommentId;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER increment_parent_comment_child_count_trigger
AFTER INSERT ON ChildComments
FOR EACH ROW
EXECUTE FUNCTION increment_parent_comment_child_count();