
function cropPrintMark(pageSelector, size = '0.3cm', color = 'black',
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
  addMarks('page[sizeX="{{.SinglePage.Spec.Canvas.Width}}"][sizeY="{{.SinglePage.Spec.Canvas.Height}}"]', '0.3cm', 'black')
};

function print(sizeX, sizeY){
  const { jsPDF } = window.page.jspdf;

  const doc = new jsPDF();

  // Capture the content of the current page or a specific element
  const content = document.body; // or document.getElementById('yourElementId');

  // Use the html method to add content to the PDF
  doc.html(content, {
    callback: function (doc) {
      // Save the PDF and then print
      const pdfOutput = doc.output('blob');
      const pdfUrl = URL.createObjectURL(pdfOutput);

      const printWindow = window.open(pdfUrl);
      printWindow.onload = function() {
        printWindow.print();
      };
    },
    x: sizeX,
    y: sizeY,
  });


  document.getElementById("download").addEventListener("click", function() {
    // Get jsPDF instance
    const { jsPDF } = window.page.jspdf;
    const pdf = new jsPDF({
      orientation: "p", // Portrait
      unit: "cm",       // Use centimeters
      format: [8, 8],   // Custom format 8x8 cm
      putOnlyUsedFonts: true,
      floatPrecision: 16 // Float precision for jsPDF
    });

    // Set background color
    // pdf.setFillColor(252, 51, 0); // Equivalent to #fc3
    // pdf.rect(0, 0, 8, 8, 'F'); // Fill the entire page with the background color

    // Set padding
    const padding = 0.75;
    const contentX = padding; // X position inside the PDF
    const contentY = padding; // Y position

    // Add content to the PDF
    // Save the PDF
    pdf.save("document_with_custom_marks.pdf");
  });
}


function printCurrentPage() {
  const { jsPDF } = window.jspdf;

  const doc = new jsPDF({
    orientation: 'landscape', // Change to 'portrait' if needed
    unit: 'cm',
    format: [15, 11], // Width: 15 cm, Height: 11 cm
    putOnlyUsedFonts: true,
    floatPrecision: 16 // Optional, adjust for precision
  });

  const content = document.body; // or document.getElementById('yourElementId');

  // Use html2canvas to convert HTML to canvas
  x = html2canvas(content, {
    scale: 2, // Increase scale for better quality
    useCORS: false // Enable if loading images from another domain
  })

    const imgData = canvas.toDataURL('image/png');
    const imgWidth = 15; // Width in cm
    const pageHeight = 11; // Height in cm
    const imgHeight = (canvas.height * imgWidth) / canvas.width;
    let heightLeft = imgHeight;

    let position = 0;

    // Add the first image to the PDF
    doc.addImage(imgData, 'PNG', 0, position, imgWidth, imgHeight);
    heightLeft -= pageHeight;

    // Add additional pages if needed
    while (heightLeft >= 0) {
      position = heightLeft - imgHeight;
      doc.addPage();
      doc.addImage(imgData, 'PNG', 0, position, imgWidth, imgHeight);
      heightLeft -= pageHeight;
    }

    // Save the PDF
    doc.save('currentPage.pdf');

    // Optionally, open print dialog
    const pdfOutput = doc.output('blob');
    const pdfUrl = URL.createObjectURL(pdfOutput);
    const printWindow = window.open(pdfUrl);
    printWindow.onload = function() {
      printWindow.print();
    };

}
