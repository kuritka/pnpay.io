
function cropPrintMark(pageSelector, size = '"{{.SinglePage.Spec.Marks.Size}}"', color = 'black',
                       { top = '', right = '', bottom = '', left = ''}={},
                       { borderTop = false, borderRight = false, borderBottom = false, borderLeft = false}={}
) {
  const page = document.querySelector(pageSelector);                // Select the page element
  const mark = document.createElement('div');
  const border = "{{.SinglePage.Spec.Marks.Border}}"
  mark.style.position = 'absolute';
  if (top !== '') {
    mark.style.top = top;
  }
  if (right !== '') {
    mark.style.right = right;
  }
  if (bottom !== '') {
    mark.style.bottom = bottom;
  }
  if (left !== '') {
    mark.style.left = left;
  }
  mark.style.width = size;                                              // Set mark width
  mark.style.height = size;                                             // Set mark height
  mark.style.background = 'transparent';                                // Transparent background
  if (borderTop) {
    mark.style.borderTop = border;                                      // Set left border width
  }
  if (borderLeft) {
    mark.style.borderLeft = border;                                     // Set left border width
  }
  if (borderRight) {
    mark.style.borderRight = border;                                    // Set left border width
  }
  if (borderBottom) {
    mark.style.borderBottom = border;                                   // Set left border width
  }
  page.appendChild(mark);
}

function addMarks(pageSelector, size, color) {
  const minus  = '-'+size
  cropPrintMark(pageSelector, size, color,  {top: minus,  right:'0'}, {borderLeft: true});
  cropPrintMark(pageSelector, size, color,  {top:'-0',  right: minus}, {borderBottom: true});
  cropPrintMark(pageSelector, size, color,  {top: minus,  left:'0'}, {borderRight: true});
  cropPrintMark(pageSelector, size, color,  {top:'0',  left: minus}, {borderBottom: true});
  cropPrintMark(pageSelector, size, color,  {bottom: minus,  right:'0'}, {borderLeft: true});
  cropPrintMark(pageSelector, size, color,  {bottom:'0',  right: minus}, {borderTop: true});
  cropPrintMark(pageSelector, size, color,  {bottom:'0',  left: minus}, {borderTop: true});
  cropPrintMark(pageSelector, size, color,  {bottom: minus,  left:'0'}, {borderRight: true});
}

// Call the function with the target page selector and size for the marks
window.onload = function() {
  addMarks('page[sizeX="{{.SinglePage.Spec.Canvas.Width}}"][sizeY="{{.SinglePage.Spec.Canvas.Height}}"]', '{{.SinglePage.Spec.Canvas.ExtPadding}}', 'black')
};
